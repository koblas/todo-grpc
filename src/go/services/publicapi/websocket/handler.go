package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"go.uber.org/zap"
)

var (
	contentType     = map[string]string{"Content-Type": "application/json"}
	responseBadAuth = events.APIGatewayProxyResponse{StatusCode: http.StatusNonAuthoritativeInfo, Body: "{}", Headers: contentType}
	responseBad     = events.APIGatewayProxyResponse{StatusCode: http.StatusBadGateway, Body: "{}", Headers: contentType}
	responseOk      = events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: "{}", Headers: contentType}
)

type WebsocketHandler struct {
	jwtMaker tokenmanager.Maker
	store    websocket.ConnectionStore
}

var (
	ErrorMissingAuthentication = errors.New("bad authentication token")
	ErrorInvalidToken          = errors.New("token is invalid")
)

type Option func(*WebsocketHandler)

func WithStore(store websocket.ConnectionStore) Option {
	return func(conf *WebsocketHandler) {
		conf.store = store
	}
}

func NewWebsocketHandler(jwtSecret string, opts ...Option) *WebsocketHandler {
	maker, err := tokenmanager.NewJWTMaker(jwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	conf := WebsocketHandler{
		jwtMaker: maker,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	return &conf
}

func (svc *WebsocketHandler) isAuthenticated(ctx context.Context, log logger.Logger, req events.APIGatewayWebsocketProxyRequest) (string, error) {
	token := req.QueryStringParameters["t"]
	if token == "" {
		if len(req.Body) == 0 {
			return "", ErrorMissingAuthentication
		}

		// Check the message body
		msg := struct {
			Action *string
			Token  *string
		}{}

		if err := json.Unmarshal([]byte(req.Body), &msg); err != nil {
			log.With("error", err).Info("unmarshal message failed")
			return "", ErrorMissingAuthentication
		}

		if msg.Action == nil || msg.Token == nil || *msg.Action == "" || *msg.Token == "" {
			return "", ErrorMissingAuthentication
		}

		if *msg.Action != "authorization" {
			return "", ErrorMissingAuthentication
		}
		token = *msg.Token
	}

	payload, err := svc.jwtMaker.VerifyToken(token)
	if err != nil {
		return "", ErrorInvalidToken
	}

	if err := svc.store.Create(ctx, payload.UserId, req.RequestContext.ConnectionID); err != nil {
		log.With("error", err).Info("DB Create failed")
		return "", err
	}

	return payload.UserId, nil
}

func (svc *WebsocketHandler) HandleRequest(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.FromContext(ctx).With("eventType", req.RequestContext.EventType)

	if req.RequestContext.EventType == "CONNECT" {
		log.Info("Connect event")
		_, err := svc.isAuthenticated(ctx, log, req)
		if err != nil && err != ErrorMissingAuthentication {
			log.With("error", err).Info("Authentication error")
			return responseBadAuth, nil
		}
	} else if req.RequestContext.EventType == "DISCONNECT" {
		log.Info("Disconnect event")
		if err := svc.store.Delete(ctx, req.RequestContext.ConnectionID); err != nil {
			log.With("error", err).Info("DB Delete failed")
			return responseBad, err
		}
	} else if req.RequestContext.EventType == "MESSAGE" {
		log.Info("Message event")

		userId, err := svc.isAuthenticated(ctx, log, req)
		if err == ErrorMissingAuthentication {
			// Missing Authentication implies we're already authenticated and we need to update
			// the DB that the connection is alive
			if err := svc.store.Heartbeat(ctx, req.RequestContext.ConnectionID); err != nil {
				log.With("error", err).Info("DB heartbeat failed")
				return responseBad, err
			}
		} else if err != nil {
			log.With(zap.Error(err)).Info("Authentication failed")
			return responseBad, err
		} else if userId != "" {
			log.With(zap.String("userId", userId)).Info("Authentication success")
		}
	} else {
		log.Error("Unknown event type")
		return responseBad, nil
	}

	return responseOk, nil

}

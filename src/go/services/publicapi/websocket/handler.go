package websocket

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
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

type Option func(*WebsocketHandler)

func WithStore(store websocket.ConnectionStore) Option {
	return func(conf *WebsocketHandler) {
		conf.store = store
	}
}

func NewWebsocketHandler(config SsmConfig, opts ...Option) *WebsocketHandler {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
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

func (svc *WebsocketHandler) isAuthenticated(req events.APIGatewayWebsocketProxyRequest) (string, error) {
	token := req.QueryStringParameters["t"]

	if token == "" {
		return "", errors.New("missing authorization token")
	}

	payload, err := svc.jwtMaker.VerifyToken(token)
	if err != nil {
		return "", errors.New("token didn't validate")
	}

	return payload.UserId, nil
}

func (svc *WebsocketHandler) HandleRequest(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.FromContext(ctx).With("eventType", req.RequestContext.EventType)

	if req.RequestContext.EventType == "CONNECT" {
		log.Info("Connect event")
		userId, err := svc.isAuthenticated(req)
		if err != nil {
			log.With("error", err).Info("Authentication error")
			return responseBadAuth, nil
		}
		if err := svc.store.Create(userId, req.RequestContext.ConnectionID); err != nil {
			log.With("error", err).Info("DB Create failed")
			return responseBad, err
		}
	} else if req.RequestContext.EventType == "DISCONNECT" {
		log.Info("Disconnect event")
		if err := svc.store.Delete(req.RequestContext.ConnectionID); err != nil {
			log.With("error", err).Info("DB Delete failed")
			return responseBad, err
		}
	} else if req.RequestContext.EventType == "MESSAGE" {
		log.Info("Message event")
		if err := svc.store.Heartbeat(req.RequestContext.ConnectionID); err != nil {
			log.With("error", err).Info("DB heartbeaat failed")
			return responseBad, err
		}
		// return responseBad, nil
	} else {
		log.Error("Unknown event type")
		return responseBad, nil
	}

	return responseOk, nil

}

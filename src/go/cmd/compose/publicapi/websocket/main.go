package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	wsocket "github.com/gorilla/websocket"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/eventbus"
	redisbus "github.com/koblas/grpc-todo/pkg/eventbus/redis"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	wstore "github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/publicapi/websocket"
	"go.uber.org/zap"
)

var upgrader = wsocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type socketHandler struct {
	connections map[string]*wsocket.Conn
	consumer    eventbus.SimpleConsumer
	api         *websocket.WebsocketHandler
}

func (h *socketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	connectionID := uuid.NewString()
	log := logger.FromContext(r.Context()).With(zap.String("connectionId", connectionID))

	log.Info("Handling request")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.With(zap.Error(err)).Error("failed to upgrade")
		return
	}
	defer conn.Close()

	queryString := map[string]string{}
	for key, value := range r.URL.Query() {
		queryString[key] = value[0]
	}

	wsRes, err := h.api.HandleRequest(r.Context(), events.APIGatewayWebsocketProxyRequest{
		RequestContext: events.APIGatewayWebsocketProxyRequestContext{
			ConnectionID: connectionID,
			EventType:    "CONNECT",
		},
		QueryStringParameters: queryString,
	})
	if err != nil || wsRes.StatusCode != http.StatusOK {
		log.With(zap.Error(err)).Error("failed to connect")
		return
	}
	h.connections[connectionID] = conn
	defer delete(h.connections, connectionID)

	defer func() {
		_, err := h.api.HandleRequest(r.Context(), events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{
				ConnectionID: connectionID,
				EventType:    "DISCONNECT",
			},
			QueryStringParameters: queryString,
		})
		if err != nil {
			log.With(zap.Error(err)).Error("failed to disconnect")
		}
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if messageType != wsocket.CloseMessage {
				log.With(zap.Error(err)).Error("ReadMessage failed")
			}
			break
		}

		_, err = h.api.HandleRequest(r.Context(), events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{
				ConnectionID: connectionID,
				EventType:    "MESSAGE",
			},
			QueryStringParameters: queryString,
			Body:                  string(p),
		})
		if err != nil {
			log.With(zap.Error(err)).Error("failed to send message")
		}
	}

}

func (h *socketHandler) consume(ctx context.Context) {
	log := logger.FromContext(ctx)

	for {
		msg, err := h.consumer.Next()
		if err != nil {
			log.With(zap.Error(err)).Info("message consumer failed")
			return
		}

		apimsg := awsutil.ConvertEventbusToApiGateway(&msg)

		conn, ok := h.connections[*apimsg.ConnectionId]
		if !ok {
			continue
		}

		if err := conn.WriteMessage(wsocket.TextMessage, []byte(apimsg.Data)); err != nil {
			log.With(zap.Error(err)).Info("message consumer failed")
			return
		}
	}
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := websocket.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	consumer := redisbus.NewConsumer(mgr.Context(), ssmConfig.RedisAddr, ssmConfig.WebsocketBroadcast)

	handler := socketHandler{
		api: websocket.NewWebsocketHandler(
			ssmConfig,
			websocket.WithStore(wstore.NewRedisStore(ssmConfig.RedisAddr)),
		),
		consumer:    consumer,
		connections: map[string]*wsocket.Conn{},
	}

	go handler.consume(mgr.Context())

	mgr.Start(&handler)
}

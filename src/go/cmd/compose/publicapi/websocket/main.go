package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
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
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret                  string `validate:"min=32"`
	RedisAddr                  string `default:"redis:6379"`
	WebsocketConnectionMessage string
}

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
	connectionID := ulid.Make().String()
	log := logger.FromContext(r.Context()).With(zap.String("connectionId", connectionID))

	log.Info("Handling request")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.With(zap.Error(err)).Error("failed to upgrade")
		return
	}
	defer conn.Close()

	log.Info("websocket connection created")

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
		log.With(zap.Error(err)).Info("failed to connect")
		return
	}

	// Let's make sure we register the cleanup handling
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

	h.connections[connectionID] = conn
	defer delete(h.connections, connectionID)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if messageType != wsocket.CloseMessage {
				log.With(zap.Error(err)).Info("ReadMessage failed")
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
			log.With(zap.Error(err)).Error("failed to receive message")
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

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	consumer := redisbus.NewConsumer(mgr.Context(), config.RedisAddr, config.WebsocketConnectionMessage)

	handler := socketHandler{
		api: websocket.NewWebsocketHandler(
			config.JwtSecret,
			websocket.WithStore(wstore.NewRedisStore(config.RedisAddr)),
		),
		consumer:    consumer,
		connections: map[string]*wsocket.Conn{},
	}

	go handler.consume(mgr.Context())

	mux := http.NewServeMux()
	mux.Handle("/", &handler)
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

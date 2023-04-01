package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/websocket/broadcast"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr                   string
	RedisAddr                  string
	WebsocketConnectionMessage string
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	client := redisutil.RedisWsClient(mgr.Context(), config.RedisAddr, config.WebsocketConnectionMessage)

	s := broadcast.NewBroadcastServer(
		broadcast.WithStore(websocket.NewRedisStore(config.RedisAddr)),
		broadcast.WithClient(client),
	)
	nats := natsutil.NewNatsClient(config.NatsAddr)

	mgr.Start(nats.TopicConsumer(mgr.Context(),
		natsutil.ConnectToTopic(eventbusv1connect.BroadcastEventbusServiceName),
		s,
	))
}

package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr string
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	producer := eventbusv1connect.NewBroadcastEventbusServiceClient(nats, "")

	s := todo.NewTodoChangeServer(
		todo.WithProducer(producer),
	)

	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.ConnectToTopic(eventbusv1connect.TodoEventbusServiceName),
		s))
}

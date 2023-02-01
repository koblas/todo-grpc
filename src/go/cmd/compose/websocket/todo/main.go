package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := todo.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	producer := corepbv1.NewBroadcastEventbusProtobufClient("", nats)

	s := todo.NewTodoChangeServer(
		todo.WithProducer(producer),
	)

	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepbv1.TodoEventbusPathPrefix),
		s))
}

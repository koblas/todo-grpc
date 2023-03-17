package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	eventbus := corev1connect.NewTodoEventbusServiceClient(
		nats,
		"topic://"+config.TodoEventsTopic,
	)

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
	}

	_, api := corev1connect.NewTodoServiceHandler(todo.NewTodoServer(opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

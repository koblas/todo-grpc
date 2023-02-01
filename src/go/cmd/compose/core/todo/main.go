package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
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
	eventbus := corepbv1.NewTodoEventbusJSONClient(
		"topic://"+config.TodoEventsTopic,
		nats,
	)

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
	}

	api := corepbv1.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

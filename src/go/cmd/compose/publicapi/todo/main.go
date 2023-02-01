package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []todo.Option{
		todo.WithTodoService(
			corepbv1.NewTodoServiceProtobufClient(
				"http://"+config.TodoServiceAddr,
				&http.Client{},
			),
		),
	}

	api := apipbv1.NewTodoServiceServer(todo.NewTodoServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

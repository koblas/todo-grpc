package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig todo.SsmConfig
	if err := awsutil.LoadEnvConfig("/common/", &ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("unable to load configuration")
	}

	opts := []todo.Option{
		todo.WithTodoService(
			core.NewTodoServiceProtobufClient(
				"http://"+util.Getenv("TODO_SERVICE_ADDR", ":13005"),
				&http.Client{},
			),
		),
	}

	api := publicapi.NewTodoServiceServer(todo.NewTodoServer(ssmConfig, opts...))

	mgr.Start(api)
}

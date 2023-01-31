package main

import (
	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []todo.Option{
		todo.WithTodoService(corepb.NewTodoServiceJSONClient("lambda://core-todo", awsutil.NewTwirpCallLambda())),
	}

	api := apipb.NewTodoServiceServer(todo.NewTodoServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

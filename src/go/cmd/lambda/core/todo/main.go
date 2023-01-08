package main

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// eventbus := corepb.NewTodoEventbusProtobufClient(config.EventArn, awsutil.NewTwirpCallLambda())
	producer := corepb.NewTodoEventbusJSONClient(
		config.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	opts := []todo.Option{
		todo.WithProducer(producer),
		todo.WithTodoStore(todo.NewTodoDynamoStore()),
	}

	api := corepb.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(api)
}

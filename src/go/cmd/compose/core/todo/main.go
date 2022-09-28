package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig todo.SsmConfig
	if err := awsutil.LoadEnvConfig("/common/", &ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	eventbus := core.NewTodoEventbusProtobufClient(ssmConfig.EventArn, awsutil.NewTwirpCallLambda())

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoDynamoStore()),
	}

	api := core.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(api)
}

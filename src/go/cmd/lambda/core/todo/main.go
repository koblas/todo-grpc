package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
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
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	log.With(zap.String("eventArn", config.EventArn)).Info("Constructing producer")
	// eventbus := corepbv1.NewTodoEventbusProtobufClient(config.EventArn, awsutil.NewTwirpCallLambda())
	producer := corepbv1.NewTodoEventbusServiceJSONClient(
		config.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	opts := []todo.Option{
		todo.WithProducer(producer),
		todo.WithTodoStore(todo.NewTodoDynamoStore()),
	}

	api := corepbv1.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

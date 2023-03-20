package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
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
	producer := corev1connect.NewTodoEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.EventArn,
	)

	opts := []todo.Option{
		todo.WithProducer(producer),
		todo.WithTodoStore(todo.NewTodoDynamoStore()),
	}

	_, api := corev1connect.NewTodoServiceHandler(
		todo.NewTodoServer(opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

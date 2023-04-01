package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/todo"
	"go.uber.org/zap"
)

type Config struct {
	BusEntityArn string `validate:"required"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	cloader := confmgr.NewLoader(
		confmgr.NewLoaderEnvironment("", "_"),
		aws.NewLoaderSsm(mgr.Context(), "/common/"),
	)
	if err := cloader.Parse(mgr.Context(), &config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := eventbusv1connect.NewTodoEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.BusEntityArn,
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

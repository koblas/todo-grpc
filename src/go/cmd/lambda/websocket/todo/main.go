package main

import (
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"go.uber.org/zap"
)

type Config struct {
	BusEntityArn string `validate:"required"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := Config{}
	if err := confmgr.Parse(&conf, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := eventbusv1connect.NewBroadcastEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		conf.BusEntityArn,
	)

	h := todo.NewTodoChangeServer(todo.WithProducer(producer))

	mgr.Start(awsutil.HandleSqsLambda(h))
}

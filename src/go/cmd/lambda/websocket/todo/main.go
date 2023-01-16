package main

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := todo.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepb.NewBroadcastEventbusJSONClient(
		conf.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	h := todo.NewTodoChangeServer(todo.WithProducer(producer))

	mgr.Start(awsutil.HandleSqsLambda(h))
}

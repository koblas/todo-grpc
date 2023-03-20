package main

import (
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/websocket/user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := user.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewBroadcastEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		conf.EventArn,
	)

	// s := user.NewUserChangeServer(
	// 	user.WithProducer(producer),
	// )
	// mux := http.NewServeMux()
	// mux.Handle(corepbv1.UserEventbusPathPrefix, corepbv1.NewUserEventbusServer(s))

	handlers := user.NewUserChangeServer(user.WithProducer(producer))

	mgr.Start(awsutil.HandleSqsLambda(handlers))
}

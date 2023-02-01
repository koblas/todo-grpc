package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/websocket/user"
	"go.uber.org/zap"
)

type handler struct {
	bus corepbv1.TwirpServer
}

func (h handler) GroupName() string {
	return ""
}
func (h handler) Handler() corepbv1.TwirpServer {
	return h.bus
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := user.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepbv1.NewBroadcastEventbusJSONClient(
		conf.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	// s := user.NewUserChangeServer(
	// 	user.WithProducer(producer),
	// )
	// mux := http.NewServeMux()
	// mux.Handle(corepbv1.UserEventbusPathPrefix, corepbv1.NewUserEventbusServer(s))

	h := user.NewUserChangeServer(user.WithProducer(producer))

	mgr.Start(awsutil.HandleSqsLambda(h))
}

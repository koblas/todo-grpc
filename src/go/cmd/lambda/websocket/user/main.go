package main

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/websocket/user"
	"go.uber.org/zap"
)

type handler struct {
	bus corepb.TwirpServer
}

func (h handler) GroupName() string {
	return ""
}
func (h handler) Handler() corepb.TwirpServer {
	return h.bus
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := user.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepb.NewBroadcastEventbusJSONClient(
		conf.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	// s := user.NewUserChangeServer(
	// 	user.WithProducer(producer),
	// )
	// mux := http.NewServeMux()
	// mux.Handle(corepb.UserEventbusPathPrefix, corepb.NewUserEventbusServer(s))

	h := user.NewUserChangeServer(user.WithProducer(producer))

	mgr.Start(awsutil.HandleSqsLambda(h))
}

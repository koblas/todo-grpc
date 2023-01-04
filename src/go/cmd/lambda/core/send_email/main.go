package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := send_email.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := core.NewSendEmailEventsProtobufClient(
		config.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(config))
	mux := http.NewServeMux()
	mux.Handle(core.SendEmailServicePathPrefix, core.NewSendEmailServiceServer(s))

	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
}

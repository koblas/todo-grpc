package main

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := send_email.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepb.NewSendEmailEventsProtobufClient(
		config.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(config))
	// mux := http.NewServeMux()
	// mux.Handle(corepb.SendEmailServicePathPrefix, corepb.NewSendEmailServiceServer(s))

	mgr.Start(awsutil.HandleSqsLambda(s))
}

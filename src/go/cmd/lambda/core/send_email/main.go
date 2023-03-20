package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := send_email.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}
	smtpConfig := send_email.SmtpConfig{}
	if err := confmgr.Parse(&smtpConfig, aws.NewLoaderSsm(mgr.Context(), "/smtp/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewSendEmailEventsServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.EventArn,
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(smtpConfig))

	mgr.Start(awsutil.HandleSqsLambda(s))
}

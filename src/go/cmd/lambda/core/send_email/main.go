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

type Config struct {
	BusEntityArn string `ssm:"/common/bus_entity_arn" validate:"required"`
	Smtp         struct {
		Addr     string
		Username string
		Password string
	}
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	cloader := confmgr.NewLoader(
		confmgr.NewLoaderEnvironment("", "_"),
		aws.NewLoaderSsm(mgr.Context(), ""),
	)
	if err := cloader.Parse(mgr.Context(), &config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewSendEmailEventsServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.BusEntityArn,
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
	)

	s := send_email.NewSendEmailServer(producer,
		send_email.NewSmtpService(
			config.Smtp.Addr,
			config.Smtp.Username,
			config.Smtp.Password,
		),
	)

	mgr.Start(awsutil.HandleSqsLambda(s))
}

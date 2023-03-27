package main

import (
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	workers "github.com/koblas/grpc-todo/services/workers/workers_user"
	"go.uber.org/zap"
)

type Config struct {
	UrlBase string
}

func main() {
	mgr := manager.NewManager()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		mgr.Logger().With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []workers.Option{
		workers.WithSendEmail(
			corev1connect.NewSendEmailServiceClient(
				awsutil.NewTwirpCallLambda(),
				"sqs://send-email",
			),
		),
		workers.WithUrlBase(config.UrlBase),
	}

	mgr.Start(awsutil.HandleSqsLambda(workers.GetHandler(opts...)))
}

package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	workers "github.com/koblas/grpc-todo/services/workers/workers_user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()

	config := workers.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		mgr.Logger().With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []workers.Option{
		workers.WithSendEmail(
			corepbv1.NewSendEmailServiceProtobufClient(
				"sqs://send-email",
				awsutil.NewTwirpCallLambda(),
			),
		),
	}

	mgr.Start(awsutil.HandleSqsLambda(workers.GetHandler(config, opts...)))
}

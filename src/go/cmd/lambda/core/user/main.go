package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config user.Config
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepbv1.NewUserEventbusJSONClient(
		config.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	opts := []user.Option{
		user.WithProducer(producer),
		user.WithUserStore(user.NewUserDynamoStore()),
	}

	api := corepbv1.NewUserServiceServer(user.NewUserServer(opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

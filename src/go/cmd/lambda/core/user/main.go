package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/user"
	"go.uber.org/zap"
)

type Config struct {
	BusEntityArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	cloader := confmgr.NewLoader(
		confmgr.NewLoaderEnvironment("", "_"),
		aws.NewLoaderSsm(mgr.Context(), "/common/"),
	)
	if err := cloader.Parse(mgr.Context(), &config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewUserEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.BusEntityArn,
	)

	opts := []user.Option{
		user.WithProducer(producer),
		user.WithUserStore(user.NewUserDynamoStore()),
	}

	_, api := corev1connect.NewUserServiceHandler(
		user.NewUserServer(opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

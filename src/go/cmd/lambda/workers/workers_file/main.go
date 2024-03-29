package main

import (
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Config struct {
	BusEntityArn  string `validate:"required"`
	PublicBucket  string `validate:"required"`
	PrivateBucket string `validate:"required"`
}

func main() {
	mgr := manager.NewManager()
	defer mgr.Shutdown()

	config := Config{}

	var opts []workers_file.Option
	{
		ctx, span := otel.Tracer("test").Start(mgr.Context(), "initialize")
		defer span.End()
		if err := confmgr.ParseWithContext(ctx, &config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(ctx, "/common/")); err != nil {
			mgr.Logger().With(zap.Error(err)).Fatal("failed to load configuration")
		}

		client := awsutil.NewTwirpCallLambda()
		opts = []workers_file.Option{
			workers_file.WithProducer(eventbusv1connect.NewFileEventbusServiceClient(
				client,
				config.BusEntityArn,
			)),
			workers_file.WithFileService(filestore.NewAwsProvider()),
			workers_file.WithUserService(userv1connect.NewUserServiceClient(
				client,
				"lambda://core-user",
			)),
			workers_file.WithPublicBucket(config.PublicBucket),
		}
	}

	mgr.Start(awsutil.HandleSqsLambda(workers_file.BuildHandlers(opts...)))
}

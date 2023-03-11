package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	defer mgr.Shutdown()

	var config workers_file.Config
	var opts []workers_file.Option

	{
		ctx, span := otel.Tracer("test").Start(mgr.Context(), "initialize")
		defer span.End()
		if err := confmgr.ParseWithContext(ctx, &config, aws.NewLoaderSsm(ctx, "/common/")); err != nil {
			mgr.Logger().With(zap.Error(err)).Fatal("failed to load configuration")
		}

		client := awsutil.NewTwirpCallLambda()
		opts = []workers_file.Option{
			workers_file.WithProducer(corepbv1.NewFileEventbusJSONClient(
				config.EventArn,
				client,
			)),
			// workers_file.WithFileService(corepbv1.NewFileServiceJSONClient("lambda://core-file", client)),
			workers_file.WithFileService(filestore.NewAwsProvider()),
			workers_file.WithUserService(corepbv1.NewUserServiceJSONClient("lambda://core-user", client)),
		}
	}

	mgr.Start(awsutil.HandleSqsLambda(workers_file.BuildHandlers(config, opts...)))
}

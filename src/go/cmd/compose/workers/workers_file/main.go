package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr        string
	RedisAddr       string
	MinioEndpoint   string
	UserServiceAddr string
	PublicBucket    string `validate:"required"`
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	defer mgr.Shutdown()

	config := Config{
		MinioEndpoint: "s3.amazonaws.com",
	}

	var opts []workers_file.Option
	var nats *natsutil.Client

	{
		ctx, span := otel.Tracer("test").Start(mgr.Context(), "fishsh")
		defer span.End()
		log := mgr.Logger()

		if err := confmgr.ParseWithContext(ctx, &config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
			log.With(zap.Error(err)).Fatal("failed to load configuration")
		}

		if config.RedisAddr == "" {
			log.Fatal("redis address is missing")
		}

		nats = natsutil.NewNatsClient(config.NatsAddr)
		producer := corev1connect.NewFileEventbusServiceClient(nats, "")

		opts = []workers_file.Option{
			workers_file.WithProducer(producer),
			workers_file.WithFileService(
				filestore.NewMinioProvider(config.MinioEndpoint),
			),
			workers_file.WithUserService(
				corev1connect.NewUserServiceClient(
					bufcutil.NewHttpClient(),
					"http://"+config.UserServiceAddr,
				),
			),
			workers_file.WithPublicBucket(config.PublicBucket),
		}

	}

	mgr.Start(nats.TopicConsumer(mgr.Context(),
		natsutil.ConnectToTopic(corev1connect.FileEventbusServiceName),
		workers_file.BuildHandlers(opts...)))
}

package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	defer mgr.Shutdown()

	var config workers_file.Config
	var opts []workers_file.Option
	var nats *natsutil.Client

	{
		ctx, span := otel.Tracer("test").Start(mgr.Context(), "fishsh")
		defer span.End()
		log := mgr.Logger()

		if err := confmgr.ParseWithContext(ctx, &config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
			log.With(zap.Error(err)).Fatal("failed to load configuration")
		}

		if config.RedisAddr == "" {
			log.Fatal("redis address is missing")
		}

		nats = natsutil.NewNatsClient(config.NatsAddr)
		producer := corepbv1.NewFileEventbusProtobufClient("", nats)

		opts = []workers_file.Option{
			workers_file.WithProducer(producer),
			workers_file.WithFileService(
				filestore.NewMinioProvider(config.MinioEndpoint),
			),
			workers_file.WithUserService(
				corepbv1.NewUserServiceProtobufClient(
					"http://"+config.UserServiceAddr,
					&http.Client{},
				),
			),
		}

	}

	mgr.Start(nats.TopicConsumer(mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepbv1.FileEventbusPathPrefix),
		workers_file.BuildHandlers(config, opts...)))
}

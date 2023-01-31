package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := workers_file.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	if config.RedisAddr == "" {
		log.Fatal("redis address is missing")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	producer := corepb.NewFileEventbusProtobufClient("", nats)

	opts := []workers_file.Option{
		workers_file.WithProducer(producer),
		workers_file.WithFileService(
			filestore.NewMinioProvider(config.MinioEndpoint),
		),
		workers_file.WithUserService(
			corepb.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
	}

	handlers := workers_file.BuildHandlers(config, opts...)

	mgr.Start(nats.TopicConsumer(mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepb.FileEventbusPathPrefix),
		handlers))
}

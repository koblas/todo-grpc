package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/file"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := file.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	if config.RedisAddr == "" {
		log.Fatal("redis address is missing")
	}

	producer := corepb.NewFileEventbusProtobufClient(
		"",
		natsutil.NewNatsClient(config.NatsAddr),
	)

	opts := []file.Option{
		file.WithProducer(producer),
	}

	log.Info("Starting up with Memory store")
	opts = append(opts, file.WithFileStore(file.NewFileMemoryStore("/api/v1/fileput/")))

	api := corepb.NewFileServiceServer(file.NewFileServer(opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

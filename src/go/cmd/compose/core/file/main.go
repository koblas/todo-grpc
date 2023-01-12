package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/core/file"
	"go.uber.org/zap"
)

// type endpointResolver struct {
// 	hostname string
// }

// func (h *endpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 	return aws.Endpoint{URL: "http://" + h.hostname}, nil
// }

func main() {
	// mgr := manager.NewManager(manager.WithHealth("/health"))
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := file.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	producer := corepb.NewFileEventbusJSONClient(
		"topic://"+config.FileEventTopic,
		redis,
	)

	opts := []file.Option{
		file.WithProducer(producer),
	}

	// if config.RedisAddr == "" {
	log.Info("Starting up with Memory store")
	opts = append(opts, file.WithFileStore(file.NewFileMemoryStore("/api/v1/fileput/")))
	// } else {
	// 	log.With(
	// 		zap.String("dynamoAddr", config.RedisAddr),
	// 	).Info("Starting up with Redis store")
	// 	opts = append(opts,
	// 		file.WithFileStore(
	// 			file.NewFileRedisStore(config.RedisAddr),
	// 		),
	// 	)
	// }

	api := corepb.NewFileServiceServer(file.NewFileServer(opts...))

	mgr.Start(api)
}

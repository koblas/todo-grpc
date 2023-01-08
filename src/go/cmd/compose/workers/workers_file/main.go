package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := workers_file.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// var builder workers.SqsConsumerBuilder
	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	opts := []workers_file.Option{
		workers_file.WithProducer(
			core.NewFileEventbusJSONClient(
				"topic://"+config.FileEventsTopic,
				redis,
			),
		),
		workers_file.WithFileService(
			core.NewFileServiceProtobufClient(
				"http://"+config.FileServiceAddr,
				&http.Client{},
			),
		),
		workers_file.WithUserService(
			core.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
	}

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), config.FileEventsTopic, workers_file.GetHandler(config, opts...)))
}

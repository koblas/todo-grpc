package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := workers.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// var builder workers.SqsConsumerBuilder
	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	opts := []workers.Option{
		workers.WithSendEmail(
			core.NewSendEmailServiceProtobufClient(
				"queue://"+config.SendEmail,
				redis,
			),
		),
	}

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), config.UserEventsTopic, workers.GetHandler(config, opts...)))
}

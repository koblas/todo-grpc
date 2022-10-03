package main

import (
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := workers.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// var builder workers.SqsConsumerBuilder
	redis := redisutil.NewTwirpRedis(util.Getenv("REDIS_ADDR", "redis:6379"))

	opts := []workers.Option{
		workers.WithSendEmail(
			core.NewSendEmailServiceProtobufClient(
				"queue://send-email",
				redis,
			),
		),
	}

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), "user-events", workers.GetHandler(ssmConfig, opts...)))
}

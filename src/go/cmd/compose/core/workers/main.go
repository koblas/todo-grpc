package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := workers.SsmConfig{}
	err := awsutil.LoadEnvConfig("/common/", &ssmConfig)
	if err != nil {
		log.Fatal(err.Error())
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

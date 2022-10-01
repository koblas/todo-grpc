package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/core/user"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig user.SsmConfig
	if err := awsutil.LoadEnvConfig("/common/", &ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(util.Getenv("REDIS_ADDR", "redis:6379"))

	producer := core.NewUserEventServiceJSONClient(
		"topic://user-events",
		redis,
	)

	opts := []user.Option{
		user.WithProducer(producer),
		user.WithUserStore(user.NewUserMemoryStore()),
	}

	api := core.NewUserServiceServer(user.NewUserServer(opts...))

	mgr.Start(api)
}

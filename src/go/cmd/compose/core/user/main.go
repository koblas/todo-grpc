package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/redis"
	"github.com/koblas/grpc-todo/pkg/manager"
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

	producer, err := redis.NewRedisProducer(ssmConfig.EventArn)
	if err != nil {
		log.With(zap.Error(err)).Fatal("failed to create producer")
	}

	opts := []user.Option{
		user.WithProducer(producer),
		user.WithUserStore(user.NewUserDynamoStore()),
	}

	api := core.NewUserServiceServer(user.NewUserServer(opts...))

	mgr.Start(api)
}

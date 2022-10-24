package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := auth.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// Connect to the user service
	opts := []auth.Option{
		auth.WithUserClient(core.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
		auth.WithOAuthClient(core.NewAuthUserServiceJSONClient("lambda://core-oauth-user", awsutil.NewTwirpCallLambda())),
	}

	rdb := redisutil.NewClient(ssmConfig.RedisAddr)
	// Connect to redis
	if rdb != nil {
		log.With("address", ssmConfig.RedisAddr).Info("Redis Address")
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	api := publicapi.NewAuthenticationServiceServer(auth.NewAuthenticationServer(ssmConfig, opts...))

	mgr.Start(api)
}

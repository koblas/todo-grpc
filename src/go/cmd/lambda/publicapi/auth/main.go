package main

import (
	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := auth.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// Connect to the user service
	opts := []auth.Option{
		auth.WithUserClient(corepb.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
		auth.WithOAuthClient(corepb.NewAuthUserServiceJSONClient("lambda://core-oauth-user", awsutil.NewTwirpCallLambda())),
	}

	rdb := redisutil.NewClient(config.RedisAddr)
	// Connect to redis
	if rdb != nil {
		log.With("address", config.RedisAddr).Info("Redis Address")
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	api := apipb.NewAuthenticationServiceServer(auth.NewAuthenticationServer(config, opts...))

	mgr.Start(api)
}

package main

import (
	"github.com/koblas/grpc-todo/gen/api/auth/v1/authv1connect"
	"github.com/koblas/grpc-todo/gen/core/oauth_user/v1/oauth_userv1connect"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"go.uber.org/zap"
)

type Config struct {
	RedisAddr string
	JwtSecret string `validate:"min=32"`
	// UserServiceAddr      string
	// OauthUserServiceAddr string
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// Connect to the user service
	opts := []auth.Option{
		auth.WithUserClient(userv1connect.NewUserServiceClient(
			awsutil.NewTwirpCallLambda(),
			"lambda://core-user",
		)),
		auth.WithOAuthClient(oauth_userv1connect.NewOAuthUserServiceClient(
			awsutil.NewTwirpCallLambda(),
			"lambda://core-oauth-user",
		)),
	}

	rdb := redisutil.NewClient(config.RedisAddr)
	// Connect to redis
	if rdb != nil {
		log.With("address", config.RedisAddr).Info("Redis Address")
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	_, api := authv1connect.NewAuthenticationServiceHandler(auth.NewAuthenticationServer(config.JwtSecret, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

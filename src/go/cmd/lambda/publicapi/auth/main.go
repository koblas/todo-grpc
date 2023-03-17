package main

import (
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
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
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// Connect to the user service
	opts := []auth.Option{
		auth.WithUserClient(corev1connect.NewUserServiceClient(
			awsutil.NewTwirpCallLambda(),
			"lambda://core-user",
		)),
		auth.WithOAuthClient(corev1connect.NewAuthUserServiceClient(
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

	_, api := apiv1connect.NewAuthenticationServiceHandler(auth.NewAuthenticationServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

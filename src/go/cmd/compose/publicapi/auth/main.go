package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/confmgr"
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

	config := auth.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []auth.Option{
		auth.WithUserClient(
			core.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
		auth.WithOAuthClient(
			core.NewAuthUserServiceProtobufClient(
				"http://"+config.OauthUserServiceAddr,
				&http.Client{},
			),
		),
	}

	rdb := redisutil.NewClient(config.RedisAddr)
	if rdb != nil {
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	api := publicapi.NewAuthenticationServiceServer(auth.NewAuthenticationServer(config, opts...))

	mgr.Start(api)
}

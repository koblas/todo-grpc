package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := auth.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []auth.Option{
		auth.WithUserClient(
			corepbv1.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
		auth.WithOAuthClient(
			corepbv1.NewAuthUserServiceProtobufClient(
				"http://"+config.OauthUserServiceAddr,
				&http.Client{},
			),
		),
	}

	rdb := redisutil.NewClient(config.RedisAddr)
	if rdb != nil {
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	api := apipbv1.NewAuthenticationServiceServer(auth.NewAuthenticationServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

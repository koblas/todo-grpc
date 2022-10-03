package main

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := auth.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []auth.Option{
		auth.WithUserClient(
			core.NewUserServiceProtobufClient(
				"http://"+util.Getenv("USER_SERVICE_ADDR", ":13001"),
				&http.Client{},
			),
		),
		auth.WithOAuthClient(
			core.NewAuthUserServiceProtobufClient(
				"http://"+util.Getenv("OAUTH_USER_SERVICE_ADDR", ":13002"),
				&http.Client{},
			),
		),
	}

	if ssmConfig.RedisAddr != "" {
		log.With("address", ssmConfig.RedisAddr).Info("Redis Address")
		// TODO - re-enable this
		rdb := redis.NewClient(&redis.Options{
			Addr:        ssmConfig.RedisAddr,
			Password:    "",                     // no password set
			DB:          0,                      // use default DB
			DialTimeout: time.Millisecond * 200, // either it happens or it doesn't
		})

		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	api := publicapi.NewAuthenticationServiceServer(auth.NewAuthenticationServer(ssmConfig, opts...))

	mgr.Start(api)
}

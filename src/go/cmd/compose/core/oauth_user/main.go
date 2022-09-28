package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/redis"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/util"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := ouser.SsmConfig{}
	oauthConfig := ouser.OauthConfig{}
	err := awsutil.LoadEnvConfig("/common/", &ssmConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load configuration")
	}
	err = awsutil.LoadSsmConfig("/oauth/", &oauthConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load oauth configuration")
	}

	producer, err := redis.NewRedisProducer(ssmConfig.EventArn)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to connect to bus")
	}

	opts := []ouser.Option{
		ouser.WithUserService(
			core.NewUserServiceProtobufClient(
				util.Getenv("USER_SERVICE_ADDR", ":13001"),
				&http.Client{},
			),
		),
		ouser.WithProducer(producer),
		ouser.WithSecretManager(oauthConfig),
	}

	api := core.NewAuthUserServiceServer(ouser.NewOauthUserServer(ssmConfig, opts...))

	mgr.Start(api)
}

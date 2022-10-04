package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := ouser.SsmConfig{}
	oauthConfig := ouser.OauthConfig{}
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}
	if err := confmgr.Parse(&oauthConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(
			core.NewUserServiceProtobufClient(
				"http://"+ssmConfig.UserServiceAddr,
				&http.Client{},
			),
		),
		ouser.WithSecretManager(oauthConfig),
	}

	api := core.NewAuthUserServiceServer(ouser.NewOauthUserServer(ssmConfig, opts...))

	mgr.Start(api)
}

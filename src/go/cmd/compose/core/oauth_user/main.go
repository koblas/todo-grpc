package main

import (
	"net/http"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := ouser.Config{}
	oauthConfig := ouser.OauthConfig{}
	if err := confmgr.Parse(&config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}
	if err := confmgr.Parse(&oauthConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(
			corepbv1.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
		ouser.WithSecretManager(oauthConfig),
	}

	api := corepbv1.NewAuthUserServiceServer(ouser.NewOauthUserServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

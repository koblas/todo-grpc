package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
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
			corev1connect.NewUserServiceClient(
				&http.Client{},
				"http://"+config.UserServiceAddr,
			),
		),
		ouser.WithSecretManager(oauthConfig),
	}

	_, api := corev1connect.NewAuthUserServiceHandler(ouser.NewOauthUserServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"go.uber.org/zap"
)

type Common struct {
	JwtSecret       string `validate:"min=32"`
	UserServiceAddr string
}

type Config struct {
	Common Common
	Oauth  ouser.OauthConfig
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	cloader := confmgr.NewLoader(
		confmgr.NewLoaderEnvironment("", "_"),
		aws.NewLoaderSsm(mgr.Context(), ""),
	)
	if err := cloader.Parse(mgr.Context(), &config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	log.With(zap.Any("credentials", config.Oauth)).Info("testing")

	opts := []ouser.Option{
		ouser.WithUserService(corev1connect.NewUserServiceClient(
			awsutil.NewTwirpCallLambda(),
			"lambda://core-user",
		)),
		ouser.WithSecretManager(config.Oauth),
	}

	_, api := corev1connect.NewOAuthUserServiceHandler(
		ouser.NewOauthUserServer(config.Common.JwtSecret, opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

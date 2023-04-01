package main

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/gen/core/oauth_user/v1/oauth_userv1connect"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret       string `ssm:"/common/jwt_secret" validate:"min=32"`
	UserServiceAddr string
	Oauth           ouser.OauthConfig
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(
			userv1connect.NewUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.UserServiceAddr,
			),
		),
		ouser.WithSecretManager(config.Oauth),
	}

	mux := http.NewServeMux()
	mux.Handle(oauth_userv1connect.NewOAuthUserServiceHandler(
		ouser.NewOauthUserServer(config.JwtSecret, opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(oauth_userv1connect.OAuthUserServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

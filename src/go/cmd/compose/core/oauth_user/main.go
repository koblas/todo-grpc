package main

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
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
			corev1connect.NewUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.UserServiceAddr,
			),
		),
		ouser.WithSecretManager(config.Oauth),
	}

	mux := http.NewServeMux()
	mux.Handle(corev1connect.NewOAuthUserServiceHandler(
		ouser.NewOauthUserServer(config.JwtSecret, opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(corev1connect.UserServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

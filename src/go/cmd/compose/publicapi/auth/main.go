package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
	"go.uber.org/zap"
)

type Config struct {
	RedisAddr            string
	JwtSecret            string `validate:"min=32"`
	UserServiceAddr      string
	OauthUserServiceAddr string
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []auth.Option{
		auth.WithUserClient(
			corev1connect.NewUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.UserServiceAddr,
			),
		),
		auth.WithOAuthClient(
			corev1connect.NewOAuthUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.OauthUserServiceAddr,
			),
		),
	}

	rdb := redisutil.NewClient(config.RedisAddr)
	if rdb != nil {
		opts = append(opts, auth.WithAttemptService(auth.NewAttemptCounter("publicapi:authentication", rdb)))
	}

	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewAuthenticationServiceHandler(
		auth.NewAuthenticationServer(config.JwtSecret, opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(apiv1connect.AuthenticationServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

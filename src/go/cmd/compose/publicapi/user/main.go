package main

import (
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/user"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret       string `validate:"min=32"`
	UserServiceAddr string
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []user.Option{
		user.WithUserService(
			corev1connect.NewUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.UserServiceAddr,
			),
		),
		user.WithGetUserId(authHelper),
	}

	_, api := apiv1connect.NewUserServiceHandler(
		user.NewUserServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(mgr.WrapHttpHandler(api))
}

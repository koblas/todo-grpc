package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config user.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []user.Option{
		user.WithUserService(
			corev1connect.NewUserServiceClient(
				&http.Client{},
				"http://"+config.UserServiceAddr,
			),
		),
	}

	_, api := apiv1connect.NewUserServiceHandler(
		user.NewUserServer(config, opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
	)

	mgr.Start(mgr.WrapHttpHandler(api))
}

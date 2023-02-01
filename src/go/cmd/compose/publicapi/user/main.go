package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
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
			corepbv1.NewUserServiceProtobufClient(
				"http://"+config.UserServiceAddr,
				&http.Client{},
			),
		),
	}

	api := apipbv1.NewUserServiceServer(user.NewUserServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

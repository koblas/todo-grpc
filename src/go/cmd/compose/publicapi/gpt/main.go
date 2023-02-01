package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/gpt"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config gpt.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []gpt.Option{}

	api := apipbv1.NewGptServiceServer(gpt.NewGptServer(config, opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

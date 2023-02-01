package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/fileput"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := fileput.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []fileput.Option{
		fileput.WithFileService(
			corepbv1.NewFileServiceProtobufClient(
				"http://"+config.FileServiceAddr,
				&http.Client{},
			),
		),
	}

	api := fileput.NewFilePutServer(config, opts...)

	mgr.Start(mgr.WrapHttpHandler(api))
}

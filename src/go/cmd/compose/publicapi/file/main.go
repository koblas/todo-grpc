package main

import (
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/file"
	"go.uber.org/zap"
)

type Config struct {
	UploadBucket  string
	JwtSecret     string `validate:"min=32"`
	MinioEndpoint string
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{
		MinioEndpoint: "s3.amazonaws.com",
	}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []file.Option{
		file.WithFileStore(
			filestore.NewMinioProvider(config.MinioEndpoint),
		),
		file.WithGetUserId(authHelper),
		file.WithUploadBucket(config.UploadBucket),
	}

	_, api := apiv1connect.NewFileServiceHandler(
		file.NewFileServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(mgr.WrapHttpHandler(api))
}

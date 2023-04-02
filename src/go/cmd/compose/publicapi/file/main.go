package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/file/v1/filev1connect"
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
	mgr := manager.NewManager()
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

	mux := http.NewServeMux()
	mux.Handle(filev1connect.NewFileServiceHandler(
		file.NewFileServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(filev1connect.FileServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

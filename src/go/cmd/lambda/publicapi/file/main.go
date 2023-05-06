package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/api/file/v1/filev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
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
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []file.Option{
		file.WithFileStore(filestore.NewAwsProvider()),
		file.WithGetUserId(authHelper),
		file.WithUploadBucket(config.UploadBucket),
	}

	_, api := filev1connect.NewFileServiceHandler(
		file.NewFileServer(opts...),
		bufcutil.WithJSON(),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

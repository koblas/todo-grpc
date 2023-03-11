package main

import (
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/file"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config file.Config
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []file.Option{
		// file.WithFileService(corepbv1.NewFileServiceJSONClient("lambda://core-file", awsutil.NewTwirpCallLambda())),
		file.WithFileStore(filestore.NewAwsProvider()),
	}

	api := apipbv1.NewFileServiceServer(file.NewFileServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

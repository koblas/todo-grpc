package main

import (
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/gpt"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config gpt.Config
	if err := confmgr.ParseWithContext(mgr.Context(), &config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []gpt.Option{}

	api := apipbv1.NewGptServiceServer(gpt.NewGptServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/user"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig user.SsmConfig
	if err := confmgr.Parse(&ssmConfig, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []user.Option{
		user.WithUserService(core.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
	}

	api := publicapi.NewUserServiceServer(user.NewUserServer(ssmConfig, opts...))

	mgr.Start(api)
}

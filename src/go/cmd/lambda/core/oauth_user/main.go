package main

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := ouser.Config{}
	oauthConfig := ouser.OauthConfig{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load general configuration")
	}
	if err := confmgr.Parse(&oauthConfig, aws.NewLoaderSsm("/oauth/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load oauth configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(corepb.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
		ouser.WithSecretManager(oauthConfig),
	}

	api := corepb.NewAuthUserServiceServer(ouser.NewOauthUserServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

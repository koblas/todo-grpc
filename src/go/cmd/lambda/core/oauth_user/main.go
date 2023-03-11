package main

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
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
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load general configuration")
	}
	if err := confmgr.Parse(&oauthConfig, aws.NewLoaderSsm(mgr.Context(), "/oauth/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load oauth configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(corepbv1.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
		ouser.WithSecretManager(oauthConfig),
	}

	api := corepbv1.NewAuthUserServiceServer(ouser.NewOauthUserServer(config, opts...))

	mgr.Start(awsutil.HandleApiLambda(api))
}

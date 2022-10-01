package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/util"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := ouser.SsmConfig{}
	oauthConfig := ouser.OauthConfig{}
	err := awsutil.LoadEnvConfig("/common/", &ssmConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load configuration")
	}
	err = awsutil.LoadEnvConfig("/oauth/", &oauthConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load oauth configuration")
	}

	opts := []ouser.Option{
		ouser.WithUserService(
			core.NewUserServiceProtobufClient(
				"http://"+util.Getenv("USER_SERVICE_ADDR", ":13001"),
				&http.Client{},
			),
		),
		ouser.WithSecretManager(oauthConfig),
	}

	api := core.NewAuthUserServiceServer(ouser.NewOauthUserServer(ssmConfig, opts...))

	mgr.Start(api)
}

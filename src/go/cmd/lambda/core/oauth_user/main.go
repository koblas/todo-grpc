package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	ouser "github.com/koblas/grpc-todo/services/core/oauth_user"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := ouser.SsmConfig{}
	oauthConfig := ouser.OauthConfig{}
	err := awsutil.LoadSsmConfig("/common/", ssmConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load configuration")
	}
	err = awsutil.LoadSsmConfig("/oauth/", oauthConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to load oauth configuration")
	}

	// Connect to the user service
	producer, err := aws.NewAwsProducer(ssmConfig.EventArn)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to connect to bus")
	}

	opts := []ouser.Option{
		ouser.WithUserService(core.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())),
		ouser.WithProducer(producer),
		ouser.WithSecretManager(oauthConfig),
	}

	api := core.NewAuthUserServiceServer(ouser.NewOauthUserServer(ssmConfig, opts...))

	mgr.Start(api)
}

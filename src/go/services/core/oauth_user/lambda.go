package user

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
)

type SsmConfig struct {
	EventArn  string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET"`
}

type OauthConfig struct {
	GoogleClientId string `ssm:"google_client_id" environment:"GOOGLE_CLIENT_ID"`
	GoogleSecret   string `ssm:"google_secret" environment:"GOOGLE_SECRET"`
	GitHubClientId string `ssm:"github_client_id" environment:"GITHUB_CLIENT_ID"`
	GitHubSecret   string `ssm:"github_secret" environment:"GITHUB_SECRET"`
}

func (conf OauthConfig) GetSecret(provider string) (string, string, error) {
	if provider == "github" {
		return conf.GitHubClientId, conf.GitHubSecret, nil
	}
	if provider == "google" {
		return conf.GoogleClientId, conf.GoogleSecret, nil
	}
	return "", "", nil
}

var ssmConfig SsmConfig
var oauthConfig OauthConfig
var api core.TwirpServer

func HandleLambda() func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if api == nil {
		if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
			log.Fatal(err.Error())
		}
		if err := awsutil.LoadSsmConfig("/oauth/", &oauthConfig); err != nil {
			log.Fatal(err.Error())
		}

		producer, err := aws.NewAwsPublish(ssmConfig.EventArn)
		if err != nil {
			log.Fatal(err)
		}

		userService := core.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())

		s := NewOauthUserServer(producer, userService, ssmConfig, oauthConfig)
		api = core.NewAuthUserServiceServer(s)
	}

	linfo := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), linfo)

	return awsutil.HandleLambda(ctx, api)
}

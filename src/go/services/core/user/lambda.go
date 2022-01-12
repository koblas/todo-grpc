package user

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
)

type SsmConfig struct {
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
}

var ssmConfig *SsmConfig
var api core.TwirpServer

func HandleLambda() func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log := logger.NewZap(logger.LevelInfo)

	if ssmConfig == nil || api == nil {
		log.Info("BEGIN: lambda initialization")
		ssmConfig = &SsmConfig{}
		err := awsutil.LoadSsmConfig("/common/", ssmConfig)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Info("Getting parameters finished")

		producer, err := aws.NewAwsPublish(ssmConfig.EventArn)
		if err != nil {
			log.With("error", err).Fatal("Unable to initilaize AwsPublisher")
		}

		s := NewUserServer(producer, NewUserDynamoStore())
		api = core.NewUserServiceServer(s)
		log.Info("FINISHED: lambda initialization")
	}

	ctx := logger.ToContext(context.Background(), log)

	return awsutil.HandleLambda(ctx, api)
}

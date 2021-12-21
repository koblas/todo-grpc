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
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
}

var ssmConfig *SsmConfig
var api core.TwirpServer

func HandleLambda() func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if ssmConfig == nil || api == nil {
		ssmConfig = &SsmConfig{}
		err := awsutil.LoadSsmConfig("/common/", ssmConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		producer, err := aws.NewAwsPublish(ssmConfig.EventArn)
		if err != nil {
			log.Fatal(err)
		}

		s := NewUserServer(producer, NewUserDynamoStore())
		api = core.NewUserServiceServer(s)
	}

	linfo := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), linfo)

	return awsutil.HandleLambda(ctx, api)
}

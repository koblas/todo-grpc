package todo

import (
	"context"
	"log"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
)

type SsmConfig struct {
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
}

var lambdaHandler awsutil.TwirpHttpHandler

func init() {
	var ssmConfig SsmConfig
	var api core.TwirpServer

	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := aws.NewAwsPublish(ssmConfig.EventArn)
	if err != nil {
		log.Fatal(err)
	}

	s := NewTodoServer(producer, NewTodoDynamoStore())
	api = core.NewTodoServiceServer(s)

	linfo := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), linfo)

	lambdaHandler = awsutil.HandleLambda(ctx, api)
}

func HandleLambda() awsutil.TwirpHttpHandler {
	return lambdaHandler
}

package send_email

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
	SmtpAddr string `ssm:"smtp/addr" environment:"SMTP_ADDR"`
	SmtpUser string `ssm:"smtp/username" environment:"SMTP_USERNAME"`
	SmtpPass string `ssm:"smtp/password" environment:"SMTP_PASSWORD"`
}

var lambdaHandler awsutil.TwirpHttpSqsHandler

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
	smtp := NewSmtpService(ssmConfig)
	s := NewSendEmailServer(producer, smtp)
	api = core.NewSendEmailServiceServer(s)

	linfo := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), linfo)

	lambdaHandler = awsutil.HandleSqsLambda(ctx, api, nil)
}

func HandleLambda() awsutil.TwirpHttpSqsHandler {
	return lambdaHandler
}

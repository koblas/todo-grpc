package workers

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	awsbus "github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
)

type SsmConfig struct {
	UrlBase string `ssm:"url_base" environment:"URL_BASE_UI"`
}

var handler awsbus.SqsConsumerFunc

func HandleLambda(ctx context.Context, event events.SQSEvent) {
	mode := os.Getenv("SQS_HANDLER")
	baselogger := logger.NewZap(logger.LevelDebug)
	log := baselogger.With("SQS_HANDLER", mode)

	if mode == "" {
		log.Fatal("SQS_HANDLER environment variable must be set")
	}
	log.Info("Starting queue worker")

	if handler == nil {
		ssmConfig := &SsmConfig{}
		err := awsutil.LoadSsmConfig("/common/", ssmConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)

		if mode == "" {
			log.Fatal("SQS_HANDLER is not defiend")
		}

		var builder SqsConsumerBuilder

		for _, item := range workers {
			if item.GroupName == mode {
				builder = item.Build
				break
			}
		}

		if builder == nil {
			log.Fatal("Unable to find handler")
		}
		handler = builder(ssmConfig)
	}

	sender, err := awsbus.NewAwsSqsConsumer(handler)
	if err != nil {
		log.With("error", err).Fatal("Unable to start consumer")
	}

	sender.AddMessagesLambda(logger.ToContext(ctx, log), event)
}

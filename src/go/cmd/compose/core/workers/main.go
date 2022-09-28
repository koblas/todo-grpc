package main

import (
	"os"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	awsbus "github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mode := os.Getenv("SQS_HANDLER")

	mgr := manager.NewManager()
	log := mgr.Logger().With("SQS_HANDLER", mode)

	if mode == "" {
		log.Fatal("SQS_HANDLER environment variable must be set")
	}

	ssmConfig := workers.SsmConfig{}
	err := awsutil.LoadEnvConfig("/common/", &ssmConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	var builder workers.SqsConsumerBuilder

	for _, item := range workers.GetWorkers() {
		if item.GroupName == mode {
			builder = item.Build
			break
		}
	}

	opts := []workers.Option{
		workers.WithSendEmail(
			core.NewSendEmailServiceProtobufClient(
				"sqs://send-email",
				awsutil.NewTwirpCallLambda(),
			),
		),
	}

	if builder == nil {
		log.Fatal("Unable to find handler")
	}
	handler := builder(ssmConfig, opts...)

	sender := awsbus.NewAwsSqsConsumer(handler)
	if err != nil {
		log.With(zap.Error(err)).Fatal("Unable to start consumer")
	}

	mgr.StartConsumer(sender.AddMessagesLambda)
}

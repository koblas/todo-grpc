package main

import (
	"log"
	"os"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
)

func main() {
	mode := os.Getenv("SQS_HANDLER")
	if mode == "" {
		log.Fatal("SQS_HANDLER environment variable must be set")
	}

	mgr := manager.NewManager()
	log := mgr.Logger().With("SQS_HANDLER", mode)

	ssmConfig := workers.SsmConfig{}
	err := awsutil.LoadSsmConfig("/common/", &ssmConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// var builder workers.SqsConsumerBuilder

	// for _, item := range workers.GetWorkers() {
	// 	if item.GroupName == mode {
	// 		builder = item.Build
	// 		break
	// 	}
	// }

	// if builder == nil {
	// 	log.Fatal("Unable to find handler")
	// }

	opts := []workers.Option{
		workers.WithOnly(mode),
		workers.WithSendEmail(
			core.NewSendEmailServiceProtobufClient(
				"sqs://send-email",
				awsutil.NewTwirpCallLambda(),
			),
		),
	}

	/*
		handler := builder(ssmConfig, opts...)

		sender := awsbus.NewAwsSqsConsumer(handler)
		if err != nil {
			log.With(zap.Error(err)).Fatal("Unable to start consumer")
		}

		mgr.StartConsumer(sender.AddMessagesLambda)
	*/
	mgr.StartConsumer(awsutil.HandleSqsLambda(workers.GetHandler(ssmConfig, opts...)))
}

package main

import (
	"log"
	"os"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/workers"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mode := os.Getenv("SQS_HANDLER")
	if mode == "" {
		log.Fatal("SQS_HANDLER environment variable must be set")
	}

	mgr := manager.NewManager()
	log := mgr.Logger().With("SQS_HANDLER", mode)

	config := workers.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
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
		handler := builder(config, opts...)

		sender := awsbus.NewAwsSqsConsumer(handler)
		if err != nil {
			log.With(zap.Error(err)).Fatal("Unable to start consumer")
		}

		mgr.StartConsumer(sender.AddMessagesLambda)
	*/
	mgr.StartConsumer(awsutil.HandleSqsLambda(workers.GetHandler(config, opts...)))
}

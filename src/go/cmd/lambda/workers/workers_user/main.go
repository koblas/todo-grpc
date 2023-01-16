package main

import (
	"log"
	"os"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	workers "github.com/koblas/grpc-todo/services/workers/workers_user"
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
			corepb.NewSendEmailServiceProtobufClient(
				"sqs://send-email",
				awsutil.NewTwirpCallLambda(),
			),
		),
	}

	mgr.Start(awsutil.HandleSqsLambda(workers.GetHandler(config, opts...)))
}

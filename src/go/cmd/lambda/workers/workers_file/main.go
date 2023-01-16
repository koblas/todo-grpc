package main

import (
	"log"
	"os"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/workers/workers_file"
	"go.uber.org/zap"
)

func main() {
	mode := os.Getenv("SQS_HANDLER")
	if mode == "" {
		log.Fatal("SQS_HANDLER environment variable must be set")
	}

	mgr := manager.NewManager()
	log := mgr.Logger().With("SQS_HANDLER", mode)

	config := workers_file.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []workers_file.Option{
		workers_file.WithOnly(mode),
		workers_file.WithProducer(corepb.NewFileEventbusJSONClient(
			config.EventArn,
			awsutil.NewTwirpCallLambda(),
		)),
	}

	// mgr.StartConsumerMsg(awsutil.HandleSqsLambda(workers_file.GetHandler(config, opts...)))
	mgr.Start(awsutil.HandleSqsLambda(workers_file.BuildHandlers(config, opts...)))
}

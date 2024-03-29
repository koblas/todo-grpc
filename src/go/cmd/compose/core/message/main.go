package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/message"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr        string
	DynamoStoreAddr *string `json:"dynamo-store"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	eventbus := eventbusv1connect.NewMessageEventbusServiceClient(
		nats,
		"topic:",
	)

	opts := []message.Option{
		message.WithProducer(eventbus),
		// message.WithMessageStore(message.NewMemoryStore()),
		message.WithMessageStore(message.NewDynamoStore()),
	}

	if config.DynamoStoreAddr == nil || *config.DynamoStoreAddr == "" {
		log.Info("Starting up with Memory store")
		opts = append(opts, message.WithMessageStore(message.NewMemoryStore()))
	} else {
		log.With(
			zap.String("dynamoAddr", *config.DynamoStoreAddr),
		).Info("Starting up with DynamoDB store")
		opts = append(opts,
			message.WithMessageStore(
				message.NewDynamoStore(
					message.WithDynamoClient(
						awsutil.LocalDynamoClient(*config.DynamoStoreAddr),
					),
				),
			),
		)
	}

	mux := http.NewServeMux()
	mux.Handle(messagev1connect.NewMessageServiceHandler(
		message.NewMessageServer(opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(messagev1connect.MessageServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

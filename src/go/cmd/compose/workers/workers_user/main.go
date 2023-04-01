package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/gen/core/send_email/v1/send_emailv1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/workers/workers_user"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr  string
	SendEmail string
	UrlBase   string
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)

	opts := []workers_user.Option{
		workers_user.WithSendEmail(
			send_emailv1connect.NewSendEmailServiceClient(
				nats,
				"topic://"+config.SendEmail,
			),
		),
		workers_user.WithUrlBase(config.UrlBase),
	}
	s := workers_user.GetHandler(opts...)

	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.ConnectToTopic(eventbusv1connect.UserEventbusServiceName),
		s))
}

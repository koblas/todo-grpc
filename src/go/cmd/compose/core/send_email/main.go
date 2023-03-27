package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr string
	Smtp     struct {
		Addr     string
		Username string
		Password string
	}
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewSendEmailEventsServiceClient(
		natsutil.NewNatsClient(config.NatsAddr),
		"topic://",
	)

	s := send_email.NewSendEmailServer(producer,
		send_email.NewSmtpService(
			config.Smtp.Addr,
			config.Smtp.Username,
			config.Smtp.Password,
		),
	)

	nats := natsutil.NewNatsClient(config.NatsAddr)
	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.ConnectToTopic(corev1connect.SendEmailServiceName),
		s))
}

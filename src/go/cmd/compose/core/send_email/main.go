package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := send_email.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepb.NewSendEmailEventsProtobufClient(
		"",
		natsutil.NewNatsClient(config.NatsAddr),
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(config))

	nats := natsutil.NewNatsClient(config.NatsAddr)
	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepb.BroadcastEventbusPathPrefix),
		s))
}

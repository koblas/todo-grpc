package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/workers/workers_user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := workers_user.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	// var builder workers.SqsConsumerBuilder
	nats := natsutil.NewNatsClient(config.NatsAddr)

	opts := []workers_user.Option{
		workers_user.WithSendEmail(
			corepb.NewSendEmailServiceProtobufClient(
				"queue://"+config.SendEmail,
				nats,
			),
		),
	}
	s := workers_user.GetHandler(config, opts...)

	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepb.UserEventbusPathPrefix),
		s))
}

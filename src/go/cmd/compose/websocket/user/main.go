package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/websocket/user"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := user.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)

	producer := corepb.NewBroadcastEventbusProtobufClient(
		"",
		nats,
	)
	log.With(zap.String("nats", config.NatsAddr)).Info("Creating nats producer")

	s := user.NewUserChangeServer(
		user.WithProducer(producer),
	)

	mgr.Start(nats.TopicConsumer(mgr.Context(), natsutil.TwirpPathToNatsTopic(corepb.UserEventbusPathPrefix), s))
}

package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	config := todo.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corepb.NewBroadcastEventbusProtobufClient(
		"",
		natsutil.NewNatsClient(config.NatsAddr),
	)

	s := todo.NewTodoChangeServer(
		todo.WithProducer(producer),
	)
	// sc := corepb.NewTodoEventbusServer(s, twirp.WithServerPathPrefix(""))
	// mux := http.NewServeMux()
	// mux.Handle(corepb.TodoEventbusPathPrefix, corepb.NewTodoEventbusServer(s))

	nats := natsutil.NewNatsClient(config.NatsAddr)
	// mgr.StartConsumer(nats.TopicConsumer(mgr.Context(), strings.Trim(sc.PathPrefix(), "/")+".*", "websocket.todo", mux))
	mgr.Start(nats.TopicConsumer(
		mgr.Context(),
		natsutil.TwirpPathToNatsTopic(corepb.BroadcastEventbusPathPrefix),
		s))
}

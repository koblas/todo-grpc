package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
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

	redis := redisutil.NewTwirpRedis(config.RedisAddr)
	producer := corepb.NewBroadcastEventbusJSONClient(
		"topic://"+config.BroadcastTopic,
		redis,
	)

	s := todo.NewTodoChangeServer(
		todo.WithProducer(producer),
	)
	mux := http.NewServeMux()
	mux.Handle(corepb.TodoEventbusPathPrefix, corepb.NewTodoEventbusServer(s))

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), config.TodoEventsTopic, mux))
}

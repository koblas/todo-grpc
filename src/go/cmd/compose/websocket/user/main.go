package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
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

	client := redisutil.RedisWsClient(mgr.Context(), config.RedisAddr, config.WebsocketBroadcast)

	s := user.NewUserChangeServer(
		user.WithStore(websocket.NewRedisStore(config.RedisAddr)),
		user.WithClient(client),
	)
	mux := http.NewServeMux()
	mux.Handle(corepb.UserEventbusPathPrefix, corepb.NewUserEventbusServer(s))

	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), config.UserEventsTopic, mux))
}

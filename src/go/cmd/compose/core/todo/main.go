package main

import (
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/core/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(config.RedisAddr)
	eventbus := core.NewTodoEventbusJSONClient(
		"topic://"+config.TodoEvents,
		redis,
	)

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
	}

	api := core.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(api)
}

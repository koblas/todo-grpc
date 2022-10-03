package main

import (
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/core/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig todo.SsmConfig
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(util.Getenv("REDIS_ADDR", "redis:6379"))
	eventbus := core.NewTodoEventbusJSONClient(
		"topic://todo-events",
		redis,
	)

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
	}

	api := core.NewTodoServiceServer(todo.NewTodoServer(opts...))

	mgr.Start(api)
}

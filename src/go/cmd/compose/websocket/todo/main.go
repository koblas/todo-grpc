package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/eventbus"
	redisbus "github.com/koblas/grpc-todo/pkg/eventbus/redis"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

type redisPublish struct {
	producer eventbus.SimpleProducer
}

func (svc redisPublish) PostToConnection(ctx context.Context, params *apigatewaymanagementapi.PostToConnectionInput, optFns ...func(*apigatewaymanagementapi.Options)) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	msg := awsutil.ConvertApiGatewayToMessage(params)

	err := svc.producer.Write(ctx, &msg)

	return nil, err
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := todo.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := redisbus.NewProducer(ssmConfig.RedisAddr, "websocket-broadcast")

	s := todo.NewTodoChangeServer(
		todo.WithStore(websocket.NewRedisStore(ssmConfig.RedisAddr)),
		todo.WithClient(redisPublish{producer}),
	)
	mux := http.NewServeMux()
	mux.Handle(core.TodoEventbusPathPrefix, core.NewTodoEventbusServer(s))

	redis := redisutil.NewTwirpRedis(ssmConfig.RedisAddr)

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), "todo-events", mux))
}

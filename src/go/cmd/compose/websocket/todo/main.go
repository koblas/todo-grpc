package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
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

	config := todo.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := redisbus.NewProducer(config.RedisAddr, config.WebsocketBroadcast)

	s := todo.NewTodoChangeServer(
		todo.WithStore(websocket.NewRedisStore(config.RedisAddr)),
		todo.WithClient(redisPublish{producer}),
	)
	mux := http.NewServeMux()
	mux.Handle(core.TodoEventbusPathPrefix, core.NewTodoEventbusServer(s))

	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	mgr.StartConsumer(redis.TopicConsumer(mgr.Context(), config.TodoEventsTopic, mux))
}

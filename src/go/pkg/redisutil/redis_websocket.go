package redisutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus"
	redisbus "github.com/koblas/grpc-todo/pkg/eventbus/redis"
)

// type PostToConnectionAPI interface {
// 	PostToConnection(ctx context.Context, params *apigatewaymanagementapi.PostToConnectionInput, optFns ...func(*apigatewaymanagementapi.Options)) (*apigatewaymanagementapi.PostToConnectionOutput, error)
// }

type redisPublish struct {
	producer eventbus.SimpleProducer
}

func (svc *redisPublish) PostToConnection(ctx context.Context, params *apigatewaymanagementapi.PostToConnectionInput, optFns ...func(*apigatewaymanagementapi.Options)) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	msg := awsutil.ConvertApiGatewayToMessage(params)

	err := svc.producer.Write(ctx, &msg)

	return nil, err
}

func RedisWsClient(ctx context.Context, redisAddr string, wsBroadcast string) *redisPublish {
	producer := redisbus.NewProducer(redisAddr, wsBroadcast)

	return &redisPublish{producer}
}

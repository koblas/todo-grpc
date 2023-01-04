package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	store "github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/publicapi/websocket"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := websocket.Config{}
	if err := confmgr.Parse(&config, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	handler := websocket.NewWebsocketHandler(
		config,
		websocket.WithStore(
			store.NewUserDynamoStore(
				store.WithDynamoTable(config.ConnDb),
			),
		),
	)

	lambdaHandler := func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		log := logger.FromContext(ctx)

		log = log.With("requestId", req.RequestContext.RequestID).
			With("connectionId", req.RequestContext.ConnectionID).
			With("routeKey", req.RequestContext.RouteKey)

		return handler.HandleRequest(logger.ToContext(ctx, log), req)
	}

	lambda.StartWithContext(mgr.Context(), lambdaHandler)
}

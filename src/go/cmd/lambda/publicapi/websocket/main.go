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

type Config struct {
	JwtSecret string `validate:"min=32"`
	ConnDb    string `validate:"required"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	handler := websocket.NewWebsocketHandler(
		config.JwtSecret,
		websocket.WithStore(
			store.NewWsConnectionDynamoStore(
				store.WithDynamoTable(config.ConnDb),
			),
		),
	)

	lambdaHandler := func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		log := logger.FromContext(ctx).With(
			zap.String("requestId", req.RequestContext.RequestID),
			zap.String("connectionId", req.RequestContext.ConnectionID),
			zap.String("routeKey", req.RequestContext.RouteKey),
		)

		return handler.HandleRequest(logger.ToContext(ctx, log), req)
	}

	lambda.StartWithOptions(lambdaHandler, lambda.WithContext(mgr.Context()))
}

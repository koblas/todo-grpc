package websocket

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"go.uber.org/zap"
)

type SsmConfig struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	ConnDb    string `environment:"CONN_DB"`
}

type lambdaHandlerType func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

var lambdaHandler lambdaHandlerType

func init() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig SsmConfig
	if err := confmgr.Parse(&ssmConfig, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	handler := NewWebsocketHandler(ssmConfig)

	lambdaHandler = func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		log = log.With("requestId", req.RequestContext.RequestID).
			With("connectionId", req.RequestContext.ConnectionID).
			With("routeKey", req.RequestContext.RouteKey)

		return handler.HandleRequest(logger.ToContext(ctx, log), req)
	}
}

func HandleLambda() lambdaHandlerType {
	return lambdaHandler
}

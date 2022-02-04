package websocket

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
)

type SsmConfig struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET"`
	ConnDb    string `environment:"CONN_DB"`
}

type lambdaHandlerType func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

var lambdaHandler lambdaHandlerType

func init() {
	var ssmConfig SsmConfig

	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
	}

	linfo := logger.NewZap(logger.LevelInfo)

	handler := NewWebsocketHandler(ssmConfig)

	lambdaHandler = func(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		log := linfo.With("requestId", req.RequestContext.RequestID).
			With("connectionId", req.RequestContext.ConnectionID).
			With("routeKey", req.RequestContext.RouteKey)

		return handler.HandleRequest(logger.ToContext(ctx, log), req)
	}
}

func HandleLambda() lambdaHandlerType {
	return lambdaHandler
}

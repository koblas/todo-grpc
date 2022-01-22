package todo

import (
	"context"
	"log"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
)

type SsmConfig struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET"`
}

var lambdaHandler awsutil.TwirpHttpApiHandler

func init() {
	var ssmConfig SsmConfig
	var api core.TwirpServer

	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
	}

	todoService := core.NewTodoServiceJSONClient("lambda://core-todo", awsutil.NewTwirpCallLambda())

	s := NewTodoServer(todoService, ssmConfig)
	api = publicapi.NewTodoServiceServer(s)

	linfo := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), linfo)

	lambdaHandler = awsutil.HandleApiLambda(ctx, api)
}

func HandleLambda() awsutil.TwirpHttpApiHandler {
	return lambdaHandler
}

package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/websocket/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := todo.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	store := websocket.NewUserDynamoStore(websocket.WithDynamoTable(ssmConfig.ConnDb))

	cfg, err := config.LoadDefaultConfig(mgr.Context())
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to connect to aws")
	}

	endpointResolver := func(o *apigatewaymanagementapi.Options) {
		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(ssmConfig.WsEndpoint)
	}
	client := apigatewaymanagementapi.NewFromConfig(cfg, endpointResolver)

	s := todo.NewTodoChangeServer(
		todo.WithStore(store),
		todo.WithClient(client),
	)
	mux := http.NewServeMux()
	mux.Handle(core.TodoEventbusPathPrefix, core.NewTodoEventbusServer(s))

	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
}

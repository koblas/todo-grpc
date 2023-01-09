package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/websocket/user"
	"go.uber.org/zap"
)

// func main() {
// 	mgr := manager.NewManager()
// 	log := mgr.Logger()

// 	conf := user.Config{}
// 	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
// 		log.With(zap.Error(err)).Fatal("failed to load configuration")
// 	}

// 	store := websocket.NewUserDynamoStore(websocket.WithDynamoTable(conf.ConnDb))

// 	cfg, err := config.LoadDefaultConfig(mgr.Context())
// 	if err != nil {
// 		log.With(zap.Error(err)).Fatal("unable to connect to aws")
// 	}

// 	endpointResolver := func(o *apigatewaymanagementapi.Options) {
// 		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(conf.WsEndpoint)
// 	}
// 	client := apigatewaymanagementapi.NewFromConfig(cfg, endpointResolver)

// 	s := user.NewUserChangeServer(
// 		user.WithStore(store),
// 		user.WithClient(client),
// 	)
// 	mux := http.NewServeMux()
// 	mux.Handle(corepb.UserEventbusPathPrefix, corepb.NewUserEventbusServer(s))

// 	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
// }

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := user.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	client, err := awsutil.ApigwClient(mgr.Context(), conf.WsEndpoint)
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to connect to aws")
	}

	store := websocket.NewUserDynamoStore(websocket.WithDynamoTable(conf.ConnDb))

	s := user.NewUserChangeServer(
		user.WithStore(store),
		user.WithClient(client),
	)
	mux := http.NewServeMux()
	mux.Handle(corepb.UserEventbusPathPrefix, corepb.NewUserEventbusServer(s))

	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
}

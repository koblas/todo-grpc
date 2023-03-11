package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/services/websocket/broadcast"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	conf := broadcast.Config{}
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	client, err := awsutil.ApigwClient(mgr.Context(), conf.WsEndpoint)
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to connect to aws")
	}

	store := websocket.NewWsConnectionDynamoStore(websocket.WithDynamoTable(conf.ConnDb))

	handlers := broadcast.NewBroadcastServer(
		broadcast.WithStore(store),
		broadcast.WithClient(client),
	)

	mgr.Start(awsutil.HandleSqsLambda(handlers))
}

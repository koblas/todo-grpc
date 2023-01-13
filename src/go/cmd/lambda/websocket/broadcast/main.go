package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/corepb"
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
	if err := confmgr.Parse(&conf, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	client, err := awsutil.ApigwClient(mgr.Context(), conf.WsEndpoint)
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to connect to aws")
	}

	store := websocket.NewWsConnectionDynamoStore(websocket.WithDynamoTable(conf.ConnDb))

	s := broadcast.NewBroadcastServer(
		broadcast.WithStore(store),
		broadcast.WithClient(client),
	)
	mux := http.NewServeMux()
	mux.Handle(corepb.BroadcastEventbusPathPrefix, corepb.NewBroadcastEventbusServer(s))

	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
}

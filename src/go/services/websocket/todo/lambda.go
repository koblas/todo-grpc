package todo

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

type SsmConfig struct {
	UrlBase    string `ssm:"url_base" environment:"URL_BASE_UI"`
	ConnDb     string `environment:"CONN_DB"`
	WsEndpoint string `environment:"WS_ENDPOINT"`
}

var lambdaHandler awsutil.TwirpHttpSqsHandler

func init() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig SsmConfig
	if err := confmgr.Parse(&ssmConfig, aws.NewLoaderSsm("/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	store := websocket.NewUserDynamoStore(websocket.WithDynamoTable(ssmConfig.ConnDb))

	s := NewTodoServer(store, ssmConfig.WsEndpoint)
	mux := http.NewServeMux()
	mux.Handle(core.TodoEventbusPathPrefix, core.NewTodoEventbusServer(s))

	lambdaHandler = awsutil.HandleSqsLambda(mux)
}

func HandleLambda() awsutil.TwirpHttpSqsHandler {
	return lambdaHandler
}

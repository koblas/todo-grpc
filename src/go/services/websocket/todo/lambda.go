package todo

import (
	"log"
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/twpb/core"
)

type SsmConfig struct {
	UrlBase    string `ssm:"url_base" environment:"URL_BASE_UI"`
	ConnDb     string `environment:"CONN_DB"`
	WsEndpoint string `environment:"WS_ENDPOINT"`
}

var lambdaHandler awsutil.TwirpHttpSqsHandler

func init() {
	var ssmConfig SsmConfig

	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
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

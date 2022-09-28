package todo

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var ssmConfig todo.SsmConfig
	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	opts := []todo.Option{
		todo.WithTodoService(core.NewTodoServiceJSONClient("lambda://core-todo", awsutil.NewTwirpCallLambda())),
	}

	api := publicapi.NewTodoServiceServer(todo.NewTodoServer(ssmConfig, opts...))

	mgr.Start(api)
}

package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/api/todo/v1/todov1connect"
	core "github.com/koblas/grpc-todo/gen/core/todo/v1/todov1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"go.uber.org/zap"
)

type Config struct {
	TodoServiceAddr string
	JwtSecret       string `validate:"min=32"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{
		TodoServiceAddr: ":13005",
	}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []todo.Option{
		todo.WithTodoService(
			core.NewTodoServiceClient(
				awsutil.NewTwirpCallLambda(),
				"lambda://core-todo",
			),
		),
		todo.WithGetUserId(authHelper),
	}

	_, api := todov1connect.NewTodoServiceHandler(
		todo.NewTodoServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

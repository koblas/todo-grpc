package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config todo.Config
	if err := confmgr.Parse(&config, aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []todo.Option{
		todo.WithTodoService(
			corev1connect.NewTodoServiceClient(
				awsutil.NewTwirpCallLambda(),
				"lambda://core-todo",
			),
		),
		todo.WithGetUserId(authHelper),
	}

	_, api := apiv1connect.NewTodoServiceHandler(
		todo.NewTodoServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

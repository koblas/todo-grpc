package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
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
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []todo.Option{
		todo.WithTodoService(
			corev1connect.NewTodoServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.TodoServiceAddr,
			),
		),
		todo.WithGetUserId(authHelper),
	}

	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewTodoServiceHandler(
		todo.NewTodoServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(apiv1connect.TodoServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

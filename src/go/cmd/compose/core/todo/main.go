package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/todo"
	"go.uber.org/zap"
)

type Config struct {
	NatsAddr string
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)
	eventbus := corev1connect.NewTodoEventbusServiceClient(
		nats,
		"topic:",
	)

	opts := []todo.Option{
		todo.WithProducer(eventbus),
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
	}

	mux := http.NewServeMux()
	mux.Handle(corev1connect.NewTodoServiceHandler(
		todo.NewTodoServer(opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(corev1connect.TodoServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

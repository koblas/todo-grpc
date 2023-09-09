package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	apiv1connect "github.com/koblas/grpc-todo/gen/api/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"

	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/message"

	"go.uber.org/zap"
)

type Config struct {
	MessageServiceAddr string
	JwtSecret          string `validate:"min=32"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{
		MessageServiceAddr: "core-message",
	}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []message.Option{
		message.WithMessageService(
			messagev1connect.NewMessageServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.MessageServiceAddr,
			),
		),
		message.WithGetUserId(authHelper),
	}

	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewMessageServiceHandler(
		message.NewMessageServer(opts...),
		bufcutil.WithJSON(),
		connect.WithInterceptors(
			interceptors.NewReqidInterceptor(),
			interceptors.NewDelayInterceptor(),
			auth,
		),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(apiv1connect.MessageServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

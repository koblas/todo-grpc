package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/gpt/v1/gptv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/gpt"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret string `validate:"min=32"`
	GptApiKey string `validate:"min=2"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []gpt.Option{
		gpt.WithGetUserId(authHelper),
		gpt.WithGpiApiKey(config.GptApiKey),
	}

	mux := http.NewServeMux()
	mux.Handle(gptv1connect.NewGptServiceHandler(
		gpt.NewGptServer(opts...),
		bufcutil.WithJSON(),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(gptv1connect.GptServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

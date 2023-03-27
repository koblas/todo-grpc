package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/api/v1/apiv1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
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
	if err := confmgr.ParseWithContext(mgr.Context(), &config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []gpt.Option{
		gpt.WithGetUserId(authHelper),
		gpt.WithGpiApiKey(config.GptApiKey),
	}

	_, api := apiv1connect.NewGptServiceHandler(
		gpt.NewGptServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

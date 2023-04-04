package main

import (
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/api/message/v1/messagev1connect"
	core "github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
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
		MessageServiceAddr: ":13005",
	}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []message.Option{
		message.WithMessageService(
			core.NewMessageServiceClient(
				awsutil.NewTwirpCallLambda(),
				"lambda://core-message",
			),
		),
		message.WithGetUserId(authHelper),
	}

	_, api := messagev1connect.NewMessageServiceHandler(
		message.NewMessageServer(opts...),
		connect.WithCodec(bufcutil.NewJsonCodec()),
		connect.WithInterceptors(interceptors.NewReqidInterceptor(), auth),
	)

	mgr.Start(awsutil.HandleApiLambda(api))
}

package main

import (
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/api/user/v1/userv1connect"
	cuser "github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/publicapi/user"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret       string `validate:"min=32"`
	UserServiceAddr string
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	auth, authHelper := interceptors.NewAuthInterceptor(config.JwtSecret)

	opts := []user.Option{
		user.WithUserService(
			cuser.NewUserServiceClient(
				bufcutil.NewHttpClient(),
				"http://"+config.UserServiceAddr,
			),
		),
		user.WithGetUserId(authHelper),
	}

	svc := user.NewUserServer(opts...)
	mux := http.NewServeMux()

	userPrefix, userSvc := userv1connect.NewUserServiceHandler(
		svc,
		bufcutil.WithJSON(),
		connect.WithInterceptors(
			interceptors.NewReqidInterceptor(),
			interceptors.NewDelayInterceptor(),
			auth,
		),
		connect.WithCompressMinBytes(1024),
	)
	teamPrefix, teamSvc := userv1connect.NewTeamServiceHandler(
		svc,
		bufcutil.WithJSON(),
		connect.WithInterceptors(
			interceptors.NewReqidInterceptor(),
			interceptors.NewDelayInterceptor(),
			auth,
		),
		connect.WithCompressMinBytes(1024),
	)

	mux.Handle(bufcutil.RewriteMux("/api/v1/user/", userPrefix, userSvc))
	mux.Handle(bufcutil.RewriteMux("/api/v1/team/", teamPrefix, teamSvc))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(userv1connect.UserServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

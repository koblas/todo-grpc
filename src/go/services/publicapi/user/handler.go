package user

import (
	"github.com/go-playground/validator/v10"
	apiv1 "github.com/koblas/grpc-todo/gen/api/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/interceptors"
)

var validate = validator.New()

// Server represents the gRPC server
type UserServer struct {
	user       userv1connect.UserServiceClient
	userHelper interceptors.UserIdFromContext
}

var _ apiv1.UserServiceHandler = (*UserServer)(nil)
var _ apiv1.TeamServiceHandler = (*UserServer)(nil)

type Option func(*UserServer)

func WithUserService(client userv1connect.UserServiceClient) Option {
	return func(svr *UserServer) {
		svr.user = client
	}
}

func WithGetUserId(helper interceptors.UserIdFromContext) Option {
	return func(svr *UserServer) {
		svr.userHelper = helper
	}
}

func NewUserServer(opts ...Option) *UserServer {
	svr := UserServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.user == nil {
		panic("no user service provided")
	}
	if svr.userHelper == nil {
		panic("no user helper provided")
	}

	return &svr
}

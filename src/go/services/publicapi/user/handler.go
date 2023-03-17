package user

import (
	"errors"
	"log"

	"github.com/bufbuild/connect-go"
	"github.com/go-playground/validator/v10"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var validate = validator.New()

// Server represents the gRPC server
type UserServer struct {
	user     corev1connect.UserServiceClient
	jwtMaker tokenmanager.Maker
}

type Option func(*UserServer)

func WithUserService(client corev1connect.UserServiceClient) Option {
	return func(svr *UserServer) {
		svr.user = client
	}
}

func NewUserServer(config Config, opts ...Option) *UserServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := UserServer{
		jwtMaker: maker,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *UserServer) getUserId(ctx context.Context) (string, error) {
	return tokenmanager.UserIdFromContext(ctx, svc.jwtMaker)
}

// SayHello generates response to a Ping request
func (svc *UserServer) GetUser(ctx context.Context, _ *connect.Request[apiv1.GetUserRequest]) (*connect.Response[apiv1.GetUserResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("GetUser BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}
	log = log.With("userId", userId)
	log.Info("Looking up user")

	user, err := svc.user.FindBy(ctx, connect.NewRequest(&corev1.FindByRequest{
		FindBy: &corev1.FindBy{
			UserId: userId,
		},
	}))

	if err != nil {
		log.With(zap.Error(err)).Info("lookup failed")
		if connect.CodeOf(err) == connect.CodeNotFound {
			return nil, connect.NewError(connect.CodeNotFound, nil)
		}
		return nil, bufcutil.InternalError(err)
	}

	// return connect.NewResponse(&apipbv1.GetUserResponse{
	// User: protoutil.UserCoreToApi(user.User),
	// }), nil
	return connect.NewResponse(&apiv1.GetUserResponse{
		User: protoutil.UserCoreToApi(user.Msg.User),
	}), nil
}

func (svc *UserServer) UpdateUser(ctx context.Context, updateIn *connect.Request[apiv1.UpdateUserRequest]) (*connect.Response[apiv1.UpdateUserResponse], error) {
	update := updateIn.Msg
	log := logger.FromContext(ctx)
	log.Info("UserUpdate BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With(zap.Error(err)).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	found, err := svc.user.FindBy(ctx, connect.NewRequest(&corev1.FindByRequest{
		FindBy: &corev1.FindBy{
			UserId: userId,
		},
	}))
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	if found == nil {
		return nil, bufcutil.InternalError(nil, "unable to locate user")
	}

	if update.Email != nil {
		if err := validate.Var(update.Email, "email"); err != nil {
			return nil, bufcutil.InvalidArgumentError("email", "bad format")
		}
	}
	if update.PasswordNew != nil {
		if update.Password == nil || *update.Password == "" {
			return nil, bufcutil.InvalidArgumentError("password_new", "current password is required")
		}
		if err := validate.Var(update.PasswordNew, "min=8"); err != nil {
			return nil, bufcutil.InvalidArgumentError("password_new", "must be 8 characters or more")
		}
	}

	user, err := svc.user.Update(ctx, connect.NewRequest(&corev1.UserServiceUpdateRequest{
		UserId:      userId,
		Name:        update.Name,
		Email:       update.Email,
		Password:    update.Password,
		PasswordNew: update.PasswordNew,
	}))
	if err != nil {
		log.With(zap.Error(err)).Error("user update failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.UpdateUserResponse{
		User: protoutil.UserCoreToApi(user.Msg.User),
	}), nil
}

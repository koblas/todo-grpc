package user

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/user/v1"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// SayHello generates response to a Ping request
func (svc *UserServer) GetUser(ctx context.Context, _ *connect.Request[apiv1.GetUserRequest]) (*connect.Response[apiv1.GetUserResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("GetUser BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}
	log = log.With("userId", userId)
	log.Info("Looking up user")

	user, err := svc.user.FindBy(ctx, connect.NewRequest(&userv1.FindByRequest{
		FindBy: &userv1.FindBy{
			UserId: userId,
		},
	}))

	if err != nil {
		log.With(zap.Error(err)).Info("lookup failed")
		if connect.CodeOf(err) == connect.CodeNotFound {
			// NotFound is valid error, translate this to not-authenticated since
			// it means they don't have a valid account
			return nil, connect.NewError(connect.CodeUnauthenticated, nil)
		}
		return nil, bufcutil.InternalError(err)
	}

	// TODO -- we should consider status in this case and potentially de-authenticate

	return connect.NewResponse(&apiv1.GetUserResponse{
		User: protoutil.UserCoreToApi(user.Msg.User),
	}), nil
}

func (svc *UserServer) UpdateUser(ctx context.Context, updateIn *connect.Request[apiv1.UpdateUserRequest]) (*connect.Response[apiv1.UpdateUserResponse], error) {
	update := updateIn.Msg
	log := logger.FromContext(ctx)
	log.Info("UserUpdate BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With(zap.Error(err)).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	found, err := svc.user.FindBy(ctx, connect.NewRequest(&userv1.FindByRequest{
		FindBy: &userv1.FindBy{
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

	user, err := svc.user.Update(ctx, connect.NewRequest(&userv1.UpdateRequest{
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

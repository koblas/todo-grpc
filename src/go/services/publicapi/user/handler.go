package user

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var validate = validator.New()

// Server represents the gRPC server
type UserServer struct {
	user     core.UserService
	jwtMaker tokenmanager.Maker
}

type Option func(*UserServer)

func WithUserService(client core.UserService) Option {
	return func(svr *UserServer) {
		svr.user = client
	}
}

func NewUserServer(config SsmConfig, opts ...Option) *UserServer {
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
	headers, ok := ctx.Value(awsutil.HeaderCtxKey).(http.Header)
	if !ok {
		if ctx.Value(awsutil.HeaderCtxKey) != nil {
			log.Println("Headers are present")
		}
		return "", fmt.Errorf("headers not in context")
	}

	value := headers.Get("authorization")
	if value == "" {
		return "", fmt.Errorf("no authorization header")
	}
	parts := strings.Split(value, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("bad format")
	}

	payload, err := svc.jwtMaker.VerifyToken(parts[1])
	if err != nil {
		return "", err
	}
	if payload.UserId == "" {
		return "", fmt.Errorf("no user_id")
	}

	return payload.UserId, nil
}

// SayHello generates response to a Ping request
func (svc *UserServer) GetUser(ctx context.Context, _ *publicapi.UserGetParams) (*publicapi.UserResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GetUser BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)
	log.Info("Looking up user")

	user, err := svc.user.FindBy(ctx, &core.UserFindParam{
		UserId: userId,
	})

	if err != nil {
		log.With(zap.Error(err)).Info("lookup failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return marshalUser(user), nil
}

func (svc *UserServer) UpdateUser(ctx context.Context, update *publicapi.UserUpdateParams) (*publicapi.UserResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("UserUpdate BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With(zap.Error(err)).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	user, err := svc.user.FindBy(ctx, &core.UserFindParam{
		UserId: userId,
	})
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if user == nil {
		return nil, twirp.InternalError("unable to locate user")
	}

	if update.Email != nil {
		if err := validate.Var(update.Email, "email"); err != nil {
			return nil, twirp.InvalidArgumentError("email", "bad format")
		}
	}
	if update.PasswordNew != nil {
		if update.Password == nil || *update.Password == "" {
			return nil, twirp.InvalidArgumentError("password_new", "current password is required")
		}
		if err := validate.Var(update.PasswordNew, "min=8"); err != nil {
			return nil, twirp.InvalidArgumentError("password_new", "must be 8 characters or more")
		}
	}

	user, err = svc.user.Update(ctx, &core.UserUpdateParam{
		UserId:      userId,
		Name:        update.Name,
		Email:       update.Email,
		Password:    update.Password,
		PasswordNew: update.PasswordNew,
	})
	if err != nil {
		log.With(zap.Error(err)).Error("user update failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return marshalUser(user), nil
}

func marshalUser(user *core.User) *publicapi.UserResponse {
	if user == nil {
		return &publicapi.UserResponse{}
	}

	return &publicapi.UserResponse{
		User: &publicapi.User{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
		},
	}
}

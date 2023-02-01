package user

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var validate = validator.New()

// Server represents the gRPC server
type UserServer struct {
	user     corepbv1.UserService
	jwtMaker tokenmanager.Maker
}

type Option func(*UserServer)

func WithUserService(client corepbv1.UserService) Option {
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
func (svc *UserServer) GetUser(ctx context.Context, _ *apipbv1.UserGetParams) (*apipbv1.UserResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GetUser BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)
	log.Info("Looking up user")

	user, err := svc.user.FindBy(ctx, &corepbv1.UserFindParam{
		UserId: userId,
	})

	if err != nil {
		log.With(zap.Error(err)).Info("lookup failed")
		var twerr twirp.Error
		if errors.As(err, &twerr) {
			if twerr.Code() == twirp.NotFound {
				return nil, twirp.NotFoundError("lookup failed")
			}
		}
		return nil, twirp.InternalErrorWith(err)
	}

	return marshalUser(user), nil
}

func (svc *UserServer) UpdateUser(ctx context.Context, update *apipbv1.UserUpdateParams) (*apipbv1.UserResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("UserUpdate BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With(zap.Error(err)).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	user, err := svc.user.FindBy(ctx, &corepbv1.UserFindParam{
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

	user, err = svc.user.Update(ctx, &corepbv1.UserUpdateParam{
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

func marshalUser(user *corepbv1.User) *apipbv1.UserResponse {
	return &apipbv1.UserResponse{
		User: protoutil.UserCoreToApi(user),
	}
}

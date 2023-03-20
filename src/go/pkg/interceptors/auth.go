package interceptors

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
)

type userIdCtxKeyType string

const UserIdCtxKey userIdCtxKeyType = "userId"

var (
	ErrAuthMissingToken = errors.New("no token provided")
	ErrAuthBadFormat    = errors.New("bad format")
	ErrAuthInvalidToken = errors.New("invalid token")
	ErrAuthMissingData  = errors.New("missing required data")
)

type UserIdFromContext interface {
	GetUserId(ctx context.Context) (string, error)
}

type userIdExtractor struct {
}

var _ UserIdFromContext = (*userIdExtractor)(nil)

func (*userIdExtractor) GetUserId(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(UserIdCtxKey).(string)
	if !ok {
		return "", connect.NewError(connect.CodeUnauthenticated, nil)
	}
	return userId, nil
}

func GetUserIdFromHeaders(ctx context.Context, maker tokenmanager.Maker, req connect.AnyRequest) (string, error) {
	log := logger.FromContext(ctx)
	value := req.Header().Get("authorization")
	if value == "" {
		log.Info("authentication failed missing header")
		return "", connect.NewError(connect.CodeUnauthenticated, ErrAuthMissingToken)
	}

	parts := strings.Split(value, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		log.Info("authentication failed header malformed")
		return "", connect.NewError(connect.CodeUnauthenticated, ErrAuthBadFormat)
	}

	payload, err := maker.VerifyToken(parts[1])
	if err != nil {
		log.Info("authentication token is not valid")
		return "", connect.NewError(connect.CodeUnauthenticated, ErrAuthInvalidToken)
	}
	if payload.UserId == "" {
		log.Info("no user ID present in token")
		return "", connect.NewError(connect.CodeUnauthenticated, ErrAuthMissingData)
	}
	log.Info("authentication successful")

	return payload.UserId, nil
}

func NewAuthInterceptor(tokenSecret string) (connect.UnaryInterceptorFunc, UserIdFromContext) {
	maker, err := tokenmanager.NewJWTMaker(tokenSecret)
	if err != nil {
		panic(err)
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			userId, err := GetUserIdFromHeaders(ctx, maker, req)
			if err != nil {
				return nil, err
			}

			return next(context.WithValue(ctx, UserIdCtxKey, userId), req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor), &userIdExtractor{}
}

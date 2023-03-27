package auth

import (
	"log"
	"time"

	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"
)

const MAX_LOGIN_ATTEMPS = 5
const LOGIN_LOCKOUT_MINUTES = 15

// Server represents the gRPC server
type AuthenticationServer struct {
	// apiv1.UnimplementedAuthenticationServiceServer

	jwtMaker    tokenmanager.Maker
	userClient  corev1connect.UserServiceClient
	oauthClient corev1connect.AuthUserServiceClient
	attempts    AttemptService
}

type Option func(*AuthenticationServer)

func WithUserClient(client corev1connect.UserServiceClient) Option {
	return func(input *AuthenticationServer) {
		input.userClient = client
	}
}

func WithOAuthClient(client corev1connect.AuthUserServiceClient) Option {
	return func(input *AuthenticationServer) {
		input.oauthClient = client
	}
}

func WithAttemptService(client AttemptService) Option {
	return func(input *AuthenticationServer) {
		input.attempts = client
	}
}

func NewAuthenticationServer(jwtSecret string, opts ...Option) AuthenticationServer {
	maker, err := tokenmanager.NewJWTMaker(jwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := AuthenticationServer{
		jwtMaker: maker,
		attempts: NewAttemptsStub(),
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return svr
}

func (s AuthenticationServer) returnToken(ctx context.Context, userId string) (*apiv1.Token, error) {
	// TODO: This is an authentication event that should be pushed onto
	//  the messagebus

	bearer, err := s.jwtMaker.CreateToken(userId, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &apiv1.Token{
		AccessToken: bearer,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}

package auth

import (
	"log"
	"time"

	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"
)

const MAX_LOGIN_ATTEMPS = 5
const LOGIN_LOCKOUT_MINUTES = 15

// Server represents the gRPC server
type AuthenticationServer struct {
	// apipb.UnimplementedAuthenticationServiceServer

	jwtMaker    tokenmanager.Maker
	userClient  corepb.UserService
	oauthClient corepb.AuthUserService
	attempts    AttemptService
}

type Option func(*AuthenticationServer)

func WithUserClient(client corepb.UserService) Option {
	return func(input *AuthenticationServer) {
		input.userClient = client
	}
}

func WithOAuthClient(client corepb.AuthUserService) Option {
	return func(input *AuthenticationServer) {
		input.oauthClient = client
	}
}

func WithAttemptService(client AttemptService) Option {
	return func(input *AuthenticationServer) {
		input.attempts = client
	}
}

func NewAuthenticationServer(config Config, opts ...Option) AuthenticationServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
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

func (s AuthenticationServer) returnToken(ctx context.Context, userId string) (*apipb.Token, error) {
	// TODO: This is an authentication event that should be pushed onto
	//  the messagebus

	bearer, err := s.jwtMaker.CreateToken(userId, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &apipb.Token{
		AccessToken: bearer,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}

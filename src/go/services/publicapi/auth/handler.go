package auth

import (
	"log"
	"time"

	"github.com/koblas/grpc-todo/twpb/core"

	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"golang.org/x/net/context"
)

const MAX_LOGIN_ATTEMPS = 5
const LOGIN_LOCKOUT_MINUTES = 15

// Server represents the gRPC server
type AuthenticationServer struct {
	// publicapi.UnimplementedAuthenticationServiceServer

	jwtMaker    tokenmanager.Maker
	userClient  core.UserService
	oauthClient core.AuthUserService
	attempts    AttemptService
}

func NewAuthenticationServer(config SsmConfig, userClient core.UserService, oauthClient core.AuthUserService, attempts AttemptService) AuthenticationServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	return AuthenticationServer{
		userClient:  userClient,
		oauthClient: oauthClient,
		jwtMaker:    maker,
		attempts:    attempts,
	}
}

func (s AuthenticationServer) returnToken(ctx context.Context, userId string) (*publicapi.Token, error) {
	// TODO: This is an authentication event that should be pushed onto
	//  the messagebus

	bearer, err := s.jwtMaker.CreateToken(userId, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.Token{
		AccessToken: bearer,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}

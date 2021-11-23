package auth

import (
	"log"
	"os"
	"time"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"
)

const MAX_LOGIN_ATTEMPS = 5
const LOGIN_LOCKOUT_MINUTES = 15

// Server represents the gRPC server
type AuthenticationServer struct {
	publicapi.UnimplementedAuthenticationServiceServer

	jwtMaker    tokenmanager.Maker
	userClient  core.UserServiceClient
	oauthClient core.OauthUserServiceClient
	attempts    AttemptService
}

func NewAuthenticationServer(userClient core.UserServiceClient, oauthClient core.OauthUserServiceClient, attempts AttemptService) AuthenticationServer {
	maker, err := tokenmanager.NewJWTMaker(os.Getenv("JWT_SECRET"))
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

func (s AuthenticationServer) returnToken(ctx context.Context, user *core.User) (*publicapi.Token, error) {
	// TODO: This is an authentication event that should be pushed onto
	//  the messagebus

	bearer, err := s.jwtMaker.CreateToken(user.Id, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.Token{
		AccessToken: bearer,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}

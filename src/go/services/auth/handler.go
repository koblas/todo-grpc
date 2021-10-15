package auth

import (
	"log"

	"github.com/koblas/grpc-todo/genpb"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type AuthenticationServer struct {
	genpb.UnimplementedAuthenticationServiceServer
}

func (s *AuthenticationServer) Login(ctx context.Context, params *genpb.LoginParams) (*genpb.TokenResponse, error) {
	log.Printf("Received login %s", params.Username)

	response := genpb.TokenResponse{
		AccessToken: "abc",
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}
	return &response, nil
}

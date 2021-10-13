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

// SayHello generates response to a Ping request
func (s *AuthenticationServer) Login(ctx context.Context, params *genpb.LoginParams) (*genpb.LoginResponse, error) {
	log.Printf("Received login %s", params.Username)

	response := genpb.LoginResponse{}
	return &response, nil
}

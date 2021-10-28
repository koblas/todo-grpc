package middleware

import (
	"context"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"google.golang.org/grpc/metadata"
)

type authMiddleware struct {
	manager tokenmanager.Maker
}

func NewAuthenticator(secret string) MiddlewareProvider {
	manager, err := tokenmanager.NewJWTMaker(secret)
	if err != nil {
		log.Fatalf("Unable to create manager %v", err)
	}

	return authMiddleware{manager}
}

func (m authMiddleware) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(m.authenticator)
}

func (m authMiddleware) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_auth.StreamServerInterceptor(m.authenticator)
}

// exampleAuthFunc is used by a middleware to authenticate requests
func (s authMiddleware) authenticator(ctx context.Context) (context.Context, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	log.Printf("Metadata = %+v", md)

	// if val == "grpc.health.v1.Health" {
	// 	return ctx, nil
	// }

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	payload, err := s.manager.VerifyToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	// grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	// newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)
	newCtx := context.WithValue(ctx, "user_id", payload.UserId)

	return newCtx, nil
}

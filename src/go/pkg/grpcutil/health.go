package grpcutil

import (
	"google.golang.org/grpc/health"

	"golang.org/x/net/context"
)

type HealthNoAuth struct {
	*health.Server
}

func NewServer() *HealthNoAuth {
	return &HealthNoAuth{health.NewServer()}
}

// Disable Authentication for this
func (s *HealthNoAuth) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

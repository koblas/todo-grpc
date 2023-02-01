package workers_user

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

type EmptyServer struct{}

func (*EmptyServer) UserChange(ctx context.Context, msg *corepbv1.UserChangeEvent) (*corepbv1.EventbusEmpty, error) {
	return &corepbv1.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityPasswordChange(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	return &corepbv1.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityForgotRequest(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	return &corepbv1.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityRegisterToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	return &corepbv1.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityInviteToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	return &corepbv1.EventbusEmpty{}, nil
}

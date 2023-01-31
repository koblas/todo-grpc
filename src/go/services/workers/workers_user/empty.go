package workers_user

import (
	"context"

	"github.com/koblas/grpc-todo/gen/corepb"
)

type EmptyServer struct{}

func (*EmptyServer) UserChange(ctx context.Context, msg *corepb.UserChangeEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityPasswordChange(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityForgotRequest(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityRegisterToken(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*EmptyServer) SecurityInviteToken(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}

package workers_user

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

type EmptyServer struct{}

func (*EmptyServer) UserChange(ctx context.Context, msg *corepbv1.UserChangeEvent) (*corepbv1.UserEventbusUserChangeResponse, error) {
	return &corepbv1.UserEventbusUserChangeResponse{}, nil
}
func (*EmptyServer) SecurityPasswordChange(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityPasswordChangeResponse, error) {
	return &corepbv1.UserEventbusSecurityPasswordChangeResponse{}, nil
}
func (*EmptyServer) SecurityForgotRequest(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityForgotRequestResponse, error) {
	return &corepbv1.UserEventbusSecurityForgotRequestResponse{}, nil
}
func (*EmptyServer) SecurityRegisterToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityRegisterTokenResponse, error) {
	return &corepbv1.UserEventbusSecurityRegisterTokenResponse{}, nil
}
func (*EmptyServer) SecurityInviteToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityInviteTokenResponse, error) {
	return &corepbv1.UserEventbusSecurityInviteTokenResponse{}, nil
}

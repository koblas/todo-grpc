package workers_user

import (
	"context"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
)

type EmptyServer struct{}

func (*EmptyServer) UserChange(ctx context.Context, msg *connect.Request[corev1.UserChangeEvent]) (*connect.Response[corev1.UserEventbusUserChangeResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusUserChangeResponse{}), nil
}
func (*EmptyServer) SecurityPasswordChange(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityPasswordChangeResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityPasswordChangeResponse{}), nil
}
func (*EmptyServer) SecurityForgotRequest(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityForgotRequestResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityForgotRequestResponse{}), nil
}
func (*EmptyServer) SecurityRegisterToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityRegisterTokenResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityRegisterTokenResponse{}), nil
}
func (*EmptyServer) SecurityInviteToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityInviteTokenResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityInviteTokenResponse{}), nil
}

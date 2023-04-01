package workers_user

import (
	"context"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
)

type EmptyServer struct{}

func (*EmptyServer) UserChange(ctx context.Context, msg *connect.Request[corev1.UserChangeEvent]) (*connect.Response[eventv1.UserEventbusUserChangeResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil
}
func (*EmptyServer) SecurityPasswordChange(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityPasswordChangeResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityPasswordChangeResponse{}), nil
}
func (*EmptyServer) SecurityForgotRequest(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityForgotRequestResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityForgotRequestResponse{}), nil
}
func (*EmptyServer) SecurityRegisterToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityRegisterTokenResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil
}
func (*EmptyServer) SecurityInviteToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityInviteTokenResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityInviteTokenResponse{}), nil
}

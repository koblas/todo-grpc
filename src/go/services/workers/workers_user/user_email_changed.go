package workers_user

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

func init() {
	workers = append(workers, Worker{
		Stream:    "event:user_security",
		GroupName: "userSecurity/password_changed",
		Build:     NewUserEmailChanged,
	})
}

type userEmailChanged struct {
	EmptyServer
	WorkerConfig
}

func NewUserEmailChanged(config WorkerConfig) http.Handler {
	svc := &userEmailChanged{WorkerConfig: config}

	_, api := corev1connect.NewUserEventbusServiceHandler(svc)

	return api
}

func (cfg *userEmailChanged) SecurityPasswordChange(ctx context.Context, msg *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityPasswordChangeResponse], error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Msg.Action))).With(zap.String("email", msg.Msg.User.Email))

	log.Info("processing message")
	user := msg.Msg.User

	params := corev1.PasswordChangeMessageRequest{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corev1.EmailUser{
			UserId: user.Id,
			Name:   user.Name,
			Email:  user.Email,
		},
	}

	log.Info("Sending password change email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.PasswordChangeMessage(ctx, connect.NewRequest(&params)); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, bufcutil.InternalError(err, "failed to send")
		}
	}

	return connect.NewResponse(&corev1.UserEventbusSecurityPasswordChangeResponse{}), nil
}

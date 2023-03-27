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
		GroupName: "userSecurity/forgot",
		Build:     NewUserEmailForgot,
	})
}

type userEmailForgot struct {
	EmptyServer
	WorkerConfig
}

func NewUserEmailForgot(config WorkerConfig) http.Handler {
	svc := &userEmailForgot{WorkerConfig: config}

	_, api := corev1connect.NewUserEventbusServiceHandler(svc)
	return api
}

func (cfg *userEmailForgot) SecurityForgotRequest(ctx context.Context, msgIn *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityForgotRequestResponse], error) {
	msg := msgIn.Msg
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return connect.NewResponse(&corev1.UserEventbusSecurityForgotRequestResponse{}), err
	}

	params := corev1.PasswordRecoveryMessageRequest{
		AppInfo: buildAppInfo(cfg.config.UrlBase),
		Recipient: &corev1.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending forgot email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.PasswordRecoveryMessage(ctx, connect.NewRequest(&params)); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, bufcutil.InternalError(err, "failed to send")
		}
	}

	return connect.NewResponse(&corev1.UserEventbusSecurityForgotRequestResponse{}), err
}

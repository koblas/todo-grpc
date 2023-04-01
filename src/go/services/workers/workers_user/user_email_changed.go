package workers_user

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	emailv1 "github.com/koblas/grpc-todo/gen/core/send_email/v1"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
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

	_, api := eventbusv1connect.NewUserEventbusServiceHandler(svc)

	return api
}

func (cfg *userEmailChanged) SecurityPasswordChange(ctx context.Context, msg *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityPasswordChangeResponse], error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Msg.Action))).With(zap.String("email", msg.Msg.User.Email))

	log.Info("processing message")
	user := msg.Msg.User

	params := emailv1.PasswordChangeMessageRequest{
		AppInfo: buildAppInfo(cfg.config.UrlBase),
		Recipient: &emailv1.EmailUser{
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

	return connect.NewResponse(&eventv1.UserEventbusSecurityPasswordChangeResponse{}), nil
}

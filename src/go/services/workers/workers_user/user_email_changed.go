package workers_user

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
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

func NewUserEmailChanged(config WorkerConfig) corepbv1.TwirpServer {
	svc := &userEmailChanged{WorkerConfig: config}

	return corepbv1.NewUserEventbusServer(svc)
}

func (cfg *userEmailChanged) SecurityPasswordChange(ctx context.Context, msg *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepbv1.UserSecurity_USER_PASSWORD_CHANGE {
		return &corepbv1.EventbusEmpty{}, nil
	}
	user := msg.User

	params := corepbv1.EmailPasswordChangeParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepbv1.EmailUser{
			UserId: user.Id,
			Name:   user.Name,
			Email:  user.Email,
		},
	}

	log.Info("Sending password change email")
	if cfg.sendEmail != nil {
		_, err := cfg.sendEmail.PasswordChangeMessage(ctx, &params)
		log.With(zap.Error(err)).Info("Failed to send")
		return nil, err
	}

	return &corepbv1.EventbusEmpty{}, nil
}

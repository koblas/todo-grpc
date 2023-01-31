package workers_user

import (
	"context"

	"github.com/koblas/grpc-todo/gen/corepb"
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

func NewUserEmailChanged(config WorkerConfig) corepb.TwirpServer {
	svc := &userEmailChanged{WorkerConfig: config}

	return corepb.NewUserEventbusServer(svc)
}

func (cfg *userEmailChanged) SecurityPasswordChange(ctx context.Context, msg *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepb.UserSecurity_USER_PASSWORD_CHANGE {
		return &corepb.EventbusEmpty{}, nil
	}
	user := msg.User

	params := corepb.EmailPasswordChangeParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepb.EmailUser{
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

	return &corepb.EventbusEmpty{}, nil
}

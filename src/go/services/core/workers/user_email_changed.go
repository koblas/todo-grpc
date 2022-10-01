package workers

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
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
	WorkerConfig
}

func NewUserEmailChanged(config WorkerConfig) genpb.TwirpServer {
	svc := &userEmailChanged{WorkerConfig: config}

	return genpb.NewUserEventServiceServer(svc)
}

func (cfg *userEmailChanged) UserChange(ctx context.Context, msg *genpb.UserChangeEvent) (*genpb.EventbusEmpty, error) {
	return &genpb.EventbusEmpty{}, nil
}

func (cfg *userEmailChanged) UserSecurity(ctx context.Context, msg *genpb.UserSecurityEvent) (*genpb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != genpb.UserSecurity_USER_PASSWORD_CHANGE {
		return &genpb.EventbusEmpty{}, nil
	}
	user := msg.User

	params := genpb.EmailPasswordChangeParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &genpb.EmailUser{
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

	return &genpb.EventbusEmpty{}, nil
}

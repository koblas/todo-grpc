package workers_user

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/logger"
	twcore "github.com/koblas/grpc-todo/twpb/core"
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
	WorkerConfig
}

func NewUserEmailForgot(config WorkerConfig) twcore.TwirpServer {
	svc := &userEmailForgot{WorkerConfig: config}

	return twcore.NewUserEventbusServer(svc)
}

func (cfg *userEmailForgot) UserChange(ctx context.Context, msg *twcore.UserChangeEvent) (*twcore.EventbusEmpty, error) {
	return &twcore.EventbusEmpty{}, nil
}

func (cfg *userEmailForgot) UserSecurity(ctx context.Context, msg *twcore.UserSecurityEvent) (*twcore.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != twcore.UserSecurity_USER_FORGOT_REQUEST {
		return &twcore.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return &twcore.EventbusEmpty{}, err
	}

	params := twcore.EmailPasswordRecoveryParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &twcore.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending forgot email")
	if cfg.sendEmail != nil {
		_, err := cfg.sendEmail.PasswordRecoveryMessage(ctx, &params)
		log.With(zap.Error(err)).Info("Failed to send")
		return nil, err
	}

	return &twcore.EventbusEmpty{}, err
}

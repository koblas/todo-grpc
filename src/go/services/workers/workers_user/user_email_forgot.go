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
		GroupName: "userSecurity/forgot",
		Build:     NewUserEmailForgot,
	})
}

type userEmailForgot struct {
	WorkerConfig
}

func NewUserEmailForgot(config WorkerConfig) corepb.TwirpServer {
	svc := &userEmailForgot{WorkerConfig: config}

	return corepb.NewUserEventbusServer(svc)
}

func (cfg *userEmailForgot) UserChange(ctx context.Context, msg *corepb.UserChangeEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}

func (cfg *userEmailForgot) UserSecurity(ctx context.Context, msg *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepb.UserSecurity_USER_FORGOT_REQUEST {
		return &corepb.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return &corepb.EventbusEmpty{}, err
	}

	params := corepb.EmailPasswordRecoveryParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepb.EmailUser{
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

	return &corepb.EventbusEmpty{}, err
}

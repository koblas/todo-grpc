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
		GroupName: "userSecurity/forgot",
		Build:     NewUserEmailForgot,
	})
}

type userEmailForgot struct {
	WorkerConfig
}

func NewUserEmailForgot(config WorkerConfig) genpb.TwirpServer {
	svc := &userEmailForgot{WorkerConfig: config}

	return genpb.NewUserEventServiceServer(svc)
}

func (cfg *userEmailForgot) UserChange(ctx context.Context, msg *genpb.UserChangeEvent) (*genpb.EventbusEmpty, error) {
	return &genpb.EventbusEmpty{}, nil
}

func (cfg *userEmailForgot) UserSecurity(ctx context.Context, msg *genpb.UserSecurityEvent) (*genpb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != genpb.UserSecurity_USER_FORGOT_REQUEST {
		return &genpb.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	params := genpb.EmailPasswordRecoveryParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &genpb.EmailUser{
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

	return &genpb.EventbusEmpty{}, err
}

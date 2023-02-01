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
		GroupName: "userSecurity/forgot",
		Build:     NewUserEmailForgot,
	})
}

type userEmailForgot struct {
	EmptyServer
	WorkerConfig
}

func NewUserEmailForgot(config WorkerConfig) corepbv1.TwirpServer {
	svc := &userEmailForgot{WorkerConfig: config}

	return corepbv1.NewUserEventbusServer(svc)
}

func (cfg *userEmailForgot) SecurityForgotRequest(ctx context.Context, msg *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepbv1.UserSecurity_USER_FORGOT_REQUEST {
		return &corepbv1.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return &corepbv1.EventbusEmpty{}, err
	}

	params := corepbv1.EmailPasswordRecoveryParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepbv1.EmailUser{
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

	return &corepbv1.EventbusEmpty{}, err
}

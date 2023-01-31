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
		GroupName: "userSecurity/register",
		Build:     NewUserEmailConfirm,
	})
}

type userEmailConfirm struct {
	EmptyServer
	WorkerConfig
}

func NewUserEmailConfirm(config WorkerConfig) corepb.TwirpServer {
	svc := &userEmailConfirm{WorkerConfig: config}

	return corepb.NewUserEventbusServer(svc)
}

func (cfg *userEmailConfirm) SecurityRegisterToken(ctx context.Context, msg *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepb.UserSecurity_USER_REGISTER_TOKEN {
		return &corepb.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	log = log.With("email", msg.User.Email)

	params := corepb.EmailRegisterParam{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepb.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending registration email")
	if cfg.sendEmail != nil {
		_, err := cfg.sendEmail.RegisterMessage(ctx, &params)
		log.With(zap.Error(err)).Info("Failed to send")
		return nil, err
	}

	return &corepb.EventbusEmpty{}, err
}

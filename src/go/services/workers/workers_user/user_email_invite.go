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
		GroupName: "userSecurity/invite",
		Build:     NewUserEmailInvite,
	})
}

type userEmailInvite struct {
	WorkerConfig
}

func NewUserEmailInvite(config WorkerConfig) corepb.TwirpServer {
	svc := &userEmailInvite{WorkerConfig: config}

	return corepb.NewUserEventbusServer(svc)
}

func (cfg *userEmailInvite) UserChange(ctx context.Context, msg *corepb.UserChangeEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}

func (cfg *userEmailInvite) UserSecurity(ctx context.Context, msg *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepb.UserSecurity_USER_INVITE_TOKEN {
		return &corepb.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	params := corepb.EmailInviteUserParam{
		AppInfo: buildAppInfo(cfg.config),
		Sender: &corepb.EmailUser{
			UserId: msg.Sender.Id,
			Name:   msg.Sender.Name,
			Email:  msg.Sender.Email,
		},
		Recipient: &corepb.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending registration email")
	if cfg.sendEmail != nil {
		_, err := cfg.sendEmail.InviteUserMessage(ctx, &params)
		log.With(zap.Error(err)).Info("Failed to send")
		return nil, err
	}

	return &corepb.EventbusEmpty{}, nil
}

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
		GroupName: "userSecurity/invite",
		Build:     NewUserEmailInvite,
	})
}

type userEmailInvite struct {
	WorkerConfig
}

func NewUserEmailInvite(config WorkerConfig) genpb.TwirpServer {
	svc := &userEmailInvite{WorkerConfig: config}

	return genpb.NewUserEventServiceServer(svc)
}

func (cfg *userEmailInvite) UserChange(ctx context.Context, msg *genpb.UserChangeEvent) (*genpb.EventbusEmpty, error) {
	return &genpb.EventbusEmpty{}, nil
}

func (cfg *userEmailInvite) UserSecurity(ctx context.Context, msg *genpb.UserSecurityEvent) (*genpb.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != genpb.UserSecurity_USER_INVITE_TOKEN {
		return &genpb.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	params := genpb.EmailInviteUserParam{
		AppInfo: buildAppInfo(cfg.config),
		Sender: &genpb.EmailUser{
			UserId: msg.Sender.Id,
			Name:   msg.Sender.Name,
			Email:  msg.Sender.Email,
		},
		Recipient: &genpb.EmailUser{
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

	return &genpb.EventbusEmpty{}, nil
}

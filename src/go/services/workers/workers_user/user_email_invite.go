package workers_user

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/twitchtv/twirp"
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
	EmptyServer
	WorkerConfig
}

func NewUserEmailInvite(config WorkerConfig) corepbv1.TwirpServer {
	svc := &userEmailInvite{WorkerConfig: config}

	return corepbv1.NewUserEventbusServer(svc)
}

func (cfg *userEmailInvite) SecurityInviteToken(ctx context.Context, msg *corepbv1.UserSecurityEvent) (*corepbv1.EventbusEmpty, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepbv1.UserSecurity_USER_INVITE_TOKEN {
		return &corepbv1.EventbusEmpty{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	params := corepbv1.EmailInviteUserParam{
		AppInfo: buildAppInfo(cfg.config),
		Sender: &corepbv1.EmailUser{
			UserId: msg.Sender.Id,
			Name:   msg.Sender.Name,
			Email:  msg.Sender.Email,
		},
		Recipient: &corepbv1.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending registration email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.InviteUserMessage(ctx, &params); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, twirp.WrapError(twirp.InternalError("failed to send"), err)
		}
	}

	return &corepbv1.EventbusEmpty{}, nil
}

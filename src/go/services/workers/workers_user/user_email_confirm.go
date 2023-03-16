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
		GroupName: "userSecurity/register",
		Build:     NewUserEmailConfirm,
	})
}

type userEmailConfirm struct {
	EmptyServer
	WorkerConfig
}

func NewUserEmailConfirm(config WorkerConfig) corepbv1.TwirpServer {
	svc := &userEmailConfirm{WorkerConfig: config}

	return corepbv1.NewUserEventbusServiceServer(svc)
}

func (cfg *userEmailConfirm) SecurityRegisterToken(ctx context.Context, msg *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityRegisterTokenResponse, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepbv1.UserSecurity_USER_SECURITY_USER_REGISTER_TOKEN {
		return &corepbv1.UserEventbusSecurityRegisterTokenResponse{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	log = log.With("email", msg.User.Email)

	params := corepbv1.RegisterMessageRequest{
		AppInfo: buildAppInfo(cfg.config),
		Recipient: &corepbv1.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending registration email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.RegisterMessage(ctx, &params); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, twirp.WrapError(twirp.InternalError("failed to send"), err)
		}
	}

	return &corepbv1.UserEventbusSecurityRegisterTokenResponse{}, err
}

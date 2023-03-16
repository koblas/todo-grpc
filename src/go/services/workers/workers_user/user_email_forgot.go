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

	return corepbv1.NewUserEventbusServiceServer(svc)
}

func (cfg *userEmailForgot) SecurityForgotRequest(ctx context.Context, msg *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityForgotRequestResponse, error) {
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")
	if msg.Action != corepbv1.UserSecurity_USER_SECURITY_USER_FORGOT_REQUEST {
		return &corepbv1.UserEventbusSecurityForgotRequestResponse{}, nil
	}

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return &corepbv1.UserEventbusSecurityForgotRequestResponse{}, err
	}

	params := corepbv1.PasswordRecoveryMessageRequest{
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
		if _, err := cfg.sendEmail.PasswordRecoveryMessage(ctx, &params); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, twirp.WrapError(twirp.InternalError("failed to send"), err)
		}
	}

	return &corepbv1.UserEventbusSecurityForgotRequestResponse{}, err
}

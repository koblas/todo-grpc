package workers_user

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
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
	EmptyServer
	WorkerConfig
}

func NewUserEmailInvite(config WorkerConfig) http.Handler {
	svc := &userEmailInvite{WorkerConfig: config}

	_, api := corev1connect.NewUserEventbusServiceHandler(svc)
	return api
}

func (cfg *userEmailInvite) SecurityInviteToken(ctx context.Context, msgIn *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityInviteTokenResponse], error) {
	msg := msgIn.Msg
	log := logger.FromContext(ctx).With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	params := corev1.InviteUserMessageRequest{
		AppInfo: buildAppInfo(cfg.config.UrlBase),
		Sender: &corev1.EmailUser{
			UserId: msg.Sender.Id,
			Name:   msg.Sender.Name,
			Email:  msg.Sender.Email,
		},
		Recipient: &corev1.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending invitation email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.InviteUserMessage(ctx, connect.NewRequest(&params)); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, bufcutil.InternalError(err, "failed to send")
		}
	}

	return connect.NewResponse(&corev1.UserEventbusSecurityInviteTokenResponse{}), nil
}

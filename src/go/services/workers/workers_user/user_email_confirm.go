package workers_user

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
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

func NewUserEmailConfirm(config WorkerConfig) http.Handler {
	svc := &userEmailConfirm{WorkerConfig: config}

	_, api := eventbusv1connect.NewUserEventbusServiceHandler(svc)
	return api
}

func (cfg *userEmailConfirm) SecurityRegisterToken(ctx context.Context, msgIn *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityRegisterTokenResponse], error) {
	log := logger.FromContext(ctx)
	if msgIn.Msg.User == nil {
		log.With(zap.Any("msg", msgIn.Msg)).Info("RAW MESSAGE")
	}
	msg := msgIn.Msg
	log = log.With(zap.Int32("action", int32(msg.Action))).With(zap.String("email", msg.User.Email))

	log.Info("processing message")

	tokenValue, err := decodeSecure(log, msg.Token)
	if err != nil {
		return nil, err
	}

	log = log.With("email", msg.User.Email)

	params := corev1.RegisterMessageRequest{
		AppInfo: buildAppInfo(cfg.config.UrlBase),
		Recipient: &corev1.EmailUser{
			UserId: msg.User.Id,
			Name:   msg.User.Name,
			Email:  msg.User.Email,
		},
		Token: tokenValue,
	}

	log.Info("Sending registration email")
	if cfg.sendEmail != nil {
		if _, err := cfg.sendEmail.RegisterMessage(ctx, connect.NewRequest(&params)); err != nil {
			log.With(zap.Error(err)).Info("Failed to send")
			return nil, bufcutil.InternalError(err, "failed to send")
		}
	}

	return connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), err
}

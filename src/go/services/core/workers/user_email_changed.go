package workers

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	awsbus "github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
)

func init() {
	workers = append(workers, Worker{
		Stream:    "event:user_security",
		GroupName: "userSecurity/password_changed",
		Build:     NewUserEmailChanged,
	})
}

func NewUserEmailChanged(config *SsmConfig) awsbus.SqsConsumerFunc {
	return config.userEmailChanged
}

func (config *SsmConfig) userEmailChanged(ctx context.Context, msg *eventbus.Message) error {
	log := logger.FromContext(ctx)
	event := genpb.UserSecurityEvent{}
	action, err := extractBasic(log, msg, &event)
	if err != nil {
		log.With("error", err).Info("Unable to extract message")
		return err
	}
	log.With("action", action).Info("processing message")
	cuser := event.User
	if event.Action != genpb.UserSecurity_USER_PASSWORD_CHANGE {
		return nil
	}

	params := genpb.EmailPasswordChangeParam{
		AppInfo: buildAppInfo(config),
		Recipient: &genpb.EmailUser{
			UserId: cuser.Id,
			Name:   cuser.Name,
			Email:  cuser.Email,
		},
	}

	email, err := getEmailService(log)
	if err != nil {
		log.With("email", cuser.Email, "error", err).Info("Failed to send")
		return err
	}
	log.With("email", cuser.Email).Info("Sending registration email")
	_, err = email.PasswordChangeMessage(ctx, &params)

	if err != nil {
		log.With("email", cuser.Email, "error", err).Info("Failed to send")
	}

	return err
}

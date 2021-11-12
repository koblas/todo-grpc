package workers

import (
	"context"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/robinjoseph08/redisqueue"
)

func init() {
	workers = append(workers, Worker{
		Stream:    "event:user_security",
		GroupName: "userSecurity/invite",
		Process: func(ctx context.Context, msg *redisqueue.Message) error {
			log := logger.FromContext(ctx)
			event := genpb.UserSecurityEvent{}
			action, err := extractBasic(log, msg, &event)
			if err != nil {
				log.With("error", err).Info("Unable to extract message")
				return err
			}
			log.With("action", action).Info("processing message")
			cuser := event.User
			if event.Action != genpb.UserSecurity_USER_INVITE_TOKEN || event.Token == "" {
				return nil
			}

			params := genpb.EmailInviteUserParam{
				AppInfo: appInfo,
				Sender: &genpb.EmailUser{
					Name:  event.Sender.Name,
					Email: event.Sender.Email,
				},
				Recipient: &genpb.EmailUser{
					Name:  cuser.Name,
					Email: cuser.Email,
				},
				Token: event.Token,
			}

			email, err := getEmailService(log)
			if err != nil {
				log.With("email", cuser.Email, "error", err).Info("Failed to send")
				return err
			}
			log.With("email", cuser.Email).Info("Sending registration email")
			_, err = email.InviteUserMessage(ctx, &params)

			if err != nil {
				log.With("email", cuser.Email, "error", err).Info("Failed to send")
			}

			return err
		},
	})
}

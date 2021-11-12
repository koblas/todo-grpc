package workers

import (
	"context"
	"log"
	"os"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/robinjoseph08/redisqueue"
)

func init() {
	urlBase := os.Getenv("URL_BASE_UI")
	if urlBase == "" {
		log.Fatal("environment variable URL_BASE_UI must be set")
	}

	workers = append(workers, Worker{
		Stream:    "event:user_security",
		GroupName: "userSecurity/password_changed",
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
			if event.Action != genpb.UserSecurity_USER_PASSWORD_CHANGE || event.Token == "" {
				return nil
			}

			params := genpb.EmailPasswordChangeParam{
				AppInfo: appInfo,
				Recipient: &genpb.EmailUser{
					Name:  cuser.Name,
					Email: cuser.Email,
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
		},
	})
}

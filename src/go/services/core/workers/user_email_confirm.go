package workers

import (
	"context"
	"errors"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/robinjoseph08/redisqueue"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func extractBasic(log logger.Logger, msg *redisqueue.Message, m protoreflect.ProtoMessage) (string, error) {
	var action string

	metadata := genpb.Metadata{}
	badPayload := errors.New("Bad Payload")

	// Decode the metadata
	if iface, ok := msg.Values["metadata"]; !ok {
		log.Info("No meta data in payload")
		return action, badPayload
	} else if bdata, ok := iface.(string); !ok {
		log.Info("No meta data as string")
		return action, badPayload
	} else if err := proto.Unmarshal([]byte(bdata), &metadata); err != nil {
		return action, err
	} else {
		for _, item := range metadata.Metadata {
			if item.Key == "action" {
				action = item.Value
				break
			}
		}
		if action == "" {
			return action, badPayload
		}
	}

	// Decode the body
	if iface, ok := msg.Values["body"]; !ok {
		log.Info("No body in payload")
		return action, badPayload
	} else if bdata, ok := iface.(string); !ok {
		log.Info("No body as string")
		return action, badPayload
	} else if err := proto.Unmarshal([]byte(bdata), m); err != nil {
		return action, err
	}

	return action, nil
}

func init() {
	workers = append(workers, Worker{
		Stream:    "entity:user",
		GroupName: "userEmailConfirm",
		Process: func(ctx context.Context, msg *redisqueue.Message) error {
			log := logger.FromContext(ctx)
			user := genpb.UserChangeEvent{}
			action, err := extractBasic(log, msg, &user)
			if err != nil {
				log.With("error", err).Info("Unable to extract message")
				return err
			}
			log.With("action", action).Info("processing message")
			cuser := user.Current
			if action != "created" || cuser == nil {
				return nil
			}

			params := genpb.EmailRegisterParam{
				AppInfo: &genpb.EmailAppInfo{
					UrlBase: "",
					AppName: "TestApp",
				},
				Sender: &genpb.EmailUser{
					Name:  "Test David",
					Email: "david@koblas.com",
				},
				Recipient: &genpb.EmailUser{
					Name:  cuser.Name,
					Email: cuser.Email,
				},
				Token: cuser.VerificationToken,
			}

			email, err := getEmailService(log)
			if err != nil {
				log.With("email", cuser.Email, "error", err).Info("Failed to send")
				return err
			}
			log.With("email", cuser.Email).Info("Sending registration email")
			_, err = email.RegisterMessage(ctx, &params)

			if err != nil {
				log.With("email", cuser.Email, "error", err).Info("Failed to send")
			}

			return err
		},
	})
}

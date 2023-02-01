package send_email

import (
	"bytes"
	"context"
	"fmt"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/manager"
)

type SendEmailServer struct {
	sender Sender
	pubsub corepbv1.SendEmailEvents
}

// This is really hear to make it easy to make sure that you've
//
//	tied the correct event to the template that will be sent
var templates map[corepbv1.EmailTemplate]emailContent = map[corepbv1.EmailTemplate]emailContent{
	corepbv1.EmailTemplate_USER_REGISTERED:   registerUser,
	corepbv1.EmailTemplate_USER_INVITED:      inviteUser,
	corepbv1.EmailTemplate_PASSWORD_CHANGE:   passwordChange,
	corepbv1.EmailTemplate_PASSWORD_RECOVERY: passwordRecovery,
}

type Handler struct {
	handler corepbv1.TwirpServer
}

func (h Handler) GroupName() string {
	return "send_email"
}

func (h Handler) Handler() corepbv1.TwirpServer {
	return h.handler
}

func NewSendEmailServer(producer corepbv1.SendEmailEvents, sender Sender) []manager.MsgHandler {
	// pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
	// 	StreamMaxLength:      1000,
	// 	ApproximateMaxLength: true,
	// 	RedisOptions: &redisqueue.RedisOptions{
	// 		Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
	// 	},
	// })
	// if err != nil {
	// 	logger.With(err).Fatal("unable to start producer")
	// }

	return []manager.MsgHandler{
		Handler{
			handler: corepbv1.NewSendEmailServiceServer(&SendEmailServer{
				pubsub: producer,
				sender: sender,
			}),
		},
	}
}

func (s *SendEmailServer) RegisterMessage(ctx context.Context, params *corepbv1.EmailRegisterParam) (*corepbv1.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Id":    params.Recipient.UserId,
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, corepbv1.EmailTemplate_USER_REGISTERED, params.ReferenceId); err != nil {
		return nil, err
	}

	return &corepbv1.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) PasswordChangeMessage(ctx context.Context, params *corepbv1.EmailPasswordChangeParam) (*corepbv1.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Id":    params.Recipient.UserId,
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, corepbv1.EmailTemplate_PASSWORD_CHANGE, params.ReferenceId); err != nil {
		return nil, err
	}

	return &corepbv1.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) PasswordRecoveryMessage(ctx context.Context, params *corepbv1.EmailPasswordRecoveryParam) (*corepbv1.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Id":    params.Recipient.UserId,
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, corepbv1.EmailTemplate_PASSWORD_RECOVERY, params.ReferenceId); err != nil {
		return nil, err
	}

	return &corepbv1.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) InviteUserMessage(ctx context.Context, params *corepbv1.EmailInviteUserParam) (*corepbv1.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Id":    params.Recipient.UserId,
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"Sender": map[string]string{
			"Id":    params.Sender.UserId,
			"Email": params.Sender.Email,
			"Name":  params.Sender.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, corepbv1.EmailTemplate_USER_INVITED, params.ReferenceId); err != nil {
		return nil, err
	}

	return &corepbv1.EmailOkResponse{Ok: true}, nil
}

// One stop shop to send a message
func (svc *SendEmailServer) simpleSend(ctx context.Context, sender, recipient string, data Params, tmpl corepbv1.EmailTemplate, referenceId string) error {
	content, found := templates[tmpl]
	if !found {
		return fmt.Errorf("unable to find template id=%d name=%s", tmpl, corepbv1.EmailTemplate_name[int32(tmpl)])
	}

	subject, body, err := buildEmail(data, content)
	if err != nil {
		return err
	}

	messageId, err := svc.sender.SendEmail(ctx, sender, recipient, subject, body)
	if err != nil {
		return err
	}

	svc.notify(ctx, messageId, recipient, tmpl, referenceId)

	return nil
}

func (svc *SendEmailServer) notify(ctx context.Context, messageId string, recipient string, tmpl corepbv1.EmailTemplate, referenceId string) error {
	_, err := svc.pubsub.NotifySent(ctx, &corepbv1.EmailSentEvent{
		RecipientEmail: recipient,
		MessageId:      messageId,
		Template:       tmpl,
		ReferenceId:    referenceId,
	})

	return err
}

// Common functionality to build and email message from a template
func buildEmail(data map[string]interface{}, content emailContent) (string, string, error) {
	var subject bytes.Buffer
	var body bytes.Buffer

	if err := content.subject.Execute(&subject, data); err != nil {
		return "", "", err
	}
	if err := content.body.Execute(&body, data); err != nil {
		return "", "", err
	}

	return subject.String(), body.String(), nil
}

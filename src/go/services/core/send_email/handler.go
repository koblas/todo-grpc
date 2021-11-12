package send_email

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/robinjoseph08/redisqueue"
	"google.golang.org/protobuf/proto"
)

type SendEmailServer struct {
	genpb.UnimplementedSendEmailServiceServer

	sender Sender
	pubsub *redisqueue.Producer
}

// This is really hear to make it easy to make sure that you've
//  tied the correct event to the template that will be sent
var templates map[genpb.EmailTemplate]emailContent = map[genpb.EmailTemplate]emailContent{
	genpb.EmailTemplate_USER_REGISTERED:   registerUser,
	genpb.EmailTemplate_USER_INVITED:      inviteUser,
	genpb.EmailTemplate_PASSWORD_CHANGE:   passwordChange,
	genpb.EmailTemplate_PASSWORD_RECOVERY: passwordRecovery,
}

func NewSendEmailServer(logger logger.Logger, sender Sender) *SendEmailServer {
	pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
		},
	})
	if err != nil {
		logger.With(err).Fatal("unable to start producer")
	}

	return &SendEmailServer{
		pubsub: pubsub,
		sender: sender,
	}
}

func (s *SendEmailServer) RegisterMessage(ctx context.Context, params *genpb.EmailRegisterParam) (*genpb.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, genpb.EmailTemplate_USER_REGISTERED, params.ReferenceId); err != nil {
		return nil, err
	}

	return &genpb.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) PasswordChangeMessage(ctx context.Context, params *genpb.EmailPasswordChangeParam) (*genpb.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, genpb.EmailTemplate_PASSWORD_CHANGE, params.ReferenceId); err != nil {
		return nil, err
	}

	return &genpb.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) PasswordRecoveryMessage(ctx context.Context, params *genpb.EmailPasswordRecoveryParam) (*genpb.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, genpb.EmailTemplate_PASSWORD_RECOVERY, params.ReferenceId); err != nil {
		return nil, err
	}

	return &genpb.EmailOkResponse{Ok: true}, nil
}

func (s *SendEmailServer) InviteUserMessage(ctx context.Context, params *genpb.EmailInviteUserParam) (*genpb.EmailOkResponse, error) {
	recipient := params.Recipient.Email
	sender := params.Recipient.Email
	data := Params{
		"User": map[string]string{
			"Email": params.Recipient.Email,
			"Name":  params.Recipient.Name,
		},
		"Sender": map[string]string{
			"Email": params.Sender.Email,
			"Name":  params.Sender.Name,
		},
		"AppName": params.AppInfo.AppName,
		"URLBase": params.AppInfo.UrlBase,
		"Token":   params.Token,
	}

	if err := s.simpleSend(ctx, sender, recipient, data, genpb.EmailTemplate_USER_INVITED, params.ReferenceId); err != nil {
		return nil, err
	}

	return &genpb.EmailOkResponse{Ok: true}, nil
}

// One stop shop to send a message
func (svc *SendEmailServer) simpleSend(ctx context.Context, sender, recipient string, data Params, tmpl genpb.EmailTemplate, referenceId string) error {
	content, found := templates[tmpl]
	if !found {
		return fmt.Errorf("unable to find template id=%d name=%s", tmpl, genpb.EmailTemplate_name[int32(tmpl)])
	}

	subject, body, err := buildEmail(data, content.subject, content.body)
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

func (svc *SendEmailServer) notify(ctx context.Context, messageId string, recipient string, tmpl genpb.EmailTemplate, referenceId string) error {
	stream := "email_messages"

	values := []*genpb.MetadataEntry{
		{Key: "stream", Value: stream},
		{Key: "action", Value: "sent"},
	}
	mbytes, err := proto.Marshal(&genpb.Metadata{
		Metadata: values,
	})
	if err != nil {
		return err
	}
	bbytes, err := proto.Marshal(&genpb.EmailSentEvent{
		RecipientEmail: recipient,
		MessageId:      messageId,
		Template:       tmpl,
		ReferenceId:    referenceId,
	})
	if err != nil {
		return err
	}

	return svc.pubsub.Enqueue(&redisqueue.Message{
		Stream: stream,
		Values: map[string]interface{}{
			"metadata": mbytes,
			"body":     bbytes,
		},
	})
}

// Common functionality to build and email message from a template
func buildEmail(data map[string]interface{}, subjectTmpl, bodyTmpl string) (string, string, error) {
	funcMap := template.FuncMap{
		/*
			"Currency": func(v interface{}) string {
				ac := accounting.Accounting{Symbol: "$", Precision: 2}

				return ac.FormatMoney(v)
			},
		*/
		"Round": func(v interface{}) string {
			return fmt.Sprintf("%.2f", v)
		},
	}

	var doc bytes.Buffer

	// Create the subject
	tmpl, err := template.New("subject").Funcs(funcMap).Parse(subjectTmpl)
	if err != nil {
		return "", "", err
	}

	doc.Reset()
	if err := tmpl.Execute(&doc, data); err != nil {
		return "", "", err
	}

	subject := doc.String()

	// Create the body
	tmpl, err = Layout().Funcs(funcMap).Parse(bodyTmpl)
	if err != nil {
		return "", "", err
	}

	doc.Reset()
	if err := tmpl.Execute(&doc, data); err != nil {
		return "", "", err
	}

	body := doc.String()

	return subject, body, nil
}

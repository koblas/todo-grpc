package send_email

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	eventbusv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	emailv1 "github.com/koblas/grpc-todo/gen/core/send_email/v1"
	"github.com/koblas/grpc-todo/gen/core/send_email/v1/send_emailv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
)

type SendEmailServer struct {
	sender Sender
	pubsub eventbusv1connect.SendEmailEventsServiceClient
}

// This is really hear to make it easy to make sure that you've
//
//	tied the correct event to the template that will be sent
var templates map[emailv1.EmailTemplate]emailContent = map[emailv1.EmailTemplate]emailContent{
	emailv1.EmailTemplate_EMAIL_TEMPLATE_USER_REGISTERED:   registerUser,
	emailv1.EmailTemplate_EMAIL_TEMPLATE_USER_INVITED:      inviteUser,
	emailv1.EmailTemplate_EMAIL_TEMPLATE_PASSWORD_CHANGE:   passwordChange,
	emailv1.EmailTemplate_EMAIL_TEMPLATE_PASSWORD_RECOVERY: passwordRecovery,
}

func NewSendEmailServer(producer eventbusv1connect.SendEmailEventsServiceClient, sender Sender) map[string]http.Handler {
	_, api := send_emailv1connect.NewSendEmailServiceHandler(
		NewSendEmailServerServer(producer, sender),
	)

	return map[string]http.Handler{"queue.send_email": api}
}

func NewSendEmailServerServer(producer eventbusv1connect.SendEmailEventsServiceClient, sender Sender) *SendEmailServer {
	server := SendEmailServer{
		pubsub: producer,
		sender: sender,
	}

	return &server

}

func (s *SendEmailServer) RegisterMessage(ctx context.Context, req *connect.Request[emailv1.RegisterMessageRequest]) (*connect.Response[emailv1.RegisterMessageResponse], error) {
	params := req.Msg
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

	if err := s.simpleSend(ctx, sender, recipient, data, emailv1.EmailTemplate_EMAIL_TEMPLATE_USER_REGISTERED, params.ReferenceId); err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&emailv1.RegisterMessageResponse{}), nil
}

func (s *SendEmailServer) PasswordChangeMessage(ctx context.Context, req *connect.Request[emailv1.PasswordChangeMessageRequest]) (*connect.Response[emailv1.PasswordChangeMessageResponse], error) {
	params := req.Msg
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

	if err := s.simpleSend(ctx, sender, recipient, data, emailv1.EmailTemplate_EMAIL_TEMPLATE_PASSWORD_CHANGE, params.ReferenceId); err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&emailv1.PasswordChangeMessageResponse{}), nil
}

func (s *SendEmailServer) PasswordRecoveryMessage(ctx context.Context, req *connect.Request[emailv1.PasswordRecoveryMessageRequest]) (*connect.Response[emailv1.PasswordRecoveryMessageResponse], error) {
	params := req.Msg
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

	if err := s.simpleSend(ctx, sender, recipient, data, emailv1.EmailTemplate_EMAIL_TEMPLATE_PASSWORD_RECOVERY, params.ReferenceId); err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&emailv1.PasswordRecoveryMessageResponse{}), nil
}

func (s *SendEmailServer) InviteUserMessage(ctx context.Context, req *connect.Request[emailv1.InviteUserMessageRequest]) (*connect.Response[emailv1.InviteUserMessageResponse], error) {
	params := req.Msg
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

	if err := s.simpleSend(ctx, sender, recipient, data, emailv1.EmailTemplate_EMAIL_TEMPLATE_USER_INVITED, params.ReferenceId); err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&emailv1.InviteUserMessageResponse{}), nil
}

// One stop shop to send a message
func (svc *SendEmailServer) simpleSend(ctx context.Context, sender, recipient string, data Params, tmpl emailv1.EmailTemplate, referenceId string) error {
	content, found := templates[tmpl]
	if !found {
		return fmt.Errorf("unable to find template id=%d name=%s", tmpl, emailv1.EmailTemplate_name[int32(tmpl)])
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

func (svc *SendEmailServer) notify(ctx context.Context, messageId string, recipient string, tmpl emailv1.EmailTemplate, referenceId string) error {
	_, err := svc.pubsub.NotifyEmailSent(ctx, connect.NewRequest(&eventbusv1.NotifyEmailSentRequest{
		RecipientEmail: recipient,
		MessageId:      messageId,
		Template:       tmpl,
		ReferenceId:    referenceId,
	}))

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

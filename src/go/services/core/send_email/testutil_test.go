package send_email_test

import (
	"context"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/services/core/send_email"
)

type stubSender struct {
	subject string
	body    string
}

func (svc *stubSender) SendEmail(ctx context.Context, sender, to, subject, html string) (string, error) {
	svc.subject = subject
	svc.body = html

	return "", nil
}

type stubBus struct {
}

func (svc *stubBus) NotifyEmailSent(context.Context, *connect.Request[eventv1.NotifyEmailSentRequest]) (*connect.Response[eventv1.NotifyEmailSentResponse], error) {
	return nil, nil
}

func buildTestService() (*send_email.SendEmailServer, *stubSender) {
	senderData := &stubSender{}
	bus := &stubBus{}
	svc := send_email.NewSendEmailServerServer(bus, senderData)

	return svc, senderData
}

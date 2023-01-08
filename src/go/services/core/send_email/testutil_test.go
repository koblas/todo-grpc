package send_email_test

import (
	"context"

	"github.com/koblas/grpc-todo/gen/corepb"
	email "github.com/koblas/grpc-todo/services/core/send_email"
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

func (svc *stubBus) NotifySent(context.Context, *corepb.EmailSentEvent) (*corepb.EmailOkResponse, error) {
	return nil, nil
}

func buildTestService() (corepb.SendEmailService, *stubSender) {
	senderData := &stubSender{}
	bus := &stubBus{}
	svc := email.NewSendEmailServer(bus, senderData)

	return svc, senderData
}

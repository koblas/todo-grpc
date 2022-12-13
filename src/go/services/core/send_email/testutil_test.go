package send_email_test

import (
	"context"

	email "github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
	genpb "github.com/koblas/grpc-todo/twpb/core"
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

func (svc *stubBus) NotifySent(context.Context, *core.EmailSentEvent) (*core.EmailOkResponse, error) {
	return nil, nil
}

func buildTestService() (genpb.SendEmailService, *stubSender) {
	senderData := &stubSender{}
	bus := &stubBus{}
	svc := email.NewSendEmailServer(bus, senderData)

	return svc, senderData
}
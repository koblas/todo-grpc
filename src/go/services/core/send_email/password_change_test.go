package send_email_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	eventstub "github.com/koblas/grpc-todo/pkg/eventbus/stub"
	email "github.com/koblas/grpc-todo/services/core/send_email"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"github.com/stretchr/testify/require"
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

func buildTestService() (genpb.SendEmailService, *stubSender) {
	senderData := &stubSender{}
	bus := eventstub.NewEventbusStub()
	svc := email.NewSendEmailServer(bus, senderData)

	return svc, senderData
}

func TestPasswordChange(t *testing.T) {
	faker := faker.New()

	svc, msgData := buildTestService()

	params := genpb.EmailPasswordChangeParam{
		Recipient: &genpb.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		AppInfo: &genpb.EmailAppInfo{
			AppName: faker.Company().Name(),
			UrlBase: faker.Internet().URL(),
		},
	}
	_, err := svc.PasswordChangeMessage(context.Background(), &params)

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, msgData.subject, "No subject")
	require.NotEmpty(t, msgData.body, "No body")

	require.Contains(t, msgData.body, params.Recipient.Name, "Mesage doesn't contain sender's firstname")
}

package send_email_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/stretchr/testify/require"
)

func TestPasswordChange(t *testing.T) {
	faker := faker.New()

	svc, msgData := buildTestService()

	params := corepb.EmailPasswordChangeParam{
		Recipient: &corepb.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		AppInfo: &corepb.EmailAppInfo{
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

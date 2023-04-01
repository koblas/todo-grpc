package send_email_test

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/go-faker/faker/v4"
	emailv1 "github.com/koblas/grpc-todo/gen/core/send_email/v1"
	"github.com/stretchr/testify/require"
)

func TestPasswordChange(t *testing.T) {
	svc, msgData := buildTestService()

	params := connect.NewRequest(&emailv1.PasswordChangeMessageRequest{
		Recipient: &emailv1.EmailUser{
			Name:  faker.Name(),
			Email: faker.Email(),
		},
		AppInfo: &emailv1.EmailAppInfo{
			AppName: faker.Name(),
			UrlBase: faker.URL(),
		},
	})
	_, err := svc.PasswordChangeMessage(context.Background(), params)

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, msgData.subject, "No subject")
	require.NotEmpty(t, msgData.body, "No body")

	require.Contains(t, msgData.body, params.Msg.Recipient.Name, "Mesage doesn't contain sender's firstname")
}

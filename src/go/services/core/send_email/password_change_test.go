package send_email_test

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/jaswdr/faker"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/stretchr/testify/require"
)

func TestPasswordChange(t *testing.T) {
	faker := faker.New()

	svc, msgData := buildTestService()

	params := connect.NewRequest(&corev1.PasswordChangeMessageRequest{
		Recipient: &corev1.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		AppInfo: &corev1.EmailAppInfo{
			AppName: faker.Company().Name(),
			UrlBase: faker.Internet().URL(),
		},
	})
	_, err := svc.PasswordChangeMessage(context.Background(), params)

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, msgData.subject, "No subject")
	require.NotEmpty(t, msgData.body, "No body")

	require.Contains(t, msgData.body, params.Msg.Recipient.Name, "Mesage doesn't contain sender's firstname")
}

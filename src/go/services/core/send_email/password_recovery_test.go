package send_email_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func TestPasswordRecovery(t *testing.T) {
	faker := faker.New()

	svc, msgData := buildTestService()

	params := corepbv1.EmailPasswordRecoveryParam{
		Recipient: &corepbv1.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		AppInfo: &corepbv1.EmailAppInfo{
			AppName: faker.Company().Name(),
			UrlBase: faker.Internet().URL(),
		},
		Token: faker.Hash().MD5(),
	}

	_, err := svc.PasswordRecoveryMessage(context.Background(), &params)

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, msgData.subject, "No subject")
	require.NotEmpty(t, msgData.body, "No body")

	require.Contains(t, msgData.body, params.AppInfo.UrlBase, "Mesage doesn't contain url")
	require.Contains(t, msgData.body, params.Token, "Mesage doesn't contain token")
	require.Contains(t, msgData.body, params.Recipient.Name, "Mesage doesn't contain sender's firstname")
}

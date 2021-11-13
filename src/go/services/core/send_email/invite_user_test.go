package send_email_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/stretchr/testify/require"
)

func TestInviteUser(t *testing.T) {
	faker := faker.New()

	svc, msgData := buildTestService()

	params := genpb.EmailInviteUserParam{
		Sender: &genpb.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		Recipient: &genpb.EmailUser{
			Name:  faker.Person().Name(),
			Email: faker.Internet().Email(),
		},
		AppInfo: &genpb.EmailAppInfo{
			AppName: faker.Company().Name(),
			UrlBase: faker.Internet().URL(),
		},
		Token: faker.Hash().MD5(),
	}

	_, err := svc.InviteUserMessage(context.Background(), &params)

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, msgData.subject, "No subject")
	require.NotEmpty(t, msgData.body, "No body")
	require.Contains(t, msgData.body, params.Token, "Mesage doesn't contain url")
	require.Contains(t, msgData.body, params.Recipient.Name, "Mesage doesn't contain recipients's name")
	require.Contains(t, msgData.body, params.Sender.Name, "Mesage doesn't contain sender's name")
}

package send_email_test

import (
	"context"
	"testing"

	"github.com/koblas/projectx/server-go/pkg/services/email"
	send "github.com/koblas/projectx/server-go/pkg/services/send_email"
	"github.com/stretchr/testify/require"
)

func TestInviteUser(t *testing.T) {
	senderData := &stubSender{}

	url := "THIS_IS_THE_URL_XXX"

	svc := email.NewService(senderData)

	user := send.EmailUser{}
	user.Email = "foo@example.com"
	user.Name = "Tom Smith"
	newUser := send.EmailUser{}
	newUser.Email = "bar@example.com"
	newUser.Name = "Mary Jane"

	err := svc.InviteUser(context.Background(), "test@example.com", send.Params{
		"User":      user,
		"NewUser":   newUser,
		"InviteUrl": url,
	})

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, senderData.subject, "No subject")
	require.NotEmpty(t, senderData.body, "No body")
	require.Contains(t, senderData.body, url, "Mesage doesn't contain url")
	require.Contains(t, senderData.body, user.Name, "Mesage doesn't contain sender's firstname")
}

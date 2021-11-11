package send_email_test

import (
	"context"
	"testing"

	"github.com/koblas/projectx/server-go/pkg/entity"
	"github.com/koblas/projectx/server-go/pkg/services/email"
	"github.com/stretchr/testify/require"
)

func TestPasswordRecovery(t *testing.T) {
	senderData := &stubSender{}
	svc := email.NewService(senderData)

	url := "THIS_IS_THE_URL_XXX"
	token := "QuickBrownFox"

	user := entity.User{}
	user.Email = "foo@example.com"
	user.Name = "John Smith"

	err := svc.PasswordRecovery(context.Background(), "foo@example.com", email.Params{
		"User":    user,
		"URLBase": url,
		"Token":   token,
	})

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, senderData.subject, "No subject")
	require.NotEmpty(t, senderData.body, "No body")

	require.Contains(t, senderData.body, url+token, "Mesage doesn't contain url")
	require.Contains(t, senderData.body, user.Name, "Mesage doesn't contain sender's firstname")
}

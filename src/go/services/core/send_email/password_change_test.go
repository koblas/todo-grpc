package send_email_test

import (
	"context"
	"testing"

	"github.com/koblas/projectx/server-go/pkg/entity"
	"github.com/koblas/projectx/server-go/pkg/services/email"
	"github.com/stretchr/testify/require"
)

type stubSender struct {
	subject string
	body    string
}

func (svc *stubSender) SendEmail(ctx context.Context, to, subject, body string) error {
	svc.subject = subject
	svc.body = body
	return nil
}

func TestPasswordChange(t *testing.T) {
	senderData := &stubSender{}

	svc := email.NewService(senderData)

	user := entity.User{}
	user.Email = "foo@example.com"
	user.Name = "John Smith"

	err := svc.PasswordChange(context.Background(), "text@example.com", email.Params{"User": user})

	require.Nil(t, err, "Failed to build")
	require.NotEmpty(t, senderData.subject, "No subject")
	require.NotEmpty(t, senderData.body, "No body")

	require.Contains(t, senderData.body, user.Name, "Mesage doesn't contain sender's firstname")
}

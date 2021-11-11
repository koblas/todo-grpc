package send_email

import (
	"context"
)

type EmailUser struct {
	Email string
	Name  string
}

type Params map[string]interface{}

type emailContent struct {
	subject string
	body    string
}

type Sender interface {
	SendEmail(ctx context.Context, sender, to, subject, body string) (string, error)
}

package send_email

import (
	"context"
	"html/template"
)

type EmailUser struct {
	Email string
	Name  string
}

type Params map[string]interface{}

type emailContent struct {
	subject *template.Template
	body    *template.Template
}

type Sender interface {
	SendEmail(ctx context.Context, sender, to, subject, body string) (string, error)
}

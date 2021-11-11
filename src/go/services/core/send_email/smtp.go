package send_email

import (
	"context"
	"os"

	"github.com/koblas/grpc-todo/pkg/logger"
	"gopkg.in/gomail.v2"
)

type smtpService struct {
}

// NewService construct a default email service
func NewSmtpService() Sender {
	return &smtpService{}
}

func (svc *smtpService) SendEmail(ctx context.Context, sender, to, subject, html string) (string, error) {
	log := logger.FromContext(ctx).With("to", to)

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	host := os.Getenv("SMTP_HOST")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	d := gomail.NewDialer(host, 587, username, password)
	log.With("stmpHost", host, "smtpUser", username).Info("Sending email")

	err := d.DialAndSend(m)

	if err != nil {
		log.With("error", err).Error("Failed to send message")
	}

	return "", err
}

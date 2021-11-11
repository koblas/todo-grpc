package util

import (
	"github.com/koblas/grpc-todo/pkg/logger"
	"gopkg.in/gomail.v2"
)

type SmtpServiceConfig struct {
	Host     string
	Port     int
	Sender   string
	Username string
	Password string
}

type EmailService interface {
	SendEmail(logger logger.Logger, to, subject, body string) error
}

type smtpService struct {
	config SmtpServiceConfig
}

// NewService construct a default email service
func NewSmtpService(config SmtpServiceConfig) EmailService {
	return &smtpService{
		config: config,
	}
}

func (svc *smtpService) SendEmail(logger logger.Logger, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "david@koblas.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	logger.With("to", to).Info("Sending email")

	d := gomail.NewDialer(
		"email-smtp.us-west-2.amazonaws.com", 587,
		svc.config.Username, svc.config.Password,
	)

	return d.DialAndSend(m)
}

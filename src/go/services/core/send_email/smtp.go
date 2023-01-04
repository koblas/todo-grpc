package send_email

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/koblas/grpc-todo/pkg/logger"
	"gopkg.in/gomail.v2"
)

type smtpService struct {
	host     string
	port     int
	username string
	password string
}

// NewService construct a default email service
func NewSmtpService(config Config) Sender {
	parts := strings.Split(config.SmtpAddr, ":")
	host := parts[0]
	port := 587
	if len(parts) == 2 {
		var err error
		port, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	return &smtpService{
		host:     host,
		port:     port,
		username: config.SmtpUser,
		password: config.SmtpPass,
	}
}

func (svc *smtpService) SendEmail(ctx context.Context, sender, to, subject, html string) (string, error) {
	log := logger.FromContext(ctx).With("to", to)

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(svc.host, svc.port, svc.username, svc.password)
	log.With("stmpHost", svc.host, "smtpUser", svc.username).Info("Sending email")

	err := d.DialAndSend(m)

	if err != nil {
		log.With("error", err).Error("Failed to send message")
	}

	return "", err
}

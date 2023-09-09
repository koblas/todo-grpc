package send_email

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type smtpService struct {
	host     string
	port     int
	username string
	password string
}

// NewService construct a default email service
func NewSmtpService(addr, user, pass string) Sender {
	parts := strings.Split(addr, ":")
	host := parts[0]
	port := 587
	if len(parts) == 2 {
		var err error
		port, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if host == "" {
		panic("No SMTP host configured")
	}

	return &smtpService{
		host:     host,
		port:     port,
		username: user,
		password: pass,
	}
}

func (svc *smtpService) SendEmail(ctx context.Context, sender, to, subject, html string) (string, error) {
	log := logger.FromContext(ctx).With(
		zap.String("to", to),
	).With(
		zap.Any("smtp", map[string]any{
			"host": svc.host,
			"port": svc.port,
			"user": svc.username,
		}),
	)

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(svc.host, svc.port, svc.username, svc.password)
	log.Info("Sending email")

	err := d.DialAndSend(m)

	if err != nil {
		log.With(zap.Error(err)).Error("Failed to send message")
	}

	return "", err
}

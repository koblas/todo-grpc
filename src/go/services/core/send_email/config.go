package send_email

type Config struct {
	// Used by Lambda
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	// Used by Kubernetes
	NatsAddr       string `environment:"NATS_ADDR"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	EmailSentTopic string `json:"send-email-events"`
	// Shared
	SmtpAddr string `ssm:"smtp/addr" environment:"SMTP_ADDR"`
	SmtpUser string `ssm:"smtp/username" environment:"SMTP_USERNAME"`
	SmtpPass string `ssm:"smtp/password" environment:"SMTP_PASSWORD"`
}

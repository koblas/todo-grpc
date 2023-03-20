package send_email

type Config struct {
	// Used by Lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by Kubernetes
	NatsAddr       string `environment:"NATS_ADDR"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	EmailSentTopic string `json:"send-email-events"`
	// Shared
}

type SmtpConfig struct {
	SmtpAddr string `ssm:"addr" environment:"SMTP_ADDR"`
	SmtpUser string `ssm:"username" environment:"SMTP_USERNAME"`
	SmtpPass string `ssm:"password" environment:"SMTP_PASSWORD"`
}

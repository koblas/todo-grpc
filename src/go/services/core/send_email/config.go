package send_email

type Config struct {
	EventArn       string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	SmtpAddr       string `ssm:"smtp/addr" environment:"SMTP_ADDR"`
	SmtpUser       string `ssm:"smtp/username" environment:"SMTP_USERNAME"`
	SmtpPass       string `ssm:"smtp/password" environment:"SMTP_PASSWORD"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	EmailSentTopic string `json:"send-email-events"`
}

package send_email

type SsmConfig struct {
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	SmtpAddr string `ssm:"smtp/addr" environment:"SMTP_ADDR"`
	SmtpUser string `ssm:"smtp/username" environment:"SMTP_USERNAME"`
	SmtpPass string `ssm:"smtp/password" environment:"SMTP_PASSWORD"`
}

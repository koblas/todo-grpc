package workers_user

type Config struct {
	// Used by lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by Kubernetes for the event source
	NatsAddr        string `environment:"NATS_ADDR"`
	RedisAddr       string `environment:"REDIS_ADDR"`
	UserEventsTopic string `json:"user-events"`

	// Shared
	UrlBase   string `ssm:"url_base" environment:"URL_BASE_UI"`
	SendEmail string `json:"send-email"`
}

package user

type Config struct {
	// Used by lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by Kubernetes for event source
	NatsAddr string `environment:"NATS_ADDR"`
	// or
	RedisAddr       string `environment:"REDIS_ADDR"`
	UserEventsTopic string `json:"user-events"`
	BroadcastTopic  string `json:"websocket-broadcast-events"`
	// Shared
}

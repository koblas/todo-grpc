package user

type Config struct {
	// Used by lambda
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	// Used by Kubernetes for event source
	NatsAddr string `environment:"NATS_ADDR"`
	// or
	RedisAddr       string `environment:"REDIS_ADDR"`
	UserEventsTopic string `json:"user-events"`
	// Shared
	BroadcastTopic string `json:"websocket-broadcast-events" validate:"required"`
}

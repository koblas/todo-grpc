package todo

type Config struct {
	// Used by lambda
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	// Used by Kubernetes for the event source
	RedisAddr       string `environment:"REDIS_ADDR"`
	TodoEventsTopic string `json:"todo-events"`
	// Shared
	BroadcastTopic string `json:"websocket-broadcast-events" validate:"required"`
}

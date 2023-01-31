package todo

type Config struct {
	// Used by lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by Kubernetes for the event source
	NatsAddr        string `environment:"NATS_ADDR"`
	RedisAddr       string `environment:"REDIS_ADDR"`
	TodoEventsTopic string `json:"todo-events"`
	BroadcastTopic  string `json:"websocket-broadcast-events"`
	// Shared
}

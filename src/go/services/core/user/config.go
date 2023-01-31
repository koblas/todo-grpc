package user

type Config struct {
	// Used by lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by kubernetes
	NatsAddr       string `environment:"NATS_ADDR"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	UserEventTopic string `json:"user-events"`
	// Shared
	DynamoStore string `environment:"DYNAMO_STORE_ADDR" json:"dynamo-store"`
}

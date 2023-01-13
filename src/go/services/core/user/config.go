package user

type Config struct {
	// Used by lambda
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	// Used by kubernetes
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	UserEventTopic string `json:"user-events"`
	// Shared
	DynamoStore string `environment:"DYNAMO_STORE_ADDR" json:"dynamo-store"`
}

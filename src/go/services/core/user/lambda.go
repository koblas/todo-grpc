package user

type SsmConfig struct {
	EventArn       string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	UserEventTopic string `json:"user-events"`
	DynamoStore    string `environment:"DYNAMO_STORE_ADDR" json:"dynamo-store"`
}

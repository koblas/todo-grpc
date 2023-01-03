package file

type Config struct {
	EventArn       string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	FileEventTopic string `json:"user-events"`
}

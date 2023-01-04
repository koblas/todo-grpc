package todo

type Config struct {
	RedisAddr  string `environment:"REDIS_ADDR" default:"redis:6379"`
	EventArn   string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	TodoEvents string `json:"todo-events"`
}

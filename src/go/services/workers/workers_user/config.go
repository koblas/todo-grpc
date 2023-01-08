package workers_user

type Config struct {
	UrlBase         string `ssm:"url_base" environment:"URL_BASE_UI"`
	UserEventsTopic string `json:"user-events"`
	SendEmail       string `json:"send-email"`
	RedisAddr       string `ssm:"redis_addr" json:"redis-addr" environment:"REDIS_ADDR" default:"redis:6379"`
}

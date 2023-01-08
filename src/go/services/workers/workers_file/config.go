package workers_file

type Config struct {
	UrlBase         string `ssm:"url_base" environment:"URL_BASE_UI"`
	EventArn        string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	FileEventsTopic string `json:"file-events"`
	RedisAddr       string `ssm:"redis_addr" json:"redis-addr" environment:"REDIS_ADDR" default:"redis:6379"`
	UserServiceAddr string `environment:"USER_SERVICE_ADDR" json:"core-user-addr"`
	FileServiceAddr string `environment:"FILE_SERVICE_ADDR" json:"core-file-addr"`
}

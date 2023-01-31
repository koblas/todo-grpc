package workers_file

type Config struct {
	// Used by lambda
	EventArn string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	// Used by kubernetes
	FileEventsTopic string `json:"file-events"`
	NatsAddr        string `environment:"NATS_ADDR"`
	RedisAddr       string `ssm:"redis_addr" json:"redis-addr" environment:"REDIS_ADDR" default:"redis:6379"`
	// Shared
	MinioEndpoint   string `environment:"MINIO_ENDPOINT" default:"s3.amazonaws.com"`
	UrlBase         string `ssm:"url_base" environment:"URL_BASE_UI"`
	UserServiceAddr string `environment:"USER_SERVICE_ADDR" json:"core-user-addr"`
	FileServiceAddr string `environment:"FILE_SERVICE_ADDR" json:"core-file-addr"`
	PublicBucket    string `environment:"PUBLIC_BUCKET"`
}

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
	UserServiceAddr string `environment:"USER_SERVICE_ADDR" json:"core-user-addr"`
	FileServiceAddr string `environment:"FILE_SERVICE_ADDR" json:"core-file-addr"`
	// Files are written here after conversion
	PublicBucket string `environment:"PUBLIC_BUCKET"`
}

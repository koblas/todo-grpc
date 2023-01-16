package file

type Config struct {
	// Used by Lambda
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	// Used by kubernetes
	NatsAddr       string `environment:"NATS_ADDR"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	FileEventTopic string `json:"file-events"`
	//  Shared
	S3Bucket string `json:"s3bucket"`
	S3Prefix string `json:"s3prefix" default:"/"`
}

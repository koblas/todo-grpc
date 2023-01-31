package file

type Config struct {
	// Used by Lambda
	EventArn      string `environment:"BUS_ENTITY_ARN" ssm:"bus_entity_arn"`
	S3Bucket      string `environment:"BUCKET"`
	S3Prefix      string `json:"s3prefix" default:"/"`
	S3DomainAlias string `environment:"BUCKET_ALIAS" default:""`
	// Used by kubernetes
	NatsAddr       string `environment:"NATS_ADDR"`
	RedisAddr      string `environment:"REDIS_ADDR" default:"redis:6379"`
	FileEventTopic string `json:"file-events"`
	//  Shared
}

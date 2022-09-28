package todo

type SsmConfig struct {
	EventArn string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
}

package websocket

type SsmConfig struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	ConnDb    string `environment:"CONN_DB"`
	RedisAddr string `ssm:"redis_addr" environment:"REDIS_ADDR"`
}

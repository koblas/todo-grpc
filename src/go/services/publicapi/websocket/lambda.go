package websocket

type SsmConfig struct {
	JwtSecret          string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	ConnDb             string `environment:"CONN_DB"`
	RedisAddr          string `ssm:"redis_addr" json:"redis-addr" environment:"REDIS_ADDR" default:"redis:6379"`
	WebsocketBroadcast string `json:"websocket-broadcast"`
}

package websocket

type Config struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	// Used by the lambda version
	ConnDb string `environment:"CONN_DB"`
	// Used by the docker-compose version
	RedisAddr                  string `json:"redis-addr" environment:"REDIS_ADDR" default:"redis:6379"`
	WebsocketConnectionMessage string `json:"websocket-connection-message"`
}

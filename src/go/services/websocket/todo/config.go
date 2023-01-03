package todo

type SsmConfig struct {
	ConnDb             string `environment:"CONN_DB"`
	WsEndpoint         string `environment:"WS_ENDPOINT"`
	RedisAddr          string `environment:"REDIS_ADDR"`
	WebsocketBroadcast string `json:"websocket-broadcast"`
	TodoEventsTopic    string `json:"todo-events"`
}

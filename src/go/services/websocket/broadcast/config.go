package broadcast

type Config struct {
	// Used by Lambda
	WsEndpoint string `environment:"WS_ENDPOINT"`
	// Used by Kubernetes
	RedisAddr                  string `environment:"REDIS_ADDR"`            // Message bus
	BroadcastEventTopic        string `json:"websocket-broadcast-events"`   // Event source
	WebsocketConnectionMessage string `json:"websocket-connection-message"` // Event target
	// Shared
	ConnDb string `environment:"CONN_DB"` // Connection store
}

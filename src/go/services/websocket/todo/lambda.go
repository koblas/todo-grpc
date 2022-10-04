package todo

type SsmConfig struct {
	UrlBase    string `ssm:"url_base" environment:"URL_BASE_UI"`
	ConnDb     string `environment:"CONN_DB"`
	WsEndpoint string `environment:"WS_ENDPOINT"`
	RedisAddr  string `environment:"REDIS_ADDR"`
}

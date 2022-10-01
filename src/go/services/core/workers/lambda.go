package workers

type SsmConfig struct {
	UrlBase string `ssm:"url_base" environment:"URL_BASE_UI"`
}

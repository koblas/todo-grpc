package user

type SsmConfig struct {
	EventArn        string `ssm:"bus_entity_arn" environment:"BUS_ENTITY_ARN"`
	JwtSecret       string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	UserServiceAddr string `environment:"USER_SERVICE_ADDR" json:"core-user-addr" default:":13001"`
}

type OauthConfig struct {
	GoogleClientId string `ssm:"google_client_id" environment:"GOOGLE_CLIENT_ID"`
	GoogleSecret   string `ssm:"google_secret" environment:"GOOGLE_SECRET"`
	GitHubClientId string `ssm:"github_client_id" environment:"GITHUB_CLIENT_ID"`
	GitHubSecret   string `ssm:"github_secret" environment:"GITHUB_SECRET"`
}

func (conf OauthConfig) GetSecret(provider string) (string, string, error) {
	if provider == "github" {
		return conf.GitHubClientId, conf.GitHubSecret, nil
	}
	if provider == "google" {
		return conf.GoogleClientId, conf.GoogleSecret, nil
	}
	return "", "", nil
}

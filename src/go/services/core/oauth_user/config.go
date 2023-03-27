package user

type OauthConfig struct {
	GoogleClientId string
	GoogleSecret   string
	GitHubClientId string
	GitHubSecret   string
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

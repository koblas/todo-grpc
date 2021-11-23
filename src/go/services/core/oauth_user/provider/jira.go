package provider

import (
	"context"
	"errors"
	"net/http"

	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/oauth2"
)

type jiraProvider struct {
	providerBase
}

var _ OAuthProvider = jiraProvider{}

func init() {
	build := func(config OAuthConfig, logger logger.Logger) OAuthProvider {
		return jiraProvider{providerBase{config, logger}}
	}
	providers["jira"] = build
}

func (svc jiraProvider) BuildRedirect(ctx context.Context, redirectURI string, state string) string {
	config := &oauth2.Config{
		ClientID:     svc.config.ClientId,
		ClientSecret: svc.config.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://auth.atlassian.com/authorize",
		},
		Scopes:      []string{"read:jira-user", "read:jira-work", "write:jira-work", "offline_access"},
		RedirectURL: redirectURI,
	}

	// @TODO - fix scope
	return config.AuthCodeURL(state,
		oauth2.SetAuthURLParam("audience", "api.atlassian.com"),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("prompt", "consent"),
	)
}

func (svc jiraProvider) GetAccessToken(ctx context.Context, code string, redirectURI string, state string) (TokenResult, error) {
	logger := svc.logger.With("method", "jiraProvider.GetAccessToken")

	return svc.httpTokenRequest(ctx, logger, "https://auth.atlassian.com/oauth/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"code":          code,
		"redirect_uri":  redirectURI,
		"grant_type":    "authorization_code",
	})
}

func (svc jiraProvider) RefreshToken(ctx context.Context, refreshToken string) (TokenResult, error) {
	logger := svc.logger.With("method", "jiraProvider.RefreshToken")

	logger.Info("Refreshing OAuthToken")

	return svc.httpTokenRequest(ctx, logger, "https://auth.atlassian.com/oauth/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	})
}

func (svc jiraProvider) GetInfo(ctx context.Context, tokenResult TokenResult) (*OAuthInfo, error) {
	logger := svc.logger.With("method", "jiraProvider.GetInfo")

	logger.Info("Doing a GetInfo")

	var info OAuthInfo

	req, err := http.NewRequest("GET", "https://api.atlassian.com/oauth/token/accessible-resources", nil)
	if err != nil {
		return nil, err
	}

	var resourcesResult []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatarUrl"`
	}

	source := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: tokenResult.AccessToken,
	})

	err = makeHTTPcall(ctx, source, logger, req, &resourcesResult)
	if err != nil {
		return nil, err
	}

	// @TODO handle more than one response
	if len(resourcesResult) == 0 {
		return nil, errors.New("No resources returned")
	}

	guid := resourcesResult[0].ID

	// New Try for the user
	req, err = http.NewRequest("GET", "https://api.atlassian.com/ex/jira/"+guid+"/rest/api/2/myself", nil)
	if err != nil {
		return nil, err
	}

	var userResult struct {
		AccountID   string `json:"accountId"`
		Name        string `json:"name"`
		Email       string `json:"emailAddress"`
		DisplayName string `json:"displayName"`
	}

	err = makeHTTPcall(ctx, source, logger, req, &userResult)
	if err != nil {
		return nil, err
	}

	info.Id = userResult.AccountID
	info.Name = userResult.DisplayName
	info.Email = userResult.Email

	logger.With(
		"oauth_id", info.Id,
		"oauth_name", info.Name,
		"oauth_email", info.Email,
	).Info("Got Fields")

	return &info, nil
}

func (svc jiraProvider) AssociateUser(ctx context.Context, tokenResult TokenResult) error {
	// logger := svc.logger.WithField("method", "githubProvider.GetAccesToken")
	return nil
}

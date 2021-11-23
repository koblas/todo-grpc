package provider

import (
	"context"
	"net/http"

	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleProvider struct {
	providerBase
}

var _ OAuthProvider = googleProvider{}

func init() {
	build := func(config OAuthConfig, logger logger.Logger) OAuthProvider {
		return googleProvider{providerBase{config, logger}}
	}

	providers["google"] = build
}

func (svc googleProvider) BuildRedirect(ctx context.Context, redirectURI string, state string) string {
	// logger := svc.logger.WithField("method", "googleProvider.BuildRedirect")
	config := &oauth2.Config{
		ClientID:     svc.config.ClientId,
		ClientSecret: svc.config.Secret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"email", "profile"},
		RedirectURL:  redirectURI,
	}

	return config.AuthCodeURL(state,
		oauth2.SetAuthURLParam("approval_prompt", "force"),
		oauth2.SetAuthURLParam("access_type", "offline"),
	)
}

func (svc googleProvider) GetAccessToken(ctx context.Context, code string, redirectURI string, state string) (TokenResult, error) {
	logger := svc.logger.With("method", "googleProvider.GetAccessToken")

	return svc.httpTokenRequest(ctx, logger, "https://www.googleapis.com/oauth2/v4/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"code":          code,
		"redirect_uri":  redirectURI,
		"grant_type":    "authorization_code",
	})
}

func (svc googleProvider) RefreshToken(ctx context.Context, refreshToken string) (TokenResult, error) {
	logger := svc.logger.With("method", "googleProvider.RefreshToken")

	return svc.httpTokenRequest(ctx, logger, "https://www.googleapis.com/oauth2/v4/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	})
}

func (svc googleProvider) GetInfo(ctx context.Context, tokenResult TokenResult) (*OAuthInfo, error) {
	logger := svc.logger.With("method", "googleProvider.GetInfo")

	var info OAuthInfo

	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}

	goResult := struct {
		ID         string `json:"id"`
		Email      string `json:"email"`
		Name       string `json:"name"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}{}

	if err := makeHTTPcall(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: tokenResult.AccessToken,
	}), logger, req, &goResult); err != nil {
		return nil, err
	}

	logger.With("data", goResult).Info("Got google data")

	info.Id = goResult.ID
	info.Email = goResult.Email
	info.Name = goResult.Name

	logger.With(
		"oauth_id", info.Id,
		"oauth_name", info.Name,
		"oauth_email", info.Email,
	).Info("Got Fields")

	return &info, nil
}

func (svc googleProvider) AssociateUser(ctx context.Context, tokenResult TokenResult) error {
	// logger := svc.logger.WithField("method", "googleProvider.GetAccesToken")
	return nil
}

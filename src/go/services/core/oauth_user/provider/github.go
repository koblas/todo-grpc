package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/oauth2"
)

type githubProvider struct {
	providerBase
}

var _ OAuthProvider = githubProvider{}

func init() {
	build := func(config OAuthConfig, logger logger.Logger) OAuthProvider {
		return githubProvider{providerBase{config, logger}}
	}
	providers["github"] = build
}

func (svc githubProvider) BuildRedirect(ctx context.Context, redirectURI string, state string) string {
	config := &oauth2.Config{
		ClientID:     svc.config.ClientId,
		ClientSecret: svc.config.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://github.com/login/oauth/authorize",
		},
		Scopes:      []string{"user"},
		RedirectURL: redirectURI,
	}

	// @TODO - fix scope
	return config.AuthCodeURL(state,
		oauth2.SetAuthURLParam("allow_signup", "false"),
	)
}

func (svc githubProvider) GetAccessToken(ctx context.Context, code string, redirectURI string, state string) (TokenResult, error) {
	logger := svc.logger.With("method", "githubProvider.GetAccessToken")

	return svc.httpTokenRequest(ctx, logger, "https://github.com/login/oauth/access_token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"code":          code,
		"scope":         "user",
		"redirect_uri":  redirectURI,
		"state":         state, // @todo -- should match
	})
}

func (svc githubProvider) RefreshToken(ctx context.Context, refreshToken string) (TokenResult, error) {
	svc.logger.With("method", "githubProvider.RefreshToken").Error("Refresh not implemented")

	return TokenResult{}, errors.New("Refresh not implemented")
}

func (svc githubProvider) GetInfo(ctx context.Context, tokenResult TokenResult) (*OAuthInfo, error) {
	logger := svc.logger.With("method", "githubProvider.GetInfo")

	var info OAuthInfo

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	var ghResult struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	source := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: tokenResult.AccessToken,
		TokenType:   "token",
	})

	err = makeHTTPcall(ctx, source, logger, req, &ghResult)
	if err != nil {
		return nil, err
	}

	// @TODO const om2 = await OAuthModel.getByProviderId(input.provider, goResult.id);

	info.Id = fmt.Sprintf("%d", ghResult.ID)
	info.Name = ghResult.Name

	logger.With(
		"oauth_id", info.Id,
		"oauth_name", info.Name,
		"oauth_email", info.Email,
	).Info("Got Fields")

	return &info, nil
}

func (svc githubProvider) AssociateUser(ctx context.Context, tokenResult TokenResult) error {
	// logger := svc.logger.WithField("method", "githubProvider.GetAccesToken")
	return nil
}

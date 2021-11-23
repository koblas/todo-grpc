package provider

import (
	"context"
	"errors"

	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/oauth2"
)

type stravaProvider struct {
	providerBase
}

var _ OAuthProvider = stravaProvider{}

func init() {
	build := func(config OAuthConfig, logger logger.Logger) OAuthProvider {
		return stravaProvider{providerBase{config, logger}}
	}
	providers["strava"] = build
}

func (svc stravaProvider) BuildRedirect(ctx context.Context, redirectURI string, state string) string {
	config := &oauth2.Config{
		ClientID:     svc.config.ClientId,
		ClientSecret: svc.config.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://www.strava.com/oauth/authorize",
		},
		Scopes:      []string{"read_all,activity:read_all"},
		RedirectURL: redirectURI,
	}

	return config.AuthCodeURL(state,
		oauth2.SetAuthURLParam("approval_prompt", "auto"),
		oauth2.SetAuthURLParam("access_type", "offline"),
	)
}

func (svc stravaProvider) GetAccessToken(ctx context.Context, code string, redirectURI string, state string) (TokenResult, error) {
	logger := svc.logger.With("method", "stravaProvider.GetAccesToken")

	// Strava returns "state" in ther response....

	token, err := svc.httpTokenRequest(ctx, logger, "https://www.strava.com/oauth/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"code":          code,
		"grant_type":    "authorization_code",
	})

	if err != nil {
		return TokenResult{}, err
	}

	if token.State != state {
		logger.With("stateExpected", state, "stateReceived", token.State).Error("State mismatch")
		return TokenResult{}, errors.New("strava oauth2 state didn't match")
	}

	return token, nil
}

func (svc stravaProvider) RefreshToken(ctx context.Context, refreshToken string) (TokenResult, error) {
	logger := svc.logger.With("method", "stravaProvider.RefreshToken")

	logger.Info("Refreshing OAuthToken")

	return svc.httpTokenRequest(ctx, logger, "https://www.strava.com/oauth/token", map[string]interface{}{
		"client_id":     svc.config.ClientId,
		"client_secret": svc.config.Secret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	})
}

func (svc stravaProvider) GetInfo(ctx context.Context, tokenResult TokenResult) (*OAuthInfo, error) {
	panic("Not implemented")
	// logger := svc.logger.With("method", "stravaProvider.GetInfo", "provider", "strava")

	// client := strava.NewClient(oauth2.StaticTokenSource(&oauth2.Token{AccessToken: tokenResult.AccessToken})).AtheleteService()

	// athelete, err := client.AtheleteGet().Do()
	// if err != nil {
	// 	return nil, err
	// }

	// info := OAuthInfo{
	// 	Id:   strconv.FormatInt(int64(athelete.Id), 10),
	// 	Name: fmt.Sprintf("%s %s", athelete.Firstname, athelete.Lastname),
	// }

	// logger.With(
	// 	"oauth_info", info,
	// 	"oauth_id", info.Id,
	// 	"oauth_name", info.Name,
	// 	"oauth_email", info.Email,
	// ).Info("Got OAUth Fields")

	// return &info, nil

	return nil, nil
}

// This make sure the Data hook is setup
func (svc stravaProvider) AssociateUser(ctx context.Context, tokenResult TokenResult) error {
	/*
		logger := svc.logger.WithField("method", "stravaProvider.AssociateUser")

		data := url.Values{
			"client_id":     {svc.config.Strava.ClientID},
			"client_secret": {svc.config.Strava.Secret},
			"callback_url":  {"https://hook.snaplabs.com/hook/strava"},
			"verify_token":  {"snaplabs"},
		}

		fmt.Println(data.Encode())

		req, err := http.NewRequest(
			"POST",
			"https://api.strava.com/api/v3/push_subscriptions",
			strings.NewReader(data.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		empty := struct {
			Id            int    `json:"id"`
			ResourceState int    `json:"resource_state"`
			ApplicationID int    `json:"application_id"`
			CallbackURL   string `json:"callback_url"`
			CreatedAt     string `json:"created_at"`
			UpdatedAt     string `json:"updated_at"`
		}{}

		if err := makeHTTPcall(ctx, logger, req, empty); err == nil {
			return nil
		} else {
			return nil
		}
	*/
	return nil
}

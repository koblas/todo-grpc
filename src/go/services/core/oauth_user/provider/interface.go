package provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/koblas/grpc-todo/pkg/logger"
)

type OAuthInfo struct {
	Id    string
	Name  string
	Email string
}

type providerBase struct {
	config OAuthConfig
	logger logger.Logger
}

type TokenResult struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token,optional"`
	Expires      *time.Time `json:"-,optional"`
	Scope        string     `json:"scope,optional"`

	// Strava returns "state" in the response
	State string `json:"state,optional"`
}

type OAuthProvider interface {
	BuildRedirect(ctx context.Context, redirectURI string, state string) string
	GetAccessToken(ctx context.Context, code string, redirectURI string, state string) (TokenResult, error)
	RefreshToken(ctx context.Context, refreshToken string) (TokenResult, error)
	GetInfo(ctx context.Context, tokenResult TokenResult) (*OAuthInfo, error)
	AssociateUser(ctx context.Context, tokenResult TokenResult) error
}

type OAuthConfig struct {
	ClientId string
	Secret   string
}

type providerBuilder func(config OAuthConfig, logger logger.Logger) OAuthProvider

var providers = map[string]providerBuilder{}

// GetOAuthProvider - the a service for handling the OAuth requests
func GetOAuthProvider(provider string, logger logger.Logger) (OAuthProvider, error) {
	up := strings.ToUpper(provider)
	clientId := os.Getenv(up + "_CLIENT_ID")
	secret := os.Getenv(up + "_SECRET")

	if clientId == "" || secret == "" {
		return nil, fmt.Errorf("Unable to get %s_SECRET or %s_CLIENT_ID from environment", up, up)
	}

	if factory, found := providers[provider]; found {
		return factory(OAuthConfig{
			ClientId: clientId,
			Secret:   secret,
		}, logger), nil
	}

	logger.With("provider", provider).Info("Unknown provider")

	return nil, fmt.Errorf("Unknown provider=%v", provider)
}

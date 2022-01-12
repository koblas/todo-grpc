package user

import (
	"time"

	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
)

type UserStatus string

type OauthUser struct {
	ID             string    `dynamodbav:"id"`
	UserId         string    `dynamodbav:"user_id"`
	Provider       string    `dynamodbav:"provider"`
	ProviderId     string    `dynamodbav:"provider_id"`
	ProviderMerged string    `dynamodbav:"provider_merged"`
	AccessToken    string    `dynamodbav:"access_token"`
	RefreshToken   string    `dynamodbav:"refresh_token"`
	ExpiresAt      time.Time `dynamodbav:"expires_at"`
}

type OAuthStore interface {
	ListByUserId(userId string) ([]OauthUser, error)
	FindByProviderId(provider string, providerId string) (*OauthUser, error)
	Associate(userId string, provider string, providerId string, token oauth_provider.TokenResult) error
}

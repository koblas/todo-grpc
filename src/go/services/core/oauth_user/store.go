package user

import (
	"time"

	"github.com/google/uuid"
	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
)

type UserStatus string

type OauthUser struct {
	ID           string
	UserId       string
	Provider     string
	ProviderId   string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type OAuthStore interface {
	ListByUserId(userId string) ([]OauthUser, error)
	FindByProviderId(provider string, providerId string) ([]OauthUser, error)
	Associate(userId string, provider string, providerId string, token oauth_provider.TokenResult) error
}

type oauthStore struct {
	database map[string][]OauthUser
}

func NewOauthMemoryStore() OAuthStore {
	return &oauthStore{
		database: map[string][]OauthUser{},
	}
}

func (store *oauthStore) ListByUserId(userId string) ([]OauthUser, error) {
	items, found := store.database[userId]
	if !found {
		return []OauthUser{}, nil
	}

	return items, nil
}

func (store *oauthStore) Associate(userId string, provider string, providerId string, token oauth_provider.TokenResult) error {
	if err := store.Remove(userId, provider); err != nil {
		return err
	}

	item := OauthUser{
		ID:           uuid.New().String(),
		UserId:       userId,
		Provider:     provider,
		ProviderId:   providerId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    *token.Expires,
	}

	store.database[userId] = append(store.database[userId], item)

	return nil
}

func (store *oauthStore) Remove(userId string, provider string) error {
	items, found := store.database[userId]
	if !found {
		return nil
	}

	for idx, item := range items {
		if item.UserId == userId && item.Provider == provider {
			store.database[userId] = append(items[:idx], items[idx+1:]...)
			return nil
		}
	}

	return nil
}

func (store *oauthStore) FindByProviderId(provider string, providerId string) ([]OauthUser, error) {
	for _, items := range store.database {
		for _, item := range items {
			if item.Provider == provider && item.ProviderId == providerId {
				return []OauthUser{item}, nil
			}
		}
	}

	return []OauthUser{}, nil
}

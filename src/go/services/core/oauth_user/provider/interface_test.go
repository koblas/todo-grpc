package provider_test

import (
	"testing"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	"github.com/stretchr/testify/require"
)

type hasSecret struct{}

func (hasSecret) GetSecret(provider string) (string, string, error) {
	return "ali", "opensaysme", nil
}

type hasNoSecret struct{}

func (hasNoSecret) GetSecret(provider string) (string, string, error) {
	return "", "", nil
}

func TestGetOAuthProvider(t *testing.T) {
	p, err := provider.GetOAuthProvider("github", hasSecret{}, logger.NewNopLogger())

	require.Nil(t, err)
	require.NotNil(t, p)
}

func TestGetOAuthProviderBad(t *testing.T) {
	p, err := provider.GetOAuthProvider("unkown", hasNoSecret{}, logger.NewNopLogger())

	require.Nil(t, p)
	require.NotNil(t, err)
}

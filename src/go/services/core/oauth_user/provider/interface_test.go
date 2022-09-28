package provider_test

import (
	"testing"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	"github.com/stretchr/testify/require"
)

type basicSecret struct{}

func (basicSecret) GetSecret(provider string) (string, string, error) {
	return "", "", nil
}

func TestGetOAuthProvider(t *testing.T) {
	p, err := provider.GetOAuthProvider("github", basicSecret{}, logger.NewNopLogger())

	require.Nil(t, err)
	require.NotNil(t, p)
}

func TestGetOAuthProviderBad(t *testing.T) {
	p, err := provider.GetOAuthProvider("unkown", basicSecret{}, logger.NewNopLogger())

	require.Nil(t, p)
	require.NotNil(t, err)
}

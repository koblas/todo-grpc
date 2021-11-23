package provider_test

import (
	"testing"

	"github.com/koblas/projectx/server-go/pkg/config"
	"github.com/koblas/projectx/server-go/pkg/log"
	"github.com/koblas/projectx/server-go/pkg/services/oauth/provider"
	"github.com/stretchr/testify/require"
)

func TestGetOAuthProvider(t *testing.T) {
	p, err := provider.GetOAuthProvider("github", config.OAuth{}, log.NewNopLogger())

	require.Nil(t, err)
	require.NotNil(t, p)
}

func TestGetOAuthProviderBad(t *testing.T) {
	p, err := provider.GetOAuthProvider("unkown", config.OAuth{}, log.NewNopLogger())

	require.Nil(t, p)
	require.NotNil(t, err)
}

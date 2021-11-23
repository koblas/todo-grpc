package key_manager_test

import (
	"testing"

	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	svc := key_manager.NewSecureEnvironment()

	value := []byte{1, 2, 3}

	result, err := svc.Encode(value)

	require.Nil(t, err, "Error returned", err)
	require.NotZero(t, len(result.KmsUri), "No KMS key")
	require.NotZero(t, len(result.DataKey), "No data key")
	require.NotZero(t, len(result.Data), "No data ")
}

func TestDecrypt(t *testing.T) {
	svc := key_manager.NewSecureEnvironment()

	value := []byte{1, 2, 3}

	result, err := svc.Encode(value)

	require.Nil(t, err, "Error returned")

	data, err := svc.Decode(result)

	require.Nil(t, err, "Error returned")
	require.Equal(t, value, data, "Round trip good")
}

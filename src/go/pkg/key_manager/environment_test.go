package key_manager_test

import (
	"fmt"
	"testing"

	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/stretchr/testify/assert"
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

func TestDecrypt_1(t *testing.T) {
	svc := key_manager.NewSecureEnvironment()

	value := []byte{1, 2, 3}

	result, err := svc.Encode(value)

	fmt.Printf("uri=%s\n", result.KmsUri)
	fmt.Printf("dataKey=%v\n", result.DataKey)
	fmt.Printf("data=%v\n", string(result.Data))

	require.Nil(t, err, "Error returned")

	data, err := svc.Decode(result)

	require.Nil(t, err, "Error returned")
	require.Equal(t, value, data, "Round trip good")
}

func TestDecrypt_Empty(t *testing.T) {
	svc := key_manager.NewSecureEnvironment()

	value := []byte{}

	result, err := svc.Encode(value)

	fmt.Printf("uri=%s\n", result.KmsUri)
	fmt.Printf("dataKey=%v\n", result.DataKey)
	fmt.Printf("data=%v\n", string(result.Data))

	require.Nil(t, err, "Error returned")

	data, err := svc.Decode(result)

	require.Nil(t, err, "Error returned")
	require.Equal(t, value, data, "Round trip good")
}

func TestDecrypt_Mismatch(t *testing.T) {
	s1 := key_manager.NewSecureEnvironment()
	s2 := key_manager.NewSecureEnvironment()

	value := []byte{1, 2, 3}

	result, err := s1.Encode(value)
	assert.NoError(t, err)

	_, err = s2.Decode(result)

	assert.ErrorIs(t, err, key_manager.ErrorKeyLookupFailed)
}

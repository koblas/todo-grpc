package password_test

import (
	"testing"

	"github.com/koblas/grpc-todo/pkg/password"
	"github.com/stretchr/testify/assert"
)

func TestArgonSmoke(t *testing.T) {
	argon2ID := password.NewArgon2ID()

	plain := "inanzzz"

	hash1, err := argon2ID.Hash(plain)
	assert.NoError(t, err, "hash failed")
	assert.NotEmpty(t, hash1, "invalid hash")

	hash2, err := argon2ID.Hash(plain)
	assert.NoError(t, err, "hash failed")
	assert.NotEmpty(t, hash1, "invalid hash")
	assert.NotEqual(t, hash1, hash2, "two passwords should be different")

	ok, err := argon2ID.Verify(plain, hash1)
	assert.NoError(t, err)
	assert.True(t, ok, "passwords didn't verify")

	ok, err = argon2ID.Verify(plain+"bad", hash1)
	assert.NoError(t, err)
	assert.False(t, ok, "different passwords should be different")
}

package password_test

import (
	"testing"

	"github.com/koblas/grpc-todo/pkg/password"
	"github.com/stretchr/testify/assert"
)

func TestBcryptSmoke(t *testing.T) {
	svc := password.NewBcrypt()

	plain := "inanzzz"

	hash1, err := svc.Hash(plain)
	assert.NoError(t, err, "hash failed")
	assert.NotEmpty(t, hash1, "invalid hash")

	hash2, err := svc.Hash(plain)
	assert.NoError(t, err, "hash failed")
	assert.NotEmpty(t, hash1, "invalid hash")
	assert.NotEqual(t, hash1, hash2, "two passwords should be different")

	ok, err := svc.Verify(plain, hash1)
	assert.NoError(t, err)
	assert.True(t, ok, "passwords didn't verify")

	ok, err = svc.Verify(plain+"bad", hash1)
	assert.NoError(t, err)
	assert.False(t, ok, "different passwords should be different")
}

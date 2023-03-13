package confmgr

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type baseTestCase struct {
	Foo   string `environment:"TEST_FOO" default:"xyzzy"`
	Bar   int    `environment:"TEST_BAR" default:"3"`
	Inner struct {
		Baz bool `environment:"TEST_BAZ" default:"true"`
	}
}

func envSet(key, value string) func() {
	origValue, origSet := os.LookupEnv(key)

	os.Setenv(key, value)

	return func() {
		if origSet {
			os.Setenv(key, origValue)
		} else {
			os.Unsetenv(key)
		}
	}
}

func TestBasic(t *testing.T) {
	item := baseTestCase{}
	spec, err := walkInput(&item)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(spec))
}

func TestBadInput(t *testing.T) {
	item := struct{}{}
	value := 7

	err := NewLoader().Parse(context.TODO(), item)
	assert.Error(t, err)

	err = NewLoader().Parse(context.TODO(), value)
	assert.Error(t, err)

	err = NewLoader().Parse(context.TODO(), &value)
	assert.Error(t, err)
}

func TestDefault(t *testing.T) {
	item := baseTestCase{}
	err := NewLoader().Parse(context.TODO(), &item)

	assert.NoError(t, err)

	assert.Equal(t, "xyzzy", item.Foo)
	assert.Equal(t, 3, item.Bar)
}

func TestEnvironment(t *testing.T) {
	defer envSet("TEST_FOO", "fishfood")()
	defer envSet("TEST_BAR", "99")()

	item := baseTestCase{}
	err := NewLoader().Parse(context.TODO(), &item)

	assert.NoError(t, err)

	assert.Equal(t, "fishfood", item.Foo)
	assert.Equal(t, 99, item.Bar)
}

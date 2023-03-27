package confmgr_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvCommon struct {
	Foo string
	Bar int
}
type EnvSmtp struct {
	Foo string
	Bar int
}

type baseEnvTest struct {
	Foo string
	Bar int

	Common EnvCommon
	Smtp   EnvSmtp
}

func (EnvTests) envSet(key, value string) func() {
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

type EnvTests struct {
	suite.Suite
}

func TestEnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvTests))
}

func (suite *EnvTests) TestBasic() {
	prefix := faker.UUIDHyphenated()

	defer suite.envSet(prefix+"_FOO", "fishfood")()
	defer suite.envSet(prefix+"_BAR", "99")()
	defer suite.envSet(prefix+"_COMMON_BAR", "11")()
	defer suite.envSet(prefix+"_SMTP_BAR", "33")()

	item := baseEnvTest{
		Common: EnvCommon{
			Foo: "cat",
		},
		Smtp: EnvSmtp{
			Foo: "dog",
		},
	}
	err := confmgr.NewLoader(confmgr.NewLoaderEnvironment(prefix, "_")).Parse(context.TODO(), &item)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "fishfood", item.Foo)
	assert.Equal(suite.T(), 99, item.Bar)
	assert.Equal(suite.T(), 11, item.Common.Bar)
	assert.Equal(suite.T(), "cat", item.Common.Foo)
	assert.Equal(suite.T(), 33, item.Smtp.Bar)
	assert.Equal(suite.T(), "dog", item.Smtp.Foo)
}

func (suite *EnvTests) TestPointer() {
	prefix := faker.UUIDHyphenated()

	defer suite.envSet(prefix+"_FOO", "fishfood")()
	defer suite.envSet(prefix+"_BAR", "8")()

	item := struct {
		Bar *int8
		Foo *string
	}{}
	err := confmgr.NewLoader(confmgr.NewLoaderEnvironment(prefix, "_")).Parse(context.TODO(), &item)

	assert.NoError(suite.T(), err)

	assert.NotNil(suite.T(), item.Foo)
	assert.Equal(suite.T(), "fishfood", *item.Foo)
	assert.NotNil(suite.T(), item.Bar)
	assert.EqualValues(suite.T(), 8, *item.Bar)
}

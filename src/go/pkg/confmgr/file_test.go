package confmgr_test

import (
	"strings"
	"testing"

	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/stretchr/testify/assert"
)

type baseTestCase struct {
	Foo   string `json:"foo"`
	Bar   int
	Inner struct {
		Baz bool   `json:"TEST_BAZ"`
		Foo string `json:"TEST_INNER_BAZ"`
	}
}

func TestFileSmoke(t *testing.T) {
	base := baseTestCase{}
	err := confmgr.Parse(&base, confmgr.NewJsonReader(strings.NewReader("")))

	assert.NoError(t, err)

	err = confmgr.Parse(&base, confmgr.NewJsonReader(strings.NewReader("{}")))
	assert.NoError(t, err)

	err = confmgr.Parse(&base, confmgr.NewJsonReader(strings.NewReader(`{"base":77}`)))
	assert.NoError(t, err)
}

func TestFileBasic(t *testing.T) {
	base := baseTestCase{}

	input := `
	{
		"foo": "car",
		"Bar": 77
	}
	`

	err := confmgr.Parse(&base, confmgr.NewJsonReader(strings.NewReader(input)))

	assert.NoError(t, err)
	assert.Equal(t, "car", base.Foo)
	assert.Equal(t, "", base.Inner.Foo)
}

package confmgr

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type baseTestCase struct {
	Foo   string
	Bar   int
	Inner struct {
		Baz bool
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
	item := baseTestCase{
		Bar: 3,
		Foo: "xyzzy",
	}
	err := NewLoader().Parse(context.TODO(), &item)

	assert.NoError(t, err)

	assert.Equal(t, "xyzzy", item.Foo)
	assert.Equal(t, 3, item.Bar)
}

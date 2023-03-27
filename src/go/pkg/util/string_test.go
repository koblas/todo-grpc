package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// here is the test
func TestToSnake(t *testing.T) {
	pairs := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "MyLIFEIsAwesomE",
			Expect: "my_life_is_awesom_e",
		},
		{
			Input:  "SomeURL",
			Expect: "some_url",
		},
	}

	for _, test := range pairs {
		got := ToSnake(test.Input)

		assert.Equal(t, test.Expect, got, "failed")
	}
}

package send_email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayout(t *testing.T) {
	layout := Layout()

	assert.NotNil(t, layout)
}

package redisutil

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuePayload(t *testing.T) {
	payload := queuePayload{
		Path: "One",
		Body: "Two",
		Headers: payloadHeaders{
			"foo": []string{"bar"},
			"big": []string{"bar", "baz"},
		},
	}

	bytes, err := json.Marshal(&payload)
	assert.NoError(t, err)

	result := queuePayload{}
	err = json.Unmarshal(bytes, &result)

	assert.NoError(t, err)

	assert.Equal(t, &payload, &result)
}

func TestQueueMap(t *testing.T) {
	payloadStruct := queuePayload{
		Path: "One",
		Body: "Two",
		Headers: payloadHeaders{
			"foo": []string{"bar"},
			"big": []string{"bar", "baz"},
		},
	}

	payload := map[string]interface{}{
		"Path": "One",
		"Body": "Two",
		"Headers": map[string][]string{
			"foo": {"bar"},
			"big": {"bar", "baz"},
		},
	}

	bytes, err := json.Marshal(&payload)
	assert.NoError(t, err)

	result := queuePayload{
		// Headers: payloadHeaders{},
	}
	err = json.Unmarshal(bytes, &result)

	assert.NoError(t, err)

	assert.Equal(t, &payloadStruct, &result)
}

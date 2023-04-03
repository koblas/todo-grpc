package util

import (
	"encoding/json"

	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"google.golang.org/protobuf/proto"
)

type SocketMessage struct {
	ObjectId string           `json:"object_id"`
	Action   string           `json:"action"`
	Topic    string           `json:"topic"`
	Body     *json.RawMessage `json:"body"`
}

func MarshalData(topic, objectId, action string, msg proto.Message) ([]byte, error) {
	codec := bufcutil.NewJsonCodec()
	bodyString, err := codec.Marshal(msg)
	if err != nil {
		return nil, err
	}
	raw := json.RawMessage(bodyString)
	data, err := json.Marshal(SocketMessage{
		Topic:    topic,
		ObjectId: objectId,
		Action:   action,
		Body:     &raw,
	})

	return data, err
}

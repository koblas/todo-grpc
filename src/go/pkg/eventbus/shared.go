package eventbus

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func EnqueuePb(ctx context.Context, producer Producer, msg ChangeMessage) error {
	body := map[string]string{}

	if msg.Current != nil {
		value, err := proto.Marshal(msg.Current)
		if err != nil {
			return errors.Wrap(err, "failed to marshal Current")
		}
		body["current"] = base64.RawStdEncoding.EncodeToString(value)
	}
	if msg.Original != nil {
		value, err := proto.Marshal(msg.Original)
		if err != nil {
			return errors.Wrap(err, "failed to marshal Original")
		}
		body["original"] = base64.RawStdEncoding.EncodeToString(value)
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return errors.Wrap(err, "unable to serialize to json")
	}

	attributes := map[string]string{
		"stream": msg.Stream,
		"action": msg.Action,
	}
	for k, v := range msg.Attributes {
		attributes[k] = v
	}

	return producer.Enqueue(ctx, &Message{
		Attributes: attributes,
		BodyString: string(bytes),
	})
}

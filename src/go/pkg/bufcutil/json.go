package bufcutil

import (
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func NewJsonCodec() *protoJSONCodec {
	return &protoJSONCodec{"json"}
}

type protoJSONCodec struct {
	name string
}

var _ connect.Codec = (*protoJSONCodec)(nil)

func (c *protoJSONCodec) Name() string { return c.name }

func (c *protoJSONCodec) Marshal(message any) ([]byte, error) {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return nil, errNotProto(message)
	}
	options := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	return options.Marshal(protoMessage)
}

func (c *protoJSONCodec) Unmarshal(binary []byte, message any) error {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return errNotProto(message)
	}
	if len(binary) == 0 {
		return errors.New("zero-length payload is not a valid JSON object")
	}
	options := protojson.UnmarshalOptions{DiscardUnknown: true}
	return options.Unmarshal(binary, protoMessage)
}

func errNotProto(message any) error {
	return fmt.Errorf("%T doesn't implement proto.Message", message)
}

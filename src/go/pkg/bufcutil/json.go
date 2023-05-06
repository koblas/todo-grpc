package bufcutil

import (
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func WithJSON() connect.Option {
	return connect.WithOptions(
		connect.WithCodec(newJsonCodec("json")),
		connect.WithCodec(newJsonCodec("json; charset=utf-8")),
	)
}

func NewJsonCodec() *JsonCodec {
	return newJsonCodec("json")
}

func newJsonCodec(name string) *JsonCodec {
	marshal := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	unmarshal := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	return &JsonCodec{name: name, marshal: marshal, unmarshal: unmarshal}
}

type JsonCodec struct {
	name string

	marshal   protojson.MarshalOptions
	unmarshal protojson.UnmarshalOptions
}

var _ connect.Codec = (*JsonCodec)(nil)

func (c *JsonCodec) Name() string   { return c.name }
func (c *JsonCodec) IsBinary() bool { return false }

func (c *JsonCodec) Marshal(message any) ([]byte, error) {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return nil, errNotProto(message)
	}
	return c.marshal.Marshal(protoMessage)
}

func (c *JsonCodec) Unmarshal(binary []byte, message any) error {
	protoMessage, ok := message.(proto.Message)
	if !ok {
		return errNotProto(message)
	}
	if len(binary) == 0 {
		return errors.New("zero-length payload is not a valid JSON object")
	}
	return c.unmarshal.Unmarshal(binary, protoMessage)
}

func errNotProto(message any) error {
	if _, ok := message.(protoiface.MessageV1); ok {
		return fmt.Errorf("%T uses github.com/golang/protobuf, but connect-go only supports google.golang.org/protobuf: see https://go.dev/blog/protobuf-apiv2", message)
	}

	return fmt.Errorf("%T doesn't implement proto.Message", message)
}

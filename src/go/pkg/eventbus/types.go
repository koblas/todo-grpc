package eventbus

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Message struct {
	ID         string
	Stream     string
	Attributes map[string]string
	BodyBytes  []byte
	BodyString string
}

type ChangeMessage struct {
	ID         string
	Stream     string
	Action     string
	Attributes map[string]string
	Current    proto.Message
	Original   proto.Message
}

type Producer interface {
	Enqueue(ctx context.Context, msg *Message) error
}

type Consumer struct {
	Messages chan *Message
	Errors   chan error
	Closing  chan struct{}
}

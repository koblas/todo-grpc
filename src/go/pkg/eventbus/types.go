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

// SimpleMessage is just a basic building block message, but gets the job done
type SimpleMessage struct {
	ID         string
	Attributes map[string]string
	Body       string
}

type Producer interface {
	Enqueue(ctx context.Context, msg *Message) error
}

type SimpleProducer interface {
	Write(ctx context.Context, msg *SimpleMessage) error
}

type SimpleConsumer interface {
	Next() (SimpleMessage, error)
}

type Consumer struct {
	Messages chan *Message
	Errors   chan error
	Closing  chan struct{}
}

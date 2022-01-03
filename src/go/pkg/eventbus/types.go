package eventbus

import "context"

type Message struct {
	ID         string
	Stream     string
	Attributes map[string]string
	BodyBytes  []byte
	BodyString string
}

type Producer interface {
	Enqueue(ctx context.Context, msg *Message) error
}

type Consumer struct {
	Messages chan *Message
	Errors   chan error
	Closing  chan struct{}
}

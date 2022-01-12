package stub

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/eventbus"
)

type stubBus struct{}

func NewEventbusStub() eventbus.Producer {
	return stubBus{}
}

func (stubBus) Enqueue(ctx context.Context, msg *eventbus.Message) error {
	return nil
}

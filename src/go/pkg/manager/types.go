package manager

import (
	"context"

	"github.com/koblas/grpc-todo/gen/corepb"
)

type HandlerStart interface {
	Start(ctx context.Context) error
}

type MsgHandler interface {
	GroupName() string
	Handler() corepb.TwirpServer
}

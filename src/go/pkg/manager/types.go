package manager

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

type HandlerStart interface {
	Start(ctx context.Context) error
}

type MsgHandler interface {
	GroupName() string
	Handler() corepbv1.TwirpServer
}

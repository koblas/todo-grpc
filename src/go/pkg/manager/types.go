package manager

import (
	"context"
	"net/http"
)

type HandlerStart interface {
	Start(ctx context.Context) error
}

type MsgHandler interface {
	GroupName() string
	Handler() http.Handler
}

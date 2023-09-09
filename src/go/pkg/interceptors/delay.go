package interceptors

import (
	"context"
	"strconv"
	"time"

	"github.com/bufbuild/connect-go"
)

type delayCtxKeyType string

const DelayCtxKey delayCtxKeyType = ""
const DelayHeader = "x-sleep-delay"
const MAX_DELAY = 60 // seconds

func GetDelayFromHeaders(ctx context.Context, req connect.AnyRequest) int {
	value := req.Header().Get(DelayHeader)
	if value == "" {
		return 0
	}

	delay, err := strconv.Atoi(value)
	// If it's non-sensical (traveling backwards in time...)
	if err != nil || delay < 0 {
		return 0
	}
	// You can't delay more than 60 seconds
	if delay > MAX_DELAY {
		return MAX_DELAY
	}

	return delay
}

func NewDelayInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			delay := GetDelayFromHeaders(ctx, req)

			if delay != 0 {
				time.Sleep(time.Second * time.Duration(delay))
			}

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

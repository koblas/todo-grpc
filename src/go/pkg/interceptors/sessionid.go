package interceptors

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type requestIdCtxKeyType string

const RequestIdCtxKey requestIdCtxKeyType = "requestId"
const RequestIdHeader = "x-request-id"

func GetReqIdFromHeaders(ctx context.Context, req connect.AnyRequest) string {
	value := req.Header().Get(RequestIdHeader)
	if value != "" && len(value) < 40 {
		return value
	}

	return xid.New().String()
}

func NewReqidInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			log := logger.FromContext(ctx)
			reqId := GetReqIdFromHeaders(ctx, req)

			ctx = context.WithValue(ctx, RequestIdCtxKey, reqId)
			ctx = logger.ToContext(ctx, log.With(zap.String("reqid", reqId)))

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

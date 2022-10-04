package manager

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"
)

type Manager struct {
	log logger.Logger
	ctx context.Context
}

func NewManager() *Manager {
	log := logger.NewZap(logger.LevelInfo)

	return &Manager{
		log: log,
		ctx: logger.ToContext(context.Background(), log),
	}
}

func (mgr *Manager) Logger() logger.Logger {
	return mgr.log
}

func (mgr *Manager) Context() context.Context {
	return mgr.ctx
}

func (mgr *Manager) Start(api http.Handler) {
	mgr.StartWithContext(mgr.ctx, api)
}

func (mgr *Manager) StartWithContext(ctx context.Context, api http.Handler) {
	mgr.log.Info("starting service")

	if isLambda() {
		lambda.StartWithContext(ctx, awsutil.HandleApiLambda(ctx, api))
	} else {
		server := &http.Server{
			Addr:        ":" + util.Getenv("PORT", "14586"),
			Handler:     withHeaders(api),
			BaseContext: func(net.Listener) context.Context { return ctx },
		}

		if err := server.ListenAndServe(); err != nil {
			mgr.log.With(zap.Error(err)).Fatal("failed to serve")
		}
	}
}

func (mgr *Manager) StartConsumer(handler awsutil.TwirpHttpSqsHandler) {
	mgr.StartConsumerWithContext(mgr.ctx, handler)
}

func (mgr *Manager) StartConsumerWithContext(ctx context.Context, handler awsutil.TwirpHttpSqsHandler) {
	if isLambda() {
		lambda.StartWithContext(ctx, handler)
	} else {
		// A little funky in that we're assuming this never returns until all messages are consumed
		//
		// It would probably be better to re-abstract this to a channel based system
		handler(ctx, events.SQSEvent{})
	}
}

func withHeaders(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), awsutil.HeaderCtxKey, r.Header)

		base.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isLambda() bool {
	_, found := os.LookupEnv("LAMBDA_TASK_ROOT")

	return found
}

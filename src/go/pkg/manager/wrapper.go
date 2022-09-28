package manager

import (
	"context"
	"net"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/redis"
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

	if util.Getenv("LAMBDA_TASK_ROOT", "") != "" {
		lambda.StartWithContext(ctx, awsutil.HandleApiLambda(ctx, api))
	} else {
		server := &http.Server{
			Addr:        ":" + util.Getenv("PORT", "14586"),
			Handler:     api,
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
	if util.Getenv("LAMBDA_TASK_ROOT", "") != "" {
		lambda.StartWithContext(ctx, handler)
	} else {
		svr := redis.NewRedisConsumer(util.Getenv("REDIS_ADDR", "redis:6379"))

		worker, err := svr.BuildWorker("", "")
		if err != nil {
			mgr.log.With(zap.Error(err)).Fatal("unable to start redis consumer")
		}

		for {
			select {
			case msg := <-worker.Messages:
				event := events.SQSEvent{
					Records: []events.SQSMessage{
						{
							MessageId:  msg.ID,
							Attributes: msg.Attributes,
							Body:       msg.BodyString,
						},
					},
				}

				handler(ctx, event)
			case err := <-worker.Errors:
				mgr.log.With(zap.Error(err)).Error("protocol error")
			}
		}
	}
}

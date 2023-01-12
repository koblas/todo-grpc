package manager

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	// healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type Manager struct {
	log            logger.Logger
	ctx            context.Context
	port           string
	healthPrefix   string
	grpcHealthPort string
}

type grpcServer struct {
	manager *Manager
}

type Option func(*Manager)

func WithPort(port string) Option {
	return func(mgr *Manager) {
		mgr.port = port
	}
}

func WithHealth(prefix string) Option {
	return func(mgr *Manager) {
		mgr.healthPrefix = prefix
	}
}

func WithGrpcHealth(port string) Option {
	return func(mgr *Manager) {
		mgr.grpcHealthPort = port
	}
}

func NewManager(opts ...Option) *Manager {
	log := logger.NewZap(logger.LevelInfo)

	mgr := &Manager{
		log:  log,
		ctx:  logger.ToContext(context.Background(), log),
		port: util.Getenv("PORT", "14586"),
	}

	for _, opt := range opts {
		opt(mgr)
	}

	return mgr
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
	if isLambda() {
		mgr.log.Info("starting service - lambda")
		lambda.StartWithContext(ctx, awsutil.HandleApiLambda(ctx, api))
	} else {
		mgr.startGrpcHealthCheck()
		mgr.log.Info("starting service", zap.String("port", mgr.port))
		server := &http.Server{
			Addr:        ":" + mgr.port,
			Handler:     mgr.withHeaders(api),
			BaseContext: func(net.Listener) context.Context { return ctx },
		}

		if err := server.ListenAndServe(); err != nil {
			mgr.log.With(zap.Error(err)).Fatal("failed to serve")
		}
	}
}

func (mgr *Manager) StartConsumer(handler awsutil.TwirpHttpSqsHandler) {
	go func() {
		server := &http.Server{
			Addr:        ":" + mgr.port,
			Handler:     http.HandlerFunc(mgr.healthCheck),
			BaseContext: func(net.Listener) context.Context { return mgr.ctx },
		}
		if err := server.ListenAndServe(); err != nil {
			mgr.log.With(zap.Error(err)).Fatal("failed to serve")
		}
	}()

	mgr.StartConsumerWithContext(mgr.ctx, handler)
}

func (mgr *Manager) StartConsumerWithContext(ctx context.Context, handler awsutil.TwirpHttpSqsHandler) {
	if isLambda() {
		lambda.StartWithContext(ctx, handler)
	} else {
		// A little funky in that we're assuming this never returns until all messages are consumed
		//
		// It would probably be better to re-abstract this to a channel based system
		mgr.startGrpcHealthCheck()

		handler(ctx, events.SQSEvent{})
	}
}

func (mgr *Manager) startGrpcHealthCheck() error {
	if mgr.grpcHealthPort == "" {
		return nil
	}
	log := mgr.log.With(zap.String("port", mgr.grpcHealthPort))

	// server := grpcServer{mgr}

	lis, err := net.Listen("tcp", ":"+mgr.grpcHealthPort)
	if err != nil {
		mgr.log.With(zap.Error(err)).Error("failed to listen on port")
		return err
	}
	s := grpc.NewServer()
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)
	log.Info("gRPC healthcheck server is running")

	go func() {
		if err := s.Serve(lis); err != nil {
			log.With(zap.Error(err)).Error("failed to serve")
		}
	}()

	return nil
}

func (mgr *Manager) withHeaders(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mgr.healthPrefix != "" && strings.HasPrefix(r.URL.Path, mgr.healthPrefix) && r.Method == "GET" {
			mgr.healthCheck(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), awsutil.HeaderCtxKey, r.Header)

		base.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mgr *Manager) healthCheck(w http.ResponseWriter, r *http.Request) {
	if mgr.healthPrefix == "" || !strings.HasPrefix(r.URL.Path, mgr.healthPrefix) || r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	path := strings.Trim(strings.TrimPrefix(r.URL.Path, mgr.healthPrefix), "/")
	if path == "ready" || path == "live" {
		fmt.Println("Got ", path, " check")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func isLambda() bool {
	_, found := os.LookupEnv("LAMBDA_TASK_ROOT")

	return found
}

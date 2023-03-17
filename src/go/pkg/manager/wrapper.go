package manager

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"

	// healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type headerCtxKeyType string

const HttpHeaderCtxKey headerCtxKeyType = "headers"

type Manager struct {
	log            logger.Logger
	ctx            context.Context
	tracer         *sdktrace.TracerProvider
	port           string
	healthPrefix   string
	grpcHealthPort string

	//
	grpcHealthServer *http.Server
	httpServer       *http.Server
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

	exp, err := newExporter()
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to construct tracer")
	}

	sp := sdktrace.NewBatchSpanProcessor(exp)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sp),
		// sdktrace.WithBatcher(exp),
		sdktrace.WithResource(newResource()),
	)
	otel.SetTracerProvider(tp)

	mgr := &Manager{
		log:    log,
		ctx:    logger.ToContext(context.Background(), log),
		port:   util.Getenv("PORT", "14586"),
		tracer: tp,
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

func (mgr *Manager) Shutdown() {
	fmt.Println("CALLING SHUTDOWN")
	mgr.tracer.Shutdown(mgr.ctx)
}

// Tracer setup
func newExporter() (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
	// stdouttrace.WithWriter(w),
	// Use human-readable output.
	// stdouttrace.WithPrettyPrint(),
	// Do not print timestamps for the demo.
	// stdouttrace.WithoutTimestamps(),
	)
}

func newResource() *resource.Resource {
	// r, _ := resource.Merge(
	// 	resource.Default(),
	// 	resource.NewWithAttributes(
	// 		semconv.SchemaURL,
	// 		semconv.ServiceName("fib"),
	// 		semconv.ServiceVersion("v0.1.0"),
	// 		attribute.String("environment", "demo"),
	// 	),
	// )
	// return r
	return resource.Default()
}

// func (mgr *Manager) Start(api http.Handler) {
// 	mgr.StartWithContext(mgr.ctx, api)
// }

// func (mgr *Manager) StartWithContext(ctx context.Context, api http.Handler) {
// 	if isLambda() {
// 		mgr.log.Info("starting service - lambda")
// 		lambda.StartWithContext(ctx, awsutil.HandleApiLambda(ctx, api))
// 	} else {
// 		mgr.startGrpcHealthCheck()
// 		mgr.log.Info("starting service", zap.String("port", mgr.port))
// 		server := &http.Server{
// 			Addr:        ":" + mgr.port,
// 			Handler:     mgr.withHeaders(api),
// 			BaseContext: func(net.Listener) context.Context { return ctx },
// 		}

// 		if err := server.ListenAndServe(); err != nil {
// 			mgr.log.With(zap.Error(err)).Fatal("failed to serve")
// 		}
// 	}
// }

func (mgr *Manager) Start(handler HandlerStart) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	if isLambda() {
		// If we receive a SIGTERM call the shutdown
		go handler.Start(mgr.ctx)

		<-interrupt
	} else {
		// A little funky in that we're assuming this never returns until all messages are consumed
		//
		// It would probably be better to re-abstract this to a channel based system
		ctx, cancel := context.WithCancel(mgr.ctx)
		defer cancel()

		group, ctx := errgroup.WithContext(ctx)
		mgr.ctx = ctx

		group.Go(func() error { return mgr.startGrpcHealthCheck(ctx) })
		group.Go(func() error { return handler.Start(ctx) })

		select {
		case <-interrupt:
			mgr.log.Info("Got interrupt")
			break
		case <-ctx.Done():
			mgr.log.Info("Got context done")
			break
		}

		mgr.log.Info("received shutdown")

		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		if mgr.httpServer != nil {
			mgr.log.Info("calling http Shutdown()")
			go mgr.httpServer.Shutdown(shutdownCtx)
		}
		if mgr.grpcHealthServer != nil {
			mgr.log.Info("calling grpc Shutdown()")
			go mgr.grpcHealthServer.Shutdown(shutdownCtx)
			// mgr.grpcHealthServer.GracefulStop()
		}
		<-shutdownCtx.Done()

		mgr.log.Info("waiting on group")

		if err := group.Wait(); err != nil {
			mgr.log.With(zap.Error(err)).Error("unable to shutdown")
		}
	}
}

type WrapHttp struct {
	api     http.Handler
	manager *Manager
}

func (wrap *WrapHttp) Start(ctx context.Context) error {
	mgr := wrap.manager
	mgr.log.Info("starting service", zap.String("port", mgr.port))
	server := &http.Server{
		Addr:        ":" + mgr.port,
		Handler:     h2c.NewHandler(mgr.withHeaders(wrap.api), &http2.Server{}),
		BaseContext: func(net.Listener) context.Context { return mgr.ctx },
	}

	mgr.httpServer = server

	return server.ListenAndServe()
	// if err := server.ListenAndServe(); err != nil {
	// 	mgr.log.With(zap.Error(err)).Fatal("failed to serve")
	// }

	// return nil
}

func (mgr *Manager) WrapHttpHandler(api http.Handler) HandlerStart {
	return &WrapHttp{
		api:     api,
		manager: mgr,
	}
}

// func (mgr *Manager) StartConsumerSqs(handler awsutil.TwirpHttpSqsHandler) {
// 	if isLambda() {
// 		lambda.StartWithContext(mgr.ctx, handler)
// 	} else {
// 		panic("sqs consumer not supported")
// 	}
// }

func (mgr *Manager) startGrpcHealthCheck(ctx context.Context) error {
	if mgr.grpcHealthPort == "" {
		return nil
	}
	log := mgr.log.With(zap.String("port", mgr.grpcHealthPort))

	log.Info("Starting health service")

	// server := grpcServer{mgr}

	// lis, err := net.Listen("tcp", ":"+mgr.grpcHealthPort)
	// if err != nil {
	// 	mgr.log.With(zap.Error(err)).Error("failed to listen on port")
	// 	return err
	// }

	mux := http.NewServeMux()
	checker := grpchealth.NewStaticChecker()
	mux.Handle(grpchealth.NewHandler(checker))
	// healthcheck := health.NewServer()
	// s := grpc.NewServer()
	// healthgrpc.RegisterHealthServer(s, healthcheck)
	// log.Info("gRPC healthcheck server is running")

	server := &http.Server{
		Addr:        ":" + mgr.grpcHealthPort,
		Handler:     h2c.NewHandler(mux, &http2.Server{}),
		BaseContext: func(net.Listener) context.Context { return mgr.ctx },
	}

	mgr.grpcHealthServer = server
	return server.ListenAndServe()

	// if err := s.Serve(lis); err != http.ErrServerClosed {
	// log.With(zap.Error(err)).Error("failed to serve")
	// }

	// return nil
}

func (mgr *Manager) withHeaders(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), HttpHeaderCtxKey, r.Header)

		base.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func (mgr *Manager) healthCheck(w http.ResponseWriter, r *http.Request) {
// 	if mgr.healthPrefix == "" || !strings.HasPrefix(r.URL.Path, mgr.healthPrefix) || r.Method != "GET" {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	path := strings.Trim(strings.TrimPrefix(r.URL.Path, mgr.healthPrefix), "/")
// 	if path == "ready" || path == "live" {
// 		fmt.Println("Got ", path, " check")
// 		w.WriteHeader(http.StatusOK)
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

func isLambda() bool {
	_, found := os.LookupEnv("LAMBDA_TASK_ROOT")

	return found
}

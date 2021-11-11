package todo

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_tracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/grpcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Server() {
	logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)
	runServer()
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func runServer() {
	ctx := context.Background()
	port := getenv("PORT", "14586")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	zapMiddleware := middleware.NewZapLogger(logger.ZapLogger)
	authMiddleware := middleware.NewAuthenticator(os.Getenv("JWT_SECRET"))

	// add middleware
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_tracing.StreamServerInterceptor(),
			zapMiddleware.StreamServerInterceptor(),
			authMiddleware.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_tracing.UnaryServerInterceptor(),
			zapMiddleware.UnaryServerInterceptor(),
			authMiddleware.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	// opts = middleware.AddLogging(logger.Log, opts)
	// opts = middleware.AddAuthentication(os.Getenv("JWT_SECRET"), opts)

	grpcServer := grpc.NewServer(opts...)

	// attach the Todo service
	s := NewTodoServer()
	publicapi.RegisterTodoServiceServer(grpcServer, s)
	grpc_health_v1.RegisterHealthServer(grpcServer, grpcutil.NewServer())

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.ZapLogger.Info("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	logger.ZapLogger.Info("staring gRPC server... port:" + port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

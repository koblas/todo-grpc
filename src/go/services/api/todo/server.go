package todo

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_tracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Server() {
	logger.Init(-1, time.RFC3339Nano)
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

	zapMiddleware := middleware.NewZapLogger(logger.Log)
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

	server := grpc.NewServer(opts...)

	// attach the Todo service
	s := NewTodoServer()
	publicapi.RegisterTodoServiceServer(server, s)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Info("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	logger.Log.Info("staring gRPC server... port:" + port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

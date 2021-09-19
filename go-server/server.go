package main

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	"github.com/koblas/grpc-todo/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	logger.Init(-1, time.RFC3339Nano)
	runServer()
}

func runServer() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 14586))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	server := grpc.NewServer(opts...)

	// attach the Todo service
	s := todo.Server{}
	todo.RegisterTodoServiceServer(server, &s)

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

	logger.Log.Info("staring gRPC server...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

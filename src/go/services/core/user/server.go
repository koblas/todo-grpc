package user

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	"github.com/koblas/grpc-todo/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Server() {
	logger.Init(-1, time.RFC3339Nano)
	runServer()
}

func runServer() {
	ctx := context.Background()
	port := util.Getenv("PORT", "14586")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	grpcServer := grpc.NewServer(opts...)

	// attach the Todo service
	s := NewUserServer()
	genpb.RegisterUserServiceServer(grpcServer, s)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Info("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	logger.Log.Info("staring gRPC server... port=" + port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

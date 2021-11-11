package user

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
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
	logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)
	runServer()
}

func runServer() {
	log := logger.NewZap(logger.LevelInfo)
	ctx := context.Background()
	port := util.Getenv("PORT", "14586")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.With(err).Fatal("failed to listen")
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.ZapLogger, opts)

	grpcServer := grpc.NewServer(opts...)

	// attach the Todo service
	s := NewUserServer(log)
	genpb.RegisterUserServiceServer(grpcServer, s)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Info("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Info("staring gRPC server... port=" + port)

	if err := grpcServer.Serve(lis); err != nil {
		log.With(err).Fatal("failed to serve: %s", err)
	} else {
		log.Info("Server started successfully")
	}
}

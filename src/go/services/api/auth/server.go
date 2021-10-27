package auth

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
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

func connectUserService() (core.UserServiceClient, *grpc.ClientConn) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	host := util.Getenv("USER_SERVICE_ADDR", ":13001")
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		fmt.Println(err)
	}

	svc := core.NewUserServiceClient(conn)

	return svc, conn
}

func runServer() {
	ctx := context.Background()
	port := util.Getenv("PORT", "14586")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Connect to the user service
	userService, userConn := connectUserService()
	defer userConn.Close()
	s := NewAuthenticationServer(userService)

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	server := grpc.NewServer(opts...)

	// attach the Todo service
	publicapi.RegisterAuthenticationServiceServer(server, &s)

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

	logger.Log.Info("staring gRPC server... port=" + port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

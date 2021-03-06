package extauth

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/koblas/grpc-todo/pkg/grpcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	authToken "github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"

	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/health/grpc_health_v1"

	// core "github.com/envoyproxy/go-control-plane/envoy/api/v3/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"

	// rpc "github.com/gogo/googleapis/google/rpc"
	"github.com/koblas/grpc-todo/pkg/util"
	// rpc "google.golang.org/genproto/googleapis/rpc"
	// status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
)

// empty struct because this isn't a fancy example
type AuthorizationServer struct {
	maker authToken.Maker
}

func NewAuthorizationServer(secret string) (*AuthorizationServer, error) {
	maker, err := authToken.NewJWTMaker(secret)
	if err != nil {
		return nil, err
	}

	return &AuthorizationServer{maker}, nil
}

// inject a header that can be used for future rate limiting
func (a *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	logger.ZapLogger.Info("Doing check")
	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string
	if ok {
		splitToken = strings.Split(authHeader, "Bearer ")
	}

	if len(splitToken) != 2 {
		return &auth.CheckResponse{
			Status: &status.Status{
				Code: int32(code.Code_UNAUTHENTICATED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: &auth.DeniedHttpResponse{
					Status: &v3.HttpStatus{
						Code: v3.StatusCode_Unauthorized,
					},
					Body: "Need an Authorization Header with a Bearer token! #secure",
				},
			},
		}, nil
	}

	_, err := a.maker.VerifyToken(splitToken[1])

	if err != nil {
		return &auth.CheckResponse{
			Status: &status.Status{
				Code: int32(code.Code_UNAUTHENTICATED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: &auth.DeniedHttpResponse{
					Status: &v3.HttpStatus{
						Code: v3.StatusCode_Unauthorized,
					},
					Body: "Token didn't verify",
				},
			},
		}, nil
	}

	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				// Headers: []*v31.HeaderValueOption{
				// 	{
				// 		Header: &v31.HeaderValue{
				// 			Key:   "x-ext-auth-ratelimit",
				// 			Value: tokenSha,
				// 		},
				// 	},
				// },
			},
		},
	}, nil

}

func runServer() {
	ctx := context.Background()
	port := util.Getenv("PORT", "14586")
	// create a TCP listener on port 4000
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	opts := []grpc.ServerOption{}
	// add middleware
	opts = middleware.AddLogging(logger.ZapLogger, opts)

	grpcServer := grpc.NewServer(opts...)
	authServer, err := NewAuthorizationServer(util.Getenv("JWT_SECRET", ""))
	if err != nil {
		log.Fatalf("Failed to create AuthServer: %v", err)
	}
	auth.RegisterAuthorizationServer(grpcServer, authServer)
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

func Server() {
	logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)
	runServer()
}

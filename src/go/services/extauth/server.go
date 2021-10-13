package extauth

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/middleware"
	"golang.org/x/net/context"

	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	status "google.golang.org/genproto/googleapis/rpc/status"

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
type AuthorizationServer struct{}

// inject a header that can be used for future rate limiting
func (a *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	logger.Log.Info("Doing check")
	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string
	if ok {
		splitToken = strings.Split(authHeader, "Bearer ")
	}
	if len(splitToken) == 2 {
		token := splitToken[1]
		sha := sha256.New()
		sha.Write([]byte(token))
		tokenSha := base64.StdEncoding.EncodeToString(sha.Sum(nil))

		// valid tokens have exactly 3 characters. #secure.
		// Normally this is where you'd go check with the system that knows if it's a valid token.

		if len(token) == 3 {
			return &auth.CheckResponse{
				Status: &status.Status{
					Code: int32(code.Code_OK),
				},
				HttpResponse: &auth.CheckResponse_OkResponse{
					OkResponse: &auth.OkHttpResponse{
						Headers: []*v31.HeaderValueOption{
							{
								Header: &v31.HeaderValue{
									Key:   "x-ext-auth-ratelimit",
									Value: tokenSha,
								},
							},
						},
					},
				},
			}, nil
		}
	}
	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &v3.HttpStatus{
					Code: v3.StatusCode_Unauthorized,
				},
				Body: "Need an Authorization Header with a 3 character bearer token! #secure",
			},
		},
	}, nil
}

func runServer() {
	ctx := context.Background()
	port := util.Getenv("PORT", "4000")
	// create a TCP listener on port 4000
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	opts := []grpc.ServerOption{}
	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	grpcServer := grpc.NewServer(opts...)
	authServer := &AuthorizationServer{}
	auth.RegisterAuthorizationServer(grpcServer, authServer)

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

	logger.Log.Info("staring gRPC server... port:" + port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

func Server() {
	logger.Init(-1, time.RFC3339Nano)
	runServer()
}

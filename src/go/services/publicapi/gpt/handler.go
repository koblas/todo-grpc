package gpt

import (
	"log"

	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type GptServer struct {
	jwtMaker tokenmanager.Maker
}

type Option func(*GptServer)

func NewGptServer(config Config, opts ...Option) *GptServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := GptServer{
		jwtMaker: maker,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *GptServer) getUserId(ctx context.Context) (string, error) {
	return tokenmanager.UserIdFromContext(ctx, svc.jwtMaker)
}

// SayHello generates response to a Ping request
func (svc *GptServer) Create(ctx context.Context, params *apipb.GptCreateParams) (*apipb.GptCreateResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GptCreate BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)
	log.Info("Looking up user")

	text := params.Prompt

	return &apipb.GptCreateResponse{
		Text: text,
	}, nil
}

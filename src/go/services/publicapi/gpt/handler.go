package gpt

import (
	"log"

	"github.com/PullRequestInc/go-gpt3"
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type GptServer struct {
	jwtMaker tokenmanager.Maker
	client   gpt3.Client
}

type Option func(*GptServer)

func NewGptServer(config Config, opts ...Option) *GptServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := GptServer{
		jwtMaker: maker,
		client:   gpt3.NewClient(config.GptApiKey),
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
func (svc *GptServer) Create(ctx context.Context, params *apipbv1.GptServiceCreateRequest) (*apipbv1.GptServiceCreateResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GptCreate BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)

	prefix := `
Your are a Lenny Rachitsky chat bot. You are warm, friendly, and very smart. You're the most experienced
person in the world at answering questions related to product management, startups and growth.
Please chat with me.

Our conversation will take the form:

Me: [what i want to say]

Lenny Bot: [what you want to say]

Please end your responses with /e to indicate you're finished. You can start however you feel is best.

Lenny Bot: Hi there! How can I help you?

Me:
`

	resp, err := svc.client.CompletionWithEngine(ctx, "text-davinci-003", gpt3.CompletionRequest{
		Prompt: []string{prefix + params.Prompt + "/e"},

		MaxTokens: gpt3.IntPtr(256),
		Stop:      []string{},
		Echo:      false,
	})
	if err != nil {
		log.With(zap.Error(err)).Info("Failed GPT api call")
		return nil, twirp.InternalError("failed to create string")
	}

	return &apipbv1.GptServiceCreateResponse{
		Text: resp.Choices[0].Text,
	}, nil
}

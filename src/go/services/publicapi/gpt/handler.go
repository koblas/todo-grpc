package gpt

import (
	"errors"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type GptServer struct {
	userHelper interceptors.UserIdFromContext
	client     gpt3.Client
	gptApiKey  string
}

type Option func(*GptServer)

func WithGetUserId(helper interceptors.UserIdFromContext) Option {
	return func(svr *GptServer) {
		svr.userHelper = helper
	}
}

func WithGpiApiKey(apiKey string) Option {
	return func(svr *GptServer) {
		svr.gptApiKey = apiKey
	}
}

func NewGptServer(opts ...Option) *GptServer {
	svr := GptServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.userHelper == nil {
		panic("no user helper provided")
	}
	if svr.gptApiKey == "" {
		panic("no GPT api key provided")
	}

	svr.client = gpt3.NewClient(svr.gptApiKey)

	return &svr
}

// SayHello generates response to a Ping request
func (svc *GptServer) Create(ctx context.Context, params *connect.Request[apiv1.GptServiceCreateRequest]) (*connect.Response[apiv1.GptServiceCreateResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("GptCreate BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("missing userid"))
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
		Prompt: []string{prefix + params.Msg.Prompt + "/e"},

		MaxTokens: gpt3.IntPtr(256),
		Stop:      []string{},
		Echo:      false,
	})
	if err != nil {
		log.With(zap.Error(err)).Info("Failed GPT api call")
		return nil, bufcutil.InternalError(err, "failed to create string")
	}

	return connect.NewResponse(&apiv1.GptServiceCreateResponse{
		Text: resp.Choices[0].Text,
	}), nil
}

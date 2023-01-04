package file

import (
	"log"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/context"
)

// var validate = validator.New()

// Server represents the gRPC server
type FileServer struct {
	file     core.FileService
	jwtMaker tokenmanager.Maker
}

type Option func(*FileServer)

func WithFileService(client core.FileService) Option {
	return func(svr *FileServer) {
		svr.file = client
	}
}

func NewFileServer(config Config, opts ...Option) *FileServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := FileServer{
		jwtMaker: maker,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *FileServer) getUserId(ctx context.Context) (string, error) {
	return tokenmanager.UserIdFromContext(ctx, svc.jwtMaker)
}

func (svc *FileServer) UploadUrl(ctx context.Context, input *publicapi.UploadUrlParams) (*publicapi.UploadUrlResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("UploadUrl BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)

	log.Info("Just a test")

	req := core.FileUploadUrlParams{
		Type: input.Type,
	}

	res, err := svc.file.UploadUrl(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &publicapi.UploadUrlResponse{
		Url: res.Url,
	}, err
}

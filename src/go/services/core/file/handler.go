package file

import (
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type FileServer struct {
	files  FileStore
	pubsub core.FileEventbus
	kms    key_manager.Encoder
}

type Option func(*FileServer)

func WithFileStore(store FileStore) Option {
	return func(cfg *FileServer) {
		cfg.files = store
	}
}

func WithProducer(bus core.FileEventbus) Option {
	return func(cfg *FileServer) {
		cfg.pubsub = bus
	}
}

func NewFileServer(opts ...Option) *FileServer {
	svr := FileServer{
		kms: key_manager.NewSecureClear(),
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (s *FileServer) UploadUrl(ctx context.Context, params *genpb.FileUploadUrlParams) (*genpb.FileUploadUrlResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("UploadUrl")

	result, err := s.files.CreateUploadUrl(ctx, params.UserId, params.Type)
	if err != nil {
		log.With(zap.Error(err)).Error("UploadUrl failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &genpb.FileUploadUrlResponse{
		Url: result,
	}, nil
}

func (s *FileServer) VerifyUrl(ctx context.Context, params *genpb.FileVerifyUrlParams) (*genpb.FileVerifyUrlResponse, error) {
	log := logger.FromContext(ctx)

	result, err := s.files.LookupUploadUrl(ctx, params.Url)
	if err != nil {
		log.With(zap.Error(err)).Error("VerifyUrl failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &genpb.FileVerifyUrlResponse{
		Type:   result.FileType,
		UserId: result.UserId,
	}, nil
}

// Put - write non-authenticated bytes to a persistent store and triggers a notification.
// If we have S3 (or similar) then this is not needed
func (s *FileServer) Put(ctx context.Context, params *genpb.FilePutParams) (*genpb.FilePutResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "Put"))

	name, entry, err := s.files.StoreFile(ctx, params.Path, params.Query, params.Data)
	if err != nil {
		log.With(zap.Error(err)).Error("StoreFile failed")
		return nil, twirp.InternalErrorWith(err)
	}
	log.With(zap.String("path", params.Path)).Info("Accepted file")

	s.pubsub.FileUploaded(ctx, &genpb.FileUploadEvent{
		IdemponcyId: entry.Id,
		UserId:      &entry.UserId,
		FileType:    entry.FileType,
		Url:         entry.InternalUrl,
	})

	return &genpb.FilePutResponse{
		Path: name,
	}, nil
}

func (s *FileServer) Get(ctx context.Context, params *genpb.FileGetParams) (*genpb.FileGetResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "Get"))

	bytes, err := s.files.GetFile(ctx, params.Path)
	if err != nil {
		log.With(zap.Error(err)).Error("GetFile failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &genpb.FileGetResponse{
		Data: bytes,
	}, nil
}

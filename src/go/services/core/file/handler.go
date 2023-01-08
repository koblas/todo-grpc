package file

import (
	"strings"

	"github.com/google/uuid"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type FileServer struct {
	files  FileStore
	pubsub genpb.FileEventbus
	kms    key_manager.Encoder
}

var _ genpb.FileService = (*FileServer)(nil)

type Option func(*FileServer)

func WithFileStore(store FileStore) Option {
	return func(cfg *FileServer) {
		cfg.files = store
	}
}

func WithProducer(bus genpb.FileEventbus) Option {
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
	if params.UserId == "" {
		return nil, twirp.InvalidArgumentError("userId", "missing")
	}
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
func (s *FileServer) Upload(ctx context.Context, params *genpb.FileUploadParams) (*genpb.FileUploadResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "Upload"))

	if err := s.files.VerifyUploadUrl(ctx, params.Path, params.Query); err != nil {
		log.Info("StoreFile failed - signature mismatch")
		return nil, twirp.NewError(twirp.PermissionDenied, "signature mismatch")
	}
	entry, err := s.files.LookupUploadUrl(ctx, params.Path)
	if err != nil {
		log.Info("StoreFile failed - signature mismatch")
		return nil, twirp.NewError(twirp.PermissionDenied, "signature mismatch")
	}
	if _, err := s.files.StoreFile(ctx, params.Path, params.Data); err != nil {
		log.With(zap.Error(err)).Error("StoreFile failed")
		return nil, twirp.InternalErrorWith(err)
	}
	log.With(zap.String("path", params.Path)).Info("Accepted file")

	s.pubsub.FileUploaded(ctx, &genpb.FileUploadEvent{
		IdemponcyId: entry.Id,
		Info: &genpb.FileUploadInfo{
			UserId:   &entry.UserId,
			FileType: entry.FileType + ":upload",
			Url:      entry.InternalUrl,
		},
	})

	return &genpb.FileUploadResponse{
		Path: params.Path,
	}, nil
}

func (s *FileServer) Put(ctx context.Context, params *genpb.FilePutParams) (*genpb.FilePutResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "Put"))

	path := strings.Join([]string{
		params.UserId,
		params.FileType,
		uuid.NewString() + params.Suffix,
	}, "/")

	url, err := s.files.StoreFile(ctx, path, params.Data)
	if err != nil {
		log.With(zap.Error(err)).Error("StoreFile failed")
	}

	return &genpb.FilePutResponse{
		Path: url,
	}, nil
}

func (s *FileServer) Get(ctx context.Context, params *genpb.FileGetParams) (*genpb.FileGetResponse, error) {
	log := logger.FromContext(ctx).With(
		zap.String("method", "Get"),
		zap.String("path", params.Path),
	)

	bytes, err := s.files.GetFile(ctx, params.Path)
	if err != nil {
		if err == ErrorLookupNotFound {
			log.Info("file not found")
			return nil, twirp.NotFoundError("")
		}
		log.With(zap.Error(err)).Error("GetFile failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &genpb.FileGetResponse{
		Data: bytes,
	}, nil
}

package file

import (
	"strings"

	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/oklog/ulid/v2"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type FileServer struct {
	files  FileStore
	pubsub corepbv1.FileEventbus
	kms    key_manager.Encoder
}

var _ corepbv1.FileService = (*FileServer)(nil)

type Option func(*FileServer)

func WithFileStore(store FileStore) Option {
	return func(cfg *FileServer) {
		cfg.files = store
	}
}

func WithProducer(bus corepbv1.FileEventbus) Option {
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

func (s *FileServer) UploadUrl(ctx context.Context, params *corepbv1.FileUploadUrlParams) (*corepbv1.FileUploadUrlResponse, error) {
	if params.UserId == "" {
		return nil, twirp.InvalidArgumentError("userId", "missing")
	}
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("UploadUrl")

	result, err := s.files.CreateUploadUrl(logger.ToContext(ctx, log), params.UserId, params.Type)
	if err != nil {
		log.With(zap.Error(err)).Error("UploadUrl failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &corepbv1.FileUploadUrlResponse{
		Url: result,
	}, nil
}

func (s *FileServer) VerifyUrl(ctx context.Context, params *corepbv1.FileVerifyUrlParams) (*corepbv1.FileVerifyUrlResponse, error) {
	log := logger.FromContext(ctx)

	result, err := s.files.LookupUploadUrl(ctx, params.Url)
	if err != nil {
		log.With(zap.Error(err)).Error("VerifyUrl failed")
		return nil, twirp.InternalErrorWith(err)
	}

	return &corepbv1.FileVerifyUrlResponse{
		Type:   result.FileType,
		UserId: result.UserId,
	}, nil
}

// Put - write non-authenticated bytes to a persistent store and triggers a notification.
// If we have S3 (or similar) then this is not needed
func (s *FileServer) Upload(ctx context.Context, params *corepbv1.FileUploadParams) (*corepbv1.FileUploadResponse, error) {
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

	s.pubsub.FileUploaded(ctx, &corepbv1.FileUploadEvent{
		IdemponcyId: entry.Id,
		Info: &corepbv1.FileUploadInfo{
			UserId:   &entry.UserId,
			FileType: entry.FileType + ":upload",
			Url:      entry.InternalUrl,
		},
	})

	return &corepbv1.FileUploadResponse{
		Path: params.Path,
	}, nil
}

func (s *FileServer) Put(ctx context.Context, params *corepbv1.FilePutParams) (*corepbv1.FilePutResponse, error) {
	log := logger.FromContext(ctx).With(
		zap.String("method", "Put"),
		zap.Int("length", len(params.Data)),
	)

	path := strings.Join([]string{
		params.UserId,
		params.FileType,
		ulid.Make().String() + params.Suffix,
	}, "/")

	url, err := s.files.StoreFile(ctx, path, params.Data)
	if err != nil {
		log.With(zap.Error(err)).Error("StoreFile failed")
	}

	log.Info("File stored success")

	return &corepbv1.FilePutResponse{
		Path: url,
	}, nil
}

func (s *FileServer) Get(ctx context.Context, params *corepbv1.FileGetParams) (*corepbv1.FileGetResponse, error) {
	log := logger.FromContext(ctx).With(
		zap.String("method", "Get"),
		zap.String("path", params.Path),
	)

	data, err := s.files.GetFile(ctx, params.Path)
	if err != nil {
		if err == ErrorLookupNotFound {
			log.Info("file not found")
			return nil, twirp.NotFoundError("")
		}
		log.With(zap.Error(err)).Error("GetFile failed")
		return nil, twirp.InternalErrorWith(err)
	}
	log.With(zap.Int("length", len(data))).Info("Sending data")

	return &corepbv1.FileGetResponse{
		Data: data,
	}, nil
}

package file

import (
	"log"

	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// var validate = validator.New()

// Server represents the gRPC server
type FileServer struct {
	// file     corepb.FileService
	uploadBucket string
	file         filestore.Filestore
	jwtMaker     tokenmanager.Maker
}

type Option func(*FileServer)

func WithFileStore(client filestore.Filestore) Option {
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
		jwtMaker:     maker,
		uploadBucket: config.UploadBucket,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *FileServer) getUserId(ctx context.Context) (string, error) {
	return tokenmanager.UserIdFromContext(ctx, svc.jwtMaker)
}

func (svc *FileServer) UploadUrl(ctx context.Context, input *apipb.UploadUrlParams) (*apipb.UploadUrlResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("UploadUrl BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}
	log = log.With("userId", userId)

	req := filestore.FilePutParams{
		Bucket:   svc.uploadBucket,
		UserId:   userId,
		FileType: input.Type + ".upload",
	}

	res, err := svc.file.UploadUrl(ctx, &req)
	if err != nil {
		log.With(zap.Error(err)).Error("Unable to build PUT url")
		return nil, twirp.InternalError("unable to construct url")
	}

	urlStr := res.Url
	// if true {
	// 	u, err := url.Parse(urlStr)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	urlStr = "/minio_" + u.Path + "?" + u.RawQuery
	// }

	return &apipb.UploadUrlResponse{
		Url: urlStr,
	}, nil
}

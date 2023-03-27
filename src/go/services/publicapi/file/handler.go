package file

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// var validate = validator.New()

// Server represents the gRPC server
type FileServer struct {
	// file     corepbv1.FileService
	uploadBucket string
	file         filestore.Filestore
	userHelper   interceptors.UserIdFromContext
}

type Option func(*FileServer)

func WithFileStore(client filestore.Filestore) Option {
	return func(svr *FileServer) {
		svr.file = client
	}
}

func WithGetUserId(helper interceptors.UserIdFromContext) Option {
	return func(svr *FileServer) {
		svr.userHelper = helper
	}
}

func WithUploadBucket(uploadBucket string) Option {
	return func(svr *FileServer) {
		svr.uploadBucket = uploadBucket
	}
}

func NewFileServer(opts ...Option) *FileServer {
	svr := FileServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.uploadBucket == "" {
		panic("no upload bucket provided")
	}
	if svr.userHelper == nil {
		panic("no user helper provided")
	}

	return &svr
}

func (svc *FileServer) UploadUrl(ctx context.Context, input *connect.Request[apiv1.FileServiceUploadUrlRequest]) (*connect.Response[apiv1.FileServiceUploadUrlResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("UploadUrl BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}
	log = log.With("userId", userId)

	req := filestore.FilePutParams{
		Bucket:      svc.uploadBucket,
		UserId:      userId,
		FileType:    input.Msg.Type + ".upload",
		ContentType: input.Msg.ContentType,
	}

	res, err := svc.file.UploadUrl(ctx, &req)
	if err != nil {
		log.With(zap.Error(err)).Error("Unable to build PUT url")
		return nil, bufcutil.InternalError(err, "unable to construct url")
	}

	// urlStr := res.Url
	// if true {
	// 	u, err := url.Parse(urlStr)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	urlStr = "/minio_" + u.Path + "?" + u.RawQuery
	// }

	return connect.NewResponse(&apiv1.FileServiceUploadUrlResponse{
		Url: res.Url,
		Id:  res.Id,
	}), nil
}

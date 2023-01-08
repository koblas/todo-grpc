package fileput

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

// Server represents the gRPC server
type FilePutServer struct {
	file corepb.FileService
}

type Option func(*FilePutServer)

func WithFileService(client corepb.FileService) Option {
	return func(svr *FilePutServer) {
		svr.file = client
	}
}

func NewFilePutServer(config Config, opts ...Option) *FilePutServer {
	svr := FilePutServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *FilePutServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log := logger.FromContext(ctx).With(
		zap.String("method", req.Method),
		zap.String("urlPath", req.URL.Path),
	)

	if req.Method != "PUT" {
		log.Info("Invalid method sent")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	url := req.URL
	contentType := req.Header.Get("content-type")
	contentLength := req.Header.Get("content-length")

	log.With(
		zap.String("contentType", contentType),
		zap.String("contentLength", contentLength),
	).Info("Upload Info")

	data := []byte{}
	buffer := bytes.NewBuffer(data)
	if _, err := buffer.ReadFrom(req.Body); err != nil {
		log.With(zap.Error(err)).Info("Unable to read body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := svc.file.Upload(ctx, &corepb.FileUploadParams{
		Path:        url.Path,
		Query:       url.RawQuery,
		ContentType: contentType,
		Data:        buffer.Bytes(),
	})

	if err != nil {
		log.With(zap.Error(err)).Info("unable to put")

		var twerr twirp.Error
		if errors.As(err, &twerr) && twerr.Code() == twirp.InvalidArgument {
			writer.WriteHeader(http.StatusForbidden)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

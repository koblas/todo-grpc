package fileput

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

// Server represents the gRPC server
type FilePutServer struct {
	file core.FileService
}

type Option func(*FilePutServer)

func WithFileService(client core.FileService) Option {
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
	log := logger.FromContext(ctx)
	log.Info("FILE PUT BEGIN")

	if req.Method != "PUT" {
		log.With(zap.String("method", req.Method)).Info("Invalid method sent")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	url := req.URL
	contentType := req.Header.Get("content-type")
	contentLength := req.Header.Get("content-length")

	log.With(
		zap.String("urlPath", url.Path),
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

	_, err := svc.file.Put(ctx, &core.FilePutParams{
		Path:        url.Path,
		Query:       url.RawQuery,
		ContentType: contentType,
		Data:        buffer.Bytes(),
	})

	if err != nil {
		log.With(zap.Error(err)).Info("unable to put")

		var twerr twirp.Error
		if errors.As(err, &twerr) {
			if twerr.Code() == twirp.InvalidArgument {
				writer.WriteHeader(http.StatusForbidden)
			}
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

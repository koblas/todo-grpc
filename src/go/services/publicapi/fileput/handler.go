package fileput

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"strings"

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

	switch req.Method {
	case "PUT":
		svc.handlePUT(ctx, log, writer, req)
	case "GET":
		svc.handleGET(ctx, log, writer, req)
	default:
		log.Info("Invalid method sent")
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *FilePutServer) handleGET(ctx context.Context, log logger.Logger, writer http.ResponseWriter, req *http.Request) {
	url := req.URL

	result, err := svc.file.Get(ctx, &corepb.FileGetParams{
		Path: url.Path,
	})
	if err != nil {
		log.With(zap.Error(err)).Info("unable to get")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if strings.HasSuffix(url.Path, ".png") {
		writer.Header().Add("content-type", "image/png")
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(result.Data)
}

func (svc *FilePutServer) handlePUT(ctx context.Context, log logger.Logger, writer http.ResponseWriter, req *http.Request) {
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

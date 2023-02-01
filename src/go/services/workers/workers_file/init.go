package workers_file

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
)

type SqsConsumerBuilder func(WorkerConfig) corepbv1.TwirpServer

type Worker struct {
	Stream    string
	GroupName string
	Build     SqsConsumerBuilder
}

// Some generic handling of definition

type WorkerConfig struct {
	config      Config
	onlyHandler string
	pubsub      corepbv1.FileEventbus
	// fileService corepbv1.FileService
	fileService filestore.Filestore
	userService corepbv1.UserService
}

type Option func(*WorkerConfig)

func WithOnly(item string) Option {
	return func(cfg *WorkerConfig) {
		cfg.onlyHandler = item
	}
}

func WithProducer(bus corepbv1.FileEventbus) Option {
	return func(cfg *WorkerConfig) {
		cfg.pubsub = bus
	}
}

func WithFileService(svc filestore.Filestore) Option {
	return func(cfg *WorkerConfig) {
		cfg.fileService = svc
	}
}

func WithUserService(svc corepbv1.UserService) Option {
	return func(cfg *WorkerConfig) {
		cfg.userService = svc
	}
}

func buildServiceConfig(config Config, opts ...Option) WorkerConfig {
	cfg := WorkerConfig{
		config: config,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

var workers = []Worker{}

type HandlerData struct {
	group   string
	handler corepbv1.TwirpServer
}

type Handler interface {
	GroupName() string
	Handler() corepbv1.TwirpServer
}

func (h *HandlerData) GroupName() string {
	return h.group
}
func (h *HandlerData) Handler() corepbv1.TwirpServer {
	return h.handler
}

func BuildHandlers(config Config, opts ...Option) []manager.MsgHandler {
	handlers := []manager.MsgHandler{}

	cfg := buildServiceConfig(config, opts...)

	for _, worker := range workers {
		if cfg.onlyHandler != "" && cfg.onlyHandler != worker.Stream {
			continue
		}
		data := HandlerData{
			group:   worker.GroupName,
			handler: worker.Build(cfg),
		}

		handlers = append(handlers, manager.MsgHandler(&data))
	}

	return handlers
}

func GetHandlers(config Config, opts ...Option) http.HandlerFunc {
	handlers := []corepbv1.TwirpServer{}

	cfg := buildServiceConfig(config, opts...)

	for _, worker := range workers {
		if cfg.onlyHandler != "" && cfg.onlyHandler != worker.Stream {
			continue
		}

		handlers = append(handlers, worker.Build(cfg))
	}

	return func(w http.ResponseWriter, req *http.Request) {
		// We need to copy the input such that we can read multiple times
		buf := bytes.Buffer{}
		_, err := io.Copy(&buf, req.Body)
		if err != nil {
			// TODO
			return
		}

		for _, handler := range handlers {
			if !strings.HasPrefix(req.URL.Path, handler.PathPrefix()) {
				continue
			}

			writer := httptest.NewRecorder()
			reqCopy := *req
			reqCopy.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))

			handler.ServeHTTP(writer, &reqCopy)

			res := writer.Result()
			if res.StatusCode != http.StatusOK {
				log := logger.FromContext(req.Context())
				buf, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
				log.With("statusCode", res.StatusCode).With("statusMsg", string(buf)).Info("handler invoke error")
			}
		}
	}
}

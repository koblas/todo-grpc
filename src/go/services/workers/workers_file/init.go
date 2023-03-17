package workers_file

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/filestore"
)

type SqsConsumerBuilder func(WorkerConfig) http.Handler

type Worker struct {
	Stream    string
	GroupName string
	Build     SqsConsumerBuilder
}

// Some generic handling of definition

type WorkerConfig struct {
	config      Config
	onlyHandler string
	pubsub      corev1connect.FileEventbusServiceClient
	// fileService corepbv1.FileService
	fileService filestore.Filestore
	userService corev1connect.UserServiceClient
}

type Option func(*WorkerConfig)

func WithOnly(item string) Option {
	return func(cfg *WorkerConfig) {
		cfg.onlyHandler = item
	}
}

func WithProducer(bus corev1connect.FileEventbusServiceClient) Option {
	return func(cfg *WorkerConfig) {
		cfg.pubsub = bus
	}
}

func WithFileService(svc filestore.Filestore) Option {
	return func(cfg *WorkerConfig) {
		cfg.fileService = svc
	}
}

func WithUserService(svc corev1connect.UserServiceClient) Option {
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

func BuildHandlers(config Config, opts ...Option) []http.Handler {
	handlers := []http.Handler{}

	cfg := buildServiceConfig(config, opts...)

	for _, worker := range workers {
		if cfg.onlyHandler != "" && cfg.onlyHandler != worker.Stream {
			continue
		}

		handlers = append(handlers, worker.Build(cfg))
	}

	return handlers
}

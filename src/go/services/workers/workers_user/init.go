package workers_user

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
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
	sendEmail   corev1connect.SendEmailServiceClient
}

type Option func(*WorkerConfig)

func WithOnly(item string) Option {
	return func(cfg *WorkerConfig) {
		cfg.onlyHandler = item
	}
}

func WithSendEmail(sender corev1connect.SendEmailServiceClient) Option {
	return func(cfg *WorkerConfig) {
		cfg.sendEmail = sender
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

type WorkerHandler struct {
	group string
}

func GetHandler(config Config, opts ...Option) []http.Handler {
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

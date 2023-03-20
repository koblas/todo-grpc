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
	config    Config
	sendEmail corev1connect.SendEmailServiceClient
}

type Option func(*WorkerConfig)

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

// type WorkerHandler struct {
// 	group string
// }

func GetHandler(config Config, opts ...Option) map[string]http.Handler {
	handlers := map[string]http.Handler{}

	cfg := buildServiceConfig(config, opts...)

	for _, worker := range workers {
		handlers[worker.GroupName] = worker.Build(cfg)
	}

	return handlers
}

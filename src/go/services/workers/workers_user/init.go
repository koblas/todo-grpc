package workers_user

import (
	"net/http"

	"github.com/koblas/grpc-todo/gen/core/send_email/v1/send_emailv1connect"
)

type SqsConsumerBuilder func(WorkerConfig) http.Handler

type Worker struct {
	Stream    string
	GroupName string
	Build     SqsConsumerBuilder
}

// Some generic handling of definition

type WorkerConfig struct {
	config struct {
		UrlBase string
	}
	sendEmail send_emailv1connect.SendEmailServiceClient
}

type Option func(*WorkerConfig)

func WithSendEmail(sender send_emailv1connect.SendEmailServiceClient) Option {
	return func(cfg *WorkerConfig) {
		cfg.sendEmail = sender
	}
}

func WithUrlBase(base string) Option {
	return func(cfg *WorkerConfig) {
		cfg.config.UrlBase = base
	}
}

func buildServiceConfig(opts ...Option) WorkerConfig {
	cfg := WorkerConfig{}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

var workers = []Worker{}

// type WorkerHandler struct {
// 	group string
// }

func GetHandler(opts ...Option) map[string]http.Handler {
	handlers := map[string]http.Handler{}

	cfg := buildServiceConfig(opts...)

	for _, worker := range workers {
		handlers[worker.GroupName] = worker.Build(cfg)
	}

	return handlers
}

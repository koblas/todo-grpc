package workers

import (
	awsbus "github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/twpb/core"
)

type SqsConsumerBuilder func(SsmConfig, ...Option) awsbus.SqsConsumerFunc

type Worker struct {
	Stream    string
	GroupName string
	Build     SqsConsumerBuilder
}

// Some generic handling of definition

type WorkerConfig struct {
	config    SsmConfig
	sendEmail core.SendEmailService
}

type Option func(*WorkerConfig)

func WithSendEmail(sender core.SendEmailService) Option {
	return func(cfg *WorkerConfig) {
		cfg.sendEmail = sender
	}
}

func buildService(config SsmConfig, opts ...Option) WorkerConfig {
	cfg := WorkerConfig{
		config: config,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

var workers = []Worker{}

func GetWorkers() []Worker {
	return workers
}

// func startWorker(log logger.Logger, item Worker) {
// 	c, err := redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
// 		VisibilityTimeout: 60 * time.Second,
// 		BlockingTimeout:   5 * time.Second,
// 		ReclaimInterval:   1 * time.Second,
// 		BufferSize:        100,
// 		Concurrency:       10,
// 		GroupName:         item.GroupName,
// 		RedisOptions: &redisqueue.RedisOptions{
// 			Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	processor := func(msg *redisqueue.Message) error {
// 		ctx := context.Background()
// 		return item.Process(logger.ToContext(ctx, log), msg)
// 	}

// 	c.Register(item.Stream, processor)

// 	go func() {
// 		for err := range c.Errors {
// 			log.With(err).Error("consumer error")
// 		}
// 	}()

// 	c.Run()
// }

// func RunWorkers() {
// 	logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)

// 	group := sync.WaitGroup{}

// 	logger := logger.NewZap(logger.LevelDebug)
// 	logger.Info("Starting all worker")

// 	for _, item := range workers {
// 		group.Add(1)

// 		go func(entry Worker) {
// 			logger.With("stream", entry.Stream, "streamGroup", entry.GroupName).Info("Starting worker")
// 			defer group.Done()
// 			startWorker(logger.With(
// 				"workerStream", entry.Stream,
// 				"workerGroup", entry.GroupName,
// 			), entry)
// 		}(item)
// 	}

// 	group.Wait()
// }

package redis

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/robinjoseph08/redisqueue"
)

const CONTENT_TRANSFER_ENCODING = "content-transfer-encoding"

type redisBus struct {
	pubsub *redisqueue.Producer
	topic  string
}

var _ eventbus.Producer = (*redisBus)(nil)

func NewRedisProducer(topic string) (eventbus.Producer, error) {
	pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
		},
	})
	if err != nil {
		return nil, err
	}

	return &redisBus{
		topic:  topic,
		pubsub: pubsub,
	}, nil
}

func (svc *redisBus) Enqueue(ctx context.Context, msg *eventbus.Message) error {
	attr := msg.Attributes

	var mbody string
	if len(msg.BodyBytes) != 0 {
		attr[CONTENT_TRANSFER_ENCODING] = "base64"
		mbody = base64.StdEncoding.EncodeToString([]byte(msg.BodyBytes))
	} else {
		mbody = msg.BodyString
	}

	mvalue := map[string]interface{}{
		"body":       mbody,
		"attributes": attr,
	}

	return svc.pubsub.Enqueue(&redisqueue.Message{
		Stream: svc.topic,
		Values: mvalue,
	})
}

type redisConsumerConfig struct {
	redisAddr string
}

func NewRedisConsumer(redisAddr string) *redisConsumerConfig {
	return &redisConsumerConfig{
		redisAddr: redisAddr,
	}
}

func (svc *redisConsumerConfig) BuildWorker(groupName string, streamName string) (*eventbus.Consumer, error) {
	c, err := redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
		VisibilityTimeout: 60 * time.Second,
		BlockingTimeout:   5 * time.Second,
		ReclaimInterval:   1 * time.Second,
		BufferSize:        100,
		Concurrency:       10,
		GroupName:         groupName,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: svc.redisAddr,
		},
	})
	if err != nil {
		panic(err)
	}

	messages := make(chan *eventbus.Message)
	errchan := make(chan error)

	processor := func(msg *redisqueue.Message) error {
		var attr map[string]string
		var body string

		if avalue, ok := msg.Values["attributes"]; !ok {
			errchan <- fmt.Errorf("missing attributes on message")
			return nil
		} else if av, ok := avalue.(map[string]string); !ok {
			errchan <- fmt.Errorf("attributes wrong type")
			return nil
		} else {
			attr = av
		}

		if bvalue, ok := msg.Values["body"]; !ok {
			errchan <- fmt.Errorf("missing body on message")
			return nil
		} else if bv, ok := bvalue.(string); !ok {
			errchan <- fmt.Errorf("body wrong type")
			return nil
		} else {
			body = bv
		}

		mout := eventbus.Message{
			Attributes: attr,
		}

		if v, found := attr[CONTENT_TRANSFER_ENCODING]; found && v == "base64" {
			dec, err := base64.StdEncoding.DecodeString(body)
			if err != nil {
				errchan <- err
				return nil
			}

			mout.BodyBytes = dec
		} else {
			mout.BodyString = body
		}

		messages <- &mout

		return nil
	}

	c.Register(streamName, processor)

	go func() {
		for err := range c.Errors {
			errchan <- err
		}
	}()

	go c.Run()

	return &eventbus.Consumer{
		Messages: messages,
		Errors:   errchan,
		Closing:  make(chan struct{}),
	}, nil
}

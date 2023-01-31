package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/renstrom/shortuuid"
	"github.com/robinjoseph08/redisqueue"
	"go.uber.org/zap"
)

type simpleProducer struct {
	client *redisqueue.Producer
	topic  string
}

type simpleConsumer struct {
	client   *redisqueue.Consumer
	ctx      context.Context
	channel  chan eventbus.SimpleMessage
	shutdown chan struct{}
}

// --- Producer

func NewProducer(redisAddr, topic string) *simpleProducer {
	client, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: redisAddr,
		},
	})
	if err != nil {
		return nil
	}

	return &simpleProducer{
		client: client,
		topic:  topic,
	}
}

func (svc *simpleProducer) Write(ctx context.Context, msg *eventbus.SimpleMessage) error {
	id := msg.ID
	if id == "" {
		id = shortuuid.New()
	}

	attr, err := json.Marshal(msg.Attributes)
	if err != nil {
		return err
	}

	return svc.client.Enqueue(&redisqueue.Message{
		Stream: svc.topic,
		Values: map[string]interface{}{
			"id":         id,
			"body":       msg.Body,
			"attributes": attr,
		},
	})
}

// --- Consumer

func NewConsumer(ctx context.Context, redisAddr, topic string) *simpleConsumer {
	client, err := redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
		VisibilityTimeout: 60 * time.Second,
		BlockingTimeout:   5 * time.Second,
		ReclaimInterval:   1 * time.Second,
		BufferSize:        100,
		Concurrency:       10,
		// GroupName:         groupName,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: redisAddr,
		},
	})
	if err != nil {
		panic(err)
	}

	svc := simpleConsumer{
		client:   client,
		ctx:      ctx,
		channel:  make(chan eventbus.SimpleMessage),
		shutdown: make(chan struct{}),
	}

	client.Register(topic, svc.processor)
	go func(c *simpleConsumer) {
		defer func() {
			c.client = nil
			c.shutdown <- struct{}{}
		}()
		c.client.Run()
	}(&svc)

	return &svc
}

func (svc *simpleConsumer) processor(msg *redisqueue.Message) error {
	log := logger.FromContext(svc.ctx).With(zap.String("msg_id", msg.ID))
	message := eventbus.SimpleMessage{}

	if value, found := msg.Values["id"].(string); found {
		message.ID = value
	} else {
		log.Error("message missing id")
		return nil
	}
	if value, found := msg.Values["attributes"].(string); found {
		attr := map[string]string{}
		if err := json.Unmarshal([]byte(value), &attr); err != nil {
			return err
		}
		message.Attributes = attr
	} else {
		log.Error("message missing attributes")
		return nil
	}
	if value, found := msg.Values["body"].(string); found {
		message.Body = value
	} else {
		log.Error("message missing body")
		return nil
	}

	svc.channel <- message

	return nil
}

func (svc *simpleConsumer) Next() (eventbus.SimpleMessage, error) {
	if svc.client == nil {
		return eventbus.SimpleMessage{}, errors.New("channel closed")
	}

	select {
	case msg := <-svc.channel:
		return msg, nil
	case <-svc.shutdown:
		return eventbus.SimpleMessage{}, errors.New("channel closed")
	}
}

func (svc *simpleConsumer) Close() {
	if svc.client == nil {
		return
	}

	svc.client.Shutdown()
}

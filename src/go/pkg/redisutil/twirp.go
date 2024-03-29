package redisutil

import (
	"bytes"
	"context"
	"encoding"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/go-redis/redis/v8"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/pkg/errors"
	"github.com/robinjoseph08/redisqueue"
	"go.uber.org/zap"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	redisAddr string

	redis  *redis.Client
	pubsub *redisqueue.Producer
	queue  map[string]rmq.Queue
}

type payloadHeaders map[string][]string

var _ encoding.BinaryMarshaler = (payloadHeaders)(nil)

func (m payloadHeaders) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *payloadHeaders) UnmarshalJSON(data []byte) error {
	cleanValue := data
	if data[0] == '"' {
		// Double encoded
		unwrap := ""
		if err := json.Unmarshal(data, &unwrap); err != nil {
			return err
		}

		cleanValue = []byte(unwrap)
	}

	value := map[string][]string{}
	if err := json.Unmarshal(cleanValue, &value); err != nil {
		return err
	}

	*m = value

	return nil
}

// General Purpose payload -- keep types "stock" to make sure that redisqueue can
// serialize and deserialize correctly
type queuePayload struct {
	Path    string
	Headers payloadHeaders
	Body    string
}

var _ Client = (*client)(nil)

func NewTwirpRedis(addr string) *client {
	svc := client{
		redisAddr: addr,
		queue:     map[string]rmq.Queue{},
	}

	return &svc
}

func (svc *client) openRedisQueue(name string) (rmq.Queue, error) {
	if queue, found := svc.queue[name]; found {
		return queue, nil
	}

	svc.openRedis()

	errChan := make(chan error)
	connection, err := rmq.OpenConnectionWithRedisClient("twirp-queue", svc.redis, errChan)
	if err != nil {
		return nil, err
	}

	queue, err := connection.OpenQueue(name)
	if err != nil {
		return nil, err
	}

	svc.queue[name] = queue

	return queue, nil
}

func (svc *client) openRedisTopicProducer(name string) (*redisqueue.Producer, error) {
	if svc.pubsub != nil {
		return svc.pubsub, nil
	}

	pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: svc.redisAddr,
		},
	})

	// Not sure this is the best way, but fails if we cannot contact redis
	//  maybe a retry poool?
	if err != nil {
		return nil, err
	}

	svc.pubsub = pubsub

	return pubsub, nil
}

func (svc *client) openRedisTopicConsumer(name string) (*redisqueue.Consumer, error) {
	return redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
		VisibilityTimeout: 60 * time.Second,
		BlockingTimeout:   5 * time.Second,
		ReclaimInterval:   1 * time.Second,
		BufferSize:        100,
		Concurrency:       10,
		GroupName:         name,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: svc.redisAddr,
		},
	})
}

func (svc *client) openRedis() error {
	svc.redis = NewClient(svc.redisAddr)

	return nil
}

func (svc *client) Do(req *http.Request) (*http.Response, error) {
	buf := strings.Builder{}
	_, err := io.Copy(&buf, req.Body)
	if err != nil {
		return nil, err
	}
	bodyString := base64.RawStdEncoding.EncodeToString([]byte(buf.String()))

	mvalue := queuePayload{
		Path:    req.URL.Path,
		Body:    bodyString,
		Headers: payloadHeaders(req.Header),
	}
	payload, err := json.Marshal(mvalue)
	if err != nil {
		return nil, errors.Wrap(err, "twirp-redis-topic: unable to encode payload")
	}

	if req.URL.Scheme == "topic" {
		pubsub, err := svc.openRedisTopicProducer(req.URL.Host)
		if err != nil {
			return nil, errors.Wrap(err, "twirp-redis-topic: unable to connect to producer")
		}

		// jsonHdrs, err := json.Marshal(mvalue.Headers.Header)
		// if err != nil {
		// 	return nil, errors.Wrap(err, "twirp-redis-topic: unable to encode headers")
		// }
		err = pubsub.Enqueue(&redisqueue.Message{
			Stream: req.URL.Host,
			Values: map[string]interface{}{
				"Payload": string(payload),
			},
		})
		if err != nil {
			return nil, errors.Wrap(err, "twirp-redis-topic: failed to send")
		}
	} else if req.URL.Scheme == "queue" {
		queue, err := svc.openRedisQueue(req.URL.Host)
		if err != nil {
			return nil, errors.Wrap(err, "twirp-redis-queue: unable to open queue")
		}
		err = queue.PublishBytes(payload)
		if err != nil {
			return nil, errors.Wrap(err, "twirp-redis-queue: failed to publish")
		}
	} else {
		return nil, errors.New("unimplemented redis scheme")
	}

	res := http.Response{
		StatusCode: http.StatusOK,
	}
	ctype := http.Header(mvalue.Headers).Get("content-type")
	if strings.Contains(ctype, "application/json") {
		res.Body = io.NopCloser(strings.NewReader("{}"))
	} else {
		res.Body = io.NopCloser(bytes.NewReader([]byte{}))
	}

	return &res, nil
}

func buildRequestFromPayload(host string, payloadString string) (http.Request, error) {
	payload := queuePayload{}
	err := json.Unmarshal([]byte(payloadString), &payload)
	if err != nil {
		return http.Request{}, errors.Wrap(err, "redis-topic-consumer: unable to unmarshal json")
	}

	bodyBytes, err := base64.RawStdEncoding.DecodeString(payload.Body)
	if err != nil {
		return http.Request{}, errors.Wrap(err, "redis-topic-consumer: unable to unmarshal base64")
	}

	req := http.Request{
		URL: &url.URL{
			Scheme: "queue",
			Host:   host,
			Path:   payload.Path,
		},
		Method: "POST",
		Header: http.Header(payload.Headers),
		Body:   io.NopCloser(bytes.NewReader(bodyBytes)),
	}

	return req, nil
}

type topicConsumer struct {
	client   *client
	handlers []http.Handler
	Topic    string
}

type queueConsumer struct {
	client   *client
	handlers []http.Handler
	Queue    string
}

func (con *topicConsumer) Start(ctx context.Context) error {
	topicName := con.Topic

	log := logger.FromContext(ctx).With(zap.String("topic", topicName))

	consumer, err := con.client.openRedisTopicConsumer(topicName)
	if err != nil {
		return err
	}

	consumer.Register(topicName, func(msg *redisqueue.Message) error {
		payloadValue, found := msg.Values["Payload"]
		if !found {
			return errors.New("redis-topic-consumer: Missing payload")
		}
		payloadStr, found := payloadValue.(string)
		if !found {
			return errors.New("redis-topic-consumer: payload bad type")
		}

		req, err := buildRequestFromPayload(topicName, payloadStr)
		if err != nil {
			return err
		}

		for _, item := range con.handlers {
			w := httptest.NewRecorder()
			item.ServeHTTP(w, req.WithContext(ctx))

			res := w.Result()
			if res.StatusCode >= http.StatusOK || res.StatusCode < http.StatusBadRequest {
				// ignore
			} else {
				buf, _ := io.ReadAll(io.LimitReader(res.Body, 256))
				log.With("statusCode", res.StatusCode).With("statusMsg", string(buf)).Info("SQS Message error")
			}
		}

		return nil
	})

	go func(c *redisqueue.Consumer) {
		for err := range c.Errors {
			// handle errors accordingly
			log.With(zap.Error(err)).Error("Unable to consume")
		}
	}(consumer)

	consumer.Run()

	return nil
}

func (svc *client) TopicConsumer(ctx context.Context, topic string, handlers []http.Handler) manager.HandlerStart {
	consumer := topicConsumer{
		client:   svc,
		Topic:    topic,
		handlers: handlers,
	}

	return &consumer
}

func (svc *client) QueueConsumer(ctx context.Context, queue string, handlers []http.Handler) manager.HandlerStart {
	consumer := queueConsumer{
		client:   svc,
		Queue:    queue,
		handlers: handlers,
	}

	return &consumer
}

func (con *queueConsumer) Start(ctx context.Context) error {
	// func (svc *client) QueueConsumer(ctx context.Context, queueName string, handler http.Handler) awsutil.TwirpHttpSqsHandler {
	log := logger.FromContext(ctx).With(zap.String("queueName", con.Queue))

	queue, err := con.client.openRedisQueue(con.Queue)

	if err != nil {
		log.With(zap.Error(err)).Error("unable to open queue")
		return err
	}

	if err = queue.StartConsuming(10, time.Second); err != nil {
		log.With(zap.Error(err)).Error("queue.StartConsuming failed")
		return err
	}

	finished := make(chan bool)

	_, err = queue.AddConsumerFunc(con.Queue, func(delivery rmq.Delivery) {
		// defer func() {
		// 	finished <- true
		// }()

		req, err := buildRequestFromPayload(con.Queue, delivery.Payload())
		if err != nil {
			log.With(zap.Error(err)).Error("failed to unmarshal")
			return
		}

		oneSuccess := false
		for _, item := range con.handlers {
			w := httptest.NewRecorder()
			item.ServeHTTP(w, req.WithContext(ctx))

			res := w.Result()
			if res.StatusCode >= http.StatusOK || res.StatusCode < http.StatusBadRequest {
				oneSuccess = true
			} else {
				buf, _ := io.ReadAll(io.LimitReader(res.Body, 256))
				log.With("statusCode", res.StatusCode).With("statusMsg", string(buf)).Info("SQS Message error")
			}
		}

		if oneSuccess {
			if err := delivery.Ack(); err != nil {
				log.With(zap.Error(err)).Error("failed to Ack")
			}
		} else {
			if err := delivery.Reject(); err != nil {
				log.With(zap.Error(err)).Error("failed to Reject")
			}
		}
	})

	if err != nil {
		log.With(zap.Error(err)).Error("failed to add consumer")
		return err
	}

	// Wait for the finished event
	<-finished

	log.Info("all done")

	return nil
}

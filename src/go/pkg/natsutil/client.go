package natsutil

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type client struct {
	conn *nats.Conn
	url  string
}

func NewNatsClient(addr string) *client {
	return &client{
		url: "nats://" + addr,
	}
}

func (svc *client) connect(ctx context.Context) error {
	if svc.conn != nil {
		return nil
	}

	log := logger.FromContext(ctx)

	conn, err := nats.Connect(svc.url,
		// nats.Name("dalong-reply"),
		nats.DisconnectHandler(func(conn *nats.Conn) {
			log.Info("nats.DisconnectHandler")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			log.Info("nats.ReconnectHandler")
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			log.Info("nats.ClosedHandler")
		}),
		nats.DiscoveredServersHandler(func(conn *nats.Conn) {
			log.Info("nats.DiscoveredServersHandler")
		}),
		nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, e error) {
			log.Info("nats.ErrorHandler" + e.Error())
		}),
		nats.DisconnectHandler(func(conn *nats.Conn) {
			log.Info("nats.DisconnectHandler")
		}),
	)

	if err != nil {
		return errors.Wrap(err, "unable to connect to nats")
	}
	svc.conn = conn

	return nil
}

func (svc *client) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	if err := svc.connect(ctx); err != nil {
		return nil, err
	}

	log := logger.FromContext(ctx)

	// Copy the body
	buf := []byte{}
	buffer := bytes.NewBuffer(buf)
	_, err := io.Copy(buffer, req.Body)
	if err != nil {
		return nil, err
	}

	subject := TwirpPathToNatsPath(req.URL.Path)
	// subject := strings.Trim(strings.Replace(req.URL.Path, "/", ".", -1), ".")
	log.With(zap.String("subject", subject)).Info("Sending to subject")

	msg := nats.Msg{
		Subject: subject,
		Data:    buffer.Bytes(),
		Header:  nats.Header(req.Header),
	}

	if err := svc.conn.PublishMsg(&msg); err != nil {
		return nil, errors.Wrap(err, "unable to send on nats")
	}

	res := http.Response{
		StatusCode: http.StatusOK,
	}
	ctype := req.Header.Get("content-type")

	if strings.Contains(ctype, "application/json") {
		res.Body = io.NopCloser(strings.NewReader("{}"))
	} else {
		res.Body = io.NopCloser(bytes.NewReader([]byte{}))
	}

	return &res, nil
}

type TopicHandler interface {
	GroupName() string
	Handler() corepb.TwirpServer
}

type Consumer struct {
	url      string
	conn     *nats.Conn
	Topic    string
	handlers []manager.MsgHandler
}

func (svc *Consumer) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	wg := sync.WaitGroup{}

	for _, item := range svc.handlers {
		wg.Add(1)
		go func(handler manager.MsgHandler) {
			log := log.With("group", handler.GroupName())
			log.Info("Creating queue subscription")
			_, err := svc.conn.QueueSubscribe(svc.Topic, handler.GroupName(), func(msg *nats.Msg) {
				parts := strings.Split(msg.Subject, ".")
				path := "/" + strings.Join(parts[0:len(parts)-1], ".") + "/" + parts[len(parts)-1]
				if parts[0] == "twirp" {
					path = strings.Replace(path, "/twirp.", "/twirp/", 1)
				}
				log.With(
					zap.String("subject", msg.Subject),
					zap.String("path", path),
				).Info("Got nats message")
				req := http.Request{
					URL: &url.URL{
						Scheme: "queue",
						Host:   "",
						Path:   path,
					},
					Method: "POST",
					Header: http.Header(msg.Header),
					Body:   io.NopCloser(bytes.NewReader(msg.Data)),
				}

				w := httptest.NewRecorder()
				handler.Handler().ServeHTTP(w, req.WithContext(ctx))

				res := w.Result()
				if res.StatusCode != http.StatusOK {
					buf, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
					log.With("statusCode", res.StatusCode).With("statusMsg", string(buf)).Info("nats consumer error")
				}
				msg.Ack()
			})

			if err != nil {
				// This is really not good -- TODO better handling
				wg.Add(-1)
				log.With(zap.Error(err), zap.String("addr", svc.url)).Fatal("unable to connect")
			}
		}(item)
	}

	wg.Wait()

	return nil
}

func (svc *client) TopicConsumer(ctx context.Context, topic string, handlers []manager.MsgHandler) manager.HandlerStart {
	log := logger.FromContext(ctx)

	log.With(zap.String("topic", topic)).Info("Consuming on topic")
	if err := svc.connect(ctx); err != nil {
		log.With(zap.Error(err)).Fatal("unable to connect")
	}

	consumer := Consumer{
		conn:     svc.conn,
		url:      svc.url,
		Topic:    topic,
		handlers: handlers,
	}

	return &consumer
}

func TwirpPathToNatsTopic(path string) string {
	return "twirp." + strings.Trim(strings.TrimPrefix(path, "/twirp"), "/") + ".*"
}

func TwirpPathToNatsPath(path string) string {
	return strings.Trim(strings.Replace(path, "/", ".", -1), ".")
}

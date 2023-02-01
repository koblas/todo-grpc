package main

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const EVENT_GROUP = "trigger-minio"
const MINIO_EVENT_BUS = "minioevents"

type Config struct {
	// Used by lambda
	NatsAddr string `environment:"NATS_ADDR"`
}

type handler struct {
	nats     *natsutil.Client
	producer corepbv1.FileEventbus
}

func newHandler(nats *natsutil.Client, producer corepbv1.FileEventbus) *handler {
	return &handler{nats: nats, producer: producer}
}

func (h *handler) Start(ctx context.Context) error {
	if err := h.nats.Connect(ctx); err != nil {
		return err
	}

	h.nats.Conn.QueueSubscribe(MINIO_EVENT_BUS, EVENT_GROUP, func(msg *nats.Msg) {
		log := logger.FromContext(ctx)

		var request events.S3Event
		if err := json.Unmarshal(msg.Data, &request); err != nil {
			log.With(zap.Error(err)).Error("Unable to unmarshal event")
			return
		}

		for _, record := range request.Records {
			if !strings.HasPrefix(record.EventName, "s3:ObjectCreated") {
				continue
			}
			bucket := record.S3.Bucket.Name
			// Minio leaves this as URL encoded
			key, err := url.QueryUnescape(record.S3.Object.Key)
			if err != nil {
				log.Error("bad URL encoding")
				continue
			}
			log = log.With(zap.String("bucket", bucket), zap.String("key", key))
			log.Info("translating event")

			parts := strings.SplitN(key, "/", 3)
			if len(parts) != 3 {
				log.Error("insufficent components")
				continue
			}
			log = log.With(zap.String("bucket", bucket), zap.String("key", key))
			log.Info("translating event")
			// got message
			h.producer.FileUploaded(ctx, &corepbv1.FileUploadEvent{
				Info: &corepbv1.FileUploadInfo{
					UserId:   &parts[1],
					FileType: parts[0],
					Url:      "s3://" + bucket + "/" + key,
				},
			})
		}
	})
	return nil
}

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config Config

	if err := confmgr.Parse(&config); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	nats := natsutil.NewNatsClient(config.NatsAddr)

	producer := corepbv1.NewFileEventbusProtobufClient(
		"",
		nats,
	)

	mgr.Start(newHandler(nats, producer))
}

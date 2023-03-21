package main

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/minio/minio-go/v7/pkg/notification"

	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/filestore"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const EVENT_GROUP = "trigger-minio"
const MINIO_EVENT_BUS = "minioevents"

type Config struct {
	NatsAddr      string `environment:"NATS_ADDR"`
	MinioEndpoint string `environment:"MINIO_ENDPOINT" default:"s3.amazonaws.com"`
	UploadBucket  string `environment:"UPLOAD_BUCKET"`
}

type handler struct {
	conf     Config
	nats     *natsutil.Client
	producer corev1connect.FileEventbusServiceClient
	store    *filestore.MinioProvider
}

func newHandler(conf Config, nats *natsutil.Client, producer corev1connect.FileEventbusServiceClient) *handler {
	return &handler{
		conf:     conf,
		nats:     nats,
		producer: producer,
		store:    filestore.NewMinioProvider(conf.MinioEndpoint),
	}
}

func (h *handler) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	// log.With(zap.Any("ENV", os.Environ())).Info("ENV")

	if err := h.nats.Connect(ctx); err != nil {
		log.With(zap.Error(err)).Fatal("Unable to connect to nats")
	}
	if err := h.store.BuildClient(ctx); err != nil {
		log.With(zap.Error(err)).Fatal("Unable to connect to minio")
	}

	topicArn := notification.NewArn("minio", "sqs", "", "PRIMARY", "nats")
	topicConfig := notification.NewConfig(topicArn)
	topicConfig.AddEvents(notification.ObjectCreatedAll)

	nconf := notification.Configuration{}
	nconf.AddQueue(topicConfig)

	log.With(zap.String("TOPIC", topicArn.String())).Info("Got Topic")

	if err := h.store.VerifyBucket(ctx, h.conf.UploadBucket); err != nil {
		log.With(zap.Error(err), zap.String("bucket", h.conf.UploadBucket)).Fatal("Unable to verify bucket")
	}

	if err := h.store.Client.SetBucketNotification(ctx, h.conf.UploadBucket, nconf); err != nil {
		log.With(zap.Error(err)).Fatal("Unable to create notification")
	}

	// initialization complete
	log.With(
		zap.String("bus", MINIO_EVENT_BUS),
		zap.String("group", EVENT_GROUP),
	).Info("Subscribing to NATS events")

	h.nats.Conn.QueueSubscribe(MINIO_EVENT_BUS, EVENT_GROUP, func(msg *nats.Msg) {
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
			h.producer.FileUploaded(ctx, connect.NewRequest(&corev1.FileServiceUploadEvent{
				Info: &corev1.FileServiceUploadInfo{
					UserId:   &parts[1],
					FileType: parts[0],
					Url:      "s3://" + bucket + "/" + key,
				},
			}))
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

	producer := corev1connect.NewFileEventbusServiceClient(
		nats,
		"",
	)

	mgr.Start(newHandler(config, nats, producer))
}

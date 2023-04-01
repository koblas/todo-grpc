package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambdaGo "github.com/aws/aws-lambda-go/lambda"
	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"go.uber.org/zap"
)

type Config struct {
	BusEntityArn string `validate:"required"`
}

type Handler struct {
	produer eventbusv1connect.FileEventbusServiceClient
}

func (state *Handler) Start(ctx context.Context) error {
	lambdaGo.StartWithContext(ctx, func(ctx context.Context, request events.S3Event) error {
		log := logger.FromContext(ctx)
		for _, record := range request.Records {
			bucket := record.S3.Bucket.Name
			key := record.S3.Object.Key
			log = log.With(zap.String("bucket", bucket), zap.String("key", key))
			log.Info("translating event")

			parts := strings.SplitN(key, "/", 3)
			if len(parts) != 3 {
				log.Error("insufficent components")
				continue
			}

			_, err := state.produer.FileUploaded(ctx, connect.NewRequest(&corev1.FileServiceUploadEvent{
				Info: &corev1.FileServiceUploadInfo{
					UserId:   &parts[1],
					FileType: parts[0],
					Url:      "s3://" + bucket + "/" + key,
				},
			}))
			if err != nil {
				log.With(zap.Error(err)).Info("failed to publish")
			}

		}
		return nil
	})
	return nil
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	var config Config

	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), aws.NewLoaderSsm(mgr.Context(), "/common/")); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := eventbusv1connect.NewFileEventbusServiceClient(
		awsutil.NewTwirpCallLambda(),
		config.BusEntityArn,
	)

	mgr.Start(&Handler{producer})
}

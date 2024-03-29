package aws

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/koblas/grpc-todo/pkg/eventbus"
	busutil "github.com/koblas/grpc-todo/pkg/eventbus/util"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

type awsBus struct {
	topic  string
	client *sns.Client
}

var _ eventbus.Producer = (*awsBus)(nil)

func NewAwsProducer(topic string) (eventbus.Producer, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	client := sns.NewFromConfig(cfg)

	return &awsBus{client: client, topic: topic}, nil
}

func (svc *awsBus) Enqueue(ctx context.Context, msg *eventbus.Message) error {
	log := logger.FromContext(ctx)

	input := busutil.MessageToSns(svc.topic, *msg)

	log.With("topicArn", svc.topic).With("bodyLen", len(*input.Message)).With("messageAttributes", input.MessageAttributes).Info("sending message")

	_, err := svc.client.Publish(ctx, &input)

	return err
}

type SqsConsumerFunc func(context.Context, *eventbus.Message) error

type SqsConsumer struct {
	handler SqsConsumerFunc
}

// This runs by insertion
func NewAwsSqsConsumer(handler SqsConsumerFunc) *SqsConsumer {
	// todo
	return &SqsConsumer{
		handler: handler,
	}
}

func (svc *SqsConsumer) AddMessages(ctx context.Context, messages []*sqstypes.Message) error {
	for _, item := range messages {
		mout, err := busutil.SqsToMessage(*item)
		if err != nil {
			return err
		}

		err = svc.handler(ctx, &mout)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *SqsConsumer) AddMessagesLambda(ctx context.Context, event events.SQSEvent) (events.SQSEventResponse, error) {
	failed := []events.SQSBatchItemFailure{}
	log := logger.FromContext(ctx)

	for _, item := range event.Records {
		msg := busutil.EventToSqs(item)

		mout, err := busutil.SqsToMessage(msg)
		if err != nil {
			failed = append(failed, events.SQSBatchItemFailure{ItemIdentifier: item.MessageId})

			log.With(zap.Error(err)).With(zap.String("messageId", item.MessageId)).Error("unmarshaling failed")

			continue
		}

		if err := svc.handler(ctx, &mout); err != nil {
			failed = append(failed, events.SQSBatchItemFailure{ItemIdentifier: item.MessageId})

			log.With(zap.Error(err)).With(zap.String("messageId", item.MessageId)).Error("unmarshaling failed")
		}
	}

	return events.SQSEventResponse{BatchItemFailures: failed}, nil
}

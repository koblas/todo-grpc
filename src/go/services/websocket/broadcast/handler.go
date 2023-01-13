package broadcast

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws"
	smithy "github.com/aws/smithy-go"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
)

type BroadcastServer struct {
	store  websocket.ConnectionStore
	client PostToConnectionAPI
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type PostToConnectionAPI interface {
	PostToConnection(ctx context.Context, params *apigatewaymanagementapi.PostToConnectionInput, optFns ...func(*apigatewaymanagementapi.Options)) (*apigatewaymanagementapi.PostToConnectionOutput, error)
}

type Option func(*BroadcastServer)

func WithStore(store websocket.ConnectionStore) Option {
	return func(conf *BroadcastServer) {
		conf.store = store
	}
}

func WithClient(client PostToConnectionAPI) Option {
	return func(conf *BroadcastServer) {
		conf.client = client
	}
}

// Convert websocket-broadcast events into per-connection events
func NewBroadcastServer(opts ...Option) corepb.BroadcastEventbus {
	svr := BroadcastServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *BroadcastServer) Send(ctx context.Context, event *corepb.BroadcastEvent) (*corepb.EventbusEmpty, error) {
	log := logger.FromContext(ctx)
	log.Info("received broadcast event")

	conns, err := svc.store.ForUser(ctx, event.Filter.UserId)
	if err != nil {
		return nil, err
	}

	for _, connection := range conns {
		log.With("connectionId", connection).Info("Sending to connection")
		_, err = svc.client.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(connection),
			Data:         event.Data,
		})
		if err != nil {
			var ae smithy.APIError

			if errors.As(err, &ae) {
				// TODO:  There has to be a more-Go way to do this
				//   however with a bit of errors.Is / errors.As cannot get the goneexception
				if ae.ErrorCode() == "GoneException" {
					log.With("connectionId", connection).Info("Connect is Gone - deleting")
					// Connection is no longer present it should be removed
					if err = svc.store.Delete(ctx, connection); err != nil {
						log.With("error", err).Info("Unable to delete connection")
					}
				} else {
					log.With("status", ae.ErrorCode()).With("message", ae.ErrorMessage()).Info("Unable to send")
				}
			} else {
				log.With("connectionId", connection).With("error", err).Info("Unable to send message")
			}
		}
	}

	log.With("count", len(conns)).Info("Found connections")

	return &corepb.EventbusEmpty{}, nil
}

package broadcast

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws"
	smithy "github.com/aws/smithy-go"
	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	websocketv1 "github.com/koblas/grpc-todo/gen/core/websocket/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"go.uber.org/zap"
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

type BroadcastServerHandler struct {
}

// Convert websocket-broadcast events into per-connection events
func NewBroadcastServer(opts ...Option) map[string]http.Handler {
	svr := BroadcastServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := eventbusv1connect.NewBroadcastEventbusServiceHandler(&svr)

	return map[string]http.Handler{"websocket.broadcast": api}
}

func (svc *BroadcastServer) Send(ctx context.Context, event *connect.Request[websocketv1.BroadcastEvent]) (*connect.Response[eventv1.BroadcastEventbusSendResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("userId", event.Msg.Filter.UserId))
	log.Info("received broadcast event")

	conns, err := svc.store.ForUser(ctx, event.Msg.Filter.UserId)
	if err != nil {
		log.With(zap.Error(err)).Info("Failed to lookup user")
		return nil, err
	}
	success := 0

	for _, connection := range conns {
		clog := log.With("connectionId", connection)
		clog.Info("Sending to connection")
		_, err = svc.client.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(connection),
			Data:         event.Msg.Data,
		})
		if err != nil {
			var ae smithy.APIError

			if errors.As(err, &ae) {
				// TODO:  There has to be a more-Go way to do this
				//   however with a bit of errors.Is / errors.As cannot get the goneexception
				if ae.ErrorCode() == "GoneException" {
					clog.Info("Connect is Gone - deleting")
					// Connection is no longer present it should be removed
					if err = svc.store.Delete(ctx, connection); err != nil {
						clog.With("error", err).Info("Unable to delete connection")
					}
				} else {
					clog.With("status", ae.ErrorCode()).With("message", ae.ErrorMessage()).Info("Unable to send")
				}
			} else {
				clog.With("error", err).Info("Unable to send message")
			}
		} else {
			success += 1
		}
	}

	if len(conns) == success {
		log.With("count", len(conns), "success", success).Info("send successful")
	} else {
		log.With("count", len(conns), "success", success).Info("send partial failure")
	}

	return connect.NewResponse(&eventv1.BroadcastEventbusSendResponse{}), nil
}

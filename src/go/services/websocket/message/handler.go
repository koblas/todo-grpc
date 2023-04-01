package message

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	messagev1 "github.com/koblas/grpc-todo/gen/api/message/v1"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	mcorev1 "github.com/koblas/grpc-todo/gen/core/message/v1"
	websocketv1 "github.com/koblas/grpc-todo/gen/core/websocket/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

type MessageServer struct {
	producer eventbusv1connect.BroadcastEventbusServiceClient
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*MessageServer)

func WithProducer(producer eventbusv1connect.BroadcastEventbusServiceClient) Option {
	return func(conf *MessageServer) {
		conf.producer = producer
	}
}

func NewMessageChangeServer(opts ...Option) map[string]http.Handler {
	svr := MessageServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := eventbusv1connect.NewMessageEventbusServiceHandler(&svr)
	return map[string]http.Handler{"message.change": api}
}

func (svc *MessageServer) Change(ctx context.Context, eventIn *connect.Request[mcorev1.MessageChangeEvent]) (*connect.Response[eventv1.MessageEventbusServiceChangeResponse], error) {
	event := eventIn.Msg
	log := logger.FromContext(ctx)
	log.Info("received message event")

	obj := event.Current
	action := "update"
	if event.Current == nil {
		obj = event.Original
		action = "delete"
	} else if event.Original == nil {
		action = "create"
	}

	if obj == nil {
		return nil, errors.New("missing object")
	}
	data, err := json.Marshal(SocketMessage{
		Topic:    "message",
		ObjectId: obj.Id,
		Action:   action,
		Body: messagev1.MessageItem{
			Id:     obj.Id,
			RoomId: obj.RoomId,
			Sender: obj.UserId,
			Text:   obj.Text,
		},
	})

	if err != nil {
		return nil, err
	}

	for _, userId := range event.UserId {
		if _, err := svc.producer.Send(ctx, connect.NewRequest(&websocketv1.BroadcastEvent{
			Filter: &websocketv1.BroadcastFilter{
				UserId: userId,
			},
			Data: data,
		})); err != nil {
			log.With(zap.Error(err)).Error("failed to send to websocket")
		}
	}

	return connect.NewResponse(&eventv1.MessageEventbusServiceChangeResponse{}), nil
}

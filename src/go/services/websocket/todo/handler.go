package todo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

type TodoServer struct {
	producer eventbusv1connect.BroadcastEventbusServiceClient
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*TodoServer)

func WithProducer(producer eventbusv1connect.BroadcastEventbusServiceClient) Option {
	return func(conf *TodoServer) {
		conf.producer = producer
	}
}

func NewTodoChangeServer(opts ...Option) map[string]http.Handler {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := eventbusv1connect.NewTodoEventbusServiceHandler(&svr)
	return map[string]http.Handler{"todo.change": api}
}

func (svc *TodoServer) TodoChange(ctx context.Context, eventIn *connect.Request[corev1.TodoChangeEvent]) (*connect.Response[eventv1.TodoEventbusTodoChangeResponse], error) {
	event := eventIn.Msg
	log := logger.FromContext(ctx)
	log.Info("received todo event")

	userId := ""
	if event.Current != nil {
		userId = event.Current.UserId
	} else if event.Original != nil {
		userId = event.Original.UserId
	} else {
		return nil, errors.New("no user found")
	}

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
		Topic:    "todo",
		ObjectId: obj.Id,
		Action:   action,
		Body: apiv1.TodoObject{
			Id:   obj.Id,
			Task: obj.Task,
		},
	})

	if err != nil {
		return nil, err
	}

	if _, err := svc.producer.Send(ctx, connect.NewRequest(&corev1.BroadcastEvent{
		Filter: &corev1.BroadcastFilter{
			UserId: userId,
		},
		Data: data,
	})); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return connect.NewResponse(&eventv1.TodoEventbusTodoChangeResponse{}), nil
}

package todo

import (
	"context"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/todo/v1"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	todov1 "github.com/koblas/grpc-todo/gen/core/todo/v1"
	websocketv1 "github.com/koblas/grpc-todo/gen/core/websocket/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"
)

type TodoServer struct {
	producer eventbusv1connect.BroadcastEventbusServiceClient
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

func (svc *TodoServer) TodoChange(ctx context.Context, eventIn *connect.Request[todov1.TodoChangeEvent]) (*connect.Response[eventv1.TodoEventbusTodoChangeResponse], error) {
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

	data, err := util.MarshalData("todo", obj.Id, action, &apiv1.TodoObject{
		Id:   obj.Id,
		Task: obj.Task,
	})

	if err != nil {
		return nil, err
	}

	if _, err := svc.producer.Send(ctx, connect.NewRequest(&websocketv1.BroadcastEvent{
		Filter: &websocketv1.BroadcastFilter{
			UserId: userId,
		},
		Data: data,
	})); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return connect.NewResponse(&eventv1.TodoEventbusTodoChangeResponse{}), nil
}

package todo

import (
	"context"
	"encoding/json"
	"errors"

	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"go.uber.org/zap"
)

type TodoServer struct {
	producer corepbv1.BroadcastEventbusService
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*TodoServer)

func WithProducer(producer corepbv1.BroadcastEventbusService) Option {
	return func(conf *TodoServer) {
		conf.producer = producer
	}
}

type TodoServerHandler struct {
	handler corepbv1.TwirpServer
}

func NewTodoChangeServer(opts ...Option) []manager.MsgHandler {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return []manager.MsgHandler{
		&TodoServerHandler{
			handler: corepbv1.NewTodoEventbusServiceServer(&svr),
		},
	}
}

func (svc *TodoServerHandler) GroupName() string {
	return "websocket.todo"
}

func (svc *TodoServerHandler) Handler() corepbv1.TwirpServer {
	return svc.handler
}

func (svc *TodoServer) TodoChange(ctx context.Context, event *corepbv1.TodoChangeEvent) (*corepbv1.TodoEventbusTodoChangeResponse, error) {
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
		Body: apipbv1.TodoObject{
			Id:   obj.Id,
			Task: obj.Task,
		},
	})

	if err != nil {
		return nil, err
	}

	if _, err := svc.producer.Send(ctx, &corepbv1.BroadcastEvent{
		Filter: &corepbv1.BroadcastFilter{
			UserId: userId,
		},
		Data: data,
	}); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return &corepbv1.TodoEventbusTodoChangeResponse{}, nil
}

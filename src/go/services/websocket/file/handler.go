package file

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	filev1 "github.com/koblas/grpc-todo/gen/core/file/v1"
	websocketv1 "github.com/koblas/grpc-todo/gen/core/websocket/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

type FileServer struct {
	producer eventbusv1connect.BroadcastEventbusServiceClient
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*FileServer)

func WithProducer(producer eventbusv1connect.BroadcastEventbusServiceClient) Option {
	return func(conf *FileServer) {
		conf.producer = producer
	}
}

func NewFileChangeServer(opts ...Option) map[string]http.Handler {
	svr := FileServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := eventbusv1connect.NewFileEventbusServiceHandler(&svr)
	return map[string]http.Handler{"websocket.file": api}
}

func (svc *FileServer) FileUploaded(ctx context.Context, eventIn *connect.Request[filev1.FileServiceUploadEvent]) (*connect.Response[eventv1.FileEventbusFileUploadedResponse], error) {
	return connect.NewResponse(&eventv1.FileEventbusFileUploadedResponse{}), nil
}

func (svc *FileServer) FileComplete(ctx context.Context, eventIn *connect.Request[filev1.FileServiceCompleteEvent]) (*connect.Response[eventv1.FileEventbusFileCompleteResponse], error) {
	event := eventIn.Msg
	log := logger.FromContext(ctx)
	log.Info("received file event")

	body := map[string]string{
		"id": event.Id,
	}
	msg := SocketMessage{
		Topic:    "file",
		ObjectId: event.Id,
		Action:   "create",
		Body:     body,
	}

	if event.ErrorMessage != nil {
		body["error"] = *event.ErrorMessage
		msg.Action = "error"
	}

	data, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	if _, err := svc.producer.Send(ctx, connect.NewRequest(&websocketv1.BroadcastEvent{
		Filter: &websocketv1.BroadcastFilter{
			UserId: *event.Info.UserId,
		},
		Data: data,
	})); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return connect.NewResponse(&eventv1.FileEventbusFileCompleteResponse{}), nil
}

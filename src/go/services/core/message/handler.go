package message

import (
	"context"

	"github.com/bufbuild/connect-go"
	messagev1 "github.com/koblas/grpc-todo/gen/core/message/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type MessageServer struct {
	store MessageStore
	// producer corev1.TodoEventbus
	pubsub corev1connect.MessageEventbusServiceClient
}

type Option func(*MessageServer)

func WithMessageStore(store MessageStore) Option {
	return func(cfg *MessageServer) {
		cfg.store = store
	}
}

func WithProducer(bus corev1connect.MessageEventbusServiceClient) Option {
	return func(cfg *MessageServer) {
		cfg.pubsub = bus
	}
}

func NewMessageServer(opts ...Option) *MessageServer {
	svr := MessageServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.store == nil {
		panic("Must provide a store")
	}

	return &svr
}

func (svc *MessageServer) Add(ctx context.Context, bufreq *connect.Request[messagev1.AddRequest]) (*connect.Response[messagev1.AddResponse], error) {
	msg := bufreq.Msg
	log := logger.FromContext(ctx).With(zap.String("method", "AddTodo"))
	log.Info("creating message")

	if msg.UserId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if msg.RoomId == "" {
		return nil, bufcutil.InvalidArgumentError("roomId", "missing")
	}
	if msg.Text == "" {
		return nil, bufcutil.InvalidArgumentError("task", "empty")
	}

	task, err := svc.store.Create(ctx, Message{
		ID:     xid.New().String(),
		RoomId: msg.RoomId,
		Text:   msg.Text,
		UserId: msg.UserId,
	})

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	message := messagev1.MessageItem{
		Id:     task.ID,
		Text:   task.Text,
		RoomId: task.RoomId,
		UserId: task.UserId,
	}

	users, err := svc.store.Members(ctx, msg.RoomId)
	if err != nil {
		return nil, bufcutil.InternalError(err, "Unable to get room members")
	}
	svc.store.Join(ctx, msg.RoomId, msg.UserId)

	if svc.pubsub != nil {
		if _, err := svc.pubsub.Change(ctx, connect.NewRequest(&messagev1.MessageChangeEvent{
			IdemponcyId: xid.New().String(),
			UserId:      users,
			Current:     &message,
		})); err != nil {
			log.With("error", err).Info("user entity publish failed")
		}
	}

	return connect.NewResponse(&messagev1.AddResponse{Message: &message}), nil
}

func (svc *MessageServer) List(ctx context.Context, find *connect.Request[messagev1.ListRequest]) (*connect.Response[messagev1.ListResponse], error) {
	out, err := svc.store.FindByRoom(ctx, find.Msg.RoomId)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	messages := []*messagev1.MessageItem{}
	for _, item := range out {
		messages = append(messages, &messagev1.MessageItem{
			Id:     item.ID,
			UserId: item.UserId,
			RoomId: item.RoomId,
			Text:   item.Text,
		})
	}

	return connect.NewResponse(&messagev1.ListResponse{Messages: messages}), nil
}

func (svc *MessageServer) Delete(ctx context.Context, bufreq *connect.Request[messagev1.DeleteRequest]) (*connect.Response[messagev1.DeleteResponse], error) {
	msg := bufreq.Msg
	log := logger.FromContext(ctx).With(zap.String("method", "DeleteMessage"))
	log.Info("delete todo event")

	if msg.RoomId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if msg.UserId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if msg.MsgId == "" {
		return nil, bufcutil.InvalidArgumentError("id", "empty")
	}

	message, err := svc.store.DeleteOne(ctx, msg.RoomId, msg.MsgId)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	if message != nil && svc.pubsub != nil {
		if _, err := svc.pubsub.Change(ctx, connect.NewRequest(&messagev1.MessageChangeEvent{
			IdemponcyId: xid.New().String(),
			Original: &messagev1.MessageItem{
				Id:     message.ID,
				RoomId: message.RoomId,
				Text:   message.Text,
				UserId: message.UserId,
			},
		})); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return connect.NewResponse(&messagev1.DeleteResponse{}), nil
}

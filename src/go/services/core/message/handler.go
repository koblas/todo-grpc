package message

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	messagev1 "github.com/koblas/grpc-todo/gen/core/message/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type MessageServer struct {
	store MessageStore
	// producer corev1.TodoEventbus
	pubsub eventbusv1connect.MessageEventbusServiceClient
}

type Option func(*MessageServer)

func WithMessageStore(store MessageStore) Option {
	return func(cfg *MessageServer) {
		cfg.store = store
	}
}

func WithProducer(bus eventbusv1connect.MessageEventbusServiceClient) Option {
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
	log := logger.FromContext(ctx).With(zap.String("method", "AddMessage"))
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

	task, err := svc.store.CreateMessage(ctx, msg.OrgId, msg.RoomId, Message{
		ID:     xid.New().String(),
		RoomId: msg.RoomId,
		Text:   msg.Text,
		UserId: msg.UserId,
	})

	if err != nil {
		log.With(zap.Error(err)).Error("CreateMessage failed")
		return nil, bufcutil.InternalError(err)
	}

	message := messagev1.MessageItem{
		Id:     task.ID,
		Text:   task.Text,
		RoomId: task.RoomId,
		UserId: task.UserId,
	}

	users, err := svc.store.Members(ctx, msg.OrgId, msg.RoomId)
	if err != nil {
		log.With(zap.Error(err)).Error("Members failed")
		return nil, bufcutil.InternalError(err, "Unable to get room members")
	}
	svc.store.Join(ctx, msg.RoomId, msg.OrgId, msg.UserId)

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
	msg := find.Msg
	out, err := svc.store.ListMessages(ctx, msg.OrgId, msg.RoomId)
	log := logger.FromContext(ctx).With(zap.String("method", "ListMessages"))
	log.Info("begin call")

	if err != nil {
		log.With(zap.Error(err)).Error("ListMessages failed")
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
	log.Info("begin call")

	if msg.RoomId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if msg.UserId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if msg.MsgId == "" {
		return nil, bufcutil.InvalidArgumentError("id", "empty")
	}

	message, err := svc.store.GetMessage(ctx, msg.OrgId, msg.RoomId, msg.MsgId)
	if err != nil {
		log.With(zap.Error(err)).Error("GetMessage failed")
		return nil, bufcutil.InternalError(err)
	}

	if err := svc.store.DeleteOne(ctx, msg.OrgId, msg.RoomId, msg.MsgId); err != nil {
		log.With(zap.Error(err)).Error("DeleteOne failed")
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

func (svc *MessageServer) RoomList(ctx context.Context, find *connect.Request[messagev1.RoomListRequest]) (*connect.Response[messagev1.RoomListResponse], error) {
	msg := find.Msg
	log := logger.FromContext(ctx).With(zap.String("method", "RoomList"))
	log.Info("Begin call")
	rooms, err := svc.store.ListRooms(ctx, msg.OrgId, nil)
	if err != nil {
		log.With(zap.Error(err)).Error("ListRooms failed")
		return nil, bufcutil.InternalError(err)
	}

	result := []*messagev1.RoomItem{}
	for _, item := range rooms {
		result = append(result, &messagev1.RoomItem{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	return connect.NewResponse(&messagev1.RoomListResponse{
		Rooms: result,
	}), nil
}

func (svc *MessageServer) RoomJoin(ctx context.Context, find *connect.Request[messagev1.RoomJoinRequest]) (*connect.Response[messagev1.RoomJoinResponse], error) {
	msg := find.Msg
	log := logger.FromContext(ctx).With(
		zap.String("method", "RoomJoin"),
		zap.String("roomName", msg.Name),
		zap.String("orgId", msg.OrgId),
	)
	log.Info("Calling Room Join")
	rooms, err := svc.store.ListRooms(ctx, msg.OrgId, nil)
	if err != nil {
		log.With(zap.Error(err)).Error("ListRooms failed")
		return nil, bufcutil.InternalError(err)
	}
	log.With(zap.Int("count", len(rooms))).Info("Got rooms")

	var room *Room
	for _, item := range rooms {
		if item.Name == msg.Name {
			room = item
			break
		}
	}

	// TODO Not found -- create for now
	if room == nil {
		log.Info("Room not found creating")
		room, err = svc.store.CreateRoom(ctx, msg.OrgId, msg.UserId, msg.Name)
		if err != nil {
			log.With(zap.Error(err)).Error("CreateRoom failed")
			return nil, bufcutil.InternalError(err)
		}
	}
	if err := svc.store.Join(ctx, msg.OrgId, room.ID, msg.UserId); err != nil {
		log.With(zap.Error(err)).Error("Join failed")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&messagev1.RoomJoinResponse{
		Room: &messagev1.RoomItem{
			Id:   room.ID,
			Name: room.Name,
		},
	}), nil
}

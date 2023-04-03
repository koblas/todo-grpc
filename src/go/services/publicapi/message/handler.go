package message

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/message/v1"
	apicon "github.com/koblas/grpc-todo/gen/api/message/v1/messagev1connect"
	corev1 "github.com/koblas/grpc-todo/gen/core/message/v1"
	"github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type MessageServer struct {
	client     messagev1connect.MessageServiceClient
	userHelper interceptors.UserIdFromContext
}

type Option func(*MessageServer)

func WithMessageService(client messagev1connect.MessageServiceClient) Option {
	return func(svr *MessageServer) {
		svr.client = client
	}
}

func WithGetUserId(helper interceptors.UserIdFromContext) Option {
	return func(svr *MessageServer) {
		svr.userHelper = helper
	}
}

var _ apicon.MessageServiceHandler = (*MessageServer)(nil)

func NewMessageServer(opts ...Option) *MessageServer {
	svr := MessageServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.client == nil {
		panic("No backend message service")
	}
	if svr.userHelper == nil {
		panic("no user helper provided")
	}

	return &svr
}

// SayHello generates response to a Ping request
func (svc *MessageServer) MsgCreate(ctx context.Context, bufreq *connect.Request[apiv1.MsgCreateRequest]) (*connect.Response[apiv1.MsgCreateResponse], error) {
	msg := bufreq.Msg
	log := logger.FromContext(ctx)
	log.Info("AddMessage BEGIN")
	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	resp, err := svc.client.Add(ctx, connect.NewRequest(&corev1.AddRequest{
		UserId: userId,
		RoomId: msg.RoomId,
		Text:   msg.Text,
	}))

	if err != nil {
		log.With(zap.Error(err)).Error("Unable to create message")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.MsgCreateResponse{
		Message: &apiv1.MessageItem{
			Id:     resp.Msg.Message.Id,
			Sender: resp.Msg.Message.UserId,
			RoomId: resp.Msg.Message.RoomId,
			Text:   resp.Msg.Message.Text,
		},
	}), nil
}

func (svc *MessageServer) MsgList(ctx context.Context, bufreq *connect.Request[apiv1.MsgListRequest]) (*connect.Response[apiv1.MsgListResponse], error) {
	msg := bufreq.Msg
	log := logger.FromContext(ctx)
	log.Info("List Message BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	out, err := svc.client.List(ctx, connect.NewRequest(&corev1.ListRequest{
		RoomId: msg.RoomId,
		UserId: userId,
	}))

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	messages := []*apiv1.MessageItem{}
	if out != nil {
		for _, item := range out.Msg.Messages {
			messages = append(messages, &apiv1.MessageItem{
				Id:     item.Id,
				RoomId: item.RoomId,
				Sender: item.UserId,
				Text:   item.Text,
			})
		}
	}

	return connect.NewResponse(&apiv1.MsgListResponse{Messages: messages}), nil
}

func (svc *MessageServer) RoomJoin(ctx context.Context, bufreq *connect.Request[apiv1.RoomJoinRequest]) (*connect.Response[apiv1.RoomJoinResponse], error) {
	msg := bufreq.Msg
	log := logger.FromContext(ctx)
	log.Info("Room Join BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	out, err := svc.client.RoomJoin(ctx, connect.NewRequest(&corev1.RoomJoinRequest{
		UserId: userId,
		Name:   msg.Name,
	}))
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.RoomJoinResponse{
		Room: &apiv1.RoomItem{
			Id:   out.Msg.Room.Id,
			Name: out.Msg.Room.Name,
		},
	}), nil
}

func (svc *MessageServer) RoomList(ctx context.Context, bufreq *connect.Request[apiv1.RoomListRequest]) (*connect.Response[apiv1.RoomListResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("Room List BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	out, err := svc.client.RoomList(ctx, connect.NewRequest(&corev1.RoomListRequest{
		UserId: userId,
	}))
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	rooms := []*apiv1.RoomItem{}
	for _, item := range out.Msg.Rooms {
		rooms = append(rooms, &apiv1.RoomItem{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return connect.NewResponse(&apiv1.RoomListResponse{
		Rooms: rooms,
	}), nil
}

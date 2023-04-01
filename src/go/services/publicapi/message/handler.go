package message

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/message/v1"
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
func (svc *MessageServer) Add(ctx context.Context, bufreq *connect.Request[apiv1.AddRequest]) (*connect.Response[apiv1.AddResponse], error) {
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

	return connect.NewResponse(&apiv1.AddResponse{
		Message: &apiv1.MessageItem{
			Id:     resp.Msg.Message.Id,
			Sender: resp.Msg.Message.UserId,
			RoomId: resp.Msg.Message.RoomId,
			Text:   resp.Msg.Message.Text,
		},
	}), nil
}

func (svc *MessageServer) List(ctx context.Context, bufreq *connect.Request[apiv1.ListRequest]) (*connect.Response[apiv1.ListResponse], error) {
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

	return connect.NewResponse(&apiv1.ListResponse{Messages: messages}), nil
}

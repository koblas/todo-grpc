package user

import (
	"context"
	"encoding/json"
	"errors"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"go.uber.org/zap"
)

type UserServer struct {
	producer corepbv1.BroadcastEventbusService
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*UserServer)

type UserServerHandler struct {
	handler corepbv1.TwirpServer
}

func WithProducer(producer corepbv1.BroadcastEventbusService) Option {
	return func(conf *UserServer) {
		conf.producer = producer
	}
}

func NewUserChangeServer(opts ...Option) []manager.MsgHandler {
	svr := UserServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return []manager.MsgHandler{
		&UserServerHandler{
			handler: corepbv1.NewUserEventbusServiceServer(&svr),
		},
	}
}

func (svc *UserServerHandler) GroupName() string {
	return "websocket.user"
}

func (svc *UserServerHandler) Handler() corepbv1.TwirpServer {
	return svc.handler
}

func (svc *UserServer) UserChange(ctx context.Context, event *corepbv1.UserChangeEvent) (*corepbv1.UserEventbusUserChangeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("received user event")

	userId := ""
	if event.Current != nil {
		userId = event.Current.Id
	} else if event.Original != nil {
		userId = event.Original.Id
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
		Topic:    "user",
		ObjectId: obj.Id,
		Action:   action,
		Body:     protoutil.UserCoreToApi(obj),
	})

	if err != nil {
		log.With(zap.Error(err)).Error("failed to marshal")
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

	return &corepbv1.UserEventbusUserChangeResponse{}, nil
}

func (*UserServer) SecurityPasswordChange(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityPasswordChangeResponse, error) {
	return &corepbv1.UserEventbusSecurityPasswordChangeResponse{}, nil
}
func (*UserServer) SecurityForgotRequest(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityForgotRequestResponse, error) {
	return &corepbv1.UserEventbusSecurityForgotRequestResponse{}, nil
}
func (*UserServer) SecurityRegisterToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityRegisterTokenResponse, error) {
	return &corepbv1.UserEventbusSecurityRegisterTokenResponse{}, nil
}
func (*UserServer) SecurityInviteToken(context.Context, *corepbv1.UserSecurityEvent) (*corepbv1.UserEventbusSecurityInviteTokenResponse, error) {
	return &corepbv1.UserEventbusSecurityInviteTokenResponse{}, nil
}

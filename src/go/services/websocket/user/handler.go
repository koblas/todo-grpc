package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"go.uber.org/zap"
)

type UserServer struct {
	producer corepb.BroadcastEventbus
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type Option func(*UserServer)

type UserServerHandler struct {
	handler corepb.TwirpServer
}

func WithProducer(producer corepb.BroadcastEventbus) Option {
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
			handler: corepb.NewUserEventbusServer(&svr),
		},
	}
}

func (svc *UserServerHandler) GroupName() string {
	return "websocket.user"
}

func (svc *UserServerHandler) Handler() corepb.TwirpServer {
	return svc.handler
}

func (svc *UserServer) UserChange(ctx context.Context, event *corepb.UserChangeEvent) (*corepb.EventbusEmpty, error) {
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

	if _, err := svc.producer.Send(ctx, &corepb.BroadcastEvent{
		Filter: &corepb.BroadcastFilter{
			UserId: userId,
		},
		Data: data,
	}); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return &corepb.EventbusEmpty{}, nil
}

func (*UserServer) SecurityPasswordChange(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*UserServer) SecurityForgotRequest(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*UserServer) SecurityRegisterToken(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
func (*UserServer) SecurityInviteToken(context.Context, *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}

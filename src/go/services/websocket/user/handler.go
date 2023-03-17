package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"go.uber.org/zap"
)

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type UserServer struct {
	producer corev1connect.BroadcastEventbusServiceClient
}

type Option func(*UserServer)

func WithProducer(producer corev1connect.BroadcastEventbusServiceClient) Option {
	return func(conf *UserServer) {
		conf.producer = producer
	}
}

func NewUserChangeServer(opts ...Option) []http.Handler {
	svr := UserServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := corev1connect.NewUserEventbusServiceHandler(&svr)

	return []http.Handler{api}
}

func (svc *UserServer) UserChange(ctx context.Context, eventIn *connect.Request[corev1.UserChangeEvent]) (*connect.Response[corev1.UserEventbusUserChangeResponse], error) {
	event := eventIn.Msg
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

	if _, err := svc.producer.Send(ctx, connect.NewRequest(&corev1.BroadcastEvent{
		Filter: &corev1.BroadcastFilter{
			UserId: userId,
		},
		Data: data,
	})); err != nil {
		log.With(zap.Error(err)).Error("failed to send to websocket")
	}

	return connect.NewResponse(&corev1.UserEventbusUserChangeResponse{}), nil
}

func (*UserServer) SecurityPasswordChange(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityPasswordChangeResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityPasswordChangeResponse{}), nil
}
func (*UserServer) SecurityForgotRequest(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityForgotRequestResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityForgotRequestResponse{}), nil
}
func (*UserServer) SecurityRegisterToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityRegisterTokenResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityRegisterTokenResponse{}), nil
}
func (*UserServer) SecurityInviteToken(context.Context, *connect.Request[corev1.UserSecurityEvent]) (*connect.Response[corev1.UserEventbusSecurityInviteTokenResponse], error) {
	return connect.NewResponse(&corev1.UserEventbusSecurityInviteTokenResponse{}), nil
}

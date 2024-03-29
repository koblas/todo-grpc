package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	websocketv1 "github.com/koblas/grpc-todo/gen/core/websocket/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"
)

type UserServer struct {
	producer eventbusv1connect.BroadcastEventbusServiceClient
}

type Option func(*UserServer)

func WithProducer(producer eventbusv1connect.BroadcastEventbusServiceClient) Option {
	return func(conf *UserServer) {
		conf.producer = producer
	}
}

func NewUserChangeServer(opts ...Option) map[string]http.Handler {
	svr := UserServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	_, api := eventbusv1connect.NewUserEventbusServiceHandler(&svr)

	return map[string]http.Handler{"websocket.user": api}
}

func (svc *UserServer) UserChange(ctx context.Context, eventIn *connect.Request[userv1.UserChangeEvent]) (*connect.Response[eventv1.UserEventbusUserChangeResponse], error) {
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

	data, err := util.MarshalData("user", obj.Id, action, protoutil.UserCoreToApi(obj))

	if err != nil {
		log.With(zap.Error(err)).Error("failed to marshal")
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

	return connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil
}

func (*UserServer) SecurityPasswordChange(context.Context, *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityPasswordChangeResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityPasswordChangeResponse{}), nil
}
func (*UserServer) SecurityForgotRequest(context.Context, *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityForgotRequestResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityForgotRequestResponse{}), nil
}
func (*UserServer) SecurityRegisterToken(context.Context, *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityRegisterTokenResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil
}
func (*UserServer) SecurityInviteToken(context.Context, *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityInviteTokenResponse], error) {
	return connect.NewResponse(&eventv1.UserEventbusSecurityInviteTokenResponse{}), nil
}

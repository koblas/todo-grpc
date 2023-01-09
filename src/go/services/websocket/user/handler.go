package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws"
	smithy "github.com/aws/smithy-go"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
)

type UserServer struct {
	store  websocket.ConnectionStore
	client PostToConnectionAPI
}

type SocketMessage struct {
	ObjectId string      `json:"object_id"`
	Action   string      `json:"action"`
	Topic    string      `json:"topic"`
	Body     interface{} `json:"body"`
}

type PostToConnectionAPI interface {
	PostToConnection(ctx context.Context, params *apigatewaymanagementapi.PostToConnectionInput, optFns ...func(*apigatewaymanagementapi.Options)) (*apigatewaymanagementapi.PostToConnectionOutput, error)
}

type Option func(*UserServer)

func WithStore(store websocket.ConnectionStore) Option {
	return func(conf *UserServer) {
		conf.store = store
	}
}

func WithClient(client PostToConnectionAPI) Option {
	return func(conf *UserServer) {
		conf.client = client
	}
}

func NewUserChangeServer(opts ...Option) corepb.UserEventbus {
	svr := UserServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
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
	conns, err := svc.store.ForUser(ctx, userId)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	for _, connection := range conns {
		log.With("connectionId", connection).Info("Sending to connection")
		_, err = svc.client.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(connection),
			Data:         data,
		})
		if err != nil {
			var ae smithy.APIError

			if errors.As(err, &ae) {
				// TODO:  There has to be a more-Go way to do this
				//   however with a bit of errors.Is / errors.As cannot get the goneexception
				if ae.ErrorCode() == "GoneException" {
					log.With("connectionId", connection).Info("Connect is Gone - deleting")
					// Connection is no longer present it should be removed
					if err = svc.store.Delete(ctx, connection); err != nil {
						log.With("error", err).Info("Unable to delete connection")
					}
				} else {
					log.With("status", ae.ErrorCode()).With("message", ae.ErrorMessage()).Info("Unable to send")
				}
			} else {
				log.With("connectionId", connection).With("error", err).Info("Unable to send message")
			}
		}
	}

	log.With("count", len(conns)).Info("Found connections")

	return &corepb.EventbusEmpty{}, nil
}

func (svc *UserServer) UserSecurity(ctx context.Context, event *corepb.UserSecurityEvent) (*corepb.EventbusEmpty, error) {
	return &corepb.EventbusEmpty{}, nil
}
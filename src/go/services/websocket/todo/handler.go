package todo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws"
	smithy "github.com/aws/smithy-go"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/store/websocket"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
)

type TodoServer struct {
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

type Option func(*TodoServer)

func WithStore(store websocket.ConnectionStore) Option {
	return func(conf *TodoServer) {
		conf.store = store
	}
}

func WithClient(client PostToConnectionAPI) Option {
	return func(conf *TodoServer) {
		conf.client = client
	}
}

func NewTodoChangeServer(opts ...Option) core.TodoEventbus {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) TodoChange(ctx context.Context, event *core.TodoChangeEvent) (*core.EventbusEmpty, error) {
	log := logger.FromContext(ctx)
	log.Info("received todo event")

	userId := ""
	if event.Current != nil {
		userId = event.Current.UserId
	} else if event.Original != nil {
		userId = event.Original.UserId
	} else {
		return nil, errors.New("no user found")
	}
	action := "create"
	if event.Current == nil {
		action = "delete"
	} else if event.Original != nil {
		action = "update"
	}

	conns, err := svc.store.ForUser(userId)
	if err != nil {
		return nil, err
	}

	var obj *core.TodoObject
	if event.Current != nil {
		obj = event.Current
	} else if event.Original != nil {
		obj = event.Original
	} else {
		return nil, errors.New("missing object")
	}
	data, err := json.Marshal(SocketMessage{
		Topic:    "todo",
		ObjectId: obj.Id,
		Action:   action,
		Body: publicapi.TodoObject{
			Id:   obj.Id,
			Task: obj.Task,
		},
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
					if err = svc.store.Delete(connection); err != nil {
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

	return &core.EventbusEmpty{}, nil
}

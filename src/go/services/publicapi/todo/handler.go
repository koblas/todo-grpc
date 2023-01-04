package todo

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	todos    core.TodoService
	jwtMaker tokenmanager.Maker
}

type Option func(*TodoServer)

func WithTodoService(client core.TodoService) Option {
	return func(svr *TodoServer) {
		svr.todos = client
	}
}

func NewTodoServer(config Config, opts ...Option) *TodoServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	svr := TodoServer{
		jwtMaker: maker,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) getUserId(ctx context.Context) (string, error) {
	headers, ok := ctx.Value(awsutil.HeaderCtxKey).(http.Header)
	if !ok {
		if ctx.Value(awsutil.HeaderCtxKey) != nil {
			log.Println("Headers are present")
		}
		return "", fmt.Errorf("headers not in context")
	}

	value := headers.Get("authorization")
	if value == "" {
		return "", fmt.Errorf("no authorization header")
	}
	parts := strings.Split(value, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("bad format")
	}

	payload, err := svc.jwtMaker.VerifyToken(parts[1])
	if err != nil {
		return "", err
	}
	if payload.UserId == "" {
		return "", fmt.Errorf("no user_id")
	}

	return payload.UserId, nil
}

// SayHello generates response to a Ping request
func (svc *TodoServer) AddTodo(ctx context.Context, newTodo *publicapi.TodoAddParams) (*publicapi.TodoObject, error) {
	log := logger.FromContext(ctx)
	log.Info("AddTodo BEGIN")
	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	task, err := svc.todos.AddTodo(ctx, &core.TodoAddParams{
		UserId: userId, // TODO
		Task:   newTodo.Task,
	})

	if err != nil {
		log.With(zap.Error(err)).Error("Unable to create todo")
		return nil, twirp.InternalErrorWith(err)
	}
	log.With("task", newTodo.Task).Info("Received new task")

	return &publicapi.TodoObject{
		Id:   task.Id,
		Task: task.Task,
	}, nil
}

func (svc *TodoServer) GetTodos(ctx context.Context, _ *publicapi.TodoGetParams) (*publicapi.TodoResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GetTodos BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	out, err := svc.todos.GetTodos(ctx, &core.TodoGetParams{
		UserId: userId,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*publicapi.TodoObject{}
	if out != nil {
		for _, item := range out.Todos {
			todos = append(todos, &publicapi.TodoObject{
				Id:   item.Id,
				Task: item.Task,
			})
		}
	}

	return &publicapi.TodoResponse{Todos: todos}, nil
}

func (svc *TodoServer) DeleteTodo(ctx context.Context, delTodo *publicapi.TodoDeleteParams) (*publicapi.TodoDeleteResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("DeleteTodo BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	_, err = svc.todos.DeleteTodo(ctx, &core.TodoDeleteParams{
		UserId: userId,
		Id:     delTodo.Id,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &publicapi.TodoDeleteResponse{Message: "success"}, nil
}

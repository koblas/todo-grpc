package todo

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	todos    corev1connect.TodoServiceClient
	jwtMaker tokenmanager.Maker
}

type Option func(*TodoServer)

func WithTodoService(client corev1connect.TodoServiceClient) Option {
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
	headers, ok := ctx.Value(manager.HttpHeaderCtxKey).(http.Header)
	if !ok {
		if ctx.Value(manager.HttpHeaderCtxKey) != nil {
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
func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *connect.Request[apiv1.TodoAddRequest]) (*connect.Response[apiv1.TodoAddResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("AddTodo BEGIN")
	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	task, err := svc.todos.TodoAdd(ctx, connect.NewRequest(&corev1.TodoAddRequest{
		UserId: userId, // TODO
		Task:   newTodo.Msg.Task,
	}))

	if err != nil {
		log.With(zap.Error(err)).Error("Unable to create todo")
		return nil, bufcutil.InternalError(err)
	}
	log.With("task", newTodo.Msg.Task).Info("Received new task")

	return connect.NewResponse(&apiv1.TodoAddResponse{
		Todo: &apiv1.TodoObject{
			Id:   task.Msg.Todo.Id,
			Task: task.Msg.Todo.Task,
		},
	}), nil
}

func (svc *TodoServer) TodoList(ctx context.Context, _ *connect.Request[apiv1.TodoListRequest]) (*connect.Response[apiv1.TodoListResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("GetTodos BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	out, err := svc.todos.TodoList(ctx, connect.NewRequest(&corev1.TodoListRequest{
		UserId: userId,
	}))

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todos := []*apiv1.TodoObject{}
	if out != nil {
		for _, item := range out.Msg.Todos {
			todos = append(todos, &apiv1.TodoObject{
				Id:   item.Id,
				Task: item.Task,
			})
		}
	}

	return connect.NewResponse(&apiv1.TodoListResponse{Todos: todos}), nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, delTodo *connect.Request[apiv1.TodoDeleteRequest]) (*connect.Response[apiv1.TodoDeleteResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("DeleteTodo BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	_, err = svc.todos.TodoDelete(ctx, connect.NewRequest(&corev1.TodoDeleteRequest{
		UserId: userId,
		Id:     delTodo.Msg.Id,
	}))

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.TodoDeleteResponse{Message: "success"}), nil
}

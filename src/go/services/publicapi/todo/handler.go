package todo

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	todos    corepbv1.TodoService
	jwtMaker tokenmanager.Maker
}

type Option func(*TodoServer)

func WithTodoService(client corepbv1.TodoService) Option {
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
func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *apipbv1.TodoAddRequest) (*apipbv1.TodoAddResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("AddTodo BEGIN")
	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	task, err := svc.todos.TodoAdd(ctx, &corepbv1.TodoAddRequest{
		UserId: userId, // TODO
		Task:   newTodo.Task,
	})

	if err != nil {
		log.With(zap.Error(err)).Error("Unable to create todo")
		return nil, twirp.InternalErrorWith(err)
	}
	log.With("task", newTodo.Task).Info("Received new task")

	return &apipbv1.TodoAddResponse{
		Todo: &apipbv1.TodoObject{
			Id:   task.Todo.Id,
			Task: task.Todo.Task,
		},
	}, nil
}

func (svc *TodoServer) TodoList(ctx context.Context, _ *apipbv1.TodoListRequest) (*apipbv1.TodoListResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("GetTodos BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	out, err := svc.todos.TodoList(ctx, &corepbv1.TodoListRequest{
		UserId: userId,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*apipbv1.TodoObject{}
	if out != nil {
		for _, item := range out.Todos {
			todos = append(todos, &apipbv1.TodoObject{
				Id:   item.Id,
				Task: item.Task,
			})
		}
	}

	return &apipbv1.TodoListResponse{Todos: todos}, nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, delTodo *apipbv1.TodoDeleteRequest) (*apipbv1.TodoDeleteResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("DeleteTodo BEGIN")

	userId, err := svc.getUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, twirp.Unauthenticated.Error("missing userid")
	}

	_, err = svc.todos.TodoDelete(ctx, &corepbv1.TodoDeleteRequest{
		UserId: userId,
		Id:     delTodo.Id,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &apipbv1.TodoDeleteResponse{Message: "success"}, nil
}

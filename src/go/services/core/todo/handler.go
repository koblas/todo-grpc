package todo

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	todov1 "github.com/koblas/grpc-todo/gen/core/todo/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type TodoServer struct {
	todos TodoStore
	// producer corev1.TodoEventbus
	pubsub eventbusv1connect.TodoEventbusServiceClient
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus eventbusv1connect.TodoEventbusServiceClient) Option {
	return func(cfg *TodoServer) {
		cfg.pubsub = bus
	}
}

func NewTodoServer(opts ...Option) *TodoServer {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.todos == nil {
		panic("Must provide a store")
	}

	return &svr
}

func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *connect.Request[todov1.TodoAddRequest]) (*connect.Response[todov1.TodoAddResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("method", "AddTodo"))
	log.Info("creating todo event")

	if newTodo.Msg.UserId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if newTodo.Msg.Task == "" {
		return nil, bufcutil.InvalidArgumentError("task", "empty")
	}

	task, err := svc.todos.Create(ctx, Todo{
		ID:     xid.New().String(),
		Task:   newTodo.Msg.Task,
		UserId: newTodo.Msg.UserId,
	})

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todo := todov1.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}

	if svc.pubsub != nil {
		if _, err := svc.pubsub.TodoChange(ctx, connect.NewRequest(&todov1.TodoChangeEvent{
			Current: &todo,
		})); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return connect.NewResponse(&todov1.TodoAddResponse{Todo: &todo}), nil
}

func (svc *TodoServer) TodoList(ctx context.Context, find *connect.Request[todov1.TodoListRequest]) (*connect.Response[todov1.TodoListResponse], error) {
	out, err := svc.todos.FindByUser(ctx, find.Msg.UserId)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todos := []*todov1.TodoObject{}
	for _, item := range out {
		todos = append(todos, &todov1.TodoObject{
			Id:   item.ID,
			Task: item.Task,
		})
	}

	return connect.NewResponse(&todov1.TodoListResponse{Todos: todos}), nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, params *connect.Request[todov1.TodoDeleteRequest]) (*connect.Response[todov1.TodoDeleteResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("method", "DeleteTodo"))
	log.Info("delete todo event")

	if params.Msg.UserId == "" {
		return nil, bufcutil.InvalidArgumentError("userId", "missing")
	}
	if params.Msg.Id == "" {
		return nil, bufcutil.InvalidArgumentError("id", "empty")
	}

	todo, err := svc.todos.DeleteOne(ctx, params.Msg.UserId, params.Msg.Id)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	if todo != nil && svc.pubsub != nil {
		if _, err := svc.pubsub.TodoChange(ctx, connect.NewRequest(&todov1.TodoChangeEvent{
			Original: &todov1.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		})); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return connect.NewResponse(&todov1.TodoDeleteResponse{Message: "ok"}), nil
}

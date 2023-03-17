package todo

import (
	"context"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type TodoServer struct {
	todos TodoStore
	// producer corev1.TodoEventbus
	pubsub corev1connect.TodoEventbusServiceClient
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus corev1connect.TodoEventbusServiceClient) Option {
	return func(cfg *TodoServer) {
		cfg.pubsub = bus
	}
}

func NewTodoServer(opts ...Option) *TodoServer {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *connect.Request[corev1.TodoAddRequest]) (*connect.Response[corev1.TodoAddResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("method", "AddTodo"))
	log.Info("creating todo event")

	task, err := svc.todos.Create(ctx, Todo{
		ID:     xid.New().String(),
		Task:   newTodo.Msg.Task,
		UserId: newTodo.Msg.UserId,
	})

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todo := corev1.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}

	if _, err := svc.pubsub.TodoChange(ctx, connect.NewRequest(&corev1.TodoChangeEvent{
		Current: &todo,
	})); err != nil {
		log.With("error", err).Info("todo entity publish failed")
	}

	return connect.NewResponse(&corev1.TodoAddResponse{Todo: &todo}), nil
}

func (svc *TodoServer) TodoList(ctx context.Context, find *connect.Request[corev1.TodoListRequest]) (*connect.Response[corev1.TodoListResponse], error) {
	out, err := svc.todos.FindByUser(ctx, find.Msg.UserId)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todos := []*corev1.TodoObject{}
	for _, item := range out {
		todos = append(todos, &corev1.TodoObject{
			Id:   item.ID,
			Task: item.Task,
		})
	}

	return connect.NewResponse(&corev1.TodoListResponse{Todos: todos}), nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, params *connect.Request[corev1.TodoDeleteRequest]) (*connect.Response[corev1.TodoDeleteResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("method", "DeleteTodo"))
	log.Info("delete todo event")

	todo, err := svc.todos.DeleteOne(ctx, params.Msg.UserId, params.Msg.Id)

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	if todo != nil {
		if _, err := svc.pubsub.TodoChange(ctx, connect.NewRequest(&corev1.TodoChangeEvent{
			Original: &corev1.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		})); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return connect.NewResponse(&corev1.TodoDeleteResponse{Message: "ok"}), nil
}

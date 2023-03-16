package todo

import (
	"context"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

type TodoServer struct {
	todos TodoStore
	// producer corepbv1.TodoEventbus
	pubsub corepbv1.TodoEventbusService
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus corepbv1.TodoEventbusService) Option {
	return func(cfg *TodoServer) {
		cfg.pubsub = bus
	}
}

func NewTodoServer(opts ...Option) corepbv1.TodoService {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *corepbv1.TodoAddRequest) (*corepbv1.TodoAddResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "AddTodo"))
	log.Info("creating todo event")

	task, err := svc.todos.Create(ctx, Todo{
		ID:     xid.New().String(),
		Task:   newTodo.Task,
		UserId: newTodo.UserId,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todo := corepbv1.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}

	if _, err := svc.pubsub.TodoChange(ctx, &corepbv1.TodoChangeEvent{
		Current: &todo,
	}); err != nil {
		log.With("error", err).Info("todo entity publish failed")
	}

	return &corepbv1.TodoAddResponse{Todo: &todo}, nil
}

func (svc *TodoServer) TodoList(ctx context.Context, find *corepbv1.TodoListRequest) (*corepbv1.TodoListResponse, error) {
	out, err := svc.todos.FindByUser(ctx, find.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*corepbv1.TodoObject{}
	for _, item := range out {
		todos = append(todos, &corepbv1.TodoObject{
			Id:   item.ID,
			Task: item.Task,
		})
	}

	return &corepbv1.TodoListResponse{Todos: todos}, nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, params *corepbv1.TodoDeleteRequest) (*corepbv1.TodoDeleteResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "DeleteTodo"))
	log.Info("delete todo event")

	todo, err := svc.todos.DeleteOne(ctx, params.UserId, params.Id)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if todo != nil {
		if _, err := svc.pubsub.TodoChange(ctx, &corepbv1.TodoChangeEvent{
			Original: &corepbv1.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		}); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return &corepbv1.TodoDeleteResponse{Message: "ok"}, nil
}

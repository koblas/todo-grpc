package todo

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/rs/xid"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	todos TodoStore
	// producer core.TodoEventbus
	pubsub core.TodoEventbus
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus core.TodoEventbus) Option {
	return func(cfg *TodoServer) {
		cfg.pubsub = bus
	}
}

func NewTodoServer(opts ...Option) core.TodoService {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) AddTodo(ctx context.Context, newTodo *core.TodoAddParams) (*core.TodoObject, error) {
	log := logger.FromContext(ctx)
	log.Info("creating todo event")

	task, err := svc.todos.Create(ctx, Todo{
		ID:     xid.New().String(),
		Task:   newTodo.Task,
		UserId: newTodo.UserId,
	})

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todo := core.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}

	if _, err := svc.pubsub.TodoChange(ctx, &core.TodoChangeEvent{
		Current: &todo,
	}); err != nil {
		log.With("error", err).Info("todo entity publish failed")
	}

	return &todo, nil
}

func (svc *TodoServer) GetTodos(ctx context.Context, find *core.TodoGetParams) (*core.TodoResponse, error) {
	out, err := svc.todos.FindByUser(ctx, find.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*core.TodoObject{}
	for _, item := range out {
		todos = append(todos, &core.TodoObject{
			Id:   item.ID,
			Task: item.Task,
		})
	}

	return &core.TodoResponse{Todos: todos}, nil
}

func (svc *TodoServer) DeleteTodo(ctx context.Context, params *core.TodoDeleteParams) (*core.TodoDeleteResponse, error) {
	log := logger.FromContext(ctx)
	log.Info("delete todo event")

	todo, err := svc.todos.DeleteOne(ctx, params.UserId, params.Id)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if todo != nil {
		if _, err := svc.pubsub.TodoChange(ctx, &core.TodoChangeEvent{
			Original: &core.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		}); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return &core.TodoDeleteResponse{Message: "ok"}, nil
}

package todo

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/rs/xid"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

type TodoServer struct {
	todos TodoStore
	// producer core.TodoEventbus
	producer eventbus.Producer
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus eventbus.Producer) Option {
	return func(cfg *TodoServer) {
		cfg.producer = bus
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

	task, err := svc.todos.Create(Todo{
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

	if err := eventbus.EnqueuePb(ctx, svc.producer, eventbus.ChangeMessage{
		ID:      todo.Id,
		Action:  "create",
		Current: &todo,
	}); err != nil {
		log.With(zap.Error(err)).Error("unable to send event")
	}

	return &todo, nil
}

func (svc *TodoServer) GetTodos(ctx context.Context, find *core.TodoGetParams) (*core.TodoResponse, error) {
	out, err := svc.todos.FindByUser(find.UserId)

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

	todo, err := svc.todos.DeleteOne(params.UserId, params.Id)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if todo != nil {
		if err := eventbus.EnqueuePb(ctx, svc.producer, eventbus.ChangeMessage{
			ID:     todo.ID,
			Action: "delete",
			Original: &core.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		}); err != nil {
			log.With(zap.Error(err)).Error("unable to send event")
		}
	}

	return &core.TodoDeleteResponse{Message: "ok"}, nil
}

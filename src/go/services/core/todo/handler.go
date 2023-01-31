package todo

import (
	"context"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
)

type TodoServer struct {
	todos TodoStore
	// producer corepb.TodoEventbus
	pubsub corepb.TodoEventbus
}

type Option func(*TodoServer)

func WithTodoStore(store TodoStore) Option {
	return func(cfg *TodoServer) {
		cfg.todos = store
	}
}

func WithProducer(bus corepb.TodoEventbus) Option {
	return func(cfg *TodoServer) {
		cfg.pubsub = bus
	}
}

func NewTodoServer(opts ...Option) corepb.TodoService {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *TodoServer) AddTodo(ctx context.Context, newTodo *corepb.TodoAddParams) (*corepb.TodoObject, error) {
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

	todo := corepb.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}

	if _, err := svc.pubsub.TodoChange(ctx, &corepb.TodoChangeEvent{
		Current: &todo,
	}); err != nil {
		log.With("error", err).Info("todo entity publish failed")
	}

	return &todo, nil
}

func (svc *TodoServer) GetTodos(ctx context.Context, find *corepb.TodoGetParams) (*corepb.TodoResponse, error) {
	out, err := svc.todos.FindByUser(ctx, find.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*corepb.TodoObject{}
	for _, item := range out {
		todos = append(todos, &corepb.TodoObject{
			Id:   item.ID,
			Task: item.Task,
		})
	}

	return &corepb.TodoResponse{Todos: todos}, nil
}

func (svc *TodoServer) DeleteTodo(ctx context.Context, params *corepb.TodoDeleteParams) (*corepb.TodoDeleteResponse, error) {
	log := logger.FromContext(ctx).With(zap.String("method", "DeleteTodo"))
	log.Info("delete todo event")

	todo, err := svc.todos.DeleteOne(ctx, params.UserId, params.Id)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if todo != nil {
		if _, err := svc.pubsub.TodoChange(ctx, &corepb.TodoChangeEvent{
			Original: &corepb.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		}); err != nil {
			log.With("error", err).Info("todo entity publish failed")
		}
	}

	return &corepb.TodoDeleteResponse{Message: "ok"}, nil
}

package todo

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/renstrom/shortuuid"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	todos    TodoStore
	producer core.TodoEventbus
}

func NewTodoServer(producer core.TodoEventbus, store TodoStore) core.TodoService {
	return &TodoServer{
		todos:    store,
		producer: producer,
	}
}

func (svc *TodoServer) AddTodo(ctx context.Context, newTodo *core.TodoAddParams) (*core.TodoObject, error) {
	log := logger.FromContext(ctx)
	log.Info("creating todo event")

	task, err := svc.todos.Create(Todo{
		ID:     shortuuid.New(),
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

	if _, err := svc.producer.Message(ctx, &core.TodoEvent{
		IdemponcyId: todo.Id,
		Action:      "create",
		Current:     &todo,
	}); err != nil {
		log.With("error", err).Error("Eventbus message failed")
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
		if _, err := svc.producer.Message(ctx, &core.TodoEvent{
			IdemponcyId: shortuuid.New(),
			Action:      "delete",
			Previous: &core.TodoObject{
				Id:     todo.ID,
				Task:   todo.Task,
				UserId: todo.UserId,
			},
		}); err != nil {
			log.With("error", err).Error("Eventbus message failed")
		}
	}

	return &core.TodoDeleteResponse{Message: "ok"}, nil
}

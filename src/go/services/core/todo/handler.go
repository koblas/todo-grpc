package todo

import (
	"context"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/renstrom/shortuuid"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	todos  TodoStore
	pubsub eventbus.Producer
}

func NewTodoServer(producer eventbus.Producer, store TodoStore) core.TodoService {
	return &TodoServer{
		todos:  store,
		pubsub: producer,
	}
}

func (svc *TodoServer) AddTodo(ctx context.Context, newTodo *core.TodoAddParams) (*core.TodoObject, error) {
	todo := Todo{
		ID:     shortuuid.New(),
		Task:   newTodo.Task,
		UserId: newTodo.UserId,
	}
	task, err := svc.todos.Create(todo)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &core.TodoObject{
		Id:     task.ID,
		Task:   task.Task,
		UserId: task.UserId,
	}, nil
}

func (svc *TodoServer) GetTodos(ctx context.Context, find *core.TodoGetParams) (*core.TodoResponse, error) {
	out, err := svc.todos.FindByUser(find.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	todos := []*core.TodoObject{}
	if out != nil {
		for _, item := range out {
			todos = append(todos, &core.TodoObject{
				Id:   item.ID,
				Task: item.Task,
			})
		}
	}

	return &core.TodoResponse{Todos: todos}, nil
}

func (svc *TodoServer) DeleteTodo(ctx context.Context, params *core.TodoDeleteParams) (*core.TodoDeleteResponse, error) {
	err := svc.todos.DeleteOne(params.UserId, params.Id)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &core.TodoDeleteResponse{Message: "ok"}, nil
}

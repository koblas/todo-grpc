package todo

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type memoryTodo struct {
	todos map[string][]*Todo
}

var _ TodoStore = (*memoryTodo)(nil)

func NewTodoMemoryStore() *memoryTodo {
	return &memoryTodo{
		todos: map[string][]*Todo{},
	}
}

func (store *memoryTodo) FindByUser(ctx context.Context, user_id string) ([]*Todo, error) {
	if todos, found := store.todos[user_id]; found {
		return todos, nil
	}

	return []*Todo{}, nil
}

func (store *memoryTodo) DeleteOne(ctx context.Context, user_id string, id string) (*Todo, error) {
	todos, found := store.todos[user_id]
	if !found {
		return nil, nil
	}

	filtered := []*Todo{}
	var matched *Todo
	for _, todo := range todos {
		if todo.ID == id {
			matched = todo
			continue
		}
		filtered = append(filtered, todo)
	}

	store.todos[user_id] = filtered

	return matched, nil
}

func (store *memoryTodo) Create(ctx context.Context, todo Todo) (*Todo, error) {
	todos, found := store.todos[todo.UserId]
	if !found {
		todos = []*Todo{}
	}

	todo.ID = ulid.Make().String()
	store.todos[todo.UserId] = append(todos, &todo)

	return &todo, nil
}

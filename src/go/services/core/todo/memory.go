package todo

import (
	"github.com/google/uuid"
)

type memoryTodo struct {
	todos map[string][]Todo
}

var _ TodoStore = (*memoryTodo)(nil)

func NewTodoMemoryStore() *memoryTodo {
	return &memoryTodo{
		todos: map[string][]Todo{},
	}
}

func (store *memoryTodo) FindByUser(user_id string) ([]Todo, error) {
	if todos, found := store.todos[user_id]; found {
		return todos, nil
	}

	return []Todo{}, nil
}

func (store *memoryTodo) DeleteOne(user_id string, id string) (*Todo, error) {
	todos, found := store.todos[user_id]
	if !found {
		return &Todo{}, nil
	}

	replace := []Todo{}
	matchIdx := -1
	for idx, todo := range todos {
		if todo.ID == id {
			matchIdx = idx
			continue
		}
		replace = append(replace, todo)
	}
	if matchIdx == -1 {
		return nil, nil
	}

	store.todos[user_id] = replace

	return &todos[matchIdx], nil
}

func (store *memoryTodo) Create(todo Todo) (*Todo, error) {
	todos, found := store.todos[todo.UserId]
	if !found {
		todos = []Todo{}
	}

	todo.ID = uuid.NewString()
	store.todos[todo.UserId] = append(todos, todo)

	return &todo, nil
}

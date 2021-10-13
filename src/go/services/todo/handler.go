package todo

import (
	"log"

	"github.com/koblas/grpc-todo/genpb"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	genpb.UnimplementedTodoServiceServer
	Todos []*genpb.TodoObject
}

// SayHello generates response to a Ping request
func (s *TodoServer) AddTodo(ctx context.Context, newTodo *genpb.AddTodoParams) (*genpb.TodoObject, error) {
	log.Printf("Received new task %s", newTodo.Task)
	todoObject := &genpb.TodoObject{
		Id:   uuid.NewV1().String(),
		Task: newTodo.Task,
	}
	s.Todos = append(s.Todos, todoObject)
	return todoObject, nil
}

func (s *TodoServer) GetTodos(ctx context.Context, _ *genpb.GetTodoParams) (*genpb.TodoResponse, error) {
	log.Printf("get tasks")
	return &genpb.TodoResponse{Todos: s.Todos}, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, delTodo *genpb.DeleteTodoParams) (*genpb.DeleteResponse, error) {
	var updatedTodos []*genpb.TodoObject
	for index, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos
	return &genpb.DeleteResponse{Message: "success"}, nil
}

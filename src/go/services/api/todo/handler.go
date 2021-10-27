package todo

import (
	"log"

	"github.com/koblas/grpc-todo/genpb/publicapi"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	publicapi.UnimplementedTodoServiceServer
	Todos map[string][]*publicapi.TodoObject
}

func NewTodoServer() *TodoServer {
	return &TodoServer{
		Todos: map[string][]*publicapi.TodoObject{},
	}
}

// SayHello generates response to a Ping request
func (s *TodoServer) AddTodo(ctx context.Context, newTodo *publicapi.AddTodoParams) (*publicapi.TodoObject, error) {
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		log.Panic("expected user_id")
	}

	todos, found := s.Todos[userId]
	if !found {
		todos = []*publicapi.TodoObject{}
	}

	log.Printf("Received new task %s", newTodo.Task)
	todoObject := &publicapi.TodoObject{
		Id:   uuid.NewV1().String(),
		Task: newTodo.Task,
	}

	todos = append(todos, todoObject)
	s.Todos[userId] = todos
	return todoObject, nil
}

func (s *TodoServer) GetTodos(ctx context.Context, _ *publicapi.GetTodoParams) (*publicapi.TodoResponse, error) {
	log.Printf("get tasks")

	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		log.Panic("expected userId")
	}

	todos, found := s.Todos[userId]
	if !found {
		todos = []*publicapi.TodoObject{}
	}

	return &publicapi.TodoResponse{Todos: todos}, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, delTodo *publicapi.DeleteTodoParams) (*publicapi.DeleteResponse, error) {
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		log.Panic("expected user_id")
	}

	todos, found := s.Todos[userId]
	if !found {
		todos = []*publicapi.TodoObject{}
	}

	var filtered []*publicapi.TodoObject
	for index, todo := range todos {
		if todo.Id == delTodo.Id {
			filtered = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	s.Todos[userId] = filtered

	return &publicapi.DeleteResponse{Message: "success"}, nil
}

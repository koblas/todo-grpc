package todo

import "context"

type Todo struct {
	ID     string `dynamodbav:"id"`
	UserId string `dynamodbav:"user_id"`
	Task   string `dynamodbav:"task"`
}

type TodoStore interface {
	FindByUser(ctx context.Context, user_id string) ([]Todo, error)
	DeleteOne(ctx context.Context, user_id string, id string) (*Todo, error)
	Create(ctx context.Context, todo Todo) (*Todo, error)
}

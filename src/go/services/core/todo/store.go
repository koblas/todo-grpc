package todo

type Todo struct {
	ID     string `dynamodbav:"id"`
	UserId string `dynamodbav:"user_id"`
	Task   string `dynamodbav:"task"`
}

type TodoStore interface {
	FindByUser(user_id string) ([]Todo, error)
	DeleteOne(user_id string, id string) error
	Create(todo Todo) (*Todo, error)
}

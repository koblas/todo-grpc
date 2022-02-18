package websocket

// type ConnectionValue struct {
// 	Pk       string     `dynamodbav:"pk"`
// 	Sk       string     `dynamodbav:"sk"`
// 	DeleteAt *time.Time `dynamodbav:"delete_at",nullempty`
// }

type ConnectionStore interface {
	Create(userId string, connectionId string) error
	Delete(connectionId string) error
	ForUser(userId string) ([]string, error)
	Heartbeat(connectionId string) error
}

package message

import "context"

type Message struct {
	ID     string `dynamodbav:"id"`
	RoomId string `dynamodbav:"room_id"`
	UserId string `dynamodbav:"user_id"`
	Text   string `dynamodbav:"text"`
}

type MessageStore interface {
	FindByRoom(ctx context.Context, roomId string) ([]*Message, error)
	DeleteOne(ctx context.Context, roomId string, id string) (*Message, error)
	Create(ctx context.Context, msg Message) (*Message, error)

	Members(ctx context.Context, roomId string) ([]string, error)
	Join(ctx context.Context, roomId string, userId string) error
	Leave(ctx context.Context, roomId string, userId string) error
}

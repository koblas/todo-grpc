package message

import "context"

type Message struct {
	ID     string
	OrgId  string
	RoomId string
	UserId string
	Text   string
}

type Room struct {
	ID    string
	OrgId string
	Name  string
}

type MessageStore interface {
	ListMessages(ctx context.Context, orgId, roomId string) ([]*Message, error)
	DeleteOne(ctx context.Context, orgId, roomId string, msgId string) error
	GetMessage(ctx context.Context, orgId, roomId string, msgId string) (*Message, error)
	CreateMessage(ctx context.Context, orgId, roomId string, msg Message) (*Message, error)

	// If the user is not passed this will list all rooms for the org
	ListRooms(ctx context.Context, orgId string, userId *string) ([]*Room, error)
	CreateRoom(ctx context.Context, orgId, userId string, name string) (*Room, error)
	Members(ctx context.Context, orgId, roomId string) ([]string, error)
	Join(ctx context.Context, orgId, roomId string, userId string) error
	Leave(ctx context.Context, orgId, roomId string, userId string) error
}

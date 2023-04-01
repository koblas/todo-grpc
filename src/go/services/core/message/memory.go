package message

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type memoryMessages struct {
	messages map[string][]*Message
	members  map[string]map[string]struct{}
}

var _ MessageStore = (*memoryMessages)(nil)

func NewMessageMemoryStore() *memoryMessages {
	return &memoryMessages{
		messages: map[string][]*Message{},
		members:  map[string]map[string]struct{}{},
	}
}

func (store *memoryMessages) Join(ctx context.Context, roomId string, userId string) error {
	members, found := store.members[roomId]
	if !found {
		members = map[string]struct{}{}
		store.members[roomId] = members
	}

	members[userId] = struct{}{}

	return nil
}

func (store *memoryMessages) Leave(ctx context.Context, roomId string, userId string) error {
	members, found := store.members[roomId]
	if !found {
		return nil
	}

	delete(members, userId)

	return nil
}

func (store *memoryMessages) Members(ctx context.Context, roomId string) ([]string, error) {
	members, found := store.members[roomId]
	if !found {
		return []string{}, nil
	}

	users := []string{}
	for k := range members {
		users = append(users, k)
	}

	return users, nil
}

func (store *memoryMessages) FindByRoom(ctx context.Context, roomId string) ([]*Message, error) {
	if todos, found := store.messages[roomId]; found {
		return todos, nil
	}

	return []*Message{}, nil
}

func (store *memoryMessages) DeleteOne(ctx context.Context, roomId string, id string) (*Message, error) {
	todos, found := store.messages[roomId]
	if !found {
		return nil, nil
	}

	filtered := []*Message{}
	var matched *Message
	for _, todo := range todos {
		if todo.ID == id {
			matched = todo
			continue
		}
		filtered = append(filtered, todo)
	}

	store.messages[roomId] = filtered

	return matched, nil
}

func (store *memoryMessages) Create(ctx context.Context, message Message) (*Message, error) {
	todos, found := store.messages[message.RoomId]
	if !found {
		todos = []*Message{}
	}

	message.ID = ulid.Make().String()
	store.messages[message.RoomId] = append(todos, &message)

	return &message, nil
}

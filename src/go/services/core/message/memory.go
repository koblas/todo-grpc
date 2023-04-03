package message

import (
	"context"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

type memoryRoom struct {
	id       string
	orgId    string
	name     string
	messages []*Message
	members  map[string]struct{}
}

func (room memoryRoom) unmarshal() *Room {
	return &Room{
		ID:    room.id,
		OrgId: room.orgId,
		Name:  room.name,
	}
}

type memoryData struct {
	rooms map[[2]string]*memoryRoom
}

var _ MessageStore = (*memoryData)(nil)

func NewMemoryStore() *memoryData {
	return &memoryData{
		rooms: map[[2]string]*memoryRoom{},
	}
}

func (store *memoryData) CreateRoom(ctx context.Context, orgId, userId string, name string) (*Room, error) {
	id := xid.New().String()

	// TODO -- for a given org cannot create a duplciate named room

	data := &memoryRoom{
		id:       id,
		orgId:    orgId,
		name:     name,
		messages: []*Message{},
		members:  map[string]struct{}{},
	}

	store.rooms[[2]string{orgId, id}] = data

	return data.unmarshal(), nil
}

func (store *memoryData) ListRooms(ctx context.Context, orgId string, userId *string) ([]*Room, error) {
	rooms := []*Room{}

	for _, room := range store.rooms {
		if room.orgId != orgId {
			continue
		}
		if userId != nil {
			_, found := room.members[*userId]
			if !found {
				continue
			}
		}
		rooms = append(rooms, room.unmarshal())
	}

	return rooms, nil
}

func (store *memoryData) fetchRoom(ctx context.Context, orgId, roomId string) (*memoryRoom, bool) {
	room, found := store.rooms[[2]string{orgId, roomId}]

	return room, found
}

func (store *memoryData) Join(ctx context.Context, orgId, roomId string, userId string) error {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return errors.New("room not found")
	}

	room.members[userId] = struct{}{}

	return nil
}

func (store *memoryData) Leave(ctx context.Context, orgId, roomId string, userId string) error {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return nil
	}

	delete(room.members, userId)

	return nil
}

func (store *memoryData) Members(ctx context.Context, orgId, roomId string) ([]string, error) {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return []string{}, nil
	}

	users := []string{}
	for k := range room.members {
		users = append(users, k)
	}

	return users, nil
}

func (store *memoryData) ListMessages(ctx context.Context, orgId, roomId string) ([]*Message, error) {
	if room, found := store.fetchRoom(ctx, orgId, roomId); found {
		return room.messages, nil
	}

	return []*Message{}, nil
}

func (store *memoryData) GetMessage(ctx context.Context, orgId, roomId string, msgId string) (*Message, error) {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return nil, nil
	}

	for _, item := range room.messages {
		if item.ID == msgId {
			return item, nil
		}
	}

	return nil, nil
}

func (store *memoryData) DeleteOne(ctx context.Context, orgId, roomId string, msgId string) error {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return nil
	}

	filtered := []*Message{}
	for _, item := range room.messages {
		if item.ID == msgId {
			continue
		}
		filtered = append(filtered, item)
	}

	room.messages = filtered

	return nil
}

func (store *memoryData) CreateMessage(ctx context.Context, orgId, roomId string, message Message) (*Message, error) {
	room, found := store.fetchRoom(ctx, orgId, roomId)
	if !found {
		return nil, errors.New("room does not exist")
	}

	message.ID = ulid.Make().String()
	room.messages = append(room.messages, &message)

	return &message, nil
}

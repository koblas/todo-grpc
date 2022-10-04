package websocket

type memoryStore struct {
	data map[string]string
}

func NewMemoryStore() *memoryStore {
	return &memoryStore{
		data: map[string]string{},
	}
}

func (store *memoryStore) Create(userId string, connectionId string) error {
	store.data[connectionId] = userId
	return nil
}

func (store *memoryStore) Delete(connectionId string) error {
	delete(store.data, connectionId)
	return nil
}

func (store *memoryStore) ForUser(userId string) ([]string, error) {
	conns := []string{}

	for cId, uId := range store.data {
		if uId == userId {
			conns = append(conns, cId)
		}
	}

	return conns, nil
}

func (store *memoryStore) Heartbeat(connectionId string) error {
	return nil
}

package websocket

type stubStore struct{}

func NewStubStore() *stubStore {
	return &stubStore{}
}

func (store *stubStore) Create(userId string, connectionId string) error {
	return nil
}

func (store *stubStore) Delete(connectionId string) error {
	return nil
}

func (store *stubStore) ForUser(userId string) ([]string, error) {
	return []string{}, nil
}

func (store *stubStore) Heartbeat(connectionId string) error {
	return nil
}

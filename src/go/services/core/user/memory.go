package user

type memoryStore struct {
	database []User
}

func NewUserMemoryStore() UserStore {
	return &memoryStore{
		database: []User{},
	}
}

func (store *memoryStore) GetById(id string) (*User, error) {
	for _, u := range store.database {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, nil
}

func (store *memoryStore) GetByEmail(email string) (*User, error) {
	for _, u := range store.database {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, nil
}

func (store *memoryStore) CreateUser(user *User) error {
	store.database = append(store.database, *user)

	return nil
}

func (store *memoryStore) UpdateUser(user *User) error {
	for idx, u := range store.database {
		if u.ID == user.ID {
			store.database[idx] = *user
			return nil
		}
	}

	return nil
}

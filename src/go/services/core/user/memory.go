package user

type memoryStore struct {
	database []User
	auth     map[string]UserAuth
}

func NewUserMemoryStore() UserStore {
	return &memoryStore{
		database: []User{},
		auth:     map[string]UserAuth{},
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

func (store *memoryStore) CreateUser(user User) error {
	store.database = append(store.database, user)

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

func (store *memoryStore) buildAuthKey(provider, provider_id string) string {
	return provider + "#" + provider_id
}

func (store *memoryStore) AuthGet(provider, provider_id string) (*UserAuth, error) {
	auth, found := store.auth[store.buildAuthKey(provider, provider_id)]

	if !found {
		return nil, nil
	}
	return &auth, nil
}

func (store *memoryStore) AuthUpsert(provider, provider_id string, auth UserAuth) error {
	store.auth[store.buildAuthKey(provider, provider_id)] = auth

	return nil
}

func (store *memoryStore) AuthDelete(provider, provider_id string, auth UserAuth) error {
	delete(store.auth, store.buildAuthKey(provider, provider_id))

	return nil
}

package user

import (
	"time"
)

type UserStatus string

const (
	UserStatus_REGISTERED = "registered"
	UserStatus_ACTIVE     = "active"
	UserStatus_INVITED    = "invited"
	UserStatus_DISABLED   = "disabled"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Status   UserStatus
	Password []byte
	Settings map[string]map[string]string

	VerificationToken   *[]byte
	VerificationExpires *time.Time
}

type UserStore interface {
	GetById(id string) *User
	GetByEmail(email string) *User
	CreateUser(user *User) error
	UpdateUser(user *User) error
}

type userStore struct {
	database []User
}

func NewUserMemoryStore() UserStore {
	return &userStore{
		database: []User{},
	}
}

func (store *userStore) GetById(id string) *User {
	for _, u := range store.database {
		if u.ID == id {
			return &u
		}
	}

	return nil
}

func (store *userStore) GetByEmail(email string) *User {
	for _, u := range store.database {
		if u.Email == email {
			return &u
		}
	}

	return nil
}

func (store *userStore) CreateUser(user *User) error {
	store.database = append(store.database, *user)

	return nil
}

func (store *userStore) UpdateUser(user *User) error {
	for idx, u := range store.database {
		if u.ID == user.ID {
			store.database[idx] = *user
			return nil
		}
	}

	return nil
}

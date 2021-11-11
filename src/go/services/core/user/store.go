package user

import (
	"log"
	"time"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Status   int
	Password []byte
	Settings map[string]map[string]string

	VerificationToken   string
	VerificationExpires time.Time
}

func (s *UserServer) getById(id string) *User {
	for _, u := range s.users {
		if u.ID == id {
			return &u
		}
	}

	return nil
}

func (s *UserServer) updateUser(user *User) error {
	for idx, u := range s.users {
		if u.ID == user.ID {
			s.users[idx] = *user
			return nil
		}
	}

	return nil
}

func (s *UserServer) getByEmail(email string) *User {
	for _, u := range s.users {
		log.Print("CHECKING ", email, u.Email)
		if u.Email == email {
			return &u
		}
	}

	return nil
}

package user

import (
	"log"
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

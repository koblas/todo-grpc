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
	ID       string                       `dynamodbav:"id"`
	Name     string                       `dynamodbav:"name"`
	Email    string                       `dynamodbav:"email"`
	Status   UserStatus                   `dynamodbav:"status"`
	Password []byte                       `dynamodbav:"password"`
	Settings map[string]map[string]string `dynamodbav:"settings"`

	VerificationToken   *[]byte    `dynamodbav:"verification_token,nullempty"`
	VerificationExpires *time.Time `dynamodbav:"verification_expires,nullempty"`
}

type UserStore interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
}

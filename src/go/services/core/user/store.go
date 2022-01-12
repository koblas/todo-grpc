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
	ID             string                       `dynamodbav:"id"`
	Name           string                       `dynamodbav:"name"`
	Email          string                       `dynamodbav:"email"`
	VerifiedEmails []string                     `dynamodbav:"verified_email"`
	Status         UserStatus                   `dynamodbav:"status"`
	Settings       map[string]map[string]string `dynamodbav:"settings"`

	// For email address confirmation
	EmailVerifyToken     []byte     `dynamodbav:"email_verify_token",nullempty`
	EmailVerifyExpiresAt *time.Time `dynamodbav:"email_verify_expires_at",nullempty`
}

type UserAuth struct {
	UserID    string     `dynamodbav:"user_id"`
	Password  []byte     `dynamodbav:"password"`
	ExpiresAt *time.Time `dynamodbav:"expires_at",nullempty`
}

type UserStore interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	CreateUser(user User) error
	UpdateUser(user *User) error

	AuthGet(provider, provider_id string) (*UserAuth, error)
	AuthUpsert(provider, provider_id string, auth UserAuth) error
	AuthDelete(provider, provider_id string, auth UserAuth) error
}

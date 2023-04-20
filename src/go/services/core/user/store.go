package user

import (
	"context"
	"errors"
	"time"
)

type UserStatus string

const (
	UserStatus_REGISTERED = "registered"
	UserStatus_ACTIVE     = "active"
	UserStatus_INVITED    = "invited"
	UserStatus_DISABLED   = "disabled"
)

var (
	ErrorTeamNotFound = errors.New("team not found")
	ErrorUserNotFound = errors.New("user not found")
)

type User struct {
	ID             string                       `dynamodbav:"id"`
	Name           string                       `dynamodbav:"name"`
	Email          string                       `dynamodbav:"email"`
	VerifiedEmails []string                     `dynamodbav:"verified_email"`
	Status         UserStatus                   `dynamodbav:"status"`
	Settings       map[string]map[string]string `dynamodbav:"settings"`
	AvatarUrl      *string                      `dynamodbav:"avatar_url,nullempty"`

	// For email address confirmation
	EmailVerifyToken     []byte     `dynamodbav:"email_verify_token,nullempty"`
	EmailVerifyExpiresAt *time.Time `dynamodbav:"email_verify_expires_at,nullempty"`
}

type UserAuth struct {
	UserID    string     `dynamodbav:"user_id"`
	Password  []byte     `dynamodbav:"password"`
	ExpiresAt *time.Time `dynamodbav:"expires_at,nullempty"`
}

type Team struct {
	TeamId string
	Name   string
}

type TeamUser struct {
	UserId string
	TeamId string
	Role   string
}

// User operations
type UserStore interface {
	// Get the user by their unique ID
	GetById(ctx context.Context, id string) (*User, error)

	// Get the user by their normailzed email address
	GetByEmail(ctx context.Context, email string) (*User, error)

	// Create a user given their basic user information,
	// secondary behavior: this will also add them to the user's default team
	CreateUser(ctx context.Context, user User) error

	// Update a user record
	UpdateUser(ctx context.Context, user *User) error
}

type OAuthStore interface {
	// Get a user by their provider + provider_id identifier
	AuthGet(ctx context.Context, provider, provider_id string) (*UserAuth, error)

	// Create a user association based on their provider + provider_id identifier
	AuthUpsert(ctx context.Context, provider, provider_id string, auth UserAuth) error

	// Remove a user association based on their provider + provider_id identifier
	AuthDelete(ctx context.Context, provider, provider_id string, auth UserAuth) error
}

type TeamStore interface {
	// Create the given team with one or more members
	TeamCreate(ctx context.Context, name string, tuser ...TeamUser) (*Team, error)

	// Get details about this team
	TeamGet(ctx context.Context, teamId string) (*Team, error)

	// Add a given User to the team
	// Note - all teamIDs must match
	TeamAddUsers(ctx context.Context, tuser ...TeamUser) error

	// Remove one or more users from the given team
	TeamListUsers(ctx context.Context, teamId string) ([]TeamUser, error)

	// Remove one or more users from the given team
	TeamDeleteUsers(ctx context.Context, teamId string, userIds ...string) error

	// List all of the teams that this user is a member of
	TeamList(ctx context.Context, userId string) ([]*Team, error)

	// Delete the given team
	TeamDelete(ctx context.Context, teamId string) error
}

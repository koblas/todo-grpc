package user

/*
** Some solid thoughts on user/team/membership modeling
**   https://blog.bullettrain.co/teams-should-be-an-mvp-feature/
**/

import (
	"context"
	"errors"
	"time"
)

// UserStatus enum
type UserStatus int

const (
	UserStatus_UNSET UserStatus = iota
	UserStatus_REGISTERED
	UserStatus_ACTIVE
	UserStatus_INVITED
)

func (s UserStatus) String() string {
	switch s {
	case UserStatus_REGISTERED:
		return "registered"
	case UserStatus_ACTIVE:
		return "active"
	case UserStatus_INVITED:
		return "invited"
	}
	return "unknown"
}

func UserStatusFromString(value string) UserStatus {
	switch value {
	case "registered":
		return UserStatus_REGISTERED
	case "active":
		return UserStatus_ACTIVE
	case "invited":
		return UserStatus_INVITED
	}
	panic("bad user status string")
}

// Closed Status

type ClosedStatus int

const (
	ClosedStatus_ACTIVE ClosedStatus = iota
	ClosedStatus_DISABLED
	ClosedStatus_DELETED
)

func (s ClosedStatus) String() string {
	switch s {
	case ClosedStatus_ACTIVE:
		return ""
	case ClosedStatus_DISABLED:
		return "disabled"
	case ClosedStatus_DELETED:
		return "deleted"
	}
	return "unknown"
}

func ClosedStatusFromString(value string) ClosedStatus {
	switch value {
	case "":
		return ClosedStatus_ACTIVE
	case "deleted":
		return ClosedStatus_DELETED
	case "disabled":
		return ClosedStatus_DISABLED
	}
	panic("bad user status string: " + value)
}

// TeamStatus enum
type TeamStatus int

const (
	TeamStatus_UNSET TeamStatus = iota
	TeamStatus_ACTIVE
	TeamStatus_INVITED
)

func (s TeamStatus) String() string {
	switch s {
	case TeamStatus_ACTIVE:
		return "active"
	case TeamStatus_INVITED:
		return "invited"
	}
	panic("bad team status")
}

func TeamStatusFromString(value string) TeamStatus {
	switch value {
	case "active":
		return TeamStatus_ACTIVE
	case "invited":
		return TeamStatus_INVITED
	}
	panic("bad team status string")

}

// Standard errors

var (
	ErrorTeamNotFound = errors.New("team not found")
	ErrorUserNotFound = errors.New("user not found")
	ErrorAuthNotFound = errors.New("auth not found")
)

type User struct {
	ID             string
	Name           string
	Email          string
	VerifiedEmails []string
	Status         UserStatus
	ClosedStatus   ClosedStatus
	Settings       map[string]map[string]string
	AvatarUrl      *string

	// For email address confirmation
	// The Nonce is a one time secret that is used
	EmailVerifyNonce []byte
	// When presented the Nonce + SECRET will result in this TOKEN
	EmailVerifyToken     []byte
	EmailVerifyExpiresAt *time.Time
}

type UserAuth struct {
	UserID    string
	Password  []byte
	ExpiresAt *time.Time
}

type Team struct {
	TeamId string
	Name   string
}

type TeamMember struct {
	MemberId  string
	UserId    string
	TeamId    string
	Status    TeamStatus
	Role      string
	InvitedBy *string
	InvitedOn *time.Time
}

// User operations
type UserStore interface {
	// Get the user by their unique ID
	GetById(ctx context.Context, id string) (*User, error)

	// Get the user by their normailzed email address
	GetByEmail(ctx context.Context, email string) (*User, error)

	// Create a user given their basic user information,
	// secondary behavior: this will also add them to the user's default team
	CreateUser(ctx context.Context, user User) (*User, error)

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
	TeamCreate(ctx context.Context, name string, tuser ...TeamMember) (*Team, error)

	// Get details about this team
	TeamGet(ctx context.Context, teamId string) (*Team, error)

	//  Delete the given team
	TeamDelete(ctx context.Context, teamId string) error

	// Team membership operations

	// List all of the teams that this user is a member of
	TeamList(ctx context.Context, userId string) ([]TeamMember, error)

	// Get a specific team member (e.g. checking permissions)
	TeamGetMember(ctx context.Context, teamId string, userId string) (*TeamMember, error)

	// Add users to the membership, in the given status
	//  typically it will be INVITED but for the first user in the team
	//  they're added directly
	// MemberId will be assigned once the record is added -- you must set
	//  Role, Status and UserId in a TeamMember
	TeamAddMembers(ctx context.Context, tuser ...TeamMember) error

	// Accept an invite for a given team (it must exist)
	TeamAcceptInvite(ctx context.Context, teamId, userId string) error

	// Remove one or more users from the given team
	TeamListMembers(ctx context.Context, teamId string) ([]TeamMember, error)

	// Remove one or more users from the given team
	TeamDeleteMembers(ctx context.Context, teamId string, userIds ...string) error
}

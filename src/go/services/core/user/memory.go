package user

import (
	"context"
	"strings"

	"github.com/rs/xid"
)

type memoryTeamInfo struct {
	role   string
	status TeamStatus
}

type memoryTeam struct {
	id    string
	name  string
	users map[string]*memoryTeamInfo
}

type memoryStore struct {
	database map[string]User
	auth     map[string]UserAuth
	team     map[string]*memoryTeam
}

var _ UserStore = (*memoryStore)(nil)
var _ OAuthStore = (*memoryStore)(nil)
var _ TeamStore = (*memoryStore)(nil)

func NewUserMemoryStore() *memoryStore {
	return &memoryStore{
		database: map[string]User{},
		auth:     map[string]UserAuth{},
		team:     map[string]*memoryTeam{},
	}
}

func (store *memoryStore) GetById(ctx context.Context, userId string) (*User, error) {
	u, found := store.database[userId]
	if !found {
		return nil, ErrorUserNotFound
	}
	return &u, nil
}

func (store *memoryStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	for _, u := range store.database {
		if strings.EqualFold(u.Email, email) {
			return &u, nil
		}
	}

	return nil, ErrorUserNotFound
}

func (store *memoryStore) CreateUser(ctx context.Context, user User) (*User, error) {
	user.ID = xid.New().String()
	store.database[user.ID] = user
	store.TeamCreate(ctx, "", TeamMember{UserId: user.ID, Role: "admin"})

	return &user, nil
}

func (store *memoryStore) UpdateUser(ctx context.Context, user *User) error {
	store.database[user.ID] = *user

	return nil
}

//
// OAuth handling
//

func (store *memoryStore) buildAuthKey(provider, provider_id string) string {
	return provider + "#" + provider_id
}

func (store *memoryStore) AuthGet(ctx context.Context, provider, provider_id string) (*UserAuth, error) {
	auth, found := store.auth[store.buildAuthKey(provider, provider_id)]

	if !found {
		return nil, nil
	}
	return &auth, nil
}

func (store *memoryStore) AuthUpsert(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	store.auth[store.buildAuthKey(provider, provider_id)] = auth

	return nil
}

func (store *memoryStore) AuthDelete(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	delete(store.auth, store.buildAuthKey(provider, provider_id))

	return nil
}

// Team handling
func (team *memoryTeam) unmarshal() *Team {
	return &Team{
		TeamId: team.id,
		Name:   team.name,
	}
}

func (store *memoryStore) TeamCreate(ctx context.Context, name string, tuser ...TeamMember) (*Team, error) {
	teamId := xid.New().String()

	users := map[string]*memoryTeamInfo{}
	for _, item := range tuser {
		users[item.UserId] = &memoryTeamInfo{
			role:   item.Role,
			status: item.Status,
		}
	}

	store.team[teamId] = &memoryTeam{
		id:    teamId,
		name:  name,
		users: users,
	}

	return nil, nil
}

func (store *memoryStore) TeamGet(ctx context.Context, teamId string) (*Team, error) {
	team, found := store.team[teamId]
	if !found {
		return nil, ErrorTeamNotFound
	}

	return (team).unmarshal(), nil
}

func (store *memoryStore) TeamAddMembers(ctx context.Context, tuser ...TeamMember) error {
	if len(tuser) == 0 {
		return nil
	}
	team, found := store.team[tuser[0].TeamId]
	if !found {
		return ErrorTeamNotFound
	}
	for _, item := range tuser {
		team.users[item.UserId] = &memoryTeamInfo{
			role:   item.Role,
			status: item.Status,
		}
	}
	return nil
}

func (store *memoryStore) TeamAcceptInvite(ctx context.Context, teamId, userId string) error {
	team, found := store.team[teamId]
	if !found {
		return ErrorTeamNotFound
	}
	member, found := team.users[userId]
	if !found {
		return ErrorUserNotFound
	}
	member.status = TeamStatus_ACTIVE

	return nil
}

func (store *memoryStore) TeamDeleteMembers(ctx context.Context, teamId string, userIds ...string) error {
	team, found := store.team[teamId]
	if !found {
		return ErrorTeamNotFound
	}
	for _, userId := range userIds {
		delete(team.users, userId)
	}
	return nil
}

func (store *memoryStore) TeamGetMember(ctx context.Context, teamId, userId string) (*TeamMember, error) {
	team, found := store.team[teamId]
	if !found {
		return nil, ErrorTeamNotFound
	}

	for userId, data := range team.users {
		if _, found := store.database[userId]; !found {
			// This is bad
			return nil, ErrorUserNotFound
		}
		return &TeamMember{
			TeamId: teamId,
			UserId: userId,
			Role:   data.role,
			Status: data.status,
		}, nil
	}

	return nil, ErrorUserNotFound
}

func (store *memoryStore) TeamListMembers(ctx context.Context, teamId string) ([]TeamMember, error) {
	team, found := store.team[teamId]
	if !found {
		return nil, ErrorTeamNotFound
	}
	users := []TeamMember{}
	for userId, data := range team.users {
		if _, found := store.database[userId]; !found {
			// This is bad
			return nil, ErrorUserNotFound
		}
		users = append(users, TeamMember{
			TeamId: teamId,
			UserId: userId,
			Role:   data.role,
			Status: data.status,
		})
	}
	return users, nil
}

func (store *memoryStore) TeamList(ctx context.Context, userId string) ([]TeamMember, error) {
	members := []TeamMember{}

	for _, team := range store.team {
		if member, found := team.users[userId]; found {
			members = append(members, TeamMember{
				MemberId: team.id + "#" + userId,
				UserId:   userId,
				TeamId:   team.id,
				Status:   member.status,
				Role:     member.role,
			})
		}
	}

	return members, nil
}

func (store *memoryStore) TeamDelete(ctx context.Context, teamId string) error {
	delete(store.team, teamId)
	return nil
}

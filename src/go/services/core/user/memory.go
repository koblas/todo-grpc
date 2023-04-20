package user

import (
	"context"

	"github.com/rs/xid"
)

type memoryTeam struct {
	id    string
	name  string
	users map[string]string
}

type memoryStore struct {
	database map[string]User
	auth     map[string]UserAuth
	team     map[string]memoryTeam
}

var _ UserStore = (*memoryStore)(nil)
var _ OAuthStore = (*memoryStore)(nil)
var _ TeamStore = (*memoryStore)(nil)

func NewUserMemoryStore() *memoryStore {
	return &memoryStore{
		database: map[string]User{},
		auth:     map[string]UserAuth{},
		team:     map[string]memoryTeam{},
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
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, ErrorUserNotFound
}

func (store *memoryStore) CreateUser(ctx context.Context, user User) error {
	store.database[user.ID] = user
	store.TeamCreate(ctx, "", TeamUser{UserId: user.ID, Role: "admin"})

	return nil
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

func (store *memoryStore) TeamCreate(ctx context.Context, name string, tuser ...TeamUser) (*Team, error) {
	teamId := xid.New().String()

	users := map[string]string{}
	for _, item := range tuser {
		users[item.UserId] = item.Role
	}

	store.team[teamId] = memoryTeam{
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

	return (&team).unmarshal(), nil
}

func (store *memoryStore) TeamAddUsers(ctx context.Context, tuser ...TeamUser) error {
	if len(tuser) == 0 {
		return nil
	}
	team, found := store.team[tuser[0].TeamId]
	if !found {
		return ErrorTeamNotFound
	}
	for _, item := range tuser {
		team.users[item.UserId] = item.Role
	}
	return nil
}

func (store *memoryStore) TeamDeleteUsers(ctx context.Context, teamId string, userIds ...string) error {
	team, found := store.team[teamId]
	if !found {
		return ErrorTeamNotFound
	}
	for _, userId := range userIds {
		delete(team.users, userId)
	}
	return nil
}

func (store *memoryStore) TeamListUsers(ctx context.Context, teamId string) ([]TeamUser, error) {
	team, found := store.team[teamId]
	if !found {
		return nil, ErrorTeamNotFound
	}
	users := []TeamUser{}
	for userId, role := range team.users {
		if _, found := store.database[userId]; !found {
			// This is bad
			return nil, ErrorUserNotFound
		}
		users = append(users, TeamUser{
			TeamId: teamId,
			UserId: userId,
			Role:   role,
		})
	}
	return users, nil
}

func (store *memoryStore) TeamList(ctx context.Context, userId string) ([]*Team, error) {
	teams := []*Team{}

	for _, team := range store.team {
		if _, found := team.users[userId]; found {
			teams = append(teams, (&team).unmarshal())
		}
	}

	return teams, nil
}

func (store *memoryStore) TeamDelete(ctx context.Context, teamId string) error {
	delete(store.team, teamId)
	return nil
}

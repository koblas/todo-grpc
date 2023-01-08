package user

import "github.com/koblas/grpc-todo/gen/corepb"

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

func (s *UserServer) toProtoUser(user *User) *corepb.User {
	if user == nil {
		return nil
	}

	isVerified := false
	for _, v := range user.VerifiedEmails {
		isVerified = isVerified || (v == user.Email)
	}

	return &corepb.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Status:          statusToPbStatus[user.Status],
		EmailIsVerified: isVerified,
		AvatarUrl:       user.AvatarUrl,
	}
}

func (s *UserServer) toProtoSettings(user *User) *corepb.UserSettings {
	if user == nil {
		return nil
	}

	output := map[string]*corepb.UserSettingGroup{}
	for key, value := range user.Settings {
		subgroup := corepb.UserSettingGroup{}
		output[key] = &subgroup
		for subkey, subvalue := range value {
			subgroup.Values[subkey] = subvalue
		}
	}

	return &corepb.UserSettings{
		UserId:   user.ID,
		Settings: output,
	}
}

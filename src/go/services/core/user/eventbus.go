package user

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

func (s *UserServer) toProtoUser(user *User) *corepbv1.User {
	if user == nil {
		return nil
	}

	isVerified := false
	for _, v := range user.VerifiedEmails {
		isVerified = isVerified || (v == user.Email)
	}

	return &corepbv1.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Status:          statusToPbStatus[user.Status],
		EmailIsVerified: isVerified,
		AvatarUrl:       user.AvatarUrl,
	}
}

func (s *UserServer) toProtoSettings(user *User) *corepbv1.UserSettings {
	if user == nil {
		return nil
	}

	output := map[string]*corepbv1.UserSettingGroup{}
	for key, value := range user.Settings {
		subgroup := corepbv1.UserSettingGroup{}
		output[key] = &subgroup
		for subkey, subvalue := range value {
			subgroup.Values[subkey] = subvalue
		}
	}

	return &corepbv1.UserSettings{
		UserId:   user.ID,
		Settings: output,
	}
}

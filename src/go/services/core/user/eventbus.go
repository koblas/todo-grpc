package user

import (
	genpb "github.com/koblas/grpc-todo/twpb/core"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

func (s *UserServer) toProtoUser(user *User) *genpb.User {
	if user == nil {
		return nil
	}

	isVerified := false
	for _, v := range user.VerifiedEmails {
		isVerified = isVerified || (v == user.Email)
	}

	return &genpb.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Status:          statusToPbStatus[user.Status],
		EmailIsVerified: isVerified,
		AvatarUrl:       user.AvatarUrl,
	}
}

func (s *UserServer) toProtoSettings(user *User) *genpb.UserSettings {
	if user == nil {
		return nil
	}

	output := map[string]*genpb.UserSettingGroup{}
	for key, value := range user.Settings {
		subgroup := genpb.UserSettingGroup{}
		output[key] = &subgroup
		for subkey, subvalue := range value {
			subgroup.Values[subkey] = subvalue
		}
	}

	return &genpb.UserSettings{
		UserId:   user.ID,
		Settings: output,
	}
}

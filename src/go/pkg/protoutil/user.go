package protoutil

import (
	"net/url"

	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
)

func UserCoreToApi(user *userv1.User) *apiv1.User {
	if user == nil {
		return nil
	}

	avatarUrl := user.AvatarUrl
	if avatarUrl != nil {
		if u, err := url.Parse(*user.AvatarUrl); err == nil {
			if u.Scheme == "s3" {
				path := "https://files.iqvine.com" + u.Path
				avatarUrl = &path
			} else if u.Scheme == "minio" {
				path := "http://minio:9000" + u.Path
				avatarUrl = &path
			}
		}
	}

	return &apiv1.User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: avatarUrl,
	}
}

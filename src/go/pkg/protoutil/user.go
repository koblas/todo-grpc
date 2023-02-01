package protoutil

import (
	"net/url"

	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

func UserCoreToApi(user *corepbv1.User) *apipbv1.User {
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

	return &apipbv1.User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: avatarUrl,
	}
}

package protoutil

import (
	"net/url"

	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
)

func UserCoreToApi(user *corepb.User) *apipb.User {
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

	return &apipb.User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: avatarUrl,
	}
}

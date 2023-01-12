package protoutil

import (
	"strings"

	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
)

func UserCoreToApi(user *corepb.User) *apipb.User {
	if user == nil {
		return nil
	}

	avatarUrl := user.AvatarUrl
	if avatarUrl != nil && strings.HasPrefix(*user.AvatarUrl, "corefile:") {
		url := "/api/v1/fileput/" + strings.TrimPrefix(*user.AvatarUrl, "corefile:")
		avatarUrl = &url
	}

	return &apipb.User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: avatarUrl,
	}
}

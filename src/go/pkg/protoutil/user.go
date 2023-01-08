package protoutil

import (
	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
)

func UserCoreToApi(user *corepb.User) *apipb.User {
	if user == nil {
		return nil
	}

	return &apipb.User{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: user.AvatarUrl,
	}
}

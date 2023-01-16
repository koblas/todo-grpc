package tokenmanager

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/pkg/manager"
)

type VerifyJwt interface {
	VerifyToken(string) (*Payload, error)
}

func UserIdFromContext(ctx context.Context, verify VerifyJwt) (string, error) {
	headers, ok := ctx.Value(manager.HttpHeaderCtxKey).(http.Header)
	if !ok {
		if ctx.Value(manager.HttpHeaderCtxKey) != nil {
			log.Println("Headers are present")
		}
		return "", fmt.Errorf("headers not in context")
	}

	value := headers.Get("authorization")
	if value == "" {
		return "", fmt.Errorf("no authorization header")
	}
	parts := strings.Split(value, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("bad format")
	}

	payload, err := verify.VerifyToken(parts[1])
	if err != nil {
		return "", err
	}
	if payload.UserId == "" {
		return "", fmt.Errorf("no user_id")
	}

	return payload.UserId, nil
}

package user

import (
	"context"
	"log"
	"reflect"

	"github.com/koblas/grpc-todo/pkg/eventbus"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

// Common method to create wire version
func (s *OauthUserServer) publishMsg(ctx context.Context, stream string, action string, body []byte) error {
	attr := map[string]string{
		"stream":       stream,
		"action":       action,
		"content-type": "application/protobuf",
	}
	return s.pubsub.Enqueue(ctx, &eventbus.Message{
		Stream:     stream,
		Attributes: attr,
		BodyBytes:  body,
	})
}

// Can't wait for generics...  func buildAction[T any](orig, current *T) string {}
func buildAction(orig, current interface{}) string {
	origNil := orig == nil || (reflect.ValueOf(orig).Kind() == reflect.Ptr && reflect.ValueOf(orig).IsNil())
	currentNil := current == nil || (reflect.ValueOf(current).Kind() == reflect.Ptr && reflect.ValueOf(current).IsNil())

	log.Printf("buildAction origNil=%v currentNil=%v", origNil, currentNil)

	if origNil && !currentNil {
		return "created"
	}
	if !origNil && currentNil {
		return "deleted"
	}

	return "updated"
}

package user

import (
	"log"
	"reflect"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/robinjoseph08/redisqueue"
	"google.golang.org/protobuf/proto"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

// Common method to create wire version
func (s *OauthUserServer) publishMsg(stream string, action string, body []byte) error {
	values := []*genpb.MetadataEntry{
		{Key: "stream", Value: stream},
		{Key: "action", Value: action},
	}
	mbytes, err := proto.Marshal(&genpb.Metadata{
		Metadata: values,
	})
	if err != nil {
		return err
	}
	return s.pubsub.Enqueue(&redisqueue.Message{
		Stream: stream,
		Values: map[string]interface{}{
			"metadata": mbytes,
			"body":     body,
		},
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

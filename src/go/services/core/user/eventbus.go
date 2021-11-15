package user

import (
	"errors"
	"log"
	"reflect"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/robinjoseph08/redisqueue"
	"google.golang.org/protobuf/proto"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"
const ENTITY_SECURITY = "event:user_security"

func (s *UserServer) toProtoUser(user *User) *genpb.User {
	if user == nil {
		return nil
	}

	return &genpb.User{
		Id:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: statusToPbStatus[user.Status],
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

// Common method to create wire version
func (s *UserServer) publishMsg(stream string, action string, body []byte) error {
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

func (s *UserServer) publishUser(stream string, orig, current *User) error {
	action := buildAction(orig, current)
	pbchange := genpb.UserChangeEvent{
		Current: s.toProtoUser(current),
		Orig:    s.toProtoUser(orig),
	}
	body, err := proto.Marshal(&pbchange)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("body is nil")
	}

	return s.publishMsg(stream, action, body)
}

func (s *UserServer) publishSettings(stream string, orig, current *User) error {
	action := buildAction(orig, current)
	change := genpb.UserSettingsChangeEvent{
		Current: s.toProtoSettings(current),
		Orig:    s.toProtoSettings(orig),
	}
	body, err := proto.Marshal(&change)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("body is nil")
	}

	return s.publishMsg(stream, action, body)
}

func (s *UserServer) publishSecurity(log logger.Logger, action genpb.UserSecurity, user User, token string) error {
	log.With("userSecurity", action).Info("Sending security event")
	svalue, err := s.kms.Encode([]byte(token))
	if err != nil {
		return err
	}
	stoken := genpb.SecureValue{
		KeyUri:  svalue.KmsUri,
		DataKey: string(svalue.DataKey),
		Value:   string(svalue.Data),
	}
	event := genpb.UserSecurityEvent{
		Action: action,
		User:   s.toProtoUser(&user),
		Token:  &stoken,
	}
	body, err := proto.Marshal(&event)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("body is nil")
	}

	return s.publishMsg(ENTITY_SECURITY, genpb.UserSecurity_name[int32(action)], body)
}

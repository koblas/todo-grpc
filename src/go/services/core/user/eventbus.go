package user

import (
	"errors"
	"log"
	"reflect"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
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
func (s *UserServer) publishMsg(ctx context.Context, stream string, action string, body []byte) error {
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

func (s *UserServer) publishUser(ctx context.Context, stream string, orig, current *User) error {
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

	return s.publishMsg(ctx, stream, action, body)
}

func (s *UserServer) publishSettings(ctx context.Context, stream string, orig, current *User) error {
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

	return s.publishMsg(ctx, stream, action, body)
}

func (s *UserServer) publishSecurity(ctx context.Context, log logger.Logger, action genpb.UserSecurity, user User, token string) error {
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

	return s.publishMsg(ctx, ENTITY_SECURITY, genpb.UserSecurity_name[int32(action)], body)
}

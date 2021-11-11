package user

import (
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/renstrom/shortuuid"
	"github.com/robinjoseph08/redisqueue"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const ENTITY_USER = "entity:user"
const ENTITY_SETTINGS = "entity:user_settings"

// Server represents the gRPC server
type UserServer struct {
	genpb.UnimplementedUserServiceServer

	logger logger.Logger
	users  []User
	pubsub *redisqueue.Producer
}

func NewUserServer(log logger.Logger) *UserServer {
	pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
		},
	})
	if err != nil {
		log.With(err).Fatal("unable to start producer")
	}
	return &UserServer{
		logger: log,
		users:  []User{},
		pubsub: pubsub,
	}
}

func (s *UserServer) FindBy(ctx context.Context, params *genpb.FindParam) (*genpb.User, error) {
	s.logger.With("email", params.Email).Info("Received find")

	var user *User
	if params.Email != "" {
		user = s.getByEmail(params.Email)
	} else if params.UserId != "" {
		user = s.getById(params.Email)
	}

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "User not found")
	}

	s.logger.With("ID", user.ID).Info("found")

	return s.toProtoUser(user), nil
}

func (s *UserServer) Create(ctx context.Context, params *genpb.CreateParam) (*genpb.User, error) {
	s.logger.With("Email", params.Email).Info("Received Create")

	if s.getByEmail(params.Email) != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Email address not found")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		ID:       uuid.New().String(),
		Name:     params.Name,
		Email:    params.Email,
		Password: pass,
		Status:   int(genpb.UserStatus_ACTIVE),
		Settings: map[string]map[string]string{},

		VerificationToken:   shortuuid.New(),
		VerificationExpires: time.Now().Add(time.Duration(24 * time.Hour)),
	}

	s.users = append(s.users, user)

	if err := s.publishUser(ENTITY_USER, nil, &user); err != nil {
		s.logger.With(err).Info("Publish failed")
	}

	return s.toProtoUser(&user), nil
}

func (s *UserServer) Update(ctx context.Context, params *genpb.UpdateParam) (*genpb.User, error) {
	s.logger.With("UserId", params.UserId).Info("User Update")

	orig := s.getById(params.UserId)

	if orig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "User not found")
	}

	updated := *orig
	if len(params.Email) == 1 && params.Email[0] != "" {
		updated.Email = params.Email[0]
	}
	if len(params.Name) == 1 && params.Name[0] != "" {
		updated.Name = params.Name[0]
	}
	if len(params.Status) == 1 {
		updated.Status = int(params.Status[0])
	}
	if len(params.Password) == 1 && params.Password[0] != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(params.Password[0]), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updated.Password = pass
	}

	s.updateUser(&updated)
	s.publishUser(ENTITY_USER, orig, &updated)

	return s.toProtoUser(&updated), nil
}

func (s *UserServer) ComparePassword(ctx context.Context, params *genpb.AuthenticateParam) (*genpb.User, error) {
	s.logger.With("UserId", params.UserId).Info("Compare Passwords")

	user := s.getById(params.UserId)

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID address not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Password mismatch")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) GetSettings(ctx context.Context, params *genpb.UserIdParam) (*genpb.UserSettings, error) {
	user := s.getById(params.UserId)

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID address not found")
	}

	return s.toProtoSettings(user), nil
}

func (s *UserServer) SetSettings(ctx context.Context, params *genpb.UserSettingsUpdate) (*genpb.UserSettings, error) {
	orig := s.getById(params.UserId)

	if orig == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID address not found")
	}

	updated := *orig

	for key, value := range params.Settings {
		if value == nil {
			continue
		}
		if _, found := updated.Settings[key]; !found {
			updated.Settings[key] = map[string]string{}
		}

		for subkey, subvalue := range value.Values {
			updated.Settings[key][subkey] = subvalue
		}
	}

	s.updateUser(&updated)
	s.publishSettings(ENTITY_SETTINGS, orig, &updated)

	return s.toProtoSettings(&updated), nil
}

func (s *UserServer) toProtoUser(user *User) *genpb.User {
	if user == nil {
		return nil
	}

	return &genpb.User{
		Id:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		Status:            genpb.UserStatus(user.Status),
		VerificationToken: user.VerificationToken,
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
	change := genpb.UserChangeEvent{
		Current: s.toProtoUser(current),
		Orig:    s.toProtoUser(orig),
	}
	body, err := proto.Marshal(&change)
	if err != nil {
		return err
	}
	if body == nil {
		return errors.New("Body is nil")
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
		return errors.New("Body is nil")
	}

	return s.publishMsg(stream, action, body)
}

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

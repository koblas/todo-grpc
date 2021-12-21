package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"github.com/renstrom/shortuuid"
	"github.com/twitchtv/twirp"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type UserServer struct {
	users  UserStore
	pubsub eventbus.Producer
	kms    key_manager.Encoder
}

var pbStatusToStatus = map[genpb.UserStatus]UserStatus{
	genpb.UserStatus_ACTIVE:     UserStatus_ACTIVE,
	genpb.UserStatus_INVITED:    UserStatus_INVITED,
	genpb.UserStatus_DISABLED:   UserStatus_DISABLED,
	genpb.UserStatus_REGISTERED: UserStatus_REGISTERED,
}
var statusToPbStatus = map[UserStatus]genpb.UserStatus{
	UserStatus_ACTIVE:     genpb.UserStatus_ACTIVE,
	UserStatus_INVITED:    genpb.UserStatus_INVITED,
	UserStatus_DISABLED:   genpb.UserStatus_DISABLED,
	UserStatus_REGISTERED: genpb.UserStatus_REGISTERED,
}

func NewUserServer(producer eventbus.Producer, store UserStore) *UserServer {
	return &UserServer{
		users:  store,
		pubsub: producer,
		kms:    key_manager.NewSecureClear(),
	}
}

func (s *UserServer) FindBy(ctx context.Context, params *genpb.FindParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("email", params.Email).With("userId", params.UserId)
	log.Info("FindBy")

	var user *User
	var err error
	if params.Email != "" {
		user, err = s.users.GetByEmail(params.Email)
	} else if params.UserId != "" {
		user, err = s.users.GetById(params.UserId)
	}
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("User not found")
	}

	log.With("ID", user.ID).Info("found")

	return s.toProtoUser(user), nil
}

func (s *UserServer) Create(ctx context.Context, params *genpb.CreateParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Received Create")

	if params.Status != genpb.UserStatus_REGISTERED && params.Status != genpb.UserStatus_INVITED {
		log.Error("Bad user status")
		return nil, fmt.Errorf("invalid user status = %s", params.Status)
	}

	if u, err := s.users.GetByEmail(params.Email); err != nil {
		return nil, twirp.InternalErrorWith(err)
	} else if u != nil {
		return nil, twirp.AlreadyExists.Error("Email address not found")
	}

	pass := []byte{}
	if params.Password != "" {
		var errmsg string

		pass, errmsg = passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, twirp.InvalidArgument.Error(errmsg)
		}
	}

	var vExpires time.Time
	vToken := []byte{}
	var secret string
	if params.Status != genpb.UserStatus_ACTIVE {
		var err error
		vExpires = time.Now().Add(time.Duration(24 * time.Hour))
		vToken, secret, err = tokenEncrypt(shortuuid.New())
		if err != nil {
			return nil, err
		}
	}

	user := User{
		ID:       uuid.New().String(),
		Name:     params.Name,
		Email:    params.Email,
		Password: pass,
		Status:   pbStatusToStatus[params.Status],
		Settings: map[string]map[string]string{},

		VerificationToken:   &vToken,
		VerificationExpires: &vExpires,
	}

	if err := s.users.CreateUser(&user); err != nil {
		return nil, err
	}

	log.With("userId", user.ID, "email", user.Email).Info("User Created")

	if err := s.publishUser(ctx, ENTITY_USER, nil, &user); err != nil {
		log.With("error", err).Info("user entity publish failed")
	}
	if secret != "" {
		if params.Status == genpb.UserStatus_REGISTERED {
			if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_REGISTER_TOKEN, user, secret); err != nil {
				log.With("error", err).Info("register user publish failed")
			}
		} else if params.Status == genpb.UserStatus_INVITED {
			if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_INVITE_TOKEN, user, secret); err != nil {
				log.With("error", err).Info("invite user publish failed")
			}
		}
	}

	return s.toProtoUser(&user), nil
}

func (s *UserServer) Update(ctx context.Context, params *genpb.UpdateParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("User Update")

	orig, err := s.users.GetById(params.UserId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if orig == nil {
		return nil, twirp.NotFound.Error("User not found")
	}
	// Some basic validation
	if len(params.PasswordNew) != 0 || len(params.Password) != 0 {
		if len(params.PasswordNew) == 1 {
			// If you're setting a new password, you must provide the old
			if len(params.Password) == 0 {
				return nil, twirp.InvalidArgument.Error("Password missing")
			}
			if errmsg := validatePassword(params.PasswordNew[0]); errmsg != "" {
				return nil, twirp.InvalidArgument.Error("Password too short")
			}
		}
		// If you provided a password, we will check it...
		if len(params.Password) != 0 && !passwordCompare(orig, params.Password[0]) {
			return nil, twirp.InvalidArgument.Error("Password mismatch")
		}
	}

	// Now do the updates
	updated := *orig
	if len(params.Email) == 1 && params.Email[0] != "" {
		updated.Email = params.Email[0]
	}
	if len(params.Name) == 1 && params.Name[0] != "" {
		updated.Name = params.Name[0]
	}
	if len(params.Status) == 1 {
		updated.Status = pbStatusToStatus[params.Status[0]]
	}
	if len(params.PasswordNew) == 1 && params.Password[0] != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(params.PasswordNew[0]), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updated.Password = pass
	}

	s.users.UpdateUser(&updated)
	if err := s.publishUser(ctx, ENTITY_USER, orig, &updated); err != nil {
		log.With("error", err).Error("unable to publish user event")
	}
	if len(params.PasswordNew) != 0 {
		if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_PASSWORD_CHANGE, updated, ""); err != nil {
			log.With("error", err).Error("unable to publish user security event")
		}
	}

	return s.toProtoUser(&updated), nil
}

func (s *UserServer) ComparePassword(ctx context.Context, params *genpb.AuthenticateParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("check password")

	user, err := s.users.GetById(params.UserId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	if !passwordCompare(user, params.Password) {
		return nil, twirp.InvalidArgument.Error("Password mismatch")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) GetSettings(ctx context.Context, params *genpb.UserIdParam) (*genpb.UserSettings, error) {
	user, err := s.users.GetById(params.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	return s.toProtoSettings(user), nil
}

func (s *UserServer) SetSettings(ctx context.Context, params *genpb.UserSettingsUpdate) (*genpb.UserSettings, error) {
	log := logger.FromContext(ctx)
	orig, err := s.users.GetById(params.UserId)

	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if orig == nil {
		return nil, twirp.NotFound.Error("user ID not found")
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

	s.users.UpdateUser(&updated)
	if err := s.publishSettings(ctx, ENTITY_SETTINGS, orig, &updated); err != nil {
		log.With("error", err).Error("unable to publish user security event")
	}

	return s.toProtoSettings(&updated), nil
}

func (s *UserServer) getUserByVerification(ctx context.Context, params *genpb.VerificationParam) (*User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("getUserByVerification BEGIN")

	user, err := s.users.GetById(params.UserId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		log.Info("User not found")
		return nil, twirp.NotFound.Error("user ID not found")
	}
	if user.VerificationToken == nil || user.VerificationExpires == nil {
		log.Info("User has no verification token")
		return nil, twirp.InvalidArgument.Error("user not found")
	}
	if user.VerificationExpires.Before(time.Now()) {
		log.Info("Token is expired")
		// Should we remove it at this point?
		return nil, twirp.InvalidArgument.Error("expired")
	}
	if !tokenCompare(user, params.Token) {
		log.Info("Token mismatch")
		return nil, twirp.InvalidArgument.Error("bad token")
	}

	return user, nil
}

func (s *UserServer) VerificationVerify(ctx context.Context, params *genpb.VerificationParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("Verification email")

	user, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}

	update := *user

	// No longer valid
	update.VerificationExpires = nil
	update.VerificationToken = nil
	// You could be "REGISTERED" or "INVITED"
	update.Status = UserStatus_ACTIVE
	s.users.UpdateUser(&update)

	if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
		log.With("error", err).Info("Publish failed")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) VerificationUpdate(ctx context.Context, params *genpb.VerificationParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("Verification update")

	user, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}
	if user.Status == UserStatus_DISABLED {
		return nil, twirp.NotFound.Error("user is disabled")
	}

	update := *user

	if params.Password != "" {
		pass, errmsg := passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, twirp.InvalidArgument.Error(errmsg)
		}

		update.Password = pass
		s.users.UpdateUser(&update)
	}
	// No longer valid
	update.VerificationExpires = nil
	update.VerificationToken = nil
	// You could be "REGISTERED" or "INVITED"
	update.Status = UserStatus_ACTIVE
	s.users.UpdateUser(&update)

	if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
		log.With("error", err).Info("Publish failed")
	}
	if params.Password != "" {
		if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_PASSWORD_CHANGE, *user, ""); err != nil {
			log.With("error", err).Info("Publish security failed")
		}
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotSend(ctx context.Context, params *genpb.FindParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("Forgot send")

	if params.Email == "" {
		return nil, twirp.InvalidArgument.Error("Must provide email")
	}
	user, err := s.users.GetByEmail(params.Email)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if user == nil {
		return nil, twirp.NotFound.Error("user not found")
	}

	update := *user
	vExpires := time.Now().Add(time.Duration(24 * time.Hour))
	vToken, secret, err := tokenEncrypt(shortuuid.New())
	if err != nil {
		return nil, err
	}
	update.VerificationToken = &vToken
	update.VerificationExpires = &vExpires
	s.users.UpdateUser(&update)

	if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
		log.With("error", err).Info("Publish failed")
	}
	if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_FORGOT_REQUEST, update, secret); err != nil {
		log.With("error", err).Info("Publish security failed")
	}

	return nil, nil
}

package user

import (
	"fmt"
	"strings"
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
	var auth *UserAuth
	var err error
	if params.Email != "" {
		user, err = s.users.GetByEmail(params.Email)
	} else if params.UserId != "" {
		user, err = s.users.GetById(params.UserId)
	} else if params.Auth.Provider != "" && params.Auth.ProviderId != "" {
		auth, err = s.users.AuthGet(params.Auth.Provider, params.Auth.ProviderId)
		if auth != nil {
			user, err = s.users.GetById(auth.UserID)
		}
	} else {
		return nil, twirp.NotFoundError("no query provided")
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

	log.With("email", params.Email).Info("Checking for duplicate")
	if u, err := s.users.GetByEmail(params.Email); err != nil {
		return nil, twirp.InternalErrorWith(err)
	} else if u != nil {
		return nil, twirp.AlreadyExists.Error("Email address not found")
	}
	log.With("email", params.Email).Info("DONE Checking for duplicate")

	pass := []byte{}
	if params.Password != "" {
		var errmsg string

		pass, errmsg = passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, twirp.InvalidArgument.Error(errmsg)
		}
	}
	log.With("email", params.Email).Info("DONE encrypt password")

	userId := uuid.New().String()

	var vExpires time.Time
	vToken := []byte{}
	var secret string
	if params.Status != genpb.UserStatus_ACTIVE {
		var err error
		vExpires = time.Now().Add(time.Duration(24 * time.Hour))
		vToken, secret, err = hmacCreate(userId, shortuuid.New())
		if err != nil {
			return nil, err
		}
	}
	log.With("email", params.Email).Info("DONE token encrypt")

	user := User{
		ID:       userId,
		Name:     params.Name,
		Email:    params.Email,
		Status:   pbStatusToStatus[params.Status],
		Settings: map[string]map[string]string{},

		EmailVerifyToken:     vToken,
		EmailVerifyExpiresAt: &vExpires,
	}
	auth := UserAuth{
		UserID:   user.ID,
		Password: pass,
	}

	log.With("userId", user.ID, "email", user.Email).Info("Saving user to store")

	if err := s.users.CreateUser(user); err != nil {
		return nil, err
	}
	if err := s.users.AuthUpsert("email", strings.ToLower(user.Email), auth); err != nil {
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
		auth, err := s.users.AuthGet("email", strings.ToLower(orig.Email))
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		if len(params.PasswordNew) == 1 && auth != nil {
			// If you're setting a new password, you must provide the old
			if len(params.Password) == 0 {
				return nil, twirp.InvalidArgument.Error("Password missing")
			}
			if errmsg := validatePassword(params.PasswordNew[0]); errmsg != "" {
				return nil, twirp.InvalidArgument.Error("Password too short")
			}
		}
		// If you provided a password, we will check it...
		if len(params.Password) != 0 && !passwordCompare(auth.Password, params.Password[0]) {
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

		auth := UserAuth{
			UserID:   params.UserId,
			Password: pass,
		}
		// Hmm... This is a good case where they both shouldn't be updated at the same time
		if err := s.users.AuthUpsert("email", strings.ToLower(updated.Email), auth); err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
	}

	s.users.UpdateUser(&updated)

	// If the email changed, so move the password to the new authentication
	if updated.Email != orig.Email {
		oldAuth, err := s.users.AuthGet("email", strings.ToLower(orig.Email))
		if err != nil {
			log.With("email", orig.Email, "error", err).Error("unable to get old authentication")
		} else if oldAuth != nil {
			err = s.users.AuthUpsert("email", strings.ToLower(updated.Email), *oldAuth)
			if err != nil {
				// this is "bad"
				log.With("email", orig.Email, "error", err).Error("unable to delete old authentication")
			} else {
				err = s.users.AuthDelete("email", strings.ToLower(orig.Email), *oldAuth)
				if err != nil {
					log.With("email", orig.Email, "error", err).Error("unable to delete old authentication")
				}
			}
		}
	}

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

func (s *UserServer) ComparePassword(ctx context.Context, params *genpb.AuthenticateParam) (*genpb.UserIdParam, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("check password")

	auth, err := s.users.AuthGet("email", strings.ToLower(params.Email))
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if auth == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	if !passwordCompare(auth.Password, params.Password) {
		return nil, twirp.InvalidArgument.Error("Password mismatch")
	}

	return &genpb.UserIdParam{
		UserId: auth.UserID,
	}, nil
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

func (s *UserServer) getUserByVerification(ctx context.Context, params *genpb.VerificationParam) (*UserAuth, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("getUserByVerification BEGIN")

	auth, err := s.users.AuthGet("forgot", params.UserId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	if auth == nil {
		log.Info("User not found")
		return nil, twirp.NotFound.Error("user ID not found")
	}
	if len(auth.Password) == 0 || auth.ExpiresAt == nil {
		log.Info("User has no verification token")
		return nil, twirp.InvalidArgument.Error("user not found")
	}
	if auth.ExpiresAt.Before(time.Now()) {
		log.Info("Token is expired")
		// Should we remove it at this point?
		return nil, twirp.InvalidArgument.Error("expired")
	}
	if !passwordCompare(auth.Password, params.Token) {
		log.Info("Token mismatch")
		return nil, twirp.InvalidArgument.Error("bad token")
	}

	return auth, nil
}

// Verify the email address is "owned" by you
func (s *UserServer) VerificationVerify(ctx context.Context, params *genpb.VerificationParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("Verification email")

	user, err := s.users.GetById(params.UserId)
	if err != nil {
		return nil, err
	}
	if len(user.EmailVerifyToken) == 0 || user.EmailVerifyExpiresAt == nil {
		return nil, twirp.NotFoundError("email may already be verified")
	}
	if user.EmailVerifyExpiresAt.Before(time.Now()) {
		log.Info("Token is expired")
		// Should we remove it at this point?
		return nil, twirp.InvalidArgument.Error("expired")
	}
	if ok, err := hmacCompare(user.ID, params.Token, user.EmailVerifyToken); err != nil {
		return nil, twirp.InternalErrorWith(err)
	} else if !ok {
		log.Info("Token mismatch")
		return nil, twirp.InvalidArgument.Error("bad token")
	}

	update := *user

	// No longer valid
	update.EmailVerifyToken = []byte{}
	update.EmailVerifyExpiresAt = nil
	// Mark this email as verified
	update.VerifiedEmails = append(update.VerifiedEmails, user.Email)
	// You could be "REGISTERED" or "INVITED"
	update.Status = UserStatus_ACTIVE
	s.users.UpdateUser(&update)

	if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
		log.With("error", err).Info("Publish failed")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotVerify(ctx context.Context, params *genpb.VerificationParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("START ForgotVerify")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}
	user, err := s.users.GetById(auth.UserID)
	if err != nil {
		return nil, err
	}
	if user.Status == UserStatus_DISABLED {
		return nil, twirp.NotFound.Error("user is disabled")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotUpdate(ctx context.Context, params *genpb.VerificationParam) (*genpb.User, error) {
	log := logger.FromContext(ctx).With("userId", params.UserId)
	log.Info("START ForgotUpdate")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}
	user, err := s.users.GetById(auth.UserID)
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

		auth := UserAuth{
			UserID:   user.ID,
			Password: pass,
		}
		if err := s.users.AuthUpsert("email", strings.ToLower(user.Email), auth); err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		err = s.users.AuthDelete("forgot", user.ID, UserAuth{
			UserID: user.ID,
		})
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
	}

	// You could be "REGISTERED" or "INVITED"
	if user.Status != UserStatus_ACTIVE {
		update.Status = UserStatus_ACTIVE
		err := s.users.UpdateUser(&update)
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}

		if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
			log.With("error", err).Info("Publish failed")
		}
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
		return nil, twirp.NotFound.Errorf("user not found email=%s", params.Email)
	}

	vExpires := time.Now().Add(time.Duration(24 * time.Hour))
	vToken, secret, err := tokenEncrypt(shortuuid.New())
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	err = s.users.AuthUpsert("forgot", user.ID, UserAuth{
		UserID:    user.ID,
		Password:  vToken,
		ExpiresAt: &vExpires,
	})
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	// TODO?
	// if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
	// 	log.With("error", err).Info("Publish failed")
	// }
	if err := s.publishSecurity(ctx, log, genpb.UserSecurity_USER_FORGOT_REQUEST, *user, secret); err != nil {
		log.With("error", err).Info("Publish security failed")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) AuthAssociate(ctx context.Context, params *genpb.AuthAssociateParam) (*genpb.UserIdParam, error) {
	auth := UserAuth{
		UserID: params.UserId,
	}

	if err := s.users.AuthUpsert(params.Auth.Provider, params.Auth.ProviderId, auth); err != nil {
		return nil, err
	}

	return &genpb.UserIdParam{UserId: params.UserId}, nil
}

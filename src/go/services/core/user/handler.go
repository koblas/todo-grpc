package user

import (
	"fmt"
	"strings"
	"time"

	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/pkg/types"
	"github.com/renstrom/shortuuid"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type UserId struct{}

func (UserId) Prefix() string { return "U" }

// Server represents the gRPC server
type UserServer struct {
	users  UserStore
	pubsub corepb.UserEventbus
	kms    key_manager.Encoder
}

var pbStatusToStatus = map[corepb.UserStatus]UserStatus{
	corepb.UserStatus_ACTIVE:     UserStatus_ACTIVE,
	corepb.UserStatus_INVITED:    UserStatus_INVITED,
	corepb.UserStatus_DISABLED:   UserStatus_DISABLED,
	corepb.UserStatus_REGISTERED: UserStatus_REGISTERED,
}
var statusToPbStatus = map[UserStatus]corepb.UserStatus{
	UserStatus_ACTIVE:     corepb.UserStatus_ACTIVE,
	UserStatus_INVITED:    corepb.UserStatus_INVITED,
	UserStatus_DISABLED:   corepb.UserStatus_DISABLED,
	UserStatus_REGISTERED: corepb.UserStatus_REGISTERED,
}

type Option func(*UserServer)

func WithUserStore(store UserStore) Option {
	return func(cfg *UserServer) {
		cfg.users = store
	}
}

func WithProducer(bus corepb.UserEventbus) Option {
	return func(cfg *UserServer) {
		cfg.pubsub = bus
	}
}

func NewUserServer(opts ...Option) *UserServer {
	svr := UserServer{
		kms: key_manager.NewSecureClear(),
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (s *UserServer) FindBy(ctx context.Context, params *corepb.UserFindParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("email", params.Email)).With(zap.String("userId", params.UserId))
	log.Info("FindBy")

	var user *User
	var auth *UserAuth
	var err error
	if params.Email != "" {
		user, err = s.users.GetByEmail(ctx, params.Email)
	} else if params.UserId != "" {
		user, err = s.users.GetById(ctx, params.UserId)
	} else if params.Auth.Provider != "" && params.Auth.ProviderId != "" {
		auth, err = s.users.AuthGet(ctx, params.Auth.Provider, params.Auth.ProviderId)
		if auth != nil {
			user, err = s.users.GetById(ctx, auth.UserID)
		}
	} else {
		return nil, twirp.NotFoundError("no query provided")
	}

	if err != nil {
		log.With(zap.Error(err)).Error("FindBy failed")
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("User not found")
	}

	log.With("ID", user.ID).Info("found")

	return s.toProtoUser(user), nil
}

func (s *UserServer) Create(ctx context.Context, params *corepb.UserCreateParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("email", params.Email))
	log.Info("Received Create")

	if params.Status != corepb.UserStatus_REGISTERED && params.Status != corepb.UserStatus_INVITED {
		log.Error("Bad user status")
		return nil, fmt.Errorf("invalid user status = %s", params.Status)
	}

	log.Info("Checking for duplicate")
	if u, err := s.users.GetByEmail(ctx, params.Email); err != nil {
		log.With(zap.Error(err)).Error("GetByEmail failed")
		return nil, twirp.InternalErrorWith(err)
	} else if u != nil {
		return nil, twirp.AlreadyExists.Error("Email address not found")
	}
	log.Info("DONE Checking for duplicate")

	pass := []byte{}
	if params.Password != "" {
		var errmsg string

		pass, errmsg = passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, twirp.InvalidArgument.Error(errmsg)
		}
	}
	log.Info("DONE encrypt password")

	userId := types.New[UserId]().String()
	// userId := xid.New().String()

	var vExpires time.Time
	vToken := []byte{}
	var secret string
	if params.Status != corepb.UserStatus_ACTIVE {
		var err error
		vExpires = time.Now().Add(time.Duration(24 * time.Hour))
		vToken, secret, err = hmacCreate(userId, shortuuid.New())
		if err != nil {
			return nil, err
		}
	}
	log.Info("DONE token encrypt")

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
	log = log.With(zap.String("userId", user.ID))

	log.Info("Saving user to store")

	if err := s.users.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	if err := s.users.AuthUpsert(ctx, "email", strings.ToLower(user.Email), auth); err != nil {
		return nil, err
	}

	log.Info("User Created")

	if _, err := s.pubsub.UserChange(ctx, &corepb.UserChangeEvent{
		Current: s.toProtoUser(&user),
	}); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}
	if secret != "" {
		token, err := protoutil.SecureValueEncode(s.kms, secret)
		if err != nil {
			log.With(zap.Error(err)).Info("unable to create token")
		} else {
			payload := corepb.UserSecurityEvent{
				User:  s.toProtoUser(&user),
				Token: token,
			}

			if params.Status == corepb.UserStatus_REGISTERED {
				payload.Action = corepb.UserSecurity_USER_REGISTER_TOKEN
			} else if params.Status == corepb.UserStatus_INVITED {
				payload.Action = corepb.UserSecurity_USER_INVITE_TOKEN
			}

			if _, err := s.pubsub.UserSecurity(ctx, &payload); err != nil {
				log.With(zap.Error(err)).Info("user security publish failed")
			}
		}
	}

	return s.toProtoUser(&user), nil
}

func (s *UserServer) Update(ctx context.Context, params *corepb.UserUpdateParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("User Update")

	orig, err := s.users.GetById(ctx, params.UserId)
	if err != nil {
		log.With(zap.Error(err)).Error("Update/GetById failed")
		return nil, twirp.InternalErrorWith(err)
	}

	if orig == nil {
		return nil, twirp.NotFound.Error("User not found")
	}
	// Some basic validation
	if params.Password != nil || params.PasswordNew != nil {
		auth, err := s.users.AuthGet(ctx, "email", strings.ToLower(orig.Email))
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		if params.PasswordNew != nil && auth != nil {
			// If you're setting a new password, you must provide the old
			if params.Password == nil {
				return nil, twirp.InvalidArgument.Error("Password missing")
			}
			if errmsg := validatePassword(*params.PasswordNew); errmsg != "" {
				return nil, twirp.InvalidArgument.Error("Password too short")
			}
		}
		// If you provided a password, we will check it...
		if params.Password != nil && !passwordCompare(auth.Password, *params.Password) {
			return nil, twirp.InvalidArgument.Error("Password mismatch")
		}
	}

	// Now do the updates
	updated := *orig
	if params.Email != nil && *params.Email != "" {
		updated.Email = *params.Email
	}
	if params.Name != nil && *params.Name != "" {
		updated.Name = *params.Name
	}
	if params.Status != nil {
		updated.Status = pbStatusToStatus[*params.Status]
	}
	if params.PasswordNew != nil && *params.Password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(*params.PasswordNew), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		auth := UserAuth{
			UserID:   params.UserId,
			Password: pass,
		}
		// Hmm... This is a good case where they both shouldn't be updated at the same time
		if err := s.users.AuthUpsert(ctx, "email", strings.ToLower(updated.Email), auth); err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
	}
	if params.AvatarUrl != nil {
		updated.AvatarUrl = params.AvatarUrl
	}

	s.users.UpdateUser(ctx, &updated)

	// If the email changed, so move the password to the new authentication
	if updated.Email != orig.Email {
		logInner := log.With(zap.String("email", orig.Email))
		oldAuth, err := s.users.AuthGet(ctx, "email", strings.ToLower(orig.Email))
		if err != nil {
			logInner.With(zap.Error(err)).Error("unable to get old authentication")
		} else if oldAuth != nil {
			err = s.users.AuthUpsert(ctx, "email", strings.ToLower(updated.Email), *oldAuth)
			if err != nil {
				// this is "bad"
				logInner.With(zap.Error(err)).Error("unable to delete old authentication")
			} else {
				err = s.users.AuthDelete(ctx, "email", strings.ToLower(orig.Email), *oldAuth)
				if err != nil {
					logInner.With(zap.Error(err)).Error("unable to delete old authentication")
				}
			}
		}
	}

	if _, err := s.pubsub.UserChange(ctx, &corepb.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	}); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	if params.PasswordNew != nil {
		if _, err := s.pubsub.UserSecurity(ctx, &corepb.UserSecurityEvent{
			Action: corepb.UserSecurity_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(&updated),
		}); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return s.toProtoUser(&updated), nil
}

func (s *UserServer) ComparePassword(ctx context.Context, params *corepb.AuthenticateParam) (*corepb.UserIdParam, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("check password")

	auth, err := s.users.AuthGet(ctx, "email", strings.ToLower(params.Email))
	if err != nil {
		log.With(zap.Error(err)).Error("AuthGet failed")
		return nil, twirp.InternalErrorWith(err)
	}

	if auth == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	if !passwordCompare(auth.Password, params.Password) {
		return nil, twirp.InvalidArgument.Error("Password mismatch")
	}

	return &corepb.UserIdParam{
		UserId: auth.UserID,
	}, nil
}

func (s *UserServer) GetSettings(ctx context.Context, params *corepb.UserIdParam) (*corepb.UserSettings, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	user, err := s.users.GetById(ctx, params.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("GetSettings/GetById failed")
		return nil, twirp.InternalErrorWith(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	return s.toProtoSettings(user), nil
}

func (s *UserServer) SetSettings(ctx context.Context, params *corepb.UserSettingsUpdate) (*corepb.UserSettings, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	orig, err := s.users.GetById(ctx, params.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("SetSettings/GetById failed")
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

	s.users.UpdateUser(ctx, &updated)
	if _, err := s.pubsub.UserChange(ctx, &corepb.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	}); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return s.toProtoSettings(&updated), nil
}

func (s *UserServer) getUserByVerification(ctx context.Context, params *corepb.VerificationParam) (*UserAuth, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("getUserByVerification BEGIN")

	auth, err := s.users.AuthGet(ctx, "forgot", params.UserId)
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
func (s *UserServer) VerificationVerify(ctx context.Context, params *corepb.VerificationParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("Verification email")

	user, err := s.users.GetById(ctx, params.UserId)
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
	s.users.UpdateUser(ctx, &update)

	if _, err := s.pubsub.UserChange(ctx, &corepb.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	}); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotVerify(ctx context.Context, params *corepb.VerificationParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("START ForgotVerify")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}
	user, err := s.users.GetById(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	if user.Status == UserStatus_DISABLED {
		return nil, twirp.NotFound.Error("user is disabled")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotUpdate(ctx context.Context, params *corepb.VerificationParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("START ForgotUpdate")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, err
	}
	user, err := s.users.GetById(ctx, auth.UserID)
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
		if err := s.users.AuthUpsert(ctx, "email", strings.ToLower(user.Email), auth); err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		err = s.users.AuthDelete(ctx, "forgot", user.ID, UserAuth{
			UserID: user.ID,
		})
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
	}

	// You could be "REGISTERED" or "INVITED"
	if user.Status != UserStatus_ACTIVE {
		update.Status = UserStatus_ACTIVE
		err := s.users.UpdateUser(ctx, &update)
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
	}

	if _, err := s.pubsub.UserChange(ctx, &corepb.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	}); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}
	if params.Password != "" {
		if _, err := s.pubsub.UserSecurity(ctx, &corepb.UserSecurityEvent{
			Action: corepb.UserSecurity_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(user),
		}); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) ForgotSend(ctx context.Context, params *corepb.UserFindParam) (*corepb.User, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("Forgot send")

	if params.Email == "" {
		return nil, twirp.InvalidArgument.Error("Must provide email")
	}
	user, err := s.users.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if user == nil {
		return nil, twirp.NotFound.Errorf("user not found email=%s", params.Email)
	}

	vExpires := time.Now().Add(time.Duration(24 * time.Hour))
	// use shortuuid rather than xid since it's less sequential
	vToken, secret, err := tokenEncrypt(shortuuid.New())
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	err = s.users.AuthUpsert(ctx, "forgot", user.ID, UserAuth{
		UserID:    user.ID,
		Password:  vToken,
		ExpiresAt: &vExpires,
	})
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	// TODO?
	// if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
	// 	log.With(zap.Error(err)).Info("Publish failed")
	// }
	if token, err := protoutil.SecureValueEncode(s.kms, secret); err != nil {
		log.With(zap.Error(err)).Info("failed to encrypt token")
	} else if _, err := s.pubsub.UserSecurity(ctx, &corepb.UserSecurityEvent{
		Action: corepb.UserSecurity_USER_FORGOT_REQUEST,
		User:   s.toProtoUser(user),
		Token:  token,
	}); err != nil {
		log.With(zap.Error(err)).Info("user security publish failed")
	}

	return s.toProtoUser(user), nil
}

func (s *UserServer) AuthAssociate(ctx context.Context, params *corepb.AuthAssociateParam) (*corepb.UserIdParam, error) {
	auth := UserAuth{
		UserID: params.UserId,
	}

	if err := s.users.AuthUpsert(ctx, params.Auth.Provider, params.Auth.ProviderId, auth); err != nil {
		return nil, err
	}

	return &corepb.UserIdParam{UserId: params.UserId}, nil
}

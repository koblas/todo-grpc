package user

import (
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/renstrom/shortuuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var _ userv1connect.UserServiceHandler = (*UserServer)(nil)

type UserId struct{}

func (UserId) Prefix() string { return "U" }

const EMAIL_PROVIDER = "__email__"
const FORGOT_PROVIDER = "__forgot__"

type Store interface {
	UserStore
	OAuthStore
	TeamStore
}

// Server represents the gRPC server
type UserServer struct {
	store  Store
	pubsub eventbusv1connect.UserEventbusServiceClient
	kms    key_manager.Encoder
}

var pbStatusToStatus = map[userv1.UserStatus]UserStatus{
	userv1.UserStatus_USER_STATUS_ACTIVE:     UserStatus_ACTIVE,
	userv1.UserStatus_USER_STATUS_REGISTERED: UserStatus_REGISTERED,
	userv1.UserStatus_USER_STATUS_INVITED:    UserStatus_INVITED,
}
var statusToPbStatus = map[UserStatus]userv1.UserStatus{
	UserStatus_ACTIVE:     userv1.UserStatus_USER_STATUS_ACTIVE,
	UserStatus_REGISTERED: userv1.UserStatus_USER_STATUS_REGISTERED,
	UserStatus_INVITED:    userv1.UserStatus_USER_STATUS_INVITED,
}

var pbClosedStatusToStatus = map[userv1.ClosedStatus]ClosedStatus{
	userv1.ClosedStatus_CLOSED_STATUS_UNSPECIFIED: ClosedStatus_ACTIVE,
	userv1.ClosedStatus_CLOSED_STATUS_DELETED:     ClosedStatus_DELETED,
	userv1.ClosedStatus_CLOSED_STATUS_DISABLED:    ClosedStatus_DISABLED,
}
var closedStatusToPbStatus = map[ClosedStatus]userv1.ClosedStatus{
	ClosedStatus_ACTIVE:   userv1.ClosedStatus_CLOSED_STATUS_UNSPECIFIED,
	ClosedStatus_DELETED:  userv1.ClosedStatus_CLOSED_STATUS_DELETED,
	ClosedStatus_DISABLED: userv1.ClosedStatus_CLOSED_STATUS_DISABLED,
}

type Option func(*UserServer)

func WithStore(store Store) Option {
	return func(cfg *UserServer) {
		cfg.store = store
	}
}

func WithProducer(bus eventbusv1connect.UserEventbusServiceClient) Option {
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

func (s *UserServer) FindBy(ctx context.Context, request *connect.Request[userv1.FindByRequest]) (*connect.Response[userv1.FindByResponse], error) {
	params := request.Msg.FindBy
	log := logger.FromContext(ctx).With(zap.String("email", params.Email)).With(zap.String("userId", params.UserId))
	log.Info("FindBy")

	var user *User
	var auth *UserAuth
	var err error
	if params.Email != "" {
		user, err = s.store.GetByEmail(ctx, params.Email)
	} else if params.UserId != "" {
		user, err = s.store.GetById(ctx, params.UserId)
	} else if params.Auth.Provider != "" && params.Auth.ProviderId != "" {
		auth, err = s.store.AuthGet(ctx, params.Auth.Provider, params.Auth.ProviderId)
		if auth != nil {
			user, err = s.store.GetById(ctx, auth.UserID)
		}
	} else {
		return nil, bufcutil.NotFoundError("no query provided")
	}

	if err != nil {
		log.With(zap.Error(err)).Error("FindBy failed")
		return nil, bufcutil.InternalError(err)
	}

	if user == nil {
		return nil, bufcutil.NotFoundError("user not found")
	}

	log.With("ID", user.ID).Info("found")

	return connect.NewResponse(&userv1.FindByResponse{User: s.toProtoUser(user)}), nil
}

// Create is used in the registration process to establish a new account
//
// This is intended to be the entry point for creating a user if they are initially creating an account.
// It could be called from an OAuth flow where they're creating the account for the first time, but they
// then may already have either an existing email address registered in this system, or they could have pending invites
// which means that they have an existing account
func (s *UserServer) Create(ctx context.Context, request *connect.Request[userv1.CreateRequest]) (*connect.Response[userv1.CreateResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("email", params.Email))
	log.Info("Received create")

	// ACTIVE == Create this with a "known" good email address
	// REGISTERED == Create this with an unknown email address

	if params.Status != userv1.UserStatus_USER_STATUS_ACTIVE && params.Status != userv1.UserStatus_USER_STATUS_REGISTERED && params.Status != userv1.UserStatus_USER_STATUS_INVITED {
		log.With("status", params.Status).Error("Bad user status")
		return nil, bufcutil.InvalidArgumentError("status", "invalid user status")
	}

	log.Info("Looking up user")
	user, err := s.store.GetByEmail(ctx, params.Email)
	if err != nil && err != ErrorUserNotFound {
		log.With(zap.Error(err)).Error("GetByEmail failed")
		return nil, bufcutil.InternalError(err)
	}
	// User already exists and they're trying to REGISTER new account
	if user != nil && params.Status != userv1.UserStatus_USER_STATUS_ACTIVE {
		return nil, connect.NewError(connect.CodeAlreadyExists, nil)
	}

	// Get the password encryption out of the way, this is to prevent the creation of
	//  a user account that might have an error due to password rules (e.g. length)
	pass := []byte{}
	if params.Password != "" {
		var errmsg string

		pass, errmsg = passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, bufcutil.InvalidArgumentError("password", errmsg)
		}
	}
	log.Info("DONE encrypt password")

	// If user exists, then we're coming in through an "authenticated" channel
	var secret string
	created := false
	if user == nil {
		user = &User{
			Name:         params.Name,
			Email:        params.Email,
			Status:       pbStatusToStatus[params.Status],
			ClosedStatus: ClosedStatus_ACTIVE,
			Settings:     map[string]map[string]string{},
		}

		if params.Status != userv1.UserStatus_USER_STATUS_ACTIVE {
			var err error
			vExpires := time.Now().Add(time.Duration(24 * time.Hour))
			vNonce, err := randomBytes(20)
			if err != nil {
				return nil, bufcutil.InternalError(err, "failed to generate nonce")
			}
			secret, err = randomString(20)
			if err != nil {
				return nil, bufcutil.InternalError(err, "failed to generate secret")
			}
			vToken, err := hmacCreate(vNonce, secret)
			if err != nil {
				return nil, bufcutil.InternalError(err, "failed to hash token")
			}

			user.EmailVerifyNonce = vNonce
			user.EmailVerifyToken = vToken
			user.EmailVerifyExpiresAt = &vExpires
			log.Info("verification token creation done")
		}

		log.Info("Saving user to store")
		if u, err := s.store.CreateUser(ctx, *user); err != nil {
			return nil, bufcutil.InternalError(err, "db create failed")
		} else {
			user = u
		}
		log = log.With("userId", user.ID)
		log.Info("user created")
		created = true
	} else {
		log = log.With("userId", user.ID)
		log.Info("user use existing")
	}

	// If we have a password created for this user save it
	if len(pass) != 0 {
		auth := UserAuth{
			UserID:   user.ID,
			Password: pass,
		}
		if err := s.store.AuthUpsert(ctx, EMAIL_PROVIDER, strings.ToLower(user.Email), auth); err != nil {
			return nil, bufcutil.InternalError(err, "db authentication create failed")
		}
		log.Info("Password authentication created")
	}
	if params.Status == userv1.UserStatus_USER_STATUS_INVITED {
		secret, err = s.createForgot(ctx, user)
		if err != nil {
			log.With(zap.Error(err)).Error("unable create create forgot token")
			return nil, bufcutil.InternalError(err)
		}
	}

	if created {
		if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
			Current: s.toProtoUser(user),
		})); err != nil {
			log.With(zap.Error(err)).Info("user entity publish failed")
		}

		// User Created in state REGISTERED
		if secret != "" {
			token, err := protoutil.SecureValueEncode(s.kms, secret)
			if err != nil {
				log.With(zap.Error(err)).Error("unable to create registration token")
				// FIXME -- the user is created, but they will not get a REGISTRATION token
			} else {
				payload := userv1.UserSecurityEvent{
					User:  s.toProtoUser(user),
					Token: token,
				}

				if user.Status == UserStatus_REGISTERED {
					payload.Action = userv1.UserSecurity_USER_SECURITY_USER_REGISTER_TOKEN
				} else if user.Status == UserStatus_INVITED {
					payload.Action = userv1.UserSecurity_USER_SECURITY_USER_INVITE_TOKEN
				} else {
					log.With(zap.String("status", user.Status.String())).Error("unknown user status")
				}

				if _, err := s.pubsub.SecurityRegisterToken(ctx, connect.NewRequest(&payload)); err != nil {
					log.With(zap.Error(err)).Error("user security publish failed")
				}
			}
		} else {
			// User created in state ACTIVE
		}
	} else if user.Status == UserStatus_REGISTERED && user.Status != pbStatusToStatus[params.Status] {
		// Through an OAuth flow we're creating an ACTIVE user
		orig := s.toProtoUser(user)

		user.Status = pbStatusToStatus[params.Status]
		if err := s.store.UpdateUser(ctx, user); err != nil {
			log.With(zap.Error(err)).Error("unable to update user status")
			return nil, bufcutil.InternalError(err)
		}

		// We're changing from REGISTERED to ACTIVE
		if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
			Original: orig,
			Current:  s.toProtoUser(user),
		})); err != nil {
			log.With(zap.Error(err)).Info("user entity publish failed")
		}
	}

	if user.Status == UserStatus_ACTIVE {
		// TODO - something with Invites
	}

	return connect.NewResponse(&userv1.CreateResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) Update(ctx context.Context, request *connect.Request[userv1.UpdateRequest]) (*connect.Response[userv1.UpdateResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("User Update")

	orig, err := s.store.GetById(ctx, params.UserId)
	if err != nil {
		log.With(zap.Error(err)).Error("Update/GetById failed")
		return nil, bufcutil.InternalError(err)
	}

	if orig == nil {
		return nil, bufcutil.NotFoundError("User not found")
	}
	// Some basic validation
	if params.Password != nil || params.PasswordNew != nil {
		auth, err := s.store.AuthGet(ctx, EMAIL_PROVIDER, strings.ToLower(orig.Email))
		if err != nil {
			return nil, bufcutil.InternalError(err)
		}
		if params.PasswordNew != nil && auth != nil {
			// If you're setting a new password, you must provide the old
			if params.Password == nil {
				return nil, bufcutil.InvalidArgumentError("password", "password missing")
			}
			if errmsg := validatePassword(*params.PasswordNew); errmsg != "" {
				return nil, bufcutil.InvalidArgumentError("password", "password too short")
			}
		}
		// If you provided a password, we will check it...
		if params.Password != nil && !passwordCompare(auth.Password, *params.Password) {
			return nil, bufcutil.InvalidArgumentError("password", "password mismatch")
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

		// Clear out these values if we're now registered
		if updated.Status == UserStatus_REGISTERED {
			updated.EmailVerifyNonce = nil
			updated.EmailVerifyToken = nil
			updated.EmailVerifyExpiresAt = nil
		}
	}
	if params.ClosedStatus != nil {
		// TODO
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
		if err := s.store.AuthUpsert(ctx, EMAIL_PROVIDER, strings.ToLower(updated.Email), auth); err != nil {
			return nil, bufcutil.InternalError(err)
		}
	}
	if params.AvatarUrl != nil {
		updated.AvatarUrl = params.AvatarUrl
	}

	s.store.UpdateUser(ctx, &updated)

	// If the email changed, so move the password to the new authentication
	if updated.Email != orig.Email {
		logInner := log.With(zap.String("email", orig.Email))
		oldAuth, err := s.store.AuthGet(ctx, EMAIL_PROVIDER, strings.ToLower(orig.Email))
		if err != nil {
			logInner.With(zap.Error(err)).Error("unable to get old authentication")
		} else if oldAuth != nil {
			err = s.store.AuthUpsert(ctx, EMAIL_PROVIDER, strings.ToLower(updated.Email), *oldAuth)
			if err != nil {
				// this is "bad"
				logInner.With(zap.Error(err)).Error("unable to delete old authentication")
			} else {
				err = s.store.AuthDelete(ctx, EMAIL_PROVIDER, strings.ToLower(orig.Email), *oldAuth)
				if err != nil {
					logInner.With(zap.Error(err)).Error("unable to delete old authentication")
				}
			}
		}
	}

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	if params.PasswordNew != nil {
		if _, err := s.pubsub.SecurityPasswordChange(ctx, connect.NewRequest(&userv1.UserSecurityEvent{
			Action: userv1.UserSecurity_USER_SECURITY_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(&updated),
		})); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return connect.NewResponse(&userv1.UpdateResponse{User: s.toProtoUser(&updated)}), nil
}

func (s *UserServer) ComparePassword(ctx context.Context, request *connect.Request[userv1.ComparePasswordRequest]) (*connect.Response[userv1.ComparePasswordResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("check password")

	auth, err := s.store.AuthGet(ctx, EMAIL_PROVIDER, strings.ToLower(params.Email))
	if err != nil {
		log.With(zap.Error(err)).Error("AuthGet failed")
		return nil, bufcutil.InternalError(err)
	}

	if auth == nil {
		return nil, bufcutil.NotFoundError("User not found")
	}

	if !passwordCompare(auth.Password, params.Password) {
		return nil, bufcutil.InvalidArgumentError("password", "password mismatch")
	}

	return connect.NewResponse(&userv1.ComparePasswordResponse{
		UserId: auth.UserID,
	}), nil
}

func (s *UserServer) AuthAssociate(ctx context.Context, params *connect.Request[userv1.AuthAssociateRequest]) (*connect.Response[userv1.AuthAssociateResponse], error) {
	auth := UserAuth{
		UserID: params.Msg.UserId,
	}

	if err := s.store.AuthUpsert(ctx, params.Msg.Auth.Provider, params.Msg.Auth.ProviderId, auth); err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&userv1.AuthAssociateResponse{UserId: params.Msg.UserId}), nil
}

// Settings

func (s *UserServer) GetSettings(ctx context.Context, params *connect.Request[userv1.UserServiceGetSettingsRequest]) (*connect.Response[userv1.UserServiceGetSettingsResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.Msg.UserId))
	user, err := s.store.GetById(ctx, params.Msg.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("GetSettings/GetById failed")
		return nil, bufcutil.InternalError(err)
	}

	if user == nil {
		return nil, bufcutil.NotFoundError("User not found")
	}

	return connect.NewResponse(&userv1.UserServiceGetSettingsResponse{Settings: s.toProtoSettings(user)}), nil
}

func (s *UserServer) SetSettings(ctx context.Context, request *connect.Request[userv1.UserServiceSetSettingsRequest]) (*connect.Response[userv1.UserServiceSetSettingsResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	orig, err := s.store.GetById(ctx, params.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("SetSettings/GetById failed")
		return nil, bufcutil.InternalError(err)
	}
	if orig == nil {
		return nil, bufcutil.NotFoundError("User not found")
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

	s.store.UpdateUser(ctx, &updated)
	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return connect.NewResponse(&userv1.UserServiceSetSettingsResponse{Settings: s.toProtoSettings(&updated)}), nil
}

func (s *UserServer) getUserByVerification(ctx context.Context, params *userv1.Verification) (*UserAuth, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("getUserByVerification BEGIN")

	auth, err := s.store.AuthGet(ctx, FORGOT_PROVIDER, params.UserId)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	if auth == nil {
		log.Info("User not found")
		return nil, bufcutil.NotFoundError("User not found")
	}
	if len(auth.Password) == 0 || auth.ExpiresAt == nil {
		log.Info("User has no verification token")
		return nil, bufcutil.InvalidArgumentError("token", "user not found")
	}
	if auth.ExpiresAt.Before(time.Now()) {
		log.Info("Token is expired")
		// Should we remove it at this point?
		return nil, bufcutil.InvalidArgumentError("token", "expired")
	}
	if !passwordCompare(auth.Password, params.Token) {
		log.Info("Token mismatch")
		return nil, bufcutil.InvalidArgumentError("token", "bad match")
	}

	return auth, nil
}

// Verify the email address is "owned" by you
func (s *UserServer) VerificationVerify(ctx context.Context, request *connect.Request[userv1.VerificationVerifyRequest]) (*connect.Response[userv1.VerificationVerifyResponse], error) {
	params := request.Msg.Verification
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("Verification email")

	user, err := s.store.GetById(ctx, params.UserId)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	if user.EmailVerifyExpiresAt == nil {
		return nil, bufcutil.NotFoundError("email may already be verified")
	}
	if user.EmailVerifyExpiresAt.Before(time.Now()) {
		log.Info("Token is expired")
		// Should we remove it at this point?
		return nil, bufcutil.InvalidArgumentError("token", "expired")
	}
	if ok, err := hmacCompare(user.EmailVerifyNonce, params.Token, user.EmailVerifyToken); err != nil {
		return nil, bufcutil.InternalError(err)
	} else if !ok {
		log.Info("Token mismatch")
		return nil, bufcutil.InvalidArgumentError("token", "bad token")
	}

	update := *user

	// No longer valid
	update.EmailVerifyNonce = nil
	update.EmailVerifyToken = nil
	update.EmailVerifyExpiresAt = nil
	// Mark this email as verified
	update.VerifiedEmails = append(update.VerifiedEmails, user.Email)
	// You could be "REGISTERED" or "INVITED"
	update.Status = UserStatus_ACTIVE
	s.store.UpdateUser(ctx, &update)

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return connect.NewResponse(&userv1.VerificationVerifyResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) ForgotVerify(ctx context.Context, request *connect.Request[userv1.ForgotVerifyRequest]) (*connect.Response[userv1.ForgotVerifyResponse], error) {
	params := request.Msg.Verification
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("START ForgotVerify")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	user, err := s.store.GetById(ctx, auth.UserID)
	if err != nil {
		return nil, bufcutil.NotFoundError("user not found")
	}
	if user.ClosedStatus != ClosedStatus_ACTIVE {
		return nil, bufcutil.FailedPreconditionError("user is disabled")
	}

	return connect.NewResponse(&userv1.ForgotVerifyResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) ForgotUpdate(ctx context.Context, request *connect.Request[userv1.ForgotUpdateRequest]) (*connect.Response[userv1.ForgotUpdateResponse], error) {
	params := request.Msg.Verification
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("START ForgotUpdate")

	auth, err := s.getUserByVerification(ctx, params)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	user, err := s.store.GetById(ctx, auth.UserID)
	if err != nil {
		return nil, bufcutil.NotFoundError("user not found")
	}
	if user.ClosedStatus != ClosedStatus_ACTIVE {
		return nil, bufcutil.FailedPreconditionError("user is disabled")
	}

	update := *user

	if params.Password != "" {
		pass, errmsg := passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, bufcutil.InvalidArgumentError("password", "no match")
		}

		auth := UserAuth{
			UserID:   user.ID,
			Password: pass,
		}
		if err := s.store.AuthUpsert(ctx, EMAIL_PROVIDER, strings.ToLower(user.Email), auth); err != nil {
			return nil, bufcutil.InternalError(err)
		}
		err = s.store.AuthDelete(ctx, FORGOT_PROVIDER, user.ID, UserAuth{
			UserID: user.ID,
		})
		if err != nil {
			return nil, bufcutil.InternalError(err)
		}
	}

	// You could be "REGISTERED" or "INVITED"
	//  Since you just verified an email address, we know you've validated the email address
	if user.Status != UserStatus_ACTIVE {
		update.Status = UserStatus_ACTIVE
		err := s.store.UpdateUser(ctx, &update)
		if err != nil {
			return nil, bufcutil.InternalError(err)
		}
	}

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&userv1.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}
	if params.Password != "" {
		if _, err := s.pubsub.SecurityPasswordChange(ctx, connect.NewRequest(&userv1.UserSecurityEvent{
			Action: userv1.UserSecurity_USER_SECURITY_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(user),
		})); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return connect.NewResponse(&userv1.ForgotUpdateResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) createForgot(ctx context.Context, user *User) (string, error) {
	vExpires := time.Now().Add(time.Duration(24 * time.Hour))
	// use shortuuid rather than xid since it's less sequential
	vToken, secret, err := tokenEncrypt(shortuuid.New())
	if err != nil {
		return "", err
	}
	err = s.store.AuthUpsert(ctx, FORGOT_PROVIDER, user.ID, UserAuth{
		UserID:    user.ID,
		Password:  vToken,
		ExpiresAt: &vExpires,
	})
	if err != nil {
		return "", err
	}

	return secret, nil
}

func (s *UserServer) ForgotSend(ctx context.Context, request *connect.Request[userv1.ForgotSendRequest]) (*connect.Response[userv1.ForgotSendResponse], error) {
	params := request.Msg.FindBy
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("Forgot send")

	if params.Email == "" {
		return nil, bufcutil.InvalidArgumentError("email", "must provide email")
	}
	user, err := s.store.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	if user == nil {
		return nil, bufcutil.NotFoundError("user not found")
	}

	secret, err := s.createForgot(ctx, user)
	if err != nil {
		log.With(zap.Error(err)).Error("unable create create forgot token")
		return nil, bufcutil.InternalError(err)
	}

	// TODO?
	// if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
	// 	log.With(zap.Error(err)).Info("Publish failed")
	// }
	if token, err := protoutil.SecureValueEncode(s.kms, secret); err != nil {
		log.With(zap.Error(err)).Info("failed to encrypt token")
	} else if _, err := s.pubsub.SecurityForgotRequest(ctx, connect.NewRequest(&userv1.UserSecurityEvent{
		Action: userv1.UserSecurity_USER_SECURITY_USER_FORGOT_REQUEST,
		User:   s.toProtoUser(user),
		Token:  token,
	})); err != nil {
		log.With(zap.Error(err)).Info("user security publish failed")
	}

	return connect.NewResponse(&userv1.ForgotSendResponse{User: s.toProtoUser(user)}), nil
}

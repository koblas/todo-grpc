package user

import (
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
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
	pubsub corev1connect.UserEventbusServiceClient
	kms    key_manager.Encoder
}

var pbStatusToStatus = map[corev1.UserStatus]UserStatus{
	corev1.UserStatus_USER_STATUS_ACTIVE:     UserStatus_ACTIVE,
	corev1.UserStatus_USER_STATUS_INVITED:    UserStatus_INVITED,
	corev1.UserStatus_USER_STATUS_DISABLED:   UserStatus_DISABLED,
	corev1.UserStatus_USER_STATUS_REGISTERED: UserStatus_REGISTERED,
}
var statusToPbStatus = map[UserStatus]corev1.UserStatus{
	UserStatus_ACTIVE:     corev1.UserStatus_USER_STATUS_ACTIVE,
	UserStatus_INVITED:    corev1.UserStatus_USER_STATUS_INVITED,
	UserStatus_DISABLED:   corev1.UserStatus_USER_STATUS_DISABLED,
	UserStatus_REGISTERED: corev1.UserStatus_USER_STATUS_REGISTERED,
}

type Option func(*UserServer)

func WithUserStore(store UserStore) Option {
	return func(cfg *UserServer) {
		cfg.users = store
	}
}

func WithProducer(bus corev1connect.UserEventbusServiceClient) Option {
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

func (s *UserServer) FindBy(ctx context.Context, request *connect.Request[corev1.FindByRequest]) (*connect.Response[corev1.FindByResponse], error) {
	params := request.Msg.FindBy
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
		return nil, bufcutil.InternalError(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("User not found")
	}

	log.With("ID", user.ID).Info("found")

	return connect.NewResponse(&corev1.FindByResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) Create(ctx context.Context, request *connect.Request[corev1.UserServiceCreateRequest]) (*connect.Response[corev1.UserServiceCreateResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("email", params.Email))
	log.Info("Received Create")

	if params.Status != corev1.UserStatus_USER_STATUS_REGISTERED && params.Status != corev1.UserStatus_USER_STATUS_INVITED {
		log.Error("Bad user status")
		return nil, bufcutil.InvalidArgumentError("status", "invalid user status choice")
	}

	log.Info("Checking for duplicate")
	if u, err := s.users.GetByEmail(ctx, params.Email); err != nil {
		log.With(zap.Error(err)).Error("GetByEmail failed")
		return nil, bufcutil.InternalError(err)
	} else if u != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, nil)
	}
	log.Info("DONE Checking for duplicate")

	pass := []byte{}
	if params.Password != "" {
		var errmsg string

		pass, errmsg = passwordEncrypt(params.Password)
		if errmsg != "" {
			return nil, bufcutil.InvalidArgumentError("password", errmsg)
		}
	}
	log.Info("DONE encrypt password")

	userId := types.New[UserId]().String()
	// userId := xid.New().String()

	var vExpires time.Time
	vToken := []byte{}
	var secret string
	if params.Status != corev1.UserStatus_USER_STATUS_ACTIVE {
		var err error
		vExpires = time.Now().Add(time.Duration(24 * time.Hour))
		vToken, secret, err = hmacCreate(userId, shortuuid.New())
		if err != nil {
			return nil, bufcutil.InternalError(err, "failed to hash token")
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
		return nil, bufcutil.InternalError(err, "db create failed")
	}
	if err := s.users.AuthUpsert(ctx, "email", strings.ToLower(user.Email), auth); err != nil {
		return nil, bufcutil.InternalError(err, "db authentication create failed")
	}

	log.Info("User Created")

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&corev1.UserChangeEvent{
		Current: s.toProtoUser(&user),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}
	if secret != "" {
		token, err := protoutil.SecureValueEncode(s.kms, secret)
		if err != nil {
			log.With(zap.Error(err)).Info("unable to create token")
		} else {
			payload := corev1.UserSecurityEvent{
				User:  s.toProtoUser(&user),
				Token: token,
			}

			if params.Status == corev1.UserStatus_USER_STATUS_REGISTERED {
				payload.Action = corev1.UserSecurity_USER_SECURITY_USER_REGISTER_TOKEN
				if _, err := s.pubsub.SecurityRegisterToken(ctx, connect.NewRequest(&payload)); err != nil {
					log.With(zap.Error(err)).Info("user security publish failed")
				}
			} else if params.Status == corev1.UserStatus_USER_STATUS_INVITED {
				payload.Action = corev1.UserSecurity_USER_SECURITY_USER_INVITE_TOKEN
				if _, err := s.pubsub.SecurityInviteToken(ctx, connect.NewRequest(&payload)); err != nil {
					log.With(zap.Error(err)).Info("user security publish failed")
				}
			}

		}
	}

	return connect.NewResponse(&corev1.UserServiceCreateResponse{User: s.toProtoUser(&user)}), nil
}

func (s *UserServer) Update(ctx context.Context, request *connect.Request[corev1.UserServiceUpdateRequest]) (*connect.Response[corev1.UserServiceUpdateResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("User Update")

	orig, err := s.users.GetById(ctx, params.UserId)
	if err != nil {
		log.With(zap.Error(err)).Error("Update/GetById failed")
		return nil, bufcutil.InternalError(err)
	}

	if orig == nil {
		return nil, twirp.NotFound.Error("User not found")
	}
	// Some basic validation
	if params.Password != nil || params.PasswordNew != nil {
		auth, err := s.users.AuthGet(ctx, "email", strings.ToLower(orig.Email))
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
			return nil, bufcutil.InternalError(err)
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

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&corev1.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	if params.PasswordNew != nil {
		if _, err := s.pubsub.SecurityPasswordChange(ctx, connect.NewRequest(&corev1.UserSecurityEvent{
			Action: corev1.UserSecurity_USER_SECURITY_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(&updated),
		})); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return connect.NewResponse(&corev1.UserServiceUpdateResponse{User: s.toProtoUser(&updated)}), nil
}

func (s *UserServer) ComparePassword(ctx context.Context, request *connect.Request[corev1.ComparePasswordRequest]) (*connect.Response[corev1.ComparePasswordResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("check password")

	auth, err := s.users.AuthGet(ctx, "email", strings.ToLower(params.Email))
	if err != nil {
		log.With(zap.Error(err)).Error("AuthGet failed")
		return nil, bufcutil.InternalError(err)
	}

	if auth == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	if !passwordCompare(auth.Password, params.Password) {
		return nil, bufcutil.InvalidArgumentError("password", "password mismatch")
	}

	return connect.NewResponse(&corev1.ComparePasswordResponse{
		UserId: auth.UserID,
	}), nil
}

func (s *UserServer) GetSettings(ctx context.Context, params *connect.Request[corev1.UserServiceGetSettingsRequest]) (*connect.Response[corev1.UserServiceGetSettingsResponse], error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.Msg.UserId))
	user, err := s.users.GetById(ctx, params.Msg.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("GetSettings/GetById failed")
		return nil, bufcutil.InternalError(err)
	}

	if user == nil {
		return nil, twirp.NotFound.Error("user ID not found")
	}

	return connect.NewResponse(&corev1.UserServiceGetSettingsResponse{Settings: s.toProtoSettings(user)}), nil
}

func (s *UserServer) SetSettings(ctx context.Context, request *connect.Request[corev1.UserServiceSetSettingsRequest]) (*connect.Response[corev1.UserServiceSetSettingsResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	orig, err := s.users.GetById(ctx, params.UserId)

	if err != nil {
		log.With(zap.Error(err)).Error("SetSettings/GetById failed")
		return nil, bufcutil.InternalError(err)
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
	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&corev1.UserChangeEvent{
		Current:  s.toProtoUser(&updated),
		Original: s.toProtoUser(orig),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return connect.NewResponse(&corev1.UserServiceSetSettingsResponse{Settings: s.toProtoSettings(&updated)}), nil
}

func (s *UserServer) getUserByVerification(ctx context.Context, params *corev1.Verification) (*UserAuth, error) {
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("getUserByVerification BEGIN")

	auth, err := s.users.AuthGet(ctx, "forgot", params.UserId)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	if auth == nil {
		log.Info("User not found")
		return nil, twirp.NotFound.Error("user ID not found")
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
func (s *UserServer) VerificationVerify(ctx context.Context, request *connect.Request[corev1.VerificationVerifyRequest]) (*connect.Response[corev1.VerificationVerifyResponse], error) {
	params := request.Msg.Verification
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
		return nil, bufcutil.InvalidArgumentError("token", "expired")
	}
	if ok, err := hmacCompare(user.ID, params.Token, user.EmailVerifyToken); err != nil {
		return nil, bufcutil.InternalError(err)
	} else if !ok {
		log.Info("Token mismatch")
		return nil, bufcutil.InvalidArgumentError("token", "bad token")
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

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&corev1.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}

	return connect.NewResponse(&corev1.VerificationVerifyResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) ForgotVerify(ctx context.Context, request *connect.Request[corev1.ForgotVerifyRequest]) (*connect.Response[corev1.ForgotVerifyResponse], error) {
	params := request.Msg.Verification
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

	return connect.NewResponse(&corev1.ForgotVerifyResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) ForgotUpdate(ctx context.Context, request *connect.Request[corev1.ForgotUpdateRequest]) (*connect.Response[corev1.ForgotUpdateResponse], error) {
	params := request.Msg.Verification
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
			return nil, bufcutil.InvalidArgumentError("password", "no match")
		}

		auth := UserAuth{
			UserID:   user.ID,
			Password: pass,
		}
		if err := s.users.AuthUpsert(ctx, "email", strings.ToLower(user.Email), auth); err != nil {
			return nil, bufcutil.InternalError(err)
		}
		err = s.users.AuthDelete(ctx, "forgot", user.ID, UserAuth{
			UserID: user.ID,
		})
		if err != nil {
			return nil, bufcutil.InternalError(err)
		}
	}

	// You could be "REGISTERED" or "INVITED"
	if user.Status != UserStatus_ACTIVE {
		update.Status = UserStatus_ACTIVE
		err := s.users.UpdateUser(ctx, &update)
		if err != nil {
			return nil, bufcutil.InternalError(err)
		}
	}

	if _, err := s.pubsub.UserChange(ctx, connect.NewRequest(&corev1.UserChangeEvent{
		Current:  s.toProtoUser(&update),
		Original: s.toProtoUser(user),
	})); err != nil {
		log.With(zap.Error(err)).Info("user entity publish failed")
	}
	if params.Password != "" {
		if _, err := s.pubsub.SecurityPasswordChange(ctx, connect.NewRequest(&corev1.UserSecurityEvent{
			Action: corev1.UserSecurity_USER_SECURITY_USER_PASSWORD_CHANGE,
			User:   s.toProtoUser(user),
		})); err != nil {
			log.With(zap.Error(err)).Info("user security publish failed")
		}
	}

	return connect.NewResponse(&corev1.ForgotUpdateResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) ForgotSend(ctx context.Context, request *connect.Request[corev1.ForgotSendRequest]) (*connect.Response[corev1.ForgotSendResponse], error) {
	params := request.Msg.FindBy
	log := logger.FromContext(ctx).With(zap.String("userId", params.UserId))
	log.Info("Forgot send")

	if params.Email == "" {
		return nil, bufcutil.InvalidArgumentError("email", "must provide email")
	}
	user, err := s.users.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	if user == nil {
		return nil, twirp.NotFound.Errorf("user not found email=%s", params.Email)
	}

	vExpires := time.Now().Add(time.Duration(24 * time.Hour))
	// use shortuuid rather than xid since it's less sequential
	vToken, secret, err := tokenEncrypt(shortuuid.New())
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	err = s.users.AuthUpsert(ctx, "forgot", user.ID, UserAuth{
		UserID:    user.ID,
		Password:  vToken,
		ExpiresAt: &vExpires,
	})
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	// TODO?
	// if err := s.publishUser(ctx, ENTITY_USER, user, &update); err != nil {
	// 	log.With(zap.Error(err)).Info("Publish failed")
	// }
	if token, err := protoutil.SecureValueEncode(s.kms, secret); err != nil {
		log.With(zap.Error(err)).Info("failed to encrypt token")
	} else if _, err := s.pubsub.SecurityForgotRequest(ctx, connect.NewRequest(&corev1.UserSecurityEvent{
		Action: corev1.UserSecurity_USER_SECURITY_USER_FORGOT_REQUEST,
		User:   s.toProtoUser(user),
		Token:  token,
	})); err != nil {
		log.With(zap.Error(err)).Info("user security publish failed")
	}

	return connect.NewResponse(&corev1.ForgotSendResponse{User: s.toProtoUser(user)}), nil
}

func (s *UserServer) AuthAssociate(ctx context.Context, params *connect.Request[corev1.AuthAssociateRequest]) (*connect.Response[corev1.AuthAssociateResponse], error) {
	auth := UserAuth{
		UserID: params.Msg.UserId,
	}

	if err := s.users.AuthUpsert(ctx, params.Msg.Auth.Provider, params.Msg.Auth.ProviderId, auth); err != nil {
		return nil, err
	}

	return connect.NewResponse(&corev1.AuthAssociateResponse{UserId: params.Msg.UserId}), nil
}

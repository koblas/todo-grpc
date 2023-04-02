package auth

import (
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/auth/v1"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"golang.org/x/net/context"
)

// Authenticate the user with email and password (aka Login)
func (s AuthenticationServer) Authenticate(ctx context.Context, paramsIn *connect.Request[apiv1.AuthenticateRequest]) (*connect.Response[apiv1.AuthenticateResponse], error) {
	params := paramsIn.Msg

	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Authenticate")

	email := strings.TrimSpace(params.Email)

	if email == "" {
		return nil, bufcutil.InvalidArgumentError("email", "Bad email or password (empty_field)")
	}

	// Avoid the PII leak by storing the email in redis in the clear
	attemptsKey := util.GetMD5Hash(email)

	if count, err := s.attempts.GetTries(ctx, "login", attemptsKey); count >= MAX_LOGIN_ATTEMPS {
		return nil, bufcutil.InvalidArgumentError("email", "Too many attempts (too_many_attemps)")
	} else if err != nil {
		log.With("error", err).Error("Authenticate/redis unable to fetch attempts keys")
	}

	userParam, err := s.userClient.ComparePassword(ctx, connect.NewRequest(&userv1.ComparePasswordRequest{
		Email:    params.Email,
		Password: params.Password,
	}))
	if err != nil {
		if err := s.attempts.Incr(ctx, "login", attemptsKey, LOGIN_LOCKOUT_MINUTES); err != nil {
			log.With("error", err).Error("Authenticate/redis unable to set attempts key")
		}

		log.With("error", err).Info("Password mismatch")

		return nil, bufcutil.InvalidArgumentError("email", "Bad email or password (bad_email_password)")
	} else {
		s.attempts.Reset(ctx, "login", attemptsKey)
	}

	token, err := s.returnToken(ctx, userParam.Msg.UserId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.AuthenticateResponse{Token: token}), nil
}

func (s AuthenticationServer) Register(ctx context.Context, paramsIn *connect.Request[apiv1.RegisterRequest]) (*connect.Response[apiv1.RegisterResponse], error) {
	params := paramsIn.Msg
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Register")

	if len(params.Password) < 8 {
		return nil, bufcutil.InvalidArgumentError("password", "Password too short must be 8 characters (password_short)")
	}

	user, err := s.userClient.Create(ctx, connect.NewRequest(&userv1.UserServiceCreateRequest{
		Status:   userv1.UserStatus_USER_STATUS_REGISTERED,
		Email:    params.Email,
		Password: params.Password,
		Name:     params.Name,
	}))
	if err != nil {
		log.With("error", err).Info("Unable to register new user")
		if connect.CodeOf(err) == connect.CodeAlreadyExists {
			return nil, bufcutil.InvalidArgumentError("email", "Unable to create account (duplicate_email)")
		}
		return nil, bufcutil.InternalError(errors.New("unexpected error"))
	}

	token, err := s.returnToken(ctx, user.Msg.User.Id)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.RegisterResponse{
		Token:   token,
		Created: true,
	}), nil
}

func (s AuthenticationServer) VerifyEmail(ctx context.Context, params *connect.Request[apiv1.VerifyEmailRequest]) (*connect.Response[apiv1.VerifyEmailResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("Verify register user")

	user, err := s.userClient.VerificationVerify(ctx, connect.NewRequest(&userv1.VerificationVerifyRequest{
		Verification: &userv1.Verification{
			UserId: params.Msg.UserId,
			Token:  params.Msg.Token,
		},
	}))
	if err != nil {
		log.With("error", err).Info("Recover Send")
	}
	if user == nil {
		return nil, bufcutil.InvalidArgumentError("token", "Unable to find verification (not_found)")
	}

	return connect.NewResponse(&apiv1.VerifyEmailResponse{}), nil
}

func (s AuthenticationServer) RecoverSend(ctx context.Context, params *connect.Request[apiv1.RecoverSendRequest]) (*connect.Response[apiv1.RecoverSendResponse], error) {
	email := strings.TrimSpace(params.Msg.Email)

	log := logger.FromContext(ctx).With("email", email)
	log.Info("Recover Send")

	attemptsKey := util.GetMD5Hash(email)
	if count, err := s.attempts.GetTries(ctx, "forgotX", attemptsKey); count >= MAX_LOGIN_ATTEMPS {
		log.Info("Too many attemps")
		return nil, bufcutil.InvalidArgumentError("email", "Too many attempts (too_many_attemps)")
	} else if err != nil {
		log.With("error", err).Error("RecoverSend/redis unable to fetch attempts keys")
	}

	user, err := s.userClient.ForgotSend(ctx, connect.NewRequest(&userv1.ForgotSendRequest{
		FindBy: &userv1.FindBy{
			Email: email,
		},
	}))
	if err != nil {
		log.With("error", err).Info("Recover Send")
		return nil, err
	} else {
		if err := s.attempts.Incr(ctx, "forgot", attemptsKey, LOGIN_LOCKOUT_MINUTES); err != nil {
			log.With("error", err).Error("Authenticate/redis unable to set forgot key")
		}
		if user != nil {
			log.Info("RecoverSend - found user")
		}
	}

	return connect.NewResponse(&apiv1.RecoverSendResponse{}), nil
}

func (s AuthenticationServer) RecoverVerify(ctx context.Context, params *connect.Request[apiv1.RecoverVerifyRequest]) (*connect.Response[apiv1.RecoverVerifyResponse], error) {
	log := logger.FromContext(ctx).With("user_id", params.Msg.UserId)
	log.Info("Recover Verify")

	user, err := s.userClient.ForgotVerify(ctx, connect.NewRequest(&userv1.ForgotVerifyRequest{
		Verification: &userv1.Verification{
			UserId: params.Msg.UserId,
			Token:  params.Msg.Token,
		},
	}))
	if err != nil {
		log.With("error", err).Info("Recover Verify")
	}
	if user == nil {
		return nil, bufcutil.InvalidArgumentError("token", "Token not found (not_found)")
	}

	return connect.NewResponse(&apiv1.RecoverVerifyResponse{}), nil
}

func (s AuthenticationServer) RecoverUpdate(ctx context.Context, params *connect.Request[apiv1.RecoverUpdateRequest]) (*connect.Response[apiv1.RecoverUpdateResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("Recover Update password")

	user, err := s.userClient.ForgotUpdate(ctx, connect.NewRequest(&userv1.ForgotUpdateRequest{
		Verification: &userv1.Verification{
			UserId:   params.Msg.UserId,
			Token:    params.Msg.Token,
			Password: params.Msg.Password,
		},
	}))
	if err != nil || user == nil {
		log.With("error", err).Info("Recover Update")

		return nil, bufcutil.InvalidArgumentError("token", "Token not found (not_found)")
	}

	token, err := s.returnToken(ctx, user.Msg.User.Id)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.RecoverUpdateResponse{Token: token}), nil
}

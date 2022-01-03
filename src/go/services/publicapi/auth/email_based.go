package auth

import (
	"strings"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/context"
)

// Authenticate the user with email and password (aka Login)
func (s AuthenticationServer) Authenticate(ctx context.Context, params *publicapi.LoginParams) (*publicapi.Token, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Authenticate")

	email := strings.TrimSpace(params.Email)

	if email == "" {
		return nil, twirp.InvalidArgumentError("email", "Bad email or password").WithMeta("email", "empty_field")
	}

	// Avoid the PII leak by storing the email in redis in the clear
	attemptsKey := util.GetMD5Hash(email)

	if count, err := s.attempts.GetTries(ctx, "login", attemptsKey); count >= MAX_LOGIN_ATTEMPS {
		return nil, twirp.InvalidArgumentError("email", "Too many attempts").WithMeta("email", "too_many_attemps")
	} else if err != nil {
		log.With("error", err).Error("Authenticate/redis unable to fetch attempts keys")
	}

	// user, err := s.userClient.FindBy(ctx, &core.FindParam{
	// 	Email: email,
	// })
	// if err != nil || user == nil {
	// 	if err != nil {
	// 		if e, ok := err.(twirp.Error); ok && e.Code() == twirp.NotFound {
	// 			log.Info("User not found")

	// 			return nil, twirp.InvalidArgumentError("email", "Bad email or password").WithMeta("email", "empty_field")
	// 		} else {
	// 			log.With("error", err).With("twirpOk", ok).Error("Unable to call user service")

	// 			return nil, twirp.InternalError("user - service failure")
	// 		}
	// 	}
	// 	if user == nil {
	// 		return nil, errors.New("unexpected nil user")
	// 	}
	// 	return nil, err
	// }

	// log = log.With("user_id", user.Id)
	userParam, err := s.userClient.ComparePassword(ctx, &core.AuthenticateParam{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		if err := s.attempts.Incr(ctx, "login", attemptsKey, LOGIN_LOCKOUT_MINUTES); err != nil {
			log.With("error", err).Error("Authenticate/redis unable to set attempts key")
		}

		log.With("error", err).Info("Password mismatch")

		return nil, twirp.InvalidArgumentError("email", "Bad email or password").WithMeta("email", "empty_field")
	} else {
		s.attempts.Reset(ctx, "login", attemptsKey)
	}

	return s.returnToken(ctx, userParam.UserId)
}

func (s AuthenticationServer) Register(ctx context.Context, params *publicapi.RegisterParams) (*publicapi.TokenRegister, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Register")

	if len(params.Password) < 8 {
		return nil, twirp.InvalidArgumentError("password", "password too short").WithMeta("password", "Password must be 8 characters")
	}

	user, err := s.userClient.Create(ctx, &core.CreateParam{
		Status:   core.UserStatus_REGISTERED,
		Email:    params.Email,
		Password: params.Password,
		Name:     params.Name,
	})
	if err != nil {
		log.With("error", err).Info("Unable to register new user")
		return nil, twirp.InvalidArgumentError("email", "Unable to create").WithMeta("email", "Duplicate Email")
	}

	token, err := s.returnToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &publicapi.TokenRegister{
		Token:   token,
		Created: true,
	}, nil
}

func (s AuthenticationServer) VerifyEmail(ctx context.Context, params *publicapi.ConfirmParams) (*publicapi.Success, error) {
	log := logger.FromContext(ctx)
	log.Info("Verify register user")

	user, err := s.userClient.VerificationVerify(ctx, &core.VerificationParam{
		UserId: params.UserId,
		Token:  params.Token,
	})
	if err != nil {
		log.With("error", err).Info("Recover Send")
	}
	if user == nil {
		return nil, twirp.InvalidArgumentError("token", "not found").WithMeta("token", "Not Found")
	}

	return &publicapi.Success{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverSend(ctx context.Context, params *publicapi.RecoverySendParams) (*publicapi.Success, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Recover Send")

	email := strings.TrimSpace(params.Email)

	attemptsKey := util.GetMD5Hash(email)
	if count, err := s.attempts.GetTries(ctx, "forgot", attemptsKey); count >= MAX_LOGIN_ATTEMPS {
		log.Info("Too many attemps")
		return nil, twirp.InvalidArgumentError("email", "Too many attemps").WithMeta("email", "Too many attemps")
	} else if err != nil {
		log.With("error", err).Error("RecoverSend/redis unable to fetch attempts keys")
	}

	user, err := s.userClient.ForgotSend(ctx, &core.FindParam{
		Email: params.Email,
	})
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

	return &publicapi.Success{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverVerify(ctx context.Context, params *publicapi.RecoveryUpdateParams) (*publicapi.Success, error) {
	log := logger.FromContext(ctx).With("user_id", params.UserId)
	log.Info("Recover Verify")

	user, err := s.userClient.ForgotVerify(ctx, &core.VerificationParam{
		UserId: params.UserId,
		Token:  params.Token,
	})
	if err != nil {
		log.With("error", err).Info("Recover Verify")
	}
	if user == nil {
		return nil, twirp.InvalidArgumentError("token", "not found").WithMeta("token", "Not found")
	}

	return &publicapi.Success{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverUpdate(ctx context.Context, params *publicapi.RecoveryUpdateParams) (*publicapi.Token, error) {
	log := logger.FromContext(ctx)
	log.Info("Recover Update password")

	user, err := s.userClient.ForgotUpdate(ctx, &core.VerificationParam{
		UserId:   params.UserId,
		Token:    params.Token,
		Password: params.Password,
	})
	if err != nil || user == nil {
		log.With("error", err).Info("Recover Update")

		return nil, twirp.InvalidArgumentError("token", "not found").WithMeta("token", "Not found")
	}

	return s.returnToken(ctx, user.Id)
}

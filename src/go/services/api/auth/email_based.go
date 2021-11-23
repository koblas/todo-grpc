package auth

import (
	"errors"
	"strings"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authenticate the user with email and password (aka Login)
func (s AuthenticationServer) Authenticate(ctx context.Context, params *publicapi.LoginParams) (*publicapi.Token, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Authenticate")

	email := strings.TrimSpace(params.Email)

	if email == "" {
		if s, err := status.New(codes.InvalidArgument, "Bad email or password").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "empty_field",
			},
		); err != nil {
			panic(err)
		} else {
			return nil, s.Err()
		}
	}

	// Avoid the PII leak by storing the email in redis in the clear
	attemptsKey := util.GetMD5Hash(email)

	if count, err := s.attempts.GetTries(ctx, "login", attemptsKey); count >= MAX_LOGIN_ATTEMPS {
		if s, err := status.New(codes.InvalidArgument, "Too many attempts").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "Too many attempts to login",
			},
		); err != nil {
			panic(err)
		} else {
			return nil, s.Err()
		}

	} else {
		log.With("error", err).Error("Authenticate/redis unable to fetch attempts keys")
	}

	user, err := s.userClient.FindBy(ctx, &core.FindParam{
		Email: email,
	})
	if err != nil || user == nil {
		if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
			log.Info("User not found")
			if s, err := status.New(codes.InvalidArgument, "Bad email or password").WithDetails(
				&errdetails.BadRequest_FieldViolation{
					Field:       "Email",
					Description: "Email not associated with account",
				},
			); err != nil {
				panic(err)
			} else {
				return nil, s.Err()
			}
		}
		if user == nil {
			return nil, errors.New("unexpected nil user")
		}
		return nil, err
	}

	log = log.With("user_id", user.Id)
	user, err = s.userClient.ComparePassword(ctx, &core.AuthenticateParam{
		UserId:   user.Id,
		Password: params.Password,
	})
	if err != nil {
		if err := s.attempts.Incr(ctx, "login", attemptsKey, LOGIN_LOCKOUT_MINUTES); err != nil {
			log.With("error", err).Error("Authenticate/redis unable to set attempts key")
		}

		log.With("error", err).Info("Password mismatch")
		s, err := status.New(codes.InvalidArgument, "Bad email or password").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Password",
				Description: "Password doesn't match",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	} else {
		s.attempts.Reset(ctx, "login", attemptsKey)
	}

	return s.returnToken(ctx, user)
}

func (s AuthenticationServer) Register(ctx context.Context, params *publicapi.RegisterParams) (*publicapi.TokenRegister, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Register")

	user, err := s.userClient.Create(ctx, &core.CreateParam{
		Status:   core.UserStatus_REGISTERED,
		Email:    params.Email,
		Password: params.Password,
		Name:     params.Name,
	})
	if err != nil {
		s, err := status.New(codes.InvalidArgument, "Unable to create user").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "Duplicate email",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	}

	token, err := s.returnToken(ctx, user)
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
		s, err := status.New(codes.InvalidArgument, "token not found").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Token",
				Description: "not found",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
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
		if s, err := status.New(codes.InvalidArgument, "Too many attempts").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "Too many attempts to recover email",
			},
		); err != nil {
			panic(err)
		} else {
			return nil, s.Err()
		}
	} else {
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

	user, err := s.userClient.VerificationVerify(ctx, &core.VerificationParam{
		UserId: params.UserId,
		Token:  params.Token,
	})
	if err != nil {
		log.With("error", err).Info("Recover Send")
	}
	if user == nil {
		s, err := status.New(codes.InvalidArgument, "token not found").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Token",
				Description: "not found",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	}

	return &publicapi.Success{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverUpdate(ctx context.Context, params *publicapi.RecoveryUpdateParams) (*publicapi.Token, error) {
	log := logger.FromContext(ctx)
	log.Info("Recover Update password")

	user, err := s.userClient.VerificationUpdate(ctx, &core.VerificationParam{
		UserId:   params.UserId,
		Token:    params.Token,
		Password: params.Password,
	})
	if err != nil || user == nil {
		log.With("error", err).Info("Recover Update")

		s, err := status.New(codes.InvalidArgument, "token not found").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Password",
				Description: "Unable to update password",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	}

	return s.returnToken(ctx, user)
}

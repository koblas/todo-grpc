package auth

import (
	"log"
	"os"
	"time"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents the gRPC server
type AuthenticationServer struct {
	publicapi.UnimplementedAuthenticationServiceServer

	jwtMaker   tokenmanager.Maker
	userClient core.UserServiceClient
}

func NewAuthenticationServer(userClient core.UserServiceClient) AuthenticationServer {
	maker, err := tokenmanager.NewJWTMaker(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	return AuthenticationServer{
		userClient: userClient,
		jwtMaker:   maker,
	}
}

func (s AuthenticationServer) returnToken(ctx context.Context, user *core.User) (*publicapi.Token, error) {
	// TODO: This is an authentication event that should be pushed onto
	//  the messagebus

	bearer, err := s.jwtMaker.CreateToken(user.Id, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.Token{
		AccessToken: bearer,
		TokenType:   "Bearer",
		ExpiresIn:   24 * 3600,
	}, nil
}

func (s AuthenticationServer) Authenticate(ctx context.Context, params *publicapi.LoginParams) (*publicapi.Token, error) {
	log := logger.FromContext(ctx).With("email", params.Email)
	log.Info("Authenticate")

	user, err := s.userClient.FindBy(ctx, &core.FindParam{
		Email: params.Email,
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		log.Info("User not found")
		s, err := status.New(codes.InvalidArgument, "Bad email or password").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "Bad email or password",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	}
	user, err = s.userClient.ComparePassword(ctx, &core.AuthenticateParam{
		UserId:   user.Id,
		Password: params.Password,
	})
	if err != nil {
		log.With("error", err).Info("Password mismatch")
		s, err := status.New(codes.InvalidArgument, "Bad email or password").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Email",
				Description: "Bad email or password",
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, s.Err()
	}

	return s.returnToken(ctx, user)
}

func (s AuthenticationServer) Register(ctx context.Context, params *publicapi.RegisterParams) (*publicapi.Token, error) {
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

	return s.returnToken(ctx, user)
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

	user, err := s.userClient.ForgotSend(ctx, &core.FindParam{
		Email: params.Email,
	})
	if err != nil {
		log.With("error", err).Info("Recover Send")
	}
	if user != nil {
		log.Info("RecoverSend - found user")
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

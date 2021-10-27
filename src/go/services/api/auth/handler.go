package auth

import (
	"log"
	"os"
	"time"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"golang.org/x/net/context"
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

func (s AuthenticationServer) Login(ctx context.Context, params *publicapi.LoginParams) (*publicapi.TokenEither, error) {
	log.Printf("Received login %s", params.Email)

	either, err := s.userClient.FindBy(ctx, &core.FindParam{
		Email: params.Email,
	})
	if err != nil {
		return nil, err
	}
	if either.User != nil {
		return &publicapi.TokenEither{
			Errors: []*publicapi.ValidationError{
				{
					Field:   "email",
					Message: "Bad email or password",
				},
			},
		}, nil
	}
	user := either.User
	either, err = s.userClient.ComparePassword(ctx, &core.AuthenticateParam{
		UserId:   user.Id,
		Password: params.Password,
	})
	if err != nil {
		return nil, err
	}
	if either.User != nil {
		return &publicapi.TokenEither{
			Errors: []*publicapi.ValidationError{
				{
					Field:   "email",
					Message: "Bad email or password",
				},
			},
		}, nil
	}
	user = either.User

	bearer, err := s.jwtMaker.CreateToken(user.Id, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.TokenEither{
		Token: &publicapi.Token{
			AccessToken: bearer,
			TokenType:   "Bearer",
			ExpiresIn:   24 * 3600,
		},
	}, nil
}

func (s AuthenticationServer) Register(ctx context.Context, params *publicapi.RegisterParams) (*publicapi.TokenEither, error) {
	log.Printf("Received register %s", params.Email)

	either, err := s.userClient.Create(ctx, &core.CreateParam{
		Email:    params.Email,
		Password: params.Password,
		Name:     params.Name,
	})
	if err != nil {
		return nil, err
	}
	if either.Error != nil {
		return &publicapi.TokenEither{
			Errors: []*publicapi.ValidationError{
				{
					Field:   "Email",
					Message: either.Error.Message,
				},
			},
		}, nil
	}
	user := either.User

	bearer, err := s.jwtMaker.CreateToken(user.Id, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.TokenEither{
		Token: &publicapi.Token{
			AccessToken: bearer,
			TokenType:   "Bearer",
			ExpiresIn:   24 * 3600,
		},
	}, nil
}

func (s AuthenticationServer) VerifyEmail(ctx context.Context, params *publicapi.ConfirmParams) (*publicapi.TokenEither, error) {
	// TODO: Not Token
	log.Printf("Received verify %s", params.Token)

	bearer, err := s.jwtMaker.CreateToken(params.Token, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.TokenEither{
		Token: &publicapi.Token{
			AccessToken: bearer,
			TokenType:   "Bearer",
			ExpiresIn:   24 * 3600,
		},
	}, nil
}

func (s AuthenticationServer) RecoverSend(ctx context.Context, params *publicapi.RecoveryParams) (*publicapi.SuccessEither, error) {
	log.Fatal("RecoverVerify Not Implemented")
	// TODO: Not Token
	log.Printf("Received recovoery %s", params.Email)

	return &publicapi.SuccessEither{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverVerify(ctx context.Context, params *publicapi.RecoveryParams) (*publicapi.SuccessEither, error) {
	log.Fatal("RecoverVerify Not Implemented")
	// TODO: Not Token
	log.Printf("Received recovoery %s", params.Email)

	return &publicapi.SuccessEither{
		Success: true,
	}, nil
}

func (s AuthenticationServer) RecoverUpdate(ctx context.Context, params *publicapi.RecoveryParams) (*publicapi.TokenEither, error) {
	log.Fatal("RecoverUpdate Not Implemented")
	// TODO: Not Token
	log.Printf("Received update %s", params.Token)

	bearer, err := s.jwtMaker.CreateToken(params.Token, time.Hour*24*365)
	if err != nil {
		return nil, err
	}

	return &publicapi.TokenEither{
		Token: &publicapi.Token{
			AccessToken: bearer,
			TokenType:   "Bearer",
			ExpiresIn:   24 * 3600,
		},
	}, nil
}

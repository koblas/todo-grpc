package auth

import (
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

// func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *apipbv1.OauthAssociateRequest) (*apipbv1.Oauth, error) {
// return nil, nil
// }

func (s AuthenticationServer) OauthLogin(ctx context.Context, params *apipbv1.OauthLoginRequest) (*apipbv1.OauthLoginResponse, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, &corepbv1.AuthUserServiceUpsertUserRequest{
		Oauth: &corepbv1.AuthOauthParams{
			Provider: params.Provider,
			Code:     params.Code,
		},
		State:       params.State,
		RedirectUrl: params.RedirectUrl,
	})

	if err != nil {
		log.Info("Failed to call oauthClient.UpsertUser")
		return nil, err
	}

	user, err := s.userClient.FindBy(ctx, &corepbv1.FindByRequest{
		FindBy: &corepbv1.FindBy{
			UserId: result.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user.User.Id)
	if err != nil {
		return nil, err
	}
	return &apipbv1.OauthLoginResponse{
		Token:   token,
		Created: result.Created,
	}, nil
}

func (s AuthenticationServer) OauthUrl(ctx context.Context, params *apipbv1.OauthUrlRequest) (*apipbv1.OauthUrlResponse, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthUrl")

	// TODO -- basic validation on parameters

	result, err := s.oauthClient.GetAuthUrl(ctx, &corepbv1.AuthUserServiceGetAuthUrlRequest{
		Provider:    params.Provider,
		RedirectUrl: params.RedirectUrl,
	})
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return &apipbv1.OauthUrlResponse{
		Url: result.GetUrl(),
	}, nil
}

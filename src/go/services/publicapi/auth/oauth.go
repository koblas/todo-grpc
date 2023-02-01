package auth

import (
	apipbv1 "github.com/koblas/grpc-todo/gen/apipb/v1"
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *apipbv1.OauthAssociateParams) (*apipbv1.Success, error) {
	return nil, nil
}

func (s AuthenticationServer) OauthLogin(ctx context.Context, params *apipbv1.OauthAssociateParams) (*apipbv1.TokenRegister, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, &corepbv1.AuthUserUpsertParams{
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

	user, err := s.userClient.FindBy(ctx, &corepbv1.UserFindParam{
		UserId: result.UserId,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &apipbv1.TokenRegister{
		Token:   token,
		Created: result.Created,
	}, nil
}

func (s AuthenticationServer) OauthUrl(ctx context.Context, params *apipbv1.OauthUrlParams) (*apipbv1.OauthUrlResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthUrl")

	// TODO -- basic validation on parameters

	result, err := s.oauthClient.GetAuthURL(ctx, &corepbv1.AuthOAuthGetUrlParams{
		Provider:    params.Provider,
		RedirectUrl: params.RedirectUrl,
	})
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return &apipbv1.OauthUrlResult{
		Url: result.GetUrl(),
	}, nil
}

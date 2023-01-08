package auth

import (
	"github.com/koblas/grpc-todo/gen/apipb"
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *apipb.OauthAssociateParams) (*apipb.Success, error) {
	return nil, nil
}

func (s AuthenticationServer) OauthLogin(ctx context.Context, params *apipb.OauthAssociateParams) (*apipb.TokenRegister, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, &corepb.AuthUserUpsertParams{
		Oauth: &corepb.AuthOauthParams{
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

	user, err := s.userClient.FindBy(ctx, &corepb.UserFindParam{
		UserId: result.UserId,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &apipb.TokenRegister{
		Token:   token,
		Created: result.Created,
	}, nil
}

func (s AuthenticationServer) OauthUrl(ctx context.Context, params *apipb.OauthUrlParams) (*apipb.OauthUrlResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthUrl")

	// TODO -- basic validation on parameters

	result, err := s.oauthClient.GetAuthURL(ctx, &corepb.AuthOAuthGetUrlParams{
		Provider:    params.Provider,
		RedirectUrl: params.RedirectUrl,
	})
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return &apipb.OauthUrlResult{
		Url: result.GetUrl(),
	}, nil
}

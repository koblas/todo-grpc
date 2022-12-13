package auth

import (
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	"github.com/koblas/grpc-todo/twpb/publicapi"
	"golang.org/x/net/context"
)

func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *publicapi.OauthAssociateParams) (*publicapi.Success, error) {
	return nil, nil
}

func (s AuthenticationServer) OauthLogin(ctx context.Context, params *publicapi.OauthAssociateParams) (*publicapi.TokenRegister, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, &core.AuthUserUpsertParams{
		Oauth: &core.AuthOauthParams{
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

	user, err := s.userClient.FindBy(ctx, &core.UserFindParam{
		UserId: result.UserId,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &publicapi.TokenRegister{
		Token:   token,
		Created: result.Created,
	}, nil
}

func (s AuthenticationServer) OauthUrl(ctx context.Context, params *publicapi.OauthUrlParams) (*publicapi.OauthUrlResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthUrl")

	// TODO -- basic validation on parameters

	result, err := s.oauthClient.GetAuthURL(ctx, &core.AuthOAuthGetUrlParams{
		Provider:    params.Provider,
		RedirectUrl: params.RedirectUrl,
	})
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return &publicapi.OauthUrlResult{
		Url: result.GetUrl(),
	}, nil
}

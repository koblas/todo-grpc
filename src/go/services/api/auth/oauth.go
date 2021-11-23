package auth

import (
	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/genpb/publicapi"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *publicapi.OauthAssociateParams) (*publicapi.Success, error) {
	return nil, nil
}

func (s AuthenticationServer) OauthLogin(ctx context.Context, params *publicapi.OauthAssociateParams) (*publicapi.TokenRegister, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, &core.OauthUserUpsertParams{
		Provider:    params.Provider,
		Code:        params.Code,
		State:       params.State,
		RedirectUrl: params.RedirectUrl,
	})

	if err != nil {
		log.Info("Failed to call oauthClient.UpsertUser")
		return nil, err
	}

	user, err := s.userClient.FindBy(ctx, &core.FindParam{
		UserId: result.UserId,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user)
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

	result, err := s.oauthClient.GetAuthURL(ctx, &core.OauthUserGetUrlParams{
		Provider:    params.Provider,
		RedirectUrl: params.RedirectUrl,
		State:       params.State,
	})
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return &publicapi.OauthUrlResult{
		Url: result.GetUrl(),
	}, nil
}

package auth

import (
	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"golang.org/x/net/context"
)

// func (s AuthenticationServer) OauthAssociate(ctx context.Context, params *apiv1.OauthAssociateRequest) (*apiv1.Oauth, error) {
// return nil, nil
// }

func (s AuthenticationServer) OauthLogin(ctx context.Context, paramsIn *connect.Request[apiv1.OauthLoginRequest]) (*connect.Response[apiv1.OauthLoginResponse], error) {
	params := paramsIn.Msg
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("OauthLogin")

	result, err := s.oauthClient.UpsertUser(ctx, connect.NewRequest(&corev1.AuthUserServiceUpsertUserRequest{
		Oauth: &corev1.AuthOauthParams{
			Provider: params.Provider,
			Code:     params.Code,
		},
		State:       params.State,
		RedirectUrl: params.RedirectUrl,
	}))

	if err != nil {
		log.Info("Failed to call oauthClient.UpsertUser")
		return nil, err
	}

	user, err := s.userClient.FindBy(ctx, connect.NewRequest(&corev1.FindByRequest{
		FindBy: &corev1.FindBy{
			UserId: result.Msg.UserId,
		},
	}))
	if err != nil {
		return nil, err
	}

	token, err := s.returnToken(ctx, user.Msg.User.Id)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.OauthLoginResponse{
		Token:   token,
		Created: result.Msg.Created,
	}), nil
}

func (s AuthenticationServer) OauthUrl(ctx context.Context, params *connect.Request[apiv1.OauthUrlRequest]) (*connect.Response[apiv1.OauthUrlResponse], error) {
	log := logger.FromContext(ctx).With("provider", params.Msg.Provider)
	log.Info("OauthUrl")

	// TODO -- basic validation on parameters

	result, err := s.oauthClient.GetAuthUrl(ctx, connect.NewRequest(&corev1.AuthUserServiceGetAuthUrlRequest{
		Provider:    params.Msg.Provider,
		RedirectUrl: params.Msg.RedirectUrl,
	}))
	if err != nil {
		log.Info("Failed to call oauthClient.GetAuthURL")
		return nil, err
	}

	return connect.NewResponse(&apiv1.OauthUrlResponse{
		Url: result.Msg.GetUrl(),
	}), nil
}

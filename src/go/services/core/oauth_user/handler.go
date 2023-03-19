package user

import (
	"log"
	"time"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/pkg/util"
	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type OauthUserServer struct {
	kms      key_manager.Encoder
	user     corev1connect.UserServiceClient
	jwtMaker tokenmanager.Maker
	smanager oauth_provider.SecretManager
}

type Option func(*OauthUserServer)

func WithUserService(client corev1connect.UserServiceClient) Option {
	return func(cfg *OauthUserServer) {
		cfg.user = client
	}
}

func WithSecretManager(client oauth_provider.SecretManager) Option {
	return func(cfg *OauthUserServer) {
		cfg.smanager = client
	}
}

func NewOauthUserServer(config Config, opts ...Option) *OauthUserServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)

	if err != nil {
		log.Fatal(err)
	}

	svr := OauthUserServer{
		kms:      key_manager.NewSecureClear(),
		jwtMaker: maker,
		// user:     user,
		// pubsub:   producer,
		// smanager: smanager,
	}

	for _, opt := range opts {
		opt(&svr)
	}

	return &svr
}

func (svc *OauthUserServer) GetAuthUrl(ctx context.Context, request *connect.Request[corev1.AuthUserServiceGetAuthUrlRequest]) (*connect.Response[corev1.AuthUserServiceGetAuthUrlResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("Calling GetAuthURL")
	oprovider, err := oauth_provider.GetOAuthProvider(params.GetProvider(), svc.smanager, log)
	if err != nil {
		log.With(zap.Error(err)).Info("failed to get provider")
		return nil, bufcutil.InternalError(err)
	}

	// Build a "STATE" value
	value, err := util.GenerateRandomString(20)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}
	state, err := svc.jwtMaker.CreateToken(value, time.Minute*10)
	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	url := oprovider.BuildRedirect(ctx, params.RedirectUrl, state)

	return connect.NewResponse(&corev1.AuthUserServiceGetAuthUrlResponse{Url: url}), nil
}

func (svc *OauthUserServer) UpsertUser(ctx context.Context, request *connect.Request[corev1.AuthUserServiceUpsertUserRequest]) (*connect.Response[corev1.AuthUserServiceUpsertUserResponse], error) {
	params := request.Msg
	log := logger.FromContext(ctx).With("provider", params.Oauth.Provider)
	log.Info("Calling UpsertUser")
	oprovider, err := oauth_provider.GetOAuthProvider(params.Oauth.Provider, svc.smanager, log)

	if err != nil {
		log.With(zap.Error(err)).Info("failed to get provider")
		return nil, bufcutil.InternalError(err)
	}

	// Verify the state matches
	_, err = svc.jwtMaker.VerifyToken(params.State)
	if err != nil {
		log.With("state", params.State).Info("Falied JWT validation")
		return nil, bufcutil.InvalidArgumentError("state", "state does not validate")
	}

	tokenResult, err := oprovider.GetAccessToken(ctx, params.Oauth.Code, params.RedirectUrl)
	if err != nil {
		log.With("error", err).Info("Failed to get access token")
		return nil, bufcutil.InternalError(err)
	}
	log.Info("Getting OAuth User information")

	info, err := oprovider.GetInfo(ctx, tokenResult)
	if err != nil {
		log.With("error", err).Info("Failed to get user info")
		return nil, bufcutil.InternalError(err)
	}

	if info.Id == "" {
		log.Info("Unable to get ID from provider")
		return nil, bufcutil.InternalError(err, "Unable to get ID from provider")
	}

	findBy, err := svc.user.FindBy(ctx, connect.NewRequest(&corev1.FindByRequest{
		FindBy: &corev1.FindBy{
			Auth: &corev1.AuthInfo{
				Provider:   params.Oauth.Provider,
				ProviderId: info.Id,
			},
		},
	}))
	if err != nil {
		if connect.CodeOf(err) != connect.CodeNotFound {
			return nil, bufcutil.InternalError(err)
		}
	}

	if findBy != nil && findBy.Msg != nil {
		return connect.NewResponse(&corev1.AuthUserServiceUpsertUserResponse{UserId: findBy.Msg.User.Id, Created: false}), nil
	}

	if info.Email == "" {
		log.With("error", err).Info("Failed to get user email")
		return nil, bufcutil.InvalidArgumentError("email", "provider didn't send email address")
	}

	findBy, err = svc.user.FindBy(ctx, connect.NewRequest(&corev1.FindByRequest{
		FindBy: &corev1.FindBy{
			Email: info.Email,
		},
	}))
	if err != nil {
		if connect.CodeOf(err) != connect.CodeNotFound {
			log.With("error", err).Info("Failed to lookup user")
			return nil, bufcutil.InternalError(err)
		}
	}

	userId := ""
	created := false
	if findBy == nil || findBy.Msg.User.Id == "" {
		log.With("email", info.Email).Info("Creating new user")
		created = true
		newUser, err := svc.user.Create(ctx, connect.NewRequest(&corev1.UserServiceCreateRequest{
			Email: info.Email,
			Name:  info.Name,
			// TODO - create as "ACTIVE" since we "know" the email is good
		}))
		if err != nil {
			log.With("error", err).Info("Unable to create user")
			return nil, bufcutil.InternalError(err)
		}
		userId = newUser.Msg.User.Id
	} else {
		userId = findBy.Msg.User.Id
		log.With("email", findBy.Msg.User.Email).Info("User already exists, associating")
	}

	// Now associate the OAuth token and the UserId
	// TODO - save the token!
	_, err = svc.user.AuthAssociate(ctx, connect.NewRequest(&corev1.AuthAssociateRequest{
		UserId: userId,
		Auth: &corev1.AuthInfo{
			Provider:   params.Oauth.Provider,
			ProviderId: info.Id,
		},
	}))
	if err != nil {
		log.With("error", err).Info("Unable to associate")
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&corev1.AuthUserServiceUpsertUserResponse{UserId: userId, Created: created}), nil
}

func (s *OauthUserServer) ListAssociations(ctx context.Context, params *connect.Request[corev1.AuthUserServiceListAssociationsRequest]) (*connect.Response[corev1.AuthUserServiceListAssociationsResponse], error) {
	return connect.NewResponse(&corev1.AuthUserServiceListAssociationsResponse{}), nil
}

func (s *OauthUserServer) RemoveAssociation(ctx context.Context, params *connect.Request[corev1.AuthUserServiceRemoveAssociationRequest]) (*connect.Response[corev1.AuthUserServiceRemoveAssociationResponse], error) {
	return connect.NewResponse(&corev1.AuthUserServiceRemoveAssociationResponse{}), nil
}

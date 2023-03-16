package user

import (
	"log"
	"time"

	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/pkg/util"
	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	"github.com/twitchtv/twirp"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type OauthUserServer struct {
	kms      key_manager.Encoder
	user     corepbv1.UserService
	jwtMaker tokenmanager.Maker
	smanager oauth_provider.SecretManager
}

type Option func(*OauthUserServer)

func WithUserService(client corepbv1.UserService) Option {
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

func (svc *OauthUserServer) GetAuthUrl(ctx context.Context, params *corepbv1.AuthUserServiceGetAuthUrlRequest) (*corepbv1.AuthUserServiceGetAuthUrlResponse, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("Calling GetAuthURL")
	oprovider, err := oauth_provider.GetOAuthProvider(params.GetProvider(), svc.smanager, log)
	if err != nil {
		log.With(zap.Error(err)).Info("failed to get provider")
		return nil, twirp.InternalErrorWith(err)
	}

	// Build a "STATE" value
	value, err := util.GenerateRandomString(20)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	state, err := svc.jwtMaker.CreateToken(value, time.Minute*10)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	url := oprovider.BuildRedirect(ctx, params.RedirectUrl, state)

	return &corepbv1.AuthUserServiceGetAuthUrlResponse{Url: url}, nil
}

func (svc *OauthUserServer) UpsertUser(ctx context.Context, params *corepbv1.AuthUserServiceUpsertUserRequest) (*corepbv1.AuthUserServiceUpsertUserResponse, error) {
	log := logger.FromContext(ctx).With("provider", params.Oauth.Provider)
	log.Info("Calling UpsertUser")
	oprovider, err := oauth_provider.GetOAuthProvider(params.Oauth.Provider, svc.smanager, log)

	if err != nil {
		log.With(zap.Error(err)).Info("failed to get provider")
		return nil, twirp.InternalErrorWith(err)
	}

	// Verify the state matches
	_, err = svc.jwtMaker.VerifyToken(params.State)
	if err != nil {
		log.With("state", params.State).Info("Falied JWT validation")
		// TODO - check it's invalid vs library failure
		return nil, twirp.InvalidArgumentError("state", "state does not validate")
	}

	tokenResult, err := oprovider.GetAccessToken(ctx, params.Oauth.Code, params.RedirectUrl)
	if err != nil {
		log.With("error", err).Info("Failed to get access token")
		return nil, err
	}
	log.Info("Getting OAuth User information")

	info, err := oprovider.GetInfo(ctx, tokenResult)
	if err != nil {
		log.With("error", err).Info("Failed to get user info")
		return nil, twirp.InternalErrorWith(err)
	}

	if info.Id == "" {
		log.Info("Unable to get ID from provider")
		return nil, twirp.InternalError("Unable to get ID from provider")
	}

	findBy, err := svc.user.FindBy(ctx, &corepbv1.FindByRequest{
		FindBy: &corepbv1.FindBy{
			Auth: &corepbv1.AuthInfo{
				Provider:   params.Oauth.Provider,
				ProviderId: info.Id,
			},
		},
	})
	if err != nil {
		if e, ok := err.(twirp.Error); !ok || e.Code() != twirp.NotFound {
			log.With("error", err).Info("Failed to get oauth user")
			return nil, twirp.InternalErrorWith(err)
		}
	}

	if findBy != nil {
		return &corepbv1.AuthUserServiceUpsertUserResponse{UserId: findBy.User.Id, Created: false}, nil
	}

	if info.Email == "" {
		log.With("error", err).Info("Failed to get user email")
		return nil, twirp.InvalidArgumentError("email", "provider didn't send email address")
	}

	findBy, err = svc.user.FindBy(ctx, &corepbv1.FindByRequest{
		FindBy: &corepbv1.FindBy{
			Email: info.Email,
		},
	})
	if err != nil {
		if e, ok := err.(twirp.Error); !ok || e.Code() != twirp.NotFound {
			log.With("error", err).Info("Failed to lookup user")
			return nil, twirp.InternalErrorWith(err)
		}
	}

	userId := ""
	created := false
	if findBy == nil || findBy.User.Id == "" {
		log.With("email", info.Email).Info("Creating new user")
		created = true
		newUser, err := svc.user.Create(ctx, &corepbv1.UserServiceCreateRequest{
			Email: info.Email,
			Name:  info.Name,
			// TODO - create as "ACTIVE" since we "know" the email is good
		})
		if err != nil {
			log.With("error", err).Info("Unable to create user")
			return nil, twirp.InternalErrorWith(err)
		}
		userId = newUser.User.Id
	} else {
		userId = findBy.User.Id
		log.With("email", findBy.User.Email).Info("User already exists, associating")
	}

	// Now associate the OAuth token and the UserId
	// TODO - save the token!
	_, err = svc.user.AuthAssociate(ctx, &corepbv1.AuthAssociateRequest{
		UserId: userId,
		Auth: &corepbv1.AuthInfo{
			Provider:   params.Oauth.Provider,
			ProviderId: info.Id,
		},
	})
	if err != nil {
		log.With("error", err).Info("Unable to associate")
		return nil, twirp.InternalErrorWith(err)
	}

	return &corepbv1.AuthUserServiceUpsertUserResponse{UserId: userId, Created: created}, nil
}

func (s *OauthUserServer) ListAssociations(ctx context.Context, params *corepbv1.AuthUserServiceListAssociationsRequest) (*corepbv1.AuthUserServiceListAssociationsResponse, error) {
	return &corepbv1.AuthUserServiceListAssociationsResponse{}, nil
}

func (s *OauthUserServer) RemoveAssociation(ctx context.Context, params *corepbv1.AuthUserServiceRemoveAssociationRequest) (*corepbv1.AuthUserServiceRemoveAssociationResponse, error) {
	return &corepbv1.AuthUserServiceRemoveAssociationResponse{}, nil
}

package user

import (
	"log"
	"time"

	"github.com/koblas/grpc-todo/pkg/eventbus"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/tokenmanager"
	"github.com/koblas/grpc-todo/pkg/util"
	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	genpb "github.com/koblas/grpc-todo/twpb/core"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type OauthUserServer struct {
	pubsub   eventbus.Producer
	kms      key_manager.Encoder
	user     genpb.UserService
	jwtMaker tokenmanager.Maker
	smanager oauth_provider.SecretManager
}

func NewOauthUserServer(producer eventbus.Producer, user genpb.UserService, config SsmConfig, smanager oauth_provider.SecretManager) *OauthUserServer {
	maker, err := tokenmanager.NewJWTMaker(config.JwtSecret)

	if err != nil {
		log.Fatal(err)
	}

	return &OauthUserServer{
		pubsub:   producer,
		kms:      key_manager.NewSecureClear(),
		jwtMaker: maker,
		user:     user,
		smanager: smanager,
	}
}

func (svc *OauthUserServer) GetAuthURL(ctx context.Context, params *genpb.AuthOAuthGetUrlParams) (*genpb.AuthUserGetUrlResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("Calling GetAuthURL")
	oprovider, err := oauth_provider.GetOAuthProvider(params.GetProvider(), svc.smanager, log)
	if err != nil {
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

	if err != nil {
		return nil, err
	}

	url := oprovider.BuildRedirect(ctx, params.RedirectUrl, state)

	return &genpb.AuthUserGetUrlResult{Url: url}, nil
}

func (svc *OauthUserServer) UpsertUser(ctx context.Context, params *genpb.AuthUserUpsertParams) (*genpb.AuthUserUpsertResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Oauth.Provider)
	log.Info("Calling UpsertUser")
	oprovider, err := oauth_provider.GetOAuthProvider(params.Oauth.Provider, svc.smanager, log)

	if err != nil {
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

	user, err := svc.user.FindBy(ctx, &genpb.FindParam{
		Auth: &genpb.AuthInfo{
			Provider:   params.Oauth.Provider,
			ProviderId: info.Id,
		},
	})
	if err != nil {
		if e, ok := err.(twirp.Error); !ok || e.Code() != twirp.NotFound {
			log.With("error", err).Info("Failed to get oauth user")
			return nil, twirp.InternalErrorWith(err)
		}
	}

	if user != nil {
		return &genpb.AuthUserUpsertResult{UserId: user.Id, Created: false}, nil
	}

	if info.Email == "" {
		log.With("error", err).Info("Failed to get user email")
		return nil, twirp.InvalidArgumentError("email", "provider didn't send email address")
	}

	user, err = svc.user.FindBy(ctx, &genpb.FindParam{Email: info.Email})
	if err != nil {
		if e, ok := err.(twirp.Error); !ok || e.Code() != twirp.NotFound {
			log.With("error", err).Info("Failed to lookup user")
			return nil, twirp.InternalErrorWith(err)
		}
	}
	created := false
	if user == nil || user.Id == "" {
		log.With("email", info.Email).Info("Creating new user")
		created = true
		user, err = svc.user.Create(ctx, &genpb.CreateParam{
			Email: info.Email,
			Name:  info.Name,
			// TODO - create as "ACTIVE" since we "know" the email is good
		})
		if err != nil {
			log.With("error", err).Info("Unable to create user")
			return nil, twirp.InternalErrorWith(err)

		}
	} else {
		log.With("email", user.Email).Info("User already exists, associating")
	}

	// Now associate the OAuth token and the UserId
	// TODO - save the token!
	_, err = svc.user.AuthAssociate(ctx, &genpb.AuthAssociateParam{
		UserId: user.Id,
		Auth: &genpb.AuthInfo{
			Provider:   params.Oauth.Provider,
			ProviderId: info.Id,
		},
	})
	if err != nil {
		log.With("error", err).Info("Unable to associate")
		return nil, twirp.InternalErrorWith(err)
	}

	return &genpb.AuthUserUpsertResult{UserId: user.Id, Created: created}, nil
}

func (s *OauthUserServer) ListAssociations(ctx context.Context, params *genpb.AuthUserGetParams) (*genpb.AuthUserListAssoicationsResponse, error) {
	return &genpb.AuthUserListAssoicationsResponse{}, nil
}

func (s *OauthUserServer) RemoveAssociation(ctx context.Context, params *genpb.AuthUserGetParams) (*genpb.Empty, error) {
	return &genpb.Empty{}, nil
}

package user

import (
	"fmt"

	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	oauth_provider "github.com/koblas/grpc-todo/services/core/oauth_user/provider"
	"github.com/robinjoseph08/redisqueue"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents the gRPC server
type OauthUserServer struct {
	genpb.UnimplementedOauthUserServiceServer

	logger logger.Logger
	pubsub *redisqueue.Producer
	kms    key_manager.Encoder
	user   genpb.UserServiceClient
	store  OAuthStore
}

func NewOauthUserServer(log logger.Logger, user genpb.UserServiceClient) *OauthUserServer {
	pubsub, err := redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      1000,
		ApproximateMaxLength: true,
		RedisOptions: &redisqueue.RedisOptions{
			Addr: util.Getenv("REDIS_ADDR", "redis:6379"),
		},
	})
	if err != nil {
		log.With(err).Fatal("unable to start producer")
	}
	return &OauthUserServer{
		logger: log,
		pubsub: pubsub,
		kms:    key_manager.NewSecureClear(),
		store:  NewOauthMemoryStore(),
		user:   user,
	}
}

func (s *OauthUserServer) GetAuthURL(ctx context.Context, params *genpb.OauthUserGetUrlParams) (*genpb.OauthUserGetUrlResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("Calling GetAuthURL")
	oprovider, err := oauth_provider.GetOAuthProvider(params.GetProvider(), log)

	if err != nil {
		return nil, err
	}

	url := oprovider.BuildRedirect(ctx, params.RedirectUrl, params.State)

	return &genpb.OauthUserGetUrlResult{Url: url}, nil
}

func (s *OauthUserServer) UpsertUser(ctx context.Context, params *genpb.OauthUserUpsertParams) (*genpb.OauthUserUpsertResult, error) {
	log := logger.FromContext(ctx).With("provider", params.Provider)
	log.Info("Calling UpsertUser")
	oprovider, err := oauth_provider.GetOAuthProvider(params.GetProvider(), log)

	if err != nil {
		return nil, err
	}

	// TODO State?
	tokenResult, err := oprovider.GetAccessToken(ctx, params.Code, params.RedirectUrl, params.State)
	if err != nil {
		log.With("error", err).Info("Failed to get access token")
		return nil, err
	}
	log.Info("Getting OAuth User information")

	info, err := oprovider.GetInfo(ctx, tokenResult)
	if err != nil {
		log.With("error", err).Info("Failed to get user info")
		return nil, err
	}

	existList, err := s.store.FindByProviderId(params.Provider, info.Id)
	if err != nil {
		log.With("error", err).Info("Failed to get oauth user")
		return nil, fmt.Errorf("provider didn't send email address")
	}

	if len(existList) != 0 {
		return &genpb.OauthUserUpsertResult{UserId: existList[0].UserId, Created: false}, nil
	}

	if info.Email == "" {
		log.With("error", err).Info("Failed to get user email")
		return nil, fmt.Errorf("provider didn't send email address")
	}

	user, err := s.user.FindBy(ctx, &genpb.FindParam{Email: info.Email})
	if err != nil {
		status := status.Convert(err)
		if status.Code() != codes.NotFound {
			log.With("error", err).Info("Failed to lookup user")
			return nil, fmt.Errorf("unable to check email")
		}
	}
	created := false
	if user == nil || user.Id == "" {
		log.With("email", info.Email).Info("Creating new user")
		created = true
		user, err = s.user.Create(ctx, &genpb.CreateParam{
			Email: info.Email,
			Name:  info.Name,
			// TODO - create as "ACTIVE" since we "know" the email is good
		})
		if err != nil {
			log.With("error", err).Info("Unable to create user")
			return nil, fmt.Errorf("unable to create user")

		}
	} else {
		log.With("email", user.Email).Info("User already exists, associating")
	}

	// Now associate the OAuth token and the UserId
	err = s.store.Associate(user.Id, params.Provider, info.Id, tokenResult)
	if err != nil {
		log.With("error", err).Info("Unable to associate")
		return nil, fmt.Errorf("unable to create oauth assoication")
	}

	return &genpb.OauthUserUpsertResult{UserId: user.Id, Created: created}, nil
}

func (s *OauthUserServer) ListAssociations(ctx context.Context, params *genpb.OauthUserGetParams) (*genpb.OauthUserListAssoicationsResponse, error) {
	return &genpb.OauthUserListAssoicationsResponse{}, nil
}

func (s *OauthUserServer) RemoveAssociation(ctx context.Context, params *genpb.OauthUserGetParams) (*genpb.Empty, error) {
	return &genpb.Empty{}, nil
}

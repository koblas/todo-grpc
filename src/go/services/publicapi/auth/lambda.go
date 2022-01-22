package auth

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-redis/redis/v8"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
	publicapi "github.com/koblas/grpc-todo/twpb/publicapi"
)

type SsmConfig struct {
	RedisAddr string `ssm:"redis_addr" environment:"REDIS_ADDR"`
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET"`
}

var ssmConfig *SsmConfig
var api core.TwirpServer

func HandleLambda() func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log := logger.NewZap(logger.LevelInfo)
	ctx := logger.ToContext(context.Background(), log)

	if ssmConfig == nil || api == nil {
		ssmConfig = &SsmConfig{}
		err := awsutil.LoadSsmConfig("/common/", ssmConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Connect to the user service
		userService := core.NewUserServiceJSONClient("lambda://core-user", awsutil.NewTwirpCallLambda())
		oauthService := core.NewAuthUserServiceJSONClient("lambda://core-oauth-user", awsutil.NewTwirpCallLambda())

		// Connect to redis
		var rdb *redis.Client
		if ssmConfig.RedisAddr != "" {
			log.With("address", ssmConfig.RedisAddr).Info("Redis Address")
			// TODO - re-enable this
			rdb = redis.NewClient(&redis.Options{
				Addr:        ssmConfig.RedisAddr,
				Password:    "",                     // no password set
				DB:          0,                      // use default DB
				DialTimeout: time.Millisecond * 200, // either it happens or it doesn't
			})
		}

		s := NewAuthenticationServer(*ssmConfig, userService, oauthService, NewAttemptCounter("publicapi:authentication", rdb))
		api = publicapi.NewAuthenticationServiceServer(s)
	}

	return awsutil.HandleApiLambda(ctx, api)
}

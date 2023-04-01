package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	"github.com/koblas/grpc-todo/gen/core/user/v1/userv1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/natsutil"
	"github.com/koblas/grpc-todo/services/core/user"

	"go.uber.org/zap"
)

type endpointResolver struct {
	hostname string
}

func (h *endpointResolver) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return aws.Endpoint{URL: "http://" + h.hostname}, nil
}

func dynamoClient(endpoint string) *dynamodb.Client {
	if endpoint == "" {
		return nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptions(&endpointResolver{endpoint})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

type Config struct {
	NatsAddr        string
	DynamoStoreAddr *string `environment:"DYNAMO_STORE" json:"dynamo-store"`
}

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := Config{}
	if err := confmgr.Parse(&config, confmgr.NewLoaderEnvironment("", "_"), confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := eventbusv1connect.NewUserEventbusServiceClient(
		natsutil.NewNatsClient(config.NatsAddr),
		"",
	)

	opts := []user.Option{
		user.WithProducer(producer),
	}

	if config.DynamoStoreAddr == nil || *config.DynamoStoreAddr == "" {
		log.Info("Starting up with Memory store")
		opts = append(opts, user.WithUserStore(user.NewUserMemoryStore()))
	} else {
		log.With(
			zap.String("dynamoAddr", *config.DynamoStoreAddr),
		).Info("Starting up with DynamoDB store")
		opts = append(opts,
			user.WithUserStore(
				user.NewUserDynamoStore(
					user.WithDynamoClient(
						dynamoClient(*config.DynamoStoreAddr),
					),
				),
			),
		)
	}

	mux := http.NewServeMux()
	mux.Handle(userv1connect.NewUserServiceHandler(
		user.NewUserServer(opts...),
		connect.WithInterceptors(interceptors.NewReqidInterceptor()),
		connect.WithCompressMinBytes(1024),
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(userv1connect.UserServiceName),
		connect.WithCompressMinBytes(1024),
	))

	mgr.Start(mgr.WrapHttpHandler(mux))
}

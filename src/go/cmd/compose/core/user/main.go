package main

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/confmgr"
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

func main() {
	mgr := manager.NewManager(manager.WithGrpcHealth("15050"))
	log := mgr.Logger()

	var config user.Config
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	producer := corev1connect.NewUserEventbusServiceClient(
		natsutil.NewNatsClient(config.NatsAddr),
		"",
	)

	opts := []user.Option{
		user.WithProducer(producer),
	}

	if config.DynamoStore == "" || config.DynamoStore == "false" {
		log.Info("Starting up with Memory store")
		opts = append(opts, user.WithUserStore(user.NewUserMemoryStore()))
	} else {
		log.With(
			zap.String("dynamoAddr", config.DynamoStore),
		).Info("Starting up with DynamoDB store")
		opts = append(opts,
			user.WithUserStore(
				user.NewUserDynamoStore(
					user.WithDynamoClient(
						dynamoClient(config.DynamoStore),
					),
				),
			),
		)
	}

	_, api := corev1connect.NewUserServiceHandler(user.NewUserServer(opts...))

	mgr.Start(mgr.WrapHttpHandler(api))
}

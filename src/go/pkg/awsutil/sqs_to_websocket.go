package awsutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func ApigwClient(ctx context.Context, endpoint string) (*apigatewaymanagementapi.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	endpointResolver := func(o *apigatewaymanagementapi.Options) {
		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(endpoint)
	}

	return apigatewaymanagementapi.NewFromConfig(cfg, endpointResolver), nil
}

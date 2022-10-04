package awsutil

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/koblas/grpc-todo/pkg/eventbus"
)

func ConvertEventbusToApiGateway(msg *eventbus.SimpleMessage) apigatewaymanagementapi.PostToConnectionInput {
	return apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(msg.Attributes["connectionId"]),
		Data:         []byte(msg.Body),
	}
}

func ConvertApiGatewayToMessage(msg *apigatewaymanagementapi.PostToConnectionInput) eventbus.SimpleMessage {
	return eventbus.SimpleMessage{
		Attributes: map[string]string{
			"connectionId": *msg.ConnectionId,
		},
		Body: string(msg.Data),
	}
}

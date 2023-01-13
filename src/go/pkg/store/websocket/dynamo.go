package websocket

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type dynamoStore struct {
	client *dynamodb.Client
	table  *string
}

type dynamoConnection struct {
	Pk       string    `dynamodbav:"pk"`
	Sk       string    `dynamodbav:"sk"`
	DeleteAt time.Time `dynamodbav:"delete_at,unixtime"`
}

type DynamoOption struct {
	Client *dynamodb.Client
	Table  *string
}

func WithDynamoClient(client *dynamodb.Client) DynamoOption {
	return DynamoOption{
		Client: client,
	}
}
func WithDynamoTable(value string) DynamoOption {
	return DynamoOption{
		Table: &value,
	}
}

func NewWsConnectionDynamoStore(opts ...DynamoOption) ConnectionStore {
	state := dynamoStore{
		table: aws.String("ws-connection"),
	}

	for _, opt := range opts {
		if opt.Client != nil {
			state.client = opt.Client
		}
		if opt.Table != nil {
			state.table = opt.Table
		}
	}

	if state.client == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO())

		if err != nil {
			panic(err)
		}

		state.client = dynamodb.NewFromConfig(cfg)
	}

	return &state
}

func (store *dynamoStore) Create(ctx context.Context, userId string, connectionId string) error {
	expires := time.Now().Add(2 * time.Hour)

	byUser, err := attributevalue.MarshalMap(dynamoConnection{
		Pk:       "USER#" + userId,
		Sk:       connectionId,
		DeleteAt: expires,
	})
	if err != nil {
		return err
	}
	byConn, err := attributevalue.MarshalMap(dynamoConnection{
		Pk:       "CONN#" + connectionId,
		Sk:       userId,
		DeleteAt: expires,
	})
	if err != nil {
		return err
	}

	_, err = store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: {
				{
					PutRequest: &types.PutRequest{
						Item: byUser,
					},
				},
				{
					PutRequest: &types.PutRequest{
						Item: byConn,
					},
				},
			},
		},
	})

	return err
}

/**
 * Heartbeat will update the expires on the DynamoDB table such that they don't go away
 */
func (store *dynamoStore) Heartbeat(ctx context.Context, connectionId string) error {
	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: "CONN#" + connectionId},
		},
	})

	if err != nil || len(out.Items) == 0 {
		return err
	}

	obj := dynamoConnection{}
	if err := attributevalue.UnmarshalMap(out.Items[0], &obj); err != nil {
		return err
	}

	// Update the rows

	expires := time.Now().Add(2 * time.Hour)

	byUser, err := attributevalue.MarshalMap(dynamoConnection{
		Pk:       "USER#" + obj.Sk,
		Sk:       connectionId,
		DeleteAt: expires,
	})
	if err != nil {
		return err
	}
	byConn, err := attributevalue.MarshalMap(dynamoConnection{
		Pk:       obj.Pk,
		Sk:       obj.Sk,
		DeleteAt: expires,
	})
	if err != nil {
		return err
	}

	_, err = store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: {
				{
					PutRequest: &types.PutRequest{
						Item: byUser,
					},
				},
				{
					PutRequest: &types.PutRequest{
						Item: byConn,
					},
				},
			},
		},
	})

	return err
}

func (store *dynamoStore) Delete(ctx context.Context, connectionId string) error {
	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: "CONN#" + connectionId},
		},
	})

	if err != nil || len(out.Items) == 0 {
		return err
	}

	obj := dynamoConnection{}
	if err := attributevalue.UnmarshalMap(out.Items[0], &obj); err != nil {
		return err
	}

	_, err = store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: {
				{
					DeleteRequest: &types.DeleteRequest{
						Key: map[string]types.AttributeValue{
							"pk": &types.AttributeValueMemberS{Value: "USER#" + obj.Sk},
							"sk": &types.AttributeValueMemberS{Value: connectionId},
						},
					},
				},
				{
					DeleteRequest: &types.DeleteRequest{
						Key: map[string]types.AttributeValue{
							"pk": &types.AttributeValueMemberS{Value: "CONN#" + connectionId},
							"sk": &types.AttributeValueMemberS{Value: obj.Sk},
						},
					},
				},
			},
		},
	})

	return err
}

func (store *dynamoStore) ForUser(ctx context.Context, user_id string) ([]string, error) {
	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: "USER#" + user_id},
		},
	})

	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, item := range out.Items {
		data := dynamoConnection{}
		if err := attributevalue.UnmarshalMap(item, &data); err != nil {
			return nil, err
		}

		result = append(result, data.Sk)
	}

	return result, nil
}

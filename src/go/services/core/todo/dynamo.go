package todo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/koblas/grpc-todo/pkg/awsutil"
)

type dynamoTodo struct {
	Pk string `dynamodbav:"pk"`
	Sk string `dynamodbav:"sk"`
	Todo
}

type dynamoStore struct {
	client *dynamodb.Client
	table  *string
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

func NewTodoDynamoStore(opts ...DynamoOption) TodoStore {
	state := dynamoStore{
		table: aws.String("app-todo"),
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

func (store *dynamoStore) FindByUser(ctx context.Context, user_id string) ([]*Todo, error) {
	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: user_id},
		},
		// ExpressionAttributeNames: map[string]string{
		// 	"#date": "date",
		// },
	})

	if err != nil {
		return nil, err
	}

	// return awsutil.UnmarshalList[*Todo](out.Items)

	todos := []*Todo{}
	for _, item := range out.Items {
		todo := Todo{}
		if err := attributevalue.UnmarshalMap(item, &todo); err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}
	return todos, nil
}

func (store *dynamoStore) DeleteOne(ctx context.Context, user_id string, id string) (*Todo, error) {
	item, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: user_id},
			"sk": &types.AttributeValueMemberS{Value: id},
		},
	})

	// Already deleted (or error)
	if item == nil {
		return nil, err
	}

	todo, err := awsutil.UnmarshalItem[Todo](item.Item)
	if err != nil {
		return nil, err
	}
	// todo := Todo{}
	// if err := attributevalue.UnmarshalMap(item.Item, &todo); err != nil {
	// 	return nil, err
	// }

	_, err = store.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: user_id},
			"sk": &types.AttributeValueMemberS{Value: id},
		},
	})

	return todo, err
}

func (store *dynamoStore) Create(ctx context.Context, todo Todo) (*Todo, error) {
	av, err := attributevalue.MarshalMap(dynamoTodo{
		Todo: todo,
		Pk:   todo.UserId,
		Sk:   todo.ID,
	})
	if err != nil {
		return nil, err
	}

	_, err = store.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: store.table,
		Item:      av,
	})

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

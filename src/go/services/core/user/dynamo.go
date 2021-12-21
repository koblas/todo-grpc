package user

// https://dynobase.dev/dynamodb-golang-query-examples/

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamoStore struct {
	client *dynamodb.Client
	table  *string
}

type dynamoUser struct {
	Pk      string `dynamodbav:"pk"`
	EmailLc string `dynamodbav:"email_lc"`
	User
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

func NewUserDynamoStore(opts ...DynamoOption) UserStore {
	state := dynamoStore{
		table: aws.String("app-user"),
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

func (store *dynamoStore) buildIdKey(id string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: "id#" + id}
}
func (store *dynamoStore) buildEmailKey(email string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: "userEmail#" + strings.ToLower(email)}
}

func (store *dynamoStore) GetById(id string) (*User, error) {
	out, err := store.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": store.buildIdKey(id),
			// "sk": store.buildIdKey(id),
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	user := User{}
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (store *dynamoStore) GetByEmail(email string) (*User, error) {
	ctx := context.TODO()

	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		IndexName:              aws.String(*store.table + "-by-email"),
		KeyConditionExpression: aws.String("email_lc = :email_lc"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email_lc": &types.AttributeValueMemberS{Value: strings.ToLower(email)},
		},
	})

	if err != nil || out == nil || len(out.Items) == 0 {
		return nil, err
	}

	user := dynamoUser{}
	if err := attributevalue.UnmarshalMap(out.Items[0], &user); err != nil {
		return nil, err
	}

	return store.GetById(user.Pk[3:])
}

func (store *dynamoStore) CreateUser(user *User) error {
	av, err := attributevalue.MarshalMap(dynamoUser{
		Pk:      "id#" + user.ID,
		EmailLc: strings.ToLower(user.Email),
		User:    *user,
	})
	if err != nil {
		return err
	}

	_, err = store.client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName: store.table,
					Item:      av,
				},
			},
			{
				Put: &types.Put{
					TableName: store.table,
					Item: map[string]types.AttributeValue{
						"pk":      store.buildEmailKey(user.Email),
						"user_id": &types.AttributeValueMemberS{Value: strings.ToLower(user.Email)},
					},
				},
			},
		},
	})

	return err
}

func (store *dynamoStore) UpdateUser(user *User) error {
	av, err := attributevalue.MarshalMap(dynamoUser{
		User: *user,
		Pk:   user.ID,
	})
	if err != nil {
		return err
	}

	// TODO -- if email changes, we need to update it...
	_, err = store.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: store.table,
		Item:      av,
	})

	return err
}

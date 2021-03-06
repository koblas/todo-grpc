package user

// https://dynobase.dev/dynamodb-golang-query-examples/

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/twitchtv/twirp"
)

type dynamoStore struct {
	client *dynamodb.Client
	table  *string
}

type dynamoUser struct {
	Pk      string `dynamodbav:"pk"`
	EmailLc string `dynamodbav:"email_lc"`
	UserId  string `dynamodbav:"user_id"`
	User
}

type dynamoAuth struct {
	Pk string `dynamodbav:"pk"`
	Sk string `dynamodbav:"sk"`
	UserAuth
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

func (store *dynamoStore) buildAuthKey(provider, provider_id string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: "auth#" + provider + "#" + provider_id}
}
func (store *dynamoStore) buildIdKey(id string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: "user#" + id}
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

	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": store.buildEmailKey(email),
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	user := dynamoUser{}
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return store.GetById(user.UserId)
}

func (store *dynamoStore) CreateUser(user User) error {
	av, err := attributevalue.MarshalMap(dynamoUser{
		Pk:      store.buildIdKey(user.ID).Value,
		EmailLc: strings.ToLower(user.Email),
		User:    user,
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
						"user_id": &types.AttributeValueMemberS{Value: user.ID},
					},
				},
			},
		},
	})

	return err
}

func (store *dynamoStore) UpdateUser(user *User) error {
	old, err := store.GetById(user.ID)

	if err != nil {
		return err
	}

	av, err := attributevalue.MarshalMap(dynamoUser{
		User: *user,
		Pk:   store.buildIdKey(user.ID).Value,
	})
	if err != nil {
		return err
	}

	transact := []types.TransactWriteItem{}

	transact = append(transact, types.TransactWriteItem{
		Put: &types.Put{
			TableName: store.table,
			Item:      av,
		},
	})
	if old.Email != user.Email {
		transact = append(transact,
			types.TransactWriteItem{
				Delete: &types.Delete{
					TableName: store.table,
					Key: map[string]types.AttributeValue{
						"pk": store.buildEmailKey(old.Email),
					},
				},
			},
			types.TransactWriteItem{
				Put: &types.Put{
					TableName: store.table,
					Item: map[string]types.AttributeValue{
						"pk":      store.buildEmailKey(user.Email),
						"user_id": &types.AttributeValueMemberS{Value: user.ID},
					},
					ConditionExpression: aws.String("attribute_not_exists(pk)"),
				},
			},
		)
	}

	_, err = store.client.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: transact,
	})

	return err
}

func (store *dynamoStore) AuthUpsert(provider, provider_id string, auth UserAuth) error {
	if provider == "" || provider_id == "" {
		return twirp.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	av, err := attributevalue.MarshalMap(dynamoAuth{
		Pk:       store.buildAuthKey(provider, provider_id).Value,
		UserAuth: auth,
	})
	if err != nil {
		return err
	}

	request := map[string][]types.WriteRequest{}

	request[*store.table] = []types.WriteRequest{
		{
			PutRequest: &types.PutRequest{
				Item: av,
			},
		},
	}

	out, err := store.client.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: request,
	})

	if len(out.UnprocessedItems) != 0 {
		return fmt.Errorf("write failed - uprocessed items")
	}

	return err
}

func (store *dynamoStore) AuthGet(provider, provider_id string) (*UserAuth, error) {
	if provider == "" || provider_id == "" {
		return nil, twirp.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	out, err := store.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": store.buildAuthKey(provider, provider_id),
			// "sk": store.buildIdKey(id),
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	obj := UserAuth{}
	if err := attributevalue.UnmarshalMap(out.Item, &obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func (store *dynamoStore) AuthDelete(provider, provider_id string, auth UserAuth) error {
	if provider == "" || provider_id == "" {
		return twirp.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	_, err := store.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": store.buildAuthKey(provider, provider_id),
			// "sk": store.buildIdKey(id),
		},
		ConditionExpression: aws.String("user_id = :user_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_id": &types.AttributeValueMemberS{Value: auth.UserID},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

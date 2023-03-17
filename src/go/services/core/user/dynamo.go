package user

// https://dynobase.dev/dynamodb-golang-query-examples/

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
)

type dynamoStore struct {
	client      *dynamodb.Client
	table       *string
	tableExists bool
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

var resourceNotFound = new(types.ResourceNotFoundException)

func (store *dynamoStore) checkCreateTable(ctx context.Context) error {
	if store.tableExists {
		return nil
	}
	_, err := store.client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: store.table,
	})
	if err == nil {
		store.tableExists = true
		return nil
	}

	log := logger.FromContext(ctx).With(zap.String("tableName", *store.table))
	if !errors.As(err, &resourceNotFound) {
		log.With(zap.Error(err)).Error("DescribeTable failed")
		return err
	}
	log.Info("DescribeTable failed, creating table")

	_, err = store.client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName:   store.table,
		BillingMode: types.BillingModePayPerRequest,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
	})

	if err != nil {
		log.With(zap.Error(err)).Error("CreateTable failed")
	}
	log.Info("create complete")

	return err
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

func (store *dynamoStore) GetById(ctx context.Context, id string) (*User, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
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

func (store *dynamoStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
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

	return store.GetById(ctx, user.UserId)
}

func (store *dynamoStore) CreateUser(ctx context.Context, user User) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	av, err := attributevalue.MarshalMap(dynamoUser{
		Pk:      store.buildIdKey(user.ID).Value,
		EmailLc: strings.ToLower(user.Email),
		User:    user,
	})
	if err != nil {
		return err
	}

	_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
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

func (store *dynamoStore) UpdateUser(ctx context.Context, user *User) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	old, err := store.GetById(ctx, user.ID)

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

	_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transact,
	})

	return err
}

func (store *dynamoStore) AuthUpsert(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	if provider == "" || provider_id == "" {
		return bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
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

	out, err := store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: request,
	})

	if len(out.UnprocessedItems) != 0 {
		return fmt.Errorf("write failed - uprocessed items")
	}

	return err
}

func (store *dynamoStore) AuthGet(ctx context.Context, provider, provider_id string) (*UserAuth, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	if provider == "" || provider_id == "" {
		return nil, bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
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

func (store *dynamoStore) AuthDelete(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	if provider == "" || provider_id == "" {
		return bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	_, err := store.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
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

package user

// https://dynobase.dev/dynamodb-golang-query-examples/

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type DynamoStore struct {
	client      *dynamodb.Client
	table       *string
	tableExists bool
}

type dynamoUser struct {
	Pk      string `dynamodbav:"pk"`
	Sk      string `dynamodbav:"sk"`
	EmailLc string `dynamodbav:"email_lc"`
	UserId  string `dynamodbav:"user_id"`
	Type    string `dynamodbav:"type"`

	// User fields
	ID             string                       `dynamodbav:"id"`
	Name           string                       `dynamodbav:"name"`
	Email          string                       `dynamodbav:"email"`
	VerifiedEmails []string                     `dynamodbav:"verified_email"`
	Status         UserStatus                   `dynamodbav:"status"`
	Settings       map[string]map[string]string `dynamodbav:"settings"`
	AvatarUrl      *string                      `dynamodbav:"avatar_url,nullempty"`

	// For email address confirmation
	EmailVerifyToken     []byte     `dynamodbav:"email_verify_token,nullempty"`
	EmailVerifyExpiresAt *time.Time `dynamodbav:"email_verify_expires_at,nullempty"`
}

func (store *DynamoStore) marshalUser(user *User) *dynamoUser {
	return &dynamoUser{
		Pk:                   store.makeUserKey(user.ID),
		Sk:                   "A",
		EmailLc:              strings.ToLower(user.Email),
		Type:                 "user",
		ID:                   user.ID,
		Name:                 user.Name,
		Email:                user.Email,
		VerifiedEmails:       user.VerifiedEmails,
		Status:               user.Status,
		Settings:             user.Settings,
		AvatarUrl:            user.AvatarUrl,
		EmailVerifyToken:     user.EmailVerifyToken,
		EmailVerifyExpiresAt: user.EmailVerifyExpiresAt,
	}
}

func (obj dynamoUser) unmarshal() *User {
	return &User{
		ID:                   obj.ID,
		Name:                 obj.Name,
		Email:                obj.Email,
		VerifiedEmails:       obj.VerifiedEmails,
		Status:               obj.Status,
		Settings:             obj.Settings,
		AvatarUrl:            obj.AvatarUrl,
		EmailVerifyToken:     obj.EmailVerifyToken,
		EmailVerifyExpiresAt: obj.EmailVerifyExpiresAt,
	}
}

// Dynamo version of the authentication information
type dynamoAuth struct {
	Pk   string `dynamodbav:"pk"`
	Sk   string `dynamodbav:"sk"`
	Type string `dynamodbav:"type"`

	UserID    string     `dynamodbav:"user_id"`
	Password  []byte     `dynamodbav:"password"`
	ExpiresAt *time.Time `dynamodbav:"expires_at,nullempty"`
}

func (store *DynamoStore) marshalAuth(provider, provider_id string, auth *UserAuth) *dynamoAuth {
	return &dynamoAuth{
		Pk:        store.makeUserAuthKey(provider, provider_id),
		Sk:        "A",
		Type:      "userAuth",
		UserID:    auth.UserID,
		Password:  auth.Password,
		ExpiresAt: auth.ExpiresAt,
	}

}

func (obj dynamoAuth) unmarshal() *UserAuth {
	return &UserAuth{
		UserID:    obj.UserID,
		Password:  obj.Password,
		ExpiresAt: obj.ExpiresAt,
	}
}

// Dynamo version of the authentication information
type dynamoUserEmail struct {
	Pk   string `dynamodbav:"pk"`
	Sk   string `dynamodbav:"sk"`
	Type string `dynamodbav:"type"`

	UserID string `dynamodbav:"user_id"`
}

func (store *DynamoStore) marshalUserEmail(user *User) *dynamoUserEmail {
	return &dynamoUserEmail{
		Pk:     store.makeUserEmailKey(user.Email),
		Sk:     "A",
		Type:   "userEmail",
		UserID: user.ID,
	}

}

// the Data store

var _ UserStore = (*DynamoStore)(nil)
var _ OAuthStore = (*DynamoStore)(nil)
var _ TeamStore = (*DynamoStore)(nil)

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

func NewDynamoStore(opts ...DynamoOption) *DynamoStore {
	state := DynamoStore{
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

func (store *DynamoStore) checkCreateTable(ctx context.Context) error {
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
			{
				AttributeName: aws.String("sk"),
				KeyType:       types.KeyTypeRange,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
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

func (store *DynamoStore) GetById(ctx context.Context, id string) (*User, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	av, err := attributevalue.MarshalMap(store.marshalUser(&User{ID: id}))
	if err != nil {
		return nil, err
	}
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	user := dynamoUser{}
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return user.unmarshal(), nil
}

func (store *DynamoStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	av, err := attributevalue.MarshalMap(store.marshalUserEmail(&User{Email: email}))
	if err != nil {
		return nil, err
	}
	fmt.Println("GETTING  pk=", av["pk"])
	fmt.Println("GETTING  sk=", av["sk"])
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	user := dynamoUserEmail{}
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return store.GetById(ctx, user.UserID)
}

func (store *DynamoStore) CreateUser(ctx context.Context, user User) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	userAv, err := attributevalue.MarshalMap(store.marshalUser(&user))
	if err != nil {
		return err
	}
	emailAv, err := attributevalue.MarshalMap(store.marshalUserEmail(&user))
	if err != nil {
		return err
	}

	// By default every user is a member of their own "default" team
	_, teamItems, err := store.teamCreateItems(ctx, "", TeamUser{
		UserId: user.ID,
		TeamId: "",
		Role:   "admin",
	})
	if err != nil {
		return err
	}

	transaction := []types.TransactWriteItem{
		{
			Put: &types.Put{
				TableName: store.table,
				Item:      userAv,
			},
		},
		{
			Put: &types.Put{
				TableName:           store.table,
				Item:                emailAv,
				ConditionExpression: aws.String("attribute_not_exists(pk)"),
			},
		},
	}
	// Convert from Batch item to Transaction item
	for _, item := range teamItems {
		transaction = append(transaction, types.TransactWriteItem{
			Put: &types.Put{
				TableName: store.table,
				Item:      item.PutRequest.Item,
			},
		})
	}

	_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transaction,
	})

	return err
}

func (store *DynamoStore) UpdateUser(ctx context.Context, user *User) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}

	av, err := attributevalue.MarshalMap(store.marshalUser(user))
	if err != nil {
		return err
	}

	old, err := store.GetById(ctx, user.ID)
	if err != nil {
		return err
	}

	// We only need to use a transaction if we're updating the email address
	if old.Email == user.Email {
		_, err := store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				*store.table: {
					{
						PutRequest: &types.PutRequest{
							Item: av,
						},
					},
				},
			},
		})

		return err
	} else {
		oldAv, err := attributevalue.MarshalMap(store.marshalUserEmail(old))
		if err != nil {
			return err
		}
		emailAv, err := attributevalue.MarshalMap(store.marshalUserEmail(user))
		if err != nil {
			return err
		}

		transact := []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName: store.table,
					Item:      av,
				},
			},

			{
				Delete: &types.Delete{
					TableName: store.table,
					Key: map[string]types.AttributeValue{
						"pk": oldAv["pk"],
						"sk": oldAv["pk"],
					},
				},
			},
			{
				Put: &types.Put{
					TableName:           store.table,
					Item:                emailAv,
					ConditionExpression: aws.String("attribute_not_exists(pk)"),
				},
			},
		}

		_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
			TransactItems: transact,
		})

		return err
	}
}

func (store *DynamoStore) AuthUpsert(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	if provider == "" || provider_id == "" {
		return bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	av, err := attributevalue.MarshalMap(store.marshalAuth(provider, provider_id, &auth))
	if err != nil {
		return err
	}

	out, err := store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: {
				{
					PutRequest: &types.PutRequest{
						Item: av,
					},
				},
			},
		},
	})

	if len(out.UnprocessedItems) != 0 {
		return fmt.Errorf("write failed - uprocessed items")
	}

	return err
}

func (store *DynamoStore) AuthGet(ctx context.Context, provider, provider_id string) (*UserAuth, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	if provider == "" || provider_id == "" {
		return nil, bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	av, err := attributevalue.MarshalMap(store.marshalAuth(provider, provider_id, &UserAuth{}))
	if err != nil {
		return nil, err
	}
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	obj := dynamoAuth{}
	if err := attributevalue.UnmarshalMap(out.Item, &obj); err != nil {
		return nil, err
	}

	return obj.unmarshal(), nil
}

func (store *DynamoStore) AuthDelete(ctx context.Context, provider, provider_id string, auth UserAuth) error {
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}
	if provider == "" || provider_id == "" {
		return bufcutil.InvalidArgumentError("provider", "provider or provider_id is empty")
	}
	av, err := attributevalue.MarshalMap(store.marshalAuth(provider, provider_id, &auth))
	if err != nil {
		return err
	}
	_, err = store.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
		ConditionExpression: aws.String("user_id = :user_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_id": av["user_id"],
		},
	})

	if err != nil {
		return err
	}

	return nil
}

//
// Team handling
//

func (*DynamoStore) makeTeamKey(teamId string) string {
	return "TEAM#" + teamId
}

func (*DynamoStore) makeUserKey(userId string) string {
	return "USER#" + userId
}

func (*DynamoStore) makeUserAuthKey(provider, provider_id string) string {
	return "AUTH#" + provider + "#" + provider_id
}

func (*DynamoStore) makeUserEmailKey(userId string) string {
	return "USEREMAIL#" + userId
}

type dynamoTeam struct {
	Pk   string `dynamodbav:"pk"`
	Sk   string `dynamodbav:"sk"`
	Type string `dynamodbav:"type"`
	Id   string `dynamodbav:"team_id"`
	Name string `dynamodbav:"name"`
}
type dynamoTeamUser struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Type   string `dynamodbav:"type"`
	UserId string `dynamodbav:"user_id"`
	TeamId string `dynamodbav:"team_id"`
	Role   string `dynamodbav:"role"`
}

func (store *DynamoStore) marshalTeam(team *Team) *dynamoTeam {
	return &dynamoTeam{
		Pk:   store.makeTeamKey(team.TeamId),
		Sk:   "A",
		Type: "team",
		Id:   team.TeamId,
		Name: team.Name,
	}
}

func (store *DynamoStore) marshalTeamUser(teamId string, userId string, role string) *dynamoTeamUser {
	return &dynamoTeamUser{
		Pk:     store.makeTeamKey(teamId),
		Sk:     store.makeUserKey(userId),
		Type:   "teamUser",
		TeamId: teamId,
		UserId: userId,
		Role:   role,
	}
}

func (store *DynamoStore) marshalUserTeam(teamId string, userId string, role string) *dynamoTeamUser {
	return &dynamoTeamUser{
		Pk:     store.makeUserKey(userId),
		Sk:     store.makeTeamKey(teamId),
		Type:   "userTeam",
		TeamId: teamId,
		UserId: userId,
		Role:   role,
	}
}

func (team dynamoTeam) unmarshal() *Team {
	return &Team{
		TeamId: team.Id,
		Name:   team.Name,
	}
}

func (store *DynamoStore) createUserItems(ctx context.Context, teamId string, tuser ...TeamUser) ([]types.WriteRequest, error) {
	items := []types.WriteRequest{}
	for _, item := range tuser {
		tuserAv, err := attributevalue.MarshalMap(store.marshalTeamUser(teamId, item.UserId, item.Role))
		if err != nil {
			return nil, err
		}
		uteamAv, err := attributevalue.MarshalMap(store.marshalUserTeam(teamId, item.UserId, item.Role))
		if err != nil {
			return nil, err
		}

		// Write both keys
		items = append(items,
			types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: tuserAv,
				},
			},
			types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: uteamAv,
				},
			},
		)
	}

	return items, nil
}

func (store *DynamoStore) teamCreateItems(ctx context.Context, name string, tuser ...TeamUser) (*Team, []types.WriteRequest, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, nil, err
	}
	teamId := xid.New().String()
	team := Team{
		TeamId: teamId,
		Name:   name,
	}
	av, err := attributevalue.MarshalMap(store.marshalTeam(&team))
	if err != nil {
		return nil, nil, err
	}

	items, err := store.createUserItems(ctx, teamId, tuser...)
	if err != nil {
		return nil, nil, err
	}

	items = append(items, types.WriteRequest{
		PutRequest: &types.PutRequest{
			Item: av,
		},
	})

	return &team, items, nil
}

func (store *DynamoStore) TeamCreate(ctx context.Context, name string, tuser ...TeamUser) (*Team, error) {
	team, items, err := store.teamCreateItems(ctx, name, tuser...)
	if err != nil {
		return nil, err
	}

	_, err = store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: items,
		},
	})
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (store *DynamoStore) TeamGet(ctx context.Context, teamId string) (*Team, error) {
	av, err := attributevalue.MarshalMap(store.marshalTeam(&Team{TeamId: teamId}))
	if err != nil {
		return nil, err
	}

	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
	})
	if err != nil || len(out.Item) == 0 {
		return nil, err
	}
	msg := dynamoTeam{}
	if err := attributevalue.UnmarshalMap(out.Item, &msg); err != nil {
		return nil, err
	}

	return msg.unmarshal(), nil
}

func (store *DynamoStore) TeamAddUsers(ctx context.Context, tuser ...TeamUser) error {
	if len(tuser) == 0 {
		return nil
	}
	for _, item := range tuser {
		if item.TeamId != tuser[0].Role {
			return errors.New("team id mismatch")
		}
	}
	teamId := tuser[0].TeamId

	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}

	items, err := store.createUserItems(ctx, teamId, tuser...)
	if err != nil {
		return err
	}

	_, err = store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: items,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (store *DynamoStore) TeamDeleteUsers(ctx context.Context, teamId string, userIds ...string) error {
	items := []types.WriteRequest{}
	for _, userId := range userIds {
		tuserAv, err := attributevalue.MarshalMap(store.marshalTeamUser(teamId, userId, ""))
		if err != nil {
			return err
		}
		uteamAv, err := attributevalue.MarshalMap(store.marshalTeamUser(teamId, userId, ""))
		if err != nil {
			return err
		}

		items = append(items,
			types.WriteRequest{
				DeleteRequest: &types.DeleteRequest{
					Key: map[string]types.AttributeValue{
						"pk": tuserAv["pk"],
						"sk": tuserAv["sk"],
					},
				},
			},
			types.WriteRequest{
				DeleteRequest: &types.DeleteRequest{
					Key: map[string]types.AttributeValue{
						"pk": uteamAv["pk"],
						"sk": uteamAv["sk"],
					},
				},
			},
		)
	}
	_, err := store.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			*store.table: items,
		},
	})

	return err
}

func (store *DynamoStore) TeamListUsers(ctx context.Context, teamId string) ([]TeamUser, error) {
	tuserAv, err := attributevalue.MarshalMap(store.marshalTeamUser(teamId, "", ""))
	if err != nil {
		return nil, err
	}

	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :pk AND begins_with(sk, :prefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":     tuserAv["pk"],
			":prefix": tuserAv["sk"],
		},
	})

	if err != nil {
		return nil, err
	}

	users := []TeamUser{}
	for _, item := range out.Items {
		member := dynamoTeamUser{}
		if err := attributevalue.UnmarshalMap(item, &member); err != nil {
			return nil, err
		}

		users = append(users, TeamUser{
			TeamId: teamId,
			UserId: member.UserId,
			Role:   member.Role,
		})
	}

	return users, nil
}

func (store *DynamoStore) TeamList(ctx context.Context, userId string) ([]*Team, error) {
	uteamAv, err := attributevalue.MarshalMap(store.marshalUserTeam("", userId, ""))
	if err != nil {
		return nil, err
	}

	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		KeyConditionExpression: aws.String("pk = :pk AND begins_with(sk, :prefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":     uteamAv["pk"],
			":prefix": uteamAv["sk"],
		},
	})

	if err != nil {
		return nil, err
	}

	teams := []*Team{}
	for _, item := range out.Items {
		dteam := dynamoTeam{}
		if err := attributevalue.UnmarshalMap(item, &dteam); err != nil {
			return nil, err
		}

		teams = append(teams, dteam.unmarshal())
	}

	return teams, nil
}

// TeamDelete will remove all the users from a given team
func (store *DynamoStore) TeamDelete(ctx context.Context, teamId string) error {
	// Batch operations are max of 25 items, which means our chunk size must respect
	// the number of batch records and scale approprately
	const CHUNKSIZE = 25 / 2

	tuser, err := store.TeamListUsers(ctx, teamId)
	if err != nil {
		return err
	}

	for ; len(tuser) > 0; tuser = tuser[CHUNKSIZE:] {
		userIds := []string{}
		for _, item := range tuser[0:CHUNKSIZE] {
			userIds = append(userIds, item.UserId)
		}

		if err := store.TeamDeleteUsers(ctx, teamId, userIds...); err != nil {
			return err
		}
	}

	return nil
}

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
	Status         string                       `dynamodbav:"status"`
	ClosedStatus   string                       `dynamodbav:"closed_status"`
	Settings       map[string]map[string]string `dynamodbav:"settings"`
	AvatarUrl      *string                      `dynamodbav:"avatar_url,nullempty"`

	// For email address confirmation
	EmailVerifyNonce     []byte     `dynamodbav:"email_verify_nonce,nullempty"`
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
		Status:               user.Status.String(),
		ClosedStatus:         user.ClosedStatus.String(),
		Settings:             user.Settings,
		AvatarUrl:            user.AvatarUrl,
		EmailVerifyNonce:     user.EmailVerifyNonce,
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
		Status:               UserStatusFromString(obj.Status),
		ClosedStatus:         ClosedStatusFromString(obj.ClosedStatus),
		Settings:             obj.Settings,
		AvatarUrl:            obj.AvatarUrl,
		EmailVerifyNonce:     obj.EmailVerifyNonce,
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
		Pk:     store.makeUserEmailKey(strings.ToLower(user.Email)),
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
	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": av["pk"],
			"sk": av["sk"],
		},
	})

	if err != nil {
		return nil, err
	}
	if len(out.Item) == 0 {
		return nil, ErrorUserNotFound
	}

	user := dynamoUserEmail{}
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return store.GetById(ctx, user.UserID)
}

func (store *DynamoStore) CreateUser(ctx context.Context, user User) (*User, error) {
	user.ID = xid.New().String()

	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}
	userAv, err := attributevalue.MarshalMap(store.marshalUser(&user))
	if err != nil {
		return nil, err
	}
	emailAv, err := attributevalue.MarshalMap(store.marshalUserEmail(&user))
	if err != nil {
		return nil, err
	}

	// By default every user is a member of their own "default" team
	membership := TeamMember{
		UserId: user.ID,
		Status: TeamStatus_ACTIVE,
		Role:   "admin",
	}
	_, teamItems, err := store.teamCreateItems(ctx, "", membership)
	if err != nil {
		return nil, err
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

	return &user, err
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
						"sk": oldAv["sk"],
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

func (store *DynamoStore) marshalTeam(team *Team) *dynamoTeam {
	return &dynamoTeam{
		Pk:   store.makeTeamKey(team.TeamId),
		Sk:   "A",
		Type: "team",
		Id:   team.TeamId,
		Name: team.Name,
	}
}

func (team dynamoTeam) unmarshal() *Team {
	return &Team{
		TeamId: team.Id,
		Name:   team.Name,
	}
}

func (store *DynamoStore) teamCreateItems(ctx context.Context, name string, tuser ...TeamMember) (*Team, []types.WriteRequest, error) {
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

func (store *DynamoStore) TeamCreate(ctx context.Context, name string, tuser ...TeamMember) (*Team, error) {
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

func (store *DynamoStore) TeamList(ctx context.Context, userId string) ([]TeamMember, error) {
	_, uteam := store.marshalTeamMemberKey("", userId)
	uteamAv, err := attributevalue.MarshalMap(uteam)
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

	teams := []TeamMember{}
	for _, item := range out.Items {
		dteam := dynamoTeamMember{}
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

	tuser, err := store.TeamListMembers(ctx, teamId)
	if err != nil {
		return err
	}

	for ; len(tuser) > 0; tuser = tuser[CHUNKSIZE:] {
		userIds := []string{}
		for _, item := range tuser[0:CHUNKSIZE] {
			userIds = append(userIds, item.UserId)
		}

		if err := store.TeamDeleteMembers(ctx, teamId, userIds...); err != nil {
			return err
		}
	}

	return nil
}

// Team membership
//  this is stored as a pair of record in the forward and backwards direction (e.g. Team=>[]User  and User=>[]Team ])
// This avoids the GSI for the list, since it's typically not needed and we can batch write easy enough

type dynamoTeamMember struct {
	Pk       string `dynamodbav:"pk"`
	Sk       string `dynamodbav:"sk"`
	Type     string `dynamodbav:"type"`
	Status   string `dynamodbav:"status"`
	MemberId string `dynamodbav:"member_id"`
	UserId   string `dynamodbav:"user_id"`
	TeamId   string `dynamodbav:"team_id"`
	Role     string `dynamodbav:"role"`
}

func (store *DynamoStore) marshalTeamMemberKey(teamId string, userId string) (*dynamoTeamMember, *dynamoTeamMember) {
	teamUser := &dynamoTeamMember{
		Pk:     store.makeTeamKey(teamId),
		Sk:     store.makeUserKey(userId),
		Type:   "memberTeamUser",
		TeamId: teamId,
		UserId: userId,
	}
	userTeam := &dynamoTeamMember{
		Pk:     store.makeUserKey(userId),
		Sk:     store.makeTeamKey(teamId),
		Type:   "memberUserTeam",
		TeamId: teamId,
		UserId: userId,
	}

	return teamUser, userTeam
}

func (store *DynamoStore) marshalTeamUser(teamId string, member TeamMember) *dynamoTeamMember {
	value, _ := store.marshalTeamMemberKey(teamId, member.UserId)

	value.Status = member.Status.String()
	value.Role = member.Role

	return value
}

func (store *DynamoStore) marshalUserTeam(teamId string, member TeamMember) *dynamoTeamMember {
	_, value := store.marshalTeamMemberKey(teamId, member.UserId)

	value.Status = member.Status.String()
	value.Role = member.Role

	return value
}

func (item dynamoTeamMember) unmarshal() TeamMember {
	return TeamMember{
		MemberId: item.MemberId,
		UserId:   item.UserId,
		TeamId:   item.TeamId,
		Status:   TeamStatusFromString(item.Status),
		Role:     item.Role,
	}
}

func (store *DynamoStore) createUserItems(ctx context.Context, teamId string, tuser ...TeamMember) ([]types.WriteRequest, error) {
	items := []types.WriteRequest{}
	for _, item := range tuser {
		// We're using a computed identifier, such that if you're removed and re-added you restore all of
		//  the content that was linked to a specific user in a team
		item.MemberId = "member#" + teamId + "#" + item.UserId

		tuserAv, err := attributevalue.MarshalMap(store.marshalTeamUser(teamId, item))
		if err != nil {
			return nil, err
		}
		uteamAv, err := attributevalue.MarshalMap(store.marshalUserTeam(teamId, item))
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

func (store *DynamoStore) TeamAddMembers(ctx context.Context, tuser ...TeamMember) error {
	if len(tuser) == 0 {
		return nil
	}

	// Make sure they're all being added to the same team
	for _, item := range tuser {
		if item.TeamId != tuser[0].TeamId {
			return errors.New("team id mismatch")
		}
		if item.Status == TeamStatus_UNSET {
			return errors.New("missing status")
		}
	}
	if err := store.checkCreateTable(ctx); err != nil {
		return err
	}

	items, err := store.createUserItems(ctx, tuser[0].TeamId, tuser...)
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

func (store *DynamoStore) TeamAcceptInvite(ctx context.Context, teamId, userId string) error {
	tuser, uteam := store.marshalTeamMemberKey(teamId, userId)
	tuserAv, err := attributevalue.MarshalMap(tuser)
	if err != nil {
		return err
	}
	uteamAv, err := attributevalue.MarshalMap(uteam)
	if err != nil {
		return err
	}

	// We're using a transaction here just to make sure we always
	//  update both elements
	transaction := []types.TransactWriteItem{
		{
			Update: &types.Update{
				TableName: store.table,
				Key: map[string]types.AttributeValue{
					"pk": tuserAv["pk"],
					"sk": tuserAv["sk"],
				},
				UpdateExpression: aws.String("SET status = :status"),
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":status": &types.AttributeValueMemberS{
						Value: TeamStatus_ACTIVE.String(),
					},
				},
			},
		},
		{
			Update: &types.Update{
				TableName: store.table,
				Key: map[string]types.AttributeValue{
					"pk": uteamAv["pk"],
					"sk": uteamAv["sk"],
				},
				UpdateExpression: aws.String("SET status = :status"),
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":status": &types.AttributeValueMemberS{
						Value: TeamStatus_ACTIVE.String(),
					},
				},
			},
		},
	}

	_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transaction,
	})

	return err
}

func (store *DynamoStore) TeamDeleteMembers(ctx context.Context, teamId string, userIds ...string) error {
	items := []types.WriteRequest{}
	for _, userId := range userIds {
		tuser, uteam := store.marshalTeamMemberKey(teamId, userId)
		tuserAv, err := attributevalue.MarshalMap(tuser)
		if err != nil {
			return err
		}
		uteamAv, err := attributevalue.MarshalMap(uteam)
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

func (store *DynamoStore) TeamGetMember(ctx context.Context, teamId, userId string) (*TeamMember, error) {
	tuser, _ := store.marshalTeamMemberKey(teamId, userId)
	tuserAv, err := attributevalue.MarshalMap(tuser)
	if err != nil {
		return nil, err
	}

	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": tuserAv["pk"],
			"sk": tuserAv["sk"],
		},
	})

	if err != nil {
		return nil, err
	}
	if len(out.Item) == 0 {
		return nil, ErrorUserNotFound
	}

	dmember := dynamoTeamMember{}
	if err := attributevalue.UnmarshalMap(out.Item, &dmember); err != nil {
		return nil, err
	}

	member := dmember.unmarshal()

	return &member, nil
}

func (store *DynamoStore) TeamListMembers(ctx context.Context, teamId string) ([]TeamMember, error) {
	tuser, _ := store.marshalTeamMemberKey(teamId, "")
	tuserAv, err := attributevalue.MarshalMap(tuser)
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

	users := []TeamMember{}
	for _, item := range out.Items {
		member := dynamoTeamMember{}
		if err := attributevalue.UnmarshalMap(item, &member); err != nil {
			return nil, err
		}

		users = append(users, (member).unmarshal())
	}

	return users, nil
}

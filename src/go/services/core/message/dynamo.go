package message

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

// https://stackoverflow.com/questions/67933217/dynamo-db-erd-for-chat-application

// This table supports the following access patterns:
// fetch rooms for user (Query for PK=USER#userId SK beings_with ROOM)
// fetch message by ID (Get Item where PK=MSG#messageId SK=MSG#messageId

// I also created a Global Secondary Index named GSI1 on attributes GSI1PK and GSI1SK.
// fetch messages by room, sorted by creation date (Query GSI1 for PK=ROOM#roomId SK begins_with MSG#).
//   The messages will come back sorted by creation date/time since we're using KSUIDs!
// fetch all users by room (Query GSI1 for PK=ROOM#roomId SK begins_with USER)

type dynamoRoom struct {
	Pk      string `dynamodbav:"pk"`
	Sk      string `dynamodbav:"sk"`
	Type    string `dynamodbav:"type"`
	Name    string `dynamodbav:"name"`
	OrgId   string `dynamodbav:"org_id"`
	RoomId  string `dynamodbav:"room_id"`
	Private bool   `dynamodbav:"private"`
}
type dynamoMessage struct {
	Pk        string `dynamodbav:"pk"`
	Sk        string `dynamodbav:"sk"`
	Type      string `dynamodbav:"type"`
	MsgId     string `dynamodbav:"msg_id"`
	OrgId     string `dynamodbav:"org_id"`
	RoomId    string `dynamodbav:"room_id"`
	UserId    string `dynamodbav:"user_id"`
	Message   string `dynamodbav:"message"`
	CreatedAt string `dynamodbav:"created_at"`
	Gs1pk     string `dynamodbav:"gs1pk"`
	Gs1sk     string `dynamodbav:"gs1sk"`
}
type dynamoUser struct {
	Pk   string `dynamodbav:"pk"`
	Sk   string `dynamodbav:"sk"`
	Type string `dynamodbav:"type"`
}

type dynamoMembership struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Type   string `dynamodbav:"type"`
	Gs1pk  string `dynamodbav:"gs1pk"`
	Gs1sk  string `dynamodbav:"gs1sk"`
	UserId string `dynamodbav:"user_id"`
	OrgId  string `dynamodbav:"org_id"`
	RoomId string `dynamodbav:"room_id"`
}

func (room dynamoRoom) unmarshal() *Room {
	return &Room{
		ID:    room.RoomId,
		OrgId: room.OrgId,
		Name:  room.Name,
	}
}

func (msg dynamoMessage) unmarshal() *Message {
	return &Message{
		ID:     msg.MsgId,
		RoomId: msg.RoomId,
		UserId: msg.UserId,
		OrgId:  msg.OrgId,
		Text:   msg.Message,
	}
}

type dynamoStore struct {
	client      *dynamodb.Client
	table       *string
	index       *string
	tableExists bool
}

var _ MessageStore = (*dynamoStore)(nil)

type DynamoOption func(*dynamoStore)

func WithDynamoClient(client *dynamodb.Client) DynamoOption {
	return func(cfg *dynamoStore) {
		cfg.client = client
	}
}
func WithDynamoTable(value string) DynamoOption {
	return func(cfg *dynamoStore) {
		cfg.table = &value
	}
}

func NewDynamoStore(opts ...DynamoOption) *dynamoStore {
	state := dynamoStore{
		table: aws.String("app-messages"),
		index: aws.String("gs1"),
	}

	for _, opt := range opts {
		opt(&state)
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

func (*dynamoStore) makeOrgKey(orgId string) string {
	return "ORG#" + orgId
}
func (*dynamoStore) makeUserKey(orgId, userId string) string {
	return "USER#" + orgId + "#" + userId
}
func (*dynamoStore) makeRoomKey(orgId, roomId string) string {
	return "ROOM#" + orgId + "#" + roomId
}

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
			{
				AttributeName: aws.String("gs1pk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("gs1sk"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: store.index,
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("gs1pk"),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String("gs1sk"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{
					// NonKeyAttributes: []string{},
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
	})
	if err != nil {
		log.With(zap.Error(err)).Error("CreateTable failed")
	}
	store.tableExists = true

	log.Info("create complete")

	return err
}

func (store *dynamoStore) CreateRoom(ctx context.Context, orgId, userId string, name string) (*Room, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}

	// TODO Cannot create a duplicate named room

	roomId := xid.New().String()
	roomKey := store.makeRoomKey(orgId, roomId)
	data := dynamoRoom{
		Pk:      roomKey,
		Sk:      "A",
		Type:    "room",
		Private: true,
		OrgId:   orgId,
		RoomId:  roomId,
		Name:    name,
	}
	dataOrg := dynamoRoom{
		Pk:      store.makeOrgKey(orgId),
		Sk:      roomKey,
		Type:    "orgRoom",
		Private: true,
		OrgId:   orgId,
		RoomId:  roomId,
		Name:    name,
	}

	avRoom, err := attributevalue.MarshalMap(&data)
	if err != nil {
		return nil, err
	}
	avOrg, err := attributevalue.MarshalMap(&dataOrg)
	if err != nil {
		return nil, err
	}

	hash := md5.New()
	hash.Write([]byte(name))
	uniqueKey := "UNIQUE#ROOM#" + orgId + "#" + hex.EncodeToString(hash.Sum(nil))
	uniqueAv, err := attributevalue.MarshalMap(map[string]string{
		"pk": uniqueKey,
		"sk": uniqueKey,
	})
	if err != nil {
		return nil, err
	}

	// Rooms have unique names in an org
	_, err = store.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName: store.table,
					Item:      avRoom,
				},
			},
			{
				Put: &types.Put{
					TableName: store.table,
					Item:      avOrg,
				},
			},
			{
				Put: &types.Put{
					TableName:           store.table,
					Item:                uniqueAv,
					ConditionExpression: aws.String("attribute_not_exists(pk)"),
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return data.unmarshal(), nil
}

func (store *dynamoStore) ListRooms(ctx context.Context, orgId string, userId *string) ([]*Room, error) {
	if err := store.checkCreateTable(ctx); err != nil {
		return nil, err
	}

	// fetch rooms for user (Query for PK=USER#userId SK beings_with ROOM)
	roomKey := store.makeRoomKey(orgId, "")
	var out *dynamodb.QueryOutput
	var err error

	if userId != nil {
		userKey := store.makeUserKey(orgId, *userId)

		log := logger.FromContext(ctx)
		log.With(
			zap.String(":userId", userKey),
			zap.String(":prefix", roomKey),
		).Info("keys")

		out, err = store.client.Query(ctx, &dynamodb.QueryInput{
			TableName:              store.table,
			KeyConditionExpression: aws.String("pk = :userId AND begins_with(sk, :prefix)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":userId": &types.AttributeValueMemberS{Value: userKey},
				":prefix": &types.AttributeValueMemberS{Value: roomKey},
			},
		})

		if err != nil {
			return nil, err
		}
	} else {
		orgKey := store.makeOrgKey(orgId)

		// List all rooms for an org
		out, err = store.client.Query(ctx, &dynamodb.QueryInput{
			TableName:              store.table,
			KeyConditionExpression: aws.String("pk = :orgId AND begins_with(sk, :prefix)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":orgId":  &types.AttributeValueMemberS{Value: orgKey},
				":prefix": &types.AttributeValueMemberS{Value: roomKey},
			},
		})

		if err != nil {
			return nil, err
		}
	}

	rooms := []*Room{}
	for _, item := range out.Items {
		room := dynamoRoom{}
		if err := attributevalue.UnmarshalMap(item, &room); err != nil {
			return nil, err
		}

		rooms = append(rooms, room.unmarshal())
	}

	return rooms, nil
}

func (store *dynamoStore) Join(ctx context.Context, orgId, roomId string, userId string) error {
	userKey := store.makeUserKey(orgId, userId)
	roomKey := store.makeRoomKey(orgId, roomId)

	data := dynamoMembership{
		Pk:     userKey,
		Sk:     roomKey,
		Type:   "userRoom",
		Gs1pk:  roomKey,
		Gs1sk:  userKey,
		OrgId:  orgId,
		UserId: userId,
		RoomId: roomId,
	}

	av, err := attributevalue.MarshalMap(&data)
	if err != nil {
		return err
	}

	_, err = store.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: store.table,
		Item:      av,
	})

	if err != nil {
		return err
	}

	return nil
}

func (store *dynamoStore) Leave(ctx context.Context, orgId, roomId string, userId string) error {
	userKey := store.makeUserKey(orgId, userId)
	roomKey := store.makeRoomKey(orgId, roomId)

	_, err := store.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: userKey},
			"sk": &types.AttributeValueMemberS{Value: roomKey},
		},
	})

	return err
}

func (store *dynamoStore) Members(ctx context.Context, orgId, roomId string) ([]string, error) {
	// fetch all users by room (Query GSI1 for PK=ROOM#roomId SK begins_with USER)

	roomKey := store.makeRoomKey(orgId, roomId)
	userKey := store.makeUserKey(orgId, "")

	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		IndexName:              store.index,
		KeyConditionExpression: aws.String("gs1pk = :roomId AND begins_with(gs1sk, :prefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":roomId": &types.AttributeValueMemberS{Value: roomKey},
			":prefix": &types.AttributeValueMemberS{Value: userKey},
		},
	})

	if err != nil {
		return nil, err
	}

	users := []string{}
	for _, item := range out.Items {
		member := dynamoMembership{}
		if err := attributevalue.UnmarshalMap(item, &member); err != nil {
			return nil, err
		}

		users = append(users, member.UserId)
	}

	return users, nil
}

func (store *dynamoStore) ListMessages(ctx context.Context, orgId, roomId string) ([]*Message, error) {
	// fetch messages by room, sorted by creation date (Query GSI1 for PK=ROOM#roomId SK begins_with MSG#).
	// The messages will come back sorted by creation date/time since we're using KSUIDs!
	roomKey := store.makeRoomKey(orgId, roomId)

	out, err := store.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              store.table,
		IndexName:              store.index,
		KeyConditionExpression: aws.String("gs1pk = :roomId AND begins_with(gs1sk, :prefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":roomId": &types.AttributeValueMemberS{Value: roomKey},
			":prefix": &types.AttributeValueMemberS{Value: "MSG#"},
		},
	})

	if err != nil {
		return nil, err
	}

	messages := []*Message{}
	for _, item := range out.Items {
		msg := dynamoMessage{}
		if err := attributevalue.UnmarshalMap(item, &msg); err != nil {
			return nil, err
		}

		messages = append(messages, msg.unmarshal())
	}

	return messages, nil
}

func (store *dynamoStore) DeleteOne(ctx context.Context, orgId, roomId string, msgId string) error {
	msgKey := "MSG#" + msgId

	_, err := store.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: msgKey},
			"sk": &types.AttributeValueMemberS{Value: msgKey},
		},
	})

	return err
}

func (store *dynamoStore) CreateMessage(ctx context.Context, orgId, roomId string, message Message) (*Message, error) {
	// fetch message by ID (Get Item where PK=MSG#messageId SK=MSG#messageId
	msgId := xid.New().String()
	msgKey := "MSG#" + msgId
	roomKey := store.makeRoomKey(orgId, roomId)

	data := dynamoMessage{
		Pk:        msgKey,
		Sk:        msgKey,
		Type:      "message",
		MsgId:     msgId,
		OrgId:     orgId,
		RoomId:    message.RoomId,
		Message:   message.Text,
		UserId:    message.UserId,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		Gs1pk:     roomKey,
		Gs1sk:     msgKey,
	}

	av, err := attributevalue.MarshalMap(&data)
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

	return data.unmarshal(), nil
}

func (store *dynamoStore) GetMessage(ctx context.Context, orgId, roomId string, msgId string) (*Message, error) {
	msgKey := "MSG#" + msgId

	out, err := store.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: store.table,
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: msgKey},
			"sk": &types.AttributeValueMemberS{Value: msgKey},
		},
	})

	if err != nil || len(out.Item) == 0 {
		return nil, err
	}

	msg := dynamoMessage{}
	if err := attributevalue.UnmarshalMap(out.Item, &msg); err != nil {
		return nil, err
	}

	return (msg).unmarshal(), nil
}

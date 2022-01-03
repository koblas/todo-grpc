package user_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/koblas/grpc-todo/services/core/user"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestDynamo(t *testing.T) {
	if os.Getenv("AWS_PROFILE") != "" {
		suite.Run(t, new(DynamoBasicSuite))
	}
}

type DynamoBasicSuite struct {
	suite.Suite
	client    *dynamodb.Client
	tableName string
	store     user.UserStore
}

func (suite *DynamoBasicSuite) SetupSuite() {
	t := suite.T()
	cfg, err := config.LoadDefaultConfig(context.TODO())

	require.NoError(t, err, "Unable to load AWS")

	suite.client = dynamodb.NewFromConfig(cfg)
	suite.tableName = "test-user" + uuid.NewString()

	_, err = suite.client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		BillingMode: types.BillingModePayPerRequest,
		TableName:   aws.String(suite.tableName),

		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("pk"), AttributeType: types.ScalarAttributeTypeS},
			// {AttributeName: aws.String("sk"), AttributeType: types.ScalarAttributeTypeS},
			// {AttributeName: aws.String("user_id"), AttributeType: types.ScalarAttributeTypeS},
			// {AttributeName: aws.String("email_lc"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("pk"), KeyType: types.KeyTypeHash},
		},

		// GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
		// 	{
		// 		IndexName:  aws.String(suite.tableName + "-by-email"),
		// 		Projection: &types.Projection{ProjectionType: types.ProjectionTypeKeysOnly},
		// 		KeySchema: []types.KeySchemaElement{
		// 			{AttributeName: aws.String("email_lc"), KeyType: types.KeyTypeHash},
		// 		},
		// 	},
		// },
	})

	if err != nil {
		fmt.Print(err.Error())
	}
	require.NoError(suite.T(), err, "Unable to create table")

	var status *dynamodb.DescribeTableOutput
	for i := 0; i < 2*60 && status == nil || status.Table.TableStatus != types.TableStatusActive; i += 1 {
		time.Sleep(time.Millisecond * 500)

		status, err = suite.client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(suite.tableName),
		})
		require.NoError(suite.T(), err, "Unable to stat table")
	}

	require.Equal(t, status.Table.TableStatus, types.TableStatusActive)

	suite.store = user.NewUserDynamoStore(
		user.WithDynamoClient(suite.client),
		user.WithDynamoTable(suite.tableName),
	)
}

func (suite *DynamoBasicSuite) TearDownSuite() {
	_, err := suite.client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(suite.tableName),
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	require.NoError(suite.T(), err, "Unable to tear down table")
}

func (suite *DynamoBasicSuite) createUser(t *testing.T) user.User {
	user := user.User{
		ID:    uuid.NewString(),
		Email: faker.New().Internet().Email(),
	}
	err := suite.store.CreateUser(&user)
	require.NoError(t, err, "Create user error")

	time.Sleep(time.Millisecond * 1000)

	return user
}

func (suite *DynamoBasicSuite) TestByEmail() {
	t := suite.T()
	user := suite.createUser(t)

	// give the secondary index a chance to catchup
	// time.Sleep(2 * time.Second)

	u1, err := suite.store.GetByEmail("bad-email")
	require.NoError(t, err, "bad email failed")
	require.Nil(t, u1, "bad email - got user")

	u2, err := suite.store.GetByEmail(user.Email)
	require.NoError(t, err, "good email failed")
	require.NotNil(t, u2, "good email - user not found")
	require.Equal(t, u2.ID, user.ID, "IDs should match")
}

func (suite *DynamoBasicSuite) TestById() {
	t := suite.T()
	user := suite.createUser(t)

	u1, err := suite.store.GetById("bad-id")
	require.NoError(t, err, "bad-id returned error")
	require.Nil(t, u1, "bad-id returned object")

	u2, err := suite.store.GetById(user.ID)
	require.NoError(t, err, "good-id returned error")
	require.NotNil(t, u2, "good-id not found")
	require.Equal(t, u2.ID, user.ID, "IDs should match")
}

func (suite *DynamoBasicSuite) TestUpdateUser() {
	t := suite.T()
	u0 := suite.createUser(t)

	u0.Status = user.UserStatus(uuid.NewString())

	err := suite.store.UpdateUser(&u0)
	require.NoError(t, err, "update user 1")

	u, err := suite.store.GetById(u0.ID)
	require.NoError(t, err, "get by ID success")
	require.NotNil(t, u, "record found")

	require.Equal(t, u.ID, u0.ID, "IDs should match")
	require.Equal(t, u.Status, u0.Status, "status updated")
}

func (suite *DynamoBasicSuite) TestUpdateEmail() {
	t := suite.T()
	u0 := suite.createUser(t)

	oldEmail := u0.Email

	u0.Email = faker.New().Internet().Email()

	err := suite.store.UpdateUser(&u0)
	require.NoError(t, err, "update user 1")

	u, err := suite.store.GetById(u0.ID)
	require.NoError(t, err, "get by ID success")
	require.NotNil(t, u, "record found")
	require.Equal(t, u.ID, u0.ID, "IDs should match")

	u, err = suite.store.GetByEmail(u0.Email)
	require.NoError(t, err, "get by Email success")
	require.NotNil(t, u, "record found")
	require.Equal(t, u.ID, u0.ID, "IDs should match")

	u, err = suite.store.GetByEmail(oldEmail)
	require.NoError(t, err, "get by oldEmail success")
	require.Nil(t, u, "record not found")
}

func (suite *DynamoBasicSuite) TestUpdateDupEmail() {
	t := suite.T()
	u0 := suite.createUser(t)
	u1 := suite.createUser(t)

	u1.Email = u0.Email

	err := suite.store.UpdateUser(&u1)
	require.Error(t, err, "Update duplicate email")
}

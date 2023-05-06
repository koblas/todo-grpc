package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/go-faker/faker/v4"
	"github.com/gojuno/minimock/v3"
	eventv1 "github.com/koblas/grpc-todo/gen/core/eventbus/v1"
	"github.com/koblas/grpc-todo/gen/core/eventbus/v1/eventbusv1connect"
	userv1 "github.com/koblas/grpc-todo/gen/core/user/v1"
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	"github.com/koblas/grpc-todo/services/core/user"
	"github.com/rs/xid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestHandler(t *testing.T) {
	suite.Run(t, new(UserSuite))
	suite.Run(t, new(OAuthSuite))
}

type sharedSuite struct {
	suite.Suite
}

type UserSuite struct {
	sharedSuite
}

type OAuthSuite struct {
	sharedSuite
}

func (*sharedSuite) buildServer(t *testing.T) (*user.UserServer, *eventbusv1connect.UserEventbusServiceClientMock) {
	mc := minimock.NewController(t)
	producer := eventbusv1connect.NewUserEventbusServiceClientMock(mc)
	opts := []user.Option{
		user.WithProducer(producer),
	}

	if storeAddr := os.Getenv("DYNAMO_ADDR"); storeAddr == "" {
		opts = append(opts,
			user.WithStore(user.NewUserMemoryStore()),
		)
	} else {
		t.Logf("Using DynamoDB at %s", storeAddr)
		tableName := "test-user-" + xid.New().String()

		opts = append(opts,
			user.WithStore(
				user.NewDynamoStore(
					user.WithDynamoClient(
						awsutil.LocalDynamoClient(storeAddr),
					),
					user.WithDynamoTable(tableName),
				),
			))
	}

	return user.NewUserServer(opts...), producer
}

func (suite *UserSuite) TestCreateBasic() {
	server, mock := suite.buildServer(suite.T())

	ctx := context.TODO()
	// ctx = logger.ToContext(ctx, logger.NewZap(logger.LevelDebug))

	mock.UserChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil)
	mock.SecurityRegisterTokenMock.Return(connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil)

	type testcase struct {
		Status     userv1.UserStatus
		TokenCalls int
	}

	statues := []testcase{
		{Status: userv1.UserStatus_USER_STATUS_ACTIVE, TokenCalls: 0},
		{Status: userv1.UserStatus_USER_STATUS_REGISTERED, TokenCalls: 1},
		{Status: userv1.UserStatus_USER_STATUS_INVITED, TokenCalls: 1},
	}

	for _, testcase := range statues {
		suite.T().Run("status="+testcase.Status.String(), func(t *testing.T) {
			startTokenCalls := mock.SecurityRegisterTokenAfterCounter()

			request := userv1.CreateRequest{
				Name:   faker.Name(),
				Email:  faker.Email(),
				Status: testcase.Status,
			}
			user, err := server.Create(ctx, connect.NewRequest(&request))

			tokenCalls := mock.SecurityRegisterTokenAfterCounter() - startTokenCalls

			require.NoError(t, err, "call failed")
			require.NotNil(t, user.Msg.User, "no user returned")
			require.Equal(t, request.Name, user.Msg.User.Name, "name mismatch")
			require.Equal(t, testcase.Status, user.Msg.User.Status, "status mismatch")
			require.EqualValues(t, testcase.TokenCalls, tokenCalls, "token called")

			// Find by Email test
			findRequest := userv1.FindByRequest{
				FindBy: &userv1.FindBy{
					Email: request.Email,
				},
			}
			found, err := server.FindBy(ctx, connect.NewRequest(&findRequest))

			require.NoError(t, err, "findEmail: call failed")
			require.NotNil(t, found.Msg.User, "findEmail: no user returned")

			// Find by ID test
			findRequest = userv1.FindByRequest{
				FindBy: &userv1.FindBy{
					UserId: user.Msg.User.Id,
				},
			}
			found, err = server.FindBy(ctx, connect.NewRequest(&findRequest))

			require.NoError(t, err, "findId: call failed")
			require.NotNil(t, found.Msg.User, "findId: no user returned")
		})
	}
}

func (suite *UserSuite) TestValidate() {
	server, mock := suite.buildServer(suite.T())

	ctx := context.TODO()
	// ctx = logger.ToContext(ctx, logger.NewZap(logger.LevelDebug))

	forgotToken := ""

	mock.UserChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil)
	mock.SecurityRegisterTokenMock.Set(func(ctx context.Context, pp1 *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityRegisterTokenResponse], error) {
		decoder := key_manager.NewSecureClear()
		token, err := protoutil.SecureValueDecode(decoder, pp1.Msg.Token)
		forgotToken = token
		return nil, err
	})

	request := userv1.CreateRequest{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Status:   userv1.UserStatus_USER_STATUS_REGISTERED,
		Password: faker.Password(),
	}
	user, err := server.Create(ctx, connect.NewRequest(&request))
	require.NoError(suite.T(), err, "create: call failed")

	_, err = server.VerificationVerify(ctx, connect.NewRequest(&userv1.VerificationVerifyRequest{
		Verification: &userv1.Verification{
			UserId: user.Msg.User.Id,
			Token:  forgotToken,
		},
	}))

	require.NoError(suite.T(), err, "verificationVerify: call failed")

	// Check that the token is removed
	_, err = server.VerificationVerify(ctx, connect.NewRequest(&userv1.VerificationVerifyRequest{
		Verification: &userv1.Verification{
			UserId: user.Msg.User.Id,
			Token:  forgotToken,
		},
	}))

	require.Error(suite.T(), err, "verificationVerify2: call failed")

	// Check that the user is update to ACTIVE
	findResponse, err := server.FindBy(ctx, connect.NewRequest(&userv1.FindByRequest{
		FindBy: &userv1.FindBy{
			UserId: user.Msg.User.Id,
		},
	}))

	require.NoError(suite.T(), err, "findBy: call failed")
	require.Equal(suite.T(), userv1.UserStatus_USER_STATUS_ACTIVE, findResponse.Msg.User.Status, "findBy: user wrong state")
}

func (suite *UserSuite) TestForgotBasic() {
	server, mock := suite.buildServer(suite.T())

	ctx := context.TODO()
	// ctx = logger.ToContext(ctx, logger.NewZap(logger.LevelDebug))

	forgotToken := ""

	mock.UserChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil)
	mock.SecurityRegisterTokenMock.Return(connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil)
	mock.SecurityForgotRequestMock.Set(func(ctx context.Context, pp1 *connect.Request[userv1.UserSecurityEvent]) (*connect.Response[eventv1.UserEventbusSecurityForgotRequestResponse], error) {
		decoder := key_manager.NewSecureClear()
		token, err := protoutil.SecureValueDecode(decoder, pp1.Msg.Token)
		forgotToken = token
		return nil, err
	})
	mock.SecurityPasswordChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusSecurityPasswordChangeResponse{}), nil)

	request := userv1.CreateRequest{
		Name:   faker.Name(),
		Email:  faker.Email(),
		Status: userv1.UserStatus_USER_STATUS_REGISTERED,
	}
	user, err := server.Create(ctx, connect.NewRequest(&request))
	require.NoError(suite.T(), err, "create: call failed")

	_, err = server.ForgotSend(ctx, connect.NewRequest(&userv1.ForgotSendRequest{
		FindBy: &userv1.FindBy{
			Email: user.Msg.User.Email,
		},
	}))

	require.NoError(suite.T(), err, "forgotSend: call failed")
	require.NotEmpty(suite.T(), forgotToken, "forgotSend: password not captured")

	_, err = server.ForgotVerify(ctx, connect.NewRequest(&userv1.ForgotVerifyRequest{
		Verification: &userv1.Verification{
			UserId: user.Msg.User.Id,
			Token:  forgotToken,
		},
	}))

	require.NoError(suite.T(), err, "forgotVerify: call failed")

	_, err = server.ForgotUpdate(ctx, connect.NewRequest(&userv1.ForgotUpdateRequest{
		Verification: &userv1.Verification{
			UserId:   user.Msg.User.Id,
			Token:    forgotToken,
			Password: faker.Password(),
		},
	}))

	require.NoError(suite.T(), err, "forgotUpdate: call failed")

	// Check to make sure the forgot token is cleared
	_, err = server.ForgotVerify(ctx, connect.NewRequest(&userv1.ForgotVerifyRequest{
		Verification: &userv1.Verification{
			UserId: user.Msg.User.Id,
			Token:  forgotToken,
		},
	}))

	require.Error(suite.T(), err, "forgotVerify2: call failed")

	// 1 for the create, 1 for the update
	require.EqualValues(suite.T(), 2, len(mock.UserChangeMock.Calls()), "user update")
	require.EqualValues(suite.T(), 1, len(mock.SecurityPasswordChangeMock.Calls()), "password change event")
}

func (suite *OAuthSuite) TestNotFound() {
	// This tests the calls that UpsertUser does

	// FindBy {provider, provider.id}
	//  -> if found, then a user exists
	//  -> not found, then create
	// Extract email from OAuth information
	// FindBy {email}
	//  -> user not found do a Create{ email, name, ACTIVE }
	//  -> user found if status == REGISTERED -> Update { ACTIVE }
	// AuthAssociate { UserId, Provider, ProviderId }

	ctx := context.TODO()

	provider := "testprovider"
	providerId := faker.UUIDDigit()
	emailAddr := faker.Email()

	server, mock := suite.buildServer(suite.T())

	mock.UserChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil)
	mock.SecurityRegisterTokenMock.Return(connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil)

	findByRequest := userv1.FindByRequest{
		FindBy: &userv1.FindBy{
			Auth: &userv1.AuthInfo{
				Provider:   provider,
				ProviderId: providerId,
			},
		},
	}

	_, err := server.FindBy(ctx, connect.NewRequest(&findByRequest))
	require.Error(suite.T(), err, "findBy: call failed")
	require.EqualValues(suite.T(), connect.CodeNotFound, connect.CodeOf(err), "findBy: invalid error")

	request := userv1.CreateRequest{
		Name:   faker.Name(),
		Email:  emailAddr,
		Status: userv1.UserStatus_USER_STATUS_ACTIVE,
	}
	user, err := server.Create(ctx, connect.NewRequest(&request))
	require.NoError(suite.T(), err, "create: call failed")

	auth, err := server.AuthAssociate(ctx, connect.NewRequest(&userv1.AuthAssociateRequest{
		UserId: user.Msg.User.Id,
		Auth: &userv1.AuthInfo{
			Provider:   provider,
			ProviderId: providerId,
		},
	}))
	require.NoError(suite.T(), err, "create: call failed")
	require.EqualValues(suite.T(), auth.Msg.UserId, user.Msg.User.Id, "associate: value mismatch")

	found, err := server.FindBy(ctx, connect.NewRequest(&findByRequest))
	require.NoError(suite.T(), err, "findBy2: user not found")
	require.EqualValues(suite.T(), user.Msg.User.Id, found.Msg.User.Id, "findBy2: mismatch")
}

func (suite *OAuthSuite) TestFound() {
	// This tests the calls that UpsertUser does

	// FindBy {provider, provider.id}
	//  -> if found, then a user exists
	//  -> not found, then create
	// Extract email from OAuth information
	// FindBy {email}
	//  -> user not found do a Create{ email, name, ACTIVE }
	//  -> user found if status == REGISTERED -> Update { ACTIVE }
	// AuthAssociate { UserId, Provider, ProviderId }

	ctx := context.TODO()

	server, mock := suite.buildServer(suite.T())

	mock.UserChangeMock.Return(connect.NewResponse(&eventv1.UserEventbusUserChangeResponse{}), nil)
	mock.SecurityRegisterTokenMock.Return(connect.NewResponse(&eventv1.UserEventbusSecurityRegisterTokenResponse{}), nil)

	statues := []struct {
		Status     userv1.UserStatus
		TokenCalls int
	}{
		{Status: userv1.UserStatus_USER_STATUS_ACTIVE, TokenCalls: 0},
		{Status: userv1.UserStatus_USER_STATUS_REGISTERED, TokenCalls: 1},
		{Status: userv1.UserStatus_USER_STATUS_INVITED, TokenCalls: 1},
	}

	for _, testcase := range statues {
		suite.T().Run("status="+testcase.Status.String(), func(t *testing.T) {
			provider := "testprovider"
			providerId := faker.UUIDDigit()

			request := userv1.CreateRequest{
				Name:   faker.Name(),
				Email:  faker.Email(),
				Status: testcase.Status,
			}
			user, err := server.Create(ctx, connect.NewRequest(&request))
			require.NoError(suite.T(), err, "create: call failed")

			findByRequest := userv1.FindByRequest{
				FindBy: &userv1.FindBy{
					Auth: &userv1.AuthInfo{
						Provider:   provider,
						ProviderId: providerId,
					},
				},
			}
			_, err = server.FindBy(ctx, connect.NewRequest(&findByRequest))
			require.Error(suite.T(), err, "findBy: call failed")
			require.EqualValues(suite.T(), connect.CodeNotFound, connect.CodeOf(err), "findBy: invalid error")

			status := userv1.UserStatus_USER_STATUS_ACTIVE
			_, err = server.Update(ctx, connect.NewRequest(&userv1.UpdateRequest{
				UserId: user.Msg.User.Id,
				Status: &status,
			}))
			require.NoError(suite.T(), err, "create: call failed")

			auth, err := server.AuthAssociate(ctx, connect.NewRequest(&userv1.AuthAssociateRequest{
				UserId: user.Msg.User.Id,
				Auth: &userv1.AuthInfo{
					Provider:   provider,
					ProviderId: providerId,
				},
			}))
			require.NoError(suite.T(), err, "create: call failed")
			require.EqualValues(suite.T(), auth.Msg.UserId, user.Msg.User.Id, "associate: value mismatch")

			found, err := server.FindBy(ctx, connect.NewRequest(&findByRequest))
			require.NoError(suite.T(), err, "findBy2: user not found")
			require.EqualValues(suite.T(), user.Msg.User.Id, found.Msg.User.Id, "findBy2: mismatch")

			require.EqualValues(suite.T(), userv1.UserStatus_USER_STATUS_ACTIVE, found.Msg.User.Status, "findBy2: mismatch")
		})
	}
}

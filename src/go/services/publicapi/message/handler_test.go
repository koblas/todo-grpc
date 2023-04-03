package message_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/go-faker/faker/v4"
	"github.com/gojuno/minimock/v3"
	apiv1 "github.com/koblas/grpc-todo/gen/api/message/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/message/v1"
	"github.com/koblas/grpc-todo/gen/core/message/v1/messagev1connect"
	"github.com/koblas/grpc-todo/services/publicapi/message"
	"github.com/rs/xid"

	// "github.com/rs/xid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userIdProvider struct {
	userId string
}

func (p *userIdProvider) GetUserId(ctx context.Context) (string, error) {
	if p.userId == "" {
		return "", connect.NewError(connect.CodeUnauthenticated, nil)
	}
	return p.userId, nil
}

func buildServer(t *testing.T, userId string) (*message.MessageServer, *messagev1connect.MessageServiceClientMock) {
	mc := minimock.NewController(t)
	mocked := messagev1connect.NewMessageServiceClientMock(mc)

	return message.NewMessageServer(
		message.WithMessageService(mocked),
		message.WithGetUserId(&userIdProvider{userId}),
	), mocked
}

type MessageAddSuite struct {
	suite.Suite
}

func TestMessage(t *testing.T) {
	suite.Run(t, new(MessageAddSuite))
	// suite.Run(t, new(MessageListSuite))
	// suite.Run(t, new(MessageDeleteSuite))
}

func (suite *MessageAddSuite) TestMessageAddSmoke() {
	server, mock := buildServer(suite.T(), faker.UUIDHyphenated())

	mock.AddMock.Set(func(ctx context.Context, msg *connect.Request[corev1.AddRequest]) (*connect.Response[corev1.AddResponse], error) {
		return connect.NewResponse(&corev1.AddResponse{
			Message: &corev1.MessageItem{
				Text: msg.Msg.Text,
				Id:   xid.New().String(),
			},
		}), nil
	})

	text := faker.UUIDHyphenated()
	req := apiv1.MsgCreateRequest{
		Text: text + "test",
	}
	msg, err := server.MsgCreate(context.TODO(), connect.NewRequest(
		&req,
	))

	require.Nil(suite.T(), err, "should not return error")
	require.NotEmpty(suite.T(), msg.Msg.Message.Id, "Expected ID")
	require.Equal(suite.T(), req.Text, msg.Msg.Message.Text, "Expected task to match response task")
}

func (suite *MessageAddSuite) TestMessageAddNoUser() {
	server, mock := buildServer(suite.T(), "")
	mock.AddMock.Return(nil, errors.New("unknown"))

	_, err := server.MsgCreate(context.TODO(), connect.NewRequest(
		&apiv1.MsgCreateRequest{
			Text: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
	require.Equal(suite.T(), connect.CodeUnauthenticated, connect.CodeOf(err), "require authentication")
}

func (suite *MessageAddSuite) TestMessageAddError() {
	server, mock := buildServer(suite.T(), faker.UUIDHyphenated())
	mock.AddMock.Return(nil, errors.New("unknown"))

	_, err := server.MsgCreate(context.TODO(), connect.NewRequest(
		&apiv1.MsgCreateRequest{
			Text: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
}

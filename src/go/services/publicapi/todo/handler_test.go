package todo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/go-faker/faker/v4"
	"github.com/gojuno/minimock/v3"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/services/publicapi/todo"
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

func buildServer(client corev1connect.TodoServiceClient, userId string) *todo.TodoServer {
	return todo.NewTodoServer(
		todo.WithTodoService(client),
		todo.WithGetUserId(&userIdProvider{userId}),
	)
}

type TodoAddSuite struct {
	suite.Suite
}

func TestTodo(t *testing.T) {
	suite.Run(t, new(TodoAddSuite))
	// suite.Run(t, new(TodoListSuite))
	// suite.Run(t, new(TodoDeleteSuite))
}

func (suite *TodoAddSuite) TestTodoAddSmoke() {
	mc := minimock.NewController(suite.T())
	client := corev1connect.NewTodoServiceClientMock(mc)

	task := faker.UUIDHyphenated()

	client.TodoAddMock.Set(func(ctx context.Context, msg *connect.Request[corev1.TodoAddRequest]) (*connect.Response[corev1.TodoAddResponse], error) {
		return connect.NewResponse(&corev1.TodoAddResponse{
			Todo: &corev1.TodoObject{
				Task: msg.Msg.Task,
				Id:   xid.New().String(),
			},
		}), nil
	})

	server := buildServer(client, faker.UUIDHyphenated())

	req := apiv1.TodoAddRequest{
		Task: task + "test",
	}
	msg, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&req,
	))

	require.Nil(suite.T(), err, "should not return error")
	require.NotEmpty(suite.T(), msg.Msg.Todo.Id, "Expected ID")
	require.Equal(suite.T(), req.Task, msg.Msg.Todo.Task, "Expected task to match response task")
}

func (suite *TodoAddSuite) TestTodoAddNoUser() {
	mc := minimock.NewController(suite.T())
	client := corev1connect.NewTodoServiceClientMock(mc)

	client.TodoAddMock.Return(nil, errors.New("unknown"))

	server := buildServer(client, "")

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&apiv1.TodoAddRequest{
			Task: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
	require.Equal(suite.T(), connect.CodeUnauthenticated, connect.CodeOf(err), "require authentication")
}

func (suite *TodoAddSuite) TestTodoAddError() {
	mc := minimock.NewController(suite.T())
	client := corev1connect.NewTodoServiceClientMock(mc)

	client.TodoAddMock.Return(nil, errors.New("unknown"))

	server := buildServer(client, faker.UUIDHyphenated())

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&apiv1.TodoAddRequest{
			Task: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
}

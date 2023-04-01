package todo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/go-faker/faker/v4"
	"github.com/gojuno/minimock/v3"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	todov1 "github.com/koblas/grpc-todo/gen/core/todo/v1"
	"github.com/koblas/grpc-todo/gen/core/todo/v1/todov1connect"
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

func buildServer(t *testing.T, userId string) (*todo.TodoServer, *todov1connect.TodoServiceClientMock) {
	mc := minimock.NewController(t)
	mocked := todov1connect.NewTodoServiceClientMock(mc)

	return todo.NewTodoServer(
		todo.WithTodoService(mocked),
		todo.WithGetUserId(&userIdProvider{userId}),
	), mocked
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
	server, mock := buildServer(suite.T(), faker.UUIDHyphenated())

	mock.TodoAddMock.Set(func(ctx context.Context, msg *connect.Request[todov1.TodoAddRequest]) (*connect.Response[todov1.TodoAddResponse], error) {
		return connect.NewResponse(&todov1.TodoAddResponse{
			Todo: &todov1.TodoObject{
				Task: msg.Msg.Task,
				Id:   xid.New().String(),
			},
		}), nil
	})

	task := faker.UUIDHyphenated()
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
	server, mock := buildServer(suite.T(), "")
	mock.TodoAddMock.Return(nil, errors.New("unknown"))

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&apiv1.TodoAddRequest{
			Task: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
	require.Equal(suite.T(), connect.CodeUnauthenticated, connect.CodeOf(err), "require authentication")
}

func (suite *TodoAddSuite) TestTodoAddError() {
	server, mock := buildServer(suite.T(), faker.UUIDHyphenated())
	mock.TodoAddMock.Return(nil, errors.New("unknown"))

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&apiv1.TodoAddRequest{
			Task: faker.UUIDHyphenated(),
		},
	))

	require.NotNil(suite.T(), err, "return error")
}

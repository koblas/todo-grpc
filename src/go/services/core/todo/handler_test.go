package todo_test

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/services/core/todo"
	"github.com/rs/xid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func buildServer() (*todo.TodoServer, *todoEventbus) {
	bus := &todoEventbus{}

	return todo.NewTodoServer(
		todo.WithTodoStore(todo.NewTodoMemoryStore()),
		todo.WithProducer(bus),
	), bus
}

type todoEventbus struct {
	counter int
}

func (bus *todoEventbus) TodoChange(context.Context, *connect.Request[corev1.TodoChangeEvent]) (*connect.Response[corev1.TodoEventbusTodoChangeResponse], error) {
	bus.counter += 1
	return connect.NewResponse(&corev1.TodoEventbusTodoChangeResponse{}), nil
}

type TodoAddSuite struct {
	suite.Suite
}
type TodoListSuite struct {
	suite.Suite
}
type TodoDeleteSuite struct {
	suite.Suite
}

func TestTodo(t *testing.T) {
	suite.Run(t, new(TodoAddSuite))
	suite.Run(t, new(TodoListSuite))
	suite.Run(t, new(TodoDeleteSuite))
}

func (suite *TodoAddSuite) TestTodoNoUser() {
	server, _ := buildServer()

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			Task: "test",
		},
	))

	require.NotNil(suite.T(), err, "should return error")
	require.Equal(suite.T(), connect.CodeOf(err), connect.CodeInvalidArgument, "error response")
}

func (suite *TodoAddSuite) TestTodoText() {
	server, _ := buildServer()

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: "1234",
			Task:   "",
		},
	))

	require.NotNil(suite.T(), err, "should return error")
	require.Equal(suite.T(), connect.CodeOf(err), connect.CodeInvalidArgument, "error response")
}

func (suite *TodoAddSuite) TestTodoAdd() {
	server, bus := buildServer()

	res, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: "1234",
			Task:   "test",
		},
	))

	require.Nil(suite.T(), err, "no error")
	require.NotEmpty(suite.T(), res.Msg.Todo.Id, "expect id")
	require.Equal(suite.T(), 1, bus.counter, "event publish counts")
}

// List tests

func (suite *TodoListSuite) TestTodoAdd() {
	// Note: this test is close to a store test
	//  not a API test..
	server, _ := buildServer()
	user1 := xid.New().String()
	user2 := xid.New().String()

	_, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user1,
			Task:   "test",
		},
	))
	require.Nil(suite.T(), err, "no error")
	_, err = server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user1,
			Task:   "test",
		},
	))
	require.Nil(suite.T(), err, "no error")
	_, err = server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user2,
			Task:   "test",
		},
	))
	require.Nil(suite.T(), err, "no error")

	res, err := server.TodoList(context.TODO(), connect.NewRequest(&corev1.TodoListRequest{
		UserId: user1,
	}))
	require.Nil(suite.T(), err, "no error")
	require.Equal(suite.T(), 2, len(res.Msg.Todos), "user1")

	res, err = server.TodoList(context.TODO(), connect.NewRequest(&corev1.TodoListRequest{
		UserId: user2,
	}))
	require.Nil(suite.T(), err, "no error")
	require.Equal(suite.T(), 1, len(res.Msg.Todos), "user1")
}

func (suite *TodoDeleteSuite) TestTodoDelete() {
	// Note: this test is close to a store test
	//  not a API test..
	server, bus := buildServer()
	user1 := xid.New().String()
	user2 := xid.New().String()

	radd, err := server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user1,
			Task:   "test",
		},
	))
	id1 := radd.Msg.Todo.Id
	require.Nil(suite.T(), err, "no error")
	radd, err = server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user1,
			Task:   "test",
		},
	))
	id2 := radd.Msg.Todo.Id
	require.Nil(suite.T(), err, "no error")
	_, err = server.TodoAdd(context.TODO(), connect.NewRequest(
		&corev1.TodoAddRequest{
			UserId: user2,
			Task:   "test",
		},
	))
	// id3 := radd.Msg.Todo.Id
	require.Nil(suite.T(), err, "no error")

	_, err = server.TodoDelete(context.TODO(), connect.NewRequest(&corev1.TodoDeleteRequest{
		UserId: user1,
		Id:     id1,
	}))
	require.Nil(suite.T(), err, "no error")

	resl, err := server.TodoList(context.TODO(), connect.NewRequest(&corev1.TodoListRequest{
		UserId: user1,
	}))
	require.Nil(suite.T(), err, "no error")
	require.Equal(suite.T(), 1, len(resl.Msg.Todos), "count is right")
	require.Equal(suite.T(), id2, resl.Msg.Todos[0].Id, "check ids")

	resl, err = server.TodoList(context.TODO(), connect.NewRequest(&corev1.TodoListRequest{
		UserId: user2,
	}))
	require.Nil(suite.T(), err, "no error")
	require.Equal(suite.T(), 1, len(resl.Msg.Todos), "count is right")

	// 3 adds 1 delete
	require.Equal(suite.T(), 4, bus.counter, "event publish counts")
}

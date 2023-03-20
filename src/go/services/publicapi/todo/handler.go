package todo

import (
	"errors"

	"github.com/bufbuild/connect-go"
	apiv1 "github.com/koblas/grpc-todo/gen/api/v1"
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/gen/core/v1/corev1connect"
	"github.com/koblas/grpc-todo/pkg/bufcutil"
	"github.com/koblas/grpc-todo/pkg/interceptors"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TodoServer struct {
	todos      corev1connect.TodoServiceClient
	userHelper interceptors.UserIdFromContext
}

type Option func(*TodoServer)

func WithTodoService(client corev1connect.TodoServiceClient) Option {
	return func(svr *TodoServer) {
		svr.todos = client
	}
}

func WithGetUserId(helper interceptors.UserIdFromContext) Option {
	return func(svr *TodoServer) {
		svr.userHelper = helper
	}
}

func NewTodoServer(opts ...Option) *TodoServer {
	svr := TodoServer{}

	for _, opt := range opts {
		opt(&svr)
	}

	if svr.todos == nil {
		panic("No backend todo service")
	}
	if svr.userHelper == nil {
		panic("no user helper provided")
	}

	return &svr
}

// SayHello generates response to a Ping request
func (svc *TodoServer) TodoAdd(ctx context.Context, newTodo *connect.Request[apiv1.TodoAddRequest]) (*connect.Response[apiv1.TodoAddResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("AddTodo BEGIN")
	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	task, err := svc.todos.TodoAdd(ctx, connect.NewRequest(&corev1.TodoAddRequest{
		UserId: userId, // TODO
		Task:   newTodo.Msg.Task,
	}))

	if err != nil {
		log.With(zap.Error(err)).Error("Unable to create todo")
		return nil, bufcutil.InternalError(err)
	}
	log.With("task", newTodo.Msg.Task).Info("Received new task")

	return connect.NewResponse(&apiv1.TodoAddResponse{
		Todo: &apiv1.TodoObject{
			Id:   task.Msg.Todo.Id,
			Task: task.Msg.Todo.Task,
		},
	}), nil
}

func (svc *TodoServer) TodoList(ctx context.Context, _ *connect.Request[apiv1.TodoListRequest]) (*connect.Response[apiv1.TodoListResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("GetTodos BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	out, err := svc.todos.TodoList(ctx, connect.NewRequest(&corev1.TodoListRequest{
		UserId: userId,
	}))

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	todos := []*apiv1.TodoObject{}
	if out != nil {
		for _, item := range out.Msg.Todos {
			todos = append(todos, &apiv1.TodoObject{
				Id:   item.Id,
				Task: item.Task,
			})
		}
	}

	return connect.NewResponse(&apiv1.TodoListResponse{Todos: todos}), nil
}

func (svc *TodoServer) TodoDelete(ctx context.Context, delTodo *connect.Request[apiv1.TodoDeleteRequest]) (*connect.Response[apiv1.TodoDeleteResponse], error) {
	log := logger.FromContext(ctx)
	log.Info("DeleteTodo BEGIN")

	userId, err := svc.userHelper.GetUserId(ctx)
	if err != nil {
		log.With("error", err).Info("No user id found")
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing userid"))
	}

	_, err = svc.todos.TodoDelete(ctx, connect.NewRequest(&corev1.TodoDeleteRequest{
		UserId: userId,
		Id:     delTodo.Msg.Id,
	}))

	if err != nil {
		return nil, bufcutil.InternalError(err)
	}

	return connect.NewResponse(&apiv1.TodoDeleteResponse{Message: "success"}), nil
}

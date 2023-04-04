// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/koblas/grpc-todo/gen/api/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "api.user.v1.UserService"
)

// UserServiceClient is a client for the api.user.v1.UserService service.
type UserServiceClient interface {
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
}

// NewUserServiceClient constructs a client for the api.user.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+"/api.user.v1.UserService/get_user",
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+"/api.user.v1.UserService/update_user",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUser    *connect_go.Client[v1.GetUserRequest, v1.GetUserResponse]
	updateUser *connect_go.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
}

// GetUser calls api.user.v1.UserService.get_user.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// UpdateUser calls api.user.v1.UserService.update_user.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the api.user.v1.UserService service.
type UserServiceHandler interface {
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.user.v1.UserService/get_user", connect_go.NewUnaryHandler(
		"/api.user.v1.UserService/get_user",
		svc.GetUser,
		opts...,
	))
	mux.Handle("/api.user.v1.UserService/update_user", connect_go.NewUnaryHandler(
		"/api.user.v1.UserService/update_user",
		svc.UpdateUser,
		opts...,
	))
	return "/api.user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.UserService.get_user is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.UserService.update_user is not implemented"))
}
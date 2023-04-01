// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: core/message/v1/message.proto

package messagev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/koblas/grpc-todo/gen/core/message/v1"
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
	// MessageServiceName is the fully-qualified name of the MessageService service.
	MessageServiceName = "core.message.v1.MessageService"
)

// MessageServiceClient is a client for the core.message.v1.MessageService service.
type MessageServiceClient interface {
	Add(context.Context, *connect_go.Request[v1.AddRequest]) (*connect_go.Response[v1.AddResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
}

// NewMessageServiceClient constructs a client for the core.message.v1.MessageService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMessageServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MessageServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &messageServiceClient{
		add: connect_go.NewClient[v1.AddRequest, v1.AddResponse](
			httpClient,
			baseURL+"/core.message.v1.MessageService/Add",
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+"/core.message.v1.MessageService/Delete",
			opts...,
		),
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/core.message.v1.MessageService/List",
			opts...,
		),
	}
}

// messageServiceClient implements MessageServiceClient.
type messageServiceClient struct {
	add    *connect_go.Client[v1.AddRequest, v1.AddResponse]
	delete *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
	list   *connect_go.Client[v1.ListRequest, v1.ListResponse]
}

// Add calls core.message.v1.MessageService.Add.
func (c *messageServiceClient) Add(ctx context.Context, req *connect_go.Request[v1.AddRequest]) (*connect_go.Response[v1.AddResponse], error) {
	return c.add.CallUnary(ctx, req)
}

// Delete calls core.message.v1.MessageService.Delete.
func (c *messageServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// List calls core.message.v1.MessageService.List.
func (c *messageServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// MessageServiceHandler is an implementation of the core.message.v1.MessageService service.
type MessageServiceHandler interface {
	Add(context.Context, *connect_go.Request[v1.AddRequest]) (*connect_go.Response[v1.AddResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
}

// NewMessageServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMessageServiceHandler(svc MessageServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/core.message.v1.MessageService/Add", connect_go.NewUnaryHandler(
		"/core.message.v1.MessageService/Add",
		svc.Add,
		opts...,
	))
	mux.Handle("/core.message.v1.MessageService/Delete", connect_go.NewUnaryHandler(
		"/core.message.v1.MessageService/Delete",
		svc.Delete,
		opts...,
	))
	mux.Handle("/core.message.v1.MessageService/List", connect_go.NewUnaryHandler(
		"/core.message.v1.MessageService/List",
		svc.List,
		opts...,
	))
	return "/core.message.v1.MessageService/", mux
}

// UnimplementedMessageServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMessageServiceHandler struct{}

func (UnimplementedMessageServiceHandler) Add(context.Context, *connect_go.Request[v1.AddRequest]) (*connect_go.Response[v1.AddResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.message.v1.MessageService.Add is not implemented"))
}

func (UnimplementedMessageServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.message.v1.MessageService.Delete is not implemented"))
}

func (UnimplementedMessageServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.message.v1.MessageService.List is not implemented"))
}

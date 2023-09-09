// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/user/v1/team.proto

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
	// TeamServiceName is the fully-qualified name of the TeamService service.
	TeamServiceName = "api.user.v1.TeamService"
)

// TeamServiceClient is a client for the api.user.v1.TeamService service.
type TeamServiceClient interface {
	TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error)
	TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error)
	TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error)
	TeamInvite(context.Context, *connect_go.Request[v1.TeamInviteRequest]) (*connect_go.Response[v1.TeamInviteResponse], error)
}

// NewTeamServiceClient constructs a client for the api.user.v1.TeamService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTeamServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TeamServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &teamServiceClient{
		teamCreate: connect_go.NewClient[v1.TeamCreateRequest, v1.TeamCreateResponse](
			httpClient,
			baseURL+"/api.user.v1.TeamService/team_create",
			opts...,
		),
		teamDelete: connect_go.NewClient[v1.TeamDeleteRequest, v1.TeamDeleteResponse](
			httpClient,
			baseURL+"/api.user.v1.TeamService/team_delete",
			opts...,
		),
		teamList: connect_go.NewClient[v1.TeamListRequest, v1.TeamListResponse](
			httpClient,
			baseURL+"/api.user.v1.TeamService/team_list",
			opts...,
		),
		teamInvite: connect_go.NewClient[v1.TeamInviteRequest, v1.TeamInviteResponse](
			httpClient,
			baseURL+"/api.user.v1.TeamService/team_invite",
			opts...,
		),
	}
}

// teamServiceClient implements TeamServiceClient.
type teamServiceClient struct {
	teamCreate *connect_go.Client[v1.TeamCreateRequest, v1.TeamCreateResponse]
	teamDelete *connect_go.Client[v1.TeamDeleteRequest, v1.TeamDeleteResponse]
	teamList   *connect_go.Client[v1.TeamListRequest, v1.TeamListResponse]
	teamInvite *connect_go.Client[v1.TeamInviteRequest, v1.TeamInviteResponse]
}

// TeamCreate calls api.user.v1.TeamService.team_create.
func (c *teamServiceClient) TeamCreate(ctx context.Context, req *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error) {
	return c.teamCreate.CallUnary(ctx, req)
}

// TeamDelete calls api.user.v1.TeamService.team_delete.
func (c *teamServiceClient) TeamDelete(ctx context.Context, req *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error) {
	return c.teamDelete.CallUnary(ctx, req)
}

// TeamList calls api.user.v1.TeamService.team_list.
func (c *teamServiceClient) TeamList(ctx context.Context, req *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error) {
	return c.teamList.CallUnary(ctx, req)
}

// TeamInvite calls api.user.v1.TeamService.team_invite.
func (c *teamServiceClient) TeamInvite(ctx context.Context, req *connect_go.Request[v1.TeamInviteRequest]) (*connect_go.Response[v1.TeamInviteResponse], error) {
	return c.teamInvite.CallUnary(ctx, req)
}

// TeamServiceHandler is an implementation of the api.user.v1.TeamService service.
type TeamServiceHandler interface {
	TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error)
	TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error)
	TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error)
	TeamInvite(context.Context, *connect_go.Request[v1.TeamInviteRequest]) (*connect_go.Response[v1.TeamInviteResponse], error)
}

// NewTeamServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTeamServiceHandler(svc TeamServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.user.v1.TeamService/team_create", connect_go.NewUnaryHandler(
		"/api.user.v1.TeamService/team_create",
		svc.TeamCreate,
		opts...,
	))
	mux.Handle("/api.user.v1.TeamService/team_delete", connect_go.NewUnaryHandler(
		"/api.user.v1.TeamService/team_delete",
		svc.TeamDelete,
		opts...,
	))
	mux.Handle("/api.user.v1.TeamService/team_list", connect_go.NewUnaryHandler(
		"/api.user.v1.TeamService/team_list",
		svc.TeamList,
		opts...,
	))
	mux.Handle("/api.user.v1.TeamService/team_invite", connect_go.NewUnaryHandler(
		"/api.user.v1.TeamService/team_invite",
		svc.TeamInvite,
		opts...,
	))
	return "/api.user.v1.TeamService/", mux
}

// UnimplementedTeamServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTeamServiceHandler struct{}

func (UnimplementedTeamServiceHandler) TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.TeamService.team_create is not implemented"))
}

func (UnimplementedTeamServiceHandler) TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.TeamService.team_delete is not implemented"))
}

func (UnimplementedTeamServiceHandler) TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.TeamService.team_list is not implemented"))
}

func (UnimplementedTeamServiceHandler) TeamInvite(context.Context, *connect_go.Request[v1.TeamInviteRequest]) (*connect_go.Response[v1.TeamInviteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.user.v1.TeamService.team_invite is not implemented"))
}
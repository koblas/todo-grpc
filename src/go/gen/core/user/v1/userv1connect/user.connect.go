// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: core/user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/koblas/grpc-todo/gen/core/user/v1"
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
	UserServiceName = "core.user.v1.UserService"
)

// UserServiceClient is a client for the core.user.v1.UserService service.
type UserServiceClient interface {
	FindBy(context.Context, *connect_go.Request[v1.FindByRequest]) (*connect_go.Response[v1.FindByResponse], error)
	// Create a new user (e.g. registration flow)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	// Update user information
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	ComparePassword(context.Context, *connect_go.Request[v1.ComparePasswordRequest]) (*connect_go.Response[v1.ComparePasswordResponse], error)
	AuthAssociate(context.Context, *connect_go.Request[v1.AuthAssociateRequest]) (*connect_go.Response[v1.AuthAssociateResponse], error)
	GetSettings(context.Context, *connect_go.Request[v1.UserServiceGetSettingsRequest]) (*connect_go.Response[v1.UserServiceGetSettingsResponse], error)
	SetSettings(context.Context, *connect_go.Request[v1.UserServiceSetSettingsRequest]) (*connect_go.Response[v1.UserServiceSetSettingsResponse], error)
	// Email address verification
	VerificationVerify(context.Context, *connect_go.Request[v1.VerificationVerifyRequest]) (*connect_go.Response[v1.VerificationVerifyResponse], error)
	// Forgot password flow
	ForgotSend(context.Context, *connect_go.Request[v1.ForgotSendRequest]) (*connect_go.Response[v1.ForgotSendResponse], error)
	ForgotVerify(context.Context, *connect_go.Request[v1.ForgotVerifyRequest]) (*connect_go.Response[v1.ForgotVerifyResponse], error)
	ForgotUpdate(context.Context, *connect_go.Request[v1.ForgotUpdateRequest]) (*connect_go.Response[v1.ForgotUpdateResponse], error)
	// Create a new team
	TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error)
	// Create a new team
	TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error)
	// List all teams this user is a member of
	TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error)
	// Add members to a given team (in INVITED status)
	TeamAddMembers(context.Context, *connect_go.Request[v1.TeamAddMembersRequest]) (*connect_go.Response[v1.TeamAddMembersResponse], error)
	// Add members to a given team
	TeamAcceptInvite(context.Context, *connect_go.Request[v1.TeamAcceptInviteRequest]) (*connect_go.Response[v1.TeamAcceptInviteResponse], error)
	// Delete members from a given team
	TeamRemoveMembers(context.Context, *connect_go.Request[v1.TeamRemoveMembersRequest]) (*connect_go.Response[v1.TeamRemoveMembersResponse], error)
	// For a given Team show get all members
	TeamListMembers(context.Context, *connect_go.Request[v1.TeamListMembersRequest]) (*connect_go.Response[v1.TeamListMembersResponse], error)
}

// NewUserServiceClient constructs a client for the core.user.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		findBy: connect_go.NewClient[v1.FindByRequest, v1.FindByResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/FindBy",
			opts...,
		),
		create: connect_go.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/Create",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/Update",
			opts...,
		),
		comparePassword: connect_go.NewClient[v1.ComparePasswordRequest, v1.ComparePasswordResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/ComparePassword",
			opts...,
		),
		authAssociate: connect_go.NewClient[v1.AuthAssociateRequest, v1.AuthAssociateResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/AuthAssociate",
			opts...,
		),
		getSettings: connect_go.NewClient[v1.UserServiceGetSettingsRequest, v1.UserServiceGetSettingsResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/GetSettings",
			opts...,
		),
		setSettings: connect_go.NewClient[v1.UserServiceSetSettingsRequest, v1.UserServiceSetSettingsResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/SetSettings",
			opts...,
		),
		verificationVerify: connect_go.NewClient[v1.VerificationVerifyRequest, v1.VerificationVerifyResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/VerificationVerify",
			opts...,
		),
		forgotSend: connect_go.NewClient[v1.ForgotSendRequest, v1.ForgotSendResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/ForgotSend",
			opts...,
		),
		forgotVerify: connect_go.NewClient[v1.ForgotVerifyRequest, v1.ForgotVerifyResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/ForgotVerify",
			opts...,
		),
		forgotUpdate: connect_go.NewClient[v1.ForgotUpdateRequest, v1.ForgotUpdateResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/ForgotUpdate",
			opts...,
		),
		teamCreate: connect_go.NewClient[v1.TeamCreateRequest, v1.TeamCreateResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamCreate",
			opts...,
		),
		teamDelete: connect_go.NewClient[v1.TeamDeleteRequest, v1.TeamDeleteResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamDelete",
			opts...,
		),
		teamList: connect_go.NewClient[v1.TeamListRequest, v1.TeamListResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamList",
			opts...,
		),
		teamAddMembers: connect_go.NewClient[v1.TeamAddMembersRequest, v1.TeamAddMembersResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamAddMembers",
			opts...,
		),
		teamAcceptInvite: connect_go.NewClient[v1.TeamAcceptInviteRequest, v1.TeamAcceptInviteResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamAcceptInvite",
			opts...,
		),
		teamRemoveMembers: connect_go.NewClient[v1.TeamRemoveMembersRequest, v1.TeamRemoveMembersResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamRemoveMembers",
			opts...,
		),
		teamListMembers: connect_go.NewClient[v1.TeamListMembersRequest, v1.TeamListMembersResponse](
			httpClient,
			baseURL+"/core.user.v1.UserService/TeamListMembers",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	findBy             *connect_go.Client[v1.FindByRequest, v1.FindByResponse]
	create             *connect_go.Client[v1.CreateRequest, v1.CreateResponse]
	update             *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	comparePassword    *connect_go.Client[v1.ComparePasswordRequest, v1.ComparePasswordResponse]
	authAssociate      *connect_go.Client[v1.AuthAssociateRequest, v1.AuthAssociateResponse]
	getSettings        *connect_go.Client[v1.UserServiceGetSettingsRequest, v1.UserServiceGetSettingsResponse]
	setSettings        *connect_go.Client[v1.UserServiceSetSettingsRequest, v1.UserServiceSetSettingsResponse]
	verificationVerify *connect_go.Client[v1.VerificationVerifyRequest, v1.VerificationVerifyResponse]
	forgotSend         *connect_go.Client[v1.ForgotSendRequest, v1.ForgotSendResponse]
	forgotVerify       *connect_go.Client[v1.ForgotVerifyRequest, v1.ForgotVerifyResponse]
	forgotUpdate       *connect_go.Client[v1.ForgotUpdateRequest, v1.ForgotUpdateResponse]
	teamCreate         *connect_go.Client[v1.TeamCreateRequest, v1.TeamCreateResponse]
	teamDelete         *connect_go.Client[v1.TeamDeleteRequest, v1.TeamDeleteResponse]
	teamList           *connect_go.Client[v1.TeamListRequest, v1.TeamListResponse]
	teamAddMembers     *connect_go.Client[v1.TeamAddMembersRequest, v1.TeamAddMembersResponse]
	teamAcceptInvite   *connect_go.Client[v1.TeamAcceptInviteRequest, v1.TeamAcceptInviteResponse]
	teamRemoveMembers  *connect_go.Client[v1.TeamRemoveMembersRequest, v1.TeamRemoveMembersResponse]
	teamListMembers    *connect_go.Client[v1.TeamListMembersRequest, v1.TeamListMembersResponse]
}

// FindBy calls core.user.v1.UserService.FindBy.
func (c *userServiceClient) FindBy(ctx context.Context, req *connect_go.Request[v1.FindByRequest]) (*connect_go.Response[v1.FindByResponse], error) {
	return c.findBy.CallUnary(ctx, req)
}

// Create calls core.user.v1.UserService.Create.
func (c *userServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls core.user.v1.UserService.Update.
func (c *userServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// ComparePassword calls core.user.v1.UserService.ComparePassword.
func (c *userServiceClient) ComparePassword(ctx context.Context, req *connect_go.Request[v1.ComparePasswordRequest]) (*connect_go.Response[v1.ComparePasswordResponse], error) {
	return c.comparePassword.CallUnary(ctx, req)
}

// AuthAssociate calls core.user.v1.UserService.AuthAssociate.
func (c *userServiceClient) AuthAssociate(ctx context.Context, req *connect_go.Request[v1.AuthAssociateRequest]) (*connect_go.Response[v1.AuthAssociateResponse], error) {
	return c.authAssociate.CallUnary(ctx, req)
}

// GetSettings calls core.user.v1.UserService.GetSettings.
func (c *userServiceClient) GetSettings(ctx context.Context, req *connect_go.Request[v1.UserServiceGetSettingsRequest]) (*connect_go.Response[v1.UserServiceGetSettingsResponse], error) {
	return c.getSettings.CallUnary(ctx, req)
}

// SetSettings calls core.user.v1.UserService.SetSettings.
func (c *userServiceClient) SetSettings(ctx context.Context, req *connect_go.Request[v1.UserServiceSetSettingsRequest]) (*connect_go.Response[v1.UserServiceSetSettingsResponse], error) {
	return c.setSettings.CallUnary(ctx, req)
}

// VerificationVerify calls core.user.v1.UserService.VerificationVerify.
func (c *userServiceClient) VerificationVerify(ctx context.Context, req *connect_go.Request[v1.VerificationVerifyRequest]) (*connect_go.Response[v1.VerificationVerifyResponse], error) {
	return c.verificationVerify.CallUnary(ctx, req)
}

// ForgotSend calls core.user.v1.UserService.ForgotSend.
func (c *userServiceClient) ForgotSend(ctx context.Context, req *connect_go.Request[v1.ForgotSendRequest]) (*connect_go.Response[v1.ForgotSendResponse], error) {
	return c.forgotSend.CallUnary(ctx, req)
}

// ForgotVerify calls core.user.v1.UserService.ForgotVerify.
func (c *userServiceClient) ForgotVerify(ctx context.Context, req *connect_go.Request[v1.ForgotVerifyRequest]) (*connect_go.Response[v1.ForgotVerifyResponse], error) {
	return c.forgotVerify.CallUnary(ctx, req)
}

// ForgotUpdate calls core.user.v1.UserService.ForgotUpdate.
func (c *userServiceClient) ForgotUpdate(ctx context.Context, req *connect_go.Request[v1.ForgotUpdateRequest]) (*connect_go.Response[v1.ForgotUpdateResponse], error) {
	return c.forgotUpdate.CallUnary(ctx, req)
}

// TeamCreate calls core.user.v1.UserService.TeamCreate.
func (c *userServiceClient) TeamCreate(ctx context.Context, req *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error) {
	return c.teamCreate.CallUnary(ctx, req)
}

// TeamDelete calls core.user.v1.UserService.TeamDelete.
func (c *userServiceClient) TeamDelete(ctx context.Context, req *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error) {
	return c.teamDelete.CallUnary(ctx, req)
}

// TeamList calls core.user.v1.UserService.TeamList.
func (c *userServiceClient) TeamList(ctx context.Context, req *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error) {
	return c.teamList.CallUnary(ctx, req)
}

// TeamAddMembers calls core.user.v1.UserService.TeamAddMembers.
func (c *userServiceClient) TeamAddMembers(ctx context.Context, req *connect_go.Request[v1.TeamAddMembersRequest]) (*connect_go.Response[v1.TeamAddMembersResponse], error) {
	return c.teamAddMembers.CallUnary(ctx, req)
}

// TeamAcceptInvite calls core.user.v1.UserService.TeamAcceptInvite.
func (c *userServiceClient) TeamAcceptInvite(ctx context.Context, req *connect_go.Request[v1.TeamAcceptInviteRequest]) (*connect_go.Response[v1.TeamAcceptInviteResponse], error) {
	return c.teamAcceptInvite.CallUnary(ctx, req)
}

// TeamRemoveMembers calls core.user.v1.UserService.TeamRemoveMembers.
func (c *userServiceClient) TeamRemoveMembers(ctx context.Context, req *connect_go.Request[v1.TeamRemoveMembersRequest]) (*connect_go.Response[v1.TeamRemoveMembersResponse], error) {
	return c.teamRemoveMembers.CallUnary(ctx, req)
}

// TeamListMembers calls core.user.v1.UserService.TeamListMembers.
func (c *userServiceClient) TeamListMembers(ctx context.Context, req *connect_go.Request[v1.TeamListMembersRequest]) (*connect_go.Response[v1.TeamListMembersResponse], error) {
	return c.teamListMembers.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the core.user.v1.UserService service.
type UserServiceHandler interface {
	FindBy(context.Context, *connect_go.Request[v1.FindByRequest]) (*connect_go.Response[v1.FindByResponse], error)
	// Create a new user (e.g. registration flow)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	// Update user information
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	ComparePassword(context.Context, *connect_go.Request[v1.ComparePasswordRequest]) (*connect_go.Response[v1.ComparePasswordResponse], error)
	AuthAssociate(context.Context, *connect_go.Request[v1.AuthAssociateRequest]) (*connect_go.Response[v1.AuthAssociateResponse], error)
	GetSettings(context.Context, *connect_go.Request[v1.UserServiceGetSettingsRequest]) (*connect_go.Response[v1.UserServiceGetSettingsResponse], error)
	SetSettings(context.Context, *connect_go.Request[v1.UserServiceSetSettingsRequest]) (*connect_go.Response[v1.UserServiceSetSettingsResponse], error)
	// Email address verification
	VerificationVerify(context.Context, *connect_go.Request[v1.VerificationVerifyRequest]) (*connect_go.Response[v1.VerificationVerifyResponse], error)
	// Forgot password flow
	ForgotSend(context.Context, *connect_go.Request[v1.ForgotSendRequest]) (*connect_go.Response[v1.ForgotSendResponse], error)
	ForgotVerify(context.Context, *connect_go.Request[v1.ForgotVerifyRequest]) (*connect_go.Response[v1.ForgotVerifyResponse], error)
	ForgotUpdate(context.Context, *connect_go.Request[v1.ForgotUpdateRequest]) (*connect_go.Response[v1.ForgotUpdateResponse], error)
	// Create a new team
	TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error)
	// Create a new team
	TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error)
	// List all teams this user is a member of
	TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error)
	// Add members to a given team (in INVITED status)
	TeamAddMembers(context.Context, *connect_go.Request[v1.TeamAddMembersRequest]) (*connect_go.Response[v1.TeamAddMembersResponse], error)
	// Add members to a given team
	TeamAcceptInvite(context.Context, *connect_go.Request[v1.TeamAcceptInviteRequest]) (*connect_go.Response[v1.TeamAcceptInviteResponse], error)
	// Delete members from a given team
	TeamRemoveMembers(context.Context, *connect_go.Request[v1.TeamRemoveMembersRequest]) (*connect_go.Response[v1.TeamRemoveMembersResponse], error)
	// For a given Team show get all members
	TeamListMembers(context.Context, *connect_go.Request[v1.TeamListMembersRequest]) (*connect_go.Response[v1.TeamListMembersResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/core.user.v1.UserService/FindBy", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/FindBy",
		svc.FindBy,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/Create", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/Create",
		svc.Create,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/Update", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/ComparePassword", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/ComparePassword",
		svc.ComparePassword,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/AuthAssociate", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/AuthAssociate",
		svc.AuthAssociate,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/GetSettings", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/GetSettings",
		svc.GetSettings,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/SetSettings", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/SetSettings",
		svc.SetSettings,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/VerificationVerify", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/VerificationVerify",
		svc.VerificationVerify,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/ForgotSend", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/ForgotSend",
		svc.ForgotSend,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/ForgotVerify", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/ForgotVerify",
		svc.ForgotVerify,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/ForgotUpdate", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/ForgotUpdate",
		svc.ForgotUpdate,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamCreate", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamCreate",
		svc.TeamCreate,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamDelete", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamDelete",
		svc.TeamDelete,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamList", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamList",
		svc.TeamList,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamAddMembers", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamAddMembers",
		svc.TeamAddMembers,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamAcceptInvite", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamAcceptInvite",
		svc.TeamAcceptInvite,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamRemoveMembers", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamRemoveMembers",
		svc.TeamRemoveMembers,
		opts...,
	))
	mux.Handle("/core.user.v1.UserService/TeamListMembers", connect_go.NewUnaryHandler(
		"/core.user.v1.UserService/TeamListMembers",
		svc.TeamListMembers,
		opts...,
	))
	return "/core.user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) FindBy(context.Context, *connect_go.Request[v1.FindByRequest]) (*connect_go.Response[v1.FindByResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.FindBy is not implemented"))
}

func (UnimplementedUserServiceHandler) Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.Create is not implemented"))
}

func (UnimplementedUserServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.Update is not implemented"))
}

func (UnimplementedUserServiceHandler) ComparePassword(context.Context, *connect_go.Request[v1.ComparePasswordRequest]) (*connect_go.Response[v1.ComparePasswordResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.ComparePassword is not implemented"))
}

func (UnimplementedUserServiceHandler) AuthAssociate(context.Context, *connect_go.Request[v1.AuthAssociateRequest]) (*connect_go.Response[v1.AuthAssociateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.AuthAssociate is not implemented"))
}

func (UnimplementedUserServiceHandler) GetSettings(context.Context, *connect_go.Request[v1.UserServiceGetSettingsRequest]) (*connect_go.Response[v1.UserServiceGetSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.GetSettings is not implemented"))
}

func (UnimplementedUserServiceHandler) SetSettings(context.Context, *connect_go.Request[v1.UserServiceSetSettingsRequest]) (*connect_go.Response[v1.UserServiceSetSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.SetSettings is not implemented"))
}

func (UnimplementedUserServiceHandler) VerificationVerify(context.Context, *connect_go.Request[v1.VerificationVerifyRequest]) (*connect_go.Response[v1.VerificationVerifyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.VerificationVerify is not implemented"))
}

func (UnimplementedUserServiceHandler) ForgotSend(context.Context, *connect_go.Request[v1.ForgotSendRequest]) (*connect_go.Response[v1.ForgotSendResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.ForgotSend is not implemented"))
}

func (UnimplementedUserServiceHandler) ForgotVerify(context.Context, *connect_go.Request[v1.ForgotVerifyRequest]) (*connect_go.Response[v1.ForgotVerifyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.ForgotVerify is not implemented"))
}

func (UnimplementedUserServiceHandler) ForgotUpdate(context.Context, *connect_go.Request[v1.ForgotUpdateRequest]) (*connect_go.Response[v1.ForgotUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.ForgotUpdate is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamCreate(context.Context, *connect_go.Request[v1.TeamCreateRequest]) (*connect_go.Response[v1.TeamCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamCreate is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamDelete(context.Context, *connect_go.Request[v1.TeamDeleteRequest]) (*connect_go.Response[v1.TeamDeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamDelete is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamList(context.Context, *connect_go.Request[v1.TeamListRequest]) (*connect_go.Response[v1.TeamListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamList is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamAddMembers(context.Context, *connect_go.Request[v1.TeamAddMembersRequest]) (*connect_go.Response[v1.TeamAddMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamAddMembers is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamAcceptInvite(context.Context, *connect_go.Request[v1.TeamAcceptInviteRequest]) (*connect_go.Response[v1.TeamAcceptInviteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamAcceptInvite is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamRemoveMembers(context.Context, *connect_go.Request[v1.TeamRemoveMembersRequest]) (*connect_go.Response[v1.TeamRemoveMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamRemoveMembers is not implemented"))
}

func (UnimplementedUserServiceHandler) TeamListMembers(context.Context, *connect_go.Request[v1.TeamListMembersRequest]) (*connect_go.Response[v1.TeamListMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("core.user.v1.UserService.TeamListMembers is not implemented"))
}

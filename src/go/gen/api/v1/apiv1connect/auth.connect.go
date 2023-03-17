// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/auth.proto

package apiv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/koblas/grpc-todo/gen/api/v1"
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
	// AuthenticationServiceName is the fully-qualified name of the AuthenticationService service.
	AuthenticationServiceName = "api.v1.AuthenticationService"
)

// AuthenticationServiceClient is a client for the api.v1.AuthenticationService service.
type AuthenticationServiceClient interface {
	Register(context.Context, *connect_go.Request[v1.RegisterRequest]) (*connect_go.Response[v1.RegisterResponse], error)
	Authenticate(context.Context, *connect_go.Request[v1.AuthenticateRequest]) (*connect_go.Response[v1.AuthenticateResponse], error)
	VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error)
	RecoverSend(context.Context, *connect_go.Request[v1.RecoverSendRequest]) (*connect_go.Response[v1.RecoverSendResponse], error)
	RecoverVerify(context.Context, *connect_go.Request[v1.RecoverVerifyRequest]) (*connect_go.Response[v1.RecoverVerifyResponse], error)
	RecoverUpdate(context.Context, *connect_go.Request[v1.RecoverUpdateRequest]) (*connect_go.Response[v1.RecoverUpdateResponse], error)
	OauthLogin(context.Context, *connect_go.Request[v1.OauthLoginRequest]) (*connect_go.Response[v1.OauthLoginResponse], error)
	OauthUrl(context.Context, *connect_go.Request[v1.OauthUrlRequest]) (*connect_go.Response[v1.OauthUrlResponse], error)
}

// NewAuthenticationServiceClient constructs a client for the api.v1.AuthenticationService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthenticationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthenticationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authenticationServiceClient{
		register: connect_go.NewClient[v1.RegisterRequest, v1.RegisterResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/register",
			opts...,
		),
		authenticate: connect_go.NewClient[v1.AuthenticateRequest, v1.AuthenticateResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/authenticate",
			opts...,
		),
		verifyEmail: connect_go.NewClient[v1.VerifyEmailRequest, v1.VerifyEmailResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/verify_email",
			opts...,
		),
		recoverSend: connect_go.NewClient[v1.RecoverSendRequest, v1.RecoverSendResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/recover_send",
			opts...,
		),
		recoverVerify: connect_go.NewClient[v1.RecoverVerifyRequest, v1.RecoverVerifyResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/recover_verify",
			opts...,
		),
		recoverUpdate: connect_go.NewClient[v1.RecoverUpdateRequest, v1.RecoverUpdateResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/recover_update",
			opts...,
		),
		oauthLogin: connect_go.NewClient[v1.OauthLoginRequest, v1.OauthLoginResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/oauth_login",
			opts...,
		),
		oauthUrl: connect_go.NewClient[v1.OauthUrlRequest, v1.OauthUrlResponse](
			httpClient,
			baseURL+"/api.v1.AuthenticationService/oauth_url",
			opts...,
		),
	}
}

// authenticationServiceClient implements AuthenticationServiceClient.
type authenticationServiceClient struct {
	register      *connect_go.Client[v1.RegisterRequest, v1.RegisterResponse]
	authenticate  *connect_go.Client[v1.AuthenticateRequest, v1.AuthenticateResponse]
	verifyEmail   *connect_go.Client[v1.VerifyEmailRequest, v1.VerifyEmailResponse]
	recoverSend   *connect_go.Client[v1.RecoverSendRequest, v1.RecoverSendResponse]
	recoverVerify *connect_go.Client[v1.RecoverVerifyRequest, v1.RecoverVerifyResponse]
	recoverUpdate *connect_go.Client[v1.RecoverUpdateRequest, v1.RecoverUpdateResponse]
	oauthLogin    *connect_go.Client[v1.OauthLoginRequest, v1.OauthLoginResponse]
	oauthUrl      *connect_go.Client[v1.OauthUrlRequest, v1.OauthUrlResponse]
}

// Register calls api.v1.AuthenticationService.register.
func (c *authenticationServiceClient) Register(ctx context.Context, req *connect_go.Request[v1.RegisterRequest]) (*connect_go.Response[v1.RegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Authenticate calls api.v1.AuthenticationService.authenticate.
func (c *authenticationServiceClient) Authenticate(ctx context.Context, req *connect_go.Request[v1.AuthenticateRequest]) (*connect_go.Response[v1.AuthenticateResponse], error) {
	return c.authenticate.CallUnary(ctx, req)
}

// VerifyEmail calls api.v1.AuthenticationService.verify_email.
func (c *authenticationServiceClient) VerifyEmail(ctx context.Context, req *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error) {
	return c.verifyEmail.CallUnary(ctx, req)
}

// RecoverSend calls api.v1.AuthenticationService.recover_send.
func (c *authenticationServiceClient) RecoverSend(ctx context.Context, req *connect_go.Request[v1.RecoverSendRequest]) (*connect_go.Response[v1.RecoverSendResponse], error) {
	return c.recoverSend.CallUnary(ctx, req)
}

// RecoverVerify calls api.v1.AuthenticationService.recover_verify.
func (c *authenticationServiceClient) RecoverVerify(ctx context.Context, req *connect_go.Request[v1.RecoverVerifyRequest]) (*connect_go.Response[v1.RecoverVerifyResponse], error) {
	return c.recoverVerify.CallUnary(ctx, req)
}

// RecoverUpdate calls api.v1.AuthenticationService.recover_update.
func (c *authenticationServiceClient) RecoverUpdate(ctx context.Context, req *connect_go.Request[v1.RecoverUpdateRequest]) (*connect_go.Response[v1.RecoverUpdateResponse], error) {
	return c.recoverUpdate.CallUnary(ctx, req)
}

// OauthLogin calls api.v1.AuthenticationService.oauth_login.
func (c *authenticationServiceClient) OauthLogin(ctx context.Context, req *connect_go.Request[v1.OauthLoginRequest]) (*connect_go.Response[v1.OauthLoginResponse], error) {
	return c.oauthLogin.CallUnary(ctx, req)
}

// OauthUrl calls api.v1.AuthenticationService.oauth_url.
func (c *authenticationServiceClient) OauthUrl(ctx context.Context, req *connect_go.Request[v1.OauthUrlRequest]) (*connect_go.Response[v1.OauthUrlResponse], error) {
	return c.oauthUrl.CallUnary(ctx, req)
}

// AuthenticationServiceHandler is an implementation of the api.v1.AuthenticationService service.
type AuthenticationServiceHandler interface {
	Register(context.Context, *connect_go.Request[v1.RegisterRequest]) (*connect_go.Response[v1.RegisterResponse], error)
	Authenticate(context.Context, *connect_go.Request[v1.AuthenticateRequest]) (*connect_go.Response[v1.AuthenticateResponse], error)
	VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error)
	RecoverSend(context.Context, *connect_go.Request[v1.RecoverSendRequest]) (*connect_go.Response[v1.RecoverSendResponse], error)
	RecoverVerify(context.Context, *connect_go.Request[v1.RecoverVerifyRequest]) (*connect_go.Response[v1.RecoverVerifyResponse], error)
	RecoverUpdate(context.Context, *connect_go.Request[v1.RecoverUpdateRequest]) (*connect_go.Response[v1.RecoverUpdateResponse], error)
	OauthLogin(context.Context, *connect_go.Request[v1.OauthLoginRequest]) (*connect_go.Response[v1.OauthLoginResponse], error)
	OauthUrl(context.Context, *connect_go.Request[v1.OauthUrlRequest]) (*connect_go.Response[v1.OauthUrlResponse], error)
}

// NewAuthenticationServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthenticationServiceHandler(svc AuthenticationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.v1.AuthenticationService/register", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/register",
		svc.Register,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/authenticate", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/authenticate",
		svc.Authenticate,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/verify_email", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/verify_email",
		svc.VerifyEmail,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/recover_send", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/recover_send",
		svc.RecoverSend,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/recover_verify", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/recover_verify",
		svc.RecoverVerify,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/recover_update", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/recover_update",
		svc.RecoverUpdate,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/oauth_login", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/oauth_login",
		svc.OauthLogin,
		opts...,
	))
	mux.Handle("/api.v1.AuthenticationService/oauth_url", connect_go.NewUnaryHandler(
		"/api.v1.AuthenticationService/oauth_url",
		svc.OauthUrl,
		opts...,
	))
	return "/api.v1.AuthenticationService/", mux
}

// UnimplementedAuthenticationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthenticationServiceHandler struct{}

func (UnimplementedAuthenticationServiceHandler) Register(context.Context, *connect_go.Request[v1.RegisterRequest]) (*connect_go.Response[v1.RegisterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.register is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) Authenticate(context.Context, *connect_go.Request[v1.AuthenticateRequest]) (*connect_go.Response[v1.AuthenticateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.authenticate is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.verify_email is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) RecoverSend(context.Context, *connect_go.Request[v1.RecoverSendRequest]) (*connect_go.Response[v1.RecoverSendResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.recover_send is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) RecoverVerify(context.Context, *connect_go.Request[v1.RecoverVerifyRequest]) (*connect_go.Response[v1.RecoverVerifyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.recover_verify is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) RecoverUpdate(context.Context, *connect_go.Request[v1.RecoverUpdateRequest]) (*connect_go.Response[v1.RecoverUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.recover_update is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) OauthLogin(context.Context, *connect_go.Request[v1.OauthLoginRequest]) (*connect_go.Response[v1.OauthLoginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.oauth_login is not implemented"))
}

func (UnimplementedAuthenticationServiceHandler) OauthUrl(context.Context, *connect_go.Request[v1.OauthUrlRequest]) (*connect_go.Response[v1.OauthUrlResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1.AuthenticationService.oauth_url is not implemented"))
}

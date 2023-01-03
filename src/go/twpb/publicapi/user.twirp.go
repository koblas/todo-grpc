// Code generated by protoc-gen-twirp v8.1.1, DO NOT EDIT.
// source: publicapi/user.proto

package publicapi

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// =====================
// UserService Interface
// =====================

type UserService interface {
	GetUser(context.Context, *UserGetParams) (*UserResponse, error)

	UpdateUser(context.Context, *UserUpdateParams) (*UserResponse, error)
}

// ===========================
// UserService Protobuf Client
// ===========================

type userServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewUserServiceProtobufClient creates a Protobuf client that implements the UserService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewUserServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "api.user", "UserService")
	urls := [2]string{
		serviceURL + "GetUser",
		serviceURL + "UpdateUser",
	}
	if literalURLs {
		urls = [2]string{
			serviceURL + "get_user",
			serviceURL + "update_user",
		}
	}

	return &userServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *userServiceProtobufClient) GetUser(ctx context.Context, in *UserGetParams) (*UserResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "api.user")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "GetUser")
	caller := c.callGetUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserGetParams) (*UserResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserGetParams) when calling interceptor")
					}
					return c.callGetUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceProtobufClient) callGetUser(ctx context.Context, in *UserGetParams) (*UserResponse, error) {
	out := new(UserResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *userServiceProtobufClient) UpdateUser(ctx context.Context, in *UserUpdateParams) (*UserResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "api.user")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "UpdateUser")
	caller := c.callUpdateUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserUpdateParams) (*UserResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserUpdateParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserUpdateParams) when calling interceptor")
					}
					return c.callUpdateUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceProtobufClient) callUpdateUser(ctx context.Context, in *UserUpdateParams) (*UserResponse, error) {
	out := new(UserResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// UserService JSON Client
// =======================

type userServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewUserServiceJSONClient creates a JSON client that implements the UserService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewUserServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "api.user", "UserService")
	urls := [2]string{
		serviceURL + "GetUser",
		serviceURL + "UpdateUser",
	}
	if literalURLs {
		urls = [2]string{
			serviceURL + "get_user",
			serviceURL + "update_user",
		}
	}

	return &userServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *userServiceJSONClient) GetUser(ctx context.Context, in *UserGetParams) (*UserResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "api.user")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "GetUser")
	caller := c.callGetUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserGetParams) (*UserResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserGetParams) when calling interceptor")
					}
					return c.callGetUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceJSONClient) callGetUser(ctx context.Context, in *UserGetParams) (*UserResponse, error) {
	out := new(UserResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *userServiceJSONClient) UpdateUser(ctx context.Context, in *UserUpdateParams) (*UserResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "api.user")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "UpdateUser")
	caller := c.callUpdateUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserUpdateParams) (*UserResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserUpdateParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserUpdateParams) when calling interceptor")
					}
					return c.callUpdateUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceJSONClient) callUpdateUser(ctx context.Context, in *UserUpdateParams) (*UserResponse, error) {
	out := new(UserResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// UserService Server Handler
// ==========================

type userServiceServer struct {
	UserService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewUserServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewUserServiceServer(svc UserService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &userServiceServer{
		UserService:      svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *userServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *userServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// UserServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const UserServicePathPrefix = "/twirp/api.user.UserService/"

func (s *userServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "api.user")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "api.user.UserService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "get_user", "GetUser":
		s.serveGetUser(ctx, resp, req)
		return
	case "update_user", "UpdateUser":
		s.serveUpdateUser(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *userServiceServer) serveGetUser(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUserJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetUserProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveGetUserJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(UserGetParams)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.UserService.GetUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserGetParams) (*UserResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserGetParams) when calling interceptor")
					}
					return s.UserService.GetUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *UserResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserResponse and nil error while calling GetUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveGetUserProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(UserGetParams)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.UserService.GetUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserGetParams) (*UserResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserGetParams) when calling interceptor")
					}
					return s.UserService.GetUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *UserResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserResponse and nil error while calling GetUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveUpdateUser(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUpdateUserJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUpdateUserProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveUpdateUserJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdateUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(UserUpdateParams)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.UserService.UpdateUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserUpdateParams) (*UserResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserUpdateParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserUpdateParams) when calling interceptor")
					}
					return s.UserService.UpdateUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *UserResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserResponse and nil error while calling UpdateUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveUpdateUserProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdateUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(UserUpdateParams)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.UserService.UpdateUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserUpdateParams) (*UserResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserUpdateParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserUpdateParams) when calling interceptor")
					}
					return s.UserService.UpdateUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*UserResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*UserResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *UserResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserResponse and nil error while calling UpdateUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor3, 0
}

func (s *userServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.1"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *userServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "api.user", "UserService")
}

var twirpFileDescriptor3 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xfd, 0x6d, 0x29, 0xfc, 0xca, 0x94, 0x3f, 0x66, 0x43, 0xa4, 0x72, 0xd1, 0xf4, 0x60, 0x3c,
	0x95, 0x04, 0x8f, 0x5c, 0xb4, 0x17, 0x3d, 0x19, 0x53, 0xc3, 0xc5, 0x4b, 0xb3, 0xd0, 0x89, 0xd9,
	0x04, 0xe8, 0x66, 0xb7, 0xc8, 0x57, 0xe0, 0xe2, 0x77, 0xf2, 0xa3, 0x99, 0x9d, 0xb2, 0x0d, 0x98,
	0x78, 0x7c, 0xf3, 0xde, 0xbe, 0x9d, 0x37, 0x0f, 0x46, 0x6a, 0xb7, 0x5c, 0xcb, 0x95, 0x50, 0x72,
	0xba, 0x33, 0xa8, 0x13, 0xa5, 0xcb, 0xaa, 0xe4, 0x81, 0x50, 0x32, 0xb1, 0x38, 0x1e, 0x42, 0x7f,
	0x61, 0x50, 0x3f, 0x61, 0xf5, 0x2a, 0xb4, 0xd8, 0x98, 0xf8, 0x9b, 0xc1, 0x85, 0x9d, 0x2c, 0x54,
	0x21, 0x2a, 0xac, 0x87, 0xfc, 0x0a, 0xda, 0xb8, 0x11, 0x72, 0x1d, 0xb1, 0x1b, 0x76, 0xd7, 0x7d,
	0xfe, 0x97, 0xd5, 0xf0, 0xc0, 0x18, 0x1f, 0x83, 0xbf, 0x15, 0x1b, 0x8c, 0x3c, 0x62, 0x58, 0x46,
	0xc8, 0x12, 0xd7, 0x10, 0x28, 0x61, 0xcc, 0xbe, 0xd4, 0x45, 0xd4, 0x22, 0xd2, 0xcb, 0x9a, 0x89,
	0x15, 0xdc, 0x42, 0xcf, 0xc1, 0x7c, 0x8b, 0xfb, 0xc8, 0x27, 0x51, 0x2b, 0x0b, 0xdd, 0xf4, 0x05,
	0xf7, 0x07, 0xc6, 0xd2, 0x00, 0x3a, 0x39, 0x7d, 0x97, 0xfe, 0x87, 0x76, 0x6e, 0xed, 0xd3, 0x10,
	0xba, 0xb9, 0x53, 0xa5, 0x43, 0xe8, 0xe7, 0xa7, 0x46, 0xf1, 0x03, 0xf8, 0x36, 0x01, 0x1f, 0x80,
	0x27, 0x8b, 0x7a, 0xe5, 0xcc, 0x93, 0x05, 0x1f, 0xb9, 0x14, 0xb4, 0xeb, 0x31, 0x03, 0xe7, 0xc7,
	0x00, 0xb4, 0x63, 0xbd, 0x7e, 0x3c, 0x83, 0x9e, 0x75, 0xc8, 0xd0, 0xa8, 0x72, 0x6b, 0x90, 0xc7,
	0xe0, 0xdb, 0x6b, 0x91, 0x57, 0x38, 0x1b, 0x24, 0xee, 0x7c, 0x09, 0xa9, 0x88, 0x9b, 0x7d, 0x31,
	0x08, 0x2d, 0x7c, 0x43, 0xfd, 0x29, 0x57, 0xc8, 0xe7, 0x10, 0x7c, 0x60, 0x95, 0x5b, 0x8e, 0x8f,
	0xcf, 0x5f, 0x34, 0xd7, 0x9e, 0x5c, 0xfe, 0xb2, 0x72, 0x1f, 0x3e, 0x42, 0xb8, 0xa3, 0x02, 0xea,
	0xf7, 0x93, 0x73, 0xd9, 0x69, 0x37, 0x7f, 0x59, 0xa4, 0x83, 0xf7, 0x5e, 0x32, 0x9d, 0x37, 0xf5,
	0x2f, 0x3b, 0x54, 0xfd, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xe3, 0xb6, 0xb5, 0x12,
	0x02, 0x00, 0x00,
}

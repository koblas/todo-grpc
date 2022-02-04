// Code generated by protoc-gen-twirp v8.1.1, DO NOT EDIT.
// source: core/todo.proto

package core

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
// TodoService Interface
// =====================

type TodoService interface {
	AddTodo(context.Context, *TodoAddParams) (*TodoObject, error)

	DeleteTodo(context.Context, *TodoDeleteParams) (*TodoDeleteResponse, error)

	GetTodos(context.Context, *TodoGetParams) (*TodoResponse, error)
}

// ===========================
// TodoService Protobuf Client
// ===========================

type todoServiceProtobufClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTodoServiceProtobufClient creates a Protobuf client that implements the TodoService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTodoServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TodoService {
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
	serviceURL += baseServicePath(pathPrefix, "core.todo", "TodoService")
	urls := [3]string{
		serviceURL + "AddTodo",
		serviceURL + "DeleteTodo",
		serviceURL + "GetTodos",
	}

	return &todoServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *todoServiceProtobufClient) AddTodo(ctx context.Context, in *TodoAddParams) (*TodoObject, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "AddTodo")
	caller := c.callAddTodo
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoAddParams) (*TodoObject, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoAddParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoAddParams) when calling interceptor")
					}
					return c.callAddTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoObject)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoObject) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceProtobufClient) callAddTodo(ctx context.Context, in *TodoAddParams) (*TodoObject, error) {
	out := new(TodoObject)
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

func (c *todoServiceProtobufClient) DeleteTodo(ctx context.Context, in *TodoDeleteParams) (*TodoDeleteResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteTodo")
	caller := c.callDeleteTodo
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoDeleteParams) (*TodoDeleteResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoDeleteParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoDeleteParams) when calling interceptor")
					}
					return c.callDeleteTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoDeleteResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoDeleteResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceProtobufClient) callDeleteTodo(ctx context.Context, in *TodoDeleteParams) (*TodoDeleteResponse, error) {
	out := new(TodoDeleteResponse)
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

func (c *todoServiceProtobufClient) GetTodos(ctx context.Context, in *TodoGetParams) (*TodoResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "GetTodos")
	caller := c.callGetTodos
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoGetParams) (*TodoResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoGetParams) when calling interceptor")
					}
					return c.callGetTodos(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceProtobufClient) callGetTodos(ctx context.Context, in *TodoGetParams) (*TodoResponse, error) {
	out := new(TodoResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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
// TodoService JSON Client
// =======================

type todoServiceJSONClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTodoServiceJSONClient creates a JSON client that implements the TodoService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTodoServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TodoService {
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
	serviceURL += baseServicePath(pathPrefix, "core.todo", "TodoService")
	urls := [3]string{
		serviceURL + "AddTodo",
		serviceURL + "DeleteTodo",
		serviceURL + "GetTodos",
	}

	return &todoServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *todoServiceJSONClient) AddTodo(ctx context.Context, in *TodoAddParams) (*TodoObject, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "AddTodo")
	caller := c.callAddTodo
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoAddParams) (*TodoObject, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoAddParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoAddParams) when calling interceptor")
					}
					return c.callAddTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoObject)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoObject) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceJSONClient) callAddTodo(ctx context.Context, in *TodoAddParams) (*TodoObject, error) {
	out := new(TodoObject)
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

func (c *todoServiceJSONClient) DeleteTodo(ctx context.Context, in *TodoDeleteParams) (*TodoDeleteResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteTodo")
	caller := c.callDeleteTodo
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoDeleteParams) (*TodoDeleteResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoDeleteParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoDeleteParams) when calling interceptor")
					}
					return c.callDeleteTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoDeleteResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoDeleteResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceJSONClient) callDeleteTodo(ctx context.Context, in *TodoDeleteParams) (*TodoDeleteResponse, error) {
	out := new(TodoDeleteResponse)
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

func (c *todoServiceJSONClient) GetTodos(ctx context.Context, in *TodoGetParams) (*TodoResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
	ctx = ctxsetters.WithMethodName(ctx, "GetTodos")
	caller := c.callGetTodos
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoGetParams) (*TodoResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoGetParams) when calling interceptor")
					}
					return c.callGetTodos(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoServiceJSONClient) callGetTodos(ctx context.Context, in *TodoGetParams) (*TodoResponse, error) {
	out := new(TodoResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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
// TodoService Server Handler
// ==========================

type todoServiceServer struct {
	TodoService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewTodoServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewTodoServiceServer(svc TodoService, opts ...interface{}) TwirpServer {
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

	return &todoServiceServer{
		TodoService:      svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *todoServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *todoServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// TodoServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const TodoServicePathPrefix = "/twirp/core.todo.TodoService/"

func (s *todoServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "core.todo")
	ctx = ctxsetters.WithServiceName(ctx, "TodoService")
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
	if pkgService != "core.todo.TodoService" {
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
	case "AddTodo":
		s.serveAddTodo(ctx, resp, req)
		return
	case "DeleteTodo":
		s.serveDeleteTodo(ctx, resp, req)
		return
	case "GetTodos":
		s.serveGetTodos(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *todoServiceServer) serveAddTodo(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveAddTodoJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveAddTodoProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *todoServiceServer) serveAddTodoJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "AddTodo")
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
	reqContent := new(TodoAddParams)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TodoService.AddTodo
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoAddParams) (*TodoObject, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoAddParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoAddParams) when calling interceptor")
					}
					return s.TodoService.AddTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoObject)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoObject) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoObject
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoObject and nil error while calling AddTodo. nil responses are not supported"))
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

func (s *todoServiceServer) serveAddTodoProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "AddTodo")
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
	reqContent := new(TodoAddParams)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TodoService.AddTodo
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoAddParams) (*TodoObject, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoAddParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoAddParams) when calling interceptor")
					}
					return s.TodoService.AddTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoObject)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoObject) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoObject
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoObject and nil error while calling AddTodo. nil responses are not supported"))
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

func (s *todoServiceServer) serveDeleteTodo(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDeleteTodoJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDeleteTodoProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *todoServiceServer) serveDeleteTodoJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteTodo")
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
	reqContent := new(TodoDeleteParams)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TodoService.DeleteTodo
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoDeleteParams) (*TodoDeleteResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoDeleteParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoDeleteParams) when calling interceptor")
					}
					return s.TodoService.DeleteTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoDeleteResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoDeleteResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoDeleteResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoDeleteResponse and nil error while calling DeleteTodo. nil responses are not supported"))
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

func (s *todoServiceServer) serveDeleteTodoProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteTodo")
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
	reqContent := new(TodoDeleteParams)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TodoService.DeleteTodo
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoDeleteParams) (*TodoDeleteResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoDeleteParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoDeleteParams) when calling interceptor")
					}
					return s.TodoService.DeleteTodo(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoDeleteResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoDeleteResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoDeleteResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoDeleteResponse and nil error while calling DeleteTodo. nil responses are not supported"))
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

func (s *todoServiceServer) serveGetTodos(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetTodosJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetTodosProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *todoServiceServer) serveGetTodosJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetTodos")
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
	reqContent := new(TodoGetParams)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TodoService.GetTodos
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoGetParams) (*TodoResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoGetParams) when calling interceptor")
					}
					return s.TodoService.GetTodos(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoResponse and nil error while calling GetTodos. nil responses are not supported"))
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

func (s *todoServiceServer) serveGetTodosProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetTodos")
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
	reqContent := new(TodoGetParams)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TodoService.GetTodos
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoGetParams) (*TodoResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoGetParams)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoGetParams) when calling interceptor")
					}
					return s.TodoService.GetTodos(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*TodoResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*TodoResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *TodoResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TodoResponse and nil error while calling GetTodos. nil responses are not supported"))
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

func (s *todoServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor3, 0
}

func (s *todoServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.1"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *todoServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "core.todo", "TodoService")
}

var twirpFileDescriptor3 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0xa9, 0x36, 0x66, 0xea, 0x17, 0x0b, 0xd2, 0x10, 0x11, 0x4a, 0x4e, 0x01, 0x61, 0x0b,
	0xf5, 0x66, 0xf5, 0x50, 0x11, 0x6a, 0x4f, 0x4a, 0xf4, 0xe4, 0x45, 0xd2, 0xcc, 0x20, 0x51, 0xeb,
	0x96, 0xdd, 0xd5, 0x1f, 0xea, 0x2f, 0x92, 0x49, 0xda, 0xad, 0xc6, 0x42, 0xbd, 0x65, 0xdf, 0xc7,
	0xcc, 0xcb, 0x63, 0xe0, 0xa0, 0x50, 0x9a, 0xfa, 0x56, 0xa1, 0x92, 0x73, 0xad, 0xac, 0x12, 0x21,
	0x03, 0x92, 0x81, 0x24, 0x85, 0xbd, 0x07, 0x85, 0x6a, 0x4c, 0xf6, 0x2e, 0xd7, 0xf9, 0xcc, 0x88,
	0x2e, 0x04, 0x1f, 0x86, 0xf4, 0x53, 0x89, 0x91, 0xd7, 0xf3, 0xd2, 0x30, 0x6b, 0xf3, 0x73, 0x82,
	0xc9, 0x45, 0xad, 0x1c, 0x21, 0x6e, 0x50, 0x0a, 0x01, 0x5b, 0x36, 0x37, 0xaf, 0x91, 0x5f, 0xa1,
	0xd5, 0x77, 0x32, 0x84, 0x43, 0x76, 0x5f, 0xd3, 0x1b, 0x59, 0xda, 0x34, 0x60, 0x1f, 0xfc, 0x12,
	0x17, 0x76, 0xbf, 0xc4, 0x64, 0x02, 0xc0, 0xe6, 0xdb, 0xe9, 0x0b, 0x15, 0xf6, 0xdf, 0x36, 0x97,
	0xa3, 0xf5, 0x2b, 0xc7, 0x2e, 0x8f, 0xca, 0xc8, 0xcc, 0xd5, 0xbb, 0x21, 0x71, 0x0a, 0xdb, 0xdc,
	0x83, 0x89, 0xbc, 0x5e, 0x2b, 0xed, 0x0c, 0x8e, 0xa4, 0xab, 0x46, 0xae, 0x56, 0x66, 0xb5, 0x26,
	0x91, 0x20, 0x56, 0x3f, 0xe1, 0x46, 0x44, 0x10, 0xcc, 0xc8, 0x98, 0xfc, 0x99, 0x16, 0x79, 0x96,
	0xcf, 0xc1, 0x97, 0x07, 0x1d, 0x36, 0xdc, 0x93, 0xfe, 0x2c, 0x0b, 0x12, 0xe7, 0x10, 0x8c, 0x10,
	0x19, 0x11, 0x51, 0x63, 0x91, 0xab, 0x35, 0x5e, 0x1f, 0x41, 0xdc, 0x00, 0xd4, 0x7b, 0x2b, 0xfb,
	0x71, 0x43, 0xf4, 0xb3, 0xd7, 0xf8, 0x64, 0x2d, 0xe9, 0xf2, 0x5e, 0xc2, 0xce, 0x98, 0x2c, 0x13,
	0xe6, 0x4f, 0x0c, 0x77, 0x07, 0x71, 0xb7, 0xc1, 0x2c, 0xed, 0x57, 0xe1, 0x63, 0x20, 0xfb, 0x43,
	0x26, 0xa7, 0xed, 0xea, 0x9c, 0xce, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x91, 0x3a, 0x70, 0xc8,
	0x61, 0x02, 0x00, 0x00,
}

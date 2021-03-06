// Code generated by protoc-gen-twirp v8.1.1, DO NOT EDIT.
// source: core/eventbus.proto

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

// ======================
// TodoEventbus Interface
// ======================

type TodoEventbus interface {
	Message(context.Context, *TodoEvent) (*EventbusEmpty, error)
}

// ============================
// TodoEventbus Protobuf Client
// ============================

type todoEventbusProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTodoEventbusProtobufClient creates a Protobuf client that implements the TodoEventbus interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTodoEventbusProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TodoEventbus {
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
	serviceURL += baseServicePath(pathPrefix, "core.eventbus", "TodoEventbus")
	urls := [1]string{
		serviceURL + "Message",
	}

	return &todoEventbusProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *todoEventbusProtobufClient) Message(ctx context.Context, in *TodoEvent) (*EventbusEmpty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.eventbus")
	ctx = ctxsetters.WithServiceName(ctx, "TodoEventbus")
	ctx = ctxsetters.WithMethodName(ctx, "Message")
	caller := c.callMessage
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoEvent) (*EventbusEmpty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoEvent)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoEvent) when calling interceptor")
					}
					return c.callMessage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*EventbusEmpty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*EventbusEmpty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoEventbusProtobufClient) callMessage(ctx context.Context, in *TodoEvent) (*EventbusEmpty, error) {
	out := new(EventbusEmpty)
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

// ========================
// TodoEventbus JSON Client
// ========================

type todoEventbusJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTodoEventbusJSONClient creates a JSON client that implements the TodoEventbus interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTodoEventbusJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TodoEventbus {
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
	serviceURL += baseServicePath(pathPrefix, "core.eventbus", "TodoEventbus")
	urls := [1]string{
		serviceURL + "Message",
	}

	return &todoEventbusJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *todoEventbusJSONClient) Message(ctx context.Context, in *TodoEvent) (*EventbusEmpty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "core.eventbus")
	ctx = ctxsetters.WithServiceName(ctx, "TodoEventbus")
	ctx = ctxsetters.WithMethodName(ctx, "Message")
	caller := c.callMessage
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *TodoEvent) (*EventbusEmpty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoEvent)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoEvent) when calling interceptor")
					}
					return c.callMessage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*EventbusEmpty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*EventbusEmpty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *todoEventbusJSONClient) callMessage(ctx context.Context, in *TodoEvent) (*EventbusEmpty, error) {
	out := new(EventbusEmpty)
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

// ===========================
// TodoEventbus Server Handler
// ===========================

type todoEventbusServer struct {
	TodoEventbus
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewTodoEventbusServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewTodoEventbusServer(svc TodoEventbus, opts ...interface{}) TwirpServer {
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

	return &todoEventbusServer{
		TodoEventbus:     svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *todoEventbusServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *todoEventbusServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// TodoEventbusPathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const TodoEventbusPathPrefix = "/twirp/core.eventbus.TodoEventbus/"

func (s *todoEventbusServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "core.eventbus")
	ctx = ctxsetters.WithServiceName(ctx, "TodoEventbus")
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
	if pkgService != "core.eventbus.TodoEventbus" {
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
	case "Message":
		s.serveMessage(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *todoEventbusServer) serveMessage(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveMessageJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveMessageProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *todoEventbusServer) serveMessageJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Message")
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
	reqContent := new(TodoEvent)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TodoEventbus.Message
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoEvent) (*EventbusEmpty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoEvent)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoEvent) when calling interceptor")
					}
					return s.TodoEventbus.Message(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*EventbusEmpty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*EventbusEmpty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *EventbusEmpty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *EventbusEmpty and nil error while calling Message. nil responses are not supported"))
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

func (s *todoEventbusServer) serveMessageProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Message")
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
	reqContent := new(TodoEvent)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TodoEventbus.Message
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *TodoEvent) (*EventbusEmpty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*TodoEvent)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*TodoEvent) when calling interceptor")
					}
					return s.TodoEventbus.Message(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*EventbusEmpty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*EventbusEmpty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *EventbusEmpty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *EventbusEmpty and nil error while calling Message. nil responses are not supported"))
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

func (s *todoEventbusServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *todoEventbusServer) ProtocGenTwirpVersion() string {
	return "v8.1.1"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *todoEventbusServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "core.eventbus", "TodoEventbus")
}

var twirpFileDescriptor1 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0x2f, 0x4a,
	0xd5, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x49, 0x2a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x05, 0x09, 0xea, 0xc1, 0x04, 0xa5, 0xf8, 0xc1, 0x6a, 0x4a, 0xf2, 0x53, 0xf2, 0x21, 0xf2,
	0x4a, 0xfc, 0x5c, 0xbc, 0xae, 0x50, 0x49, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x25, 0x57, 0x2e, 0x4e,
	0xa7, 0xc4, 0xe2, 0x54, 0xb0, 0xa0, 0x90, 0x02, 0x17, 0x77, 0x66, 0x4a, 0x6a, 0x6e, 0x41, 0x7e,
	0x5e, 0x72, 0xa5, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0xb2, 0x90, 0x90, 0x18,
	0x17, 0x5b, 0x62, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58, 0x12, 0xca, 0x53, 0x5a, 0xc9,
	0xc8, 0xc5, 0x19, 0x92, 0x9f, 0x92, 0x4f, 0xa1, 0x39, 0x42, 0xfa, 0x5c, 0xec, 0xc9, 0xa5, 0x45,
	0x45, 0xa9, 0x79, 0x25, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xa2, 0x7a, 0x60, 0x1f, 0x81,
	0xbd, 0x00, 0xb2, 0xc0, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0x24, 0x08, 0xa6, 0x4a, 0xc8, 0x90, 0x8b,
	0xa3, 0xa0, 0x28, 0xb5, 0x2c, 0x33, 0xbf, 0xb4, 0x58, 0x82, 0x05, 0x9f, 0x0e, 0xb8, 0x32, 0xa3,
	0x40, 0x2e, 0x1e, 0xb8, 0x53, 0x93, 0x4a, 0x8b, 0x85, 0x1c, 0xb9, 0xd8, 0x7d, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x85, 0x24, 0xf4, 0x50, 0xc2, 0x4f, 0x0f, 0xae, 0x4e, 0x4a, 0x06, 0x4d, 0x06,
	0x25, 0x14, 0x9d, 0x38, 0xa3, 0xd8, 0xf5, 0xf4, 0xad, 0x41, 0x2a, 0x92, 0xd8, 0xc0, 0x01, 0x6d,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x67, 0x50, 0xfb, 0x6c, 0x9f, 0x01, 0x00, 0x00,
}

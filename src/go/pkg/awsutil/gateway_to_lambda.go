package awsutil

import (
	"context"
	"net/http"
	"strings"

	"encoding/base64"
	"io"
	"net/url"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koblas/grpc-todo/pkg/manager"
)

/** Convert an Lambda APIGatewayProxyRequest to an HTTPRequest */
func GatewayToHttpRequest(ctx context.Context, event events.APIGatewayProxyRequest, useProxyPath bool) (*http.Request, error) {
	// Build request URL.
	params := url.Values{}
	for k, v := range event.QueryStringParameters {
		params.Set(k, v)
	}
	for k, vals := range event.MultiValueQueryStringParameters {
		params[k] = vals
	}

	u := url.URL{
		Host:     event.Headers["Host"],
		RawPath:  event.Path,
		RawQuery: params.Encode(),
	}
	if useProxyPath {
		u.RawPath = path.Join("/", event.PathParameters["proxy"])
	}

	// Unescape request path
	p, err := url.PathUnescape(u.RawPath)
	if err != nil {
		return nil, err
	}
	u.Path = p

	if u.Path == u.RawPath {
		u.RawPath = ""
	}

	// Handle base64 encoded body.
	var body io.Reader = strings.NewReader(event.Body)
	if event.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}

	// Create a new request.
	r, err := http.NewRequest(event.HTTPMethod, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set headers.
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-proxy-integrations.html
	// If you specify values for both headers and multiValueHeaders, API Gateway merges them into a single list.
	// If the same key-value pair is specified in both, only the values from multiValueHeaders will appear
	// the merged list.
	for k, v := range event.Headers {
		r.Header.Set(k, v)
	}
	for k, vals := range event.MultiValueHeaders {
		r.Header[http.CanonicalHeaderKey(k)] = vals
	}

	// Set remote IP address.
	r.RemoteAddr = event.RequestContext.Identity.SourceIP

	// Set request URI
	r.RequestURI = u.RequestURI()

	return r.WithContext(ctx), nil
}

/** Convert an Lambda APIGatewayProxyRequest to an HTTPRequest */
func GatewayToHttpRequestV2(ctx context.Context, event events.APIGatewayV2HTTPRequest, useProxyPath bool) (*http.Request, error) {
	// Build request URL.
	params := url.Values{}
	for k, v := range event.QueryStringParameters {
		params.Set(k, v)
	}
	// TODO - reparse the raw query string
	for k, vals := range event.QueryStringParameters {
		params[k] = []string{vals}
	}

	u := url.URL{
		Host:     event.Headers["Host"],
		Path:     event.RequestContext.HTTP.Path,
		RawPath:  event.RawPath,
		RawQuery: event.RawQueryString,
	}
	if useProxyPath {
		u.RawPath = path.Join("/", event.PathParameters["proxy"])
		u.Path = path.Join("/", event.PathParameters["proxy"])
	}

	// Handle base64 encoded body.
	var body io.Reader = strings.NewReader(event.Body)
	if event.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}

	headers := http.Header{}
	for k, v := range event.Headers {
		headers.Add(k, v)
	}

	// Create a new request.
	lctx := context.WithValue(ctx, manager.HttpHeaderCtxKey, headers)
	r, err := http.NewRequestWithContext(lctx, event.RequestContext.HTTP.Method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set headers.
	r.Header = headers

	// Set remote IP address.
	r.RemoteAddr = event.RequestContext.HTTP.SourceIP

	// Set request URI
	r.RequestURI = u.RequestURI()

	return r, nil
}

func SqsEventToHttpRequest(ctx context.Context, event events.SQSMessage) (*http.Request, error) {
	headers := http.Header{}
	for k, v := range event.MessageAttributes {
		if v.DataType == "String" {
			headers.Add(k, *v.StringValue)
		} else if v.DataType == "StringList" && len(v.StringListValues) != 0 {
			headers.Add(k, v.StringListValues[0])
		}
	}

	parts := strings.Split(event.EventSourceARN, ":")

	u := url.URL{
		Scheme: "sqs",
		Host:   parts[len(parts)-1],
	}
	u.Path = headers.Get("twirp.path")
	u.RawPath = headers.Get("twirp.path")

	var body io.Reader = strings.NewReader(event.Body)
	if headers.Get("content-transfer-encoding") == "base64" {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}

	lctx := context.WithValue(ctx, manager.HttpHeaderCtxKey, headers)
	r, err := http.NewRequestWithContext(lctx, http.MethodPost, u.String(), body)

	if err != nil {
		return nil, err
	}

	// Set headers.
	r.Header = headers

	// Set remote IP address.
	r.RemoteAddr = "0.0.0.0"

	// Set request URI
	r.RequestURI = u.RequestURI()

	return r, nil
}

package awsutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/koblas/grpc-todo/pkg/logger"
)

var errorResponse = events.APIGatewayV2HTTPResponse{
	StatusCode: 500,
}

type TwirpHttpHandler func(lambdaCtx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

func HandleLambda(ctx context.Context, api http.Handler) TwirpHttpHandler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.ServeHTTP(w, r)
	})

	return func(lambdaCtx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		log := logger.FromContext(ctx).With(
			"awsHttpMethod", request.RequestContext.HTTP.Method,
			"awsRequestID", request.RequestContext.RequestID,
			"awsPath", request.RequestContext.HTTP.Path,
		)

		log.Info("Processing Lambda request")

		req, err := GatewayToHttpRequestV2(logger.ToContext(lambdaCtx, log), request, false)
		if err != nil {
			log.With("error", err).Error("Unable to build request")
			return errorResponse, nil
		}
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		result := w.Result()

		buf := bytes.Buffer{}
		if _, err := buf.ReadFrom(result.Body); err != nil {
			log.With("error", err).Error("Unable to read buffer")
			return errorResponse, nil
		}

		simpleHeaders := map[string]string{}
		for k, v := range result.Header {
			simpleHeaders[k] = strings.Join(v, ",")
		}

		return events.APIGatewayV2HTTPResponse{
			StatusCode:        result.StatusCode,
			Body:              buf.String(),
			Headers:           simpleHeaders,
			MultiValueHeaders: result.Header,
		}, nil
	}
}

// Matches the Twirp HTTPClient interface
type TwClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type twClient struct {
	client *lambda.Client
}

func NewTwirpCallLambda() TwClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := lambda.NewFromConfig(cfg)

	return twClient{
		client: client,
	}
}

func (svc twClient) Do(req *http.Request) (*http.Response, error) {
	buf := strings.Builder{}
	_, err := io.Copy(&buf, req.Body)
	if err != nil {
		return nil, err
	}

	basicHeaders := map[string]string{}
	for k, v := range req.Header {
		basicHeaders[k] = strings.Join(v, ",")
	}

	lambdaRequest := events.APIGatewayV2HTTPRequest{
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Path:   req.URL.Path,
				Method: req.Method,
			},
		},
		RawPath:         req.URL.Path,
		RawQueryString:  req.URL.Query().Encode(),
		Headers:         basicHeaders,
		Body:            buf.String(),
		IsBase64Encoded: false,
		// Resource:              "",
		// QueryStringParameters: map[string]string{},
		// PathParameters:        map[string]string{},
		// StageVariables:        map[string]string{},
	}

	// Marshal
	payload, err := json.Marshal(&lambdaRequest)
	if err != nil {
		return nil, err
	}

	log := logger.FromContext(req.Context()).With("lambda", req.URL.Host).With("rpcMethod", req.URL.Path)
	log.Info("BEGIN: calling twirp service")
	start := time.Now()
	output, err := svc.client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String(req.URL.Host),
		Payload:      payload,
	})
	elapsed := int64(time.Since(start) / time.Millisecond)
	log.With("durationInMs", elapsed).Info("END: calling twirp service")

	if err != nil {
		return nil, err
	}

	// Unmarshal
	lambdaResponse := events.APIGatewayProxyResponse{}
	if err := json.Unmarshal(output.Payload, &lambdaResponse); err != nil {
		return nil, err
	}

	res := http.Response{
		StatusCode: lambdaResponse.StatusCode,
		Header:     lambdaResponse.MultiValueHeaders,
		Body:       ioutil.NopCloser(strings.NewReader(lambdaResponse.Body)),
	}

	return &res, nil
}

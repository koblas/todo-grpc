package awsutil

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snstypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/pkg/errors"
)

var errorResponse = events.APIGatewayV2HTTPResponse{
	StatusCode: 500,
}

type TwirpHttpApiHandler func(lambdaCtx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)
type TwirpHttpSqsHandler func(lambdaCtx context.Context, request events.SQSEvent) (events.SQSEventResponse, error)

func HandleApiLambda(ctx context.Context, api http.Handler) TwirpHttpApiHandler {
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

type SqsHandlers map[string]http.Handler

func HandleSqsLambda(handlers SqsHandlers, forcePath *string) TwirpHttpSqsHandler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range handlers {
			if strings.HasPrefix(r.URL.Path, k) {
				v.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	})

	return func(ctx context.Context, request events.SQSEvent) (events.SQSEventResponse, error) {
		parentLog := logger.FromContext(ctx).With("eventItemCount", len(request.Records))

		result := events.SQSEventResponse{}
		for _, record := range request.Records {
			log := parentLog.With("messageId", record.MessageId)
			log.Info("Handling SQS Message")

			req, err := SqsEventToHttpRequest(logger.ToContext(ctx, log), record, forcePath)
			if err != nil {
				log.With("error", err).Error("Unable to decode")
				result.BatchItemFailures = append(result.BatchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: record.MessageId})
				continue
			}

			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			res := w.Result()
			if res.StatusCode != http.StatusOK {
				buf, _ := io.ReadAll(io.LimitReader(res.Body, 256))
				log.With("statusCode", res.StatusCode).With("statusMsg", string(buf)).Info("SQS Message error")
				result.BatchItemFailures = append(result.BatchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: record.MessageId})
			}
		}

		return result, nil
	}
}

// Matches the Twirp HTTPClient interface
type TwClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type twClient struct {
	lambda *lambda.Client
	sqs    *sqs.Client
	sns    *sns.Client
}

func NewTwirpCallLambda() TwClient {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return twClient{
		lambda: lambda.NewFromConfig(cfg),
		sqs:    sqs.NewFromConfig(cfg),
		sns:    sns.NewFromConfig(cfg),
	}
}

func (svc twClient) Do(req *http.Request) (*http.Response, error) {
	scheme := req.URL.Scheme
	path := req.URL.Path
	arn := ""
	if req.URL.Scheme == "arn" {
		parts := strings.Split(req.URL.Opaque, ":")
		if parts[0] != "aws" {
			return nil, errors.New("unknown ARN format")
		}
		scheme = parts[1]
		idx := strings.Index(req.URL.Opaque, "/")
		if idx >= 0 {
			arn = "arn:" + req.URL.Opaque[:idx]
			path = req.URL.Opaque[idx:]
		} else {
			arn = "arn:" + req.URL.Opaque
			path = "/"
		}
	}

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
				Path:   path,
				Method: req.Method,
			},
		},
		RawPath:         path,
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

	log := logger.FromContext(req.Context()).With(
		"host", req.URL.Host,
		"arn", arn,
		"rpcMethod", path,
		"scheme", scheme,
	)
	log.Info("BEGIN: calling twirp service")
	start := time.Now()

	var res *http.Response

	if scheme == "http" || scheme == "https" {
		if arn != "" {
			return nil, errors.New("arn not supported for http")
		}
		client := http.Client{}
		res, err = client.Do(req)
		elapsed := int64(time.Since(start) / time.Millisecond)
		log.With("durationInMs", elapsed).Info("END: calling twirp service")

		if err != nil {
			return nil, err
		}
	} else if scheme == "lambda" {
		functionName := req.URL.Host
		if arn != "" {
			functionName = arn
		}
		output, err := svc.lambda.Invoke(context.TODO(), &lambda.InvokeInput{
			FunctionName: aws.String(functionName),
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

		res = &http.Response{
			StatusCode: lambdaResponse.StatusCode,
			Header:     lambdaResponse.MultiValueHeaders,
			Body:       ioutil.NopCloser(strings.NewReader(lambdaResponse.Body)),
		}
	} else if scheme == "sqs" {
		if arn != "" {
			return nil, errors.New("arn not supported for sqs")
		}
		isJson := false

		attributes := map[string]sqstypes.MessageAttributeValue{
			"twirp.path": {DataType: aws.String("String"), StringValue: &path},
		}
		for k, v := range req.URL.Query() {
			attributes[k] = sqstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String(v[0])}
		}
		for k, v := range basicHeaders {
			attributes[k] = sqstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String(v)}
			if strings.ToLower(k) == "content-type" {
				isJson = strings.Contains(v, "application/json")
			}
		}

		var body *string
		if isJson {
			body = &lambdaRequest.Body
		} else {
			attributes["Content-Transfer-Encoding"] = sqstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("base64")}
			uEnc := base64.URLEncoding.EncodeToString([]byte(lambdaRequest.Body))
			body = &uEnc
		}

		qurl, err := svc.sqs.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
			QueueName: &req.URL.Host,
		})
		if err != nil {
			return nil, fmt.Errorf("unable to get queue url='%s' %w", req.URL.Host, err)
		}

		_, err = svc.sqs.SendMessage(context.TODO(), &sqs.SendMessageInput{
			QueueUrl:          qurl.QueueUrl,
			MessageAttributes: attributes,
			MessageBody:       body,
		})
		elapsed := int64(time.Since(start) / time.Millisecond)
		log.With("durationInMs", elapsed).Info("END: calling twirp service")

		if err != nil {
			return nil, err
		}

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	} else if scheme == "sns" {
		isJson := false

		attributes := map[string]snstypes.MessageAttributeValue{
			"twirp.path": {DataType: aws.String("String"), StringValue: &path},
		}
		for k, v := range req.URL.Query() {
			attributes[k] = snstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String(v[0])}
		}
		for k, v := range basicHeaders {
			attributes[k] = snstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String(v)}
			if strings.ToLower(k) == "content-type" {
				isJson = strings.Contains(v, "application/json")
			}
		}

		var body *string
		if isJson {
			body = &lambdaRequest.Body
		} else {
			attributes["Content-Transfer-Encoding"] = snstypes.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("base64")}
			uEnc := base64.URLEncoding.EncodeToString([]byte(lambdaRequest.Body))
			body = &uEnc
		}

		_, err = svc.sns.Publish(context.TODO(), &sns.PublishInput{
			TopicArn:          aws.String(arn),
			MessageAttributes: attributes,
			Message:           body,
		})
		elapsed := int64(time.Since(start) / time.Millisecond)
		log.With("durationInMs", elapsed).Info("END: calling twirp service")

		if err != nil {
			return nil, err
		}

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	} else {
		return nil, errors.New("unknown scheme")
	}

	return res, nil
}

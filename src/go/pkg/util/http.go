package util

import (
	// "fmt"

	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/koblas/grpc-todo/pkg/logger"
)

func NewHttpTransportContext(ctx context.Context) context.Context {
	return ctx
}

func NewHttpClient(ctx context.Context) *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
	}
}

// DoHTTPRequest -- wrapper to make a "standard" http request in the system
func DoHTTPRequest(ctx context.Context, client *http.Client, logger logger.Logger, req *http.Request, result interface{}) error {
	if req.Header.Get("User-Agent") == "" {
		req.Header.Add("User-Agent", "Project X 0.1")
	}
	if req.Header.Get("Accept") == "" {
		req.Header.Add("Accept", "application/json")
	}

	logger = logger.With("url", req.URL.String())
	logLatency := logger.Latency()
	resp, err := client.Do(req)
	logLatency("DoHTTPRequest", "status", resp.StatusCode)

	if err != nil {
		logger.With("err", err).Error("HTTP request failed")
		return err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	switch resp.StatusCode / 100 {
	case 1: // 3xx response
		logger.With("status", resp.Status).Info("HTTP error: 1xx response")
		return errors.New("Unexpected response")
	case 2:
		{ // 2xx response
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return err
			}
			return json.Unmarshal(body, result)
		}
	case 3: // 3xx response
		logger.With("status", resp.Status).Info("HTTP error: 3xx response")
		return errors.New("Unexpected response")
	case 4:
		logger.With("status", resp.Status).Info("HTTP error: 4xx response")
		return errors.New("Unexpected response")
	case 5: // 5xx response
		logger.With("status", resp.Status).Info("HTTP error: 5xx response")
		return errors.New("Unexpected response")
	}

	logger.With("status", resp.Status).Info("HTTP error: bad status")
	return errors.New("Unexpected invalid status code")
}

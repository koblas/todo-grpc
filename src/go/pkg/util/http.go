package util

import (
	// "fmt"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/koblas/grpc-todo/pkg/logger"
	"go.uber.org/zap"
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

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return json.Unmarshal(body, result)
	}

	logger.With(
		zap.String("status", resp.Status),
		zap.String("body", string(body)),
	).Info("HTTP error: non-200 status")
	return fmt.Errorf("unexpected status=%s", resp.Status)
}

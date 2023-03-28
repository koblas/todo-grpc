package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// WrappedHTTPClient - builds an http client with token support
func WrappedHTTPClient(ctx context.Context, source oauth2.TokenSource) *http.Client {
	httpClient := util.NewHttpClient(ctx)
	client := httpClient
	if source != nil {
		client = oauth2.NewClient(context.WithValue(ctx, oauth2.HTTPClient, httpClient), source)
	}

	return client
}

func makeHTTPcall(ctx context.Context, source oauth2.TokenSource, logger logger.Logger, req *http.Request, result interface{}) error {
	client := WrappedHTTPClient(ctx, source)

	return util.DoHTTPRequest(ctx, client, logger, req, result)
}

func (svc providerBase) httpTokenRequest(ctx context.Context, logger logger.Logger, baseURL string, params map[string]interface{}) (TokenResult, error) {
	tokenResult := struct {
		TokenResult
		ExpiresAt *int64 `json:"expires_at,optional"`
		ExpiresIn *int64 `json:"expires_in,optional"`
	}{}

	u, err := url.Parse(baseURL)
	if err != nil {
		logger.Fatal("Failed to parse internal URL")
	}

	q := u.Query()

	for k, v := range params {
		if slist, ok := v.([]string); ok {
			q.Set(k, strings.Join(slist, ","))
		} else {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}

	u.RawQuery = q.Encode()

	logger = logger.With("url", baseURL)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(q.Encode()))
	if err != nil {
		logger.With(zap.Error(err)).Info("Building Request Failed")
		return tokenResult.TokenResult, err
	}

	req.Header.Add("Content-type", "application/x-www-form-urlencoded")

	logger.Info("OAuth Token HTTP call start")
	if err = makeHTTPcall(ctx, nil, logger, req, &tokenResult); err != nil {
		logger.With(zap.Error(err)).Info("OAuth Token HTTP call failed")
	}

	if tokenResult.ExpiresIn != nil {
		tval := time.Now().Add(time.Second * time.Duration(*tokenResult.ExpiresIn))
		tokenResult.Expires = &tval
	} else if tokenResult.ExpiresAt != nil {
		tval := time.Unix(*tokenResult.ExpiresAt, 0)
		tokenResult.Expires = &tval
	}

	return tokenResult.TokenResult, nil
}

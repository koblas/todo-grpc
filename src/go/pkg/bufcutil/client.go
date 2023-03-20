package bufcutil

import (
	"net"
	"net/http"
	"time"

	"github.com/koblas/grpc-todo/pkg/interceptors"
)

type httpClient struct {
	http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		Client: http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 1 * time.Second,
					// KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   1 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				// ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: 30 * time.Second,
		},
	}
}

func (svc *httpClient) Do(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	reqId := ctx.Value(interceptors.RequestIdCtxKey)
	if val, ok := reqId.(string); ok && val != "" {
		req.Header.Add((interceptors.RequestIdHeader), val)
	}
	return svc.Client.Do(req)
}

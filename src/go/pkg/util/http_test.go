package util_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestHttp(t *testing.T) {
	// req, _ := http.NewRequest(http.MethodGet, "https://www.paymentrails.com", strings.NewReader(""))
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8000", strings.NewReader(""))

	client := util.NewHttpClient(context.Background())

	res, err := client.Do(req)

	// require.Nil(t, err, "Client returned err", err)
	if err != nil {
		require.NotNil(t, err)
	} else {
		require.Equal(t, res.StatusCode, 200, "Non-200 response")
	}
}

package freelancer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientOptions(t *testing.T) {
	c := NewClient("token", WithSandBox(), WithDebug(true))

	assert.Equal(t, c.baseURL, BaseAPISandBoxURL)
	assert.Equal(t, c.debugMode, true)
}

func TestClientDoSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	data, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `{"message": "success"}`)
}

func TestClientDoError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	_, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	apiErr, ok := err.(*APIError)
	assert.True(t, ok)
	assert.Equal(t, apiErr.StatusCode, 400)
	assert.Equal(t, "error", apiErr.Message)
}

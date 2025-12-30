package freelancer

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestNewClientOptions(t *testing.T) {
	c := NewClient("token", WithSandBox(), WithDebug(true))

	assert.Equal(t, c.baseURL, endpoints.BaseAPISandBoxURL)
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

func TestClientHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "token", r.Header.Get("freelancer-oauth-v1"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "GoFreelancerSDK/1.2 (+github.com/cushydigit/go-freelancer-sdk)", r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	_, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
}

func TestClientQueryParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		assert.Equal(t, "value1", values.Get("param1"))
		assert.Equal(t, "value2", values.Get("param"))
		assert.Equal(t, true, values.Get("param3"))
		assert.Equal(t, 1, values.Get("param4"))
		assert.Equal(t, 1.3, values.Get("param5"))
		assert.Equal(t, 1092, values.Get("param6"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	query := url.Values{}
	addString(query, "param1", String("value1"))
	addString(query, "param2", String("value2"))
	addBool(query, "param3", Bool(true))
	addInt(query, "param4", Int(1))
	addFloat(query, "param5", Float64(1.3))
	addInt64(query, "param6", Int64(1092))

}

func TestClientRequestBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, `{"name": "Alice", "age": 30}`, string(body))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	_, err := c.do(context.Background(), http.MethodPost, "/test", nil, bytes.NewReader([]byte(`{"name": "Alice", "age": 30}`)))
	assert.NoError(t, err)
}

func TestExecuteInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{message : success}`))
	}))
	defer ts.Close()

	type Person struct {
		Name string
		Age  int
	}

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	_, err := execute[Person](context.Background(), c, http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode error")
}

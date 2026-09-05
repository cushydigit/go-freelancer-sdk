package freelancer

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestNewClientOptions(t *testing.T) {
	c := NewClient("token", WithSandBox(), WithDebug(true))

	assert.Equal(t, c.baseURL, endpoints.APISandBoxURL)
	assert.Equal(t, c.debugMode, true)
}

func TestClientDoSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()), WithDebug(true))
	c.SetBaseUrl(ts.URL)

	data, _, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
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

	_, _, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
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
		assert.Equal(t, "GoFreelancerSDK/1.4 (+github.com/cushydigit/go-freelancer-sdk)", r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	_, _, err := c.do(context.Background(), http.MethodGet, "/test", nil, nil)
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

	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	q := query.Values(struct {
		Param1 string
		Param2 string
		Param3 bool
		Param4 int
		Param5 float64
		Param6 int64
	}{
		Param1: "value1",
		Param2: "value2",
		Param3: true,
		Param4: 1,
		Param5: 1.3,
		Param6: 1092,
	})

	c.do(context.Background(), http.MethodGet, "test", q, nil)
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
	_, _, err := c.do(context.Background(), http.MethodPost, "/test", nil, bytes.NewReader([]byte(`{"name": "Alice", "age": 30}`)))
	assert.NoError(t, err)
}

func TestClient_GetBaseUrl(t *testing.T) {
	c := NewClient("token")
	assert.Equal(t, endpoints.APIMainURL, c.GetBaseUrl())
	c.SetBaseUrl(endpoints.APISandBoxURL)
	assert.Equal(t, endpoints.APISandBoxURL, c.GetBaseUrl())
}

func TestClient_NewClient(t *testing.T) {
	apiToken := "token"
	opts := []ClientOption{
		WithSandBox(),
		WithHttpClient(&http.Client{}),
		WithDebug(true),
	}
	c := NewClient(apiToken, opts...)
	if c.httpClient == nil || !c.debugMode || c.baseURL != endpoints.APISandBoxURL {
		t.Errorf("unexpected configuration: %v", c)
	}
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

	_, _, err := execute[Person](context.Background(), c, http.MethodGet, "/test", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode error")
}

func TestExecuteGeneric(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Name":"Alice","Age":30}`))
	}))
	defer ts.Close()

	type Person struct {
		Name string
		Age  int
	}

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	result, _, err := execute[Person](context.Background(), c, http.MethodGet, "/test", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", result.Name)
	assert.Equal(t, 30, result.Age)
}

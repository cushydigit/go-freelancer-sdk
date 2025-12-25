package freelancer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_CallAPI_ErrorHandling(t *testing.T) {
	// setup a mock server that returns 404 error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"code": "not_found", "message": "Project not found"}`))
	}))
	defer server.Close()

	client := NewClient("fake_token")
	client.SetBaseUrl(server.URL)

	req := &request{
		method:   http.MethodGet,
		endpoint: "/test",
	}

	_, err := client.callAPI(context.Background(), req)

	if err == nil {
		t.Fatal("expected an error for 404 status code, got nil")
	}

	if _, ok := err.(*APIError); !ok {
		t.Fatalf("expected an APIError, got %T", err)
	}

}

package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestProjectsService_Create(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Method
		assert.Equal(t, http.MethodPost, r.Method)

		// Path
		assert.Equal(t, endpoints.Projects, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"request_id": "123456789",
			"result": {
				"id": 123,
				"owner_id": 45,
				"title": "My Project",
				"status": "active",
				"deleted": false
			}
		}`))
	}))

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	body := CreateProjectBody{
		Title:       "My Project",
		Description: "Test",
	}

	resp, err := c.Services.Projects.Create(context.Background(), body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(123), resp.Result.ID)
	assert.Equal(t, "My Project", resp.Result.Title)
	assert.False(t, resp.Result.Deleted)
	defer ts.Close()

}

func TestProjectsService_SearchActive_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, endpoints.ProjectsActive, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"request_id": "123456789",
			"result": {
				"projects": [
					{
						"id": 123,
						"owner_id": 45,
						"title": "My Project",
						"status": "active",
						"deleted": false
					}
				]
			}
		}`))
	}))

	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	resp, err := c.Services.Projects.SearchActive(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(123), resp.Result.Projects[0].ID)
	assert.Equal(t, "My Project", resp.Result.Projects[0].Title)
	assert.False(t, resp.Result.Projects[0].Deleted)
}

func TestProjectsService_SearchActive_ScalarParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		assert.Equal(t, "python golang", q.Get("query"))
		assert.Equal(t, "100.5", q.Get("min_price"))
		assert.Equal(t, "true", q.Get("full_description"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"projects": []
			}
		}`))
	}))

	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	opts := &SearchActiveProjectsOptions{
		Query:           String("python golang"),
		MinPrice:        Float64(100.5),
		FullDescription: Bool(true),
	}

	res, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Result.Projects), 0)

}

func TestProjectsService_SearchActive_ArrayParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		fmt.Println(r.URL.String())

		assert.ElementsMatch(t, []string{"1", "2"}, q["jobs[]"])
		assert.ElementsMatch(t, []string{"US", "DE"}, q["countries[]"])
		assert.ElementsMatch(t, []string{"en", "de"}, q["languages[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"projects": []
			}
		}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	opts := &SearchActiveProjectsOptions{
		Jobs:      []int64{1, 2},
		Countries: []string{"US", "DE"},
		Languages: []string{"en", "de"},
	}

	_, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)

}

func TestProjectsService_SearchActive_EnumParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		assert.Equal(t, string(SortFieldsTimeUpdated), q.Get("sort_field"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"projects": []
			}
		}`))
		// method
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	opts := &SearchActiveProjectsOptions{
		SortField: Enum(SortFieldsTimeUpdated),
	}

	res, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectsService_SearchActive_OmitNilParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.URL.Query())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"projects": []
			}
		}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	_, err := c.Services.Projects.SearchActive(context.Background(), &SearchActiveProjectsOptions{})
	assert.NoError(t, err)

}

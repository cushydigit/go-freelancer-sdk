package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/utils"
	"github.com/stretchr/testify/assert"
)

func TestProjectsService_Create_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)

		// path
		assert.Equal(t, endpoints.Projects, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Create(context.Background(), utils.CreateProjectBody{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectsService_Create_Body(t *testing.T) {
	b := utils.CreateProjectBody{
		Title:       "Project Test Title",
		Description: "Project Description Test",
		Budget: utils.Budget{
			Minimum: 10.5,
			Maximum: utils.Float64(100.6),
		},
		Jobs: []int64{1, 2},
		Type: utils.Enum(utils.ProjectFixed),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res utils.CreateProjectBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, res.Title, b.Title)
		assert.Equal(t, res.Description, b.Description)
		assert.Equal(t, res.Budget.Minimum, b.Budget.Minimum)
		assert.Equal(t, res.Budget.Maximum, b.Budget.Maximum)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	_, err := c.Services.Projects.Create(context.Background(), b)
	assert.NoError(t, err)

}

func TestProjectsService_Create_Response(t *testing.T) {
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
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	body := utils.CreateProjectBody{
		Title:       "My Project",
		Description: "Test",
	}

	resp, err := c.Services.Projects.Create(context.Background(), body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(123), resp.Result.ID)
	assert.Equal(t, body.Title, resp.Result.Title)
	assert.False(t, resp.Result.Deleted)
}

func TestProjectsService_Action_Base(t *testing.T) {
	projectID := 100
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)

		// path
		expected := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
		assert.Equal(t, expected, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Action(context.Background(), int64(projectID), utils.ActionProject{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
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

	opts := &utils.SearchActiveProjectsOptions{
		Query:           utils.String("python golang"),
		MinPrice:        utils.Float64(100.5),
		FullDescription: utils.Bool(true),
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

	opts := &utils.SearchActiveProjectsOptions{
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

		assert.Equal(t, string(utils.SortFieldsTimeUpdated), q.Get("sort_field"))

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

	opts := &utils.SearchActiveProjectsOptions{
		SortField: utils.Enum(utils.SortFieldsTimeUpdated),
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

	_, err := c.Services.Projects.SearchActive(context.Background(), &utils.SearchActiveProjectsOptions{})
	assert.NoError(t, err)

}

func TestProjectsService_SearchActive_ReturnsPointers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a mock JSON response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"projects": [
					{"id": 1, "title": "Project 1"},
					{"id": 2, "title": "Project 2"}
				],
				"total_count": 2
			}
		}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	opts := &utils.SearchActiveProjectsOptions{}

	resp, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, resp.Result.TotalCount, 2)

	// Check pointer types for projects
	for _, p := range resp.Result.Projects {
		assert.NotNil(t, p)
		assert.Equal(t, reflect.TypeOf(p).Kind(), reflect.Ptr)
	}
}

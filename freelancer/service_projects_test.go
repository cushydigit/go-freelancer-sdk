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
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
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

	res, err := c.Services.Projects.Create(context.Background(), rr.CreateProjectBody{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectsService_Create_Body(t *testing.T) {
	b := rr.CreateProjectBody{
		Title:       "Project Test Title",
		Description: "Project Description Test",
		Budget: rr.Budget{
			Minimum: 10.5,
			Maximum: *rr.Float64(100.2),
		},
		Jobs: []int64{1, 2},
		Type: rr.Enum(rr.ProjectBudgetFixed),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res rr.CreateProjectBody
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

	body := rr.CreateProjectBody{
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

	res, err := c.Services.Projects.Action(context.Background(), int64(projectID), rr.ActionProject{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectsService_Action_Body(t *testing.T) {
	action := rr.ActionProject{
		ProjectID: 100,
		Action:    rr.ProjectActionUpgrade,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check the body
		var a rr.ActionProject
		err := json.NewDecoder(r.Body).Decode(&a)
		assert.NoError(t, err)
		defer r.Body.Close()
		assert.Equal(t, action.Action, a.Action)
		assert.Equal(t, action.ProjectID, a.ProjectID)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Action(context.Background(), action.ProjectID, action)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectsService_List_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.Projects, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, err := c.Services.Projects.List(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectsService_List_Options(t *testing.T) {
	opts := rr.ListProjectsOptions{
		Projects:                []int64{100, 101, 103},
		FrontendProjectStatuses: []rr.ProjectFrontendStatus{rr.ProjectFrontendStatusComplete, rr.ProjectFrontendStatusDraft},
		FullDescription:         rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		assert.ElementsMatch(t, []string{"100", "101", "103"}, q["projects[]"])
		assert.ElementsMatch(t, []string{string(opts.FrontendProjectStatuses[0]), string(opts.FrontendProjectStatuses[1])}, q["frontend_project_statuses[]"])
		assert.Equal(t, "true", q.Get("full_description"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectsService_ListSelf_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, endpoints.ProjectsSelf, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ListSelf(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_ListSelf_Options(t *testing.T) {
	opts := rr.ListSelfProjectsOptions{
		Status: rr.Enum(rr.ProjectStatusActive),
		Types:  []rr.ProjectType{rr.Projects, rr.Contests},
		Query:  rr.String("python golang"),
		Offset: rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Equal(t, string(*opts.Status), q.Get("status"))
		assert.Equal(t, string(*opts.Query), q.Get("query"))
		assert.Equal(t, "10", q.Get("offset"))
		assert.ElementsMatch(t, []string{string(opts.Types[0]), string(opts.Types[1])}, q["type[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ListSelf(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectService_Get_Base(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Get(context.Background(), projectID, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_Get_Options(t *testing.T) {
	projectID := int64(102)
	opts := rr.GetProjectOptions{
		FullDescription: rr.Bool(true),
		Limit:           rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Equal(t, "true", q.Get("full_description"))
		assert.Equal(t, "10", q.Get("limit"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Get(context.Background(), projectID, &opts)
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

	opts := &rr.SearchActiveProjectsOptions{
		Query:           rr.String("python golang"),
		MinPrice:        rr.Float64(100.5),
		FullDescription: rr.Bool(true),
	}

	res, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Result.Projects), 0)

}

func TestProjectsService_SearchActive_ArrayParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

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

	opts := &rr.SearchActiveProjectsOptions{
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

		assert.Equal(t, string(rr.SortFieldsTimeUpdated), q.Get("sort_field"))

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

	opts := &rr.SearchActiveProjectsOptions{
		SortField: rr.Enum(rr.SortFieldsTimeUpdated),
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

	_, err := c.Services.Projects.SearchActive(context.Background(), &rr.SearchActiveProjectsOptions{})
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

	opts := &rr.SearchActiveProjectsOptions{}

	resp, err := c.Services.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, resp.Result.TotalCount, 2)

	// Check pointer types for projects
	for _, p := range resp.Result.Projects {
		assert.NotNil(t, p)
		assert.Equal(t, reflect.TypeOf(p).Kind(), reflect.Ptr)
	}
}

func TestProjectService_SearchAll_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsAll, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.SearchAll(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_SearchAll_Options(t *testing.T) {
	opts := rr.SearchAllProjectsOptions{
		Query:        rr.String("golang excel"),
		ProjectTypes: []rr.ProjectBudgetType{rr.ProjectBudgetFixed},
		Jobs:         []int64{1, 2, 3},
		MaxPrice:     rr.Float64(100.1),
		Countries:    []string{"USA"},
		ReverseSort:  rr.Bool(false),
		Limit:        rr.Int(9),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Equal(t, "false", q.Get("reverse_sort"))
		assert.Equal(t, *opts.Query, q.Get("query"))
		assert.Equal(t, "100.1", q.Get("max_price"))
		assert.ElementsMatch(t, []string{string(rr.ProjectBudgetFixed)}, q["project_types[]"])
		assert.ElementsMatch(t, []string{"USA"}, q["countries[]"])
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["jobs[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.SearchAll(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_InviteFreelancer_Base(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		path := fmt.Sprintf("%s/%d/invite", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, err := c.Services.Projects.InviteFreelancer(context.Background(), projectID, rr.InviteFreelancersBody{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_InviteFreelancer_Body(t *testing.T) {
	projectID := int64(100)
	body := rr.InviteFreelancersBody{
		FreelancerID: int64(20000),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		assert.NotNil(t, r.Body)
		var req rr.InviteFreelancersBody
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)
		assert.Equal(t, body.FreelancerID, req.FreelancerID)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, err := c.Services.Projects.InviteFreelancer(context.Background(), projectID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

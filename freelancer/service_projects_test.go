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

	res, err := c.Services.Projects.Action(context.Background(), int64(projectID), rr.ActionProjectBody{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectsService_Action_Body(t *testing.T) {
	action := rr.ActionProjectBody{
		ProjectID: 100,
		Action:    rr.ProjectActionUpgrade,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check the body
		var a rr.ActionProjectBody
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

func TestProjectService_ListUpgradesFees(t *testing.T) {
	opts := rr.ListUpgradesFeesOptions{
		Currencies:  []int64{1, 2, 3},
		Project:     rr.Int64(100),
		TaxIncluded: rr.Bool(false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsFees, r.URL.Path)
		// options
		assert.Equal(t, "false", q.Get("tax_included"))
		assert.Equal(t, "100", q.Get("project"))
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["currencies[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, err := c.Services.Projects.ListUpgradesFees(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectService_ListBids(t *testing.T) {
	projectID := int64(100)
	opts := rr.ListProjectBidsOptions{
		IsShortlisted: rr.Bool(true),
		Limit:         rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/bids", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("is_shortlisted"))
		assert.Equal(t, "10", q.Get("limit"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ListBids(context.Background(), projectID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectService_GetBidInfo(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/bids_info", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.GetBidInfo(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_ListMilestones(t *testing.T) {
	projectID := int64(100)
	opts := rr.ListProjectMilestonesOptions{
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/milestones", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ListMilestones(context.Background(), projectID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_List(t *testing.T) {
	projectID := int64(100)
	opts := rr.ListProjectsMilestoneRequestsOptions{
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/milestone_requests", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ListMilestoneRequests(context.Background(), projectID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_GetHourlyContractInfo(t *testing.T) {
	opts := rr.GetHourlyContractInfoOptions{
		ProjectIDs:     []int64{1, 2, 3},
		BillingDetails: rr.Bool(false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsHourlyContract, r.URL.Path)
		// options
		assert.Equal(t, "false", q.Get("billing_details"))
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["project_ids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.GetHourlyContractInfo(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_GetIPContractInfo(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/ip_contract_info", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.GetIPContractInfo(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_Delete(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodDelete, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Delete(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCollaborationService_List(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/collaborations", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Collaborations.List(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
func TestCollaborationService_Create(t *testing.T) {
	projectID := int64(100)
	body := rr.CreateCollaborationBody{
		Email: "test@test.com",
		Permissions: rr.Permissions{
			Chat:     true,
			BidAward: false,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/collaborations", endpoints.Projects, projectID)
		assert.Equal(t, path, r.URL.Path)
		// body
		assert.NotNil(t, r.Body)
		defer r.Body.Close()
		var res rr.CreateCollaborationBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Email, res.Email)
		assert.Equal(t, body.Permissions.Chat, res.Permissions.Chat)
		assert.Equal(t, body.Permissions.BidAward, res.Permissions.BidAward)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Collaborations.Create(context.Background(), projectID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCollaborationService_Action(t *testing.T) {
	projectID := int64(100)
	collaborationID := int64(4)
	body := rr.ActionCollaborationBody{
		Action: rr.CollaborationActionRevoke,
		Permissions: rr.Permissions{
			Chat:     true,
			BidAward: false,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/collaborations/%d/actions", endpoints.Projects, projectID, collaborationID)
		assert.Equal(t, path, r.URL.Path)
		// body
		assert.NotNil(t, r.Body)
		defer r.Body.Close()
		var res rr.ActionCollaborationBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Action, res.Action)
		assert.Equal(t, body.Permissions.Chat, res.Permissions.Chat)
		assert.Equal(t, body.Permissions.BidAward, res.Permissions.BidAward)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Collaborations.Action(context.Background(), projectID, collaborationID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCollaborationService_ListAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsCollaborations, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Collaborations.ListAll(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServicesService_Order(t *testing.T) {
	serviceID := int64(150)
	serviceType := rr.ServiceLocal
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := fmt.Sprintf("%s/%s/%d/order", endpoints.ProjectsServices, serviceType, serviceID)
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Services.Order(context.Background(), serviceID, serviceType)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestServicesService_List(t *testing.T) {
	opts := rr.ListServicesOptions{
		Services: []int64{1, 2},
		Statuses: []rr.ServiceStatusType{rr.ServiceStatusActive},
		Titles:   []string{"t1", "t2"},
		Compact:  rr.Bool(true),
		Offset:   rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsServices, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["services[]"])
		assert.ElementsMatch(t, []string{string(rr.ServiceStatusActive)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"t1", "t2"}, q["titles[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Services.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestServicesService_SearchActive(t *testing.T) {
	opts := rr.SearchActiveServicesOptions{
		Query:   rr.String("test"),
		Sort:    rr.Enum(rr.SortNewest),
		Compact: rr.Bool(true),
		Offset:  rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsServicesActive, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))
		assert.Equal(t, "test", q.Get("query"))
		assert.Equal(t, string(rr.SortNewest), q.Get("sort"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Services.SearchActive(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_List(t *testing.T) {
	opts := rr.ListBidsOptions{
		Bids:          []int64{1, 2},
		AwardStatuses: []rr.BidAwardStatus{rr.BidAwardStatusAwarded},
		Compact:       rr.Bool(true),
		Offset:        rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsBids, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["bids[]"])
		assert.ElementsMatch(t, []string{string(rr.BidAwardStatusAwarded)}, q["award_statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_Get(t *testing.T) {
	bidID := int64(100)
	opts := rr.GetBidOptions{
		Compact: rr.Bool(true),
		Offset:  rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.Get(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_Create(t *testing.T) {
	body := rr.CreateBidBody{
		ProjectID:   int64(100),
		BidderID:    int64(200),
		Amount:      100,
		Description: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsBids, r.URL.Path)
		// body
		var res rr.CreateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.BidderID, res.BidderID)
		assert.Equal(t, body.Amount, res.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_Action(t *testing.T) {
	bidID := int64(100)
	body := rr.ActionBidBody{
		Action: rr.BidActionAward,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.Action(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_Update(t *testing.T) {
	bidID := int64(100)
	body := rr.UpdateBidBody{
		Amount:      100,
		Description: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.UpdateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Amount, res.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.Update(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_GetTimeTracking(t *testing.T) {
	bidID := int64(100)
	opts := rr.GetTimeTrackingOptions{
		Invoiced: rr.Bool(false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "false", q.Get("invoiced"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.GetTimeTracking(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidsService_CreateTimeTracking(t *testing.T) {
	bidID := int64(100)
	body := rr.CreateTimeTrackingBody{
		Seconds: 1000,
		Note:    "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path

		path := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.CreateTimeTrackingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Seconds, res.Seconds)
		assert.Equal(t, body.Note, res.Note)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Bids.CreateTimeTracking(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidEditRequestsService_List(t *testing.T) {
	bidID := int64(100)
	opts := rr.ListBidEditRequestsOptions{
		Statuses:          []rr.BidStatus{rr.BidStatusAccepted},
		BidEditRequestIDs: []int64{1, 2},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/edit_requests", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.ElementsMatch(t, []string{string(rr.BidStatusAccepted)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["bid_edit_request_ids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidEditRequests.List(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidEditRequestsService_Create(t *testing.T) {
	body := rr.CreateBidEditRequestBody{
		BidID:     int64(1),
		NewAmount: 100,
		Comment:   "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsBidEditRequests, r.URL.Path)
		// body
		var res rr.CreateBidEditRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.BidID, res.BidID)
		assert.Equal(t, body.NewAmount, res.NewAmount)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidEditRequests.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidEditRequestsService_Action(t *testing.T) {
	bidID := int64(100)
	bidEditRequestID := int64(200)
	body := rr.ActionBidEditRequestBody{
		Action: rr.BidEditRequestActionAccept,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/edit_requests/%d", endpoints.ProjectsBids, bidID, bidEditRequestID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionBidEditRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidEditRequests.Action(context.Background(), bidID, bidEditRequestID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidRatingsService_Get(t *testing.T) {
	bidID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidRatings.Get(context.Background(), bidID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidRatingsService_GetByListOfBids(t *testing.T) {
	opts := rr.GetByListOfBidsOptions{
		Bids: []int64{1, 2, 3},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsBidRatings, r.URL.Path)
		// options
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["bids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidRatings.GetByListOfBids(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidRatingsService_Create(t *testing.T) {
	bidID := int64(199)
	body := rr.CreateBidRatingBody{
		Rating:  10,
		Comment: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.CreateBidRatingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Rating, res.Rating)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidRatings.Create(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBidRatingsService_Update(t *testing.T) {
	bidID := int64(199)
	bidRatingID := int64(9)
	body := rr.UpdateBidRatingBody{
		Rating:  10,
		Comment: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d/bid_ratings/%d", endpoints.ProjectsBids, bidID, bidRatingID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.UpdateBidRatingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Rating, res.Rating)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.BidRatings.Update(context.Background(), bidID, bidRatingID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestJobsService_List(t *testing.T) {
	opts := rr.ListJobsOptions{
		Jobs:      []int64{1, 2},
		JobNames:  []string{"test", "test2"},
		OnlyLocal: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsJobs, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("only_local"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["jobs[]"])
		assert.ElementsMatch(t, []string{"test", "test2"}, q["job_names[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Jobs.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestJobsService_Search(t *testing.T) {
	opts := rr.SearchJobsOptions{
		Jobs:      []int64{1, 2},
		JobNames:  []string{"test", "test2"},
		OnlyLocal: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsJobsSearch, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("only_local"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["jobs[]"])
		assert.ElementsMatch(t, []string{"test", "test2"}, q["job_names[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Jobs.Search(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestJobBundles_List(t *testing.T) {
	opts := rr.ListJobBundlesOptions{
		JobBundles: []int64{3, 4},
		Lang:       rr.String("en"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsJobBundles, r.URL.Path)
		// options
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"3", "4"}, q["job_bundles[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.JobBundles.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestJobBundleCategories_List(t *testing.T) {
	opts := rr.ListJobBundleCategoriesOptions{
		JobBundles: []int64{3, 4},
		Lang:       rr.String("en"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsJobBundleCategories, r.URL.Path)
		// options
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"3", "4"}, q["job_bundles[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.JobBundleCategories.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestonesService_List(t *testing.T) {
	opts := rr.ListMilestonesOptions{
		Projects:   []int64{1, 2},
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		SortField:  rr.Enum(rr.SortFieldsBidAvgUsd),
		UserStatus: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsMilestones, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Milestones.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestonesService_Get(t *testing.T) {
	milestoneID := int64(100)
	opts := rr.GetMilestoneOptions{
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Milestones.Get(context.Background(), milestoneID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestonesService_Create(t *testing.T) {
	body := rr.CreateMilestoneBody{
		ProjectID:   int64(100),
		BidderID:    int64(199),
		Amount:      200,
		Reason:      rr.MilestoneCreateReasonFullPayment,
		Description: "test2",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsMilestones, r.URL.Path)
		// body
		var res rr.CreateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.BidderID, res.BidderID)
		assert.Equal(t, body.Reason, body.Reason)
		assert.Equal(t, body.Amount, body.Amount)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Milestones.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestonesService_Action(t *testing.T) {
	milestoneID := int64(100)
	body := rr.ActionMilestoneBody{
		Action: rr.MilestoneActionCancel,
		Amount: 200,
		Reason: rr.MilestoneActionReasonAccidentallyCreated,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path

		path := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionMilestoneBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Action, res.Action)
		assert.Equal(t, body.Reason, body.Reason)
		assert.Equal(t, body.Amount, body.Amount)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Milestones.Action(context.Background(), milestoneID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestoneRequestsService_List(t *testing.T) {
	opts := rr.ListMilestoneRequestsOptions{
		Projects:   []int64{1, 2},
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		SortField:  rr.Enum(rr.SortFieldsBidAvgUsd),
		UserStatus: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsMilestoneRequests, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.MilestoneRequests.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestoneRequestsService_Get(t *testing.T) {
	milestoneRequestID := int64(100)
	opts := rr.GetMilestoneRequestOptions{
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.MilestoneRequests.Get(context.Background(), milestoneRequestID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestoneRequestsService_Create(t *testing.T) {
	body := rr.CreateMilestoneRequestBody{
		ProjectID:   int64(100),
		Amount:      200,
		Description: "test2",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsMilestoneRequests, r.URL.Path)
		// body
		var res rr.CreateMilestoneRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.Amount, body.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.MilestoneRequests.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestoneRequestsService_Action(t *testing.T) {
	milestoneRequestID := int64(100)
	body := rr.ActionMilestoneRequestBody{
		Action: rr.MilestoneActionRequestAccept,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionMilestoneRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.MilestoneRequests.Action(context.Background(), milestoneRequestID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviewsService_List(t *testing.T) {
	opts := rr.ListReviewsOptions{
		Projects:    []int64{1, 2},
		ReviewTypes: []rr.ReviewType{rr.ReviewTypeProject},
		UserStatus:  rr.Bool(true),
		Limit:       rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsReviews, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{string(rr.ReviewTypeProject)}, q["review_types[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Reviews.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviewsService_Create(t *testing.T) {
	body := rr.CreateReviewBody{
		ProjectID:  int64(100),
		Comment:    "test",
		ReviewType: rr.ReviewTypeProject,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsReviews, r.URL.Path)
		// body
		var res rr.CreateReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.Comment, body.Comment)
		assert.Equal(t, body.ReviewType, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Reviews.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReviewsService_Action(t *testing.T) {
	reviewID := int64(100)
	body := rr.ActionReviewBody{
		Action:     rr.ReviewActionFeature,
		ReviewType: rr.ReviewTypeContest,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsReviews, reviewID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionReviewBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Action, res.Action)
		assert.Equal(t, body.ReviewType, res.ReviewType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Reviews.Action(context.Background(), reviewID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertGuaranteesService_List(t *testing.T) {
	opts := rr.ListExpertGuaranteesOptions{
		Projects: []int64{1, 2},
		Limit:    rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsExpertGuarantees, r.URL.Path)
		// options
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ExpertGuarantees.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertGuaranteesService_Action(t *testing.T) {
	expertGuaranteesID := int64(100)
	body := rr.ActionExpertGuaranteesBody{
		Action: rr.ExpertGuaranteesActionRelease,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.ProjectsExpertGuarantees, expertGuaranteesID)
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.ActionExpertGuaranteesBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.ExpertGuarantees.Action(context.Background(), expertGuaranteesID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertCurrencies_List(t *testing.T) {
	opts := rr.ListCurrenciesOptions{
		CurrencyCodes:             []string{"usd", "cad"},
		CurrencyIDs:               []int64{1, 2},
		IncludeExternalCurrencies: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsCurrencies, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("include_external_currencies"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["currency_ids[]"])
		assert.ElementsMatch(t, []string{"usd", "cad"}, q["currency_codes[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Currencies.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertCategories_List(t *testing.T) {
	opts := rr.ListCategoriesOptions{
		Categories: []int64{1, 2},
		Lang:       rr.String("en"),
		SeoDetails: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsCategories, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("seo_details"))
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["categories[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Categories.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertBudgets_List(t *testing.T) {
	opts := rr.ListBudgetsOptions{
		CurrencyCodes:   []string{"usd", "cad"},
		Lang:            rr.String("en"),
		CurrencyDetails: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, endpoints.ProjectsBudgets, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("currency_details"))
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"usd", "cad"}, q["currency_codes[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, err := c.Services.Projects.Budgets.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

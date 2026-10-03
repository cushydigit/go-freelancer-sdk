package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
	"github.com/stretchr/testify/assert"
)

func TestProjects_Create_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)

		// path
		assert.Equal(t, string(endpoints.Projects), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.Create(context.Background(), rr.CreateProjectBody{})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjects_Create_Body(t *testing.T) {
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
	_, _, err := c.Resources.Projects.Create(context.Background(), b)
	assert.NoError(t, err)

}

func TestProjects_Create_Response(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Method
		assert.Equal(t, http.MethodPost, r.Method)

		// Path
		assert.Equal(t, string(endpoints.Projects), r.URL.Path)

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

	resp, _, err := c.Resources.Projects.Create(context.Background(), body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(123), resp.Result.ID)
	assert.Equal(t, body.Title, resp.Result.Title)
	assert.False(t, resp.Result.Deleted)
}

func TestProjects_Actions(t *testing.T) {
	projectID := int64(100)
	bidID := int64(200)
	description := "updated project description"

	tests := []struct {
		name string
		call func(*Projects, context.Context, int64) (*rr.RawResponse, *ResponseMeta, error)
		test func(*testing.T, *http.Request)
	}{
		{
			name: "SignNDA",
			call: func(r *Projects, ctx context.Context, id int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.SignNDA(
					ctx,
					id,
					"John Doe",
					"123 Main St",
					"New York",
					"NY",
					"+123456789",
					"US",
				)
			},
			test: func(t *testing.T, req *http.Request) {
				var body rr.ActionProjectSignNDA
				err := json.NewDecoder(req.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, rr.ProjectActionSignNDA, body.Action)
				assert.Equal(t, "John Doe", body.FullName)
				assert.Equal(t, "123 Main St", body.Address)
				assert.Equal(t, "New York", body.City)
				assert.Equal(t, "NY", body.State)
				assert.Equal(t, "+123456789", body.Phone)
				assert.Equal(t, "US", body.Country)
			},
		},
		{
			name: "Upgrades",
			call: func(r *Projects, ctx context.Context, id int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.Upgrades(
					ctx,
					id,
					[]rr.ProjectUpgradeType{
						rr.ProjectUpgradeAssisted,
						rr.ProjectUpgradeFeatured,
						rr.ProjectUpgradeIpContract,
					},
				)
			},
			test: func(t *testing.T, req *http.Request) {
				var body rr.ActionProjectUpgrade
				err := json.NewDecoder(req.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, rr.ProjectActionUpgrade, body.Action)
				expected := []rr.ProjectUpgradeType{
					rr.ProjectUpgradeAssisted,
					rr.ProjectUpgradeFeatured,
					rr.ProjectUpgradeIpContract,
				}

				assert.Equal(t, expected, body.Upgrades)
			},
		},
		{
			name: "Update",
			call: func(r *Projects, ctx context.Context, id int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.Update(
					ctx,
					id,
					[]int64{10, 20, 30},
					&description,
				)
			},
			test: func(t *testing.T, req *http.Request) {
				var body rr.ActionProjectUpdate
				err := json.NewDecoder(req.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, rr.ProjectActionClose, body.Action)
				assert.Equal(t, []int64{10, 20, 30}, body.JonIDs)
				assert.Equal(t, description, *body.Description)
			},
		},
		{
			name: "Close",
			call: func(r *Projects, ctx context.Context, id int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.Close(ctx, id)
			},
			test: func(t *testing.T, req *http.Request) {
				var body rr.ActionProjectClose
				err := json.NewDecoder(req.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, rr.ProjectActionClose, body.Action)
			},
		},
		{
			name: "End",
			call: func(r *Projects, ctx context.Context, id int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.End(
					ctx,
					id,
					bidID,
					rr.ProjectEndStatusType("complete"),
				)
			},
			test: func(t *testing.T, req *http.Request) {
				var body rr.ActionProjectEnd
				err := json.NewDecoder(req.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, rr.ProjectActionEnd, body.Action)
				assert.Equal(t, bidID, body.BidID)
				assert.Equal(t, rr.ProjectEndStatusType("complete"), body.Status)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPut, r.Method)
				assert.Equal(t, string(endpoints.Project(projectID)), r.URL.Path)

				tt.test(t, r)

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"message":"ok"}`))
			}))
			defer ts.Close()

			c := NewClient(
				"token",
				WithHttpClient(ts.Client()),
			)
			c.SetBaseUrl(ts.URL)

			res, _, err := tt.call(
				&c.Resources.Projects,
				context.Background(),
				projectID,
			)

			assert.NoError(t, err)
			assert.NotNil(t, res)
		})
	}
}

func TestProjects_List_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Projects), r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, _, err := c.Resources.Projects.List(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjects_List_Options(t *testing.T) {
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

	res, _, err := c.Resources.Projects.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjects_ListSelf_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, string(endpoints.ProjectsSelf), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.ListSelf(context.Background(), nil)
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

	res, _, err := c.Resources.Projects.ListSelf(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectService_Get_Base(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Project(projectID)), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.Get(context.Background(), projectID, nil)
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

	res, _, err := c.Resources.Projects.Get(context.Background(), projectID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjects_SearchActive_Base(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, string(endpoints.ProjectsActive), r.URL.Path)

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

	resp, _, err := c.Resources.Projects.SearchActive(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(123), resp.Result.Projects[0].ID)
	assert.Equal(t, "My Project", resp.Result.Projects[0].Title)
	assert.False(t, resp.Result.Projects[0].Deleted)
}

func TestProjects_SearchActive_ScalarParams(t *testing.T) {
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

	res, _, err := c.Resources.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Result.Projects), 0)

}

func TestProjects_SearchActive_ArrayParams(t *testing.T) {
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

	_, _, err := c.Resources.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)

}

func TestProjects_SearchActive_EnumParams(t *testing.T) {
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

	res, _, err := c.Resources.Projects.SearchActive(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjects_SearchActive_OmitNilParams(t *testing.T) {
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

	_, _, err := c.Resources.Projects.SearchActive(context.Background(), &rr.SearchActiveProjectsOptions{})
	assert.NoError(t, err)

}

func TestProjects_SearchActive_ReturnsPointers(t *testing.T) {
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

	resp, _, err := c.Resources.Projects.SearchActive(context.Background(), opts)
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
		assert.Equal(t, string(endpoints.ProjectsAll), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.SearchAll(context.Background(), nil)
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

	res, _, err := c.Resources.Projects.SearchAll(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_InviteFreelancer_Base(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, string(endpoints.ProjectInvite(projectID)), r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)
	res, _, err := c.Resources.Projects.InviteFreelancer(context.Background(), projectID, rr.InviteFreelancersBody{})
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
	res, _, err := c.Resources.Projects.InviteFreelancer(context.Background(), projectID, body)
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
		assert.Equal(t, string(endpoints.ProjectsFees), r.URL.Path)
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
	res, _, err := c.Resources.Projects.ListUpgradesFees(context.Background(), &opts)
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
		assert.Equal(t, string(endpoints.ProjectBids(projectID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("is_shortlisted"))
		assert.Equal(t, "10", q.Get("limit"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.ListBids(context.Background(), projectID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProjectService_GetBidInfo(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.ProjectBidsInfo(projectID)), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.GetBidInfo(context.Background(), projectID)
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
		assert.Equal(t, string(endpoints.ProjectMilestones(projectID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.ListMilestones(context.Background(), projectID, &opts)
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
		assert.Equal(t, string(endpoints.ProjectMilestoneRequests(projectID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.ListMilestoneRequests(context.Background(), projectID, &opts)
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
		assert.Equal(t, string(endpoints.HourlyContractInfo), r.URL.Path)
		// options
		assert.Equal(t, "false", q.Get("billing_details"))
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["project_ids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.GetHourlyContractInfo(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_GetIPContractInfo(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.ProjectIPContractInfo(projectID)), r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.GetIPContractInfo(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestProjectService_Delete(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodDelete, r.Method)
		// path
		assert.Equal(t, string(endpoints.Project(projectID)), r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Projects.Delete(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

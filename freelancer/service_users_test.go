package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
	"github.com/stretchr/testify/assert"
)

func TestUsersService_List(t *testing.T) {
	opts := rr.ListUsersOptions{
		Users:     []int64{1, 2},
		Usernames: []string{"user1", "user2"},
		Avatar:    rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.Users
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("avatar"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["users[]"])
		assert.ElementsMatch(t, []string{"user1", "user2"}, q["usernames[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersService_Get(t *testing.T) {
	userID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := fmt.Sprintf("%s/%d", endpoints.Users, userID)
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Get(context.Background(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersService_SearchFreelancer(t *testing.T) {
	opts := rr.SearchFreelancerOptions{
		Query:     rr.String("golang python r"),
		JobIDs:    []int64{1, 2},
		Countries: []string{"USA", "GERMANY"},
		Limit:     rr.Int(10),
		Compact:   rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersFreelancers
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "10", q.Get("limit"))
		assert.Equal(t, "golang python r", q.Get("query"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["jobs[]"])
		assert.ElementsMatch(t, []string{"USA", "GERMANY"}, q["countries[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.SearchFreelancer(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelfService_GetInfo(t *testing.T) {
	opts := rr.GetSelfInfoOptions{
		Offset:  rr.Int(19),
		Compact: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersSelf
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "19", q.Get("offset"))
		assert.Equal(t, "true", q.Get("compact"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Self.GetInfo(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelfService_ListDevices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersSelfDevices
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Self.ListDevices(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelfJobs_Add(t *testing.T) {
	body := rr.AddJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := endpoints.UsersSelfJobs
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.AddJobsBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Jobs, res.Jobs)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.SelfJob.Add(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestSelfJobs_Set(t *testing.T) {
	body := rr.SetJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := endpoints.UsersSelfJobs
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.SetJobsBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Jobs, res.Jobs)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.SelfJob.Set(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestSelfJobs_Delete(t *testing.T) {
	body := rr.DeleteJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodDelete, r.Method)
		// path
		path := endpoints.UsersSelfJobs
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.DeleteJobsBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Jobs, res.Jobs)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.SelfJob.Delete(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProfiles_Create(t *testing.T) {
	body := rr.CreateProfileBody{
		Tagline:     "test",
		HourlyRate:  10,
		Description: "test2",
		ProfileName: "test3",
		SkillIDs:    []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := endpoints.UsersProfiles
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.CreateProfileBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.Tagline, res.Tagline)
		assert.Equal(t, body.HourlyRate, res.HourlyRate)
		assert.Equal(t, body.Description, res.Description)
		assert.Equal(t, body.ProfileName, res.ProfileName)
		assert.Equal(t, body.SkillIDs, res.SkillIDs)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Profiles.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProfiles_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersProfiles
		assert.Equal(t, path, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Profiles.Get(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProfiles_Update(t *testing.T) {
	body := rr.UpdateProfileBody{
		ProfileID:   int64(1),
		Tagline:     "test",
		HourlyRate:  10,
		Description: "test2",
		ProfileName: "test3",
		SkillIDs:    []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		path := endpoints.UsersProfiles
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.UpdateProfileBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProfileID, res.ProfileID)
		assert.Equal(t, body.Tagline, res.Tagline)
		assert.Equal(t, body.HourlyRate, res.HourlyRate)
		assert.Equal(t, body.Description, res.Description)
		assert.Equal(t, body.ProfileName, res.ProfileName)
		assert.Equal(t, body.SkillIDs, res.SkillIDs)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Profiles.Update(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestReputations_List(t *testing.T) {
	opts := rr.ListReputationsOptions{
		Users:      []int64{1, 2},
		Role:       rr.Enum(rr.RoleEmployer),
		JobHistory: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersReputations
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("job_history"))
		assert.Equal(t, string(rr.RoleEmployer), q.Get("role"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["users[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Reputations.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestEnterprises_List(t *testing.T) {
	opts := rr.ListEnterprisesOptions{
		Enterprises:   []int64{1, 2},
		InternalNames: []string{"test1", "test2"},
		UserID:        rr.Int64(100),
		Limit:         rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersEnterprises
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "100", q.Get("user_id"))
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["enterprises[]"])
		assert.ElementsMatch(t, []string{"test1", "test2"}, q["internal_names[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Enterprises.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestPortfolios_List(t *testing.T) {
	opts := rr.ListPortfoliosOptions{
		Users: []int64{1, 2},
		Limit: rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersPortfolios
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["users[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Portfolios.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestPools_List(t *testing.T) {
	opts := rr.ListPoolsOptions{
		Pools:      []int64{1, 2},
		IgnoreTest: rr.Bool(true),
		Limit:      rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		path := endpoints.UsersPools
		assert.Equal(t, path, r.URL.Path)
		// options
		assert.Equal(t, "10", q.Get("limit"))
		assert.Equal(t, "true", q.Get("ignore_test"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["pools[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Pools.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestViolations_Create(t *testing.T) {
	body := rr.CreateViolationBody{
		ContextID:        int64(100),
		ContextType:      rr.ViolationContextBid,
		ViolatorUserID:   int64(200),
		Reason:           rr.ViolationReasonAdvertising,
		AdditionalReason: rr.ViolationAdditionalReasonCopiedFromSomeoneElse,
		Comments:         "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		path := endpoints.UsersViolationReports
		assert.Equal(t, path, r.URL.Path)
		// body
		var res rr.CreateViolationBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ContextID, res.ContextID)
		assert.Equal(t, body.ContextType, res.ContextType)
		assert.Equal(t, body.Reason, res.Reason)
		assert.Equal(t, body.ViolatorUserID, res.ViolatorUserID)
		assert.Equal(t, body.AdditionalReason, res.AdditionalReason)
		assert.Equal(t, body.Comments, res.Comments)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Services.Users.Violations.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

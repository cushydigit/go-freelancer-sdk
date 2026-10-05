package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
	"github.com/stretchr/testify/assert"
)

func TestUsers_List(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Users), r.URL.Path)
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

	res, _, err := c.Resources.Users.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsers_Get(t *testing.T) {
	userID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.User(userID)), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Users.Get(context.Background(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsers_SearchFreelancer(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Freelancers), r.URL.Path)
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

	res, _, err := c.Resources.Users.SearchFreelancer(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsers_ListReputations(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Reputations), r.URL.Path)
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

	res, _, err := c.Resources.Users.ListReputations(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsers_ListEnterprises(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Enterprises), r.URL.Path)
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

	res, _, err := c.Resources.Users.ListEnterprises(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsers_ListPortfolios(t *testing.T) {
	opts := rr.ListPortfoliosOptions{
		Users: []int64{1, 2},
		Limit: rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Portfolios), r.URL.Path)
		// options
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["users[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Users.ListPortfolios(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelf_ListPools(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Pools), r.URL.Path)
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

	res, _, err := c.Resources.Self.ListPools(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUserService_CreateViolationsReport(t *testing.T) {
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
		assert.Equal(t, string(endpoints.ViolationReports), r.URL.Path)
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

	res, _, err := c.Resources.Users.CreateViolationReport(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

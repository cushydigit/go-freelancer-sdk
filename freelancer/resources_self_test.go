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

func TestSelf_Get(t *testing.T) {
	opts := rr.GetSelfInfoOptions{
		Offset:  rr.Int(19),
		Compact: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Self), r.URL.Path)
		// options
		assert.Equal(t, "19", q.Get("offset"))
		assert.Equal(t, "true", q.Get("compact"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Self.Get(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelf_ListDevices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Devices), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Self.ListDevices(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSelf_AddJobs(t *testing.T) {
	body := rr.AddJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.SelfJobs), r.URL.Path)
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

	res, _, err := c.Resources.Self.AddJobs(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestSelf_UpdateJobs(t *testing.T) {
	body := rr.SetJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.SelfJobs), r.URL.Path)
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

	res, _, err := c.Resources.Self.UpdateJobs(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestSelf_DeleteJobs(t *testing.T) {
	body := rr.DeleteJobsBody{
		Jobs: []int64{100, 101},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodDelete, r.Method)
		// path
		assert.Equal(t, string(endpoints.SelfJobs), r.URL.Path)
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

	res, _, err := c.Resources.Self.DeleteJobs(context.Background(), body)
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

	res, _, err := c.Resources.Self.ListProjects(context.Background(), nil)
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

	res, _, err := c.Resources.Self.ListProjects(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

package freelancer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"

	"github.com/stretchr/testify/assert"
)

func TestCommon_ListCountries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, string(endpoints.CommonCountries), r.URL.Path)

		// query
		assert.Equal(t, "true", r.URL.Query().Get("extra_details"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"countries": [
					{
						"code": "US",
						"name": "United States"
					},
					{
						"code": "CA",
						"name": "Canada"
					}
				]
			}
		}`))

	}))

	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	opts := &rr.ListCountriesOptions{
		ExtraDetails: rr.Bool(true),
	}

	res, _, err := c.Resources.Common.ListCountries(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Result.Countries, 2)
	assert.Equal(t, "US", res.Result.Countries[0].Code)

}

func TestCommon_ListTimezones(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, string(endpoints.CommonTimezones), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"result": {
				"timezones": [
					{
						"id": 241,
						"country": "NL",
						"timezone": "Europe/Amsterdam",
						"offset": 2
					}
				]
			}
		}`))

	}))

	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Common.ListTimezones(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Result.Timezones, 1)
	assert.Equal(t, "NL", res.Result.Timezones[0].Country)

}

func TestCommon_ListCurrencies(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Currencies), r.URL.Path)
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

	res, _, err := c.Resources.Common.ListCurrencies(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_ListCategories(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Categories), r.URL.Path)
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

	res, _, err := c.Resources.Common.ListCategories(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_ListBudgets(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Budgets), r.URL.Path)
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

	res, _, err := c.Resources.Common.ListBudgets(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_ListJobs(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Jobs), r.URL.Path)
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

	res, _, err := c.Resources.Common.ListJobs(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_SearchJobs(t *testing.T) {
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
		assert.Equal(t, string(endpoints.JobsSearch), r.URL.Path)
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

	res, _, err := c.Resources.Common.SearchJobs(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_ListJobBundles(t *testing.T) {
	opts := rr.ListJobBundlesOptions{
		JobBundles: []int64{3, 4},
		Lang:       rr.String("en"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.JobBundles), r.URL.Path)
		// options
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"3", "4"}, q["job_bundles[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Common.ListJobBundles(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestCommon_ListJobBundleCategories(t *testing.T) {
	opts := rr.ListJobBundleCategoriesOptions{
		JobBundles: []int64{3, 4},
		Lang:       rr.String("en"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.JobBundleCategories), r.URL.Path)
		// options
		assert.Equal(t, "en", q.Get("lang"))
		assert.ElementsMatch(t, []string{"3", "4"}, q["job_bundles[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Common.ListJobBundleCategories(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

package freelancer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestCommonService_ListCountries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, endpoints.CommonCountries, r.URL.Path)

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

	opts := &ListCountriesOptions{
		ExtraDetails: Bool(true),
	}

	res, err := c.Services.Common.ListCountries(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Result.Countries, 2)
	assert.Equal(t, "US", res.Result.Countries[0].Code)

}

func TestCommonService_Timezones(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)

		// path
		assert.Equal(t, endpoints.CommonTimezones, r.URL.Path)

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

	res, err := c.Services.Common.ListTimezones(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Result.Timezones, 1)
	assert.Equal(t, "NL", res.Result.Timezones[0].Country)

}

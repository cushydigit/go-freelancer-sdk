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

func TestExpertGuarantees_List(t *testing.T) {
	opts := rr.ListExpertGuaranteesOptions{
		Projects: []int64{1, 2},
		Limit:    rr.Int(10),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.ExpertGuarantees), r.URL.Path)
		// options
		assert.Equal(t, "10", q.Get("limit"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.ExpertGuarantees.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestExpertGuarantees_Actions(t *testing.T) {
	expertGuaranteesID := int64(100)
	tests := []struct {
		name   string
		action rr.ExpertGuaranteesAction
		call   func(*ExpertGuarantees, context.Context, int64) (*rr.RawResponse, *ResponseMeta, error)
	}{
		{
			name:   "Release",
			action: rr.ExpertGuaranteesActionRelease,
			call:   (*ExpertGuarantees).Release,
		},
		{
			name:   "RequestRelease",
			action: rr.ExpertGuaranteesActionRequestRelease,
			call:   (*ExpertGuarantees).RequestRelease,
		},
	}
	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// method
			assert.Equal(t, http.MethodPut, r.Method)
			// path
			assert.Equal(t, string(endpoints.ExpertGuarantee(expertGuaranteesID)), r.URL.Path)
			// body
			var res rr.ActionExpertGuaranteesBody
			err := json.NewDecoder(r.Body).Decode(&res)
			defer r.Body.Close()
			assert.NoError(t, err)
			assert.NotNil(t, r.Body)
			assert.Equal(t, tt.action, res.Action)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"ok"}`))

		}))
		defer ts.Close()

		c := NewClient("token", WithHttpClient(ts.Client()))
		c.SetBaseUrl(ts.URL)

		res, _, err := tt.call(
			&c.Resources.ExpertGuarantees,
			context.Background(),
			expertGuaranteesID,
		)
		assert.NoError(t, err)
		assert.NotNil(t, res)

	}
}

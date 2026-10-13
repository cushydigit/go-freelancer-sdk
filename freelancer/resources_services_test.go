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

func TestServices_Order(t *testing.T) {
	serviceID := int64(150)
	serviceType := rr.ServiceLocal
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.ServicesOrder(string(serviceType), serviceID)), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Services.Order(context.Background(), serviceID, serviceType)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestServices_List(t *testing.T) {
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
		assert.Equal(t, string(endpoints.Services), r.URL.Path)
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

	res, _, err := c.Resources.Services.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestServices_SearchActive(t *testing.T) {
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
		assert.Equal(t, string(endpoints.ServicesActive), r.URL.Path)
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

	res, _, err := c.Resources.Services.SearchActive(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

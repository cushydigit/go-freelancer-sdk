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
		assert.Equal(t, string(endpoints.Profiles), r.URL.Path)
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

	res, _, err := c.Resources.Profiles.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestProfiles_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Profiles), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Profiles.Get(context.Background())
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
		assert.Equal(t, string(endpoints.Profiles), r.URL.Path)
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

	res, _, err := c.Resources.Profiles.Update(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

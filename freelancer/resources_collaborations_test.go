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

func TestCollaborations_List(t *testing.T) {
	projectID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.ProjectCollaborations(projectID)), r.URL.Path)
		// options

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Collaborations.List(context.Background(), projectID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
func TestCollaborations_Create(t *testing.T) {
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
		assert.Equal(t, string(endpoints.ProjectCollaborations(projectID)), r.URL.Path)
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

	res, _, err := c.Resources.Collaborations.Create(context.Background(), projectID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestCollaborations_Action(t *testing.T) {
	projectID := int64(100)
	collaborationID := int64(4)
	chatPermission := true
	bidAwardPermission := false
	tests := []struct {
		name   string
		action rr.CollaborationAction
		call   func(*Collaborations, context.Context, int64, int64, bool, bool) (*rr.RawResponse, *ResponseMeta, error)
	}{
		{
			name:   "Revoke",
			action: rr.CollaborationActionRevoke,
			call:   (*Collaborations).Revoke,
		},
		{
			name:   "Update-Permissions",
			action: rr.CollaborationActionUpdatePermissions,
			call:   (*Collaborations).UpdatePermissions,
		},
	}
	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// method
			assert.Equal(t, http.MethodPut, r.Method)
			// path
			assert.Equal(
				t,
				string(endpoints.ProjectCollaborationsActions(projectID, collaborationID)),
				r.URL.Path,
			)
			// body
			assert.NotNil(t, r.Body)
			defer r.Body.Close()
			var res rr.ActionCollaborationBody
			err := json.NewDecoder(r.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tt.action, res.Action)
			assert.Equal(t, chatPermission, res.Permissions.Chat)
			assert.Equal(t, bidAwardPermission, res.Permissions.BidAward)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"ok"}`))
		}))
		defer ts.Close()

		c := NewClient("token", WithHttpClient(ts.Client()))
		c.SetBaseUrl(ts.URL)

		res, _, err := tt.call(
			&c.Resources.Collaborations,
			context.Background(),
			projectID, collaborationID,
			chatPermission, bidAwardPermission,
		)

		assert.NoError(t, err)
		assert.NotNil(t, res)

	}

}

func TestCollaborations_ListAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.ProjectsCollaborations), r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Collaborations.ListAll(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

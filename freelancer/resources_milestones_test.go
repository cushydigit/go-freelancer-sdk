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

func TestMilestones_List(t *testing.T) {
	opts := rr.ListMilestonesOptions{
		Projects:   []int64{1, 2},
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		SortField:  rr.Enum(rr.SortFieldsBidAvgUsd),
		UserStatus: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Milestones), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_Get(t *testing.T) {
	milestoneID := int64(100)
	opts := rr.GetMilestoneOptions{
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Milestone(milestoneID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.Get(context.Background(), milestoneID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_Create(t *testing.T) {
	body := rr.CreateMilestoneBody{
		ProjectID:   int64(100),
		BidderID:    int64(199),
		Amount:      200,
		Reason:      rr.MilestoneCreateReasonFullPayment,
		Description: "test2",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.Milestones), r.URL.Path)
		// body
		var res rr.CreateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.BidderID, res.BidderID)
		assert.Equal(t, body.Reason, body.Reason)
		assert.Equal(t, body.Amount, body.Amount)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_Actions(t *testing.T) {
	milestoneID := int64(100)
	tests := []struct {
		name         string
		action       rr.MilestoneAction
		amount       int                             // for release
		other_reason string                          // for update
		reason       rr.MilestoneRequestCancelReason // for requestCancel
		reason_text  string                          // for requestCancel when other selected

		call func(*Milestones, context.Context, int64) (*rr.RawResponse, *ResponseMeta, error)
	}{
		{
			name:   "Cancel",
			action: rr.MilestoneActionCancel,
			call:   (*Milestones).Cancel,
		},
		{
			name:   "RejectCancel",
			action: rr.MilestoneActionRejectCancel,
			call:   (*Milestones).RejectCancel,
		},
		{
			name:   "RequestRelease",
			action: rr.MilestoneActionRequestRelease,
			call:   (*Milestones).RequestRelease,
		},
		{
			name:   "Release",
			action: rr.MilestoneActionRelease,
			amount: 150,
			call: func(r *Milestones, ctx context.Context, milestoneID int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.Release(ctx, milestoneID, 150)
			},
		},
		{
			name:         "Update",
			action:       rr.MilestoneActionUpdate,
			other_reason: "some reason",
			call: func(r *Milestones, ctx context.Context, milestoneID int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.Update(ctx, milestoneID, "some reason")
			},
		},
		{
			name:        "RequestCancel",
			reason:      rr.MilestoneRequestCancelOther,
			reason_text: "bullshit",
			action:      rr.MilestoneActionRequestCancel,
			call: func(r *Milestones, ctx context.Context, milestoneID int64) (*rr.RawResponse, *ResponseMeta, error) {
				return r.RequestCancel(ctx, milestoneID, rr.MilestoneRequestCancelOther, "bullshit")
			},
		},
	}

	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// method
			assert.Equal(t, http.MethodPut, r.Method)
			// path

			assert.Equal(t, string(endpoints.Milestone(milestoneID)), r.URL.Path)
			// body
			var res rr.ActionMilestoneBody
			err := json.NewDecoder(r.Body).Decode(&res)
			defer r.Body.Close()
			assert.NoError(t, err)
			assert.NotNil(t, r.Body)
			assert.Equal(t, tt.action, res.Action)
			assert.Equal(t, tt.amount, res.Amount)
			assert.Equal(t, tt.reason, res.Reason)
			assert.Equal(t, tt.reason_text, res.ReasonText)
			assert.Equal(t, tt.other_reason, res.OtherReason)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"ok"}`))

		}))
		defer ts.Close()

		c := NewClient("token", WithHttpClient(ts.Client()))
		c.SetBaseUrl(ts.URL)

		res, _, err := tt.call(
			&c.Resources.Milestones,
			context.Background(),
			milestoneID,
		)
		assert.NoError(t, err)
		assert.NotNil(t, res)

	}
}

func TestMilestoneService_ListRequests(t *testing.T) {
	opts := rr.ListMilestoneRequestsOptions{
		Projects:   []int64{1, 2},
		Statuses:   []rr.MilestoneStatus{rr.MilestoneStatusCanceled},
		SortField:  rr.Enum(rr.SortFieldsBidAvgUsd),
		UserStatus: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.MilestoneRequests), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_status"))
		assert.ElementsMatch(t, []string{string(rr.MilestoneStatusCanceled)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["projects[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.ListRequests(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_GetRequest(t *testing.T) {
	milestoneRequestID := int64(100)
	opts := rr.GetMilestoneRequestOptions{
		UserAvatar: rr.Bool(true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.MilestoneRequest(milestoneRequestID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("user_avatar"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.GetRequest(context.Background(), milestoneRequestID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_CreateRequest(t *testing.T) {
	body := rr.CreateMilestoneRequestBody{
		ProjectID:   int64(100),
		Amount:      200,
		Description: "test2",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.MilestoneRequests), r.URL.Path)
		// body
		var res rr.CreateMilestoneRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		defer r.Body.Close()
		assert.NoError(t, err)
		assert.NotNil(t, r.Body)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.Amount, body.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Milestones.CreateRequest(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestMilestones_ActionRequests(t *testing.T) {
	milestoneRequestID := int64(100)
	tests := []struct {
		name   string
		action rr.MilestoneActionRequest
		call   func(*Milestones, context.Context, int64) (*rr.RawResponse, *ResponseMeta, error)
	}{
		{
			name:   "AcceptRequest",
			action: rr.MilestoneActionAcceptRequest,
			call:   (*Milestones).AcceptRequest,
		},
		{
			name:   "RejectRequest",
			action: rr.MilestoneActionRejectRequest,
			call:   (*Milestones).RejectRequest,
		},
		{
			name:   "DeleteRequest",
			action: rr.MilestoneActionDeleteRequest,
			call:   (*Milestones).DeleteRequest,
		},
	}
	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// method
			assert.Equal(t, http.MethodPut, r.Method)
			// path
			assert.Equal(t, string(endpoints.MilestoneRequest(milestoneRequestID)), r.URL.Path)
			// body
			var res rr.ActionMilestoneRequestBody
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
			&c.Resources.Milestones,
			context.Background(),
			milestoneRequestID,
		)

		assert.NoError(t, err)
		assert.NotNil(t, res)

	}
}

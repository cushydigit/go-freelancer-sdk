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

func TestBids_List(t *testing.T) {
	opts := rr.ListBidsOptions{
		Bids:          []int64{1, 2},
		AwardStatuses: []rr.BidAwardStatus{rr.BidAwardStatusAwarded},
		Compact:       rr.Bool(true),
		Offset:        rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Bids), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))
		assert.ElementsMatch(t, []string{"1", "2"}, q["bids[]"])
		assert.ElementsMatch(t, []string{string(rr.BidAwardStatusAwarded)}, q["award_statuses[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.List(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_Get(t *testing.T) {
	bidID := int64(100)
	opts := rr.GetBidOptions{
		Compact: rr.Bool(true),
		Offset:  rr.Int(3),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.Bid(bidID)), r.URL.Path)
		// options
		assert.Equal(t, "true", q.Get("compact"))
		assert.Equal(t, "3", q.Get("offset"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.Get(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_Create(t *testing.T) {
	body := rr.CreateBidBody{
		ProjectID:   int64(100),
		BidderID:    int64(200),
		Amount:      100,
		Description: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.Bids), r.URL.Path)
		// body
		var res rr.CreateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.ProjectID, res.ProjectID)
		assert.Equal(t, body.BidderID, res.BidderID)
		assert.Equal(t, body.Amount, res.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.Create(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_Action(t *testing.T) {
	bidID := int64(100)
	body := rr.ActionBidBody{
		Action: rr.BidActionAward,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Bid(bidID)), r.URL.Path)
		// body
		var res rr.ActionBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.Action(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_Update(t *testing.T) {
	bidID := int64(100)
	body := rr.UpdateBidBody{
		Amount:      100,
		Description: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.Bid(bidID)), r.URL.Path)
		// body
		var res rr.UpdateBidBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Amount, res.Amount)
		assert.Equal(t, body.Description, res.Description)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.Update(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_GetTimeTracking(t *testing.T) {
	bidID := int64(100)
	opts := rr.GetTimeTrackingOptions{
		Invoiced: rr.Bool(false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidTimeTracking(bidID)), r.URL.Path)
		// options
		assert.Equal(t, "false", q.Get("invoiced"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.GetTimeTracking(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_CreateTimeTracking(t *testing.T) {
	bidID := int64(100)
	body := rr.CreateTimeTrackingBody{
		Seconds: 1000,
		Note:    "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path

		assert.Equal(t, string(endpoints.BidTimeTracking(bidID)), r.URL.Path)
		// body
		var res rr.CreateTimeTrackingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Seconds, res.Seconds)
		assert.Equal(t, body.Note, res.Note)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.CreateTimeTracking(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_ListEditRequests(t *testing.T) {
	bidID := int64(100)
	opts := rr.ListBidEditRequestsOptions{
		Statuses:          []rr.BidStatus{rr.BidStatusAccepted},
		BidEditRequestIDs: []int64{1, 2},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidEditRequests(bidID)), r.URL.Path)
		// options
		assert.ElementsMatch(t, []string{string(rr.BidStatusAccepted)}, q["statuses[]"])
		assert.ElementsMatch(t, []string{"1", "2"}, q["bid_edit_request_ids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.ListEditRequests(context.Background(), bidID, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_CreateEditRequest(t *testing.T) {
	body := rr.CreateBidEditRequestBody{
		BidID:     int64(1),
		NewAmount: 100,
		Comment:   "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidsEditRequests), r.URL.Path)
		// body
		var res rr.CreateBidEditRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.BidID, res.BidID)
		assert.Equal(t, body.NewAmount, res.NewAmount)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.CreateEditRequest(context.Background(), body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_ActionEditRequests(t *testing.T) {
	bidID := int64(100)
	bidEditRequestID := int64(200)
	body := rr.ActionBidEditRequestBody{
		Action: rr.BidEditRequestActionAccept,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidEditRequest(bidID, bidEditRequestID)), r.URL.Path)
		// body
		var res rr.ActionBidEditRequestBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, body.Action, res.Action)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.ActionEditRequest(context.Background(), bidID, bidEditRequestID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_GetRating(t *testing.T) {
	bidID := int64(100)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidsRatings(bidID)), r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.GetRating(context.Background(), bidID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_ListRatings(t *testing.T) {
	opts := rr.GetByListOfBidsOptions{
		Bids: []int64{1, 2, 3},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		// method
		assert.Equal(t, http.MethodGet, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidRatings), r.URL.Path)
		// options
		assert.ElementsMatch(t, []string{"1", "2", "3"}, q["bids[]"])

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.ListRatings(context.Background(), &opts)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_CreateRating(t *testing.T) {
	bidID := int64(199)
	body := rr.CreateBidRatingBody{
		Rating:  10,
		Comment: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPost, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidsRatings(bidID)), r.URL.Path)
		// body
		var res rr.CreateBidRatingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Rating, res.Rating)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.CreateRating(context.Background(), bidID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestBids_UpdateRating(t *testing.T) {
	bidID := int64(199)
	bidRatingID := int64(9)
	body := rr.UpdateBidRatingBody{
		Rating:  10,
		Comment: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// method
		assert.Equal(t, http.MethodPut, r.Method)
		// path
		assert.Equal(t, string(endpoints.BidRating(bidID, bidRatingID)), r.URL.Path)
		// body
		var res rr.UpdateBidRatingBody
		err := json.NewDecoder(r.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, body.Rating, res.Rating)
		assert.Equal(t, body.Comment, res.Comment)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))

	}))
	defer ts.Close()

	c := NewClient("token", WithHttpClient(ts.Client()))
	c.SetBaseUrl(ts.URL)

	res, _, err := c.Resources.Bids.UpdateRating(context.Background(), bidID, bidRatingID, body)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

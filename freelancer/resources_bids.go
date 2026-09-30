package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Bids
// --------------------------------------

// Returns a list of bids that match the specified criteria.
// It maps to the `GET` `/projects/0.1/bids` endpoint
func (r *Bids) List(
	ctx context.Context,
	opts *rr.ListBidsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Bids,
		opts,
		nil,
	)
}

// Returns a list of bids that match the specified criteria.
// it maps to the `GET` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Get(
	ctx context.Context,
	bidID int64,
	opts *rr.GetBidOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Bid(bidID),
		opts,
		nil,
	)
}

// Creates a bid on a project. Accepts a JSON object in the style described in the Bid struct (with enums as strings, and objects as dictionaries).
// It maps to the `POST` `/projects/0.1/bids` endpoint
func (r *Bids) Create(
	ctx context.Context,
	b rr.CreateBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Bids,
		nil,
		b,
	)
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Action(
	ctx context.Context,
	bidID int64,
	b rr.ActionBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Bid(bidID),
		nil,
		b,
	)
}

// Updates an existing bid on a project. An existing bids information (description,amount,milestone_percentage) can be updated by sending a JSON encoded Bid struct.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Update(
	ctx context.Context,
	bidID int64,
	b rr.UpdateBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Bid(bidID),
		nil,
		b,
	)
}

// Returns a list of aggregate time tracking data for a bid.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (r *Bids) GetTimeTracking(
	ctx context.Context,
	bidID int64,
	opts *rr.GetTimeTrackingOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidTimeTracking(bidID),
		opts,
		nil,
	)
}

// Creates a time tracking session for a specific bid.
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (r *Bids) CreateTimeTracking(
	ctx context.Context,
	bidID int64,
	b rr.CreateTimeTrackingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidTimeTracking(bidID),
		nil,
		b,
	)
}

// Return bid edit requests by bid id.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/edit_requests` endpoint
// the original name was Get but it was renamed to List due to returning list of edit requests
func (r *Bids) ListEditRequests(
	ctx context.Context,
	bidID int64,
	opts *rr.ListBidEditRequestsOptions,
) (*rr.ListBidEditRequestsResponse, *ResponseMeta, error) {
	return execute[*rr.ListBidEditRequestsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidEditRequests(bidID),
		opts,
		nil,
	)
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request.
// It maps to the `POST` `/projects/0.1/bids/edit_requests` endpoint
func (r *Bids) CreateEditRequest(
	ctx context.Context,
	b rr.CreateBidEditRequestBody,
) (*rr.CreateBidEditRequestResponse, *ResponseMeta, error) {
	return execute[*rr.CreateBidEditRequestResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidsEditRequests,
		nil,
		b,
	)
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (r *Bids) ActionEditRequest(
	ctx context.Context,
	bidID, bidEditRequestID int64,
	b rr.ActionBidEditRequestBody,
) (*rr.ActionBidEditRequestResponse, *ResponseMeta, error) {
	return execute[*rr.ActionBidEditRequestResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.BidEditRequest(bidID, bidEditRequestID),
		nil,
		b,
	)
}

// Fetch bid rating for a bid
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (r *Bids) GetRating(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidsRatings(bidID),
		nil,
		nil,
	)
}

// Fetch bid ratings for multiple bids
// it maps to the `GET` `/projects/0.1/bid_ratings` endpoint
func (r *Bids) ListRatings(
	ctx context.Context,
	opts *rr.GetByListOfBidsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidRatings,
		opts,
		nil,
	)
}

// Rates a bid (creates a bid rating)
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (r *Bids) CreateRating(
	ctx context.Context,
	bidID int64,
	b rr.CreateBidRatingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidsRatings(bidID),
		nil,
		b,
	)
}

// Updates an existing bid rating
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` endpoint
func (r *Bids) UpdateRating(
	ctx context.Context,
	bidID int64,
	bidRatingID int64,
	b rr.UpdateBidRatingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.BidRating(bidID, bidRatingID),
		nil,
		b,
	)
}

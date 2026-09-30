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

// Bid owner accepts the bid award from the project owner.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Accept(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionAccept,
		},
	)
}

// Bid owner denies the bid award from the project owner.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Deny(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionDeny,
		},
	)
}

// Bid owner retracts the bid from the project before it has been awarded.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Retract(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionRetract,
		},
	)
}

// Bid owner highlights a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Highlight(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionHighlight,
		},
	)
}

// Bid owner sponsors a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Sponsor(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionSponsor,
		},
	)
}

// Project owner awards a bid to the bid owner.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Award(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionAward,
		},
	)
}

// Project owner revoke an awarded bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Revoke(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionRevoke,
		},
	)
}

// Project owner add a bid to the shortlist.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Shortlist(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionShortlist,
		},
	)
}

// Project owner remove a bid from the shortlist.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Unshortlist(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionUnshortlist,
		},
	)
}

// Project owner hide a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Hide(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionHide,
		},
	)
}

// Project owner unhide a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Unhide(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionUnhide,
		},
	)
}

// Project owner send a request for location sharing to the freelancer.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) RequestLocationSharing(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		bidID,
		rr.ActionBidBody{
			Action: rr.BidActionRequestLocationSharing,
		},
	)
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) action(
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
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
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
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidsEditRequests,
		nil,
		b,
	)
}

// Employer accepts the purposed amount and period in the bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (r *Bids) AcceptEditRequest(
	ctx context.Context,
	bidID, editRequestID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.actionEditRequest(
		ctx,
		bidID, editRequestID,
		rr.ActionBidEditRequestBody{
			Action: rr.BidEditRequestActionAccept,
		},
	)
}

// Employer declines the purposed amount and period in the bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (r *Bids) DeclineEditRequest(
	ctx context.Context,
	bidID, editRequestID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.actionEditRequest(
		ctx,
		bidID, editRequestID,
		rr.ActionBidEditRequestBody{
			Action: rr.BidEditRequestActionDecline,
		},
	)
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (r *Bids) actionEditRequest(
	ctx context.Context,
	bidID, bidEditRequestID int64,
	b rr.ActionBidEditRequestBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
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

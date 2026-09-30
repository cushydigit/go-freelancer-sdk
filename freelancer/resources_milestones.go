package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Milestones
// --------------------------------------

// Returns a list of milestones. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/milestones` endpoint
func (r *Milestones) List(
	ctx context.Context,
	opts *rr.ListMilestonesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Milestones,
		opts,
		nil,
	)
}

// Returns information about a specific milestone.
// It maps to the `GET` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Get(
	ctx context.Context,
	milestoneID int64,
	opts *rr.GetMilestoneOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Milestone(milestoneID),
		opts,
		nil,
	)
}

// Creates a milestone.
// It maps to the `POST` `/projects/0.1/milestones` endpoint
func (r *Milestones) Create(
	ctx context.Context,
	b rr.CreateMilestoneBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Milestones,
		nil,
		b,
	)
}

// Actions to be performed on a milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Action(
	ctx context.Context,
	milestoneID int64,
	b rr.ActionMilestoneBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Milestone(milestoneID),
		nil,
		b,
	)
}

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/milestone_requests` endpoint
func (r *Milestones) ListRequests(
	ctx context.Context,
	opts *rr.ListMilestoneRequestsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.MilestoneRequests,
		opts,
		nil,
	)
}

// Returns information about a specific milestone request.
// It maps to the `GET` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) GetRequest(
	ctx context.Context,
	milestoneRequestID int64,
	opts *rr.GetMilestoneRequestOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.MilestoneRequest(milestoneRequestID),
		opts,
		nil,
	)
}

// Creates a milestone request from a given JSON object.
// It maps to the `POST` `/projects/0.1/milestone_requests` endpoint
func (r *Milestones) CreateRequest(
	ctx context.Context,
	b rr.CreateMilestoneRequestBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.MilestoneRequests,
		nil,
		b,
	)
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) ActionRequest(
	ctx context.Context,
	milestoneRequestID int64,
	b rr.ActionMilestoneRequestBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.MilestoneRequest(milestoneRequestID),
		nil,
		b,
	)
}

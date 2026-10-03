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

// Project owner release a milestone payment.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Release(
	ctx context.Context,
	milestoneID int64,
	amount int,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action: rr.MilestoneActionRelease,
			Amount: amount,
		},
	)
}

// Project owner update the description of a milestone
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Update(
	ctx context.Context,
	milestoneID int64,
	description string,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action:      rr.MilestoneActionUpdate,
			OtherReason: description,
		},
	)
}

// Project owner request the bid owner to cancel the milestone
// if project owner choose other for reason so the reasonText should be contain description.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) RequestCancel(
	ctx context.Context,
	milestoneID int64,
	reason rr.MilestoneRequestCancelReason,
	reasonText string,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action:     rr.MilestoneActionRequestCancel,
			Reason:     reason,
			ReasonText: reasonText,
		},
	)
}

// Freelancer request the release of a milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) RequestRelease(
	ctx context.Context,
	milestoneID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action: rr.MilestoneActionRequestRelease,
		},
	)
}

// Freelancer cancel a milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Cancel(
	ctx context.Context,
	milestoneID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action: rr.MilestoneActionCancel,
		},
	)
}

// Freelancer reject the project owner's request to cancel the milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) RejectCancel(
	ctx context.Context,
	milestoneID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		milestoneID,
		rr.ActionMilestoneBody{
			Action: rr.MilestoneActionRejectCancel,
		},
	)
}

// Actions to be performed on a milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) action(
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

// Project owner accept a milestone from the milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) AcceptRequest(
	ctx context.Context,
	milestoneRequestID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.actionRequest(
		ctx,
		milestoneRequestID,
		rr.ActionMilestoneRequestBody{
			Action: rr.MilestoneActionAcceptRequest,
		},
	)
}

// Project owner rejects the milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) RejectRequest(
	ctx context.Context,
	milestoneRequestID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.actionRequest(
		ctx,
		milestoneRequestID,
		rr.ActionMilestoneRequestBody{
			Action: rr.MilestoneActionRejectRequest,
		},
	)
}

// Bid owner deletes the milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) DeleteRequest(
	ctx context.Context,
	milestoneRequestID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.actionRequest(
		ctx,
		milestoneRequestID,
		rr.ActionMilestoneRequestBody{
			Action: rr.MilestoneActionDeleteRequest,
		},
	)
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) actionRequest(
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

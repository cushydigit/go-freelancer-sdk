package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-ExpertGuarantees
// --------------------------------------

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (r *ExpertGuarantees) List(
	ctx context.Context,
	opts *rr.ListExpertGuaranteesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ExpertGuarantees,
		opts,
		nil,
	)
}

// Project owner release.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (r *ExpertGuarantees) Release(
	ctx context.Context,
	expertGuaranteeID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		expertGuaranteeID,
		rr.ActionExpertGuaranteesBody{
			Action: rr.ExpertGuaranteesActionRelease,
		},
	)
}

// Guarantee creator request release.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (r *ExpertGuarantees) RequestRelease(
	ctx context.Context,
	expertGuaranteeID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		expertGuaranteeID,
		rr.ActionExpertGuaranteesBody{
			Action: rr.ExpertGuaranteesActionRequestRelease,
		},
	)
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (r *ExpertGuarantees) action(
	ctx context.Context,
	expertGuaranteeID int64,
	b rr.ActionExpertGuaranteesBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.ExpertGuarantee(expertGuaranteeID),
		nil,
		b,
	)
}

package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Reviews
// --------------------------------------

// Returns a list of project reviews.
// It maps to the `GET` `/projects/0.1/reviews` endpoint
func (r *Reviews) List(
	ctx context.Context,
	opts *rr.ListReviewsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Reviews,
		opts,
		nil,
	)
}

// Post a review for freelancer
// It maps to the `POST` `/projects/0.1/reviews` endpoint
func (r *Reviews) CreateForFreelancer(
	ctx context.Context,
	b rr.CreateReviewForFreelancerBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	b.Role = rr.RoleFreelancer
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Reviews,
		nil,
		b,
	)
}

// Post a review for employer
// It maps to the `POST` `/projects/0.1/reviews` endpoint
func (r *Reviews) CreateForEmployer(
	ctx context.Context,
	b rr.CreateReviewForEmployerBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	b.Role = rr.RoleEmployer
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Reviews,
		nil,
		b,
	)
}

// Feature a review with type of project
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (r *Reviews) FeatureProject(
	ctx context.Context,
	id int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		id,
		rr.ActionReviewBody{
			Action:     rr.ReviewActionFeature,
			ReviewType: rr.ReviewTypeProject,
		},
	)
}

// Unfeature a review with type of project
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (r *Reviews) UnfeatureProject(
	ctx context.Context,
	id int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		id,
		rr.ActionReviewBody{
			Action:     rr.ReviewActionUnfeature,
			ReviewType: rr.ReviewTypeProject,
		},
	)
}

// Feature a review with type of contest
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (r *Reviews) FeatureContest(
	ctx context.Context,
	id int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		id,
		rr.ActionReviewBody{
			Action:     rr.ReviewActionFeature,
			ReviewType: rr.ReviewTypeContest,
		},
	)
}

// Unfeature a review with type of contest
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (r *Reviews) UnfeatureContest(
	ctx context.Context,
	id int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		id,
		rr.ActionReviewBody{
			Action:     rr.ReviewActionUnfeature,
			ReviewType: rr.ReviewTypeContest,
		},
	)
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (r *Reviews) action(
	ctx context.Context,
	reviewID int64,
	b rr.ActionReviewBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Review(reviewID),
		nil,
		b,
	)
}

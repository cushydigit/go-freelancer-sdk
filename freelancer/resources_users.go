package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resources - Users
// --------------------------------------

// Returns a list of users.
// It maps to the `GET` `/users/0.1/users` endpoint.
func (r *Users) List(
	ctx context.Context,
	opts *rr.ListUsersOptions,
) (*rr.ListUsersResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Users,
		opts,
		nil,
	)
}

// Returns information about a specific user.
// It maps to the `GET` `/users/0.1/users/{user_id}` endpoint.
func (r *Users) Get(
	ctx context.Context,
	userID int64,
) (*rr.GetUserResponse, *ResponseMeta, error) {
	return execute[*rr.GetUserResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.User(userID),
		nil,
		nil,
	)
}

// Returns a list of Freelancers. The total_count field is the total number of eligible Freelancers, but these users can be limited by the offset and limit parameters.
// It maps to the `GET` `/users/0.1/users/directory` endpoint.
func (r *Users) SearchFreelancer(
	ctx context.Context,
	opts *rr.SearchFreelancerOptions,
) (*rr.SearchFreelancersResponse, *ResponseMeta, error) {
	return execute[*rr.SearchFreelancersResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Freelancers,
		opts,
		nil,
	)
}

// Gets the reputations for a list of users.
// It maps to the `GET` `/users/0.1/reputations` endpoint.
func (r *Users) ListReputations(
	ctx context.Context,
	opts *rr.ListReputationsOptions,
) (*rr.ListUsersReputationsResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersReputationsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Reputations,
		opts,
		nil,
	)
}

// Returns a list of enterprises.
// It maps to the `GET` `/users/0.1/enterprises` endpoint.
func (r *Users) ListEnterprises(
	ctx context.Context,
	opts *rr.ListEnterprisesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Enterprises,
		opts,
		nil,
	)
}

// The service gets the portfolios for a list of users
// Returns a list of portfolios of users. Number of portfolios of all users can be limited by the offset and limit parameters.
func (r *Users) ListPortfolios(
	ctx context.Context,
	opts *rr.ListPortfoliosOptions,
) (*rr.ListUsersPortfoliosResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersPortfoliosResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Portfolios,
		opts,
		nil,
	)
}

// Creates a user violation report.
// It maps to the `POST` `/users/0.1/violation_reports` endpoint.
func (r *Users) CreateViolationReport(
	ctx context.Context,
	b rr.CreateViolationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ViolationReports,
		nil,
		b,
	)
}

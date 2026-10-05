package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resources - Authenticated (SELF)
// --------------------------------------

// Returns information about the current user.
// It maps to the `GET` `/users/0.1/self` endpoint.
func (r *Self) Get(
	ctx context.Context,
	opts *rr.GetSelfInfoOptions,
) (*rr.UserResponse, *ResponseMeta, error) {
	return execute[*rr.UserResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Self,
		opts,
		nil,
	)
}

// Returns a list of user’s recent logged in devices.
// It maps to the `GET` `/users/0.1/self/devices` endpoint.
func (r *Self) ListDevices(
	ctx context.Context,
) (*rr.ListSelfDevicesResponse, *ResponseMeta, error) {
	return execute[*rr.ListSelfDevicesResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Devices,
		nil,
		nil,
	)
}

// Add a list of jobs to the job list of a current user.
// It maps to the `POST` `/users/0.1/self/jobs` endpoint.
func (r *Self) AddJobs(
	ctx context.Context,
	b rr.AddJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Sets a list of jobs to the job list of the current user.
// It maps to the `PUT` `/users/0.1/self/jobs` endpoint.
func (r *Self) UpdateJobs(
	ctx context.Context,
	b rr.SetJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Removes a list of jobs from the job list of the current user.
// It maps to the `DELETE` `/users/0.1/self/jobs` endpoint.
func (r *Self) DeleteJobs(
	ctx context.Context,
	b rr.DeleteJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodDelete,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Returns a list of pools belonging to the current user.
// It maps to the `GET` `/users/0.1/pools` endpoint.
func (r *Self) ListPools(
	ctx context.Context,
	opts *rr.ListPoolsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Pools,
		opts,
		nil,
	)
}

// Returns the logged in user’s projects/contests they either created or participated in (by bidding or submitting an entry).
// it maps to the `GET` `/projects/0.1/self` endpoint
func (r *Self) ListProjects(
	ctx context.Context,
	opts *rr.ListSelfProjectsOptions,
) (*rr.ListProjectsResponse, *ResponseMeta, error) {
	return execute[*rr.ListProjectsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsSelf,
		opts,
		nil,
	)
}

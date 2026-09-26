package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// USERS
// --------------------------------------

// Returns a list of users.
// It maps to the `GET` `/users/0.1/users` endpoint.
func (s *UsersService) List(
	ctx context.Context,
	opts *rr.ListUsersOptions,
) (*rr.ListUsersResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Users,
		opts,
		nil,
	)
}

// Returns information about a specific user.
// It maps to the `GET` `/users/0.1/users/{user_id}` endpoint.
func (s *UsersService) Get(
	ctx context.Context,
	userID int64,
) (*rr.GetUserResponse, *ResponseMeta, error) {
	return execute[*rr.GetUserResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.User(userID),
		nil,
		nil,
	)
}

// Returns a list of Freelancers. The total_count field is the total number of eligible Freelancers, but these users can be limited by the offset and limit parameters.
// It maps to the `GET` `/users/0.1/users/directory` endpoint.
func (s *UsersService) SearchFreelancer(
	ctx context.Context,
	opts *rr.SearchFreelancerOptions,
) (*rr.SearchFreelancersResponse, *ResponseMeta, error) {
	return execute[*rr.SearchFreelancersResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Freelancers,
		opts,
		nil,
	)
}

// Gets the reputations for a list of users.
// It maps to the `GET` `/users/0.1/reputations` endpoint.
func (s *UsersService) ListReputations(
	ctx context.Context,
	opts *rr.ListReputationsOptions,
) (*rr.ListUsersReputationsResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersReputationsResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Reputations,
		opts,
		nil,
	)
}

// Returns a list of enterprises.
// It maps to the `GET` `/users/0.1/enterprises` endpoint.
func (s *UsersService) ListEnterprises(
	ctx context.Context,
	opts *rr.ListEnterprisesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Enterprises,
		opts,
		nil,
	)
}

// The service gets the portfolios for a list of users
// Returns a list of portfolios of users. Number of portfolios of all users can be limited by the offset and limit parameters.
func (s *UsersService) ListPortfolios(
	ctx context.Context,
	opts *rr.ListPortfoliosOptions,
) (*rr.ListUsersPortfoliosResponse, *ResponseMeta, error) {
	return execute[*rr.ListUsersPortfoliosResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Portfolios,
		opts,
		nil,
	)
}

// Creates a user violation report.
// It maps to the `POST` `/users/0.1/violation_reports` endpoint.
func (s *UsersService) CreateViolationReport(
	ctx context.Context,
	b rr.CreateViolationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodPost,
		endpoints.ViolationReports,
		nil,
		b,
	)
}

// --------------------------------------
// USERS - Authenticated (SELF)
// --------------------------------------

// Returns information about the current user.
// It maps to the `GET` `/users/0.1/self` endpoint.
func (s *SelfService) Get(
	ctx context.Context,
	opts *rr.GetSelfInfoOptions,
) (*rr.GetSelfInfoResponse, *ResponseMeta, error) {
	return execute[*rr.GetSelfInfoResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Self,
		opts,
		nil,
	)
}

// Returns a list of user’s recent logged in devices.
// It maps to the `GET` `/users/0.1/self/devices` endpoint.
func (s *SelfService) ListDevices(
	ctx context.Context,
) (*rr.ListSelfDevicesResponse, *ResponseMeta, error) {
	return execute[*rr.ListSelfDevicesResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Devices,
		nil,
		nil,
	)
}

// Add a list of jobs to the job list of a current user.
// It maps to the `POST` `/users/0.1/self/jobs` endpoint.
func (s *SelfService) AddJobs(
	ctx context.Context,
	b rr.AddJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodPost,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Sets a list of jobs to the job list of the current user.
// It maps to the `PUT` `/users/0.1/self/jobs` endpoint.
func (s *SelfService) UpdateJobs(
	ctx context.Context,
	b rr.SetJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodPut,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Removes a list of jobs from the job list of the current user.
// It maps to the `DELETE` `/users/0.1/self/jobs` endpoint.
func (s *SelfService) DeleteJobs(
	ctx context.Context,
	b rr.DeleteJobsBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodDelete,
		endpoints.SelfJobs,
		nil,
		b,
	)
}

// Create a new profile for a user. Returns the created profile
// It maps to the `POST` `/users/0.1/profiles` endpoint.
func (s *SelfService) CreateProfile(
	ctx context.Context,
	b rr.CreateProfileBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodPost,
		endpoints.Profiles,
		nil,
		b,
	)
}

// NOTE: the api does not have solid on this endpoint (the get should not have body)

// Get profile(s)
// It maps to the `GET` `/users/0.1/profiles` endpoint.
func (s *SelfService) GetProfile(
	ctx context.Context,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Profiles,
		nil,
		nil,
	)
}

// Update a profile
// It maps to the `PUT` `/users/0.1/profiles` endpoint.
func (s *SelfService) UpdateProfile(
	ctx context.Context,
	b rr.UpdateProfileBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodPut,
		endpoints.Profiles,
		nil,
		b,
	)
}

// Returns a list of pools belonging to the current user.
// It maps to the `GET` `/users/0.1/pools` endpoint.
func (s *SelfService) ListPools(
	ctx context.Context,
	opts *rr.ListPoolsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		s.client,
		http.MethodGet,
		endpoints.Pools,
		opts,
		nil,
	)
}

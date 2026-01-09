package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/utils"
)

// --------------------------------------
// USERS
// --------------------------------------

// Returns a list of users.
// It maps to the `GET` `/users/0.1/users` endpoint.
func (s *UsersService) List(ctx context.Context, opts *utils.ListUsersOptions) (*ListUsersResponse, error) {
	p := endpoints.Users
	q := query.Values(opts)
	return execute[*ListUsersResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Returns information about a specific user.
// It maps to the `GET` `/users/0.1/users/{user_id}` endpoint.
func (s *UsersService) Get(ctx context.Context, userID int64) (*GetUserResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Users, userID)
	return execute[*GetUserResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Returns a list of Freelancers. The total_count field is the total number of eligible Freelancers, but these users can be limited by the offset and limit parameters.
// It maps to the `GET` `/users/0.1/users/directory` endpoint.
func (s *UsersService) SearchFreelancer(ctx context.Context, opts *utils.SearchFreelancerOptions) (*SearchFreelancersResponse, error) {
	p := endpoints.UsersFreelancers
	q := query.Values(opts)
	return execute[*SearchFreelancersResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// USERS - Authenticated (SELF)
// --------------------------------------

// Returns information about the current user.
// It maps to the `GET` `/users/0.1/self` endpoint.
func (s *SelfService) GetInfo(ctx context.Context, opts *utils.GetSelfInfoOptions) (*GetSelfInfoResponse, error) {
	p := endpoints.UsersSelf
	q := query.Values(opts)
	return execute[*GetSelfInfoResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Returns a list of user’s recent logged in devices.
// It maps to the `GET` `/users/0.1/self/devices` endpoint.
func (s *SelfService) ListDevices(ctx context.Context) (*ListSelfLoginDevicesResponse, error) {
	p := endpoints.UsersSelfDevices
	return execute[*ListSelfLoginDevicesResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Add a list of jobs to the job list of a current user.
// It maps to the `POST` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Add(ctx context.Context, b utils.JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Sets a list of jobs to the job list of the current user.
// It maps to the `PUT` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Set(ctx context.Context, b utils.JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// Removes a list of jobs from the job list of the current user.
// It maps to the `DELETE` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Delete(ctx context.Context, b utils.JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodDelete, p, nil, b)
}

// --------------------------------------
// USERS - PROFILES
// --------------------------------------

// Create a new profile for a user. Returns the created profile
// It maps to the `POST` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Create(ctx context.Context, b utils.CreateProfileBody) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: the api does not have solid on this endpoint (the get should not have body)
// TODO: refined with typed response
// Get profile(s)
// It maps to the `GET` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Get(ctx context.Context) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// TODO: refine with typed response

// Update a profile
// It maps to the `PUT` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Update(ctx context.Context, b utils.UpdateProfileBody) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// --------------------------------------
// USERS - EXTRAS
// --------------------------------------

// Gets the reputations for a list of users.
// It maps to the `GET` `/users/0.1/reputations` endpoint.
func (s *UsersExtrasService) ListReputations(ctx context.Context, opts *utils.ListReputationsOptions) (*ListUsersReputationsResponse, error) {
	p := endpoints.UsersReputations
	q := query.Values(opts)
	return execute[*ListUsersReputationsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of enterprises.
// It maps to the `GET` `/users/0.1/enterprises` endpoint.
func (s *UsersExtrasService) ListEnterprises(ctx context.Context, opts *utils.ListEnterprisesOptions) (*RawResponse, error) {
	p := endpoints.UsersEnterprises
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// The service gets the portfolios for a list of users
// Returns a list of portfolios of users. Number of portfolios of all users can be limited by the offset and limit parameters.
func (s *UsersExtrasService) ListPortfolios(ctx context.Context, opts *utils.ListPortfoliosOptions) (*ListUsersPortfoliosResponse, error) {
	p := endpoints.UsersPortfolios
	q := query.Values(opts)
	return execute[*ListUsersPortfoliosResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Creates a user violation report.
// It maps to the `POST` `/users/0.1/violation_reports` endpoint.
func (s *UsersExtrasService) CreateViolation(ctx context.Context, b utils.CreateViolationBody) (*RawResponse, error) {
	p := endpoints.UsersViolationReports
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: refine with typed response

// Returns a list of pools belonging to the current user.
// It maps to the `GET` `/users/0.1/pools` endpoint.
func (s *UsersExtrasService) ListPools(ctx context.Context, opts *utils.ListPoolsOptions) (*RawResponse, error) {
	p := endpoints.UsersPools
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type ListReputationsOptions struct {
	Users        []int64   `url:"users"`
	Jobs         []int64   `url:"jobs"`
	Role         *RoleType `url:"role"`
	JobHistory   *bool     `url:"job_history"`
	ProjectStats *bool     `url:"project_stats"`
	RehireRates  *bool     `url:"rehire_rates"`
}

// Gets the reputations for a list of users.
// It maps to the `GET` `/users/0.1/reputations` endpoint.
func (s *UsersExtrasService) ListReputations(ctx context.Context, opts *ListReputationsOptions) (*ListUsersReputationsResponse, error) {
	p := endpoints.UsersReputations
	q := query.Values(opts)
	return execute[*ListUsersReputationsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListEnterprisesOptions struct {
	Enterprises   []int64  `url:"enterprises[]"`
	InternalNames []string `url:"internal_names[]"`
	Names         []string `url:"names[]"`
	SeoUrls       []string `url:"seo_urls[]"`
	UserID        *int64   `url:"user_id"`
	IgnoreTest    *bool    `url:"ignore_test"`
	Limit         *int     `url:"limit"`
	Offset        *int     `url:"offset"`
}

// TODO: refine with typed response

// Returns a list of enterprises.
// It maps to the `GET` `/users/0.1/enterprises` endpoint.
func (s *UsersExtrasService) ListEnterprises(ctx context.Context, opts *ListEnterprisesOptions) (*RawResponse, error) {
	p := endpoints.UsersEnterprises
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListPortfoliosOptions struct {
	Users  []int64 `url:"users[]"`
	Limit  *int    `url:"limit"`
	Offset *int    `url:"offset"`
}

// The service gets the portfolios for a list of users
// Returns a list of portfolios of users. Number of portfolios of all users can be limited by the offset and limit parameters.
func (s *UsersExtrasService) ListPortfolios(ctx context.Context, opts *ListPortfoliosOptions) (*ListUsersPortfoliosResponse, error) {
	p := endpoints.UsersPortfolios
	q := query.Values(opts)
	return execute[*ListUsersPortfoliosResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type CreateViolationBody struct {
	ContextID        int                       `json:"context_id"`
	ContextType      ViolationContext          `json:"context_type"`
	ViolatorUserID   int64                     `json:"violator_user_id"`
	Reason           ViolationReason           `json:"reason"`
	AdditionalReason ViolationAdditionalReason `json:"additional_reason,omitempty"`
	Comments         string                    `json:"comments,omitempty"`
	Url              string                    `json:"url,omitempty"`
}

// Creates a user violation report.
// It maps to the `POST` `/users/0.1/violation_reports` endpoint.
func (s *UsersExtrasService) CreateViolation(ctx context.Context, b CreateViolationBody) (*RawResponse, error) {
	p := endpoints.UsersViolationReports
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ListPoolsOptions struct {
	Pools           []int64  `url:"pools[]"`
	Names           []string `url:"internal_names[]"`
	SeoUrls         []string `url:"seo_urls[]"`
	IgnoreTest      *bool    `url:"ignore_test"`
	IsTalentNetwork *bool    `url:"is_talent_network"`
	Limit           *int     `url:"limit"`
	Offset          *int     `url:"offset"`
}

// TODO: refine with typed response

// Returns a list of pools belonging to the current user.
// It maps to the `GET` `/users/0.1/pools` endpoint.
func (s *UsersExtrasService) ListPools(ctx context.Context, opts *ListPoolsOptions) (*RawResponse, error) {
	p := endpoints.UsersPools
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

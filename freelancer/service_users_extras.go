package freelancer

import (
	"context"
	"net/http"
	"net/url"
)

type ListReputationsOptions struct {
	Users        []int64
	Jobs         []int64
	Role         *RoleType
	JobHistory   *bool
	ProjectStats *bool
	RehireRates  *bool
}

// Gets the reputations for a list of users.
// It maps to the `GET` `/users/0.1/reputations` endpoint.
func (s *UsersExtrasService) ListReputations(ctx context.Context, opts *ListReputationsOptions) (*ListUsersReputationsResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Users {
			addInt64(query, "users[]", &val)
		}
		for _, val := range opts.Jobs {
			addInt64(query, "jobs[]", &val)
		}
		addEnum(query, "role", opts.Role)
		addBool(query, "job_history", opts.JobHistory)
		addBool(query, "project_stats", opts.ProjectStats)
		addBool(query, "rehire_rates", opts.RehireRates)
	}
	return execute[*ListUsersReputationsResponse](ctx, s.client, http.MethodGet, string(UsersReputations), query, nil)
}

type ListEnterprisesOptions struct {
	Enterprises   []int64
	InternalNames []string
	Names         []string
	SeoUrls       []string
	UserID        *int64
	IgnoreTest    *bool
	Limit         *int
	Offset        *int
}

// TODO: refine with typed response

// Returns a list of enterprises.
// It maps to the `GET` `/users/0.1/enterprises` endpoint.
func (s *UsersExtrasService) ListEnterprises(ctx context.Context, opts *ListEnterprisesOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Enterprises {
			addInt64(query, "enterprises[]", &val)
		}
		for _, val := range opts.InternalNames {
			addString(query, "internal_names[]", &val)
		}
		for _, val := range opts.Names {
			addString(query, "names[]", &val)
		}
		for _, val := range opts.SeoUrls {
			addString(query, "seo_urls[]", &val)
		}
		addInt64(query, "user_id", opts.UserID)
		addBool(query, "ignore_test", opts.IgnoreTest)
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(UsersEnterprises), query, nil)
}

type ListPortfoliosOptions struct {
	Users  []int64
	Limit  *int
	Offset *int
}

// The service gets the portfolios for a list of users
// Returns a list of portfolios of users. Number of portfolios of all users can be limited by the offset and limit parameters.
func (s *UsersExtrasService) ListPortfolios(ctx context.Context, opts *ListPortfoliosOptions) (*ListUsersPortfoliosResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Users {
			addInt64(query, "users[]", &val)
		}
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
	}
	return execute[*ListUsersPortfoliosResponse](ctx, s.client, http.MethodGet, string(UsersPortfolios), query, nil)
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
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(UsersViolationReports), nil, b)
}

type ListPoolsOptions struct {
	Pools           []int
	Names           []string
	SeoUrls         []string
	IgnoreTest      *bool
	IsTalentNetwork *bool
	Limit           *int
	Offset          *int
}

// TODO: refine with typed response

// Returns a list of pools belonging to the current user.
// It maps to the `GET` `/users/0.1/pools` endpoint.
func (s *UsersExtrasService) ListPools(ctx context.Context, opts *ListPoolsOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Pools {
			addInt(query, "pools[]", &val)
		}
		for _, val := range opts.Names {
			addString(query, "internal_names[]", &val)
		}
		for _, val := range opts.SeoUrls {
			addString(query, "seo_urls[]", &val)
		}
		addBool(query, "ignore_test", opts.IgnoreTest)
		addBool(query, "is_talent_network", opts.IsTalentNetwork)
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(UsersPools), query, nil)
}

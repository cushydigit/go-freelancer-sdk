package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type listReputations struct {
	client       *Client
	users        []int64
	jobs         []int32
	role         *RoleType
	jobHistory   *bool
	projectStats *bool
	rehireRates  *bool
}

func (s *listReputations) Do(ctx context.Context) (*ListUsersReputationsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersReputations),
	}

	for _, id := range s.users {
		r.addParam("users[]", id)
	}
	for _, id := range s.jobs {
		r.addParam("jobs[]", id)
	}
	if s.role != nil {
		r.addParam("role", *s.role)
	}
	if s.jobHistory != nil {
		r.addParam("job_history", *s.jobHistory)
	}
	if s.projectStats != nil {
		r.addParam("project_stats", *s.projectStats)
	}
	if s.rehireRates != nil {
		r.addParam("rehire_rates", *s.rehireRates)
	}
	return execute[*ListUsersReputationsResponse](ctx, s.client, r)

}

// SetUsers sets the filter user IDs for the reputation list service.
func (s *listReputations) SetUsers(users []int64) *listReputations {
	s.users = users
	return s
}

// SetJobs limits users to a set of job IDs as a filter.
func (s *listReputations) SetJobs(jobs []int32) *listReputations {
	s.jobs = jobs
	return s
}

// SetRole sets the type of reputation filter (e.g., freelancer or employer).
func (s *listReputations) SetRole(val RoleType) *listReputations {
	s.role = &val
	return s
}

// SetJobHistory sets whether to return job history projection for each user.
func (s *listReputations) SetJobHistory(val bool) *listReputations {
	s.jobHistory = &val
	return s
}

// SetProjectState sets whether to return project stats projection for each user.
func (s *listReputations) SetProjectState(val bool) *listReputations {
	s.projectStats = &val
	return s
}

// SetRehireRates sets whether to return rehire rates projection for a single freelancer.
func (s *listReputations) SetRehireRates(val bool) *listReputations {
	s.rehireRates = &val
	return s
}

type listEnterprises struct {
	client        *Client
	enterprises   []int
	internalNames []string
	names         []string
	seoUrls       []string
	userID        *int64
	ignoreTest    *bool
	limit         *int
	offset        *int
}

// TODO: refine with typed response
func (s *listEnterprises) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersEnterprises),
	}

	for _, val := range s.enterprises {
		r.addParam("enterprises[]", val)
	}

	for _, val := range s.internalNames {
		r.addParam("internal_names[]", val)
	}

	for _, val := range s.names {
		r.addParam("names[]", val)
	}

	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}

	if s.userID != nil {
		r.addParam("user_id", *s.userID)
	}

	if s.ignoreTest != nil {
		r.addParam("ignore_test", *s.ignoreTest)
	}

	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}

	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetEnterprises sets the filter to return enterprises with the specified enterprise IDs.
func (s *listEnterprises) SetEnterprises(vals []int) *listEnterprises {
	s.enterprises = vals
	return s
}

// SetInternalNames sets the filter to return enterprises with the specified internal names.
func (s *listEnterprises) SetInternalNames(vals []string) *listEnterprises {
	s.internalNames = vals
	return s
}

// SetNames sets the filter to return enterprises with the specified names.
func (s *listEnterprises) SetNames(vals []string) *listEnterprises {
	s.names = vals
	return s
}

// SetSeoUrls sets the filter to return enterprises with the specified SEO URLs.
func (s *listEnterprises) SetSeoUrls(vals []string) *listEnterprises {
	s.seoUrls = vals
	return s
}

// SetUserID sets the filter to return enterprises for the specified user ID.
func (s *listEnterprises) SetUserID(val int64) *listEnterprises {
	s.userID = &val
	return s
}

// SetIgnoreTest sets the filter to ignore test enterprises.
func (s *listEnterprises) SetIgnoreTest(val bool) *listEnterprises {
	s.ignoreTest = &val
	return s
}

// SetLimit sets the maximum number of results to return. Default is 100.
func (s *listEnterprises) SetLimit(val int) *listEnterprises {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip, for pagination. Default is 0.
func (s *listEnterprises) SetOffset(val int) *listEnterprises {
	s.offset = &val
	return s
}

type listPortfolios struct {
	client *Client
	users  []int64
	limit  *int
	offset *int
}

func (s *listPortfolios) DO(ctx context.Context) (*ListUsersPortfoliosResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersPortfolios),
	}

	for id := range s.users {
		r.addParam("users[]", id)
	}

	if s.limit != nil {
		r.addParam("limit", s.limit)
	}

	if s.offset != nil {
		r.addParam("offset", s.offset)
	}
	return execute[*ListUsersPortfoliosResponse](ctx, s.client, r)
}

// Returns portfolios of specified user ids
func (s *listPortfolios) SetUsers(users []int64) *listPortfolios {
	s.users = users
	return s
}

// Number of portfolios required for all users.
func (s *listPortfolios) SetLimit(limit int) *listPortfolios {
	s.limit = &limit
	return s
}

// starting offset of portfolios for each user
func (s *listPortfolios) SetOffset(offset int) *listPortfolios {
	s.offset = &offset
	return s
}

type createViolation struct {
	client *Client
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

func (s *createViolation) Do(ctx context.Context, b CreateViolationBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(UsersViolationReports),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type listPools struct {
	client          *Client
	pools           []int
	names           []string
	seoUrls         []string
	ignoreTest      *bool
	isTalentNetwork *bool
	limit           *int
	offset          *int
}

// TODO: refine with typed response
func (s *listPools) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersPools),
	}

	for _, val := range s.pools {
		r.addParam("pools[]", val)
	}
	for _, val := range s.names {
		r.addParam("internal_names[]", val)
	}
	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}
	if s.ignoreTest != nil {
		r.addParam("ignore_test", *s.ignoreTest)
	}
	if s.isTalentNetwork != nil {
		r.addParam("is_talent_network", *s.isTalentNetwork)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetPools sets the list of pool IDs to filter results by.
func (s *listPools) SetPools(vals []int) *listPools {
	s.pools = vals
	return s
}

// SetNames sets the list of pool names to filter results by.
func (s *listPools) SetNames(vals []string) *listPools {
	s.names = vals
	return s
}

// SetSeoUrls sets the list of SEO URLs to filter results by.
func (s *listPools) SetSeoUrls(vals []string) *listPools {
	s.seoUrls = vals
	return s
}

// SetIgnoreTest sets whether to ignore test pools.
func (s *listPools) SetIgnoreTest(val bool) *listPools {
	s.ignoreTest = &val
	return s
}

// SetIsTalentNetwork sets whether to filter pools based on the talent network flag.
func (s *listPools) SetIsTalentNetwork(val bool) *listPools {
	s.isTalentNetwork = &val
	return s
}

// SetLimit sets the maximum number of results to return. Default is 100.
func (s *listPools) SetLimit(val int) *listPools {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination. Default is 0.
func (s *listPools) SetOffset(val int) *listPools {
	s.offset = &val
	return s
}

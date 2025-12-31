package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

// --------------------------------------
// USERS
// --------------------------------------

type ListUsersOptions struct {
	Users                         []int64  `url:"users[]"`
	Usernames                     []string `url:"usernames[]"`
	Avatar                        *bool    `url:"avatar"`
	CountryDetails                *bool    `url:"country_details"`
	ProfileDescription            *bool    `url:"profile_description"`
	DisplayInfo                   *bool    `url:"display_info"`
	Jobs                          *bool    `url:"jobs"`
	BalanceDetails                *bool    `url:"balance_details"`
	QualificationDetails          *bool    `url:"qualification_details"`
	MembershipDetails             *bool    `url:"membership_details"`
	FinancialDetails              *bool    `url:"financial_details"`
	LocationDetails               *bool    `url:"location_details"`
	PortfolioDetails              *bool    `url:"portfolio_details"`
	PreferredDetails              *bool    `url:"preferred_details"`
	BadgeDetails                  *bool    `url:"badge_details"`
	Status                        *bool    `url:"status"`
	Reputation                    *bool    `url:"reputation"`
	EmployerReputation            *bool    `url:"employer_reputation"`
	ReputationExtra               *bool    `url:"reputation_extra"`
	EmployerReputationExtra       *bool    `url:"employer_reputation_extra"`
	CoverImage                    *bool    `url:"cover_image"`
	PastCoverImage                *bool    `url:"past_cover_image"`
	MobileTracking                *bool    `url:"mobile_tracking"`
	BidQualityDetails             *bool    `url:"bid_quality_details"`
	DepositMethods                *bool    `url:"deposit_methods"`
	UserRecommendations           *bool    `url:"user_recommendations"`
	MarketingMobileNumber         *bool    `url:"marketing_mobile_number"`
	SanctionDetails               *bool    `url:"sanction_details"`
	LimitedAccount                *bool    `url:"limited_account"`
	CompletedUserRelevantJobCount *bool    `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool    `url:"equipment_group_details"`
	JobRanks                      *bool    `url:"job_ranks"`
	JobSeoDetails                 *bool    `url:"job_seo_details"`
	RisingStar                    *bool    `url:"rising_star"`
	ShareholderDetails            *bool    `url:"shareholder_details"`
	StaffDetails                  *bool    `url:"staff_details"`
}

// Returns a list of users.
// It maps to the `GET` `/users/0.1/users` endpoint.
func (s *UsersService) List(ctx context.Context, opts *ListUsersOptions) (*ListUsersResponse, error) {
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

type SearchFreelancerOptions struct {
	Query                         *string  `url:"query"`
	JobsIDs                       []int64  `url:"jobs[]"`
	Skills                        []int64  `url:"skills[]"`
	Countries                     []string `url:"countries[]"`
	HourlyRateMin                 *int     `url:"hourly_rate_min"`
	HourlyRateMax                 *int     `url:"hourly_rate_max"`
	ReviewCountMin                *int     `url:"review_count_min"`
	ReviewCountMax                *int     `url:"review_count_max"`
	OnlineOnly                    *bool    `url:"online_only"`
	LocationLatitude              *float64 `url:"location_latitude"`
	LocationLongitude             *float64 `url:"location_longitude"`
	Insignias                     []int64  `url:"insignias[]"`
	PoolIds                       []int64  `url:"pool_ids[]"`
	Ratings                       *float64 `url:"ratings"`
	SortField                     *int     `url:"sort_field"`
	ReverseSort                   *bool    `url:"reverse_sort"`
	Avatar                        *bool    `url:"avatar"`
	CountryDetails                *bool    `url:"country_details"`
	ProfileDescription            *bool    `url:"profile_description"`
	DisplayInfo                   *bool    `url:"display_info"`
	Jobs                          *bool    `url:"jobs"`
	BalanceDetails                *bool    `url:"balance_details"`
	QualificationDetails          *bool    `url:"qualification_details"`
	MembershipDetails             *bool    `url:"membership_details"`
	FinancialDetails              *bool    `url:"financial_details"`
	LocationDetails               *bool    `url:"location_details"`
	PortfolioDetails              *bool    `url:"portfolio_details"`
	PreferredDetails              *bool    `url:"preferred_details"`
	BadgeDetails                  *bool    `url:"badge_details"`
	Status                        *bool    `url:"status"`
	Reputation                    *bool    `url:"reputation"`
	EmployerReputation            *bool    `url:"employer_reputation"`
	ReputationExtra               *bool    `url:"reputation_extra"`
	EmployerReputationExtra       *bool    `url:"employer_reputation_extra"`
	CoverImage                    *bool    `url:"cover_image"`
	PastCoverImage                *bool    `url:"past_cover_image"`
	MobileTracking                *bool    `url:"mobile_tracking"`
	BidQualityDetails             *bool    `url:"bid_quality_details"`
	DepositMethods                *bool    `url:"deposit_methods"`
	UserRecommendations           *bool    `url:"user_recommendations"`
	MarketingMobileNumber         *bool    `url:"marketing_mobile_number"`
	SanctionDetails               *bool    `url:"sanction_details"`
	LimitedAccount                *bool    `url:"limited_account"`
	CompletedUserRelevantJobCount *bool    `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool    `url:"equipment_group_details"`
	JobRanks                      *bool    `url:"job_ranks"`
	JobSeoDetails                 *bool    `url:"job_seo_details"`
	RisingStar                    *bool    `url:"rising_star"`
	ShareholderDetails            *bool    `url:"shareholder_details"`
	StaffDetails                  *bool    `url:"staff_details"`
	PoolDetails                   *bool    `url:"pool_details"`
	Limit                         *int     `url:"limit"`
	Offset                        *int     `url:"offset"`
	Compact                       *bool    `url:"compact"`
}

// Returns a list of Freelancers. The total_count field is the total number of eligible Freelancers, but these users can be limited by the offset and limit parameters.
// It maps to the `GET` `/users/0.1/users/directory` endpoint.
func (s *UsersService) SearchFreelancer(ctx context.Context, opts *SearchFreelancerOptions) (*SearchFreelancersResponse, error) {
	p := endpoints.UsersFreelancers
	q := query.Values(opts)
	return execute[*SearchFreelancersResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// USERS - Authenticated (SELF)
// --------------------------------------
type GetSelfInfoOptions struct {
	Avatar                        *bool `url:"avatar"`
	CountryDetails                *bool `url:"country_details"`
	ProfileDescription            *bool `url:"profile_description"`
	DisplayInfo                   *bool `url:"display_info"`
	Jobs                          *bool `url:"jobs"`
	BalanceDetails                *bool `url:"balance_details"`
	QualificationDetails          *bool `url:"qualification_details"`
	MembershipDetails             *bool `url:"membership_details"`
	FinancialDetails              *bool `url:"financial_details"`
	LocationDetails               *bool `url:"location_details"`
	PortfolioDetails              *bool `url:"portfolio_details"`
	PreferredDetails              *bool `url:"preferred_details"`
	BadgeDetails                  *bool `url:"badge_details"`
	Status                        *bool `url:"status"`
	Reputation                    *bool `url:"reputation"`
	EmployerReputation            *bool `url:"employer_reputation"`
	ReputationExtra               *bool `url:"reputation_extra"`
	EmployerReputationExtra       *bool `url:"employer_reputation_extra"`
	CoverImage                    *bool `url:"cover_image"`
	PastCoverImages               *bool `url:"past_cover_images"`
	MobileTracking                *bool `url:"mobile_tracking"`
	BidQualityDetails             *bool `url:"bid_quality_details"`
	DepositMethods                *bool `url:"deposit_methods"`
	UserRecommendations           *bool `url:"user_recommendations"`
	MarketingMobileNumber         *bool `url:"marketing_mobile_number"`
	SanctionDetails               *bool `url:"sanction_details"`
	LimitedAccount                *bool `url:"limited_account"`
	CompletedUserRelevantJobCount *bool `url:"completed_user_relevant_job_count"`
	EquipmentGroupDetails         *bool `url:"equipment_group_details"`
	JobRanks                      *bool `url:"job_ranks"`
	JobSeoDetails                 *bool `url:"job_seo_details"`
	RisingStar                    *bool `url:"rising_star"`
	ShareholderDetails            *bool `url:"shareholder_details"`
	StaffDetails                  *bool `url:"staff_details"`
	Limit                         *int  `url:"limit"`
	Offset                        *int  `url:"offset"`
	Compact                       *bool `url:"compact"`
}

// Returns information about the current user.
// It maps to the `GET` `/users/0.1/self` endpoint.
func (s *SelfService) GetInfo(ctx context.Context, opts *GetSelfInfoOptions) (*GetSelfInfoResponse, error) {
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

type JobsBody struct {
	Jobs []int64 `json:"jobs[]"`
}

// Add a list of jobs to the job list of a current user.
// It maps to the `POST` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Add(ctx context.Context, b JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Sets a list of jobs to the job list of the current user.
// It maps to the `PUT` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Set(ctx context.Context, b JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// Removes a list of jobs from the job list of the current user.
// It maps to the `DELETE` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Delete(ctx context.Context, b JobsBody) (*RawResponse, error) {
	p := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodDelete, p, nil, b)
}

// --------------------------------------
// USERS - PROFILES
// --------------------------------------

type CreateProfileBody struct {
	Tagline     string `json:"tagline"`
	HourlyRate  int    `json:"hourly_Rate"`
	Description string `json:"description"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

// Create a new profile for a user. Returns the created profile
// It maps to the `POST` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Create(ctx context.Context, b CreateProfileBody) (*RawResponse, error) {
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

type UpdateProfileBody struct {
	ProfileID   int    `json:"profile_id"`
	Tagline     string `json:"tagline,omitempty"`
	HourlyRate  int    `json:"hourly_Rate,omitempty"`
	Description string `json:"description,omitempty"`
	ProfileName string `json:"profile_name,omitempty"`
	SkillIDs    []int  `json:"skill_ids,omitempty"`
}

// TODO: refine with typed response

// Update a profile
// It maps to the `PUT` `/users/0.1/profiles` endpoint.
func (s *ProfilesService) Update(ctx context.Context, b UpdateProfileBody) (*RawResponse, error) {
	p := endpoints.UsersProfiles
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// --------------------------------------
// USERS - EXTRAS
// --------------------------------------

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

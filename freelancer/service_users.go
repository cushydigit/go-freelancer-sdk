package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

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

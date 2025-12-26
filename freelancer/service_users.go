package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ListUsersOptions struct {
	Users                         []int64
	Usernames                     []string
	Avatar                        *bool
	CountryDetails                *bool
	ProfileDescription            *bool
	DisplayInfo                   *bool
	Jobs                          *bool
	BalanceDetails                *bool
	QualificationDetails          *bool
	MembershipDetails             *bool
	FinancialDetails              *bool
	LocationDetails               *bool
	PortfolioDetails              *bool
	PreferredDetails              *bool
	BadgeDetails                  *bool
	Status                        *bool
	Reputation                    *bool
	EmployerReputation            *bool
	ReputationExtra               *bool
	EmployerReputationExtra       *bool
	CoverImage                    *bool
	PastCoverImage                *bool
	MobileTracking                *bool
	BidQualityDetails             *bool
	DepositMethods                *bool
	UserRecommendations           *bool
	MarketingMobileNumber         *bool
	SanctionDetails               *bool
	LimitedAccount                *bool
	CompletedUserRelevantJobCount *bool
	EquipmentGroupDetails         *bool
	JobRanks                      *bool
	JobSeoDetails                 *bool
	RisingStar                    *bool
	ShareholderDetails            *bool
	StaffDetails                  *bool
}

// Returns a list of users.
// It maps to the `GET` `/users/0.1/users` endpoint.
func (s *UsersService) List(ctx context.Context, opts *ListUsersOptions) (*ListUsersResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Users {
			addInt64(query, "users[]", &val)
		}
		for _, val := range opts.Usernames {
			addString(query, "usernames[]", &val)
		}
		addBool(query, "avatar", opts.Avatar)
		addBool(query, "country_details", opts.CountryDetails)
		addBool(query, "profile_description", opts.ProfileDescription)
		addBool(query, "display_info", opts.DisplayInfo)
		addBool(query, "jobs", opts.Jobs)
		addBool(query, "balance_details", opts.BalanceDetails)
		addBool(query, "qualification_details", opts.QualificationDetails)
		addBool(query, "membership_details", opts.MembershipDetails)
		addBool(query, "financial_details", opts.FinancialDetails)
		addBool(query, "location_details", opts.LocationDetails)
		addBool(query, "portfolio_details", opts.PortfolioDetails)
		addBool(query, "preferred_details", opts.PreferredDetails)
		addBool(query, "badge_details", opts.BadgeDetails)
		addBool(query, "status", opts.Status)
		addBool(query, "reputation", opts.Reputation)
		addBool(query, "employer_reputation", opts.EmployerReputation)
		addBool(query, "reputation_extra", opts.ReputationExtra)
		addBool(query, "employer_reputation_extra", opts.EmployerReputationExtra)
		addBool(query, "cover_image", opts.CoverImage)
		addBool(query, "past_cover_image", opts.PastCoverImage)
		addBool(query, "mobile_tracking", opts.MobileTracking)
		addBool(query, "bid_quality_details", opts.BidQualityDetails)
		addBool(query, "deposit_methods", opts.DepositMethods)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "completed_user_relevant_job_count", opts.CompletedUserRelevantJobCount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)
		addBool(query, "job_ranks", opts.JobRanks)
		addBool(query, "job_seo_details", opts.JobSeoDetails)
		addBool(query, "rising_star", opts.RisingStar)
		addBool(query, "shareholder_details", opts.ShareholderDetails)
		addBool(query, "staff_details", opts.StaffDetails)
	}
	return execute[*ListUsersResponse](ctx, s.client, http.MethodGet, string(UsersUsers), query, nil)
}

// Returns information about a specific user.
// It maps to the `GET` `/users/0.1/users/{user_id}` endpoint.
func (s *UsersService) Get(ctx context.Context, userID int64) (*GetUserResponse, error) {
	path := fmt.Sprintf("%s/%d", string(UsersUsers), userID)
	return execute[*GetUserResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

type SearchFreelancerOptions struct {
	Query                         *string
	JobsIDs                       []int64 // renamed from jobs[] to JobsIDs
	Skills                        []int64
	Countries                     []string
	HourlyRateMin                 *int
	HourlyRateMax                 *int
	ReviewCountMin                *int
	ReviewCountMax                *int
	OnlineOnly                    *bool
	LocationLatitude              *float64
	LocationLongitude             *float64
	Insignias                     []int64
	PoolIds                       []int64
	Ratings                       *float64
	SortField                     *int
	ReverseSort                   *bool
	Avatar                        *bool
	CountryDetails                *bool
	ProfileDescription            *bool
	DisplayInfo                   *bool
	Jobs                          *bool
	BalanceDetails                *bool
	QualificationDetails          *bool
	MembershipDetails             *bool
	FinancialDetails              *bool
	LocationDetails               *bool
	PortfolioDetails              *bool
	PreferredDetails              *bool
	BadgeDetails                  *bool
	Status                        *bool
	Reputation                    *bool
	EmployerReputation            *bool
	ReputationExtra               *bool
	EmployerReputationExtra       *bool
	CoverImage                    *bool
	PastCoverImage                *bool
	MobileTracking                *bool
	BidQualityDetails             *bool
	DepositMethods                *bool
	UserRecommendations           *bool
	MarketingMobileNumber         *bool
	SanctionDetails               *bool
	LimitedAccount                *bool
	CompletedUserRelevantJobCount *bool
	EquipmentGroupDetails         *bool
	JobRanks                      *bool
	JobSeoDetails                 *bool
	RisingStar                    *bool
	ShareholderDetails            *bool
	StaffDetails                  *bool
	PoolDetails                   *bool
	Limit                         *int
	Offset                        *int
	Compact                       *bool
}

// Returns a list of Freelancers. The total_count field is the total number of eligible Freelancers, but these users can be limited by the offset and limit parameters.
// It maps to the `GET` `/users/0.1/users/directory` endpoint.
func (s *UsersService) SearchFreelancer(ctx context.Context, opts *SearchFreelancerOptions) (*SearchFreelancersResponse, error) {
	query := url.Values{}
	if opts != nil {
		addString(query, "query", opts.Query)
		for _, val := range opts.JobsIDs {
			addInt64(query, "jobs[]", &val)
		}
		for _, val := range opts.Skills {
			addInt64(query, "skills[]", &val)
		}
		for _, val := range opts.Countries {
			addString(query, "countries[]", &val)
		}
		addInt(query, "hourly_rate_min", opts.HourlyRateMin)
		addInt(query, "hourly_rate_max", opts.HourlyRateMax)
		addInt(query, "review_count_min", opts.ReviewCountMin)
		addInt(query, "review_count_max", opts.ReviewCountMax)
		addBool(query, "online_only", opts.OnlineOnly)
		addFloat(query, "location_latitude", opts.LocationLatitude)
		addFloat(query, "location_longitude", opts.LocationLongitude)
		for _, val := range opts.Insignias {
			addInt64(query, "insignias[]", &val)
		}
		for _, val := range opts.PoolIds {
			addInt64(query, "pool_ids[]", &val)
		}
		addFloat(query, "ratings", opts.Ratings)
		addInt(query, "sort_field", opts.SortField)
		addBool(query, "reverse_sort", opts.ReverseSort)
		addBool(query, "avatar", opts.Avatar)
		addBool(query, "country_details", opts.CountryDetails)
		addBool(query, "profile_description", opts.ProfileDescription)
		addBool(query, "display_info", opts.DisplayInfo)
		addBool(query, "jobs", opts.Jobs)
		addBool(query, "balance_details", opts.BalanceDetails)
		addBool(query, "qualification_details", opts.QualificationDetails)
		addBool(query, "membership_details", opts.MembershipDetails)
		addBool(query, "financial_details", opts.FinancialDetails)
		addBool(query, "location_details", opts.LocationDetails)
		addBool(query, "portfolio_details", opts.PortfolioDetails)
		addBool(query, "preferred_details", opts.PreferredDetails)
		addBool(query, "badge_details", opts.BadgeDetails)
		addBool(query, "status", opts.Status)
		addBool(query, "reputation", opts.Reputation)
		addBool(query, "employer_reputation", opts.EmployerReputation)
		addBool(query, "reputation_extra", opts.ReputationExtra)
		addBool(query, "employer_reputation_extra", opts.EmployerReputationExtra)
		addBool(query, "cover_image", opts.CoverImage)
		addBool(query, "past_cover_image", opts.PastCoverImage)
		addBool(query, "mobile_tracking", opts.MobileTracking)
		addBool(query, "bid_quality_details", opts.BidQualityDetails)
		addBool(query, "deposit_methods", opts.DepositMethods)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "completed_user_relevant_job_count", opts.CompletedUserRelevantJobCount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)
		addBool(query, "job_ranks", opts.JobRanks)
		addBool(query, "job_seo_details", opts.JobSeoDetails)
		addBool(query, "rising_star", opts.RisingStar)
		addBool(query, "shareholder_details", opts.ShareholderDetails)
		addBool(query, "staff_details", opts.StaffDetails)
		addBool(query, "pool_details", opts.PoolDetails)
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)

	}
	return execute[*SearchFreelancersResponse](ctx, s.client, http.MethodGet, string(UsersFreelancers), query, nil)
}

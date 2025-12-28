package freelancer

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

type GetSelfInfoOptions struct {
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
	PastCoverImages               *bool
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
	Limit                         *int
	Offset                        *int
	Compact                       *bool
}

// Returns information about the current user.
// It maps to the `GET` `/users/0.1/self` endpoint.
func (s *SelfService) GetInfo(ctx context.Context, opts *GetSelfInfoOptions) (*GetSelfInfoResponse, error) {
	path := endpoints.UsersSelf
	query := url.Values{}
	if opts != nil {
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
		addBool(query, "past_cover_images", opts.PastCoverImages)
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
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)
	}
	return execute[*GetSelfInfoResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

// Returns a list of user’s recent logged in devices.
// It maps to the `GET` `/users/0.1/self/devices` endpoint.
func (s *SelfService) ListDevices(ctx context.Context) (*ListSelfLoginDevicesResponse, error) {
	path := endpoints.UsersSelfDevices
	return execute[*ListSelfLoginDevicesResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

type JobsBody struct {
	Jobs []int64 `json:"jobs[]"`
}

// Add a list of jobs to the job list of a current user.
// It maps to the `POST` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Add(ctx context.Context, b JobsBody) (*RawResponse, error) {
	path := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

// Sets a list of jobs to the job list of the current user.
// It maps to the `PUT` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Set(ctx context.Context, b JobsBody) (*RawResponse, error) {
	path := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

// Removes a list of jobs from the job list of the current user.
// It maps to the `DELETE` `/users/0.1/self/jobs` endpoint.
func (s *SelfJobsService) Delete(ctx context.Context, b JobsBody) (*RawResponse, error) {
	path := endpoints.UsersSelfJobs
	return execute[*RawResponse](ctx, s.client, http.MethodDelete, path, nil, b)
}

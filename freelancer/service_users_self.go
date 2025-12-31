package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

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

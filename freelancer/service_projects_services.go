package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (s *ServicesService) Order(ctx context.Context, serviceID int, serviceType ServiceType) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/%d/order", endpoints.ProjectsServices, serviceType, serviceID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, nil)
}

type ListServicesOptions struct {
	Services                    []int64             `url:"services[]"`
	Owners                      []int64             `url:"owners[]"`
	Statuses                    []ServiceStatusType `url:"statuses[]"`
	SubStatuses                 []string            `url:"sub_statuses[]"`
	Titles                      []string            `url:"titles[]"`
	SeoUrls                     []string            `url:"seo_urls[]"`
	ExtraDetails                *bool               `url:"extra_details"`
	FileDetails                 *bool               `url:"file_details"`
	JobDetails                  *bool               `url:"job_details"`
	UserDetails                 *bool               `url:"user_details"`
	FullDescription             *bool               `url:"full_description"`
	UserAvatar                  *bool               `url:"user_avatar"`
	UserCountryDetails          *bool               `url:"user_country_details"`
	UserProfileDescription      *bool               `url:"user_profile_description"`
	UserDisplayInfo             *bool               `url:"user_display_info"`
	UserJobs                    *bool               `url:"user_jobs"`
	UserBalanceDetails          *bool               `url:"user_balance_details"`
	UserQualificationDetails    *bool               `url:"user_qualification_details"`
	UserMembershipDetails       *bool               `url:"user_membership_details"`
	UserFinancialDetails        *bool               `url:"user_financial_details"`
	UserLocationDetails         *bool               `url:"user_location_details"`
	UserPortfolioDetails        *bool               `url:"user_portfolio_details"`
	UserPreferredDetails        *bool               `url:"user_preferred_details"`
	UserBadgeDetails            *bool               `url:"user_badge_details"`
	UserStatus                  *bool               `url:"user_status"`
	UserReputation              *bool               `url:"user_reputation"`
	UserEmployerReputation      *bool               `url:"user_employer_reputation"`
	UserReputationExtra         *bool               `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool               `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool               `url:"user_cover_image"`
	UserPastCoverImage          *bool               `url:"user_past_cover_image"`
	UserRecommendations         *bool               `url:"user_recommendations"`
	UserResponsiveness          *bool               `url:"user_responsiveness"`
	CorporateUsers              *bool               `url:"corporate_users"`
	MarketingMobileNumber       *bool               `url:"marketing_mobile_number"`
	SanctionDetails             *bool               `url:"sanction_details"`
	LimitedAccount              *bool               `url:"limited_account"`
	EquipmentGroupDetails       *bool               `url:"equipment_group_details"`
	Limit                       *int                `url:"limit"`
	Offset                      *int                `url:"offset"`
	Compact                     *bool               `url:"compact"`
}

// TODO: refine with typed response

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (s *ServicesService) List(ctx context.Context, opts *ListServicesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsServices
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchActiveServicesOptions struct {
	Query        *string   `url:"query"`
	Sort         *SortType `url:"sort"`
	ReverseSort  *bool     `url:"reverse_sort"`
	ExtraDetails *bool     `url:"extra_details"`
	FileDetails  *bool     `url:"file_details"`
	JobDetails   *bool     `url:"job_details"`
	UserDetails  *bool     `url:"user_details"`
	Offset       *int      `url:"offset"`
	Compact      *bool     `url:"compact"`
}

// TODO: refine with typed response

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (s *ServicesService) SearchActive(ctx context.Context, opts *SearchActiveServicesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsServicesActive
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

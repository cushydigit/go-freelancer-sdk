package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (s *ServicesService) Order(ctx context.Context, serviceID int, serviceType ServiceType) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/%d/order", endpoints.ProjectsServices, serviceType, serviceID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, nil)
}

type ListServicesOptions struct {
	Services                    []int
	Owners                      []int64
	Statuses                    []ServiceStatusType
	SubStatuses                 []string
	Titles                      []string
	SeoUrls                     []string
	ExtraDetails                *bool
	FileDetails                 *bool
	JobDetails                  *bool
	UserDetails                 *bool
	FullDescription             *bool
	UserAvatar                  *bool
	UserCountryDetails          *bool
	UserProfileDescription      *bool
	UserDisplayInfo             *bool
	UserJobs                    *bool
	UserBalanceDetails          *bool
	UserQualificationDetails    *bool
	UserMembershipDetails       *bool
	UserFinancialDetails        *bool
	UserLocationDetails         *bool
	UserPortfolioDetails        *bool
	UserPreferredDetails        *bool
	UserBadgeDetails            *bool
	UserStatus                  *bool
	UserReputation              *bool
	UserEmployerReputation      *bool
	UserReputationExtra         *bool
	UserEmployerReputationExtra *bool
	UserCoverImage              *bool
	UserPastCoverImage          *bool
	UserRecommendations         *bool
	UserResponsiveness          *bool
	CorporateUsers              *bool
	MarketingMobileNumber       *bool
	SanctionDetails             *bool
	LimitedAccount              *bool
	EquipmentGroupDetails       *bool
	Limit                       *int
	Offset                      *int
	Compact                     *bool
}

// TODO: refine with typed response

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (s *ServicesService) List(ctx context.Context, opts *ListServicesOptions) (*RawResponse, error) {
	path := endpoints.ProjectsServices
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.Services {
			addInt(query, "services[]", &id)
		}
		addBool(query, "extra_details", opts.ExtraDetails)
		addBool(query, "file_details", opts.FileDetails)
		addBool(query, "job_details", opts.JobDetails)
		addBool(query, "user_details", opts.UserDetails)
		addBool(query, "full_description", opts.FullDescription)
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)
		addBool(query, "user_balance_details", opts.UserBalanceDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
		addBool(query, "user_past_cover_image", opts.UserPastCoverImage)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "user_responsiveness", opts.UserResponsiveness)
		addBool(query, "corporate_users", opts.CorporateUsers)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type SearchActiveServicesOptions struct {
	Query        *string
	Sort         *SortType
	ReverseSort  *bool
	ExtraDetails *bool
	FileDetails  *bool
	JobDetails   *bool
	UserDetails  *bool
	Offset       *int
	Compact      *bool
}

// TODO: refine with typed response

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (s *ServicesService) SearchActive(ctx context.Context, opts *SearchActiveServicesOptions) (*RawResponse, error) {
	path := endpoints.ProjectsServicesActive
	query := url.Values{}
	if opts != nil {
		addString(query, "query", opts.Query)
		addEnum(query, "sort", opts.Sort)
		addBool(query, "reverse_sort", opts.ReverseSort)
		addBool(query, "extra_details", opts.ExtraDetails)
		addBool(query, "file_details", opts.FileDetails)
		addBool(query, "job_details", opts.JobDetails)
		addBool(query, "user_details", opts.UserDetails)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type orderService struct {
	client      *Client
	serviceID   int
	serviceType ServiceType
}

func (s *orderService) Do(ctx context.Context, serviceID int, serviceType ServiceType) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s/%d/order", string(PROJECTS_SERVICES), string(serviceType), serviceID),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

type listServices struct {
	client                      *Client
	services                    []int
	owners                      []int64
	statuses                    []ServiceStatusType
	subStatuses                 []string
	titles                      []string
	seoUrls                     []string
	extraDetails                *bool
	fileDetails                 *bool
	jobDetails                  *bool
	userDetails                 *bool
	fullDescription             *bool
	userAvatar                  *bool
	userCountryDetails          *bool
	userProfileDescription      *bool
	userDisplayInfo             *bool
	userJobs                    *bool
	userBalanceDetails          *bool
	userQualificationDetails    *bool
	userMembershipDetails       *bool
	userFinancialDetails        *bool
	userLocationDetails         *bool
	userPortfolioDetails        *bool
	userPreferredDetails        *bool
	userBadgeDetails            *bool
	userStatus                  *bool
	userReputation              *bool
	userEmployerReputation      *bool
	userReputationExtra         *bool
	userEmployerReputationExtra *bool
	userCoverImage              *bool
	userPastCoverImage          *bool
	userRecommendations         *bool
	userResponsiveness          *bool
	corporateUsers              *bool
	marketingMobileNumber       *bool
	sanctionDetails             *bool
	limitedAccount              *bool
	equipmentGroupDetails       *bool
	limit                       *int
	offset                      *int
	compact                     *bool
}

func (s *listServices) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SERVICES),
	}

	for _, val := range s.services {
		r.addParam("services[]", val)
	}
	for _, val := range s.owners {
		r.addParam("owners[]", val)
	}
	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	for _, val := range s.subStatuses {
		r.addParam("sub_statuses[]", val)
	}
	for _, val := range s.titles {
		r.addParam("titles[]", val)
	}
	for _, val := range s.seoUrls {
		r.addParam("seo_urls[]", val)
	}
	if s.extraDetails != nil {
		r.addParam("extra_details", *s.extraDetails)
	}
	if s.fileDetails != nil {
		r.addParam("file_details", *s.fileDetails)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.fullDescription != nil {
		r.addParam("full_description", *s.fullDescription)
	}
	if s.userAvatar != nil {
		r.addParam("user_avatar", *s.userAvatar)
	}
	if s.userCountryDetails != nil {
		r.addParam("user_country_details", *s.userCountryDetails)
	}
	if s.userProfileDescription != nil {
		r.addParam("user_profile_description", *s.userProfileDescription)
	}
	if s.userDisplayInfo != nil {
		r.addParam("user_display_info", *s.userDisplayInfo)
	}
	if s.userJobs != nil {
		r.addParam("user_jobs", *s.userJobs)
	}
	if s.userBalanceDetails != nil {
		r.addParam("user_balance_details", *s.userBalanceDetails)
	}
	if s.userQualificationDetails != nil {
		r.addParam("user_qualification_details", *s.userQualificationDetails)
	}
	if s.userMembershipDetails != nil {
		r.addParam("user_membership_details", *s.userMembershipDetails)
	}
	if s.userFinancialDetails != nil {
		r.addParam("user_financial_details", *s.userFinancialDetails)
	}
	if s.userLocationDetails != nil {
		r.addParam("user_location_details", *s.userLocationDetails)
	}
	if s.userPortfolioDetails != nil {
		r.addParam("user_portfolio_details", *s.userPortfolioDetails)
	}
	if s.userPreferredDetails != nil {
		r.addParam("user_preferred_details", *s.userPreferredDetails)
	}
	if s.userBadgeDetails != nil {
		r.addParam("user_badge_details", *s.userBadgeDetails)
	}
	if s.userStatus != nil {
		r.addParam("user_status", *s.userStatus)
	}
	if s.userReputation != nil {
		r.addParam("user_reputation", *s.userReputation)
	}
	if s.userEmployerReputation != nil {
		r.addParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.addParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.addParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.addParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.addParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.addParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.addParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.addParam("corporate_users", *s.corporateUsers)
	}
	if s.marketingMobileNumber != nil {
		r.addParam("marketing_mobile_number", *s.marketingMobileNumber)
	}
	if s.sanctionDetails != nil {
		r.addParam("sanction_details", *s.sanctionDetails)
	}
	if s.limitedAccount != nil {
		r.addParam("limited_account", *s.limitedAccount)
	}
	if s.equipmentGroupDetails != nil {
		r.addParam("equipment_group_details", *s.equipmentGroupDetails)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

// SetServices sets the list of service IDs to filter the response.
func (s *listServices) SetServices(services []int) *listServices {
	s.services = services
	return s
}

// SetOwners sets the list of owner user IDs to filter the services.
func (s *listServices) SetOwners(owners []int64) *listServices {
	s.owners = owners
	return s
}

// SetStatuses sets the service statuses to filter the response (e.g., pending, active, closed).
func (s *listServices) SetStatuses(statuses []ServiceStatusType) *listServices {
	s.statuses = statuses
	return s
}

// SetSubStatuses sets the sub-statuses to filter the response.
func (s *listServices) SetSubStatuses(subStatuses []string) *listServices {
	s.subStatuses = subStatuses
	return s
}

// SetTitles sets the service titles to filter the response.
func (s *listServices) SetTitles(titles []string) *listServices {
	s.titles = titles
	return s
}

// SetSeoUrls sets the SEO URLs to filter the response.
func (s *listServices) SetSeoUrls(seoUrls []string) *listServices {
	s.seoUrls = seoUrls
	return s
}

// SetExtraDetails enables or disables returning extra offerings of services.
func (s *listServices) SetExtraDetails(val bool) *listServices {
	s.extraDetails = &val
	return s
}

// SetFileDetails enables or disables returning files associated with the services.
func (s *listServices) SetFileDetails(val bool) *listServices {
	s.fileDetails = &val
	return s
}

// SetJobDetails enables or disables returning jobs associated with the services.
func (s *listServices) SetJobDetails(val bool) *listServices {
	s.jobDetails = &val
	return s
}

// SetUserDetails enables or disables returning user information for the service owners.
func (s *listServices) SetUserDetails(val bool) *listServices {
	s.userDetails = &val
	return s
}

// SetFullDescription enables or disables returning the full service description.
func (s *listServices) SetFullDescription(val bool) *listServices {
	s.fullDescription = &val
	return s
}

// SetUserAvatar enables or disables returning the avatar of the user(s).
func (s *listServices) SetUserAvatar(val bool) *listServices {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails enables or disables returning the country details of the user(s).
func (s *listServices) SetUserCountryDetails(val bool) *listServices {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription enables or disables returning the profile description of the user(s).
func (s *listServices) SetUserProfileDescription(val bool) *listServices {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo enables or disables returning the display information of the user(s).
func (s *listServices) SetUserDisplayInfo(val bool) *listServices {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs enables or disables returning jobs associated with the user(s).
func (s *listServices) SetUserJobs(val bool) *listServices {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails enables or disables returning the balance information of the user(s).
func (s *listServices) SetUserBalanceDetails(val bool) *listServices {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails enables or disables returning qualification exams completed by the user(s).
func (s *listServices) SetUserQualificationDetails(val bool) *listServices {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails enables or disables returning membership information of the user(s).
func (s *listServices) SetUserMembershipDetails(val bool) *listServices {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails enables or disables returning financial information of the user(s).
func (s *listServices) SetUserFinancialDetails(val bool) *listServices {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails enables or disables returning location information of the user(s).
func (s *listServices) SetUserLocationDetails(val bool) *listServices {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails enables or disables returning portfolio information of the user(s).
func (s *listServices) SetUserPortfolioDetails(val bool) *listServices {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails enables or disables returning preferred information of the user(s).
func (s *listServices) SetUserPreferredDetails(val bool) *listServices {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails enables or disables returning badge details of the user(s).
func (s *listServices) SetUserBadgeDetails(val bool) *listServices {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus enables or disables returning additional status information of the user(s).
func (s *listServices) SetUserStatus(val bool) *listServices {
	s.userStatus = &val
	return s
}

// SetUserReputation enables or disables returning freelancer reputation of the user(s).
func (s *listServices) SetUserReputation(val bool) *listServices {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation enables or disables returning employer reputation of the user(s).
func (s *listServices) SetUserEmployerReputation(val bool) *listServices {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra enables or disables returning full freelancer reputation of the user(s).
func (s *listServices) SetUserReputationExtra(val bool) *listServices {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra enables or disables returning full employer reputation of the user(s).
func (s *listServices) SetUserEmployerReputationExtra(val bool) *listServices {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage enables or disables returning the profile picture of the user(s).
func (s *listServices) SetUserCoverImage(val bool) *listServices {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage enables or disables returning previous profile pictures of the user(s).
func (s *listServices) SetUserPastCoverImage(val bool) *listServices {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations enables or disables returning the number of recommendations of the user(s).
func (s *listServices) SetUserRecommendations(val bool) *listServices {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness enables or disables returning responsiveness scores of the user(s).
func (s *listServices) SetUserResponsiveness(val bool) *listServices {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers enables or disables returning corporate accounts created by the user(s).
func (s *listServices) SetCorporateUsers(val bool) *listServices {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber enables or disables returning the mobile number used by recruiters.
func (s *listServices) SetMarketingMobileNumber(val bool) *listServices {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables or disables returning sanction end timestamps of the user(s).
func (s *listServices) SetSanctionDetails(val bool) *listServices {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables or disables returning the limited account status of the user(s).
func (s *listServices) SetLimitedAccount(val bool) *listServices {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails enables or disables returning equipment group details of the user(s).
func (s *listServices) SetEquipmentGroupDetails(val bool) *listServices {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *listServices) SetLimit(limit int) *listServices {
	s.limit = &limit
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *listServices) SetOffset(offset int) *listServices {
	s.offset = &offset
	return s
}

// SetCompact enables or disables stripping of null and empty values from the response.
func (s *listServices) SetCompact(val bool) *listServices {
	s.compact = &val
	return s
}

type searchActiveServices struct {
	client       *Client
	query        *string
	sort         *SortType
	reverseSort  *bool
	extraDetails *bool
	fileDetails  *bool
	jobDetails   *bool
	userDetails  *bool
	offset       *int
	compact      *bool
}

func (s *searchActiveServices) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SERVICES),
	}

	if s.query != nil {
		r.addParam("query", *s.query)
	}
	if s.sort != nil {
		r.addParam("sort", *s.sort)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.extraDetails != nil {
		r.addParam("extra_details", *s.extraDetails)
	}
	if s.fileDetails != nil {
		r.addParam("file_details", *s.fileDetails)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

// SetQuery sets an optional search query to filter active services by title or description.
func (s *searchActiveServices) SetQuery(val string) *searchActiveServices {
	s.query = &val
	return s
}

// SetSort sets the sorting option for the results.
// Valid choices: "newest", "quickest", "recommended".
func (s *searchActiveServices) SetSort(val SortType) *searchActiveServices {
	s.sort = &val
	return s
}

// SetReverseSort specifies whether to reverse the sorting order of the results.
func (s *searchActiveServices) SetReverseSort(val bool) *searchActiveServices {
	s.reverseSort = &val
	return s
}

// SetExtraDetails specifies whether to include extra service offerings in the response.
func (s *searchActiveServices) SetExtraDetails(val bool) *searchActiveServices {
	s.extraDetails = &val
	return s
}

// SetFileDetails specifies whether to include files associated with the services in the response.
func (s *searchActiveServices) SetFileDetails(val bool) *searchActiveServices {
	s.fileDetails = &val
	return s
}

// SetJobDetails specifies whether to include jobs associated with the services in the response.
func (s *searchActiveServices) SetJobDetails(val bool) *searchActiveServices {
	s.jobDetails = &val
	return s
}

// SetUserDetails specifies whether to include user details of all sellers in the response.
func (s *searchActiveServices) SetUserDetails(val bool) *searchActiveServices {
	s.userDetails = &val
	return s
}

// SetOffset sets the number of results to skip, useful for pagination.
func (s *searchActiveServices) SetOffset(offset int) *searchActiveServices {
	s.offset = &offset
	return s
}

// SetCompact specifies whether to strip all null and empty values from the response.
func (s *searchActiveServices) SetCompact(val bool) *searchActiveServices {
	s.compact = &val
	return s
}

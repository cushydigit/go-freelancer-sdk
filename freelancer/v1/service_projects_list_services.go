package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListProjectsServicesService struct {
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

func (s *ListProjectsServicesService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
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

func (s *ListProjectsServicesService) SetServices(services []int) *ListProjectsServicesService {
	s.services = services
	return s
}

func (s *ListProjectsServicesService) SetOwners(owners []int64) *ListProjectsServicesService {
	s.owners = owners
	return s
}

func (s *ListProjectsServicesService) SetStatuses(statuses []ServiceStatusType) *ListProjectsServicesService {
	s.statuses = statuses
	return s
}

func (s *ListProjectsServicesService) SetSubStatuses(subStatuses []string) *ListProjectsServicesService {
	s.subStatuses = subStatuses
	return s
}

func (s *ListProjectsServicesService) SetTitles(titles []string) *ListProjectsServicesService {
	s.titles = titles
	return s
}

func (s *ListProjectsServicesService) SetSeoUrls(seoUrls []string) *ListProjectsServicesService {
	s.seoUrls = seoUrls
	return s
}

func (s *ListProjectsServicesService) SetExtraDetails(val bool) *ListProjectsServicesService {
	s.extraDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetFileDetails(val bool) *ListProjectsServicesService {
	s.fileDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetJobDetails(val bool) *ListProjectsServicesService {
	s.jobDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserDetails(val bool) *ListProjectsServicesService {
	s.userDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetFullDescription(val bool) *ListProjectsServicesService {
	s.fullDescription = &val
	return s
}

func (s *ListProjectsServicesService) SetUserAvatar(val bool) *ListProjectsServicesService {
	s.userAvatar = &val
	return s
}

func (s *ListProjectsServicesService) SetUserCountryDetails(val bool) *ListProjectsServicesService {
	s.userCountryDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserProfileDescription(val bool) *ListProjectsServicesService {
	s.userProfileDescription = &val
	return s
}

func (s *ListProjectsServicesService) SetUserDisplayInfo(val bool) *ListProjectsServicesService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListProjectsServicesService) SetUserJobs(val bool) *ListProjectsServicesService {
	s.userJobs = &val
	return s
}

func (s *ListProjectsServicesService) SetUserBalanceDetails(val bool) *ListProjectsServicesService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserQualificationDetails(val bool) *ListProjectsServicesService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserMembershipDetails(val bool) *ListProjectsServicesService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserFinancialDetails(val bool) *ListProjectsServicesService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserLocationDetails(val bool) *ListProjectsServicesService {
	s.userLocationDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserPortfolioDetails(val bool) *ListProjectsServicesService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserPreferredDetails(val bool) *ListProjectsServicesService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserBadgeDetails(val bool) *ListProjectsServicesService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetUserStatus(val bool) *ListProjectsServicesService {
	s.userStatus = &val
	return s
}

func (s *ListProjectsServicesService) SetUserReputation(val bool) *ListProjectsServicesService {
	s.userReputation = &val
	return s
}

func (s *ListProjectsServicesService) SetUserEmployerReputation(val bool) *ListProjectsServicesService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListProjectsServicesService) SetUserReputationExtra(val bool) *ListProjectsServicesService {
	s.userReputationExtra = &val
	return s
}

func (s *ListProjectsServicesService) SetUserEmployerReputationExtra(val bool) *ListProjectsServicesService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListProjectsServicesService) SetUserCoverImage(val bool) *ListProjectsServicesService {
	s.userCoverImage = &val
	return s
}

func (s *ListProjectsServicesService) SetUserPastCoverImage(val bool) *ListProjectsServicesService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListProjectsServicesService) SetUserRecommendations(val bool) *ListProjectsServicesService {
	s.userRecommendations = &val
	return s
}

func (s *ListProjectsServicesService) SetUserResponsiveness(val bool) *ListProjectsServicesService {
	s.userResponsiveness = &val
	return s
}

func (s *ListProjectsServicesService) SetCorporateUsers(val bool) *ListProjectsServicesService {
	s.corporateUsers = &val
	return s
}

func (s *ListProjectsServicesService) SetMarketingMobileNumber(val bool) *ListProjectsServicesService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListProjectsServicesService) SetSanctionDetails(val bool) *ListProjectsServicesService {
	s.sanctionDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetLimitedAccount(val bool) *ListProjectsServicesService {
	s.limitedAccount = &val
	return s
}

func (s *ListProjectsServicesService) SetEquipmentGroupDetails(val bool) *ListProjectsServicesService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *ListProjectsServicesService) SetLimit(limit int) *ListProjectsServicesService {
	s.limit = &limit
	return s
}

func (s *ListProjectsServicesService) SetOffset(offset int) *ListProjectsServicesService {
	s.offset = &offset
	return s
}

func (s *ListProjectsServicesService) SetCompact(val bool) *ListProjectsServicesService {
	s.compact = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type servicesListService struct {
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

func (s *servicesListService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
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

func (s *servicesListService) SetServices(services []int) *servicesListService {
	s.services = services
	return s
}

func (s *servicesListService) SetOwners(owners []int64) *servicesListService {
	s.owners = owners
	return s
}

func (s *servicesListService) SetStatuses(statuses []ServiceStatusType) *servicesListService {
	s.statuses = statuses
	return s
}

func (s *servicesListService) SetSubStatuses(subStatuses []string) *servicesListService {
	s.subStatuses = subStatuses
	return s
}

func (s *servicesListService) SetTitles(titles []string) *servicesListService {
	s.titles = titles
	return s
}

func (s *servicesListService) SetSeoUrls(seoUrls []string) *servicesListService {
	s.seoUrls = seoUrls
	return s
}

func (s *servicesListService) SetExtraDetails(val bool) *servicesListService {
	s.extraDetails = &val
	return s
}

func (s *servicesListService) SetFileDetails(val bool) *servicesListService {
	s.fileDetails = &val
	return s
}

func (s *servicesListService) SetJobDetails(val bool) *servicesListService {
	s.jobDetails = &val
	return s
}

func (s *servicesListService) SetUserDetails(val bool) *servicesListService {
	s.userDetails = &val
	return s
}

func (s *servicesListService) SetFullDescription(val bool) *servicesListService {
	s.fullDescription = &val
	return s
}

func (s *servicesListService) SetUserAvatar(val bool) *servicesListService {
	s.userAvatar = &val
	return s
}

func (s *servicesListService) SetUserCountryDetails(val bool) *servicesListService {
	s.userCountryDetails = &val
	return s
}

func (s *servicesListService) SetUserProfileDescription(val bool) *servicesListService {
	s.userProfileDescription = &val
	return s
}

func (s *servicesListService) SetUserDisplayInfo(val bool) *servicesListService {
	s.userDisplayInfo = &val
	return s
}

func (s *servicesListService) SetUserJobs(val bool) *servicesListService {
	s.userJobs = &val
	return s
}

func (s *servicesListService) SetUserBalanceDetails(val bool) *servicesListService {
	s.userBalanceDetails = &val
	return s
}

func (s *servicesListService) SetUserQualificationDetails(val bool) *servicesListService {
	s.userQualificationDetails = &val
	return s
}

func (s *servicesListService) SetUserMembershipDetails(val bool) *servicesListService {
	s.userMembershipDetails = &val
	return s
}

func (s *servicesListService) SetUserFinancialDetails(val bool) *servicesListService {
	s.userFinancialDetails = &val
	return s
}

func (s *servicesListService) SetUserLocationDetails(val bool) *servicesListService {
	s.userLocationDetails = &val
	return s
}

func (s *servicesListService) SetUserPortfolioDetails(val bool) *servicesListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *servicesListService) SetUserPreferredDetails(val bool) *servicesListService {
	s.userPreferredDetails = &val
	return s
}

func (s *servicesListService) SetUserBadgeDetails(val bool) *servicesListService {
	s.userBadgeDetails = &val
	return s
}

func (s *servicesListService) SetUserStatus(val bool) *servicesListService {
	s.userStatus = &val
	return s
}

func (s *servicesListService) SetUserReputation(val bool) *servicesListService {
	s.userReputation = &val
	return s
}

func (s *servicesListService) SetUserEmployerReputation(val bool) *servicesListService {
	s.userEmployerReputation = &val
	return s
}

func (s *servicesListService) SetUserReputationExtra(val bool) *servicesListService {
	s.userReputationExtra = &val
	return s
}

func (s *servicesListService) SetUserEmployerReputationExtra(val bool) *servicesListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *servicesListService) SetUserCoverImage(val bool) *servicesListService {
	s.userCoverImage = &val
	return s
}

func (s *servicesListService) SetUserPastCoverImage(val bool) *servicesListService {
	s.userPastCoverImage = &val
	return s
}

func (s *servicesListService) SetUserRecommendations(val bool) *servicesListService {
	s.userRecommendations = &val
	return s
}

func (s *servicesListService) SetUserResponsiveness(val bool) *servicesListService {
	s.userResponsiveness = &val
	return s
}

func (s *servicesListService) SetCorporateUsers(val bool) *servicesListService {
	s.corporateUsers = &val
	return s
}

func (s *servicesListService) SetMarketingMobileNumber(val bool) *servicesListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *servicesListService) SetSanctionDetails(val bool) *servicesListService {
	s.sanctionDetails = &val
	return s
}

func (s *servicesListService) SetLimitedAccount(val bool) *servicesListService {
	s.limitedAccount = &val
	return s
}

func (s *servicesListService) SetEquipmentGroupDetails(val bool) *servicesListService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *servicesListService) SetLimit(limit int) *servicesListService {
	s.limit = &limit
	return s
}

func (s *servicesListService) SetOffset(offset int) *servicesListService {
	s.offset = &offset
	return s
}

func (s *servicesListService) SetCompact(val bool) *servicesListService {
	s.compact = &val
	return s
}

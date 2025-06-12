package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type milestonesListService struct {
	client                      *Client
	projects                    []int64
	projectOwners               []int64
	bidders                     []int64
	users                       []int64
	bids                        []int
	statuses                    []MilestoneStatus
	sortField                   *SortField
	sortDirection               *SortDirection
	excludedMilestones          *bool
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
}

func (s *milestonesListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_MILESTONES),
	}

	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	for _, val := range s.projects {
		r.addParam("projects[]", val)
	}
	for _, val := range s.bidders {
		r.addParam("bidders[]", val)
	}
	for _, val := range s.projectOwners {
		r.addParam("project_owners[]", val)
	}
	for _, val := range s.users {
		r.addParam("users[]", val)
	}
	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	if s.sortField != nil {
		r.addParam("sort_field", *s.sortField)
	}
	if s.sortDirection != nil {
		r.addParam("sort_direction", *s.sortDirection)
	}
	if s.excludedMilestones != nil {
		r.addParam("excluded_milestones", *s.excludedMilestones)
	}
	if s.userAvatar != nil {
		r.addParam("user_avatar", *s.userAvatar)
	}
	if s.userCountryDetails != nil {
		r.setParam("user_country_details", *s.userCountryDetails)
	}
	if s.userProfileDescription != nil {
		r.setParam("user_profile_Description", *s.userProfileDescription)
	}
	if s.userDisplayInfo != nil {
		r.setParam("user_display_info", *s.userDisplayInfo)
	}
	if s.userJobs != nil {
		r.setParam("user_jobs", *s.userJobs)
	}
	if s.userBalanceDetails != nil {
		r.setParam("user_balance_details", *s.userBalanceDetails)
	}
	if s.userQualificationDetails != nil {
		r.setParam("user_qualification_details", *s.userQualificationDetails)
	}
	if s.userMembershipDetails != nil {
		r.setParam("user_membership_details", *s.userMembershipDetails)
	}
	if s.userFinancialDetails != nil {
		r.setParam("user_financial_details", *s.userFinancialDetails)
	}
	if s.userLocationDetails != nil {
		r.setParam("user_location_details", *s.userLocationDetails)
	}
	if s.userPortfolioDetails != nil {
		r.setParam("user_portfolio_details", *s.userPortfolioDetails)
	}
	if s.userPreferredDetails != nil {
		r.setParam("user_preferred_details", *s.userPreferredDetails)
	}
	if s.userBadgeDetails != nil {
		r.setParam("user_badge_details", *s.userBadgeDetails)
	}
	if s.userStatus != nil {
		r.setParam("user_status", *s.userStatus)
	}
	if s.userReputation != nil {
		r.setParam("user_reputation", *s.userReputation)
	}
	if s.userEmployerReputation != nil {
		r.setParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.setParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.setParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.setParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.setParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.setParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.setParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.setParam("corporate_users", *s.corporateUsers)
	}
	if s.marketingMobileNumber != nil {
		r.setParam("marketing_mobile_number", *s.marketingMobileNumber)
	}
	if s.sanctionDetails != nil {
		r.setParam("sanction_details", *s.sanctionDetails)
	}
	if s.limitedAccount != nil {
		r.setParam("limited_account", *s.limitedAccount)
	}
	if s.equipmentGroupDetails != nil {
		r.setParam("equipment_group_details", *s.equipmentGroupDetails)
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
func (s *milestonesListService) SetBids(vals []int) *milestonesListService {
	s.bids = vals
	return s
}

func (s *milestonesListService) SetProjects(vals []int64) *milestonesListService {
	s.projects = vals
	return s
}

func (s *milestonesListService) SetBidders(vals []int64) *milestonesListService {
	s.bidders = vals
	return s
}

func (s *milestonesListService) SetProjectOwners(vals []int64) *milestonesListService {
	s.projectOwners = vals
	return s
}

func (s *milestonesListService) SetUsers(vals []int64) *milestonesListService {
	s.users = vals
	return s
}

func (s *milestonesListService) SetStatuses(vals []MilestoneStatus) *milestonesListService {
	s.statuses = vals
	return s
}

func (s *milestonesListService) SetSortField(val SortField) *milestonesListService {
	s.sortField = &val
	return s
}

func (s *milestonesListService) SetSortDirection(val SortDirection) *milestonesListService {
	s.sortDirection = &val
	return s
}

func (s *milestonesListService) SetExcludedMilestones(val bool) *milestonesListService {
	s.excludedMilestones = &val
	return s
}

func (s *milestonesListService) SetUserAvatar(val bool) *milestonesListService {
	s.userAvatar = &val
	return s
}

func (s *milestonesListService) SetUserCountryDetails(val bool) *milestonesListService {
	s.userCountryDetails = &val
	return s
}

func (s *milestonesListService) SetUserProfileDescription(val bool) *milestonesListService {
	s.userProfileDescription = &val
	return s
}

func (s *milestonesListService) SetUserDisplayInfo(val bool) *milestonesListService {
	s.userDisplayInfo = &val
	return s
}

func (s *milestonesListService) SetUserJobs(val bool) *milestonesListService {
	s.userJobs = &val
	return s
}

func (s *milestonesListService) SetUserBalanceDetails(val bool) *milestonesListService {
	s.userBalanceDetails = &val
	return s
}

func (s *milestonesListService) SetUserQualificationDetails(val bool) *milestonesListService {
	s.userQualificationDetails = &val
	return s
}

func (s *milestonesListService) SetUserMembershipDetails(val bool) *milestonesListService {
	s.userMembershipDetails = &val
	return s
}

func (s *milestonesListService) SetUserFinancialDetails(val bool) *milestonesListService {
	s.userFinancialDetails = &val
	return s
}

func (s *milestonesListService) SetUserLocationDetails(val bool) *milestonesListService {
	s.userLocationDetails = &val
	return s
}

func (s *milestonesListService) SetUserPortfolioDetails(val bool) *milestonesListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *milestonesListService) SetUserPreferredDetails(val bool) *milestonesListService {
	s.userPreferredDetails = &val
	return s
}

func (s *milestonesListService) SetUserBadgeDetails(val bool) *milestonesListService {
	s.userBadgeDetails = &val
	return s
}

func (s *milestonesListService) SetUserStatus(val bool) *milestonesListService {
	s.userStatus = &val
	return s
}

func (s *milestonesListService) SetUserReputation(val bool) *milestonesListService {
	s.userReputation = &val
	return s
}

func (s *milestonesListService) SetUserEmployerReputation(val bool) *milestonesListService {
	s.userEmployerReputation = &val
	return s
}

func (s *milestonesListService) SetUserReputationExtra(val bool) *milestonesListService {
	s.userReputationExtra = &val
	return s
}

func (s *milestonesListService) SetUserEmployerReputationExtra(val bool) *milestonesListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *milestonesListService) SetUserCoverImage(val bool) *milestonesListService {
	s.userCoverImage = &val
	return s
}

func (s *milestonesListService) SetUserPastCoverImage(val bool) *milestonesListService {
	s.userPastCoverImage = &val
	return s
}

func (s *milestonesListService) SetUserRecommendations(val bool) *milestonesListService {
	s.userRecommendations = &val
	return s
}

func (s *milestonesListService) SetUserResponsiveness(val bool) *milestonesListService {
	s.userResponsiveness = &val
	return s
}

func (s *milestonesListService) SetCorporateUsers(val bool) *milestonesListService {
	s.corporateUsers = &val
	return s
}

func (s *milestonesListService) SetMarketingMobileNumber(val bool) *milestonesListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *milestonesListService) SetSanctionDetails(val bool) *milestonesListService {
	s.sanctionDetails = &val
	return s
}

func (s *milestonesListService) SetLimitedAccount(val bool) *milestonesListService {
	s.limitedAccount = &val
	return s
}

func (s *milestonesListService) SetEquipmentGroupDetails(val bool) *milestonesListService {
	s.equipmentGroupDetails = &val
	return s
}

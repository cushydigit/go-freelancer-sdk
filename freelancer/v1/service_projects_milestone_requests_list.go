package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns bids for a single project.
type milestoneRequestsListService struct {
	client                      *Client
	milestoneRequests           []int
	projects                    []int64
	projectOwners               []int64
	bidders                     []int64
	users                       []int64
	bids                        []int
	statuses                    []MilestoneStatus
	fromTime                    *int64
	toTime                      *int64
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
	limit                       *int
	offset                      *int
}

func (s *milestoneRequestsListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_MILESTONE_REQUESTS),
	}

	for _, val := range s.milestoneRequests {
		r.addParam("milestone_requests[]", val)
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
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
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
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
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

func (s *milestoneRequestsListService) SetMilestoneRequests(vals []int) *milestoneRequestsListService {
	s.milestoneRequests = vals
	return s
}

func (s *milestoneRequestsListService) SetBids(vals []int) *milestoneRequestsListService {
	s.bids = vals
	return s
}

func (s *milestoneRequestsListService) SetProjects(vals []int64) *milestoneRequestsListService {
	s.projects = vals
	return s
}

func (s *milestoneRequestsListService) SetBidders(vals []int64) *milestoneRequestsListService {
	s.bidders = vals
	return s
}

func (s *milestoneRequestsListService) SetProjectOwners(vals []int64) *milestoneRequestsListService {
	s.projectOwners = vals
	return s
}

func (s *milestoneRequestsListService) SetUsers(vals []int64) *milestoneRequestsListService {
	s.users = vals
	return s
}

func (s *milestoneRequestsListService) SetStatuses(vals []MilestoneStatus) *milestoneRequestsListService {
	s.statuses = vals
	return s
}

func (s *milestoneRequestsListService) SetFromTime(val int64) *milestoneRequestsListService {
	s.fromTime = &val
	return s
}

func (s *milestoneRequestsListService) SetToTime(val int64) *milestoneRequestsListService {
	s.toTime = &val
	return s
}

func (s *milestoneRequestsListService) SetSortField(val SortField) *milestoneRequestsListService {
	s.sortField = &val
	return s
}

func (s *milestoneRequestsListService) SetSortDirection(val SortDirection) *milestoneRequestsListService {
	s.sortDirection = &val
	return s
}

func (s *milestoneRequestsListService) SetExcludedMilestones(val bool) *milestoneRequestsListService {
	s.excludedMilestones = &val
	return s
}

func (s *milestoneRequestsListService) SetUserAvatar(val bool) *milestoneRequestsListService {
	s.userAvatar = &val
	return s
}

func (s *milestoneRequestsListService) SetUserCountryDetails(val bool) *milestoneRequestsListService {
	s.userCountryDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserProfileDescription(val bool) *milestoneRequestsListService {
	s.userProfileDescription = &val
	return s
}

func (s *milestoneRequestsListService) SetUserDisplayInfo(val bool) *milestoneRequestsListService {
	s.userDisplayInfo = &val
	return s
}

func (s *milestoneRequestsListService) SetUserJobs(val bool) *milestoneRequestsListService {
	s.userJobs = &val
	return s
}

func (s *milestoneRequestsListService) SetUserBalanceDetails(val bool) *milestoneRequestsListService {
	s.userBalanceDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserQualificationDetails(val bool) *milestoneRequestsListService {
	s.userQualificationDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserMembershipDetails(val bool) *milestoneRequestsListService {
	s.userMembershipDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserFinancialDetails(val bool) *milestoneRequestsListService {
	s.userFinancialDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserLocationDetails(val bool) *milestoneRequestsListService {
	s.userLocationDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserPortfolioDetails(val bool) *milestoneRequestsListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserPreferredDetails(val bool) *milestoneRequestsListService {
	s.userPreferredDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserBadgeDetails(val bool) *milestoneRequestsListService {
	s.userBadgeDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetUserStatus(val bool) *milestoneRequestsListService {
	s.userStatus = &val
	return s
}

func (s *milestoneRequestsListService) SetUserReputation(val bool) *milestoneRequestsListService {
	s.userReputation = &val
	return s
}

func (s *milestoneRequestsListService) SetUserEmployerReputation(val bool) *milestoneRequestsListService {
	s.userEmployerReputation = &val
	return s
}

func (s *milestoneRequestsListService) SetUserReputationExtra(val bool) *milestoneRequestsListService {
	s.userReputationExtra = &val
	return s
}

func (s *milestoneRequestsListService) SetUserEmployerReputationExtra(val bool) *milestoneRequestsListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *milestoneRequestsListService) SetUserCoverImage(val bool) *milestoneRequestsListService {
	s.userCoverImage = &val
	return s
}

func (s *milestoneRequestsListService) SetUserPastCoverImage(val bool) *milestoneRequestsListService {
	s.userPastCoverImage = &val
	return s
}

func (s *milestoneRequestsListService) SetUserRecommendations(val bool) *milestoneRequestsListService {
	s.userRecommendations = &val
	return s
}

func (s *milestoneRequestsListService) SetUserResponsiveness(val bool) *milestoneRequestsListService {
	s.userResponsiveness = &val
	return s
}

func (s *milestoneRequestsListService) SetCorporateUsers(val bool) *milestoneRequestsListService {
	s.corporateUsers = &val
	return s
}

func (s *milestoneRequestsListService) SetMarketingMobileNumber(val bool) *milestoneRequestsListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *milestoneRequestsListService) SetSanctionDetails(val bool) *milestoneRequestsListService {
	s.sanctionDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetLimitedAccount(val bool) *milestoneRequestsListService {
	s.limitedAccount = &val
	return s
}

func (s *milestoneRequestsListService) SetEquipmentGroupDetails(val bool) *milestoneRequestsListService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *milestoneRequestsListService) SetLimit(val int) *milestoneRequestsListService {
	s.limit = &val
	return s
}

func (s *milestoneRequestsListService) SetOffset(val int) *milestoneRequestsListService {
	s.offset = &val
	return s
}

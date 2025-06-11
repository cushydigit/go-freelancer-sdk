package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns bids for a single project.
type listMilestoneRequestsService struct {
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

func (s *listMilestoneRequestsService) Do(ctx context.Context) (*BaseResponse, error) {
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

func (s *listMilestoneRequestsService) SetMilestoneRequests(vals []int) *listMilestoneRequestsService {
	s.milestoneRequests = vals
	return s
}

func (s *listMilestoneRequestsService) SetBids(vals []int) *listMilestoneRequestsService {
	s.bids = vals
	return s
}

func (s *listMilestoneRequestsService) SetProjects(vals []int64) *listMilestoneRequestsService {
	s.projects = vals
	return s
}

func (s *listMilestoneRequestsService) SetBidders(vals []int64) *listMilestoneRequestsService {
	s.bidders = vals
	return s
}

func (s *listMilestoneRequestsService) SetProjectOwners(vals []int64) *listMilestoneRequestsService {
	s.projectOwners = vals
	return s
}

func (s *listMilestoneRequestsService) SetUsers(vals []int64) *listMilestoneRequestsService {
	s.users = vals
	return s
}

func (s *listMilestoneRequestsService) SetStatuses(vals []MilestoneStatus) *listMilestoneRequestsService {
	s.statuses = vals
	return s
}

func (s *listMilestoneRequestsService) SetFromTime(val int64) *listMilestoneRequestsService {
	s.fromTime = &val
	return s
}

func (s *listMilestoneRequestsService) SetToTime(val int64) *listMilestoneRequestsService {
	s.toTime = &val
	return s
}

func (s *listMilestoneRequestsService) SetSortField(val SortField) *listMilestoneRequestsService {
	s.sortField = &val
	return s
}

func (s *listMilestoneRequestsService) SetSortDirection(val SortDirection) *listMilestoneRequestsService {
	s.sortDirection = &val
	return s
}

func (s *listMilestoneRequestsService) SetExcludedMilestones(val bool) *listMilestoneRequestsService {
	s.excludedMilestones = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserAvatar(val bool) *listMilestoneRequestsService {
	s.userAvatar = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserCountryDetails(val bool) *listMilestoneRequestsService {
	s.userCountryDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserProfileDescription(val bool) *listMilestoneRequestsService {
	s.userProfileDescription = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserDisplayInfo(val bool) *listMilestoneRequestsService {
	s.userDisplayInfo = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserJobs(val bool) *listMilestoneRequestsService {
	s.userJobs = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserBalanceDetails(val bool) *listMilestoneRequestsService {
	s.userBalanceDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserQualificationDetails(val bool) *listMilestoneRequestsService {
	s.userQualificationDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserMembershipDetails(val bool) *listMilestoneRequestsService {
	s.userMembershipDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserFinancialDetails(val bool) *listMilestoneRequestsService {
	s.userFinancialDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserLocationDetails(val bool) *listMilestoneRequestsService {
	s.userLocationDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserPortfolioDetails(val bool) *listMilestoneRequestsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserPreferredDetails(val bool) *listMilestoneRequestsService {
	s.userPreferredDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserBadgeDetails(val bool) *listMilestoneRequestsService {
	s.userBadgeDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserStatus(val bool) *listMilestoneRequestsService {
	s.userStatus = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserReputation(val bool) *listMilestoneRequestsService {
	s.userReputation = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserEmployerReputation(val bool) *listMilestoneRequestsService {
	s.userEmployerReputation = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserReputationExtra(val bool) *listMilestoneRequestsService {
	s.userReputationExtra = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserEmployerReputationExtra(val bool) *listMilestoneRequestsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserCoverImage(val bool) *listMilestoneRequestsService {
	s.userCoverImage = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserPastCoverImage(val bool) *listMilestoneRequestsService {
	s.userPastCoverImage = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserRecommendations(val bool) *listMilestoneRequestsService {
	s.userRecommendations = &val
	return s
}

func (s *listMilestoneRequestsService) SetUserResponsiveness(val bool) *listMilestoneRequestsService {
	s.userResponsiveness = &val
	return s
}

func (s *listMilestoneRequestsService) SetCorporateUsers(val bool) *listMilestoneRequestsService {
	s.corporateUsers = &val
	return s
}

func (s *listMilestoneRequestsService) SetMarketingMobileNumber(val bool) *listMilestoneRequestsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *listMilestoneRequestsService) SetSanctionDetails(val bool) *listMilestoneRequestsService {
	s.sanctionDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetLimitedAccount(val bool) *listMilestoneRequestsService {
	s.limitedAccount = &val
	return s
}

func (s *listMilestoneRequestsService) SetEquipmentGroupDetails(val bool) *listMilestoneRequestsService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *listMilestoneRequestsService) SetLimit(val int) *listMilestoneRequestsService {
	s.limit = &val
	return s
}

func (s *listMilestoneRequestsService) SetOffset(val int) *listMilestoneRequestsService {
	s.offset = &val
	return s
}

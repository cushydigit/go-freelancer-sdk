package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns bids for a single project.
type listMilestonesService struct {
	client        *Client
	projects      []int64
	projectOwners []int64
	bidders       []int64
	users         []int64
	bids          []int
	statuses      []MilestoneStatus
	sortField     *SortField

	awardStatuses               []AwardStatusType
	paidStatuses                []PaidStatusType
	completeStatuses            []CompleteStatusType
	frontBidStatuses            []FrontendBidStatusType
	fromTime                    *int64
	toTime                      *int64
	reputation                  *bool
	buyerProjectFee             *bool
	awardStatusPossibilities    *bool
	projectDetails              *bool
	userDetails                 *bool
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

func (s *listMilestonesService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_BIDS),
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
	for _, val := range s.awardStatuses {
		r.addParam("award_statuses[]", val)
	}
	for _, val := range s.paidStatuses {
		r.addParam("paid_statuses[]", val)
	}
	for _, val := range s.completeStatuses {
		r.addParam("complete_statuses[]", val)
	}
	for _, val := range s.frontBidStatuses {
		r.addParam("frontend_bid_statuses[]", val)
	}
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.reputation != nil {
		r.addParam("reputation", *s.reputation)
	}
	if s.buyerProjectFee != nil {
		r.addParam("buyer_project_fee", *s.buyerProjectFee)
	}
	if s.awardStatusPossibilities != nil {
		r.addParam("award_status_possibilities", *s.awardStatusPossibilities)
	}
	if s.projectDetails != nil {
		r.addParam("project_details", *s.projectDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
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
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.setParam("compact", *s.compact)
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
func (s *listMilestonesService) SetBids(vals []int) *listMilestonesService {
	s.bids = vals
	return s
}

func (s *listMilestonesService) SetProjects(vals []int64) *listMilestonesService {
	s.projects = vals
	return s
}

func (s *listMilestonesService) SetBidders(vals []int64) *listMilestonesService {
	s.bidders = vals
	return s
}

func (s *listMilestonesService) SetProjectOwners(vals []int64) *listMilestonesService {
	s.projectOwners = vals
	return s
}

func (s *listMilestonesService) SetAwardStatuses(vals []AwardStatusType) *listMilestonesService {
	s.awardStatuses = vals
	return s
}

func (s *listMilestonesService) SetPaidStatuses(vals []PaidStatusType) *listMilestonesService {
	s.paidStatuses = vals
	return s
}

func (s *listMilestonesService) SetCompleteStatuses(vals []CompleteStatusType) *listMilestonesService {
	s.completeStatuses = vals
	return s
}

func (s *listMilestonesService) SetFrontBidStatuses(vals []FrontendBidStatusType) *listMilestonesService {
	s.frontBidStatuses = vals
	return s
}

func (s *listMilestonesService) SetFromTime(val int64) *listMilestonesService {
	s.fromTime = &val
	return s
}

func (s *listMilestonesService) SetToTime(val int64) *listMilestonesService {
	s.toTime = &val
	return s
}

func (s *listMilestonesService) SetReputation(val bool) *listMilestonesService {
	s.reputation = &val
	return s
}

func (s *listMilestonesService) SetBuyerProjectFee(val bool) *listMilestonesService {
	s.buyerProjectFee = &val
	return s
}

func (s *listMilestonesService) SetAwardStatusPossibilities(val bool) *listMilestonesService {
	s.awardStatusPossibilities = &val
	return s
}

func (s *listMilestonesService) SetProjectDetails(val bool) *listMilestonesService {
	s.projectDetails = &val
	return s
}

func (s *listMilestonesService) SetUserDetails(val bool) *listMilestonesService {
	s.userDetails = &val
	return s
}

func (s *listMilestonesService) SetUserAvatar(val bool) *listMilestonesService {
	s.userAvatar = &val
	return s
}

func (s *listMilestonesService) SetUserCountryDetails(val bool) *listMilestonesService {
	s.userCountryDetails = &val
	return s
}

func (s *listMilestonesService) SetUserProfileDescription(val bool) *listMilestonesService {
	s.userProfileDescription = &val
	return s
}

func (s *listMilestonesService) SetUserDisplayInfo(val bool) *listMilestonesService {
	s.userDisplayInfo = &val
	return s
}

func (s *listMilestonesService) SetUserJobs(val bool) *listMilestonesService {
	s.userJobs = &val
	return s
}

func (s *listMilestonesService) SetUserBalanceDetails(val bool) *listMilestonesService {
	s.userBalanceDetails = &val
	return s
}

func (s *listMilestonesService) SetUserQualificationDetails(val bool) *listMilestonesService {
	s.userQualificationDetails = &val
	return s
}

func (s *listMilestonesService) SetUserMembershipDetails(val bool) *listMilestonesService {
	s.userMembershipDetails = &val
	return s
}

func (s *listMilestonesService) SetUserFinancialDetails(val bool) *listMilestonesService {
	s.userFinancialDetails = &val
	return s
}

func (s *listMilestonesService) SetUserLocationDetails(val bool) *listMilestonesService {
	s.userLocationDetails = &val
	return s
}

func (s *listMilestonesService) SetUserPortfolioDetails(val bool) *listMilestonesService {
	s.userPortfolioDetails = &val
	return s
}

func (s *listMilestonesService) SetUserPreferredDetails(val bool) *listMilestonesService {
	s.userPreferredDetails = &val
	return s
}

func (s *listMilestonesService) SetUserBadgeDetails(val bool) *listMilestonesService {
	s.userBadgeDetails = &val
	return s
}

func (s *listMilestonesService) SetUserStatus(val bool) *listMilestonesService {
	s.userStatus = &val
	return s
}

func (s *listMilestonesService) SetUserReputation(val bool) *listMilestonesService {
	s.userReputation = &val
	return s
}

func (s *listMilestonesService) SetUserEmployerReputation(val bool) *listMilestonesService {
	s.userEmployerReputation = &val
	return s
}

func (s *listMilestonesService) SetUserReputationExtra(val bool) *listMilestonesService {
	s.userReputationExtra = &val
	return s
}

func (s *listMilestonesService) SetUserEmployerReputationExtra(val bool) *listMilestonesService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *listMilestonesService) SetUserCoverImage(val bool) *listMilestonesService {
	s.userCoverImage = &val
	return s
}

func (s *listMilestonesService) SetUserPastCoverImage(val bool) *listMilestonesService {
	s.userPastCoverImage = &val
	return s
}

func (s *listMilestonesService) SetUserRecommendations(val bool) *listMilestonesService {
	s.userRecommendations = &val
	return s
}

func (s *listMilestonesService) SetUserResponsiveness(val bool) *listMilestonesService {
	s.userResponsiveness = &val
	return s
}

func (s *listMilestonesService) SetCorporateUsers(val bool) *listMilestonesService {
	s.corporateUsers = &val
	return s
}

func (s *listMilestonesService) SetMarketingMobileNumber(val bool) *listMilestonesService {
	s.marketingMobileNumber = &val
	return s
}

func (s *listMilestonesService) SetSanctionDetails(val bool) *listMilestonesService {
	s.sanctionDetails = &val
	return s
}

func (s *listMilestonesService) SetLimitedAccount(val bool) *listMilestonesService {
	s.limitedAccount = &val
	return s
}

func (s *listMilestonesService) SetEquipmentGroupDetails(val bool) *listMilestonesService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *listMilestonesService) SetLimit(val int) *listMilestonesService {
	s.limit = &val
	return s
}

func (s *listMilestonesService) SetOffset(val int) *listMilestonesService {
	s.offset = &val
	return s
}

func (s *listMilestonesService) SetCompact(val bool) *listMilestonesService {
	s.compact = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns bids for a single project.
type ListBidsService struct {
	client                      *Client
	bids                        []int
	projects                    []int64
	bidders                     []int64
	projectOwners               []int64
	awardStatuses               []AwardStatus
	paidStatuses                []PaidStatus
	completeStatuses            []CompleteStatus
	frontBidStatuses            []FrontendBidStatus
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

func (s *ListBidsService) Do(ctx context.Context) (*BaseResponse, error) {
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
func (s *ListBidsService) SetBids(vals []int) *ListBidsService {
	s.bids = vals
	return s
}

func (s *ListBidsService) SetProjects(vals []int64) *ListBidsService {
	s.projects = vals
	return s
}

func (s *ListBidsService) SetBidders(vals []int64) *ListBidsService {
	s.bidders = vals
	return s
}

func (s *ListBidsService) SetProjectOwners(vals []int64) *ListBidsService {
	s.projectOwners = vals
	return s
}

func (s *ListBidsService) SetAwardStatuses(vals []AwardStatus) *ListBidsService {
	s.awardStatuses = vals
	return s
}

func (s *ListBidsService) SetPaidStatuses(vals []PaidStatus) *ListBidsService {
	s.paidStatuses = vals
	return s
}

func (s *ListBidsService) SetCompleteStatuses(vals []CompleteStatus) *ListBidsService {
	s.completeStatuses = vals
	return s
}

func (s *ListBidsService) SetFrontBidStatuses(vals []FrontendBidStatus) *ListBidsService {
	s.frontBidStatuses = vals
	return s
}

func (s *ListBidsService) SetFromTime(val int64) *ListBidsService {
	s.fromTime = &val
	return s
}

func (s *ListBidsService) SetToTime(val int64) *ListBidsService {
	s.toTime = &val
	return s
}

func (s *ListBidsService) SetReputation(val bool) *ListBidsService {
	s.reputation = &val
	return s
}

func (s *ListBidsService) SetBuyerProjectFee(val bool) *ListBidsService {
	s.buyerProjectFee = &val
	return s
}

func (s *ListBidsService) SetAwardStatusPossibilities(val bool) *ListBidsService {
	s.awardStatusPossibilities = &val
	return s
}

func (s *ListBidsService) SetProjectDetails(val bool) *ListBidsService {
	s.projectDetails = &val
	return s
}

func (s *ListBidsService) SetUserDetails(val bool) *ListBidsService {
	s.userDetails = &val
	return s
}

func (s *ListBidsService) SetUserAvatar(val bool) *ListBidsService {
	s.userAvatar = &val
	return s
}

func (s *ListBidsService) SetUserCountryDetails(val bool) *ListBidsService {
	s.userCountryDetails = &val
	return s
}

func (s *ListBidsService) SetUserProfileDescription(val bool) *ListBidsService {
	s.userProfileDescription = &val
	return s
}

func (s *ListBidsService) SetUserDisplayInfo(val bool) *ListBidsService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListBidsService) SetUserJobs(val bool) *ListBidsService {
	s.userJobs = &val
	return s
}

func (s *ListBidsService) SetUserBalanceDetails(val bool) *ListBidsService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListBidsService) SetUserQualificationDetails(val bool) *ListBidsService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListBidsService) SetUserMembershipDetails(val bool) *ListBidsService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListBidsService) SetUserFinancialDetails(val bool) *ListBidsService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListBidsService) SetUserLocationDetails(val bool) *ListBidsService {
	s.userLocationDetails = &val
	return s
}

func (s *ListBidsService) SetUserPortfolioDetails(val bool) *ListBidsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListBidsService) SetUserPreferredDetails(val bool) *ListBidsService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListBidsService) SetUserBadgeDetails(val bool) *ListBidsService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListBidsService) SetUserStatus(val bool) *ListBidsService {
	s.userStatus = &val
	return s
}

func (s *ListBidsService) SetUserReputation(val bool) *ListBidsService {
	s.userReputation = &val
	return s
}

func (s *ListBidsService) SetUserEmployerReputation(val bool) *ListBidsService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListBidsService) SetUserReputationExtra(val bool) *ListBidsService {
	s.userReputationExtra = &val
	return s
}

func (s *ListBidsService) SetUserEmployerReputationExtra(val bool) *ListBidsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListBidsService) SetUserCoverImage(val bool) *ListBidsService {
	s.userCoverImage = &val
	return s
}

func (s *ListBidsService) SetUserPastCoverImage(val bool) *ListBidsService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListBidsService) SetUserRecommendations(val bool) *ListBidsService {
	s.userRecommendations = &val
	return s
}

func (s *ListBidsService) SetUserResponsiveness(val bool) *ListBidsService {
	s.userResponsiveness = &val
	return s
}

func (s *ListBidsService) SetCorporateUsers(val bool) *ListBidsService {
	s.corporateUsers = &val
	return s
}

func (s *ListBidsService) SetMarketingMobileNumber(val bool) *ListBidsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListBidsService) SetSanctionDetails(val bool) *ListBidsService {
	s.sanctionDetails = &val
	return s
}

func (s *ListBidsService) SetLimitedAccount(val bool) *ListBidsService {
	s.limitedAccount = &val
	return s
}

func (s *ListBidsService) SetEquipmentGroupDetails(val bool) *ListBidsService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *ListBidsService) SetLimit(val int) *ListBidsService {
	s.limit = &val
	return s
}

func (s *ListBidsService) SetOffset(val int) *ListBidsService {
	s.offset = &val
	return s
}

func (s *ListBidsService) SetCompact(val bool) *ListBidsService {
	s.compact = &val
	return s
}

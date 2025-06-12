package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns bids for a single project.
type bidsListService struct {
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

func (s *bidsListService) Do(ctx context.Context) (*BaseResponse, error) {
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
func (s *bidsListService) SetBids(vals []int) *bidsListService {
	s.bids = vals
	return s
}

func (s *bidsListService) SetProjects(vals []int64) *bidsListService {
	s.projects = vals
	return s
}

func (s *bidsListService) SetBidders(vals []int64) *bidsListService {
	s.bidders = vals
	return s
}

func (s *bidsListService) SetProjectOwners(vals []int64) *bidsListService {
	s.projectOwners = vals
	return s
}

func (s *bidsListService) SetAwardStatuses(vals []AwardStatus) *bidsListService {
	s.awardStatuses = vals
	return s
}

func (s *bidsListService) SetPaidStatuses(vals []PaidStatus) *bidsListService {
	s.paidStatuses = vals
	return s
}

func (s *bidsListService) SetCompleteStatuses(vals []CompleteStatus) *bidsListService {
	s.completeStatuses = vals
	return s
}

func (s *bidsListService) SetFrontBidStatuses(vals []FrontendBidStatus) *bidsListService {
	s.frontBidStatuses = vals
	return s
}

func (s *bidsListService) SetFromTime(val int64) *bidsListService {
	s.fromTime = &val
	return s
}

func (s *bidsListService) SetToTime(val int64) *bidsListService {
	s.toTime = &val
	return s
}

func (s *bidsListService) SetReputation(val bool) *bidsListService {
	s.reputation = &val
	return s
}

func (s *bidsListService) SetBuyerProjectFee(val bool) *bidsListService {
	s.buyerProjectFee = &val
	return s
}

func (s *bidsListService) SetAwardStatusPossibilities(val bool) *bidsListService {
	s.awardStatusPossibilities = &val
	return s
}

func (s *bidsListService) SetProjectDetails(val bool) *bidsListService {
	s.projectDetails = &val
	return s
}

func (s *bidsListService) SetUserDetails(val bool) *bidsListService {
	s.userDetails = &val
	return s
}

func (s *bidsListService) SetUserAvatar(val bool) *bidsListService {
	s.userAvatar = &val
	return s
}

func (s *bidsListService) SetUserCountryDetails(val bool) *bidsListService {
	s.userCountryDetails = &val
	return s
}

func (s *bidsListService) SetUserProfileDescription(val bool) *bidsListService {
	s.userProfileDescription = &val
	return s
}

func (s *bidsListService) SetUserDisplayInfo(val bool) *bidsListService {
	s.userDisplayInfo = &val
	return s
}

func (s *bidsListService) SetUserJobs(val bool) *bidsListService {
	s.userJobs = &val
	return s
}

func (s *bidsListService) SetUserBalanceDetails(val bool) *bidsListService {
	s.userBalanceDetails = &val
	return s
}

func (s *bidsListService) SetUserQualificationDetails(val bool) *bidsListService {
	s.userQualificationDetails = &val
	return s
}

func (s *bidsListService) SetUserMembershipDetails(val bool) *bidsListService {
	s.userMembershipDetails = &val
	return s
}

func (s *bidsListService) SetUserFinancialDetails(val bool) *bidsListService {
	s.userFinancialDetails = &val
	return s
}

func (s *bidsListService) SetUserLocationDetails(val bool) *bidsListService {
	s.userLocationDetails = &val
	return s
}

func (s *bidsListService) SetUserPortfolioDetails(val bool) *bidsListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *bidsListService) SetUserPreferredDetails(val bool) *bidsListService {
	s.userPreferredDetails = &val
	return s
}

func (s *bidsListService) SetUserBadgeDetails(val bool) *bidsListService {
	s.userBadgeDetails = &val
	return s
}

func (s *bidsListService) SetUserStatus(val bool) *bidsListService {
	s.userStatus = &val
	return s
}

func (s *bidsListService) SetUserReputation(val bool) *bidsListService {
	s.userReputation = &val
	return s
}

func (s *bidsListService) SetUserEmployerReputation(val bool) *bidsListService {
	s.userEmployerReputation = &val
	return s
}

func (s *bidsListService) SetUserReputationExtra(val bool) *bidsListService {
	s.userReputationExtra = &val
	return s
}

func (s *bidsListService) SetUserEmployerReputationExtra(val bool) *bidsListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *bidsListService) SetUserCoverImage(val bool) *bidsListService {
	s.userCoverImage = &val
	return s
}

func (s *bidsListService) SetUserPastCoverImage(val bool) *bidsListService {
	s.userPastCoverImage = &val
	return s
}

func (s *bidsListService) SetUserRecommendations(val bool) *bidsListService {
	s.userRecommendations = &val
	return s
}

func (s *bidsListService) SetUserResponsiveness(val bool) *bidsListService {
	s.userResponsiveness = &val
	return s
}

func (s *bidsListService) SetCorporateUsers(val bool) *bidsListService {
	s.corporateUsers = &val
	return s
}

func (s *bidsListService) SetMarketingMobileNumber(val bool) *bidsListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *bidsListService) SetSanctionDetails(val bool) *bidsListService {
	s.sanctionDetails = &val
	return s
}

func (s *bidsListService) SetLimitedAccount(val bool) *bidsListService {
	s.limitedAccount = &val
	return s
}

func (s *bidsListService) SetEquipmentGroupDetails(val bool) *bidsListService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *bidsListService) SetLimit(val int) *bidsListService {
	s.limit = &val
	return s
}

func (s *bidsListService) SetOffset(val int) *bidsListService {
	s.offset = &val
	return s
}

func (s *bidsListService) SetCompact(val bool) *bidsListService {
	s.compact = &val
	return s
}

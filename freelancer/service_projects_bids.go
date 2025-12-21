package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: refine with typed response
type listBids struct {
	client                      *Client
	bids                        []int
	projects                    []int64
	bidders                     []int64
	projectOwners               []int64
	awardStatuses               []BidAwardStatus
	paidStatuses                []BidPaidStatus
	completeStatuses            []BidCompleteStatus
	frontBidStatuses            []BidFrontendStatus
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

func (s *listBids) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(ProjectsBids),
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

	return execute[*RawResponse](ctx, s.client, r)
}

// SetBids filters results by specific bid IDs.
func (s *listBids) SetBids(vals []int) *listBids {
	s.bids = vals
	return s
}

// SetProjects filters results by project IDs.
func (s *listBids) SetProjects(vals []int64) *listBids {
	s.projects = vals
	return s
}

// SetBidders filters results by bidder (freelancer) user IDs.
func (s *listBids) SetBidders(vals []int64) *listBids {
	s.bidders = vals
	return s
}

// SetProjectOwners filters results by project owner user IDs.
func (s *listBids) SetProjectOwners(vals []int64) *listBids {
	s.projectOwners = vals
	return s
}

// SetAwardStatuses filters results by award statuses.
func (s *listBids) SetAwardStatuses(vals []BidAwardStatus) *listBids {
	s.awardStatuses = vals
	return s
}

// SetPaidStatuses filters results by payment statuses.
func (s *listBids) SetPaidStatuses(vals []BidPaidStatus) *listBids {
	s.paidStatuses = vals
	return s
}

// SetCompleteStatuses filters results by bid completion statuses.
func (s *listBids) SetCompleteStatuses(vals []BidCompleteStatus) *listBids {
	s.completeStatuses = vals
	return s
}

// SetFrontBidStatuses filters results by frontend bid statuses.
func (s *listBids) SetFrontBidStatuses(vals []BidFrontendStatus) *listBids {
	s.frontBidStatuses = vals
	return s
}

// SetFromTime filters bids created after the given Unix timestamp (inclusive).
func (s *listBids) SetFromTime(val int64) *listBids {
	s.fromTime = &val
	return s
}

// SetToTime filters bids created before the given Unix timestamp (inclusive).
func (s *listBids) SetToTime(val int64) *listBids {
	s.toTime = &val
	return s
}

// SetReputation includes reputation information relevant to the bid.
func (s *listBids) SetReputation(val bool) *listBids {
	s.reputation = &val
	return s
}

// SetBuyerProjectFee includes project fee information if the bid is awarded or accepted.
func (s *listBids) SetBuyerProjectFee(val bool) *listBids {
	s.buyerProjectFee = &val
	return s
}

// SetAwardStatusPossibilities includes possible transactions for award statuses.
func (s *listBids) SetAwardStatusPossibilities(val bool) *listBids {
	s.awardStatusPossibilities = &val
	return s
}

// SetProjectDetails includes basic project information for each bid.
func (s *listBids) SetProjectDetails(val bool) *listBids {
	s.projectDetails = &val
	return s
}

// SetUserDetails includes basic user information for each bid.
func (s *listBids) SetUserDetails(val bool) *listBids {
	s.userDetails = &val
	return s
}

// SetUserAvatar includes avatar of the associated user.
func (s *listBids) SetUserAvatar(val bool) *listBids {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails includes country code and flag information for users.
func (s *listBids) SetUserCountryDetails(val bool) *listBids {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription includes the user's profile description.
func (s *listBids) SetUserProfileDescription(val bool) *listBids {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo includes the user's display information.
func (s *listBids) SetUserDisplayInfo(val bool) *listBids {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs includes job history information for users.
func (s *listBids) SetUserJobs(val bool) *listBids {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails includes currency balance information for users.
func (s *listBids) SetUserBalanceDetails(val bool) *listBids {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails includes completed qualification exams for users.
func (s *listBids) SetUserQualificationDetails(val bool) *listBids {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails includes membership information for users.
func (s *listBids) SetUserMembershipDetails(val bool) *listBids {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails includes financial information for users.
func (s *listBids) SetUserFinancialDetails(val bool) *listBids {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails includes location information for users.
func (s *listBids) SetUserLocationDetails(val bool) *listBids {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails includes portfolio information for users.
func (s *listBids) SetUserPortfolioDetails(val bool) *listBids {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails includes preferred user information.
func (s *listBids) SetUserPreferredDetails(val bool) *listBids {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails includes earned badges for users.
func (s *listBids) SetUserBadgeDetails(val bool) *listBids {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus includes additional status information for users.
func (s *listBids) SetUserStatus(val bool) *listBids {
	s.userStatus = &val
	return s
}

// SetUserReputation includes freelancer reputation information.
func (s *listBids) SetUserReputation(val bool) *listBids {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation includes employer reputation information.
func (s *listBids) SetUserEmployerReputation(val bool) *listBids {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra includes full freelancer reputation details.
func (s *listBids) SetUserReputationExtra(val bool) *listBids {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra includes full employer reputation details.
func (s *listBids) SetUserEmployerReputationExtra(val bool) *listBids {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage includes the user's cover image.
func (s *listBids) SetUserCoverImage(val bool) *listBids {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage includes past user cover images.
func (s *listBids) SetUserPastCoverImage(val bool) *listBids {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations includes recommendation count for users.
func (s *listBids) SetUserRecommendations(val bool) *listBids {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness includes responsiveness scores for users.
func (s *listBids) SetUserResponsiveness(val bool) *listBids {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers includes corporate account information for users.
func (s *listBids) SetCorporateUsers(val bool) *listBids {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber includes recruiter-visible mobile number information.
func (s *listBids) SetMarketingMobileNumber(val bool) *listBids {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails includes sanction end timestamp information.
func (s *listBids) SetSanctionDetails(val bool) *listBids {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount includes limited account status information.
func (s *listBids) SetLimitedAccount(val bool) *listBids {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails includes equipment groups and items attached to users.
func (s *listBids) SetEquipmentGroupDetails(val bool) *listBids {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *listBids) SetLimit(val int) *listBids {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *listBids) SetOffset(val int) *listBids {
	s.offset = &val
	return s
}

// SetCompact removes null and empty values from the response when enabled.
func (s *listBids) SetCompact(val bool) *listBids {
	s.compact = &val
	return s
}

// TODO: refine with typed response
type getBidByID struct {
	client                      *Client
	bidID                       int
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
	quotations                  *bool
}

func (s *getBidByID) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(ProjectsBids), s.bidID),
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
	if s.quotations != nil {
		r.setParam("quotations", *s.quotations)
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// SetReputation enables or disables returning reputation information relevant to the bid.
func (s *getBidByID) SetReputation(val bool) *getBidByID {
	s.reputation = &val
	return s
}

// SetBuyerProjectFee enables or disables returning the buyer project fee
// if the bid is awarded or accepted.
func (s *getBidByID) SetBuyerProjectFee(val bool) *getBidByID {
	s.buyerProjectFee = &val
	return s
}

// SetAwardStatusPossibilities enables or disables returning possible
// states that the award_status can transition to.
func (s *getBidByID) SetAwardStatusPossibilities(val bool) *getBidByID {
	s.awardStatusPossibilities = &val
	return s
}

// SetProjectDetails enables or disables returning basic project
// information associated with the bid.
func (s *getBidByID) SetProjectDetails(val bool) *getBidByID {
	s.projectDetails = &val
	return s
}

// SetUserDetails enables or disables returning basic user information
// for the bidder.
func (s *getBidByID) SetUserDetails(val bool) *getBidByID {
	s.userDetails = &val
	return s
}

// SetUserAvatar enables or disables returning the avatar
// of the selected user.
func (s *getBidByID) SetUserAvatar(val bool) *getBidByID {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails enables or disables returning the country
// flag and country code of the selected user.
func (s *getBidByID) SetUserCountryDetails(val bool) *getBidByID {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription enables or disables returning the
// profile description (blurb) of the selected user.
func (s *getBidByID) SetUserProfileDescription(val bool) *getBidByID {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo enables or disables returning display information
// such as the display name of the selected user.
func (s *getBidByID) SetUserDisplayInfo(val bool) *getBidByID {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs enables or disables returning job history
// information for the selected user.
func (s *getBidByID) SetUserJobs(val bool) *getBidByID {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails enables or disables returning currency
// balance information of the selected user.
func (s *getBidByID) SetUserBalanceDetails(val bool) *getBidByID {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails enables or disables returning qualification
// exam details completed by the selected user.
func (s *getBidByID) SetUserQualificationDetails(val bool) *getBidByID {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails enables or disables returning membership
// information of the selected user.
func (s *getBidByID) SetUserMembershipDetails(val bool) *getBidByID {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails enables or disables returning financial
// information of the selected user.
func (s *getBidByID) SetUserFinancialDetails(val bool) *getBidByID {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails enables or disables returning location
// information of the selected user.
func (s *getBidByID) SetUserLocationDetails(val bool) *getBidByID {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails enables or disables returning portfolio
// information of the selected user.
func (s *getBidByID) SetUserPortfolioDetails(val bool) *getBidByID {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails enables or disables returning preferred
// information of the selected user.
func (s *getBidByID) SetUserPreferredDetails(val bool) *getBidByID {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails enables or disables returning badges
// earned by the selected user.
func (s *getBidByID) SetUserBadgeDetails(val bool) *getBidByID {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus enables or disables returning additional
// status information about the selected user.
func (s *getBidByID) SetUserStatus(val bool) *getBidByID {
	s.userStatus = &val
	return s
}

// SetUserReputation enables or disables returning freelancer
// reputation information of the selected user.
func (s *getBidByID) SetUserReputation(val bool) *getBidByID {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation enables or disables returning employer
// reputation information of the selected user.
func (s *getBidByID) SetUserEmployerReputation(val bool) *getBidByID {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra enables or disables returning the full
// freelancer reputation details of the selected user.
func (s *getBidByID) SetUserReputationExtra(val bool) *getBidByID {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra enables or disables returning the full
// employer reputation details of the selected user.
func (s *getBidByID) SetUserEmployerReputationExtra(val bool) *getBidByID {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage enables or disables returning the profile
// cover image of the selected user.
func (s *getBidByID) SetUserCoverImage(val bool) *getBidByID {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage enables or disables returning previous
// cover images of the selected user.
func (s *getBidByID) SetUserPastCoverImage(val bool) *getBidByID {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations enables or disables returning the
// recommendations count of the selected user.
func (s *getBidByID) SetUserRecommendations(val bool) *getBidByID {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness enables or disables returning the
// responsiveness score of the selected user.
func (s *getBidByID) SetUserResponsiveness(val bool) *getBidByID {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers enables or disables returning corporate
// accounts created or founded by the selected user.
func (s *getBidByID) SetCorporateUsers(val bool) *getBidByID {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber enables or disables returning the
// marketing mobile number used by recruiters to contact the user.
func (s *getBidByID) SetMarketingMobileNumber(val bool) *getBidByID {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables or disables returning sanction
// information, including sanction end timestamp, for the user.
func (s *getBidByID) SetSanctionDetails(val bool) *getBidByID {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables or disables returning limited
// account status information of the user.
func (s *getBidByID) SetLimitedAccount(val bool) *getBidByID {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails enables or disables returning equipment
// groups and items attached to the selected user.
func (s *getBidByID) SetEquipmentGroupDetails(val bool) *getBidByID {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *getBidByID) SetLimit(val int) *getBidByID {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip,
// enabling pagination of results.
func (s *getBidByID) SetOffset(val int) *getBidByID {
	s.offset = &val
	return s
}

// SetCompact enables or disables compact mode,
// which strips null and empty values from the response.
func (s *getBidByID) SetCompact(val bool) *getBidByID {
	s.compact = &val
	return s
}

// SetQuotations enables or disables returning quotations
// attached to the bid, when permitted by user role.
func (s *getBidByID) SetQuotations(val bool) *getBidByID {
	s.quotations = &val
	return s
}

type createBid struct {
	client *Client
}

// ProjectID, BidderID, Amount, Period (days), MilestonePercentage (0-100) required
type CreateBidBody struct {
	ProjectID           int64  `json:"project_id"`
	BidderID            int64  `json:"bidder_id"`
	Amount              int    `json:"amount"`
	Period              int    `json:"period"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
	ProfileID           int    `json:"profile_id"`
}

func (s *createBid) Do(ctx context.Context, b CreateBidBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(ProjectsBids),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)

}

type actionBid struct {
	client *Client
	bidID  int
}

type ActionBidBody struct {
	Action BidAction `json:"action"`
}

func (s *actionBid) Do(ctx context.Context, b ActionBidBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(ProjectsBids), s.bidID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type updateBid struct {
	client *Client
	bidID  int
}

type UpdateBidBody struct {
	Amount              int    `json:"amount"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
}

func (s *updateBid) Do(ctx context.Context, b UpdateBidBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(ProjectsBids), s.bidID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: refine with typed response
type getTimeTracking struct {
	client                *Client
	bidID                 int
	fromTime              *int64
	toTime                *int64
	dailyAggregateDetails *bool
	invoiced              *bool
}

func (s *getTimeTracking) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/time_tracking_sessions", string(ProjectsBids), s.bidID),
	}

	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.dailyAggregateDetails != nil {
		r.addParam("daily_aggregate_details", *s.dailyAggregateDetails)
	}
	if s.invoiced != nil {
		r.addParam("invoiced", *s.invoiced)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetFromTime sets the starting Unix timestamp to filter time tracking sessions for a bid.
// Optional parameter: if not set, sessions will not be filtered by start
func (s *getTimeTracking) SetFromTime(val int64) *getTimeTracking {
	s.fromTime = &val
	return s
}

// SetToTime sets the ending Unix timestamp to filter time tracking sessions for a bid.
// Optional parameter: if not set, sessions will not be filtered by end time.
func (s *getTimeTracking) SetToTime(val int64) *getTimeTracking {
	s.toTime = &val
	return s
}

// SetDailyAggregateDetails sets whether to return aggregate daily time tracking sessions for a bid.
// Required parameter: must be set to receive aggregated session details.
func (s *getTimeTracking) SetDailyAggregateDetails(val bool) *getTimeTracking {
	s.dailyAggregateDetails = &val
	return s
}

// SetInvoiced sets whether to filter the time tracking sessions based on invoiced status.
// Optional parameter: if not set, both invoiced and non-invoiced sessions may be returned.
func (s *getTimeTracking) SetInvoiced(val bool) *getTimeTracking {
	s.invoiced = &val
	return s
}

type createTimeTracking struct {
	client *Client
	bidID  int
}

// StartTime and Seconds (Duration of session in seconds) required
type CreateTimeTrackingBidBody struct {
	StartTime int64  `json:"time_start"`
	Seconds   int    `json:"seconds"`
	Note      string `json:"note"`
}

func (s *createTimeTracking) Do(ctx context.Context, b CreateTimeTrackingBidBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%d/time_tracking_sessions", string(ProjectsBids), s.bidID),
		body:     bytes.NewBuffer(m),
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: refine with typed response
type getBidRatingsByListOfBids struct {
	client *Client
	bids   []int
}

func (s *getBidRatingsByListOfBids) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(ProjectsBidRatings),
	}

	for _, val := range s.bids {
		r.addParam("bids[]", val)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetBids filters results by specific bid IDs.
func (s *getBidRatingsByListOfBids) SetBids(vals []int) *getBidRatingsByListOfBids {
	s.bids = vals
	return s
}

// TODO: refine with typed response
type getBidRatings struct {
	client *Client
	bidID  int
}

func (s *getBidRatings) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings", string(ProjectsBids), s.bidID),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type createBidRating struct {
	client *Client
	bidID  int
}

// Rating required
type CreateBidRatingsBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func (s *createBidRating) Do(ctx context.Context, b CreateBidRatingsBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings", string(ProjectsBids), s.bidID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type updateBidRating struct {
	client      *Client
	bidID       int
	bidRatingID int
}

type UpdateBidRatingBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func (s *updateBidRating) Do(ctx context.Context, b UpdateBidRatingBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d/bid_ratings/%d", string(ProjectsBids), s.bidID, s.bidRatingID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: refine with typed response
type getBidEditRequest struct {
	client            *Client
	bidID             int
	statuses          []BidStatus
	bidEditRequestIDs []int
}

func (s *getBidEditRequest) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d/bid_edit_requests", string(ProjectsBids), s.bidID),
	}

	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
	}
	for _, val := range s.bidEditRequestIDs {
		r.addParam("bid_edit_request_ids[]", val)
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// SetStatuses sets the filter for bid edit request statuses.
// The statuses parameter can include values like "pending" or "accepted".
// Only bid edit requests matching the specified statuses will be returned.
func (s *getBidEditRequest) SetStatuses(vals []BidStatus) *getBidEditRequest {
	s.statuses = vals
	return s
}

// SetBidEditRequestIDs sets the filter for specific bid edit request IDs.
// Only bid edit requests with the specified IDs will be returned for the bid.
func (s *getBidEditRequest) SetBidEditRequestIDs(vals []int) *getBidEditRequest {
	s.bidEditRequestIDs = vals
	return s
}

type createBidEditRequest struct {
	client *Client
}

// BidID, NewAmount, and Period (days) are required
type CreateBidEditRequestsBody struct {
	BidID     int `json:"bid_id"`
	NewAmount int `json:"new_amount"`
	NewPeriod int `json:"new_period"`
	Comment   int `json:"comment"`
}

func (s *createBidEditRequest) Do(ctx context.Context, b CreateBidEditRequestsBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(ProjectsBidsBidEditRequests),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type actionBidEditRequest struct {
	client           *Client
	bidID            int
	bidEditRequestID int
}

type BidEditRequestsActionBody struct {
	Action BidEditRequestAction `json:"action"`
}

func (s *actionBidEditRequest) Do(ctx context.Context, b BidEditRequestsActionBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d/bid_edit_requests/%d", string(ProjectsBids), s.bidID, s.bidEditRequestID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

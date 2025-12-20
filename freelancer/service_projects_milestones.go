package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: refine with typed response
type listMilestones struct {
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

func (s *listMilestones) Do(ctx context.Context) (*RawResponse, error) {
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
	return execute[*RawResponse](ctx, s.client, r)
}

// SetBids sets the IDs of the bids to filter milestones by.
func (s *listMilestones) SetBids(vals []int) *listMilestones {
	s.bids = vals
	return s
}

// SetProjects sets the IDs of the projects to filter milestones by.
func (s *listMilestones) SetProjects(vals []int64) *listMilestones {
	s.projects = vals
	return s
}

// SetBidders sets the IDs of freelancers to filter milestones by.
func (s *listMilestones) SetBidders(vals []int64) *listMilestones {
	s.bidders = vals
	return s
}

// SetProjectOwners sets the IDs of employers to filter milestones by.
func (s *listMilestones) SetProjectOwners(vals []int64) *listMilestones {
	s.projectOwners = vals
	return s
}

// SetUsers sets the IDs of users to filter milestones by.
// Cannot be used together with SetBidders or SetProjectOwners.
func (s *listMilestones) SetUsers(vals []int64) *listMilestones {
	s.users = vals
	return s
}

// SetStatuses sets the milestone statuses to filter by.
func (s *listMilestones) SetStatuses(vals []MilestoneStatus) *listMilestones {
	s.statuses = vals
	return s
}

// SetSortField sets the field by which the milestones are sorted.
// Default is time_created.
func (s *listMilestones) SetSortField(val SortField) *listMilestones {
	s.sortField = &val
	return s
}

// SetSortDirection sets the sort direction for milestones: asc or desc.
// Default is desc.
func (s *listMilestones) SetSortDirection(val SortDirection) *listMilestones {
	s.sortDirection = &val
	return s
}

// SetExcludedMilestones specifies whether to exclude certain milestone IDs.
func (s *listMilestones) SetExcludedMilestones(val bool) *listMilestones {
	s.excludedMilestones = &val
	return s
}

// SetUserAvatar specifies whether to include the avatar of selected users.
func (s *listMilestones) SetUserAvatar(val bool) *listMilestones {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails specifies whether to include country details of selected users.
func (s *listMilestones) SetUserCountryDetails(val bool) *listMilestones {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription specifies whether to include the profile blurb of selected users.
func (s *listMilestones) SetUserProfileDescription(val bool) *listMilestones {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo specifies whether to include display names of selected users.
func (s *listMilestones) SetUserDisplayInfo(val bool) *listMilestones {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs specifies whether to include jobs of selected users.
func (s *listMilestones) SetUserJobs(val bool) *listMilestones {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails specifies whether to include currency balances of selected users.
func (s *listMilestones) SetUserBalanceDetails(val bool) *listMilestones {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails specifies whether to include completed qualification exams of selected users.
func (s *listMilestones) SetUserQualificationDetails(val bool) *listMilestones {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails specifies whether to include membership information of selected users.
func (s *listMilestones) SetUserMembershipDetails(val bool) *listMilestones {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails specifies whether to include financial information of selected users.
func (s *listMilestones) SetUserFinancialDetails(val bool) *listMilestones {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails specifies whether to include location information of selected users.
func (s *listMilestones) SetUserLocationDetails(val bool) *listMilestones {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails specifies whether to include portfolio information of selected users.
func (s *listMilestones) SetUserPortfolioDetails(val bool) *listMilestones {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails specifies whether to include preferred information of selected users.
func (s *listMilestones) SetUserPreferredDetails(val bool) *listMilestones {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails specifies whether to include badges earned by selected users.
func (s *listMilestones) SetUserBadgeDetails(val bool) *listMilestones {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus specifies whether to include additional status information of selected users.
func (s *listMilestones) SetUserStatus(val bool) *listMilestones {
	s.userStatus = &val
	return s
}

// SetUserReputation specifies whether to include freelancer reputation of selected users.
func (s *listMilestones) SetUserReputation(val bool) *listMilestones {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation specifies whether to include employer reputation of selected users.
func (s *listMilestones) SetUserEmployerReputation(val bool) *listMilestones {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra specifies whether to include full freelancer reputation of selected users.
func (s *listMilestones) SetUserReputationExtra(val bool) *listMilestones {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra specifies whether to include full employer reputation of selected users.
func (s *listMilestones) SetUserEmployerReputationExtra(val bool) *listMilestones {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage specifies whether to include profile pictures of selected users.
func (s *listMilestones) SetUserCoverImage(val bool) *listMilestones {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage specifies whether to include previous profile pictures of selected users.
func (s *listMilestones) SetUserPastCoverImage(val bool) *listMilestones {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations specifies whether to include recommendation count of selected users.
func (s *listMilestones) SetUserRecommendations(val bool) *listMilestones {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness specifies whether to include responsiveness score of selected users.
func (s *listMilestones) SetUserResponsiveness(val bool) *listMilestones {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers specifies whether to include corporate accounts created/founded by selected users.
func (s *listMilestones) SetCorporateUsers(val bool) *listMilestones {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber specifies whether to include the mobile number used by recruiters to contact selected users.
func (s *listMilestones) SetMarketingMobileNumber(val bool) *listMilestones {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails specifies whether to include the end timestamp of any sanction given to selected users.
func (s *listMilestones) SetSanctionDetails(val bool) *listMilestones {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount specifies whether to include limit account status of selected users.
func (s *listMilestones) SetLimitedAccount(val bool) *listMilestones {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails specifies whether to include equipment groups and items attached to selected users.
func (s *listMilestones) SetEquipmentGroupDetails(val bool) *listMilestones {
	s.equipmentGroupDetails = &val
	return s
}

type getMilestoneByID struct {
	client                      *Client
	milestoneID                 int
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

// TODO: refine with typed response
func (s *getMilestoneByID) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), s.milestoneID),
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
	return execute[*RawResponse](ctx, s.client, r)
}

// SetUserAvatar enables returning the avatar of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserAvatar(val bool) *getMilestoneByID {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails enables returning the country flag/code of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserCountryDetails(val bool) *getMilestoneByID {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription enables returning the profile blurb of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserProfileDescription(val bool) *getMilestoneByID {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo enables returning the display name of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserDisplayInfo(val bool) *getMilestoneByID {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs enables returning the jobs of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserJobs(val bool) *getMilestoneByID {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails enables returning the currency balance of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserBalanceDetails(val bool) *getMilestoneByID {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails enables returning the qualification exams completed by the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserQualificationDetails(val bool) *getMilestoneByID {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails enables returning the membership information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserMembershipDetails(val bool) *getMilestoneByID {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails enables returning the financial information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserFinancialDetails(val bool) *getMilestoneByID {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails enables returning the location information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserLocationDetails(val bool) *getMilestoneByID {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails enables returning the portfolio information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserPortfolioDetails(val bool) *getMilestoneByID {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails enables returning the preferred information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserPreferredDetails(val bool) *getMilestoneByID {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails enables returning the badges earned by the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserBadgeDetails(val bool) *getMilestoneByID {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus enables returning additional status information of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserStatus(val bool) *getMilestoneByID {
	s.userStatus = &val
	return s
}

// SetUserReputation enables returning the freelancer reputation of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserReputation(val bool) *getMilestoneByID {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation enables returning the employer reputation of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserEmployerReputation(val bool) *getMilestoneByID {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra enables returning the full freelancer reputation of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserReputationExtra(val bool) *getMilestoneByID {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra enables returning the full employer reputation of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserEmployerReputationExtra(val bool) *getMilestoneByID {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage enables returning the current profile picture of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserCoverImage(val bool) *getMilestoneByID {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage enables returning the previous profile pictures of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserPastCoverImage(val bool) *getMilestoneByID {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations enables returning the recommendations count of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserRecommendations(val bool) *getMilestoneByID {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness enables returning the responsiveness score(s) of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetUserResponsiveness(val bool) *getMilestoneByID {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers enables returning the corporate accounts created/founded by the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetCorporateUsers(val bool) *getMilestoneByID {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber enables returning the mobile number of the user used by recruiters in the milestone response.
func (s *getMilestoneByID) SetMarketingMobileNumber(val bool) *getMilestoneByID {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables returning the end timestamp of any sanctions applied to the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetSanctionDetails(val bool) *getMilestoneByID {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables returning the limited account status of the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetLimitedAccount(val bool) *getMilestoneByID {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails enables returning the equipment groups and items attached to the selected user(s) in the milestone response.
func (s *getMilestoneByID) SetEquipmentGroupDetails(val bool) *getMilestoneByID {
	s.equipmentGroupDetails = &val
	return s
}

type createMilestone struct {
	client *Client
}

type CreateMilestoneBody struct {
	ProjectID   int64                 `json:"project_id"`
	BidderID    int64                 `json:"bidder_id"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneCreateReason `json:"reason"`
	Description string                `json:"description"`
}

func (s *createMilestone) Do(ctx context.Context, b CreateMilestoneBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_MILESTONES),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type actionMilestone struct {
	client      *Client
	milestoneID int
}

type ActionMilestoneBody struct {
	Action      MilestoneAction       `json:"action"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneActionReason `json:"reason"`
	ReasonText  string                `json:"reason_text"`
	OtherReason string                `json:"other_reason"`
}

func (s *actionMilestone) Do(ctx context.Context, b ActionMilestoneBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), s.milestoneID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type listMilestoneRequests struct {
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

// TODO: refine with typed response
func (s *listMilestoneRequests) Do(ctx context.Context) (*RawResponse, error) {
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
	return execute[*RawResponse](ctx, s.client, r)
}

// SetMilestoneRequests sets the milestone request IDs to filter the results.
func (s *listMilestoneRequests) SetMilestoneRequests(vals []int) *listMilestoneRequests {
	s.milestoneRequests = vals
	return s
}

// SetBids sets the bid IDs to filter the milestone requests by.
func (s *listMilestoneRequests) SetBids(vals []int) *listMilestoneRequests {
	s.bids = vals
	return s
}

// SetProjects sets the project IDs to filter the milestone requests by.
func (s *listMilestoneRequests) SetProjects(vals []int64) *listMilestoneRequests {
	s.projects = vals
	return s
}

// SetBidders sets the freelancer user IDs to filter the milestone requests by.
func (s *listMilestoneRequests) SetBidders(vals []int64) *listMilestoneRequests {
	s.bidders = vals
	return s
}

// SetProjectOwners sets the employer user IDs to filter the milestone requests by.
func (s *listMilestoneRequests) SetProjectOwners(vals []int64) *listMilestoneRequests {
	s.projectOwners = vals
	return s
}

// SetUsers sets the user IDs to filter the milestone requests by. Cannot be combined with SetBidders or SetProjectOwners.
func (s *listMilestoneRequests) SetUsers(vals []int64) *listMilestoneRequests {
	s.users = vals
	return s
}

// SetStatuses sets the milestone request statuses to filter by (e.g., pending, created, rejected, deleted).
func (s *listMilestoneRequests) SetStatuses(vals []MilestoneStatus) *listMilestoneRequests {
	s.statuses = vals
	return s
}

// SetFromTime sets the timestamp to filter milestone requests updated after this time (inclusive).
func (s *listMilestoneRequests) SetFromTime(val int64) *listMilestoneRequests {
	s.fromTime = &val
	return s
}

// SetToTime sets the timestamp to filter milestone requests updated before this time (inclusive).
func (s *listMilestoneRequests) SetToTime(val int64) *listMilestoneRequests {
	s.toTime = &val
	return s
}

// SetSortField sets the field by which to sort the milestone requests.
func (s *listMilestoneRequests) SetSortField(val SortField) *listMilestoneRequests {
	s.sortField = &val
	return s
}

// SetSortDirection sets the sorting direction (ascending or descending) for the milestone requests.
func (s *listMilestoneRequests) SetSortDirection(val SortDirection) *listMilestoneRequests {
	s.sortDirection = &val
	return s
}

// SetExcludedMilestones sets whether to exclude milestone details from the results.
func (s *listMilestoneRequests) SetExcludedMilestones(val bool) *listMilestoneRequests {
	s.excludedMilestones = &val
	return s
}

// SetUserAvatar sets whether to include the user's avatar in the response.
func (s *listMilestoneRequests) SetUserAvatar(val bool) *listMilestoneRequests {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether to include the user's country details (flag/code) in the response.
func (s *listMilestoneRequests) SetUserCountryDetails(val bool) *listMilestoneRequests {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether to include the user's profile description.
func (s *listMilestoneRequests) SetUserProfileDescription(val bool) *listMilestoneRequests {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo sets whether to include the user's display name.
func (s *listMilestoneRequests) SetUserDisplayInfo(val bool) *listMilestoneRequests {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether to include the user's jobs information.
func (s *listMilestoneRequests) SetUserJobs(val bool) *listMilestoneRequests {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether to include the user's currency balance.
func (s *listMilestoneRequests) SetUserBalanceDetails(val bool) *listMilestoneRequests {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether to include the user's completed qualification exams.
func (s *listMilestoneRequests) SetUserQualificationDetails(val bool) *listMilestoneRequests {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether to include the user's membership information.
func (s *listMilestoneRequests) SetUserMembershipDetails(val bool) *listMilestoneRequests {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether to include the user's financial information.
func (s *listMilestoneRequests) SetUserFinancialDetails(val bool) *listMilestoneRequests {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether to include the user's location information.
func (s *listMilestoneRequests) SetUserLocationDetails(val bool) *listMilestoneRequests {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether to include the user's portfolio information.
func (s *listMilestoneRequests) SetUserPortfolioDetails(val bool) *listMilestoneRequests {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether to include the user's preferred details.
func (s *listMilestoneRequests) SetUserPreferredDetails(val bool) *listMilestoneRequests {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether to include the user's badge details.
func (s *listMilestoneRequests) SetUserBadgeDetails(val bool) *listMilestoneRequests {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether to include additional status information about the user.
func (s *listMilestoneRequests) SetUserStatus(val bool) *listMilestoneRequests {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether to include the user's freelancer reputation.
func (s *listMilestoneRequests) SetUserReputation(val bool) *listMilestoneRequests {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether to include the user's employer reputation.
func (s *listMilestoneRequests) SetUserEmployerReputation(val bool) *listMilestoneRequests {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether to include the full freelancer reputation of the user.
func (s *listMilestoneRequests) SetUserReputationExtra(val bool) *listMilestoneRequests {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether to include the full employer reputation of the user.
func (s *listMilestoneRequests) SetUserEmployerReputationExtra(val bool) *listMilestoneRequests {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether to include the user's current profile picture.
func (s *listMilestoneRequests) SetUserCoverImage(val bool) *listMilestoneRequests {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether to include the user's past profile pictures.
func (s *listMilestoneRequests) SetUserPastCoverImage(val bool) *listMilestoneRequests {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether to include the count of user recommendations.
func (s *listMilestoneRequests) SetUserRecommendations(val bool) *listMilestoneRequests {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness sets whether to include the user's responsiveness score.
func (s *listMilestoneRequests) SetUserResponsiveness(val bool) *listMilestoneRequests {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers sets whether to include the corporate accounts created/founded by the user.
func (s *listMilestoneRequests) SetCorporateUsers(val bool) *listMilestoneRequests {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber sets whether to include the mobile number used by recruiters to contact the user.
func (s *listMilestoneRequests) SetMarketingMobileNumber(val bool) *listMilestoneRequests {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to include the sanction end timestamp for the user.
func (s *listMilestoneRequests) SetSanctionDetails(val bool) *listMilestoneRequests {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to include the user's limited account status.
func (s *listMilestoneRequests) SetLimitedAccount(val bool) *listMilestoneRequests {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to include the equipment groups and items attached to the user.
func (s *listMilestoneRequests) SetEquipmentGroupDetails(val bool) *listMilestoneRequests {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of milestone requests to return.
func (s *listMilestoneRequests) SetLimit(val int) *listMilestoneRequests {
	s.limit = &val
	return s
}

// SetOffset sets the pagination offset for milestone requests.
func (s *listMilestoneRequests) SetOffset(val int) *listMilestoneRequests {
	s.offset = &val
	return s
}

type getMilestoneRequestByID struct {
	client                      *Client
	milestoneRequestID          int
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

// TODO: refine with typed response
func (s *getMilestoneRequestByID) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONE_REQUESTS), s.milestoneRequestID),
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
	return execute[*RawResponse](ctx, s.client, r)
}

// SetUserAvatar sets whether the response should include the avatar of the selected user.
func (s *getMilestoneRequestByID) SetUserAvatar(val bool) *getMilestoneRequestByID {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether the response should include the country flag/code of the selected user.
func (s *getMilestoneRequestByID) SetUserCountryDetails(val bool) *getMilestoneRequestByID {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether the response should include the profile description of the selected user.
func (s *getMilestoneRequestByID) SetUserProfileDescription(val bool) *getMilestoneRequestByID {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo sets whether the response should include the display name of the selected user.
func (s *getMilestoneRequestByID) SetUserDisplayInfo(val bool) *getMilestoneRequestByID {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether the response should include the jobs of the selected user.
func (s *getMilestoneRequestByID) SetUserJobs(val bool) *getMilestoneRequestByID {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether the response should include the currency balance of the selected user.
func (s *getMilestoneRequestByID) SetUserBalanceDetails(val bool) *getMilestoneRequestByID {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether the response should include the qualification exams completed by the user.
func (s *getMilestoneRequestByID) SetUserQualificationDetails(val bool) *getMilestoneRequestByID {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether the response should include the membership information of the user.
func (s *getMilestoneRequestByID) SetUserMembershipDetails(val bool) *getMilestoneRequestByID {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether the response should include the financial information of the user.
func (s *getMilestoneRequestByID) SetUserFinancialDetails(val bool) *getMilestoneRequestByID {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether the response should include the location information of the user.
func (s *getMilestoneRequestByID) SetUserLocationDetails(val bool) *getMilestoneRequestByID {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether the response should include the portfolio information of the user.
func (s *getMilestoneRequestByID) SetUserPortfolioDetails(val bool) *getMilestoneRequestByID {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether the response should include the preferred information of the user.
func (s *getMilestoneRequestByID) SetUserPreferredDetails(val bool) *getMilestoneRequestByID {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether the response should include the badges earned by the user.
func (s *getMilestoneRequestByID) SetUserBadgeDetails(val bool) *getMilestoneRequestByID {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether the response should include additional status information of the user.
func (s *getMilestoneRequestByID) SetUserStatus(val bool) *getMilestoneRequestByID {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether the response should include the freelancer reputation of the user.
func (s *getMilestoneRequestByID) SetUserReputation(val bool) *getMilestoneRequestByID {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether the response should include the employer reputation of the user.
func (s *getMilestoneRequestByID) SetUserEmployerReputation(val bool) *getMilestoneRequestByID {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether the response should include the full freelancer reputation of the user.
func (s *getMilestoneRequestByID) SetUserReputationExtra(val bool) *getMilestoneRequestByID {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether the response should include the full employer reputation of the user.
func (s *getMilestoneRequestByID) SetUserEmployerReputationExtra(val bool) *getMilestoneRequestByID {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether the response should include the profile picture of the user.
func (s *getMilestoneRequestByID) SetUserCoverImage(val bool) *getMilestoneRequestByID {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether the response should include previous profile pictures of the user.
func (s *getMilestoneRequestByID) SetUserPastCoverImage(val bool) *getMilestoneRequestByID {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether the response should include the recommendations count of the user.
func (s *getMilestoneRequestByID) SetUserRecommendations(val bool) *getMilestoneRequestByID {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness sets whether the response should include the responsiveness score(s) of the user.
func (s *getMilestoneRequestByID) SetUserResponsiveness(val bool) *getMilestoneRequestByID {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers sets whether the response should include corporate accounts created/founded by the user.
func (s *getMilestoneRequestByID) SetCorporateUsers(val bool) *getMilestoneRequestByID {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber sets whether the response should include the mobile number used by recruiters to contact the user.
func (s *getMilestoneRequestByID) SetMarketingMobileNumber(val bool) *getMilestoneRequestByID {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether the response should include the end timestamp of sanctions given to the user.
func (s *getMilestoneRequestByID) SetSanctionDetails(val bool) *getMilestoneRequestByID {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether the response should include the limited account status of the user.
func (s *getMilestoneRequestByID) SetLimitedAccount(val bool) *getMilestoneRequestByID {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether the response should include the equipment groups and items attached to the user.
func (s *getMilestoneRequestByID) SetEquipmentGroupDetails(val bool) *getMilestoneRequestByID {
	s.equipmentGroupDetails = &val
	return s
}

type createMilestoneRequest struct {
	client *Client
}

// ProjectID, BidID, Amount and Description are required
type CreateMilestoneRequestBody struct {
	ProjectID   int64  `json:"project_id"`
	BidID       int    `json:"bid_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func (s *createMilestoneRequest) Do(ctx context.Context, b CreateMilestoneBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_MILESTONE_REQUESTS),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type actionMilestoneRequest struct {
	client             *Client
	milestoneRequestID int
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

func (s *actionMilestoneRequest) Do(ctx context.Context, b ActionMilestoneRequestBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), s.milestoneRequestID),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

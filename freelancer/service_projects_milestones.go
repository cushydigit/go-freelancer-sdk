package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type ListMilestonesOptions struct {
	Projects                    []int64
	ProjectOwners               []int64
	Bidders                     []int64
	Users                       []int64
	Bids                        []int64
	Statuses                    []MilestoneStatus
	SortField                   *SortField
	SortDirection               *SortDirection
	ExcludedMilestones          *bool
	UserAvatar                  *bool
	UserCountryDetails          *bool
	UserProfileDescription      *bool
	UserDisplayInfo             *bool
	UserJobs                    *bool
	UserBalanceDetails          *bool
	UserQualificationDetails    *bool
	UserMembershipDetails       *bool
	UserFinancialDetails        *bool
	UserLocationDetails         *bool
	UserPortfolioDetails        *bool
	UserPreferredDetails        *bool
	UserBadgeDetails            *bool
	UserStatus                  *bool
	UserReputation              *bool
	UserEmployerReputation      *bool
	UserReputationExtra         *bool
	UserEmployerReputationExtra *bool
	UserCoverImage              *bool
	UserPastCoverImage          *bool
	UserRecommendations         *bool
	UserResponsiveness          *bool
	CorporateUsers              *bool
	MarketingMobileNumber       *bool
	SanctionDetails             *bool
	LimitedAccount              *bool
	EquipmentGroupDetails       *bool
}

// TODO: refine with typed response

// Returns a list of milestones. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/milestones` endpoint
func (s *MilestonesService) List(ctx context.Context, opts *ListMilestonesOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Projects {
			addInt64(query, "projects[]", &val)
		}
		for _, val := range opts.ProjectOwners {
			addInt64(query, "project_owners[]", &val)
		}
		for _, val := range opts.Bidders {
			addInt64(query, "bidders[]", &val)
		}
		for _, val := range opts.Users {
			addInt64(query, "users[]", &val)
		}
		for _, val := range opts.Bids {
			addInt64(query, "bids[]", &val)
		}
		for _, val := range opts.Statuses {
			addEnum(query, "statuses[]", &val)
		}
		addEnum(query, "sort_field", opts.SortField)
		addEnum(query, "sort_direction", opts.SortDirection)
		addBool(query, "excluded_milestones", opts.ExcludedMilestones)
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_Description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)
		addBool(query, "user_balance_details", opts.UserBalanceDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
		addBool(query, "user_past_cover_image", opts.UserPastCoverImage)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "user_responsiveness", opts.UserResponsiveness)
		addBool(query, "corporate_users", opts.CorporateUsers)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsMilestones), query, nil)
}

type GetMilestoneOptions struct {
	MilestoneID                 int
	UserAvatar                  *bool
	UserCountryDetails          *bool
	UserProfileDescription      *bool
	UserDisplayInfo             *bool
	UserJobs                    *bool
	UserBalanceDetails          *bool
	UserQualificationDetails    *bool
	UserMembershipDetails       *bool
	UserFinancialDetails        *bool
	UserLocationDetails         *bool
	UserPortfolioDetails        *bool
	UserPreferredDetails        *bool
	UserBadgeDetails            *bool
	UserStatus                  *bool
	UserReputation              *bool
	UserEmployerReputation      *bool
	UserReputationExtra         *bool
	UserEmployerReputationExtra *bool
	UserCoverImage              *bool
	UserPastCoverImage          *bool
	UserRecommendations         *bool
	UserResponsiveness          *bool
	CorporateUsers              *bool
	MarketingMobileNumber       *bool
	SanctionDetails             *bool
	LimitedAccount              *bool
	EquipmentGroupDetails       *bool
}

// TODO: refine with typed response

// Returns information about a specific milestone.
// It maps to the `GET` `/projects/0.1/milestones/{milestone_id}` endpoint
func (s *MilestonesService) GetByID(ctx context.Context, milestoneID int, opts *GetMilestoneOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", ProjectsMilestones, milestoneID)
	query := url.Values{}
	if opts != nil {
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_Description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)
		addBool(query, "user_balance_details", opts.UserBalanceDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
		addBool(query, "user_past_cover_image", opts.UserPastCoverImage)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "user_responsiveness", opts.UserResponsiveness)
		addBool(query, "corporate_users", opts.CorporateUsers)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)

	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type CreateMilestoneBody struct {
	ProjectID   int64                 `json:"project_id"`
	BidderID    int64                 `json:"bidder_id"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneCreateReason `json:"reason"`
	Description string                `json:"description"`
}

// Post a review of a user.
// It maps to the `POST` `/projects/0.1/milestones` endpoint
func (s *MilestonesService) Create(ctx context.Context, b CreateMilestoneBody) (*RawResponse, error) {
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(ProjectsMilestones), nil, b)
}

type ActionMilestoneBody struct {
	Action      MilestoneAction       `json:"action"`
	Amount      int                   `json:"amount"`
	Reason      MilestoneActionReason `json:"reason"`
	ReasonText  string                `json:"reason_text"`
	OtherReason string                `json:"other_reason"`
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (s *MilestonesService) Action(ctx context.Context, milestoneID int, b ActionMilestoneBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", ProjectsMilestones, milestoneID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

type ListMilestonesRequestsOptions struct {
	MilestoneRequests           []int64
	Projects                    []int64
	ProjectOwners               []int64
	Bidders                     []int64
	Users                       []int64
	Bids                        []int
	Statuses                    []MilestoneStatus
	FromTime                    *int64
	ToTime                      *int64
	SortField                   *SortField
	SortDirection               *SortDirection
	ExcludedMilestones          *bool
	UserAvatar                  *bool
	UserCountryDetails          *bool
	UserProfileDescription      *bool
	UserDisplayInfo             *bool
	UserJobs                    *bool
	UserBalanceDetails          *bool
	UserQualificationDetails    *bool
	UserMembershipDetails       *bool
	UserFinancialDetails        *bool
	UserLocationDetails         *bool
	UserPortfolioDetails        *bool
	UserPreferredDetails        *bool
	UserBadgeDetails            *bool
	UserStatus                  *bool
	UserReputation              *bool
	UserEmployerReputation      *bool
	UserReputationExtra         *bool
	UserEmployerReputationExtra *bool
	UserCoverImage              *bool
	UserPastCoverImage          *bool
	UserRecommendations         *bool
	UserResponsiveness          *bool
	CorporateUsers              *bool
	MarketingMobileNumber       *bool
	SanctionDetails             *bool
	LimitedAccount              *bool
	EquipmentGroupDetails       *bool
	Limit                       *int
	Offset                      *int
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/milestone_requests` endpoint
func (s *MilestoneRequestsService) List(ctx context.Context, opts *ListMilestonesRequestsOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.MilestoneRequests {
			addInt64(query, "milestone_requests[]", &val)
		}
		for _, val := range opts.Projects {
			addInt64(query, "projects[]", &val)
		}
		for _, val := range opts.ProjectOwners {
			addInt64(query, "project_owners[]", &val)
		}
		for _, val := range opts.Bidders {
			addInt64(query, "bidders[]", &val)
		}
		for _, val := range opts.Users {
			addInt64(query, "users[]", &val)
		}
		for _, val := range opts.Bids {
			addInt(query, "bids[]", &val)
		}
		for _, val := range opts.Statuses {
			addEnum(query, "statuses[]", &val)
		}
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addEnum(query, "sort_field", opts.SortField)
		addEnum(query, "sort_direction", opts.SortDirection)
		addBool(query, "excluded_milestones", opts.ExcludedMilestones)
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)
		addBool(query, "user_balance_details", opts.UserBalanceDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
		addBool(query, "user_past_cover_image", opts.UserPastCoverImage)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "user_responsiveness", opts.UserResponsiveness)
		addBool(query, "corporate_users", opts.CorporateUsers)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)

	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsMilestoneRequests), query, nil)
}

type GetMilestoneRequestOptions struct {
	MilestoneRequestID          int64
	UserAvatar                  *bool
	UserCountryDetails          *bool
	UserProfileDescription      *bool
	UserDisplayInfo             *bool
	UserJobs                    *bool
	UserBalanceDetails          *bool
	UserQualificationDetails    *bool
	UserMembershipDetails       *bool
	UserFinancialDetails        *bool
	UserLocationDetails         *bool
	UserPortfolioDetails        *bool
	UserPreferredDetails        *bool
	UserBadgeDetails            *bool
	UserStatus                  *bool
	UserReputation              *bool
	UserEmployerReputation      *bool
	UserReputationExtra         *bool
	UserEmployerReputationExtra *bool
	UserCoverImage              *bool
	UserPastCoverImage          *bool
	UserRecommendations         *bool
	UserResponsiveness          *bool
	CorporateUsers              *bool
	MarketingMobileNumber       *bool
	SanctionDetails             *bool
	LimitedAccount              *bool
	EquipmentGroupDetails       *bool
}

// TODO: refine with typed response

// Returns information about a specific milestone request.
// It maps to the `GET` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Get(ctx context.Context, milestoneRequestID int, opts *GetMilestoneRequestOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", ProjectsMilestoneRequests, milestoneRequestID)
	query := url.Values{}
	if opts != nil {
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_Description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)
		addBool(query, "user_balance_details", opts.UserBalanceDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
		addBool(query, "user_past_cover_image", opts.UserPastCoverImage)
		addBool(query, "user_recommendations", opts.UserRecommendations)
		addBool(query, "user_responsiveness", opts.UserResponsiveness)
		addBool(query, "corporate_users", opts.CorporateUsers)
		addBool(query, "marketing_mobile_number", opts.MarketingMobileNumber)
		addBool(query, "sanction_details", opts.SanctionDetails)
		addBool(query, "limited_account", opts.LimitedAccount)
		addBool(query, "equipment_group_details", opts.EquipmentGroupDetails)

	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

// ProjectID, BidID, Amount and Description are required
type CreateMilestoneRequestBody struct {
	ProjectID   int64  `json:"project_id"`
	BidID       int    `json:"bid_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

// Creates a milestone request from a given JSON object.
// It maps to the `POST` `/projects/0.1/milestone_requests` endpoint
func (s *MilestoneRequestsService) Create(ctx context.Context, b CreateMilestoneRequestBody) (*RawResponse, error) {
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(ProjectsMilestoneRequests), nil, b)
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Action(ctx context.Context, milestoneRequestID int, b ActionMilestoneRequestBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", ProjectsMilestoneRequests, milestoneRequestID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

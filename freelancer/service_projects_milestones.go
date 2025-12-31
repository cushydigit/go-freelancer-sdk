package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type ListMilestonesOptions struct {
	Projects                    []int64           `url:"projects[]"`
	ProjectOwners               []int64           `url:"project_owners[]"`
	Bidders                     []int64           `url:"bidders[]"`
	Users                       []int64           `url:"users[]"`
	Bids                        []int64           `url:"bids[]"`
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	SortField                   *SortField        `url:"sort_field"`
	SortDirection               *SortDirection    `url:"sort_direction"`
	ExcludedMilestones          *bool             `url:"excluded_milestones"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
}

// TODO: refine with typed response

// Returns a list of milestones. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/milestones` endpoint
func (s *MilestonesService) List(ctx context.Context, opts *ListMilestonesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsMilestones
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type GetMilestoneOptions struct {
	MilestoneID                 int   `url:"milestone_id"`
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
}

// TODO: refine with typed response

// Returns information about a specific milestone.
// It maps to the `GET` `/projects/0.1/milestones/{milestone_id}` endpoint
func (s *MilestonesService) GetByID(ctx context.Context, milestoneID int, opts *GetMilestoneOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
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
	p := endpoints.ProjectsMilestones
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
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
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

type ListMilestonesRequestsOptions struct {
	MilestoneRequests           []int64           `url:"milestone_requests[]"`
	Projects                    []int64           `url:"projects[]"`
	ProjectOwners               []int64           `url:"project_owners[]"`
	Bidders                     []int64           `url:"bidders[]"`
	Users                       []int64           `url:"users[]"`
	Bids                        []int64           `url:"bids[]"`
	Statuses                    []MilestoneStatus `url:"statuses[]"`
	FromTime                    *int64            `url:"from_time"`
	ToTime                      *int64            `url:"to_time"`
	SortField                   *SortField        `url:"sort_field"`
	SortDirection               *SortDirection    `url:"sort_direction"`
	ExcludedMilestones          *bool             `url:"excluded_milestones"`
	UserAvatar                  *bool             `url:"user_avatar"`
	UserCountryDetails          *bool             `url:"user_country_details"`
	UserProfileDescription      *bool             `url:"user_profile_description"`
	UserDisplayInfo             *bool             `url:"user_display_info"`
	UserJobs                    *bool             `url:"user_jobs"`
	UserBalanceDetails          *bool             `url:"user_balance_details"`
	UserQualificationDetails    *bool             `url:"user_qualification_details"`
	UserMembershipDetails       *bool             `url:"user_membership_details"`
	UserFinancialDetails        *bool             `url:"user_financial_details"`
	UserLocationDetails         *bool             `url:"user_location_details"`
	UserPortfolioDetails        *bool             `url:"user_portfolio_details"`
	UserPreferredDetails        *bool             `url:"user_preferred_details"`
	UserBadgeDetails            *bool             `url:"user_badge_details"`
	UserStatus                  *bool             `url:"user_status"`
	UserReputation              *bool             `url:"user_reputation"`
	UserEmployerReputation      *bool             `url:"user_employer_reputation"`
	UserReputationExtra         *bool             `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool             `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool             `url:"user_cover_image"`
	UserPastCoverImage          *bool             `url:"user_past_cover_image"`
	UserRecommendations         *bool             `url:"user_recommendations"`
	UserResponsiveness          *bool             `url:"user_responsiveness"`
	CorporateUsers              *bool             `url:"corporate_users"`
	MarketingMobileNumber       *bool             `url:"marketing_mobile_number"`
	SanctionDetails             *bool             `url:"sanction_details"`
	LimitedAccount              *bool             `url:"limited_account"`
	EquipmentGroupDetails       *bool             `url:"equipment_group_details"`
	Limit                       *int              `url:"limit"`
	Offset                      *int              `url:"offset"`
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/milestone_requests` endpoint
func (s *MilestoneRequestsService) List(ctx context.Context, opts *ListMilestonesRequestsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsMilestoneRequests
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type GetMilestoneRequestOptions struct {
	UserAvatar                  *bool `url:"user_avatar"`
	UserCountryDetails          *bool `url:"user_country_details"`
	UserProfileDescription      *bool `url:"user_profile_description"`
	UserDisplayInfo             *bool `url:"user_display_info"`
	UserJobs                    *bool `url:"user_jobs"`
	UserBalanceDetails          *bool `url:"user_balance_details"`
	UserQualificationDetails    *bool `url:"user_qualification_details"`
	UserMembershipDetails       *bool `url:"user_membership_details"`
	UserFinancialDetails        *bool `url:"user_financial_details"`
	UserLocationDetails         *bool `url:"user_location_details"`
	UserPortfolioDetails        *bool `url:"user_portfolio_details"`
	UserPreferredDetails        *bool `url:"user_preferred_details"`
	UserBadgeDetails            *bool `url:"user_badge_details"`
	UserStatus                  *bool `url:"user_status"`
	UserReputation              *bool `url:"user_reputation"`
	UserEmployerReputation      *bool `url:"user_employer_reputation"`
	UserReputationExtra         *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool `url:"user_cover_image"`
	UserPastCoverImage          *bool `url:"user_past_cover_image"`
	UserRecommendations         *bool `url:"user_recommendations"`
	UserResponsiveness          *bool `url:"user_responsiveness"`
	CorporateUsers              *bool `url:"corporate_users"`
	MarketingMobileNumber       *bool `url:"marketing_mobile_number"`
	SanctionDetails             *bool `url:"sanction_details"`
	LimitedAccount              *bool `url:"limited_account"`
	EquipmentGroupDetails       *bool `url:"equipment_group_details"`
}

// TODO: refine with typed response

// Returns information about a specific milestone request.
// It maps to the `GET` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Get(ctx context.Context, milestoneRequestID int, opts *GetMilestoneRequestOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
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
	p := endpoints.ProjectsMilestoneRequests
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ActionMilestoneRequestBody struct {
	Action MilestoneActionRequest `json:"action"`
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Action(ctx context.Context, milestoneRequestID int, b ActionMilestoneRequestBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

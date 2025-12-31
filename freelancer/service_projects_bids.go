package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type ListBidsOptions struct {
	Bids                        []int64             `url:"bids[]"`
	Projects                    []int64             `url:"projects[]"`
	Bidders                     []int64             `url:"bidders[]"`
	ProjectOwners               []int64             `url:"project_owners[]"`
	AwardStatuses               []BidAwardStatus    `url:"award_statuses[]"`
	PaidStatuses                []BidPaidStatus     `url:"paid_statuses[]"`
	CompleteStatuses            []BidCompleteStatus `url:"complete_statuses[]"`
	FrontBidStatuses            []BidFrontendStatus `url:"front_bid_statuses[]"`
	FromTime                    *int64              `url:"from_time"`
	ToTime                      *int64              `url:"to_time"`
	Reputation                  *bool               `url:"reputation"`
	BuyerProjectFee             *bool               `url:"buyer_project_fee"`
	AwardStatusPossibilities    *bool               `url:"award_status_possibilities"`
	ProjectDetails              *bool               `url:"project_details"`
	UserDetails                 *bool               `url:"user_details"`
	UserAvatar                  *bool               `url:"user_avatar"`
	UserCountryDetails          *bool               `url:"user_country_details"`
	UserProfileDescription      *bool               `url:"user_profile_description"`
	UserDisplayInfo             *bool               `url:"user_display_info"`
	UserJobs                    *bool               `url:"user_jobs"`
	UserBalanceDetails          *bool               `url:"user_balance_details"`
	UserQualificationDetails    *bool               `url:"user_qualification_details"`
	UserMembershipDetails       *bool               `url:"user_membership_details"`
	UserFinancialDetails        *bool               `url:"user_financial_details"`
	UserLocationDetails         *bool               `url:"user_location_details"`
	UserPortfolioDetails        *bool               `url:"user_portfolio_details"`
	UserPreferredDetails        *bool               `url:"user_preferred_details"`
	UserBadgeDetails            *bool               `url:"user_badge_details"`
	UserStatus                  *bool               `url:"user_status"`
	UserReputation              *bool               `url:"user_reputation"`
	UserEmployerReputation      *bool               `url:"user_employer_reputation"`
	UserReputationExtra         *bool               `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool               `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool               `url:"user_cover_image"`
	UserPastCoverImage          *bool               `url:"user_past_cover_image"`
	UserRecommendations         *bool               `url:"user_recommendations"`
	UserResponsiveness          *bool               `url:"user_responsiveness"`
	CorporateUsers              *bool               `url:"corporate_users"`
	MarketingMobileNumber       *bool               `url:"marketing_mobile_number"`
	SanctionDetails             *bool               `url:"sanction_details"`
	LimitedAccount              *bool               `url:"limited_account"`
	EquipmentGroupDetails       *bool               `url:"equipment_group_details"`
	Limit                       *int                `url:"limit"`
	Offset                      *int                `url:"offset"`
	Compact                     *bool               `url:"compact"`
}

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// It maps to the `GET` `/projects/0.1/bids` endpoint
func (s *BidsService) List(ctx context.Context, opts *ListBidsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsBids
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type GetBidOptions struct {
	Reputation                  *bool `url:"reputation"`
	BuyerProjectFee             *bool `url:"buyer_project_fee"`
	AwardStatusPossibilities    *bool `url:"award_status_possibilities"`
	ProjectDetails              *bool `url:"project_details"`
	UserDetails                 *bool `url:"user_details"`
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
	Limit                       *int  `url:"limit"`
	Offset                      *int  `url:"offset"`
	Compact                     *bool `url:"compact"`
	Quotations                  *bool `url:"quotations"`
}

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// it maps to the `GET` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Get(ctx context.Context, bidID int64, opts *GetBidOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
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

// Creates a bid on a project. Accepts a JSON object in the style described in the Bid struct (with enums as strings, and objects as dictionaries).
// It maps to the `POST` `/projects/0.1/bids` endpoint
func (s *BidsService) Create(ctx context.Context, b CreateBidBody) (*RawResponse, error) {
	p := endpoints.ProjectsBids
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ActionBidBody struct {
	Action BidAction `json:"action"`
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Action(ctx context.Context, bidID int64, b *ActionBidBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

type UpdateBidBody struct {
	Amount              int    `json:"amount"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
}

// Updates an existing bid on a project. An existing bids information (description,amount,milestone_percentage) can be updated by sending a JSON encoded Bid struct.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Update(ctx context.Context, bidID int64, b UpdateBidBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

type GetTimeTrackingOptions struct {
	FromTime              *int64 `url:"from_time"`
	ToTime                *int64 `url:"to_time"`
	DailyAggregateDetails *bool  `url:"daily_aggregate_details"`
	Invoiced              *bool  `url:"invoiced"`
}

// TODO: refine with typed response

// Returns a list of aggregate time tracking data for a bid.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (s *BidsService) GetTimeTracking(ctx context.Context, bidID int64, opts *GetTimeTrackingOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// StartTime and Seconds (Duration of session in seconds) required
type CreateTimeTrackingBody struct {
	StartTime int64  `json:"time_start"`
	Seconds   int    `json:"seconds"`
	Note      string `json:"note"`
}

// Creates a time tracking session for a specific bid.
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (s *BidsService) CreateTimeTracking(ctx context.Context, bidID int64, b CreateTimeTrackingBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ListBidEditRequestsOptions struct {
	Statuses          []BidStatus `url:"statuses[]"`
	BidEditRequestIDs []int64     `url:"bid_edit_request_ids[]"`
}

// Return bid edit requests by bid id.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/edit_requests` endpoint
// the original name was Get but it was renamed to List due to returning list of edit requests
func (s *BidEditRequestService) List(ctx context.Context, bidID int64, opts *ListBidEditRequestsOptions) (*ListBidEditRequestsResponse, error) {
	p := fmt.Sprintf("%s/%d/edit_requests/", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*ListBidEditRequestsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Required: BidID, NewAmount, NewPeriod
type CreateBidEditRequestsBody struct {
	BidID     int64 `json:"bid_id"`
	NewAmount int   `json:"new_amount"`
	NewPeriod int   `json:"new_period"`
	Comment   int   `json:"comment"`
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request.
// It maps to the `POST` `/projects/0.1/bids/edit_requests` endpoint
func (s *BidEditRequestService) Create(ctx context.Context, b CreateBidEditRequestsBody) (*CreateBidEditRequestResponse, error) {
	p := endpoints.ProjectsBidsBidEditRequests
	return execute[*CreateBidEditRequestResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ActionBidEditRequestsBody struct {
	Action BidEditRequestAction `json:"action"`
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (s *BidEditRequestService) Action(ctx context.Context, bidID, bidEditRequestID int64, b ActionBidEditRequestsBody) (*ActionBidEditRequestResponse, error) {
	p := fmt.Sprintf("%s/%d/edit_requests/%d", endpoints.ProjectsBids, bidID, bidEditRequestID)
	return execute[*ActionBidEditRequestResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// TODO: refine with typed response

// Fetch bid rating for a bid
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Get(ctx context.Context, bidID int64) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

type GetByListOfBidsOptions struct {
	Bids []int64 `url:"bids[]"`
}

// TODO: refine with typed response

// Fetch bid ratings for multiple bids
// it maps to the `GET` `/projects/0.1/bid_ratings` endpoint
func (s *BidRatingsService) GetByListOfBids(ctx context.Context, opts *GetByListOfBidsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsBidRatings
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Rating required
type CreateBidRatingsBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

// Rates a bid (creates a bid rating)
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Create(ctx context.Context, bidID int64, b CreateBidRatingsBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type UpdateBidRatingBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

// Updates an existing bid rating
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` endpoint
func (s *BidRatingsService) Update(ctx context.Context, bidID int64, bidRatingId int64, b UpdateBidRatingBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings/%d", endpoints.ProjectsBids, bidID, bidRatingId)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

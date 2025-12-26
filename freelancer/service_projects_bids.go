package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ListBidsOptions struct {
	Bids                        []int64
	Projects                    []int64
	Bidders                     []int64
	ProjectOwners               []int64
	AwardStatuses               []BidAwardStatus
	PaidStatuses                []BidPaidStatus
	CompleteStatuses            []BidCompleteStatus
	FrontBidStatuses            []BidFrontendStatus
	FromTime                    *int64
	ToTime                      *int64
	Reputation                  *bool
	BuyerProjectFee             *bool
	AwardStatusPossibilities    *bool
	ProjectDetails              *bool
	UserDetails                 *bool
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
	Compact                     *bool
}

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// It maps to the `GET` `/projects/0.1/bids` endpoint
func (s *BidsService) List(ctx context.Context, opts *ListBidsOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.Bids {
			addInt64(query, "bids[]", &id)
		}
		for _, id := range opts.Projects {
			addInt64(query, "projects[]", &id)
		}
		for _, id := range opts.Bidders {
			addInt64(query, "bidders[]", &id)
		}
		for _, id := range opts.ProjectOwners {
			addInt64(query, "project_owners[]", &id)
		}
		for _, status := range opts.AwardStatuses {
			addEnum(query, "award_statuses[]", &status)
		}
		for _, status := range opts.PaidStatuses {

			addEnum(query, "paid_statuses[]", &status)
		}
		for _, status := range opts.CompleteStatuses {
			addEnum(query, "complete_statuses[]", &status)
		}
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addBool(query, "reputation", opts.Reputation)
		addBool(query, "buyer_project_fee", opts.BuyerProjectFee)
		addBool(query, "award_status_possibilities", opts.AwardStatusPossibilities)
		addBool(query, "project_details", opts.ProjectDetails)
		addBool(query, "user_details", opts.UserDetails)
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

		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)

	}

	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsBids), query, nil)
}

type GetBidOptions struct {
	Reputation                  *bool
	BuyerProjectFee             *bool
	AwardStatusPossibilities    *bool
	ProjectDetails              *bool
	UserDetails                 *bool
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
	Compact                     *bool
	Quotations                  *bool
}

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// it maps to the `GET` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Get(ctx context.Context, bidID int64, opts *GetBidOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s", ProjectsBids, strconv.FormatInt(bidID, 10))
	query := url.Values{}
	if opts != nil {
		addBool(query, "reputation", opts.Reputation)
		addBool(query, "buyer_project_fee", opts.BuyerProjectFee)
		addBool(query, "award_status_possibilities", opts.AwardStatusPossibilities)
		addBool(query, "project_details", opts.ProjectDetails)
		addBool(query, "user_details", opts.UserDetails)
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
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)
		addBool(query, "quotations", opts.Quotations)

	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
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
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(ProjectsBids), nil, b)
}

type ActionBidBody struct {
	Action BidAction `json:"action"`
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Action(ctx context.Context, bidID int64, b *ActionBidBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s", ProjectsBids, strconv.FormatInt(bidID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

type UpdateBidBody struct {
	Amount              int    `json:"amount"`
	MilestonePercentage int    `json:"milestone_percentage"`
	Description         string `json:"description"`
}

// Updates an existing bid on a project. An existing bids information (description,amount,milestone_percentage) can be updated by sending a JSON encoded Bid struct.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Update(ctx context.Context, bidID int64, b UpdateBidBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s", ProjectsBids, strconv.FormatInt(bidID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

type GetTimeTrackingOptions struct {
	FromTime              *int64
	ToTime                *int64
	DailyAggregateDetails *bool
	Invoiced              *bool
}

// TODO: refine with typed response

// Returns a list of aggregate time tracking data for a bid.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (s *BidsService) GetTimeTracking(ctx context.Context, bidID int64, opts *GetTimeTrackingOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/time_tracking", ProjectsBids, strconv.FormatInt(bidID, 10))
	query := url.Values{}
	if opts != nil {
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addBool(query, "daily_aggregate_details", opts.DailyAggregateDetails)
		addBool(query, "invoiced", opts.Invoiced)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
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
	path := fmt.Sprintf("%s/%s/time_tracking", ProjectsBids, strconv.FormatInt(bidID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

type GetBidEditRequestsOptions struct {
	Statuses          []BidStatus
	BidEditRequestIDs []int
}

// TODO: refine with typed response

// Return bid edit requests by bid id.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/edit_requests` endpoint
func (s *BidEditRequestService) Get(ctx context.Context, bidID int64, opts *GetBidEditRequestsOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%si/", ProjectsBids, strconv.FormatInt(bidID, 10))
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Statuses {
			query.Add("statuses[]", string(val))
		}
		for _, val := range opts.BidEditRequestIDs {
			query.Add("bid_edit_request_ids[]", strconv.Itoa(val))
		}
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type CreateBidEditRequestsBody struct {
	BidID     int64 `json:"bid_id"`
	NewAmount int   `json:"new_amount"`
	NewPeriod int   `json:"new_period"`
	Comment   int   `json:"comment"`
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request.
// It maps to the `POST` `/projects/0.1/bids/edit_requests` endpoint
func (s *BidEditRequestService) Create(ctx context.Context, bidID int64, b CreateBidEditRequestsBody) (*RawResponse, error) {
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(ProjectsBidsBidEditRequests), nil, b)
}

type BidEditRequestsActionBody struct {
	Action BidEditRequestAction `json:"action"`
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (s *BidEditRequestService) Action(ctx context.Context, bidID, bidEditRequestID int64, b BidEditRequestsActionBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%si/edit_requests/%s", ProjectsBids, strconv.FormatInt(bidID, 10), strconv.FormatInt(bidEditRequestID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

// TODO: refine with typed response

// Fetch bid rating for a bid
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Get(ctx context.Context, bidID int64) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d/bid_ratings", ProjectsBids, strconv.FormatInt(bidID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

type GetByListOfBidsOptions struct {
	Bids []int64
}

// TODO: refine with typed response

// Fetch bid ratings for multiple bids
// it maps to the `GET` `/projects/0.1/bid_ratings` endpoint
func (s *BidRatingsService) GetByListOfBids(ctx context.Context, opts *GetByListOfBidsOptions) (*RawResponse, error) {
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsBidRatings), nil, nil)
}

// Rating required
type CreateBidRatingsBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

// Rates a bid (creates a bid rating)
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Create(ctx context.Context, bidID int64, b CreateBidRatingsBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d/bid_ratings", ProjectsBids, bidID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

type UpdateBidRatingBody struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

// Updates an existing bid rating
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` endpoint
func (s *BidRatingsService) Update(ctx context.Context, bidID int64, bidRatingId int64, b UpdateBidRatingBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d/bid_ratings/%d", ProjectsBids, bidID, bidRatingId)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

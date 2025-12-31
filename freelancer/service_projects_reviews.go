package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

type ListReviewsOptions struct {
	Projects                    []int64            `url:"projects[]"`
	FromUsers                   []int64            `url:"from_users[]"`
	ToUsers                     []int64            `url:"to_users[]"`
	Contests                    []int64            `url:"contests[]"`
	ReviewTypes                 []ReviewType       `url:"review_types[]"`
	JobIds                      []int64            `url:"job_ids[]"`
	CompletionStatuses          []CompletionStatus `url:"completion_statuses[]"`
	FromTime                    *int64             `url:"from_time"`
	ToTime                      *int64             `url:"to_time"`
	ReviewStatus                []string           `url:"review_status[]"`
	ProjectDetails              *bool              `url:"project_details"`
	Ratings                     *bool              `url:"ratings"`
	ReviewCount                 *bool              `url:"review_count"`
	Role                        *RoleType          `url:"role"`
	ContestDetails              *bool              `url:"contest_details"`
	UserDetails                 *bool              `url:"user_details"`
	ProjectFullDescription      *bool              `url:"project_full_description"`
	ProjectUpgradeDetails       *bool              `url:"project_upgrade_details"`
	ProjectJobDetails           *bool              `url:"project_job_details"`
	ProjectSelectedBids         *bool              `url:"project_selected_bids"`
	ProjectQualificationDetails *bool              `url:"project_qualification_details"`
	ProjectAttachmentDetails    *bool              `url:"project_attachment_details"`
	ProjectHiremeDetails        *bool              `url:"project_hireme_details"`
	ContestJobDetails           *bool              `url:"contest_job_details"`
	UserAvatar                  *bool              `url:"user_avatar"`
	UserCountryDetails          *bool              `url:"user_country_details"`
	UserProfileDescription      *bool              `url:"user_profile_description"`
	UserDisplayInfo             *bool              `url:"user_display_info"`
	UserJobs                    *bool              `url:"user_jobs"`
	UserBalanceDetails          *bool              `url:"user_balance_details"`
	UserQualificationDetails    *bool              `url:"user_qualification_details"`
	UserMembershipDetails       *bool              `url:"user_membership_details"`
	UserFinancialDetails        *bool              `url:"user_financial_details"`
	UserLocationDetails         *bool              `url:"user_location_details"`
	UserPortfolioDetails        *bool              `url:"user_portfolio_details"`
	UserPreferredDetails        *bool              `url:"user_preferred_details"`
	UserBadgeDetails            *bool              `url:"user_badge_details"`
	UserStatus                  *bool              `url:"user_status"`
	UserReputation              *bool              `url:"user_reputation"`
	UserEmployerReputation      *bool              `url:"user_employer_reputation"`
	UserReputationExtra         *bool              `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool              `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool              `url:"user_cover_image"`
	UserPastCoverImage          *bool              `url:"user_past_cover_image"`
	UserRecommendations         *bool              `url:"user_recommendations"`
	UserResponsiveness          *bool              `url:"user_responsiveness"`
	CorporateUsers              *bool              `url:"corporate_users"`
	MarketingMobileNumber       *bool              `url:"marketing_mobile_number"`
	SanctionDetails             *bool              `url:"sanction_details"`
	LimitedAccount              *bool              `url:"limited_account"`
	EquipmentGroupDetails       *bool              `url:"equipment_group_details"`
	Limit                       *int               `url:"limit"`
	Offset                      *int               `url:"offset"`
	Compact                     *bool              `url:"compact"`
}

// TODO: refine with typed response

// Returns a list of project reviews.
// It maps to the `GET` `/projects/0.1/reviews` endpoint
func (s *ReviewsService) List(ctx context.Context, opts *ListReviewsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsReviews
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type CreateReviewBody struct {
	ProjectID  int64      `json:"project_id"`
	ToUserID   int64      `json:"to_user_id"`
	FromUserID int64      `json:"from_user_id"`
	ReviewType ReviewType `json:"review_type"`
	Comment    string     `json:"comment"`
	Role       RoleType   `json:"role"`
}

// Post a review of a user.
// It maps to the `POST` `/projects/0.1/reviews` endpoint
func (s *ReviewsService) Create(ctx context.Context, b CreateReviewBody) (*RawResponse, error) {
	p := endpoints.ProjectsReviews
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ReviewActionBody struct {
	Action     ReviewAction `json:"action"`
	ReviewType ReviewType   `json:"review_type"`
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (s *ReviewsService) Action(ctx context.Context, reviewID int64, b ReviewActionBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsReviews, reviewID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

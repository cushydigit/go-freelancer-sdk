package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/endpoints"
)

type ListReviewsOptions struct {
	Projects                    []int64
	FromUsers                   []int64
	ToUsers                     []int64
	Contests                    []int64
	ReviewTypes                 []ReviewType
	JobIds                      []int64
	CompletionStatuses          []CompletionStatus
	FromTime                    *int64
	ToTime                      *int64
	ReviewStatus                []string
	ProjectDetails              *bool
	Ratings                     *bool
	ReviewCount                 *bool
	Role                        *RoleType
	ContestDetails              *bool
	UserDetails                 *bool
	ProjectFullDescription      *bool
	ProjectUpgradeDetails       *bool
	ProjectJobDetails           *bool
	ProjectSelectedBids         *bool
	ProjectQualificationDetails *bool
	ProjectAttachmentDetails    *bool
	ProjectHiremeDetails        *bool
	ContestJobDetails           *bool
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

// Returns a list of project reviews.
// It maps to the `GET` `/projects/0.1/reviews` endpoint
func (s *ReviewsService) List(ctx context.Context, opts *ListReviewsOptions) (*RawResponse, error) {
	path := endpoints.ProjectsReviews
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.Projects {
			addInt64(query, "projects[]", &id)
		}
		for _, id := range opts.FromUsers {
			addInt64(query, "from_users[]", &id)
		}
		for _, id := range opts.ToUsers {
			addInt64(query, "to_users[]", &id)
		}
		for _, id := range opts.Contests {
			addInt64(query, "contests[]", &id)
		}
		for _, status := range opts.ReviewTypes {
			addEnum(query, "review_types[]", &status)
		}
		for _, id := range opts.JobIds {
			addInt64(query, "job_ids[]", &id)
		}
		for _, status := range opts.CompletionStatuses {
			addEnum(query, "completion_statuses[]", &status)
		}
		for _, status := range opts.ReviewStatus {
			addEnum(query, "review_status", &status)
		}
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addBool(query, "project_details", opts.ProjectDetails)
		addBool(query, "ratings", opts.Ratings)
		addBool(query, "review_count", opts.ReviewCount)
		addEnum(query, "role", opts.Role)
		addBool(query, "contest_details", opts.ContestDetails)
		addBool(query, "user_details", opts.UserDetails)
		addBool(query, "project_full_description", opts.ProjectFullDescription)
		addBool(query, "project_upgrade_details", opts.ProjectUpgradeDetails)
		addBool(query, "project_job_details", opts.ProjectJobDetails)
		addBool(query, "project_selected_bids", opts.ProjectSelectedBids)
		addBool(query, "project_qualification_details", opts.ProjectQualificationDetails)
		addBool(query, "project_attachment_details", opts.ProjectAttachmentDetails)
		addBool(query, "project_hireme_details", opts.ProjectHiremeDetails)
		addBool(query, "contest_job_details", opts.ContestJobDetails)
		addBool(query, "user_avatar", opts.UserAvatar)
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
		addBool(query, "compact", opts.Compact)

	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
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
	path := endpoints.ProjectsReviews
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

type ReviewActionBody struct {
	Action     ReviewAction `json:"action"`
	ReviewType ReviewType   `json:"review_type"`
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (s *ReviewsService) Action(ctx context.Context, reviewID int64, b ReviewActionBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%d", endpoints.ProjectsReviews, reviewID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, b)
}

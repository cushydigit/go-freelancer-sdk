package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

// --------------------------------------
// PROJECTS
// --------------------------------------

// Title, Description, Budget, Jobs are required
type CreateProjectBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Budget      struct {
		Minimum    float64  `json:"minimum"`
		Maximum    *float64 `json:"maximum,omitempty"`
		CurrencyID *int64   `json:"currency_id,omitempty"`
	} `json:"budget"`
	Jobs              []int64      `json:"jobs"`
	Type              *ProjectType `json:"type,omitempty"`
	HourlyProjectInfo *struct {
		Commitment struct {
			Hours    int          `json:"hours"`
			Interval IntervalType `json:"interval"`
		} `json:"commitment"`
	} `json:"hourly_project_info,omitempty"`
	HirMe            *bool `json:"hire_me,omitempty"`
	HiremeInitialBid *struct {
		BidderID int64   `json:"bidder_id"`
		Amount   float64 `json:"amount"`
		Period   int64   `json:"period"`
	} `json:"hireme_initial_bid,omitempty"`
}

// Create a new project
// It maps to the `POST` `/projects/0.1/projects` endpoint.
func (s *ProjectsService) Create(ctx context.Context, b CreateProjectBody) (*CreateProjectResponse, error) {
	p := endpoints.Projects
	return execute[*CreateProjectResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: check this service functionality so the projectActionBody should have better fields
// check : https://developers.freelancer.com/docs/projects/projects#projects-put
// Performs an action on a project.
type ProjectActionBody struct {
	ProjectID int64         `json:"project_id"`
	Action    ProjectAction `json:"action"`
}

// Perform an action on a project
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Action(ctx context.Context, projectID int64, action ProjectActionBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, projectID)
}

type ListProjectsOptions struct {
	Projects                     []int64  `url:"projects[]"`
	Owners                       []int64  `url:"owners[]"`
	Bidders                      []int64  `url:"bidders[]"`
	SeoUrls                      []string `url:"seo_urls[]"`
	FromTime                     *int64   `url:"from_time"`
	ToTime                       *int64   `url:"to_time"`
	FrontendProjectStatuses      []string `url:"frontend_project_statuses[]"`
	Team                         *bool    `url:"team"`
	IsNonHireMe                  *bool    `url:"is_none_hire_me"`
	HasMilestone                 *bool    `url:"has_milestone"`
	Count                        *bool    `url:"count"`
	FullDescription              *bool    `url:"full_description"`
	JobDetails                   *bool    `url:"job_details"`
	UpgradeDetails               *bool    `url:"upgrade_details"`
	AttachmentDetails            *bool    `url:"attachment_details"`
	FileDetails                  *bool    `url:"file_details"`
	QualificationDetails         *bool    `url:"qualification_details"`
	SelectedBids                 *bool    `url:"selected_bids"`
	HiremeDetails                *bool    `url:"hireme_details"`
	UserDetails                  *bool    `url:"user_details"`
	InvitedFreelancerDetails     *bool    `url:"invited_freelancer_details"`
	RecommendedFreelancerDetails *bool    `url:"recommended_freelancer_details"`
	SupportSessionDetails        *bool    `url:"support_session_details"`
	LocationDetails              *bool    `url:"location_details"`
	NdaSignatureDetails          *bool    `url:"nda_signature_details"`
	ProjectCollaborationDetails  *bool    `url:"project_collaboration_details"`
	ProximityDetails             *bool    `url:"proximity_details"`
	ReviewAvailabilityDetails    *bool    `url:"review_availability_details"`
	NegotiatedDetails            *bool    `url:"negotiated_details"`
	DriveFileDetails             *bool    `url:"drive_file_details"`
	NdaDetails                   *bool    `url:"nda_details"`
	LocalDetails                 *bool    `url:"local_details"`
	EquipmentDetails             *bool    `url:"equipment_details"`
	ClientEngagementDetails      *bool    `url:"client_engagement_details"`
	ServiceOfferingDetails       *bool    `url:"service_offering_details"`
	UserAvatar                   *bool    `url:"user_avatar"`
	UserCountryDetails           *bool    `url:"user_country_details"`
	UserProfileDescription       *bool    `url:"user_profile_description"`
	UserDisplayInfo              *bool    `url:"user_display_info"`
	UserJobs                     *bool    `url:"user_jobs"`
	UserBalanceDetails           *bool    `url:"user_balance_details"`
	UserQualificationDetails     *bool    `url:"user_qualification_details"`
	UserMembershipDetails        *bool    `url:"user_membership_details"`
	UserFinancialDetails         *bool    `url:"user_financial_details"`
	UserLocationDetails          *bool    `url:"user_location_details"`
	UserPortfolioDetails         *bool    `url:"user_portfolio_details"`
	UserPreferredDetails         *bool    `url:"user_preferred_details"`
	UserBadgeDetails             *bool    `url:"user_badge_details"`
	UserStatus                   *bool    `url:"user_status"`
	UserReputation               *bool    `url:"user_reputation"`
	UserEmployerReputation       *bool    `url:"user_employer_reputation"`
	UserReputationExtra          *bool    `url:"user_reputation_extra"`
	UserEmployerReputationExtra  *bool    `url:"user_employer_reputation_extra"`
	UserCoverImage               *bool    `url:"user_cover_image"`
	UserPastCoverImage           *bool    `url:"user_past_cover_image"`
	UserRecommendations          *bool    `url:"user_recommendations"`
	UserResponsiveness           *bool    `url:"user_responsiveness"`
	CorporateUsers               *bool    `url:"corporate_users"`
	MarketingMobileNumber        *bool    `url:"marketing_mobile_number"`
	SanctionDetails              *bool    `url:"sanction_details"`
	LimitedAccount               *bool    `url:"limited_account"`
	EquipmentGroupDetails        *bool    `url:"equipment_group_details"`
	Limit                        *int     `url:"limit"`
	Offset                       *int     `url:"offset"`
	Compact                      *bool    `url:"compact"`
}

// Returns information about multiple projects. Will be ordered by descending submit date (newest-to-oldest).
// it maps to the `GET` `/projects/0.1/projects` endpoint
func (s *ProjectsService) List(ctx context.Context, opts *ListProjectsOptions) (*ListProjectsResponse, error) {
	p := endpoints.Projects
	q := query.Values(opts)
	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListSelfProjectsOptions struct {
	Status      *ProjectStatusType `url:"status"`
	Role        *RoleType          `url:"role"`
	Types       []TypeType         `url:"type[]"`
	Query       *string            `url:"query"`
	SortField   *SortField         `url:"sort_field"`
	ReverseSort *bool              `url:"reverse_sort"`
	Recruiter   *bool              `url:"recruiter"`
	Offset      *int               `url:"offset"`
	Limit       *int               `url:"limit"`
}

// Returns the logged in user’s projects/contests they either created or participated in (by bidding or submitting an entry).
// it maps to the `GET` `/projects/0.1/self` endpoint
func (s *ProjectsService) ListSelf(ctx context.Context, opts *ListSelfProjectsOptions) (*ListProjectsResponse, error) {
	p := endpoints.ProjectsSelf
	q := query.Values(opts)
	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type GetProjectOptions struct {
	FullDescription              *bool `url:"full_description"`
	JobDetails                   *bool `url:"job_details"`
	UpgradeDetails               *bool `url:"upgrade_details"`
	AttachmentDetails            *bool `url:"attached_details"`
	FileDetails                  *bool `url:"file_details"`
	QualificationDetails         *bool `url:"qualification_details"`
	SelectedBids                 *bool `url:"selected_bid"`
	HiremeDetails                *bool `url:"hireme_details"`
	UserDetails                  *bool `url:"user_details"`
	InvitedFreelancerDetails     *bool `url:"invited_freelancer_details"`
	RecommendedFreelancerDetails *bool `url:"recommended_freelancer_details"`
	HourlyDetails                *bool `url:"hourly_details"`
	SupportSessionDetails        *bool `url:"support_session_details"`
	LocationDetails              *bool `url:"local_details"`
	NdaSignatureDetails          *bool `url:"nda_signature_details"`
	ProjectCollaborationDetails  *bool `url:"project_collaboration_details"`
	TrackDetails                 *bool `url:"track_details"`
	ProximityDetails             *bool `url:"proximity_details"`
	ReviewAvailabilityDetails    *bool `url:"review_availability_details"`
	NegotiatedDetails            *bool `url:"negotiated_details"`
	DriveFileDetails             *bool `url:"drive_file_details"`
	NdaDetails                   *bool `url:"nda_details"`
	LocalDetails                 *bool `url:"local_details"`
	EquipmentDetails             *bool `url:"equipment_details"`
	ClientEngagementDetails      *bool `url:"client_engagement_details"`
	ServiceOfferingDetails       *bool `url:"service_offering_details"`
	UserAvatar                   *bool `url:"user_avatar"`
	UserCountryDetails           *bool `url:"user_country_details"`
	UserProfileDescription       *bool `url:"user_profile_description"`
	UserDisplayInfo              *bool `url:"user_display_info"`
	UserJobs                     *bool `url:"user_jobs"`
	UserBalanceDetails           *bool `url:"user_balance_details"`
	UserQualificationDetails     *bool `url:"user_qualification_details"`
	UserMembershipDetails        *bool `url:"user_membership_details"`
	UserFinancialDetails         *bool `url:"user_financial_details"`
	UserLocationDetails          *bool `url:"user_location_details"`
	UserPortfolioDetails         *bool `url:"user_portfolio_details"`
	UserPreferredDetails         *bool `url:"user_preferred_details"`
	UserBadgeDetails             *bool `url:"user_badge_details"`
	UserStatus                   *bool `url:"user_status"`
	UserReputation               *bool `url:"user_reputation"`
	UserEmployerReputation       *bool `url:"user_employer_reputation"`
	UserReputationExtra          *bool `url:"user_reputation_extra"`
	UserEmployerReputationExtra  *bool `url:"user_employer_reputation_extra"`
	UserCoverImage               *bool `url:"user_cover_image"`
	UserPastCoverImage           *bool `url:"user_past_cover_image"`
	UserRecommendations          *bool `url:"user_recommendations"`
	UserResponsiveness           *bool `url:"user_responsiveness"`
	CorporateUsers               *bool `url:"corporate_users"`
	MarketingMobileNumber        *bool `url:"marketing_mobile_number"`
	SanctionDetails              *bool `url:"sanction_details"`
	LimitedAccount               *bool `url:"limited_account"`
	EquipmentGroupDetails        *bool `url:"equipment_group_details"`
	Limit                        *int  `url:"limit"`
	Offset                       *int  `url:"offset"`
	Compact                      *bool `url:"compact"`
}

// Get information about a specific project. The full range of users projection options can be specified as part of this request by first setting theuser_detailsparameter to true.
// It maps to the `GET` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Get(ctx context.Context, projectID int64, opts *GetProjectOptions) (*GetProjectResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*GetProjectResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchActiveProjectsOptions struct {
	Query                       *string              `url:"query"`
	ProjectTypes                []ProjectType        `url:"project_types[]"`
	ProjectUpgrades             []ProjectUpgradeType `url:"project_upgrades[]"`
	ContestUpgrades             []ContestUpgradeType `url:"contest_upgrades[]"`
	MinAvgPrice                 *float64             `url:"min_avg_price"`
	MaxAvgPrice                 *float64             `url:"max_avg_price"`
	MinAvgHourlyRate            *float64             `url:"min_avg_hourly_rate"`
	MaxAvgHourlyRate            *float64             `url:"max_avg_hourly_rate"`
	MinPrice                    *float64             `url:"min_price"`
	MaxPrice                    *float64             `url:"max_price"`
	MinHourlyRate               *float64             `url:"min_hourly_rate"`
	MaxHourlyRate               *float64             `url:"max_hourly_rate"`
	Jobs                        []int64              `url:"jobs[]"`
	Countries                   []string             `url:"countries[]"`
	Languages                   []string             `url:"languages[]"`
	Latitude                    *float64             `url:"latitude"`
	Longitude                   *float64             `url:"longitude"`
	FromTime                    *int64               `url:"from_time"`
	ToTime                      *int64               `url:"to_time"`
	SortField                   *SortField           `url:"sort_field"`
	ProjectIDs                  []int64              `url:"project_ids[]"`
	TopRightLatitude            *float64             `url:"top_right_latitude"`
	TopRightLongitude           *float64             `url:"top_right_longitude"`
	BottomLeftLatitude          *float64             `url:"bottom_left_latitude"`
	BottomLeftLongitude         *float64             `url:"bottom_left_longitude"`
	ReverseSort                 *bool                `url:"reverse_sort"`
	OrSearchQuery               *string              `url:"or_search_query"`
	HighlightPreTags            *string              `url:"highlight_pre_tags"`
	HighlightPostTags           *string              `url:"highlight_post_tags"`
	FullDescription             *bool                `url:"full_description"`
	JobDetails                  *bool                `url:"job_details"`
	UpgradeDetails              *bool                `url:"upgrade_details"`
	UserDetails                 *bool                `url:"user_details"`
	LocationDetails             *bool                `url:"location_details"`
	NDASignatureDetails         *bool                `url:"nda_signature_details"`
	ProjectCollaborationDetails *bool                `url:"project_collaboration_details"`
	UserAvatar                  *bool                `url:"user_avatar"`
	UserCountryDetails          *bool                `url:"user_country_details"`
	UserProfileDescription      *bool                `url:"user_profile_description"`
	UserDisplayInfo             *bool                `url:"user_display_info"`
	UserJobs                    *bool                `url:"user_jobs"`
	UserBalanceDetails          *bool                `url:"user_balance_details"`
	UserQualificationDetails    *bool                `url:"user_qualification_details"`
	UserMembershipDetails       *bool                `url:"user_membership_details"`
	UserFinancialDetails        *bool                `url:"user_financial_details"`
	UserLocationDetails         *bool                `url:"user_location_details"`
	UserPortfolioDetails        *bool                `url:"user_portfolio_details"`
	UserPreferredDetails        *bool                `url:"user_preferred_details"`
	UserBadgeDetails            *bool                `url:"user_badge_details"`
	UserStatus                  *bool                `url:"user_status"`
	UserReputation              *bool                `url:"user_reputation"`
	UserEmployerReputation      *bool                `url:"user_employer_reputation"`
	UserReputationExtra         *bool                `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool                `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool                `url:"user_cover_image"`
	UserPastCoverImage          *bool                `url:"user_past_cover_image"`
	UserRecommendations         *bool                `url:"user_recommendations"`
	UserResponsiveness          *bool                `url:"user_responsiveness"`
	CorporateUsers              *bool                `url:"corporate_users"`
	MarketingMobileNumber       *bool                `url:"marketing_mobile_number"`
	SanctionDetails             *bool                `url:"sanction_details"`
	LimitedAccount              *bool                `url:"limited_account"`
	EquipmentGroupDetails       *bool                `url:"equipment_group_details"`
	Limit                       *int                 `url:"limit"`
	Offset                      *int                 `url:"offset"`
	Compact                     *bool                `url:"compact"`
}

// Searches for active projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/active` endpoint
func (s *ProjectsService) SearchActive(ctx context.Context, opts *SearchActiveProjectsOptions) (*ListProjectsResponse, error) {
	p := endpoints.ProjectsActive
	q := query.Values(opts)
	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchAllProjectsOptions struct {
	Query                       *string                 `url:"query"`
	ProjectTypes                []ProjectType           `url:"project_types[]"`
	ProjectUpgrades             []ProjectUpgradeType    `url:"project_upgrades[]"`
	ContestUpgrades             []ContestUpgradeType    `url:"contest_upgrades[]"`
	MinAvgPrice                 *float64                `url:"min_avg_price"`
	MaxAvgPrice                 *float64                `url:"max_avg_price"`
	MinAvgHourlyRate            *float64                `url:"min_avg_hourly_rate"`
	MaxAvgHourlyRate            *float64                `url:"max_avg_hourly_rate"`
	MinPrice                    *float64                `url:"min_price"`
	MaxPrice                    *float64                `url:"max_price"`
	MinHourlyRate               *float64                `url:"min_hourly_rate"`
	MaxHourlyRate               *float64                `url:"max_hourly_rate"`
	Jobs                        []int64                 `url:"jobs[]"`
	Countries                   []string                `url:"countries[]"`
	Languages                   []string                `url:"languages[]"`
	Latitude                    *float64                `url:"latitude"`
	Longitude                   *float64                `url:"longitude"`
	FromTime                    *int64                  `url:"from_time"`
	ToTime                      *int64                  `url:"to_time"`
	SortField                   *SortField              `url:"sort_field"`
	BidAwardStatuses            []BidAwardStatus        `url:"bid_award_statuses[]"`
	BidCompleteStatuses         []BidCompleteStatus     `url:"bid_complete_statuses[]"`
	ProjectStatuses             []ProjectFrontendStatus `url:"project_statuses[]"`
	ProjectIDs                  []int64                 `url:"project_ids[]"`
	TopRightLatitude            *float64                `url:"top_right_latitude"`
	TopRightLongitude           *float64                `url:"top_right_longitude"`
	BottomLeftLatitude          *float64                `url:"bottom_left_latitude"`
	BottomLeftLongitude         *float64                `url:"bottom_left_longitude"`
	ReverseSort                 *bool                   `url:"reverse_sort"`
	OrSearchQuery               *string                 `url:"or_search_query"`
	HighlightPreTags            *string                 `url:"highlight_pre_tags"`
	HighlightPostTags           *string                 `url:"highlight_post_tags"`
	FullDescription             *bool                   `url:"full_description"`
	JobDetails                  *bool                   `url:"job_details"`
	UpgradeDetails              *bool                   `url:"upgrade_details"`
	UserDetails                 *bool                   `url:"user_details"`
	UserAvatar                  *bool                   `url:"user_avatar"`
	UserCountryDetails          *bool                   `url:"user_country_details"`
	UserProfileDescription      *bool                   `url:"user_profile_description"`
	UserDisplayInfo             *bool                   `url:"user_display_info"`
	UserJobs                    *bool                   `url:"user_jobs"`
	UserBalanceDetails          *bool                   `url:"user_balance_details"`
	UserQualificationDetails    *bool                   `url:"user_qualification_details"`
	UserMembershipDetails       *bool                   `url:"user_membership_details"`
	UserFinancialDetails        *bool                   `url:"user_financial_details"`
	UserLocationDetails         *bool                   `url:"user_location_details"`
	UserPortfolioDetails        *bool                   `url:"user_portfolio_details"`
	UserPreferredDetails        *bool                   `url:"user_preferred_details"`
	UserBadgeDetails            *bool                   `url:"user_badge_details"`
	UserStatus                  *bool                   `url:"user_status"`
	UserReputation              *bool                   `url:"user_reputation"`
	UserEmployerReputation      *bool                   `url:"user_employer_reputation"`
	UserReputationExtra         *bool                   `url:"user_reputation_extra"`
	UserEmployerReputationExtra *bool                   `url:"user_employer_reputation_extra"`
	UserCoverImage              *bool                   `url:"user_cover_image"`
	UserPastCoverImage          *bool                   `url:"user_past_cover_image"`
	UserRecommendations         *bool                   `url:"user_recommendations"`
	UserResponsiveness          *bool                   `url:"user_responsiveness"`
	CorporateUsers              *bool                   `url:"corporate_users"`
	MarketingMobileNumber       *bool                   `url:"marketing_mobile_number"`
	SanctionDetails             *bool                   `url:"sanction_details"`
	LimitedAccount              *bool                   `url:"limited_account"`
	EquipmentGroupDetails       *bool                   `url:"equipment_group_details"`
	Limit                       *int                    `url:"limit"`
	Offset                      *int                    `url:"offset"`
	Compact                     *bool                   `url:"compact"`
}

// Searches for all projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/all` endpoint
func (s *ProjectsService) SearchAll(ctx context.Context, opts *SearchAllProjectsOptions) (*ListProjectsResponse, error) {
	p := endpoints.ProjectsAll
	q := query.Values(opts)
	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type InviteFreelancersBody struct {
	FreelancerID int64 `json:"freelancer_id"`
}

// Invites specific freelancers to bid on a project.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/invite` endpoint
func (s *ProjectsService) InviteFreelancer(ctx context.Context, projectID int64, b InviteFreelancersBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/invite", endpoints.Projects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

type ListUpgradesFeesOptions struct {
	Currencies         []int64 `url:"currencies[]"`
	Project            *int64  `url:"project"`
	FreeUpgradeDetails *bool   `url:"fee_upgrade_details"`
	TaxIncluded        *bool   `url:"tax_included"`
}

// TODO: Refine the typed response with

// Returns the project upgrade fees for a given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
// It maps to the `GET` `/projects/0.1/projects/fees` endpoint
func (s *ProjectsService) ListUpgradesFees(ctx context.Context, opts *ListUpgradesFeesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsFees
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListProjectBidsOptions struct {
	IsShortlisted               *bool `url:"is_shortlisted"`
	Reputation                  *bool `url:"reputation"`
	RecommendedBid              *bool `url:"recommended_bid"`
	ShortlistedBid              *bool `url:"shortlisted_bid"`
	Distance                    *bool `url:"distance"`
	UserDetails                 *bool `url:"user_details"`
	ExpertGuarantees            *bool `url:"expert_guarantees"`
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

// TODO: Refine the typed response with

// Returns bids for a single project. Employers will see bids in a sorted order, which begins with sponsored bids, then by bid ranking. Freelancers will receive the bid list ordered by date. Note: This method is expensive to compute so it is recommended that reputation and user projection options are not set.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids` endpoint
func (s *ProjectsService) ListBids(ctx context.Context, projectID int64, opts *ListProjectBidsOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/bids", endpoints.Projects, strconv.FormatInt(projectID, 10))
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns information for posting bids on a project.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids_info` endpoint
func (s *ProjectsService) GetBidInfo(ctx context.Context, projectID int64) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bids_info", endpoints.Projects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

type ListProjectMilestonesOptions struct {
	Statuses                    []MilestoneStatus `url:"statuses[]"`
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

// Returns a list of milestones on a project. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/milestones` endpoint
func (s *ProjectsService) ListMilestones(ctx context.Context, projectID int64, opts *ListProjectMilestonesOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/milestones", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListProjectsMilestoneRequestsOptions struct {
	Statuses                    []MilestoneStatus `url:"statuses[]"`
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

// Returns a list of milestone requests by freelancers for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/milestone_requests` endpoint
func (s *ProjectsService) ListMilestoneRequests(ctx context.Context, projectID int64, opts *ListProjectsMilestoneRequestsOptions) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/milestone_requests", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type GetHourlyContractInfoOptions struct {
	ProjectIDs        []int64 `url:"project_ids[]"`
	BidderIDs         []int64 `url:"bidder_ids[]"`
	HourlyContractIDs []int64 `url:"hourly_contract_ids[]"`
	projectOwnerIDs   []int64 `url:"project_owner_ids[]"`
	billingDetails    *bool   `url:"billing_details"`
	invoiceDetails    *bool   `url:"invoice_details"`
}

// TODO: refine with typed response

// Fetch the hourly contract matching the desired query.
// It maps to the `GET` `/projects/0.1/hourly_contract_info` endpoint
func (s *ProjectsService) GetHourlyContractInfo(ctx context.Context, opts *GetHourlyContractInfoOptions) (*RawResponse, error) {
	p := endpoints.ProjectsHourlyContract
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Fetch the IP contract matching for the project id. If you are an employer it will return all of the contracts, ELSE we will return contract details specific to the logged-in user.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/ip_contract_info` endpoint
func (s *ProjectsService) GetIPContractInfo(ctx context.Context, projectID int64) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d/ip_contract_info", endpoints.Projects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Delete a project by id. Only projects that are in pending or rejected states may be deleted. This will close the project and remove its visibility.
// it maps to the `DELETE` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Delete(ctx context.Context, projectID int64) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodDelete, p, nil, nil)
}

// --------------------------------------
// PROJECTS-COLLABORATIONS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of project collaboration data for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) List(ctx context.Context, projectID int64) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/collaborations", endpoints.Projects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

type CreateProjectCollaborationsBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Title    string `json:"title"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

// Creates a new project collaboration.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) Create(ctx context.Context, projectID int64, b CreateProjectCollaborationsBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/collaborations", endpoints.Projects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, nil)
}

type ActionProjectCollaborationsBody struct {
	// required
	Action ProjectCollaborationAction `json:"action"`
	// required
	Permissions struct {
		Chat     bool `json:"CHAT"`
		BidAward bool `json:"BID_AWARD"`
	}
}

// Performs an action on a collaboration.
// it maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` endpoint
func (s *CollaborationsService) Action(ctx context.Context, projectID int64, collaborationID int, b ActionProjectCollaborationsBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/collaborations/%d/actions", endpoints.Projects, strconv.FormatInt(projectID, 10), collaborationID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: refine with typed response

// Returns a list of all collaboration data for a user.
// It maps toi the `GET` `/projects/0.1/collaborations` endpoint
func (s *CollaborationsService) ListAll(ctx context.Context) (*RawResponse, error) {
	p := endpoints.ProjectsCollaborations
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// --------------------------------------
// PROJECTS-SERVICES
// --------------------------------------

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (s *ServicesService) Order(ctx context.Context, serviceID int, serviceType ServiceType) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%s/%d/order", endpoints.ProjectsServices, serviceType, serviceID)
	return execute[*RawResponse](ctx, s.client, http.MethodPost, p, nil, nil)
}

type ListServicesOptions struct {
	Services                    []int64             `url:"services[]"`
	Owners                      []int64             `url:"owners[]"`
	Statuses                    []ServiceStatusType `url:"statuses[]"`
	SubStatuses                 []string            `url:"sub_statuses[]"`
	Titles                      []string            `url:"titles[]"`
	SeoUrls                     []string            `url:"seo_urls[]"`
	ExtraDetails                *bool               `url:"extra_details"`
	FileDetails                 *bool               `url:"file_details"`
	JobDetails                  *bool               `url:"job_details"`
	UserDetails                 *bool               `url:"user_details"`
	FullDescription             *bool               `url:"full_description"`
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

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (s *ServicesService) List(ctx context.Context, opts *ListServicesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsServices
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchActiveServicesOptions struct {
	Query        *string   `url:"query"`
	Sort         *SortType `url:"sort"`
	ReverseSort  *bool     `url:"reverse_sort"`
	ExtraDetails *bool     `url:"extra_details"`
	FileDetails  *bool     `url:"file_details"`
	JobDetails   *bool     `url:"job_details"`
	UserDetails  *bool     `url:"user_details"`
	Offset       *int      `url:"offset"`
	Compact      *bool     `url:"compact"`
}

// TODO: refine with typed response

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (s *ServicesService) SearchActive(ctx context.Context, opts *SearchActiveServicesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsServicesActive
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// PROJECTS-BIDS
// --------------------------------------

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

// --------------------------------------
// PROJECTS-JOBS
// --------------------------------------

type ListJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (s *JobsService) List(ctx context.Context, opts *ListJobsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobs
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type SearchJobsOptions struct {
	Jobs                      []int64  `url:"jobs[]"`
	JobNames                  []string `url:"job_names[]"`
	SeoUrls                   []string `url:"seo_urls[]"`
	SeoTexts                  []string `url:"seo_texts[]"`
	Categories                []int64  `url:"categories[]"`
	OnlyLocal                 *bool    `url:"only_local"`
	ActiveProjectCountDetails *bool    `url:"active_project_count_details"`
	SeoDetails                *bool    `url:"seo_details"`
	SeoCountryName            *string  `url:"seo_country_name"`
	Lang                      *string  `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (s *JobsService) Search(ctx context.Context, opts *SearchJobsOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobsSearch
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListJobBundlesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (s *JobsService) ListJobBundles(ctx context.Context, opts *ListJobBundlesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobBundles
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListJobBundleCategoriesOptions struct {
	JobBundles []int64 `url:"job_bundles[]"`
	Categories []int64 `url:"categories[]"`
	Lang       *string `url:"lang"`
}

// TODO: refine with typed response

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (s *JobsService) ListJobBundleCategories(ctx context.Context, opts *ListJobBundleCategoriesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsJobBundleCategories
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// PROJECTS-MILESTONES
// --------------------------------------

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

// --------------------------------------
// PROJECTS-REVIEWS
// --------------------------------------

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

// --------------------------------------
// PROJECTS-EXTRAS
// --------------------------------------
type listExpertGuaranteesOptions struct {
	ExpertGuarantees []int64                  `url:"expert_guarantees[]"`
	Projects         []int64                  `url:"projects[]"`
	ProjectOwners    []int64                  `url:"project_owners[]"`
	Bidders          []int64                  `url:"bidders[]"`
	Bids             []int64                  `url:"bids[]"`
	Statuses         []ExpertGuaranteesStatus `url:"statuses[]"`
	FromTime         *int64                   `url:"from_time"`
	ToTime           *int64                   `url:"to_time"`
	Offset           *int                     `url:"offset"`
	Limit            *int                     `url:"limit"`
}

// TODO: refine with typed response

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (s *ExpertGuaranteesService) List(ctx context.Context, opts *listExpertGuaranteesOptions) (*RawResponse, error) {
	p := endpoints.ProjectsExpertGuarantees
	q := query.Values(opts)
	return execute[*RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ExpertGuaranteesActionRequestBody struct {
	Action ExpertGuaranteesAction `json:"action"`
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (s *ExpertGuaranteesService) Action(ctx context.Context, expertGuaranteesID int64, b ExpertGuaranteesActionRequestBody) (*RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsExpertGuarantees, expertGuaranteesID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

type ListCurrenciesOptions struct {
	CurrencyCodes             []string `url:"currency_codes[]"`
	CurrencyIDs               []int64  `url:"currency_ids[]"`
	IncludeExternalCurrencies *bool    `url:"include_external_currencies"`
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (s *CurrenciesService) List(ctx context.Context, opts *ListCurrenciesOptions) (*ListCurrenciesResponse, error) {
	p := endpoints.ProjectsCurrencies
	q := query.Values(opts)
	return execute[*ListCurrenciesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListCategoriesOptions struct {
	Categories                []int64 `url:"categories[]"`
	JobDetails                *bool   `url:"job_details"`
	Lang                      *string `url:"lang"`
	ActiveProjectCountDetails *bool   `url:"active_project_count_details"`
	SeoDetails                *bool   `url:"seo_details"`
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (s *CategoriesService) List(ctx context.Context, opts *ListCategoriesOptions) (*ListCategoriesResponse, error) {
	p := endpoints.ProjectsCategories
	q := query.Values(opts)
	return execute[*ListCategoriesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

type ListBudgetsOptions struct {
	CurrencyCodes   []string     `url:"currency_codes[]"`
	CurrencyIDs     []int64      `url:"currency_ids[]"`
	ProjectType     *ProjectType `url:"project_type"`
	Lang            *string      `url:"lang"`
	CurrencyDetails *bool        `url:"currency_details"`
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (s *BudgetsService) List(ctx context.Context, opts *ListBudgetsOptions) (*ListBudgetsResponse, error) {
	p := endpoints.ProjectsBudgets
	q := query.Values(opts)
	return execute[*ListBudgetsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

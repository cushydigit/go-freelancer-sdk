package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
)

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

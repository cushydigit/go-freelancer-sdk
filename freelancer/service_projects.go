package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
func (s *ProjectsService) Create(ctx context.Context, b CreateProjectBody) (*RawResponse, error) {
	return execute[*RawResponse](ctx, s.client, http.MethodPost, string(ProjectsProjects), nil, b)
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
	path := fmt.Sprintf("%s/%d", ProjectsProjects, projectID)
	return execute[*RawResponse](ctx, s.client, http.MethodPut, path, nil, projectID)
}

type ListProjectsOptions struct {
	projects                     []int64
	owners                       []int64
	bidders                      []int64
	seoUrls                      []string
	fromTime                     *int64
	toTime                       *int64
	frontendProjectStatuses      []string
	team                         *bool
	isNonHireMe                  *bool
	hasMilestone                 *bool
	count                        *bool
	fullDescription              *bool
	jobDetails                   *bool
	upgradeDetails               *bool
	attachmentDetails            *bool
	fileDetails                  *bool
	qualificationDetails         *bool
	selectedBids                 *bool
	hiremeDetails                *bool
	userDetails                  *bool
	invitedFreelancerDetails     *bool
	recommendedFreelancerDetails *bool
	supportSessionDetails        *bool
	locationDetails              *bool
	ndaSignatureDetails          *bool
	projectCollaborationDetails  *bool
	proximityDetails             *bool
	reviewAvailabilityDetails    *bool
	negotiatedDetails            *bool
	driveFileDetails             *bool
	ndaDetails                   *bool
	localDetails                 *bool
	equipmentDetails             *bool
	clientEngagementDetails      *bool
	serviceOfferingDetails       *bool
	userAvatar                   *bool
	userCountryDetails           *bool
	userProfileDescription       *bool
	userDisplayInfo              *bool
	userJobs                     *bool
	userBalanceDetails           *bool
	userQualificationDetails     *bool
	userMembershipDetails        *bool
	userFinancialDetails         *bool
	userLocationDetails          *bool
	userPortfolioDetails         *bool
	userPreferredDetails         *bool
	userBadgeDetails             *bool
	userStatus                   *bool
	userReputation               *bool
	userEmployerReputation       *bool
	userReputationExtra          *bool
	userEmployerReputationExtra  *bool
	userCoverImage               *bool
	userPastCoverImage           *bool
	userRecommendations          *bool
	userResponsiveness           *bool
	corporateUsers               *bool
	marketingMobileNumber        *bool
	sanctionDetails              *bool
	limitedAccount               *bool
	equipmentGroupDetails        *bool
	limit                        *int
	offset                       *int
	compact                      *bool
}

// Returns information about multiple projects. Will be ordered by descending submit date (newest-to-oldest).
// it maps to the `GET` `/projects/0.1/projects` endpoint
func (s *ProjectsService) List(ctx context.Context, opts *ListProjectsOptions) (*ListProjectsResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, id := range opts.projects {
			addInt64(query, "projects[]", &id)
		}
		for _, id := range opts.owners {
			addInt64(query, "owners[]", &id)
		}
		for _, id := range opts.bidders {
			addInt64(query, "bidders[]", &id)
		}
		for _, seo := range opts.seoUrls {
			addString(query, "seo_urls[]", &seo)
		}
		for _, status := range opts.frontendProjectStatuses {
			addString(query, "frontend_project_statuses[]", &status)
		}
		addInt64(query, "from_time", opts.fromTime)
		addInt64(query, "to_time", opts.toTime)
		addBool(query, "team", opts.team)
		addBool(query, "is_none_hire_me", opts.isNonHireMe)
		addBool(query, "has_milestone", opts.hasMilestone)
		addBool(query, "count", opts.count)
		addBool(query, "full_description", opts.fullDescription)
		addBool(query, "job_details", opts.jobDetails)
		addBool(query, "upgrade_details", opts.upgradeDetails)
		addBool(query, "attachment_details", opts.attachmentDetails)
		addBool(query, "file_details", opts.fileDetails)
		addBool(query, "qualification_details", opts.qualificationDetails)
		addBool(query, "selected_bids", opts.selectedBids)
		addBool(query, "hireme_details", opts.hiremeDetails)
		addBool(query, "user_details", opts.userDetails)
		addBool(query, "invited_freelancer_details", opts.invitedFreelancerDetails)
		addBool(query, "recommended_freelancer_details", opts.recommendedFreelancerDetails)
		addBool(query, "support_session_details", opts.supportSessionDetails)
		addBool(query, "location_details", opts.locationDetails)
		addBool(query, "nda_signature_details", opts.ndaSignatureDetails)
		addBool(query, "project_collaboration_details", opts.projectCollaborationDetails)
		addBool(query, "proximity_details", opts.proximityDetails)
		addBool(query, "review_availability_details", opts.reviewAvailabilityDetails)
		addBool(query, "negotiated_details", opts.negotiatedDetails)
		addBool(query, "drive_file_details", opts.driveFileDetails)
		addBool(query, "nda_details", opts.ndaDetails)
		addBool(query, "local_details", opts.localDetails)
		addBool(query, "equipment_details", opts.equipmentDetails)
		addBool(query, "client_engagement_details", opts.clientEngagementDetails)
		addBool(query, "service_offering_details", opts.serviceOfferingDetails)
		addBool(query, "user_avatar", opts.userAvatar)
		addBool(query, "user_country_details", opts.userCountryDetails)
		addBool(query, "user_profile_description", opts.userProfileDescription)
		addBool(query, "user_display_info", opts.userDisplayInfo)
		addBool(query, "user_jobs", opts.userJobs)
		addBool(query, "user_balance_details", opts.userBalanceDetails)
		addBool(query, "user_qualification_details", opts.userQualificationDetails)
		addBool(query, "user_membership_details", opts.userMembershipDetails)
		addBool(query, "user_financial_details", opts.userFinancialDetails)
		addBool(query, "user_location_details", opts.userLocationDetails)
		addBool(query, "user_portfolio_details", opts.userPortfolioDetails)
		addBool(query, "user_preferred_details", opts.userPreferredDetails)
		addBool(query, "user_badge_details", opts.userBadgeDetails)
		addBool(query, "user_status", opts.userStatus)
		addBool(query, "user_reputation", opts.userReputation)
		addBool(query, "user_employer_reputation", opts.userReputation)
		addBool(query, "user_reputation_extra", opts.userReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.userEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.userCoverImage)
		addBool(query, "user_past_cover_image", opts.userPastCoverImage)
		addBool(query, "user_recommendations", opts.userRecommendations)
		addBool(query, "user_responsiveness", opts.userResponsiveness)
		addBool(query, "corporate_users", opts.corporateUsers)
		addBool(query, "marketing_mobile_number", opts.marketingMobileNumber)
		addBool(query, "sanction_details", opts.sanctionDetails)
		addBool(query, "limited_account", opts.limitedAccount)
		addBool(query, "compact", opts.compact)
		addInt(query, "limit", opts.limit)
		addInt(query, "offset", opts.offset)
	}

	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, string(ProjectsProjects), query, nil)
}

type ListSelfProjectsOptions struct {
	Status      *ProjectStatusType
	Role        *RoleType
	Types       []TypeType
	Query       *string
	SortField   *SortField
	ReverseSort *bool
	Recruiter   *bool
	Offset      *int
	Limit       *int
}

// Returns the logged in user’s projects/contests they either created or participated in (by bidding or submitting an entry).
// it maps to the `GET` `/projects/0.1/self` endpoint
func (s *ProjectsService) ListSelf(ctx context.Context, opts *ListSelfProjectsOptions) (*ListProjectsResponse, error) {
	query := url.Values{}
	if opts != nil {
		addEnum(query, "status", opts.Status)
		addEnum(query, "role", opts.Role)
		for _, val := range opts.Types {
			addEnum(query, "type[]", &val)
		}
		addString(query, "query", opts.Query)
		addEnum(query, "sort_field", opts.SortField)
		addBool(query, "reverse_sort", opts.ReverseSort)
		addBool(query, "recruiter", opts.Recruiter)
		addInt(query, "offset", opts.Offset)
		addInt(query, "limit", opts.Limit)
	}
	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, string(ProjectsSelf), query, nil)
}

type GetProjectOptions struct {
	FullDescription              *bool
	JobDetails                   *bool
	UpgradeDetails               *bool
	AttachmentDetails            *bool
	FileDetails                  *bool
	QualificationDetails         *bool
	SelectedBids                 *bool
	HiremeDetails                *bool
	UserDetails                  *bool
	InvitedFreelancerDetails     *bool
	RecommendedFreelancerDetails *bool
	HourlyDetails                *bool
	SupportSessionDetails        *bool
	LocationDetails              *bool
	NdaSignatureDetails          *bool
	ProjectCollaborationDetails  *bool
	TrackDetails                 *bool
	ProximityDetails             *bool
	ReviewAvailabilityDetails    *bool
	NegotiatedDetails            *bool
	DriveFileDetails             *bool
	NdaDetails                   *bool
	LocalDetails                 *bool
	EquipmentDetails             *bool
	ClientEngagementDetails      *bool
	ServiceOfferingDetails       *bool
	UserAvatar                   *bool
	UserCountryDetails           *bool
	UserProfileDescription       *bool
	UserDisplayInfo              *bool
	UserJobs                     *bool
	UserBalanceDetails           *bool
	UserQualificationDetails     *bool
	UserMembershipDetails        *bool
	UserFinancialDetails         *bool
	UserLocationDetails          *bool
	UserPortfolioDetails         *bool
	UserPreferredDetails         *bool
	UserBadgeDetails             *bool
	UserStatus                   *bool
	UserReputation               *bool
	UserEmployerReputation       *bool
	UserReputationExtra          *bool
	UserEmployerReputationExtra  *bool
	UserCoverImage               *bool
	UserPastCoverImage           *bool
	UserRecommendations          *bool
	UserResponsiveness           *bool
	CorporateUsers               *bool
	MarketingMobileNumber        *bool
	SanctionDetails              *bool
	LimitedAccount               *bool
	EquipmentGroupDetails        *bool
	Limit                        *int
	Offset                       *int
	Compact                      *bool
}

// Get information about a specific project. The full range of users projection options can be specified as part of this request by first setting theuser_detailsparameter to true.
// It maps to the `GET` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Get(ctx context.Context, projectID int64, opts *GetProjectOptions) (*GetProjectResponse, error) {
	query := url.Values{}
	if opts != nil {
		addBool(query, "full_description", opts.FullDescription)
		addBool(query, "job_details", opts.JobDetails)
		addBool(query, "upgrade_details", opts.UpgradeDetails)
		addBool(query, "attached_details", opts.AttachmentDetails)
		addBool(query, "file_details", opts.FileDetails)
		addBool(query, "qualification_details", opts.QualificationDetails)
		addBool(query, "selected_bid", opts.SelectedBids)
		addBool(query, "hireme_details", opts.HiremeDetails)
		addBool(query, "invited_freelancer_details", opts.InvitedFreelancerDetails)
		addBool(query, "recommended_freelancer_details", opts.RecommendedFreelancerDetails)
		addBool(query, "support_session_details", opts.SupportSessionDetails)
		addBool(query, "local_details", opts.LocalDetails)
		addBool(query, "nda_signature_details", opts.NdaSignatureDetails)
		addBool(query, "project_collaboration_details", opts.ProjectCollaborationDetails)
		addBool(query, "track_details", opts.TrackDetails)
		addBool(query, "proximity_details", opts.ProximityDetails)
		addBool(query, "review_availability_details", opts.ReviewAvailabilityDetails)
		addBool(query, "negotiated_details", opts.NegotiatedDetails)
		addBool(query, "drive_file_details", opts.DriveFileDetails)
		addBool(query, "nda_details", opts.NdaDetails)
		addBool(query, "local_details", opts.LocalDetails)
		addBool(query, "equipment_details", opts.EquipmentDetails)
		addBool(query, "client_engagement_details", opts.ClientEngagementDetails)
		addBool(query, "service_offering_details", opts.ServiceOfferingDetails)
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
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "user_reputation", opts.UserReputation)
		addBool(query, "user_employer_reputation", opts.UserEmployerReputation)
		addBool(query, "user_reputation_extra", opts.UserReputationExtra)
		addBool(query, "user_employer_reputation_extra", opts.UserEmployerReputationExtra)
		addBool(query, "user_cover_image", opts.UserCoverImage)
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

	path := fmt.Sprintf("%s/%s", string(ProjectsProjects), strconv.FormatInt(projectID, 10))
	return execute[*GetProjectResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

type SearchActiveProjectsOptions struct {
	Query                       *string
	ProjectTypes                []ProjectType
	ProjectUpgrades             []ProjectUpgradeType
	ContestUpgrades             []ContestUpgradeType
	MinAvgPrice                 *float64
	MaxAvgPrice                 *float64
	MinAvgHourlyRate            *float64
	MaxAvgHourlyRate            *float64
	MinPrice                    *float64
	MaxPrice                    *float64
	MinHourlyRate               *float64
	MaxHourlyRate               *float64
	Jobs                        []int64
	Countries                   []string
	Languages                   []string
	Latitude                    *float64
	Longitude                   *float64
	FromTime                    *int64
	ToTime                      *int64
	SortField                   *SortField
	ProjectIDs                  []int64
	TopRightLatitude            *float64
	TopRightLongitude           *float64
	BottomLeftLatitude          *float64
	BottomLeftLongitude         *float64
	ReverseSort                 *bool
	OrSearchQuery               *string
	HighlightPreTags            *string
	HighlightPostTags           *string
	FullDescription             *bool
	JobDetails                  *bool
	UpgradeDetails              *bool
	UserDetails                 *bool
	LocationDetails             *bool
	NDASignatureDetails         *bool
	ProjectCollaborationDetails *bool
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

// Searches for active projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/active` endpoint
func (s *ProjectsService) SearchActive(ctx context.Context, opts *SearchActiveProjectsOptions) (*ListProjectsResponse, error) {
	query := url.Values{}
	if opts != nil {
		addString(query, "query", opts.Query)
		for _, projectType := range opts.ProjectTypes {
			addEnum(query, "project_type[]", &projectType)
		}
		for _, projectUpgrade := range opts.ProjectUpgrades {
			addEnum(query, "project_upgrades[]", &projectUpgrade)
		}
		for _, contestUpgrade := range opts.ContestUpgrades {
			addEnum(query, "contest_upgrades[]", &contestUpgrade)
		}
		addFloat(query, "min_avg_price", opts.MinAvgPrice)
		addFloat(query, "max_avg_price", opts.MaxAvgPrice)
		addFloat(query, "max_avg_hourly_rate", opts.MaxAvgHourlyRate)
		addFloat(query, "min_price", opts.MinPrice)
		addFloat(query, "max_price", opts.MaxPrice)
		addFloat(query, "min_hourly_rate", opts.MinHourlyRate)
		for _, job := range opts.Jobs {
			addInt64(query, "jobs[]", &job)
		}
		for _, country := range opts.Countries {
			addString(query, "countries[]", &country)
		}
		for _, language := range opts.Languages {
			addString(query, "languages[]", &language)
		}

		addFloat(query, "latitude", opts.Latitude)
		addFloat(query, "longitude", opts.Longitude)
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addEnum(query, "sort_field", opts.SortField)

		for _, projectID := range opts.ProjectIDs {
			addInt64(query, "project_ids[]", &projectID)
		}

		addFloat(query, "top_right_latitude", opts.TopRightLatitude)
		addFloat(query, "top_right_longitude", opts.TopRightLongitude)
		addFloat(query, "bottom_left_latitude", opts.BottomLeftLatitude)
		addFloat(query, "bottom_left_longitude", opts.BottomLeftLongitude)

		addBool(query, "reverse_sort", opts.ReverseSort)
		addString(query, "or_search_query", opts.OrSearchQuery)

		addString(query, "highlight_pre_tags", opts.HighlightPreTags)
		addString(query, "highlight_post_tags", opts.HighlightPostTags)

		addBool(query, "full_description", opts.FullDescription)
		addBool(query, "job_details", opts.JobDetails)
		addBool(query, "upgrade_details", opts.UpgradeDetails)

		addBool(query, "user_status", opts.UserStatus)
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

		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)

	}

	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, string(ProjectsProjectsActive), query, nil)
}

type SearchAllProjectsOptions struct {
	Query                       *string
	ProjectTypes                []ProjectType
	ProjectUpgrades             []ProjectUpgradeType
	ContestUpgrades             []ContestUpgradeType
	MinAvgPrice                 *float64
	MaxAvgPrice                 *float64
	MinAvgHourlyRate            *float64
	MaxAvgHourlyRate            *float64
	MinPrice                    *float64
	MaxPrice                    *float64
	MinHourlyRate               *float64
	MaxHourlyRate               *float64
	Jobs                        []int64
	Countries                   []string
	Languages                   []string
	Latitude                    *float64
	Longitude                   *float64
	FromTime                    *int64
	ToTime                      *int64
	SortField                   *SortField
	BidAwardStatuses            []BidAwardStatus
	BidCompleteStatuses         []BidCompleteStatus
	ProjectStatuses             []ProjectFrontendStatus
	ProjectIDs                  []int64
	TopRightLatitude            *float64
	TopRightLongitude           *float64
	BottomLeftLatitude          *float64
	BottomLeftLongitude         *float64
	ReverseSort                 *bool
	OrSearchQuery               *string
	HighlightPreTags            *string
	HighlightPostTags           *string
	FullDescription             *bool
	JobDetails                  *bool
	UpgradeDetails              *bool
	UserDetails                 *bool
	LocationDetails             *bool
	NdaSignatureDetails         *bool
	ProjectCollaborationDetails *bool
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

// Searches for all projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/all` endpoint
func (s *ProjectsService) SearchAll(ctx context.Context, opts *SearchAllProjectsOptions) (*ListProjectsResponse, error) {
	query := url.Values{}

	if opts != nil {
		addString(query, "query", opts.Query)
		for _, projectType := range opts.ProjectTypes {
			addEnum(query, "project_types[]", &projectType)
		}
		for _, projectUpgrade := range opts.ProjectUpgrades {
			addEnum(query, "project_upgrades[]", &projectUpgrade)
		}
		for _, contestUpgrade := range opts.ContestUpgrades {
			addEnum(query, "contest_upgrades[]", &contestUpgrade)
		}
		addFloat(query, "min_avg_price", opts.MinAvgPrice)
		addFloat(query, "max_avg_price", opts.MaxAvgPrice)
		addFloat(query, "min_avg_hourly_rate", opts.MinAvgHourlyRate)
		addFloat(query, "max_avg_hourly_rate", opts.MaxAvgHourlyRate)
		addFloat(query, "min_price", opts.MinPrice)
		addFloat(query, "max_price", opts.MaxPrice)
		addFloat(query, "min_hourly_rate", opts.MinHourlyRate)
		addFloat(query, "max_hourly_rate", opts.MaxHourlyRate)
		for _, job := range opts.Jobs {
			addInt64(query, "jobs[]", &job)
		}
		for _, country := range opts.Countries {
			addString(query, "countries[]", &country)
		}
		for _, language := range opts.Languages {
			addString(query, "languages[]", &language)
		}
		addFloat(query, "latitude", opts.Latitude)
		addFloat(query, "longitude", opts.Longitude)
		addInt64(query, "from_time", opts.FromTime)
		addInt64(query, "to_time", opts.ToTime)
		addEnum(query, "sort_field", opts.SortField)
		for _, bidAwardStatus := range opts.BidAwardStatuses {
			addEnum(query, "bid_award_statuses[]", &bidAwardStatus)
		}
		for _, bidCompleteStatus := range opts.BidCompleteStatuses {
			addEnum(query, "bid_complete_statuses[]", &bidCompleteStatus)
		}
		for _, projectStatus := range opts.ProjectStatuses {
			addEnum(query, "project_statuses[]", &projectStatus)
		}
		for _, projectID := range opts.ProjectIDs {
			addInt64(query, "project_ids[]", &projectID)
		}
		addFloat(query, "top_right_latitude", opts.TopRightLatitude)
		addFloat(query, "top_right_longitude", opts.TopRightLongitude)
		addFloat(query, "bottom_left_latitude", opts.BottomLeftLatitude)
		addFloat(query, "bottom_left_longitude", opts.BottomLeftLongitude)
		addBool(query, "reverse_sort", opts.ReverseSort)
		addString(query, "or_search_query", opts.OrSearchQuery)
		addString(query, "highlight_pre_tags", opts.HighlightPreTags)
		addString(query, "highlight_post_tags", opts.HighlightPostTags)
		addBool(query, "full_description", opts.FullDescription)
		addBool(query, "job_details", opts.JobDetails)
		addBool(query, "upgrade_details", opts.UpgradeDetails)
		addBool(query, "user_status", opts.UserStatus)
		addBool(query, "nda_signature_details", opts.NdaSignatureDetails)
		addBool(query, "user_badge_details", opts.UserBadgeDetails)
		addBool(query, "user_qualification_details", opts.UserQualificationDetails)
		addBool(query, "user_membership_details", opts.UserMembershipDetails)
		addBool(query, "user_preferred_details", opts.UserPreferredDetails)
		addBool(query, "user_financial_details", opts.UserFinancialDetails)
		addBool(query, "user_location_details", opts.UserLocationDetails)
		addBool(query, "user_portfolio_details", opts.UserPortfolioDetails)
		addBool(query, "user_details", opts.UserDetails)
		addBool(query, "user_avatar", opts.UserAvatar)
		addBool(query, "user_country_details", opts.UserCountryDetails)
		addBool(query, "user_profile_description", opts.UserProfileDescription)
		addBool(query, "user_display_info", opts.UserDisplayInfo)
		addBool(query, "user_jobs", opts.UserJobs)

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
		addInt(query, "limit", opts.Limit)
		addInt(query, "offset", opts.Offset)
		addBool(query, "compact", opts.Compact)

	}

	return execute[*ListProjectsResponse](ctx, s.client, http.MethodGet, string(ProjectsProjectsAll), query, nil)
}

type InviteFreelancersBody struct {
	FreelancerID int64 `json:"freelancer_id"`
}

// Invites specific freelancers to bid on a project.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/invite` endpoint
func (s *ProjectsService) InviteFreelancer(ctx context.Context, projectID int64, b InviteFreelancersBody) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/invite", ProjectsProjects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodPost, path, nil, b)
}

type ListUpgradesFeesOptions struct {
	Currencies         []int
	Project            *int64
	FreeUpgradeDetails *bool
	TaxIncluded        *bool
}

// TODO: Refine the typed response with

// Returns the project upgrade fees for a given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
// It maps to the `GET` `/projects/0.1/projects/fees` endpoint
func (s *ProjectsService) ListUpgradesFees(ctx context.Context, opts *ListUpgradesFeesOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		if opts.Currencies != nil {
			for _, currency := range opts.Currencies {
				addInt(query, "currencies[]", &currency)
			}
		}
		addInt64(query, "project", opts.Project)
		addBool(query, "fee_upgrade_details", opts.FreeUpgradeDetails)
		addBool(query, "tax_included", opts.TaxIncluded)
	}

	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsProjectsFees), query, nil)
}

type ListProjectBidsOptions struct {
	ProjectID                   int64
	IsShortlisted               *bool
	Reputation                  *bool
	RecommendedBid              *bool
	ShortlistedBid              *bool
	Distance                    *bool
	UserDetails                 *bool
	ExpertGuarantees            *bool
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

// TODO: Refine the typed response with

// Returns bids for a single project. Employers will see bids in a sorted order, which begins with sponsored bids, then by bid ranking. Freelancers will receive the bid list ordered by date. Note: This method is expensive to compute so it is recommended that reputation and user projection options are not set.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids` endpoint
func (s *ProjectsService) ListBids(ctx context.Context, projectID int64, opts *ListProjectBidsOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		addBool(query, "is_shortlisted", opts.IsShortlisted)
		addBool(query, "reputation", opts.Reputation)
		addBool(query, "recommended_bid", opts.RecommendedBid)
		addBool(query, "shortlisted_bid", opts.ShortlistedBid)
		addBool(query, "distance", opts.Distance)
		addBool(query, "user_details", opts.UserDetails)
		addBool(query, "expert_guarantees", opts.ExpertGuarantees)
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

	path := fmt.Sprintf("%s/%s/bids", ProjectsProjects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, query, nil)
}

// TODO: refine with typed response

// Returns information for posting bids on a project.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids_info` endpoint
func (s *ProjectsService) GetBidInfo(ctx context.Context, projectID int64) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/bids_info", ProjectsProjects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

type ListProjectMilestonesOptions struct {
	Statuses                    []MilestoneStatus
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

// Returns a list of milestones on a project. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/milestones` endpoint
func (s *ProjectsService) ListMilestones(ctx context.Context, projectID int64, opts *ListProjectMilestonesOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/milestones", ProjectsProjects, strconv.FormatInt(projectID, 10))
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Statuses {
			addEnum(query, "statuses[]", &val)
		}
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

type ListProjectsMilestoneRequestsOptions struct {
	Statuses                    []MilestoneStatus
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

// Returns a list of milestone requests by freelancers for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/milestone_requests` endpoint
func (s *ProjectsService) ListMilestoneRequests(ctx context.Context, projectID int64, opts *ListProjectsMilestoneRequestsOptions) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/milestone_requests", ProjectsProjects, strconv.FormatInt(projectID, 10))
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.Statuses {
			addEnum(query, "statuses[]", &val)
		}
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

type GetHourlyContractInfoOptions struct {
	ProjectIDs        []int64
	BidderIDs         []int64
	HourlyContractIDs []int
	projectOwnerIDs   []int64
	billingDetails    *bool
	invoiceDetails    *bool
}

// TODO: refine with typed response

// Fetch the hourly contract matching the desired query.
// It maps to the `GET` `/projects/0.1/hourly_contract_info` endpoint
func (s *ProjectsService) GetHourlyContractInfo(ctx context.Context, opts *GetHourlyContractInfoOptions) (*RawResponse, error) {
	query := url.Values{}
	if opts != nil {
		for _, val := range opts.ProjectIDs {
			addInt64(query, "project_ids[]", &val)
		}
		for _, val := range opts.BidderIDs {
			addInt64(query, "bidder_ids[]", &val)
		}
		for _, val := range opts.HourlyContractIDs {
			addInt(query, "hourly_contract_ids[]", &val)
		}
		for _, val := range opts.projectOwnerIDs {
			query.Add("project_owner_ids[]", strconv.FormatInt(val, 10))
			addInt64(query, "project_owner_ids[]", &val)
		}
		addBool(query, "billing_details", opts.billingDetails)
		addBool(query, "invoice_details", opts.invoiceDetails)
	}
	return execute[*RawResponse](ctx, s.client, http.MethodGet, string(ProjectsHourlyContract), query, nil)
}

// TODO: refine with typed response

// Fetch the IP contract matching for the project id. If you are an employer it will return all of the contracts, ELSE we will return contract details specific to the logged-in user.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/ip_contract_info` endpoint
func (s *ProjectsService) GetIPContractInfo(ctx context.Context, projectID int64) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s/ip_contract_info", ProjectsProjects, strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodGet, path, nil, nil)
}

// Delete a project by id. Only projects that are in pending or rejected states may be deleted. This will close the project and remove its visibility.
// it maps to the `DELETE` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Delete(ctx context.Context, projectID int64) (*RawResponse, error) {
	path := fmt.Sprintf("%s/%s", string(ProjectsProjects), strconv.FormatInt(projectID, 10))
	return execute[*RawResponse](ctx, s.client, http.MethodDelete, path, nil, nil)
}

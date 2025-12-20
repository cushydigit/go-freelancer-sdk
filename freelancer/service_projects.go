package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type createProject struct {
	client *Client
}

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

func (s *createProject) Do(ctx context.Context, b CreateProjectBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_PROJECTS),
		body:     bytes.NewBuffer(m),
	}

	return execute[*RawResponse](ctx, s.client, r)
}

type actionProject struct {
	client    *Client
	projectID int64
}

// TODO: check this service functionality so the projectActionBody should have better fields
// check : https://developers.freelancer.com/docs/projects/projects#projects-put
// Performs an action on a project.
type ProjectActionBody struct {
	ProjectID int64         `json:"project_id"`
	Action    ProjectAction `json:"action"`
}

func (s *actionProject) Do(ctx context.Context, b ProjectActionBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
		body:     bytes.NewBuffer(m),
	}

	return execute[*RawResponse](ctx, s.client, r)
}

type listProjects struct {
	client                       *Client
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

// Return information about multiple projects. Will be ordered by descending submit data (newest-to-oldest)
func (s *listProjects) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS),
	}
	for _, project := range s.projects {
		r.addParam("projects[]", project)
	}
	for _, owner := range s.owners {
		r.addParam("owners[]", owner)
	}
	for _, bidder := range s.bidders {
		r.addParam("bidders[]", bidder)
	}
	for _, seuUrl := range s.seoUrls {
		r.setParam("seo_urls[]", seuUrl)
	}
	if s.fromTime != nil {
		r.setParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.setParam("to_time", *s.toTime)
	}
	if s.frontendProjectStatuses != nil {
		r.setParam("frontend_project_statuses[]", s.frontendProjectStatuses)
	}
	if s.team != nil {
		r.setParam("team", *s.team)
	}
	if s.isNonHireMe != nil {
		r.setParam("is_none_hire_me", *s.isNonHireMe)
	}
	if s.hasMilestone != nil {
		r.setParam("has_milestone", *s.hasMilestone)
	}
	if s.count != nil {
		r.setParam("count", *s.count)
	}
	if s.fullDescription != nil {
		r.setParam("full_description", *s.fullDescription)
	}
	if s.jobDetails != nil {
		r.setParam("job_details", *s.jobDetails)
	}
	if s.upgradeDetails != nil {
		r.setParam("upgrade_details", *s.upgradeDetails)
	}
	if s.attachmentDetails != nil {
		r.setParam("attachment_details", *s.attachmentDetails)
	}
	if s.fileDetails != nil {
		r.setParam("file_details", *s.fileDetails)
	}
	if s.qualificationDetails != nil {
		r.setParam("qualification_details", *s.qualificationDetails)
	}
	if s.selectedBids != nil {
		r.setParam("selected_bids", *s.selectedBids)
	}
	if s.hiremeDetails != nil {
		r.setParam("hireme_details", *s.hiremeDetails)
	}
	if s.userDetails != nil {
		r.setParam("user_details", *s.userDetails)
	}
	if s.invitedFreelancerDetails != nil {
		r.setParam("invited_freelancer_details", *s.invitedFreelancerDetails)
	}
	if s.recommendedFreelancerDetails != nil {
		r.setParam("recommended_freelancer_details", *s.recommendedFreelancerDetails)
	}
	if s.supportSessionDetails != nil {
		r.setParam("support_session_details", *s.supportSessionDetails)
	}
	if s.localDetails != nil {
		r.setParam("local_details", *s.localDetails)
	}
	if s.ndaSignatureDetails != nil {
		r.setParam("nda_signature_details", *s.ndaSignatureDetails)
	}
	if s.projectCollaborationDetails != nil {
		r.setParam("project_collaboration_details", *s.projectCollaborationDetails)
	}
	if s.proximityDetails != nil {
		r.setParam("proximity_details", *s.proximityDetails)
	}
	if s.reviewAvailabilityDetails != nil {
		r.setParam("review_availability_details", *s.reviewAvailabilityDetails)
	}
	if s.negotiatedDetails != nil {
		r.setParam("negotiated_details", *s.negotiatedDetails)
	}
	if s.driveFileDetails != nil {
		r.setParam("drive_file_details", *s.driveFileDetails)
	}
	if s.ndaDetails != nil {
		r.setParam("nda_details", *s.ndaDetails)
	}
	if s.localDetails != nil {
		r.setParam("local_details", *s.localDetails)
	}
	if s.equipmentDetails != nil {
		r.setParam("equipment_details", *s.equipmentDetails)
	}
	if s.clientEngagementDetails != nil {
		r.setParam("client_engagement_details", *s.clientEngagementDetails)
	}
	if s.serviceOfferingDetails != nil {
		r.setParam("service_offering_details", *s.serviceOfferingDetails)
	}
	if s.userAvatar != nil {
		r.setParam("user_avatar", *s.userAvatar)
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
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.setParam("compact", *s.compact)
	}

	return execute[*ListProjectsResponse](ctx, s.client, r)
}

// SetProjects sets the list of project IDs to filter.
func (s *listProjects) SetProjects(projects []int64) *listProjects {
	s.projects = projects
	return s
}

// SetOwners sets the list of owner user IDs to filter.
func (s *listProjects) SetOwners(owners []int64) *listProjects {
	s.owners = owners
	return s
}

// SetBidders sets the list of bidder user IDs to filter.
func (s *listProjects) SetBidders(bidders []int64) *listProjects {
	s.bidders = bidders
	return s
}

// SetSeoUrls sets the list of SEO URLs to filter.
func (s *listProjects) SetSeoUrls(seoUrls []string) *listProjects {
	s.seoUrls = seoUrls
	return s
}

// SetFrontendProjectStatuses sets the frontend project status filters.
func (s *listProjects) SetFrontendProjectStatuses(val []string) *listProjects {
	s.frontendProjectStatuses = val
	return s
}

// SetFromTime sets the start time filter (Unix timestamp).
func (s *listProjects) SetFromTime(val int64) *listProjects {
	s.fromTime = &val
	return s
}

// SetToTime sets the end time filter (Unix timestamp).
func (s *listProjects) SetToTime(val int64) *listProjects {
	s.toTime = &val
	return s
}

// SetFullDescription enables or disables full project description.
func (s *listProjects) SetFullDescription(val bool) *listProjects {
	s.fullDescription = &val
	return s
}

// SetJobDetails enables or disables inclusion of job details.
func (s *listProjects) SetJobDetails(val bool) *listProjects {
	s.jobDetails = &val
	return s
}

// SetUpgradeDetails enables or disables upgrade details.
func (s *listProjects) SetUpgradeDetails(val bool) *listProjects {
	s.upgradeDetails = &val
	return s
}

// SetAttachmentDetails enables or disables attachment details.
func (s *listProjects) SetAttachmentDetails(val bool) *listProjects {
	s.attachmentDetails = &val
	return s
}

// SetFileDetails enables or disables file details.
func (s *listProjects) SetFileDetails(val bool) *listProjects {
	s.fileDetails = &val
	return s
}

// SetQualificationDetails enables or disables qualification details.
func (s *listProjects) SetQualificationDetails(val bool) *listProjects {
	s.qualificationDetails = &val
	return s
}

// SetSelectedBids enables or disables retrieval of selected bids.
func (s *listProjects) SetSelectedBids(val bool) *listProjects {
	s.selectedBids = &val
	return s
}

// SetHiremeDetails enables or disables hireme details.
func (s *listProjects) SetHiremeDetails(val bool) *listProjects {
	s.hiremeDetails = &val
	return s
}

// SetUserDetails enables or disables user details.
func (s *listProjects) SetUserDetails(val bool) *listProjects {
	s.userDetails = &val
	return s
}

// SetInvitedFreelancerDetails enables or disables invited freelancer details.
func (s *listProjects) SetInvitedFreelancerDetails(val bool) *listProjects {
	s.invitedFreelancerDetails = &val
	return s
}

// SetRecommendedFreelancerDetails enables or disables recommended freelancer details.
func (s *listProjects) SetRecommendedFreelancerDetails(val bool) *listProjects {
	s.recommendedFreelancerDetails = &val
	return s
}

// SetSupportSessionDetails enables or disables support session details.
func (s *listProjects) SetSupportSessionDetails(val bool) *listProjects {
	s.supportSessionDetails = &val
	return s
}

// SetLocationDetails enables or disables location details.
func (s *listProjects) SetLocationDetails(val bool) *listProjects {
	s.locationDetails = &val
	return s
}

// SetNdaSignatureDetails enables or disables NDA signature details.
func (s *listProjects) SetNdaSignatureDetails(val bool) *listProjects {
	s.ndaSignatureDetails = &val
	return s
}

// SetDriveFileDetails enables or disables drive file details.
func (s *listProjects) SetDriveFileDetails(val bool) *listProjects {
	s.driveFileDetails = &val
	return s
}

// SetNdaDetails enables or disables NDA details.
func (s *listProjects) SetNdaDetails(val bool) *listProjects {
	s.ndaDetails = &val
	return s
}

// SetLocalDetails enables or disables local details.
func (s *listProjects) SetLocalDetails(val bool) *listProjects {
	s.localDetails = &val
	return s
}

// SetEquipmentDetails enables or disables equipment details.
func (s *listProjects) SetEquipmentDetails(val bool) *listProjects {
	s.equipmentDetails = &val
	return s
}

// SetClientEngagementDetails enables or disables client engagement details.
func (s *listProjects) SetClientEngagementDetails(val bool) *listProjects {
	s.clientEngagementDetails = &val
	return s
}

// SetUserResponsiveness enables or disables user responsiveness details.
func (s *listProjects) SetUserResponsiveness(val bool) *listProjects {
	s.userResponsiveness = &val
	return s
}

// SetServiceOfferingDetails enables or disables service offering details.
func (s *listProjects) SetServiceOfferingDetails(val bool) *listProjects {
	s.serviceOfferingDetails = &val
	return s
}

// SetCorporateUsers enables or disables corporate user filter.
func (s *listProjects) SetCorporateUsers(val bool) *listProjects {
	s.corporateUsers = &val
	return s
}

// SetIsNonHireMe sets whether to exclude HireMe projects.
func (s *listProjects) SetIsNonHireMe(val bool) *listProjects {
	s.isNonHireMe = &val
	return s
}

// SetHasMilestone enables or disables filtering projects with milestones.
func (s *listProjects) SetHasMilestone(val bool) *listProjects {
	s.hasMilestone = &val
	return s
}

// SetCount enables or disables counting results.
func (s *listProjects) SetCount(val bool) *listProjects {
	s.count = &val
	return s
}

// SetTeam enables or disables team details.
func (s *listProjects) SetTeam(val bool) *listProjects {
	s.team = &val
	return s
}

// SetCompact enables or disables compact result format.
func (s *listProjects) SetCompact(val bool) *listProjects {
	s.compact = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *listProjects) SetLimit(val int) *listProjects {
	s.limit = &val
	return s
}

// SetOffset sets the offset for pagination.
func (s *listProjects) SetOffset(val int) *listProjects {
	s.offset = &val
	return s
}

// SetProximityDetails enables or disables proximity details.
func (s *listProjects) SetProximityDetails(val bool) *listProjects {
	s.proximityDetails = &val
	return s
}

// SetReviewAvailabilityDetails enables or disables review availability details.
func (s *listProjects) SetReviewAvailabilityDetails(val bool) *listProjects {
	s.reviewAvailabilityDetails = &val
	return s
}

// SetNegotiatedDetails enables or disables negotiated details.
func (s *listProjects) SetNegotiatedDetails(val bool) *listProjects {
	s.negotiatedDetails = &val
	return s
}

// SetUserAvatar enables or disables user avatar inclusion.
func (s *listProjects) SetUserAvatar(val bool) *listProjects {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails enables or disables user country details.
func (s *listProjects) SetUserCountryDetails(val bool) *listProjects {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription enables or disables user profile description.
func (s *listProjects) SetUserProfileDescription(val bool) *listProjects {
	s.userProfileDescription = &val
	return s
}

// SetProjectCollaborationDetails enables or disables project collaboration details.
func (s *listProjects) SetProjectCollaborationDetails(val bool) *listProjects {
	s.projectCollaborationDetails = &val
	return s
}

// SetUserDisplayInfo enables or disables user display information.
func (s *listProjects) SetUserDisplayInfo(val bool) *listProjects {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs enables or disables user jobs details.
func (s *listProjects) SetUserJobs(val bool) *listProjects {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails enables or disables user balance details.
func (s *listProjects) SetUserBalanceDetails(val bool) *listProjects {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails enables or disables user qualification details.
func (s *listProjects) SetUserQualificationDetails(val bool) *listProjects {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails enables or disables user membership details.
func (s *listProjects) SetUserMembershipDetails(val bool) *listProjects {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails enables or disables user financial details.
func (s *listProjects) SetUserFinancialDetails(val bool) *listProjects {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails enables or disables user location details.
func (s *listProjects) SetUserLocationDetails(val bool) *listProjects {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails enables or disables user portfolio details.
func (s *listProjects) SetUserPortfolioDetails(val bool) *listProjects {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails enables or disables user preferred details.
func (s *listProjects) SetUserPreferredDetails(val bool) *listProjects {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails enables or disables user badge details.
func (s *listProjects) SetUserBadgeDetails(val bool) *listProjects {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus enables or disables user status details.
func (s *listProjects) SetUserStatus(val bool) *listProjects {
	s.userStatus = &val
	return s
}

// SetUserReputation enables or disables user reputation details.
func (s *listProjects) SetUserReputation(val bool) *listProjects {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation enables or disables user employer reputation details.
func (s *listProjects) SetUserEmployerReputation(val bool) *listProjects {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra enables or disables extra user reputation details.
func (s *listProjects) SetUserReputationExtra(val bool) *listProjects {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra enables or disables extra user employer reputation details.
func (s *listProjects) SetUserEmployerReputationExtra(val bool) *listProjects {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage enables or disables user cover image.
func (s *listProjects) SetUserCoverImage(val bool) *listProjects {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage enables or disables past user cover image.
func (s *listProjects) SetUserPastCoverImage(val bool) *listProjects {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations enables or disables user recommendations.
func (s *listProjects) SetUserRecommendations(val bool) *listProjects {
	s.userRecommendations = &val
	return s
}

// SetMarketingMobileNumber enables or disables marketing mobile number.
func (s *listProjects) SetMarketingMobileNumber(val bool) *listProjects {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables or disables sanction details.
func (s *listProjects) SetSanctionDetails(val bool) *listProjects {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables or disables limited account details.
func (s *listProjects) SetLimitedAccount(val bool) *listProjects {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails enables or disables equipment group details.
func (s *listProjects) SetEquipmentGroupDetails(val bool) *listProjects {
	s.equipmentGroupDetails = &val
	return s
}

type listSelfProjects struct {
	client      *Client
	status      *ProjectStatusType
	role        *RoleType
	types       []TypeType
	query       *string
	sortField   *SortField
	reverseSort *bool
	recruiter   *bool
	offset      *int
	limit       *int
}

func (s *listSelfProjects) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_SELF),
	}
	if s.status != nil {
		r.addParam("status", *s.status)
	}
	if s.role != nil {
		r.addParam("role", *s.role)
	}
	for _, val := range s.types {
		r.addParam("types[]", val)
	}
	if s.query != nil {
		r.addParam("query", *s.query)
	}
	if s.sortField != nil {
		r.addParam("sort_field", *s.sortField)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.recruiter != nil {
		r.addParam("recruiter", *s.recruiter)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}

	return execute[*ListProjectsResponse](ctx, s.client, r)

}

// SetStatus sets the status filter for the projects/contests.
// Example: awarded. Valid values: active, all, awarded, open, past.
func (s *listSelfProjects) SetStatus(val ProjectStatusType) *listSelfProjects {
	s.status = &val
	return s
}

// SetRole sets the role of the user in the projects/contests.
// Example: employer. Valid values: freelancer, employer.
func (s *listSelfProjects) SetRole(val RoleType) *listSelfProjects {
	s.role = &val
	return s
}

// SetTypes sets the type filter for the results, e.g., projects, contests, or both.
// Default: projects, contests. Example: projects.
func (s *listSelfProjects) SetTypes(vals []TypeType) *listSelfProjects {
	s.types = vals
	return s
}

// SetQuery sets the search terms to filter projects/contests by title, description, or usernames.
// Example: build website.
func (s *listSelfProjects) SetQuery(val string) *listSelfProjects {
	s.query = &val
	return s
}

// SetSortField sets the field by which to sort the results.
// Default: project_id. Valid values: project_id, submitdate.
func (s *listSelfProjects) SetSortField(val SortField) *listSelfProjects {
	s.sortField = &val
	return s
}

// SetReverseSort sets whether the sorting should be in ascending order.
// Example: true. If true, results are sorted in ascending order.
func (s *listSelfProjects) SetReverseSort(val bool) *listSelfProjects {
	s.reverseSort = &val
	return s
}

// SetRecruiter sets the filter to return only recruiter projects.
// Example: true. If true, only recruiter projects are returned.
func (s *listSelfProjects) SetRecruiter(val bool) *listSelfProjects {
	s.recruiter = &val
	return s
}

// SetOffset sets the number of results to skip (pagination).
// Default: 0. Example: 30.
func (s *listSelfProjects) SetOffset(val int) *listSelfProjects {
	s.offset = &val
	return s
}

// SetLimit sets the maximum number of results to return.
// Default: 100. Example: 10.
func (s *listSelfProjects) SetLimit(val int) *listSelfProjects {
	s.limit = &val
	return s
}

type getProjectByID struct {
	client                       *Client
	projectID                    int64
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
	hourlyDetails                *bool
	supportSessionDetails        *bool
	locationDetails              *bool
	ndaSignatureDetails          *bool
	projectCollaborationDetails  *bool
	trackDetails                 *bool
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

func (s *getProjectByID) Do(ctx context.Context) (*GetProjectResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}
	if s.fullDescription != nil {
		r.setParam("full_description", *s.fullDescription)
	}
	if s.jobDetails != nil {
		r.setParam("job_details", *s.jobDetails)
	}
	if s.upgradeDetails != nil {
		r.setParam("upgrade_details", *s.upgradeDetails)
	}
	if s.attachmentDetails != nil {
		r.setParam("attachment_details", *s.attachmentDetails)
	}
	if s.fileDetails != nil {
		r.setParam("file_details", *s.fileDetails)
	}
	if s.qualificationDetails != nil {
		r.setParam("qualification_details", *s.qualificationDetails)
	}
	if s.selectedBids != nil {
		r.setParam("selected_bids", *s.selectedBids)
	}
	if s.hiremeDetails != nil {
		r.setParam("hireme_details", *s.hiremeDetails)
	}
	if s.userDetails != nil {
		r.setParam("user_details", *s.userDetails)
	}
	if s.invitedFreelancerDetails != nil {
		r.setParam("invited_freelancer_details", *s.invitedFreelancerDetails)
	}
	if s.recommendedFreelancerDetails != nil {
		r.setParam("recommended_freelancer_details", *s.recommendedFreelancerDetails)
	}
	if s.supportSessionDetails != nil {
		r.setParam("support_session_details", *s.supportSessionDetails)
	}
	if s.localDetails != nil {
		r.setParam("local_details", *s.localDetails)
	}
	if s.ndaSignatureDetails != nil {
		r.setParam("nda_signature_details", *s.ndaSignatureDetails)
	}
	if s.projectCollaborationDetails != nil {
		r.setParam("project_collaboration_details", *s.projectCollaborationDetails)
	}
	if s.trackDetails != nil {
		r.setParam("track_details", *s.trackDetails)
	}
	if s.proximityDetails != nil {
		r.setParam("proximity_details", *s.proximityDetails)
	}
	if s.reviewAvailabilityDetails != nil {
		r.setParam("review_availability_details", *s.reviewAvailabilityDetails)
	}
	if s.negotiatedDetails != nil {
		r.setParam("negotiated_details", *s.negotiatedDetails)
	}
	if s.driveFileDetails != nil {
		r.setParam("drive_file_details", *s.driveFileDetails)
	}
	if s.ndaDetails != nil {
		r.setParam("nda_details", *s.ndaDetails)
	}
	if s.localDetails != nil {
		r.setParam("local_details", *s.localDetails)
	}
	if s.equipmentDetails != nil {
		r.setParam("equipment_details", *s.equipmentDetails)
	}
	if s.clientEngagementDetails != nil {
		r.setParam("client_engagement_details", *s.clientEngagementDetails)
	}
	if s.serviceOfferingDetails != nil {
		r.setParam("service_offering_details", *s.serviceOfferingDetails)
	}
	if s.userAvatar != nil {
		r.setParam("user_avatar", *s.userAvatar)
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
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.setParam("compact", *s.compact)
	}

	return execute[*GetProjectResponse](ctx, s.client, r)

}

// SetFullDescription sets whether to return the full project description.
func (s *getProjectByID) SetFullDescription(val bool) *getProjectByID {
	s.fullDescription = &val
	return s
}

// SetJobDetails sets whether to include job information in the project response.
func (s *getProjectByID) SetJobDetails(val bool) *getProjectByID {
	s.jobDetails = &val
	return s
}

// SetUpgradeDetails sets whether to include upgrade information for the project.
func (s *getProjectByID) SetUpgradeDetails(val bool) *getProjectByID {
	s.upgradeDetails = &val
	return s
}

// SetAttachmentDetails sets whether to include attachment details of the project.
func (s *getProjectByID) SetAttachmentDetails(val bool) *getProjectByID {
	s.attachmentDetails = &val
	return s
}

// SetFileDetails sets whether to include files shared between employer and freelancer.
func (s *getProjectByID) SetFileDetails(val bool) *getProjectByID {
	s.fileDetails = &val
	return s
}

// SetQualificationDetails sets whether to include exams needed to qualify for the project.
func (s *getProjectByID) SetQualificationDetails(val bool) *getProjectByID {
	s.qualificationDetails = &val
	return s
}

// SetSelectedBids sets whether to return bids that have been awarded or are pending.
func (s *getProjectByID) SetSelectedBids(val bool) *getProjectByID {
	s.selectedBids = &val
	return s
}

// SetHiremeDetails sets whether to include information about hireme offers.
func (s *getProjectByID) SetHiremeDetails(val bool) *getProjectByID {
	s.hiremeDetails = &val
	return s
}

// SetUserDetails sets whether to return basic user information.
func (s *getProjectByID) SetUserDetails(val bool) *getProjectByID {
	s.userDetails = &val
	return s
}

// SetInvitedFreelancerDetails sets whether to include a list of invited freelancer user IDs.
func (s *getProjectByID) SetInvitedFreelancerDetails(val bool) *getProjectByID {
	s.invitedFreelancerDetails = &val
	return s
}

// SetRecommendedFreelancerDetails sets whether to include a list of recommended freelancer IDs (requires project ownership and job_details set).
func (s *getProjectByID) SetRecommendedFreelancerDetails(val bool) *getProjectByID {
	s.recommendedFreelancerDetails = &val
	return s
}

// SetHourlyDetails sets whether to include hourly contract information indexed by bidder user ID.
func (s *getProjectByID) SetHourlyDetails(val bool) *getProjectByID {
	s.hourlyDetails = &val
	return s
}

// SetSupportSessionDetails sets whether to include a list of support sessions for the project.
func (s *getProjectByID) SetSupportSessionDetails(val bool) *getProjectByID {
	s.supportSessionDetails = &val
	return s
}

// SetLocationDetails sets whether to include information about the project's location.
func (s *getProjectByID) SetLocationDetails(val bool) *getProjectByID {
	s.locationDetails = &val
	return s
}

// SetNdaSignatureDetails sets whether to include a list of users who have signed an NDA for the project.
func (s *getProjectByID) SetNdaSignatureDetails(val bool) *getProjectByID {
	s.ndaSignatureDetails = &val
	return s
}

// SetDriveFileDetails sets whether the project has any drive files.
func (s *getProjectByID) SetDriveFileDetails(val bool) *getProjectByID {
	s.driveFileDetails = &val
	return s
}

// SetNdaDetails sets whether to return NDA details including hidden description and signatures.
func (s *getProjectByID) SetNdaDetails(val bool) *getProjectByID {
	s.ndaDetails = &val
	return s
}

// SetLocalDetails sets whether to include details about local projects.
func (s *getProjectByID) SetLocalDetails(val bool) *getProjectByID {
	s.localDetails = &val
	return s
}

// SetEquipmentDetails sets whether to include equipment-related details attached to the user.
func (s *getProjectByID) SetEquipmentDetails(val bool) *getProjectByID {
	s.equipmentDetails = &val
	return s
}

// SetClientEngagementDetails sets whether to include client engagement-related details.
func (s *getProjectByID) SetClientEngagementDetails(val bool) *getProjectByID {
	s.clientEngagementDetails = &val
	return s
}

// SetUserResponsiveness sets whether to include the responsiveness score of the selected user(s).
func (s *getProjectByID) SetUserResponsiveness(val bool) *getProjectByID {
	s.userResponsiveness = &val
	return s
}

// SetServiceOfferingDetails sets whether to include service offering details of the user(s).
func (s *getProjectByID) SetServiceOfferingDetails(val bool) *getProjectByID {
	s.serviceOfferingDetails = &val
	return s
}

// SetCorporateUsers sets whether to return corporate accounts created or founded by the user(s).
func (s *getProjectByID) SetCorporateUsers(val bool) *getProjectByID {
	s.corporateUsers = &val
	return s
}

// SetCompact sets whether to strip all null and empty values from the response.
func (s *getProjectByID) SetCompact(val bool) *getProjectByID {
	s.compact = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *getProjectByID) SetLimit(val int) *getProjectByID {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *getProjectByID) SetOffset(val int) *getProjectByID {
	s.offset = &val
	return s
}

// SetProximityDetails sets whether to include project proximity information.
func (s *getProjectByID) SetProximityDetails(val bool) *getProjectByID {
	s.proximityDetails = &val
	return s
}

// SetReviewAvailabilityDetails sets whether to include review availability information for the project.
func (s *getProjectByID) SetReviewAvailabilityDetails(val bool) *getProjectByID {
	s.reviewAvailabilityDetails = &val
	return s
}

// SetNegotiatedDetails sets whether to include information about negotiated offers.
func (s *getProjectByID) SetNegotiatedDetails(val bool) *getProjectByID {
	s.negotiatedDetails = &val
	return s
}

// SetUserAvatar sets whether to include the avatar of the selected user(s).
func (s *getProjectByID) SetUserAvatar(val bool) *getProjectByID {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether to include the country flag/code of selected user(s).
func (s *getProjectByID) SetUserCountryDetails(val bool) *getProjectByID {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether to include the profile blurb of selected user(s).
func (s *getProjectByID) SetUserProfileDescription(val bool) *getProjectByID {
	s.userProfileDescription = &val
	return s
}

// SetProjectCollaborationDetails sets whether to include a list of project collaborators.
func (s *getProjectByID) SetProjectCollaborationDetails(val bool) *getProjectByID {
	s.projectCollaborationDetails = &val
	return s
}

// SetTrackDetails sets whether to include a list of track IDs associated with the project.
func (s *getProjectByID) SetTrackDetails(val bool) *getProjectByID {
	s.trackDetails = &val
	return s
}

// SetUserDisplayInfo sets whether to include the display name of selected user(s).
func (s *getProjectByID) SetUserDisplayInfo(val bool) *getProjectByID {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether to include the jobs of selected user(s).
func (s *getProjectByID) SetUserJobs(val bool) *getProjectByID {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether to include currency balance information of selected user(s).
func (s *getProjectByID) SetUserBalanceDetails(val bool) *getProjectByID {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether to include qualification exams completed by selected user(s).
func (s *getProjectByID) SetUserQualificationDetails(val bool) *getProjectByID {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether to include membership information of selected user(s).
func (s *getProjectByID) SetUserMembershipDetails(val bool) *getProjectByID {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether to include financial information of selected user(s).
func (s *getProjectByID) SetUserFinancialDetails(val bool) *getProjectByID {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether to include location information of selected user(s).
func (s *getProjectByID) SetUserLocationDetails(val bool) *getProjectByID {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether to include portfolio information of selected user(s).
func (s *getProjectByID) SetUserPortfolioDetails(val bool) *getProjectByID {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether to include preferred information of selected user(s).
func (s *getProjectByID) SetUserPreferredDetails(val bool) *getProjectByID {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether to include badges earned by selected user(s).
func (s *getProjectByID) SetUserBadgeDetails(val bool) *getProjectByID {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether to include additional status information about selected user(s).
func (s *getProjectByID) SetUserStatus(val bool) *getProjectByID {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether to include freelancer reputation of selected user(s).
func (s *getProjectByID) SetUserReputation(val bool) *getProjectByID {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether to include employer reputation of selected user(s).
func (s *getProjectByID) SetUserEmployerReputation(val bool) *getProjectByID {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether to include full freelancer reputation of selected user(s).
func (s *getProjectByID) SetUserReputationExtra(val bool) *getProjectByID {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether to include full employer reputation of selected user(s).
func (s *getProjectByID) SetUserEmployerReputationExtra(val bool) *getProjectByID {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether to include the profile picture of selected user(s).
func (s *getProjectByID) SetUserCoverImage(val bool) *getProjectByID {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether to include previous profile pictures of selected user(s).
func (s *getProjectByID) SetUserPastCoverImage(val bool) *getProjectByID {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether to include recommendation counts of selected user(s).
func (s *getProjectByID) SetUserRecommendations(val bool) *getProjectByID {
	s.userRecommendations = &val
	return s
}

// SetMarketingMobileNumber sets whether to include the mobile number used by recruiters to contact the user.
func (s *getProjectByID) SetMarketingMobileNumber(val bool) *getProjectByID {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to include the end timestamp of any sanctions given to the user.
func (s *getProjectByID) SetSanctionDetails(val bool) *getProjectByID {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to include the limit account status of the user.
func (s *getProjectByID) SetLimitedAccount(val bool) *getProjectByID {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to include equipment groups and items attached to the user.
func (s *getProjectByID) SetEquipmentGroupDetails(val bool) *getProjectByID {
	s.equipmentGroupDetails = &val
	return s
}

type searchActiveProjects struct {
	client                      *Client
	query                       *string
	projectTypes                []ProjectType
	projectUpgrades             []ProjectUpgradeType
	contestUpgrades             []ContestUpgradeType
	minAvgPrice                 *float64
	maxAvgPrice                 *float64
	minAvgHourlyRate            *float64
	maxAvgHourlyRate            *float64
	minPrice                    *float64
	maxPrice                    *float64
	minHourlyRate               *float64
	maxHourlyRate               *float64
	jobs                        []int32
	countries                   []string
	languages                   []string
	latitude                    *float64
	longitude                   *float64
	fromTime                    *int64
	toTime                      *int64
	sortField                   *SortField
	projectIDs                  []int64
	topRightLatitude            *float64
	topRightLongitude           *float64
	bottomLeftLatitude          *float64
	bottomLeftLongitude         *float64
	reverseSort                 *bool
	orSearchQuery               *string
	highlightPreTags            *string
	highlightPostTags           *string
	fullDescription             *bool
	jobDetails                  *bool
	upgradeDetails              *bool
	userDetails                 *bool
	locationDetails             *bool
	ndaSignatureDetails         *bool
	projectCollaborationDetails *bool
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
	compact                     *bool
}

// Searches for active projects matching the desired query
func (s *searchActiveProjects) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS_ACTIVE),
	}
	if s.query != nil {
		r.addParam("query", s.query)
	}
	for _, projectType := range s.projectTypes {
		r.addParam("project_types[]", projectType)
	}
	for _, projectUpgrade := range s.projectUpgrades {
		r.addParam("project_upgrades[]", projectUpgrade)
	}
	for _, contestUpgrade := range s.contestUpgrades {
		r.addParam("contest_upgrades[]", contestUpgrade)
	}
	if s.contestUpgrades != nil {
		r.addParam("contest_upgrades", s.contestUpgrades)
	}
	if s.minAvgPrice != nil {
		r.addParam("min_avg_price", *s.minAvgPrice)
	}
	if s.maxAvgPrice != nil {
		r.addParam("max_avg_price", *s.maxAvgPrice)
	}
	if s.minAvgHourlyRate != nil {
		r.addParam("min_avg_hourly_rate", *s.minAvgHourlyRate)
	}
	if s.maxAvgHourlyRate != nil {
		r.addParam("max_avg_hourly_rate", *s.maxAvgHourlyRate)
	}
	if s.minPrice != nil {
		r.addParam("min_price", *s.minPrice)
	}
	if s.maxPrice != nil {
		r.addParam("max_price", *s.maxPrice)
	}
	if s.minHourlyRate != nil {
		r.addParam("min_hourly_rate", *s.minHourlyRate)
	}
	if s.maxHourlyRate != nil {
		r.addParam("max_hourly_rate", *s.maxHourlyRate)
	}
	for _, job := range s.jobs {
		r.addParam("jobs[]", job)
	}
	for _, country := range s.countries {
		r.addParam("countries[]", country)
	}
	for _, language := range s.languages {
		r.addParam("languages[]", language)
	}
	if s.latitude != nil {
		r.addParam("latitude", *s.latitude)
	}
	if s.longitude != nil {
		r.addParam("longitude", *s.longitude)
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
	for _, projectID := range s.projectIDs {
		r.addParam("project_ids[]", projectID)
	}
	if s.topRightLatitude != nil {
		r.addParam("top_right_latitude", *s.topRightLatitude)
	}
	if s.topRightLongitude != nil {
		r.addParam("top_right_longitude", *s.topRightLongitude)
	}
	if s.bottomLeftLatitude != nil {
		r.addParam("bottom_left_latitude", *s.bottomLeftLatitude)
	}
	if s.bottomLeftLongitude != nil {
		r.addParam("bottom_left_longitude", *s.bottomLeftLongitude)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.orSearchQuery != nil {
		r.addParam("or_search_query", *s.orSearchQuery)
	}
	if s.highlightPreTags != nil {
		r.addParam("highlight_pre_tags", *s.highlightPreTags)
	}
	if s.highlightPostTags != nil {
		r.addParam("highlight_post_tags", *s.highlightPostTags)
	}
	if s.fullDescription != nil {
		r.addParam("full_description", *s.fullDescription)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.upgradeDetails != nil {
		r.addParam("upgrade_details", *s.upgradeDetails)
	}
	if s.userStatus != nil {
		r.addParam("user_status", *s.userStatus)
	}
	if s.userEmployerReputation != nil {
		r.addParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.addParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.addParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.addParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.addParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.addParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.addParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.addParam("corporate_users", *s.corporateUsers)
	}
	if s.marketingMobileNumber != nil {
		r.addParam("marketing_mobile_number", *s.marketingMobileNumber)
	}
	if s.sanctionDetails != nil {
		r.addParam("sanction_details", *s.sanctionDetails)
	}
	if s.limitedAccount != nil {
		r.addParam("limited_account", *s.limitedAccount)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
	}

	return execute[*ListProjectsResponse](ctx, s.client, r)

}

// SetUserResponsiveness sets whether to return the responsiveness score(s) of the selected user/users.
func (s *searchActiveProjects) SetUserResponsiveness(val bool) *searchActiveProjects {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers sets whether to return the corporate accounts that the selected user/users has created/founded.
func (s *searchActiveProjects) SetCorporateUsers(val bool) *searchActiveProjects {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber sets whether to return the mobile number of the user used by recruiters to contact them.
func (s *searchActiveProjects) SetMarketingMobileNumber(val bool) *searchActiveProjects {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to return the end timestamp of any sanctions given to the user.
func (s *searchActiveProjects) SetSanctionDetails(val bool) *searchActiveProjects {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to return the limit account status of the user.
func (s *searchActiveProjects) SetLimitedAccount(val bool) *searchActiveProjects {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to return equipment groups and items attached to the user.
func (s *searchActiveProjects) SetEquipmentGroupDetails(val bool) *searchActiveProjects {
	s.equipmentGroupDetails = &val
	return s
}

// SetUserDetails sets whether to return basic user information.
func (s *searchActiveProjects) SetUserDetails(val bool) *searchActiveProjects {
	s.userDetails = &val
	return s
}

// SetLocationDetails sets whether to return information about a project's location.
func (s *searchActiveProjects) SetLocationDetails(val bool) *searchActiveProjects {
	s.locationDetails = &val
	return s
}

// SetNdaSignatureDetails sets whether to return a list of users who have signed an NDA for the project.
func (s *searchActiveProjects) SetNdaSignatureDetails(val bool) *searchActiveProjects {
	s.ndaSignatureDetails = &val
	return s
}

// SetProjectCollaborationDetails sets whether to return a list of the collaborators of a project.
func (s *searchActiveProjects) SetProjectCollaborationDetails(val bool) *searchActiveProjects {
	s.projectCollaborationDetails = &val
	return s
}

// SetUserAvatar sets whether to return the avatar of the selected user/users.
func (s *searchActiveProjects) SetUserAvatar(val bool) *searchActiveProjects {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether to return the country flag/code of the selected user/users.
func (s *searchActiveProjects) SetUserCountryDetails(val bool) *searchActiveProjects {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether to return the profile blurb of the selected user/users.
func (s *searchActiveProjects) SetUserProfileDescription(val bool) *searchActiveProjects {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo sets whether to return the display name of the selected user/users.
func (s *searchActiveProjects) SetUserDisplayInfo(val bool) *searchActiveProjects {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether to return the jobs of the selected user/users.
func (s *searchActiveProjects) SetUserJobs(val bool) *searchActiveProjects {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether to return the currency balance of the selected user/users.
func (s *searchActiveProjects) SetUserBalanceDetails(val bool) *searchActiveProjects {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether to return qualification exams completed by the selected user/users.
func (s *searchActiveProjects) SetUserQualificationDetails(val bool) *searchActiveProjects {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether to return the membership information of the selected user/users.
func (s *searchActiveProjects) SetUserMembershipDetails(val bool) *searchActiveProjects {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether to return the financial information of the selected user/users.
func (s *searchActiveProjects) SetUserFinancialDetails(val bool) *searchActiveProjects {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether to return the location information of the selected user/users.
func (s *searchActiveProjects) SetUserLocationDetails(val bool) *searchActiveProjects {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether to return the portfolio information of the selected user/users.
func (s *searchActiveProjects) SetUserPortfolioDetails(val bool) *searchActiveProjects {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether to return the preferred information of the selected user/users.
func (s *searchActiveProjects) SetUserPreferredDetails(val bool) *searchActiveProjects {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether to return the badges earned by the selected user/users.
func (s *searchActiveProjects) SetUserBadgeDetails(val bool) *searchActiveProjects {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether to return additional status information about the selected user/users.
func (s *searchActiveProjects) SetUserStatus(val bool) *searchActiveProjects {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether to return the freelancer reputation of the selected user/users.
func (s *searchActiveProjects) SetUserReputation(val bool) *searchActiveProjects {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether to return the employer reputation of the selected user/users.
func (s *searchActiveProjects) SetUserEmployerReputation(val bool) *searchActiveProjects {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether to return the full freelancer reputation of the selected user/users.
func (s *searchActiveProjects) SetUserReputationExtra(val bool) *searchActiveProjects {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether to return the full employer reputation of the selected user/users.
func (s *searchActiveProjects) SetUserEmployerReputationExtra(val bool) *searchActiveProjects {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether to return the profile picture of the selected user/users.
func (s *searchActiveProjects) SetUserCoverImage(val bool) *searchActiveProjects {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether to return previous profile pictures of the selected user/users.
func (s *searchActiveProjects) SetUserPastCoverImage(val bool) *searchActiveProjects {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether to return recommendations count of the selected user/users.
func (s *searchActiveProjects) SetUserRecommendations(val bool) *searchActiveProjects {
	s.userRecommendations = &val
	return s
}

// SetProjectTypes sets the project types to filter results by (e.g., fixed, hourly).
func (s *searchActiveProjects) SetProjectTypes(projectTypes []ProjectType) *searchActiveProjects {
	s.projectTypes = projectTypes
	return s
}

// SetProjectUpgrades sets the project upgrades to filter results by (e.g., featured, sealed).
func (s *searchActiveProjects) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *searchActiveProjects {
	s.projectUpgrades = projectUpgrades
	return s
}

// SetContestUpgrades sets the contest upgrades to filter results by (e.g., guaranteed, highlight).
func (s *searchActiveProjects) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *searchActiveProjects {
	s.contestUpgrades = contestUpgrades
	return s
}

// SetCompact sets whether to strip all null and empty values from the response.
func (s *searchActiveProjects) SetCompact(val bool) *searchActiveProjects {
	s.compact = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *searchActiveProjects) SetLimit(val int) *searchActiveProjects {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *searchActiveProjects) SetOffset(val int) *searchActiveProjects {
	s.offset = &val
	return s
}

// SetReverseSort sets whether results should appear in ascending order instead of descending.
func (s *searchActiveProjects) SetReverseSort(val bool) *searchActiveProjects {
	s.reverseSort = &val
	return s
}

// SetFullDescription sets whether to return the full project description.
func (s *searchActiveProjects) SetFullDescription(val bool) *searchActiveProjects {
	s.fullDescription = &val
	return s
}

// SetJobDetails sets whether to return job information.
func (s *searchActiveProjects) SetJobDetails(val bool) *searchActiveProjects {
	s.jobDetails = &val
	return s
}

// SetUpgradeDetails sets whether to return upgrade information.
func (s *searchActiveProjects) SetUpgradeDetails(val bool) *searchActiveProjects {
	s.upgradeDetails = &val
	return s
}

// SetQuery sets the search query string to filter projects.
func (s *searchActiveProjects) SetQuery(val string) *searchActiveProjects {
	s.query = &val
	return s
}

// SetMinAvgPrice sets the minimum average bid in USD to filter projects.
func (s *searchActiveProjects) SetMinAvgPrice(val float64) *searchActiveProjects {
	s.minAvgPrice = &val
	return s
}

// SetMaxAvgPrice sets the maximum average bid in USD to filter projects.
func (s *searchActiveProjects) SetMaxAvgPrice(val float64) *searchActiveProjects {
	s.maxAvgPrice = &val
	return s
}

// SetMinAvgHourlyRate sets the minimum hourly bid rate in USD to filter projects.
func (s *searchActiveProjects) SetMinAvgHourlyRate(val float64) *searchActiveProjects {
	s.minAvgHourlyRate = &val
	return s
}

// SetMaxAvgHourlyRate sets the maximum hourly bid rate in USD to filter projects.
func (s *searchActiveProjects) SetMaxAvgHourlyRate(val float64) *searchActiveProjects {
	s.maxAvgHourlyRate = &val
	return s
}

// SetMinPrice sets the minimum fixed price budget in USD to filter projects.
func (s *searchActiveProjects) SetMinPrice(val float64) *searchActiveProjects {
	s.minPrice = &val
	return s
}

// SetMaxPrice sets the maximum fixed price budget in USD to filter projects.
func (s *searchActiveProjects) SetMaxPrice(val float64) *searchActiveProjects {
	s.maxPrice = &val
	return s
}

// SetMinHourlyRate sets the minimum hourly rate budget in USD to filter projects.
func (s *searchActiveProjects) SetMinHourlyRate(val float64) *searchActiveProjects {
	s.minHourlyRate = &val
	return s
}

// SetMaxHourlyRate sets the maximum hourly rate budget in USD to filter projects.
func (s *searchActiveProjects) SetMaxHourlyRate(val float64) *searchActiveProjects {
	s.maxHourlyRate = &val
	return s
}

// SetJobs sets projects to filter by at least one of the specified job IDs.
func (s *searchActiveProjects) SetJobs(jobs []int32) *searchActiveProjects {
	s.jobs = jobs
	return s
}

// SetCountries sets projects to filter by at least one of the specified country codes.
func (s *searchActiveProjects) SetCountries(countries []string) *searchActiveProjects {
	s.countries = countries
	return s
}

// SetLanguages sets projects to filter by at least one of the specified language IDs.
func (s *searchActiveProjects) SetLanguages(languages []string) *searchActiveProjects {
	s.languages = languages
	return s
}

// SetLatitude sets the latitude to filter projects near the specified location.
func (s *searchActiveProjects) SetLatitude(val float64) *searchActiveProjects {
	s.latitude = &val
	return s
}

// SetLongitude sets the longitude to filter projects near the specified location.
func (s *searchActiveProjects) SetLongitude(val float64) *searchActiveProjects {
	s.longitude = &val
	return s
}

// SetFromTime sets the timestamp to filter projects last updated after this time (inclusive).
func (s *searchActiveProjects) SetFromTime(val int64) *searchActiveProjects {
	s.fromTime = &val
	return s
}

// SetToTime sets the timestamp to filter projects last updated before this time (inclusive).
func (s *searchActiveProjects) SetToTime(val int64) *searchActiveProjects {
	s.toTime = &val
	return s
}

// SetSortField sets the field to sort projects by (e.g., time_updated, bid_count, bid_avg_usd).
func (s *searchActiveProjects) SetSortField(val SortField) *searchActiveProjects {
	s.sortField = &val
	return s
}

// SetProjectIDs sets projects to filter by the specified project IDs.
func (s *searchActiveProjects) SetProjectIDs(projectIDs []int64) *searchActiveProjects {
	s.projectIDs = projectIDs
	return s
}

// SetTopRightLatitude sets the top-right latitude for map boundary filtering.
func (s *searchActiveProjects) SetTopRightLatitude(val float64) *searchActiveProjects {
	s.topRightLatitude = &val
	return s
}

// SetTopRightLongitude sets the top-right longitude for map boundary filtering.
func (s *searchActiveProjects) SetTopRightLongitude(val float64) *searchActiveProjects {
	s.topRightLongitude = &val
	return s
}

// SetBottomLeftLatitude sets the bottom-left latitude for map boundary filtering.
func (s *searchActiveProjects) SetBottomLeftLatitude(val float64) *searchActiveProjects {
	s.bottomLeftLatitude = &val
	return s
}

// SetBottomLeftLongitude sets the bottom-left longitude for map boundary filtering.
func (s *searchActiveProjects) SetBottomLeftLongitude(val float64) *searchActiveProjects {
	s.bottomLeftLongitude = &val
	return s
}

// SetOrSearchQuery sets an alternative query to return results that match any term in the query.
func (s *searchActiveProjects) SetOrSearchQuery(val string) *searchActiveProjects {
	s.orSearchQuery = &val
	return s
}

// SetHighlightPreTags sets the string to add before any matching term in the project description.
func (s *searchActiveProjects) SetHighlightPreTags(val string) *searchActiveProjects {
	s.highlightPreTags = &val
	return s
}

// SetHighlightPostTags sets the string to add after any matching term in the project description.
func (s *searchActiveProjects) SetHighlightPostTags(val string) *searchActiveProjects {
	s.highlightPostTags = &val
	return s
}

type searchAllProjects struct {
	client                      *Client
	query                       *string
	projectTypes                []ProjectType
	projectUpgrades             []ProjectUpgradeType
	contestUpgrades             []ContestUpgradeType
	minAvgPrice                 *float64
	maxAvgPrice                 *float64
	minAvgHourlyRate            *float64
	maxAvgHourlyRate            *float64
	minPrice                    *float64
	maxPrice                    *float64
	minHourlyRate               *float64
	maxHourlyRate               *float64
	jobs                        []int32
	countries                   []string
	languages                   []string
	latitude                    *float64
	longitude                   *float64
	fromTime                    *int64
	toTime                      *int64
	sortField                   *SortField
	bidAwardStatuses            []BidAwardStatus
	bidCompleteStatuses         []BidCompleteStatus
	projectStatuses             []ProjectFrontendStatus
	projectIDs                  []int64
	topRightLatitude            *float64
	topRightLongitude           *float64
	bottomLeftLatitude          *float64
	bottomLeftLongitude         *float64
	reverseSort                 *bool
	orSearchQuery               *string
	highlightPreTags            *string
	highlightPostTags           *string
	fullDescription             *bool
	jobDetails                  *bool
	upgradeDetails              *bool
	userDetails                 *bool
	locationDetails             *bool
	ndaSignatureDetails         *bool
	projectCollaborationDetails *bool
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
	compact                     *bool
}

func (s *searchAllProjects) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS_ACTIVE),
	}
	if s.query != nil {
		r.addParam("query", s.query)
	}
	for _, projectType := range s.projectTypes {
		r.addParam("project_types[]", projectType)
	}
	for _, projectUpgrade := range s.projectUpgrades {
		r.addParam("project_upgrades[]", projectUpgrade)
	}
	for _, contestUpgrade := range s.contestUpgrades {
		r.addParam("contest_upgrades[]", contestUpgrade)
	}
	if s.contestUpgrades != nil {
		r.addParam("contest_upgrades", s.contestUpgrades)
	}
	if s.minAvgPrice != nil {
		r.addParam("min_avg_price", *s.minAvgPrice)
	}
	if s.maxAvgPrice != nil {
		r.addParam("max_avg_price", *s.maxAvgPrice)
	}
	if s.minAvgHourlyRate != nil {
		r.addParam("min_avg_hourly_rate", *s.minAvgHourlyRate)
	}
	if s.maxAvgHourlyRate != nil {
		r.addParam("max_avg_hourly_rate", *s.maxAvgHourlyRate)
	}
	if s.minPrice != nil {
		r.addParam("min_price", *s.minPrice)
	}
	if s.maxPrice != nil {
		r.addParam("max_price", *s.maxPrice)
	}
	if s.minHourlyRate != nil {
		r.addParam("min_hourly_rate", *s.minHourlyRate)
	}
	if s.maxHourlyRate != nil {
		r.addParam("max_hourly_rate", *s.maxHourlyRate)
	}
	for _, job := range s.jobs {
		r.addParam("jobs[]", job)
	}
	for _, country := range s.countries {
		r.addParam("countries[]", country)
	}
	for _, language := range s.languages {
		r.addParam("languages[]", language)
	}
	if s.latitude != nil {
		r.addParam("latitude", *s.latitude)
	}
	if s.longitude != nil {
		r.addParam("longitude", *s.longitude)
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
	for _, bidAwardStatus := range s.bidAwardStatuses {
		r.addParam("bid_award_statuses[]", bidAwardStatus)
	}
	for _, bidCompleteStatus := range s.bidCompleteStatuses {
		r.addParam("bid_complete_statuses[]", bidCompleteStatus)
	}
	for _, projectStatus := range s.projectStatuses {
		r.addParam("project_statuses[]", projectStatus)
	}
	for _, projectID := range s.projectIDs {
		r.addParam("project_ids[]", projectID)
	}
	if s.topRightLatitude != nil {
		r.addParam("top_right_latitude", *s.topRightLatitude)
	}
	if s.topRightLongitude != nil {
		r.addParam("top_right_longitude", *s.topRightLongitude)
	}
	if s.bottomLeftLatitude != nil {
		r.addParam("bottom_left_latitude", *s.bottomLeftLatitude)
	}
	if s.bottomLeftLongitude != nil {
		r.addParam("bottom_left_longitude", *s.bottomLeftLongitude)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.orSearchQuery != nil {
		r.addParam("or_search_query", *s.orSearchQuery)
	}
	if s.highlightPreTags != nil {
		r.addParam("highlight_pre_tags", *s.highlightPreTags)
	}
	if s.highlightPostTags != nil {
		r.addParam("highlight_post_tags", *s.highlightPostTags)
	}
	if s.fullDescription != nil {
		r.addParam("full_description", *s.fullDescription)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.upgradeDetails != nil {
		r.addParam("upgrade_details", *s.upgradeDetails)
	}
	if s.userStatus != nil {
		r.addParam("user_status", *s.userStatus)
	}
	if s.userEmployerReputation != nil {
		r.addParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.addParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.addParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.addParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.addParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.addParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.addParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.addParam("corporate_users", *s.corporateUsers)
	}
	if s.marketingMobileNumber != nil {
		r.addParam("marketing_mobile_number", *s.marketingMobileNumber)
	}
	if s.sanctionDetails != nil {
		r.addParam("sanction_details", *s.sanctionDetails)
	}
	if s.limitedAccount != nil {
		r.addParam("limited_account", *s.limitedAccount)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
	}

	return execute[*ListProjectsResponse](ctx, s.client, r)

}

// SetUserResponsiveness sets whether to return the responsiveness score(s) of the selected user/users.
func (s *searchAllProjects) SetUserResponsiveness(val bool) *searchAllProjects {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers sets whether to return the corporate accounts that the selected user/users has created/founded.
func (s *searchAllProjects) SetCorporateUsers(val bool) *searchAllProjects {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber sets whether to return the mobile number of the user used by recruiters to contact them.
func (s *searchAllProjects) SetMarketingMobileNumber(val bool) *searchAllProjects {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to return the end timestamp of any sanctions given to the user.
func (s *searchAllProjects) SetSanctionDetails(val bool) *searchAllProjects {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to return the limit account status of the user.
func (s *searchAllProjects) SetLimitedAccount(val bool) *searchAllProjects {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to return equipment groups and items attached to the user.
func (s *searchAllProjects) SetEquipmentGroupDetails(val bool) *searchAllProjects {
	s.equipmentGroupDetails = &val
	return s
}

// SetUserDetails sets whether to return basic user information.
func (s *searchAllProjects) SetUserDetails(val bool) *searchAllProjects {
	s.userDetails = &val
	return s
}

// SetLocationDetails sets whether to return information about a project's location.
func (s *searchAllProjects) SetLocationDetails(val bool) *searchAllProjects {
	s.locationDetails = &val
	return s
}

// SetNdaSignatureDetails sets whether to return a list of users who have signed an NDA for the project.
func (s *searchAllProjects) SetNdaSignatureDetails(val bool) *searchAllProjects {
	s.ndaSignatureDetails = &val
	return s
}

// SetProjectCollaborationDetails sets whether to return a list of the collaborators of a project.
func (s *searchAllProjects) SetProjectCollaborationDetails(val bool) *searchAllProjects {
	s.projectCollaborationDetails = &val
	return s
}

// SetUserAvatar sets whether to return the avatar of the selected user/users.
func (s *searchAllProjects) SetUserAvatar(val bool) *searchAllProjects {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether to return the country flag/code of the selected user/users.
func (s *searchAllProjects) SetUserCountryDetails(val bool) *searchAllProjects {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether to return the profile blurb of the selected user/users.
func (s *searchAllProjects) SetUserProfileDescription(val bool) *searchAllProjects {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo sets whether to return the display name of the selected user/users.
func (s *searchAllProjects) SetUserDisplayInfo(val bool) *searchAllProjects {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether to return the jobs of the selected user/users.
func (s *searchAllProjects) SetUserJobs(val bool) *searchAllProjects {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether to return the currency balance of the selected user/users.
func (s *searchAllProjects) SetUserBalanceDetails(val bool) *searchAllProjects {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether to return qualification exams completed by the selected user/users.
func (s *searchAllProjects) SetUserQualificationDetails(val bool) *searchAllProjects {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether to return the membership information of the selected user/users.
func (s *searchAllProjects) SetUserMembershipDetails(val bool) *searchAllProjects {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether to return the financial information of the selected user/users.
func (s *searchAllProjects) SetUserFinancialDetails(val bool) *searchAllProjects {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether to return the location information of the selected user/users.
func (s *searchAllProjects) SetUserLocationDetails(val bool) *searchAllProjects {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether to return the portfolio information of the selected user/users.
func (s *searchAllProjects) SetUserPortfolioDetails(val bool) *searchAllProjects {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether to return the preferred information of the selected user/users.
func (s *searchAllProjects) SetUserPreferredDetails(val bool) *searchAllProjects {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether to return the badges earned by the selected user/users.
func (s *searchAllProjects) SetUserBadgeDetails(val bool) *searchAllProjects {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether to return additional status information about the selected user/users.
func (s *searchAllProjects) SetUserStatus(val bool) *searchAllProjects {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether to return the freelancer reputation of the selected user/users.
func (s *searchAllProjects) SetUserReputation(val bool) *searchAllProjects {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether to return the employer reputation of the selected user/users.
func (s *searchAllProjects) SetUserEmployerReputation(val bool) *searchAllProjects {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether to return the full freelancer reputation of the selected user/users.
func (s *searchAllProjects) SetUserReputationExtra(val bool) *searchAllProjects {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether to return the full employer reputation of the selected user/users.
func (s *searchAllProjects) SetUserEmployerReputationExtra(val bool) *searchAllProjects {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether to return the profile picture of the selected user/users.
func (s *searchAllProjects) SetUserCoverImage(val bool) *searchAllProjects {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether to return previous profile pictures of the selected user/users.
func (s *searchAllProjects) SetUserPastCoverImage(val bool) *searchAllProjects {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether to return recommendations count of the selected user/users.
func (s *searchAllProjects) SetUserRecommendations(val bool) *searchAllProjects {
	s.userRecommendations = &val
	return s
}

// SetProjectTypes sets the project types to filter results by (e.g., fixed, hourly).
func (s *searchAllProjects) SetProjectTypes(projectTypes []ProjectType) *searchAllProjects {
	s.projectTypes = projectTypes
	return s
}

// SetProjectUpgrades sets the project upgrades to filter results by (e.g., featured, sealed).
func (s *searchAllProjects) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *searchAllProjects {
	s.projectUpgrades = projectUpgrades
	return s
}

// SetContestUpgrades sets the contest upgrades to filter results by (e.g., guaranteed, highlight).
func (s *searchAllProjects) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *searchAllProjects {
	s.contestUpgrades = contestUpgrades
	return s
}

// SetCompact sets whether to strip all null and empty values from the response.
func (s *searchAllProjects) SetCompact(val bool) *searchAllProjects {
	s.compact = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *searchAllProjects) SetLimit(val int) *searchAllProjects {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *searchAllProjects) SetOffset(val int) *searchAllProjects {
	s.offset = &val
	return s
}

// SetReverseSort sets whether results should appear in ascending order instead of descending.
func (s *searchAllProjects) SetReverseSort(val bool) *searchAllProjects {
	s.reverseSort = &val
	return s
}

// SetFullDescription sets whether to return the full project description.
func (s *searchAllProjects) SetFullDescription(val bool) *searchAllProjects {
	s.fullDescription = &val
	return s
}

// SetJobDetails sets whether to return job information.
func (s *searchAllProjects) SetJobDetails(val bool) *searchAllProjects {
	s.jobDetails = &val
	return s
}

// SetUpgradeDetails sets whether to return upgrade information.
func (s *searchAllProjects) SetUpgradeDetails(val bool) *searchAllProjects {
	s.upgradeDetails = &val
	return s
}

// SetQuery sets the search query string to filter projects.
func (s *searchAllProjects) SetQuery(val string) *searchAllProjects {
	s.query = &val
	return s
}

// SetMinAvgPrice sets the minimum average bid in USD to filter projects.
func (s *searchAllProjects) SetMinAvgPrice(val float64) *searchAllProjects {
	s.minAvgPrice = &val
	return s
}

// SetMaxAvgPrice sets the maximum average bid in USD to filter projects.
func (s *searchAllProjects) SetMaxAvgPrice(val float64) *searchAllProjects {
	s.maxAvgPrice = &val
	return s
}

// SetMinAvgHourlyRate sets the minimum hourly bid rate in USD to filter projects.
func (s *searchAllProjects) SetMinAvgHourlyRate(val float64) *searchAllProjects {
	s.minAvgHourlyRate = &val
	return s
}

// SetMaxAvgHourlyRate sets the maximum hourly bid rate in USD to filter projects.
func (s *searchAllProjects) SetMaxAvgHourlyRate(val float64) *searchAllProjects {
	s.maxAvgHourlyRate = &val
	return s
}

// SetMinPrice sets the minimum fixed price budget in USD to filter projects.
func (s *searchAllProjects) SetMinPrice(val float64) *searchAllProjects {
	s.minPrice = &val
	return s
}

// SetMaxPrice sets the maximum fixed price budget in USD to filter projects.
func (s *searchAllProjects) SetMaxPrice(val float64) *searchAllProjects {
	s.maxPrice = &val
	return s
}

// SetMinHourlyRate sets the minimum hourly rate budget in USD to filter projects.
func (s *searchAllProjects) SetMinHourlyRate(val float64) *searchAllProjects {
	s.minHourlyRate = &val
	return s
}

// SetMaxHourlyRate sets the maximum hourly rate budget in USD to filter projects.
func (s *searchAllProjects) SetMaxHourlyRate(val float64) *searchAllProjects {
	s.maxHourlyRate = &val
	return s
}

// SetJobs sets projects to filter by at least one of the specified job IDs.
func (s *searchAllProjects) SetJobs(jobs []int32) *searchAllProjects {
	s.jobs = jobs
	return s
}

// SetCountries sets projects to filter by at least one of the specified country codes.
func (s *searchAllProjects) SetCountries(countries []string) *searchAllProjects {
	s.countries = countries
	return s
}

// SetLanguages sets projects to filter by at least one of the specified language IDs.
func (s *searchAllProjects) SetLanguages(languages []string) *searchAllProjects {
	s.languages = languages
	return s
}

// SetLatitude sets the latitude to filter projects near the specified location.
func (s *searchAllProjects) SetLatitude(val float64) *searchAllProjects {
	s.latitude = &val
	return s
}

// SetLongitude sets the longitude to filter projects near the specified location.
func (s *searchAllProjects) SetLongitude(val float64) *searchAllProjects {
	s.longitude = &val
	return s
}

// SetFromTime sets the timestamp to filter projects last updated after this time (inclusive).
func (s *searchAllProjects) SetFromTime(val int64) *searchAllProjects {
	s.fromTime = &val
	return s
}

// SetToTime sets the timestamp to filter projects last updated before this time (inclusive).
func (s *searchAllProjects) SetToTime(val int64) *searchAllProjects {
	s.toTime = &val
	return s
}

// SetSortField sets the field to sort projects by (e.g., time_updated, bid_count, bid_avg_usd).
func (s *searchAllProjects) SetSortField(val SortField) *searchAllProjects {
	s.sortField = &val
	return s
}

// SetBidAwardsStatuses sets the bid award statuses filter for the project search.
func (s *searchAllProjects) SetBidAwardsStatuses(vals []BidAwardStatus) *searchAllProjects {
	s.bidAwardStatuses = vals
	return s
}

// SetBidCompleteStatuses sets the bid complete statuses filter for the project search.
func (s *searchAllProjects) SetBidCompleteStatuses(vals []BidCompleteStatus) *searchAllProjects {
	s.bidCompleteStatuses = vals
	return s
}

// SetProjectStatuses sets the project statuses filter for the project search.
func (s *searchAllProjects) SetProjectStatuses(vals []ProjectFrontendStatus) *searchAllProjects {
	s.projectStatuses = vals
	return s
}

// SetProjectIDs sets projects to filter by the specified project IDs.
func (s *searchAllProjects) SetProjectIDs(projectIDs []int64) *searchAllProjects {
	s.projectIDs = projectIDs
	return s
}

// SetTopRightLatitude sets the top-right latitude for map boundary filtering.
func (s *searchAllProjects) SetTopRightLatitude(val float64) *searchAllProjects {
	s.topRightLatitude = &val
	return s
}

// SetTopRightLongitude sets the top-right longitude for map boundary filtering.
func (s *searchAllProjects) SetTopRightLongitude(val float64) *searchAllProjects {
	s.topRightLongitude = &val
	return s
}

// SetBottomLeftLatitude sets the bottom-left latitude for map boundary filtering.
func (s *searchAllProjects) SetBottomLeftLatitude(val float64) *searchAllProjects {
	s.bottomLeftLatitude = &val
	return s
}

// SetBottomLeftLongitude sets the bottom-left longitude for map boundary filtering.
func (s *searchAllProjects) SetBottomLeftLongitude(val float64) *searchAllProjects {
	s.bottomLeftLongitude = &val
	return s
}

// SetOrSearchQuery sets an alternative query to return results that match any term in the query.
func (s *searchAllProjects) SetOrSearchQuery(val string) *searchAllProjects {
	s.orSearchQuery = &val
	return s
}

// SetHighlightPreTags sets the string to add before any matching term in the project description.
func (s *searchAllProjects) SetHighlightPreTags(val string) *searchAllProjects {
	s.highlightPreTags = &val
	return s
}

// SetHighlightPostTags sets the string to add after any matching term in the project description.
func (s *searchAllProjects) SetHighlightPostTags(val string) *searchAllProjects {
	s.highlightPostTags = &val
	return s
}

type inviteFreelancer struct {
	client    *Client
	projectID int64
}

type InviteFreelancersRequestBody struct {
	FreelancerID int64 `json:"freelancer_id"`
}

func (s *inviteFreelancer) Do(ctx context.Context, b InviteFreelancersRequestBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("%s/%s/invite", PROJECTS_PROJECTS, strconv.FormatInt(s.projectID, 10)),
		body:     bytes.NewBuffer(m),
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: Refine the typed response with
type listUpgradeFees struct {
	client             *Client
	currencies         []int
	project            *int64
	freeUpgradeDetails *bool
	taxIncluded        *bool
}

func (s *listUpgradeFees) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS_FEES),
	}

	for _, currency := range s.currencies {
		r.addParam("currencies[]", currency)
	}
	if s.project != nil {
		r.addParam("project", *s.project)
	}
	if s.freeUpgradeDetails != nil {
		r.addParam("fee_upgrade_details", *s.freeUpgradeDetails)
	}
	if s.taxIncluded != nil {
		r.addParam("tax_included", *s.taxIncluded)
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// SetCurrencies sets the list of currency IDs to filter the upgrade fees for.
// This corresponds to the `currencies[]` parameter in the API request.
func (s *listUpgradeFees) SetCurrencies(vals []int) *listUpgradeFees {
	s.currencies = vals
	return s
}

// SetProject sets the project ID to check if the project is eligible
// for free upgrades. This corresponds to the optional `project` parameter
// in the API request.
func (s *listUpgradeFees) SetProject(val int64) *listUpgradeFees {
	s.project = &val
	return s
}

// SetFreeUpgradeDetails sets whether to return details about the current user's
// eligibility for free project upgrades. This corresponds to the optional
// `free_upgrade_details` parameter in the API request.
func (s *listUpgradeFees) SetFreeUpgradeDetails(val bool) *listUpgradeFees {
	s.freeUpgradeDetails = &val
	return s
}

// SetTaxIncluded sets whether the returned upgrade fees should include tax.
// This corresponds to the optional `tax_included` parameter in the API request.
func (s *listUpgradeFees) SetTaxIncluded(val bool) *listUpgradeFees {
	s.taxIncluded = &val
	return s
}

// TODO: Refine the typed response with
type listBidUpgradeFees struct {
	client      *Client
	currencies  []int
	taxIncluded *bool
}

func (s *listBidUpgradeFees) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_BIDS_FEES),
	}

	for _, currency := range s.currencies {
		r.addParam("currencies[]", currency)
	}
	if s.taxIncluded != nil {
		r.addParam("tax_included", *s.taxIncluded)
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// SetCurrencies sets the list of currency IDs for which bid upgrade fees
// should be returned.
func (s *listBidUpgradeFees) SetCurrencies(vals []int) *listBidUpgradeFees {
	s.currencies = vals
	return s
}

// SetTaxIncluded specifies whether bid upgrade fees should be returned with
// tax included in the price.
func (s *listBidUpgradeFees) SetTaxIncluded(val bool) *listBidUpgradeFees {
	s.taxIncluded = &val
	return s
}

// TODO: Refine the typed response with
type listProjectBids struct {
	client                      *Client
	projectID                   int64
	isShortlisted               *bool
	reputation                  *bool
	recommendedBid              *bool
	shortlistedBid              *bool
	distance                    *bool
	userDetails                 *bool
	expertGuarantees            *bool
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
	compact                     *bool
	quotations                  *bool
}

func (s *listProjectBids) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/bids", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}

	if s.isShortlisted != nil {
		r.addParam("is_shortlisted", *s.isShortlisted)
	}
	if s.reputation != nil {
		r.addParam("reputation", *s.reputation)
	}
	if s.recommendedBid != nil {
		r.addParam("recommended_bid", *s.recommendedBid)
	}
	if s.shortlistedBid != nil {
		r.addParam("shortlisted_bid", *s.shortlistedBid)
	}
	if s.distance != nil {
		r.addParam("distance", *s.distance)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.expertGuarantees != nil {
		r.addParam("expert_guarantees", *s.expertGuarantees)
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
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.setParam("compact", *s.compact)
	}
	if s.quotations != nil {
		r.setParam("quotations", *s.quotations)
	}

	return execute[*RawResponse](ctx, s.client, r)
}

// SetIsShortlisted sets whether to restrict results to shortlisted bids (project owner only).
func (s *listProjectBids) SetIsShortlisted(val bool) *listProjectBids {
	s.isShortlisted = &val
	return s
}

// SetReputation sets whether to return reputations relevant to bids.
func (s *listProjectBids) SetReputation(val bool) *listProjectBids {
	s.reputation = &val
	return s
}

// SetRecommendedBid sets whether to return the recommended bid (project owner only).
func (s *listProjectBids) SetRecommendedBid(val bool) *listProjectBids {
	s.recommendedBid = &val
	return s
}

// SetShortlistedBid sets whether to return the list of shortlisted bids (project owner only).
func (s *listProjectBids) SetShortlistedBid(val bool) *listProjectBids {
	s.shortlistedBid = &val
	return s
}

// SetDistance sets whether to return distance information about the project bids.
func (s *listProjectBids) SetDistance(val bool) *listProjectBids {
	s.distance = &val
	return s
}

// SetUserDetails sets whether to return basic user information.
func (s *listProjectBids) SetUserDetails(val bool) *listProjectBids {
	s.userDetails = &val
	return s
}

// SetExpertGuarantees sets whether to return additional expert guarantees indexed by bid id.
func (s *listProjectBids) SetExpertGuarantees(val bool) *listProjectBids {
	s.expertGuarantees = &val
	return s
}

// SetUserAvatar sets whether to return the avatar of the selected user/users.
func (s *listProjectBids) SetUserAvatar(val bool) *listProjectBids {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails sets whether to return the country flag/code of selected user/users.
func (s *listProjectBids) SetUserCountryDetails(val bool) *listProjectBids {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription sets whether to return the profile blurb of selected user/users.
func (s *listProjectBids) SetUserProfileDescription(val bool) *listProjectBids {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo sets whether to return the display name of selected user/users.
func (s *listProjectBids) SetUserDisplayInfo(val bool) *listProjectBids {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs sets whether to return the jobs of selected user/users.
func (s *listProjectBids) SetUserJobs(val bool) *listProjectBids {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails sets whether to return the currency balance of selected user/users.
func (s *listProjectBids) SetUserBalanceDetails(val bool) *listProjectBids {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails sets whether to return the qualification exams completed by selected user/users.
func (s *listProjectBids) SetUserQualificationDetails(val bool) *listProjectBids {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails sets whether to return the membership information of selected user/users.
func (s *listProjectBids) SetUserMembershipDetails(val bool) *listProjectBids {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails sets whether to return the financial information of selected user/users.
func (s *listProjectBids) SetUserFinancialDetails(val bool) *listProjectBids {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails sets whether to return the location information of selected user/users.
func (s *listProjectBids) SetUserLocationDetails(val bool) *listProjectBids {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails sets whether to return the portfolio information of selected user/users.
func (s *listProjectBids) SetUserPortfolioDetails(val bool) *listProjectBids {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails sets whether to return the preferred information of selected user/users.
func (s *listProjectBids) SetUserPreferredDetails(val bool) *listProjectBids {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails sets whether to return the badges earned by selected user/users.
func (s *listProjectBids) SetUserBadgeDetails(val bool) *listProjectBids {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus sets whether to return additional status information about selected user/users.
func (s *listProjectBids) SetUserStatus(val bool) *listProjectBids {
	s.userStatus = &val
	return s
}

// SetUserReputation sets whether to return the freelancer reputation of selected user/users.
func (s *listProjectBids) SetUserReputation(val bool) *listProjectBids {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation sets whether to return the employer reputation of selected user/users.
func (s *listProjectBids) SetUserEmployerReputation(val bool) *listProjectBids {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra sets whether to return the full freelancer reputation of selected user/users.
func (s *listProjectBids) SetUserReputationExtra(val bool) *listProjectBids {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra sets whether to return the full employer reputation of selected user/users.
func (s *listProjectBids) SetUserEmployerReputationExtra(val bool) *listProjectBids {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage sets whether to return the profile picture of selected user/users.
func (s *listProjectBids) SetUserCoverImage(val bool) *listProjectBids {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage sets whether to return previous profile pictures of selected user/users.
func (s *listProjectBids) SetUserPastCoverImage(val bool) *listProjectBids {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations sets whether to return the recommendations count of selected user/users.
func (s *listProjectBids) SetUserRecommendations(val bool) *listProjectBids {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness sets whether to return the responsiveness score(s) of selected user/users.
func (s *listProjectBids) SetUserResponsiveness(val bool) *listProjectBids {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers sets whether to return the corporate accounts created/founded by selected user/users.
func (s *listProjectBids) SetCorporateUsers(val bool) *listProjectBids {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber sets whether to return the mobile number used by the recruiter to contact the user.
func (s *listProjectBids) SetMarketingMobileNumber(val bool) *listProjectBids {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to return the end timestamp of the sanction given to the user.
func (s *listProjectBids) SetSanctionDetails(val bool) *listProjectBids {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to return the limit account status of the user.
func (s *listProjectBids) SetLimitedAccount(val bool) *listProjectBids {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to return the equipment groups and items attached to the user.
func (s *listProjectBids) SetEquipmentGroupDetails(val bool) *listProjectBids {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *listProjectBids) SetLimit(val int) *listProjectBids {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip for pagination.
func (s *listProjectBids) SetOffset(val int) *listProjectBids {
	s.offset = &val
	return s
}

// SetCompact sets whether to strip all null and empty values from the response.
func (s *listProjectBids) SetCompact(val bool) *listProjectBids {
	s.compact = &val
	return s
}

// SetQuotations sets whether to return quotation(s) attached to the bid.
func (s *listProjectBids) SetQuotations(val bool) *listProjectBids {
	s.quotations = &val
	return s
}

// TODO: refine with typed response
type getProjectBidInfo struct {
	client    *Client
	projectID int64
}

func (s *getProjectBidInfo) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/bidinfo", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

// TODO: refine with typed response
type listProjectMilestones struct {
	client                      *Client
	projectID                   int64
	statuses                    []MilestoneStatus
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

func (s *listProjectMilestones) Do(ctx context.Context, projectID int64) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/milestones", string(PROJECTS_PROJECTS), strconv.FormatInt(projectID, 10)),
	}

	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
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

// SetStatuses filters the returned project milestones by their status.
// Only milestones matching the provided status values are included in the response.
func (s *listProjectMilestones) SetStatuses(values []MilestoneStatus) *listProjectMilestones {
	s.statuses = values
	return s
}

// SetUserAvatar controls whether the avatar of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserAvatar(val bool) *listProjectMilestones {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails controls whether country flag and country code details
// of the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserCountryDetails(val bool) *listProjectMilestones {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription controls whether the profile description (blurb)
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserProfileDescription(val bool) *listProjectMilestones {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo controls whether display name information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserDisplayInfo(val bool) *listProjectMilestones {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs controls whether job information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserJobs(val bool) *listProjectMilestones {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails controls whether currency balance details
// of the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserBalanceDetails(val bool) *listProjectMilestones {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails controls whether qualification exam details
// completed by the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserQualificationDetails(val bool) *listProjectMilestones {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails controls whether membership information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserMembershipDetails(val bool) *listProjectMilestones {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails controls whether financial information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserFinancialDetails(val bool) *listProjectMilestones {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails controls whether location information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserLocationDetails(val bool) *listProjectMilestones {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails controls whether portfolio information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserPortfolioDetails(val bool) *listProjectMilestones {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails controls whether preferred information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserPreferredDetails(val bool) *listProjectMilestones {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails controls whether badge information
// earned by the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserBadgeDetails(val bool) *listProjectMilestones {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus controls whether additional status information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserStatus(val bool) *listProjectMilestones {
	s.userStatus = &val
	return s
}

// SetUserReputation controls whether freelancer reputation information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserReputation(val bool) *listProjectMilestones {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation controls whether employer reputation information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserEmployerReputation(val bool) *listProjectMilestones {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra controls whether full freelancer reputation details
// of the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserReputationExtra(val bool) *listProjectMilestones {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra controls whether full employer reputation details
// of the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserEmployerReputationExtra(val bool) *listProjectMilestones {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage controls whether the profile cover image
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserCoverImage(val bool) *listProjectMilestones {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage controls whether past profile cover images
// of the related user(s) are included in the response.
func (s *listProjectMilestones) SetUserPastCoverImage(val bool) *listProjectMilestones {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations controls whether recommendation count information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserRecommendations(val bool) *listProjectMilestones {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness controls whether responsiveness score information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetUserResponsiveness(val bool) *listProjectMilestones {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers controls whether corporate account information
// associated with the related user(s) is included in the response.
func (s *listProjectMilestones) SetCorporateUsers(val bool) *listProjectMilestones {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber controls whether the marketing mobile number
// used by recruiters to contact the related user(s) is included in the response.
func (s *listProjectMilestones) SetMarketingMobileNumber(val bool) *listProjectMilestones {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails controls whether sanction end timestamp information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetSanctionDetails(val bool) *listProjectMilestones {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount controls whether limited account status information
// of the related user(s) is included in the response.
func (s *listProjectMilestones) SetLimitedAccount(val bool) *listProjectMilestones {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails controls whether equipment group and item information
// attached to the related user(s) is included in the response.
func (s *listProjectMilestones) SetEquipmentGroupDetails(val bool) *listProjectMilestones {
	s.equipmentGroupDetails = &val
	return s
}

// TODO: refine with typed response
type listProjectMilestoneRequests struct {
	client                      *Client
	projectID                   int64
	statuses                    []MilestoneStatus
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

func (s *listProjectMilestoneRequests) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/milestone_requests", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}

	for _, val := range s.statuses {
		r.addParam("statuses[]", val)
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

// SetStatuses filters the milestone requests by their status.
// This corresponds to the `statuses[]` query parameter and allows returning
// only milestone requests matching the specified lifecycle states
// (e.g. pending, created, rejected, deleted).
func (s *listProjectMilestoneRequests) SetStatuses(values []MilestoneStatus) *listProjectMilestoneRequests {
	s.statuses = values
	return s
}

// SetUserAvatar controls whether the freelancer's avatar is included
// in the milestone requests response.
func (s *listProjectMilestoneRequests) SetUserAvatar(val bool) *listProjectMilestoneRequests {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails controls whether the freelancer's country
// details (such as flag or country code) are included in the response.
func (s *listProjectMilestoneRequests) SetUserCountryDetails(val bool) *listProjectMilestoneRequests {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription controls whether the freelancer's profile
// description (profile blurb) is included in the response.
func (s *listProjectMilestoneRequests) SetUserProfileDescription(val bool) *listProjectMilestoneRequests {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo controls whether the freelancer's display
// information (such as display name) is included in the response.
func (s *listProjectMilestoneRequests) SetUserDisplayInfo(val bool) *listProjectMilestoneRequests {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs controls whether the freelancer's job history
// is included in the response.
func (s *listProjectMilestoneRequests) SetUserJobs(val bool) *listProjectMilestoneRequests {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails controls whether the freelancer's currency
// balance details are included in the response.
func (s *listProjectMilestoneRequests) SetUserBalanceDetails(val bool) *listProjectMilestoneRequests {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails controls whether the freelancer's
// completed qualification exams are included in the response.
func (s *listProjectMilestoneRequests) SetUserQualificationDetails(val bool) *listProjectMilestoneRequests {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails controls whether the freelancer's
// membership information is included in the response.
func (s *listProjectMilestoneRequests) SetUserMembershipDetails(val bool) *listProjectMilestoneRequests {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails controls whether the freelancer's
// financial information is included in the response.
func (s *listProjectMilestoneRequests) SetUserFinancialDetails(val bool) *listProjectMilestoneRequests {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails controls whether the freelancer's
// location information is included in the response.
func (s *listProjectMilestoneRequests) SetUserLocationDetails(val bool) *listProjectMilestoneRequests {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails controls whether the freelancer's
// portfolio information is included in the response.
func (s *listProjectMilestoneRequests) SetUserPortfolioDetails(val bool) *listProjectMilestoneRequests {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails controls whether the freelancer's
// preferred information is included in the response.
func (s *listProjectMilestoneRequests) SetUserPreferredDetails(val bool) *listProjectMilestoneRequests {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails controls whether the freelancer's
// earned badges are included in the response.
func (s *listProjectMilestoneRequests) SetUserBadgeDetails(val bool) *listProjectMilestoneRequests {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus controls whether additional status
// information about the freelancer is included in the response.
func (s *listProjectMilestoneRequests) SetUserStatus(val bool) *listProjectMilestoneRequests {
	s.userStatus = &val
	return s
}

// SetUserReputation controls whether the freelancer's
// reputation summary is included in the response.
func (s *listProjectMilestoneRequests) SetUserReputation(val bool) *listProjectMilestoneRequests {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation controls whether the freelancer's
// employer reputation is included in the response.
func (s *listProjectMilestoneRequests) SetUserEmployerReputation(val bool) *listProjectMilestoneRequests {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra controls whether the freelancer's
// full reputation details are included in the response.
func (s *listProjectMilestoneRequests) SetUserReputationExtra(val bool) *listProjectMilestoneRequests {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra controls whether the freelancer's
// full employer reputation details are included in the response.
func (s *listProjectMilestoneRequests) SetUserEmployerReputationExtra(val bool) *listProjectMilestoneRequests {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage controls whether the freelancer's
// current cover image is included in the response.
func (s *listProjectMilestoneRequests) SetUserCoverImage(val bool) *listProjectMilestoneRequests {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage controls whether the freelancer's
// previous cover images are included in the response.
func (s *listProjectMilestoneRequests) SetUserPastCoverImage(val bool) *listProjectMilestoneRequests {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations controls whether the freelancer's
// recommendations count is included in the response.
func (s *listProjectMilestoneRequests) SetUserRecommendations(val bool) *listProjectMilestoneRequests {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness controls whether the freelancer's
// responsiveness score is included in the response.
func (s *listProjectMilestoneRequests) SetUserResponsiveness(val bool) *listProjectMilestoneRequests {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers controls whether corporate account
// information associated with the freelancer is included in the response.
func (s *listProjectMilestoneRequests) SetCorporateUsers(val bool) *listProjectMilestoneRequests {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber controls whether the freelancer's
// marketing mobile number used by recruiters is included in the response.
func (s *listProjectMilestoneRequests) SetMarketingMobileNumber(val bool) *listProjectMilestoneRequests {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails controls whether sanction details,
// including sanction end timestamps, are included in the response.
func (s *listProjectMilestoneRequests) SetSanctionDetails(val bool) *listProjectMilestoneRequests {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount controls whether limited account
// status information is included in the response.
func (s *listProjectMilestoneRequests) SetLimitedAccount(val bool) *listProjectMilestoneRequests {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails controls whether equipment groups
// and equipment items associated with the freelancer are included in the response.
func (s *listProjectMilestoneRequests) SetEquipmentGroupDetails(val bool) *listProjectMilestoneRequests {
	s.equipmentGroupDetails = &val
	return s
}

// TODO: refine with typed response
type getHourlyContractInfo struct {
	client            *Client
	projectIDs        []int64
	bidderIDs         []int64
	hourlyContractIDs []int
	projectOwnerIDs   []int64
	// teamsFilter  / not know what exactly this parameter do or is
	billingDetails *bool
	invoiceDetails *bool
}

func (s *getHourlyContractInfo) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_HOURLY_CONTRACTS),
	}

	for _, val := range s.projectIDs {
		r.addParam("project_ids[]", val)
	}
	for _, val := range s.bidderIDs {
		r.addParam("bidder_ids[]", val)
	}
	for _, val := range s.hourlyContractIDs {
		r.addParam("hourly_contract_ids[]", val)
	}
	for _, val := range s.projectOwnerIDs {
		r.addParam("project_owner_ids[]", val)
	}
	if s.billingDetails != nil {
		r.addParam("billing_details", *s.billingDetails)
	}
	if s.invoiceDetails != nil {
		r.addParam("invoice_details", *s.invoiceDetails)
	}

	return execute[*RawResponse](ctx, s.client, r)

}

// SetBidderIDs filters hourly contracts by the specified bidder (freelancer) IDs.
// This corresponds to the `bidder_ids[]` query parameter of the
func (s *getHourlyContractInfo) SetBidderIDs(vals []int64) *getHourlyContractInfo {
	s.bidderIDs = vals
	return s
}

// SetProjectOwnerIDs filters hourly contracts by the specified project owner IDs.
// This corresponds to the `project_owner_ids[]` query parameter of the
func (s *getHourlyContractInfo) SetProjectOwnerIDs(vals []int64) *getHourlyContractInfo {
	s.projectOwnerIDs = vals
	return s
}

// SetHourlyContractIDs filters the response to specific hourly contract IDs.
// This corresponds to the `hourly_contract_ids[]` query parameter of the
func (s *getHourlyContractInfo) SetHourlyContractIDs(vals []int) *getHourlyContractInfo {
	s.hourlyContractIDs = vals
	return s
}

// SetBillingDetails controls whether billing details are included in the response.
// When set to true, this enables the `billing_details` projection for the
func (s *getHourlyContractInfo) SetBillingDetails(val bool) *getHourlyContractInfo {
	s.billingDetails = &val
	return s
}

// SetInvoiceDetails controls whether invoice details are included in the response.
// When set to true, this enables the `invoice_details` projection for the
func (s *getHourlyContractInfo) SetInvoiceDetails(val bool) *getHourlyContractInfo {
	s.invoiceDetails = &val
	return s
}

// TODO: refine with typed response
type getIPContractInfo struct {
	client    *Client
	projectID int64
}

func (s *getIPContractInfo) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/ip_contracts", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := &RawResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type deleteProject struct {
	client    *Client
	projectID int64
}

func (s *deleteProject) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_PROJECTS), strconv.FormatInt(s.projectID, 10)),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

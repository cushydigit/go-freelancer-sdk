package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Return information about multiple projects. Will be ordered by descending submit data (newest-to-oldest)
type ListProjectsService struct {
	client                       *Client
	projects                     []int
	owners                       []int
	bidders                      []int
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
type ListProjectsResponse struct {
	Status    string         `json:"status"`
	RequsetID string         `json:"request_id"`
	Result    ProjectsResult `json:"result"`
}

type ProjectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}

// Do perform GET request on endpoint "projects/0.1/projects/"
func (s *ListProjectsService) Do(ctx context.Context) (*ListProjectsResponse, error) {
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

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListProjectsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *ListProjectsService) SetProjects(projects []int) *ListProjectsService {
	s.projects = projects
	return s
}

func (s *ListProjectsService) SetOwners(owners []int) *ListProjectsService {
	s.owners = owners
	return s
}

func (s *ListProjectsService) SetBidders(bidders []int) *ListProjectsService {
	s.bidders = bidders
	return s
}

func (s *ListProjectsService) SetSeoUrls(seoUrls []string) *ListProjectsService {
	s.seoUrls = seoUrls
	return s
}

func (s *ListProjectsService) SetFrontendProjectStatuses(val []string) *ListProjectsService {
	s.frontendProjectStatuses = val
	return s
}

func (s *ListProjectsService) SetFromTime(val int64) *ListProjectsService {
	s.fromTime = &val
	return s
}

func (s *ListProjectsService) SetToTime(val int64) *ListProjectsService {
	s.toTime = &val
	return s
}

func (s *ListProjectsService) SetFullDescription(val bool) *ListProjectsService {
	s.fullDescription = &val
	return s
}

func (s *ListProjectsService) SetJobDetails(val bool) *ListProjectsService {
	s.jobDetails = &val
	return s
}

func (s *ListProjectsService) SetUpgradeDetails(val bool) *ListProjectsService {
	s.upgradeDetails = &val
	return s
}

func (s *ListProjectsService) SetAttachmentDetails(val bool) *ListProjectsService {
	s.attachmentDetails = &val
	return s
}

func (s *ListProjectsService) SetFileDetails(val bool) *ListProjectsService {
	s.fileDetails = &val
	return s
}

func (s *ListProjectsService) SetQualificationDetails(val bool) *ListProjectsService {
	s.qualificationDetails = &val
	return s
}

func (s *ListProjectsService) SetSelectedBids(val bool) *ListProjectsService {
	s.selectedBids = &val
	return s
}

func (s *ListProjectsService) SetHireeDetails(val bool) *ListProjectsService {
	s.hiremeDetails = &val
	return s
}

func (s *ListProjectsService) SetUserDetails(val bool) *ListProjectsService {
	s.userDetails = &val
	return s
}

func (s *ListProjectsService) SetInvitedFreelancerDetails(val bool) *ListProjectsService {
	s.invitedFreelancerDetails = &val
	return s
}

func (s *ListProjectsService) SetRecommendedFreelancerDetails(val bool) *ListProjectsService {
	s.recommendedFreelancerDetails = &val
	return s
}

func (s *ListProjectsService) SetSupportSessionDetails(val bool) *ListProjectsService {
	s.supportSessionDetails = &val
	return s
}

func (s *ListProjectsService) SetLocationDetails(val bool) *ListProjectsService {
	s.locationDetails = &val
	return s
}

func (s *ListProjectsService) SetNdaSignatureDetails(val bool) *ListProjectsService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *ListProjectsService) SetDriveFileDetails(val bool) *ListProjectsService {
	s.driveFileDetails = &val
	return s
}

func (s *ListProjectsService) SetNdaDetails(val bool) *ListProjectsService {
	s.ndaDetails = &val
	return s
}

func (s *ListProjectsService) SetLocalDetails(val bool) *ListProjectsService {
	s.localDetails = &val
	return s
}

func (s *ListProjectsService) SetEquipmentDetails(val bool) *ListProjectsService {
	s.equipmentDetails = &val
	return s
}

func (s *ListProjectsService) SetClientEngagementDetails(val bool) *ListProjectsService {
	s.clientEngagementDetails = &val
	return s
}

func (s *ListProjectsService) SetUserResponsiveness(val bool) *ListProjectsService {
	s.userResponsiveness = &val
	return s
}

func (s *ListProjectsService) SetServiceOfferingDetails(val bool) *ListProjectsService {
	s.serviceOfferingDetails = &val
	return s
}

func (s *ListProjectsService) SetCorporateUsers(val bool) *ListProjectsService {
	s.corporateUsers = &val
	return s
}

func (s *ListProjectsService) SetIsNonHireMe(val bool) *ListProjectsService {
	s.isNonHireMe = &val
	return s
}

func (s *ListProjectsService) SetHasMilestone(val bool) *ListProjectsService {
	s.hasMilestone = &val
	return s
}

func (s *ListProjectsService) SetCount(val bool) *ListProjectsService {
	s.count = &val
	return s
}

func (s *ListProjectsService) SetTeam(val bool) *ListProjectsService {
	s.team = &val
	return s
}

func (s *ListProjectsService) SetCompact(val bool) *ListProjectsService {
	s.compact = &val
	return s
}

func (s *ListProjectsService) SetLimit(val int) *ListProjectsService {
	s.limit = &val
	return s
}

func (s *ListProjectsService) SetOffset(val int) *ListProjectsService {
	s.offset = &val
	return s
}

func (s *ListProjectsService) SetProximityDetails(val bool) *ListProjectsService {
	s.proximityDetails = &val
	return s
}

func (s *ListProjectsService) SetReviewAvailabilityDetails(val bool) *ListProjectsService {
	s.reviewAvailabilityDetails = &val
	return s
}

func (s *ListProjectsService) SetNegotiatedDetails(val bool) *ListProjectsService {
	s.negotiatedDetails = &val
	return s
}

func (s *ListProjectsService) SetUserAvatar(val bool) *ListProjectsService {
	s.userAvatar = &val
	return s
}

func (s *ListProjectsService) SetUserCountryDetails(val bool) *ListProjectsService {
	s.userCountryDetails = &val
	return s
}

func (s *ListProjectsService) SetUserProfileDescription(val bool) *ListProjectsService {
	s.userProfileDescription = &val
	return s
}

func (s *ListProjectsService) SetProjectCollaborationDetails(val bool) *ListProjectsService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *ListProjectsService) SetUserDisplayInfo(val bool) *ListProjectsService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListProjectsService) SetUserJobs(val bool) *ListProjectsService {
	s.userJobs = &val
	return s
}

func (s *ListProjectsService) SetUserBalanceDetails(val bool) *ListProjectsService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListProjectsService) SetUserQualificationDetails(val bool) *ListProjectsService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListProjectsService) SetUserMembershipDetails(val bool) *ListProjectsService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListProjectsService) SetUserFinancialDetails(val bool) *ListProjectsService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListProjectsService) SetUserLocationDetails(val bool) *ListProjectsService {
	s.userLocationDetails = &val
	return s
}

func (s *ListProjectsService) SetUserPortfolioDetails(val bool) *ListProjectsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListProjectsService) SetUserPreferredDetails(val bool) *ListProjectsService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListProjectsService) SetUserBadgeDetails(val bool) *ListProjectsService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListProjectsService) SetUserStatus(val bool) *ListProjectsService {
	s.userStatus = &val
	return s
}

func (s *ListProjectsService) SetUserReputation(val bool) *ListProjectsService {
	s.userReputation = &val
	return s
}

func (s *ListProjectsService) SetUserEmployerReputation(val bool) *ListProjectsService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListProjectsService) SetUserReputationExtra(val bool) *ListProjectsService {
	s.userReputationExtra = &val
	return s
}

func (s *ListProjectsService) SetUserEmployerReputationExtra(val bool) *ListProjectsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListProjectsService) SetUserCoverImage(val bool) *ListProjectsService {
	s.userCoverImage = &val
	return s
}

func (s *ListProjectsService) SetUserPastCoverImage(val bool) *ListProjectsService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListProjectsService) SetUserRecommendations(val bool) *ListProjectsService {
	s.userRecommendations = &val
	return s
}

func (s *ListProjectsService) SetMarketingMobileNumber(val bool) *ListProjectsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListProjectsService) SetSanctionDetails(val bool) *ListProjectsService {
	s.sanctionDetails = &val
	return s
}

func (s *ListProjectsService) SetLimitedAccount(val bool) *ListProjectsService {
	s.limitedAccount = &val
	return s
}

func (s *ListProjectsService) SetEquipmentGroupDetails(val bool) *ListProjectsService {
	s.equipmentGroupDetails = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type projectsListService struct {
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

func (s *projectsListService) Do(ctx context.Context) (*BaseResponse, error) {
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
	res := &BaseResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *projectsListService) SetProjects(projects []int64) *projectsListService {
	s.projects = projects
	return s
}

func (s *projectsListService) SetOwners(owners []int64) *projectsListService {
	s.owners = owners
	return s
}

func (s *projectsListService) SetBidders(bidders []int64) *projectsListService {
	s.bidders = bidders
	return s
}

func (s *projectsListService) SetSeoUrls(seoUrls []string) *projectsListService {
	s.seoUrls = seoUrls
	return s
}

func (s *projectsListService) SetFrontendProjectStatuses(val []string) *projectsListService {
	s.frontendProjectStatuses = val
	return s
}

func (s *projectsListService) SetFromTime(val int64) *projectsListService {
	s.fromTime = &val
	return s
}

func (s *projectsListService) SetToTime(val int64) *projectsListService {
	s.toTime = &val
	return s
}

func (s *projectsListService) SetFullDescription(val bool) *projectsListService {
	s.fullDescription = &val
	return s
}

func (s *projectsListService) SetJobDetails(val bool) *projectsListService {
	s.jobDetails = &val
	return s
}

func (s *projectsListService) SetUpgradeDetails(val bool) *projectsListService {
	s.upgradeDetails = &val
	return s
}

func (s *projectsListService) SetAttachmentDetails(val bool) *projectsListService {
	s.attachmentDetails = &val
	return s
}

func (s *projectsListService) SetFileDetails(val bool) *projectsListService {
	s.fileDetails = &val
	return s
}

func (s *projectsListService) SetQualificationDetails(val bool) *projectsListService {
	s.qualificationDetails = &val
	return s
}

func (s *projectsListService) SetSelectedBids(val bool) *projectsListService {
	s.selectedBids = &val
	return s
}

func (s *projectsListService) SetHireeDetails(val bool) *projectsListService {
	s.hiremeDetails = &val
	return s
}

func (s *projectsListService) SetUserDetails(val bool) *projectsListService {
	s.userDetails = &val
	return s
}

func (s *projectsListService) SetInvitedFreelancerDetails(val bool) *projectsListService {
	s.invitedFreelancerDetails = &val
	return s
}

func (s *projectsListService) SetRecommendedFreelancerDetails(val bool) *projectsListService {
	s.recommendedFreelancerDetails = &val
	return s
}

func (s *projectsListService) SetSupportSessionDetails(val bool) *projectsListService {
	s.supportSessionDetails = &val
	return s
}

func (s *projectsListService) SetLocationDetails(val bool) *projectsListService {
	s.locationDetails = &val
	return s
}

func (s *projectsListService) SetNdaSignatureDetails(val bool) *projectsListService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *projectsListService) SetDriveFileDetails(val bool) *projectsListService {
	s.driveFileDetails = &val
	return s
}

func (s *projectsListService) SetNdaDetails(val bool) *projectsListService {
	s.ndaDetails = &val
	return s
}

func (s *projectsListService) SetLocalDetails(val bool) *projectsListService {
	s.localDetails = &val
	return s
}

func (s *projectsListService) SetEquipmentDetails(val bool) *projectsListService {
	s.equipmentDetails = &val
	return s
}

func (s *projectsListService) SetClientEngagementDetails(val bool) *projectsListService {
	s.clientEngagementDetails = &val
	return s
}

func (s *projectsListService) SetUserResponsiveness(val bool) *projectsListService {
	s.userResponsiveness = &val
	return s
}

func (s *projectsListService) SetServiceOfferingDetails(val bool) *projectsListService {
	s.serviceOfferingDetails = &val
	return s
}

func (s *projectsListService) SetCorporateUsers(val bool) *projectsListService {
	s.corporateUsers = &val
	return s
}

func (s *projectsListService) SetIsNonHireMe(val bool) *projectsListService {
	s.isNonHireMe = &val
	return s
}

func (s *projectsListService) SetHasMilestone(val bool) *projectsListService {
	s.hasMilestone = &val
	return s
}

func (s *projectsListService) SetCount(val bool) *projectsListService {
	s.count = &val
	return s
}

func (s *projectsListService) SetTeam(val bool) *projectsListService {
	s.team = &val
	return s
}

func (s *projectsListService) SetCompact(val bool) *projectsListService {
	s.compact = &val
	return s
}

func (s *projectsListService) SetLimit(val int) *projectsListService {
	s.limit = &val
	return s
}

func (s *projectsListService) SetOffset(val int) *projectsListService {
	s.offset = &val
	return s
}

func (s *projectsListService) SetProximityDetails(val bool) *projectsListService {
	s.proximityDetails = &val
	return s
}

func (s *projectsListService) SetReviewAvailabilityDetails(val bool) *projectsListService {
	s.reviewAvailabilityDetails = &val
	return s
}

func (s *projectsListService) SetNegotiatedDetails(val bool) *projectsListService {
	s.negotiatedDetails = &val
	return s
}

func (s *projectsListService) SetUserAvatar(val bool) *projectsListService {
	s.userAvatar = &val
	return s
}

func (s *projectsListService) SetUserCountryDetails(val bool) *projectsListService {
	s.userCountryDetails = &val
	return s
}

func (s *projectsListService) SetUserProfileDescription(val bool) *projectsListService {
	s.userProfileDescription = &val
	return s
}

func (s *projectsListService) SetProjectCollaborationDetails(val bool) *projectsListService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *projectsListService) SetUserDisplayInfo(val bool) *projectsListService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectsListService) SetUserJobs(val bool) *projectsListService {
	s.userJobs = &val
	return s
}

func (s *projectsListService) SetUserBalanceDetails(val bool) *projectsListService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectsListService) SetUserQualificationDetails(val bool) *projectsListService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectsListService) SetUserMembershipDetails(val bool) *projectsListService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectsListService) SetUserFinancialDetails(val bool) *projectsListService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectsListService) SetUserLocationDetails(val bool) *projectsListService {
	s.userLocationDetails = &val
	return s
}

func (s *projectsListService) SetUserPortfolioDetails(val bool) *projectsListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectsListService) SetUserPreferredDetails(val bool) *projectsListService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectsListService) SetUserBadgeDetails(val bool) *projectsListService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectsListService) SetUserStatus(val bool) *projectsListService {
	s.userStatus = &val
	return s
}

func (s *projectsListService) SetUserReputation(val bool) *projectsListService {
	s.userReputation = &val
	return s
}

func (s *projectsListService) SetUserEmployerReputation(val bool) *projectsListService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectsListService) SetUserReputationExtra(val bool) *projectsListService {
	s.userReputationExtra = &val
	return s
}

func (s *projectsListService) SetUserEmployerReputationExtra(val bool) *projectsListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectsListService) SetUserCoverImage(val bool) *projectsListService {
	s.userCoverImage = &val
	return s
}

func (s *projectsListService) SetUserPastCoverImage(val bool) *projectsListService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectsListService) SetUserRecommendations(val bool) *projectsListService {
	s.userRecommendations = &val
	return s
}

func (s *projectsListService) SetMarketingMobileNumber(val bool) *projectsListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectsListService) SetSanctionDetails(val bool) *projectsListService {
	s.sanctionDetails = &val
	return s
}

func (s *projectsListService) SetLimitedAccount(val bool) *projectsListService {
	s.limitedAccount = &val
	return s
}

func (s *projectsListService) SetEquipmentGroupDetails(val bool) *projectsListService {
	s.equipmentGroupDetails = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListProjectsService struct {
	client                       *Client
	projects                     []int
	owners                       []int
	bidders                      []int
	seoUrls                      []string
	fromTime                     int64
	toTime                       int64
	frontendProjectStatuses      []string
	team                         bool
	isNonHireMe                  bool
	hasMilestone                 bool
	fullDescription              bool
	jobDetails                   bool
	upgradeDetails               bool
	attachmentDetails            bool
	fileDetails                  bool
	qualificationDetails         bool
	selectedBids                 bool
	hireMeDetails                bool
	userDetails                  bool
	invitedFreelancerDetails     bool
	recommendedFreelancerDetails bool
	supportSessionDetails        bool
	locationDetails              bool
	ndaSignatureDetails          bool
	projectCollaborationDetails  bool
	proximityDetails             bool
	reviewAvailabilityDetails    bool
	negotiatedDetails            bool
	driveFileDetails             bool
	ndaDetails                   bool
	localDetails                 bool
	equipmentDetails             bool
	clientEngagementDetails      bool
	serviceOfferingDetails       bool
	userAvatar                   bool
	userCountryDetails           bool
	userProfileDescription       bool
	userDisplayInfo              bool
	userJobs                     bool
	userBalanceDetails           bool
	userQualificationDetails     bool
	userMembershipDetails        bool
	userFinancialDetails         bool
	userLocationDetails          bool
	userPortfolioDetails         bool
	userPreferredDetails         bool
	userBadgeDetails             bool
	userStatus                   bool
	userReputation               bool
	userEmployerReputation       bool
	userEmployerReputationExtra  bool
	userCoverImage               bool
	userPastCoverImage           bool
	userRecommendations          bool
	userResponsiveness           bool
	corporateUsers               bool
	marketingMobileNumber        bool
	sanctionDetails              bool
	limitedAccount               bool
	equipmentGroupDetails        bool
	limit                        int
	offset                       int
	compact                      bool
}

// Do perform GET request on endpoint "projects/0.1/projects/"
func (s *ListProjectsService) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/projects",
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
	if s.seoUrls != nil {
		r.setParam("seo_urls", s.seoUrls)
	}
	if s.fromTime != 0 {
		r.setParam("from_time", s.fromTime)
	}
	if s.toTime != 0 {
		r.setParam("to_time", s.toTime)
	}
	if s.frontendProjectStatuses != nil {
		r.setParam("frontend_project_statuses", s.frontendProjectStatuses)
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

func (s *ListProjectsService) SetFromTime(fromTime int64) *ListProjectsService {
	s.fromTime = fromTime
	return s
}

func (s *ListProjectsService) SetToTime(toTime int64) *ListProjectsService {
	s.toTime = toTime
	return s
}

func (s *ListProjectsService) SetFullDescription(fullDescription bool) *ListProjectsService {
	s.fullDescription = fullDescription
	return s
}

func (s *ListProjectsService) SetJobDetails(jobDetails bool) *ListProjectsService {
	s.jobDetails = jobDetails
	return s
}

func (s *ListProjectsService) SetUpgradeDetails(upgradeDetails bool) *ListProjectsService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ListProjectsService) SetAttachmentDetails(attachmentDetails bool) *ListProjectsService {
	s.attachmentDetails = attachmentDetails
	return s
}

func (s *ListProjectsService) SetFileDetails(fileDetails bool) *ListProjectsService {
	s.fileDetails = fileDetails
	return s
}

func (s *ListProjectsService) SetQualificationDetails(qualificationDetails bool) *ListProjectsService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *ListProjectsService) SetSelectedBids(selectedBids bool) *ListProjectsService {
	s.selectedBids = selectedBids
	return s
}

func (s *ListProjectsService) SetHireMeDetails(hireMeDetails bool) *ListProjectsService {
	s.hireMeDetails = hireMeDetails
	return s
}

func (s *ListProjectsService) SetUserDetails(userDetails bool) *ListProjectsService {
	s.userDetails = userDetails
	return s
}

func (s *ListProjectsService) SetInvitedFreelancerDetails(invitedFreelancerDetails bool) *ListProjectsService {
	s.invitedFreelancerDetails = invitedFreelancerDetails
	return s
}

func (s *ListProjectsService) SetRecommendedFreelancerDetails(recommendedFreelancerDetails bool) *ListProjectsService {
	s.recommendedFreelancerDetails = recommendedFreelancerDetails
	return s
}

func (s *ListProjectsService) SetSupportSessionDetails(supportSessionDetails bool) *ListProjectsService {
	s.supportSessionDetails = supportSessionDetails
	return s
}

func (s *ListProjectsService) SetLocationDetails(locationDetails bool) *ListProjectsService {
	s.locationDetails = locationDetails
	return s
}

func (s *ListProjectsService) SetNdaSignatureDetails(ndaSignatureDetails bool) *ListProjectsService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ListProjectsService) SetDriveFileDetails(driveFileDetails bool) *ListProjectsService {
	s.driveFileDetails = driveFileDetails
	return s
}

func (s *ListProjectsService) SetNdaDetails(ndaDetails bool) *ListProjectsService {
	s.ndaDetails = ndaDetails
	return s
}

func (s *ListProjectsService) SetLocalDetails(localDetails bool) *ListProjectsService {
	s.localDetails = localDetails
	return s
}

func (s *ListProjectsService) SetEquipmentDetails(equipmentDetails bool) *ListProjectsService {
	s.equipmentDetails = equipmentDetails
	return s
}

func (s *ListProjectsService) SetClientEngagementDetails(clientEngagementDetails bool) *ListProjectsService {
	s.clientEngagementDetails = clientEngagementDetails
	return s
}

func (s *ListProjectsService) SetUserResponsiveness(userResponsiveness bool) *ListProjectsService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ListProjectsService) SetServiceOfferingDetails(serviceOfferingDetails bool) *ListProjectsService {
	s.serviceOfferingDetails = serviceOfferingDetails
	return s
}

func (s *ListProjectsService) SetCorporateUsers(corporateUsers bool) *ListProjectsService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ListProjectsService) SetIsNonHireMe(isNonHireMe bool) *ListProjectsService {
	s.isNonHireMe = isNonHireMe
	return s
}

func (s *ListProjectsService) SetHasMilestone(hasMilestone bool) *ListProjectsService {
	s.hasMilestone = hasMilestone
	return s
}

func (s *ListProjectsService) SetTeam(team bool) *ListProjectsService {
	s.team = team
	return s
}

func (s *ListProjectsService) SetCompact(compact bool) *ListProjectsService {
	s.compact = compact
	return s
}

func (s *ListProjectsService) SetLimit(limit int) *ListProjectsService {
	s.limit = limit
	return s
}

func (s *ListProjectsService) SetOffset(offset int) *ListProjectsService {
	s.offset = offset
	return s
}

func (s *ListProjectsService) SetFrontendProjectStatuses(frontendProjectStatuses []string) *ListProjectsService {
	s.frontendProjectStatuses = frontendProjectStatuses
	return s
}
func (s *ListProjectsService) SetProximityDetails(proximityDetails bool) *ListProjectsService {
	s.proximityDetails = proximityDetails
	return s
}

func (s *ListProjectsService) SetReviewAvailabilityDetails(reviewAvailabilityDetails bool) *ListProjectsService {
	s.reviewAvailabilityDetails = reviewAvailabilityDetails
	return s
}
func (s *ListProjectsService) SetNegotiatedDetails(negotiatedDetails bool) *ListProjectsService {
	s.negotiatedDetails = negotiatedDetails
	return s
}
func (s *ListProjectsService) SetUserAvatar(userAvatar bool) *ListProjectsService {
	s.userAvatar = userAvatar
	return s
}
func (s *ListProjectsService) SetUserCountryDetails(userCountryDetails bool) *ListProjectsService {
	s.userCountryDetails = userCountryDetails
	return s
}
func (s *ListProjectsService) SetUserProfileDescription(userProfileDescription bool) *ListProjectsService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ListProjectsService) SetProjectCollaborationDetails(projectCollaborationDetails bool) *ListProjectsService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ListProjectsService) SetUserDisplayInfo(userDisplayInfo bool) *ListProjectsService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ListProjectsService) SetUserJobs(userJobs bool) *ListProjectsService {
	s.userJobs = userJobs
	return s
}

func (s *ListProjectsService) SetUserBalanceDetails(userBalanceDetails bool) *ListProjectsService {
	s.userBalanceDetails = userBalanceDetails
	return s
}
func (s *ListProjectsService) SetUserQualificationDetails(userQualificationDetails bool) *ListProjectsService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ListProjectsService) SetUserMembershipDetails(userMembershipDetails bool) *ListProjectsService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ListProjectsService) SetUserFinancialDetails(userFinancialDetails bool) *ListProjectsService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ListProjectsService) SetUserLocationDetails(userLocationDetails bool) *ListProjectsService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ListProjectsService) SetUserPortfolioDetails(userPortfolioDetails bool) *ListProjectsService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ListProjectsService) SetUserPreferredDetails(userPreferredDetails bool) *ListProjectsService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ListProjectsService) SetUserBadgeDetails(userBadgeDetails bool) *ListProjectsService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ListProjectsService) SetUserStatus(userStatus bool) *ListProjectsService {
	s.userStatus = userStatus
	return s
}

func (s *ListProjectsService) SetUserReputation(userReputation bool) *ListProjectsService {
	s.userReputation = userReputation
	return s
}

func (s *ListProjectsService) SetUserEmployerReputation(userEmployerReputation bool) *ListProjectsService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ListProjectsService) SetUserEmployerReputationExtra(userEmployerReputationExtra bool) *ListProjectsService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ListProjectsService) SetUserCoverImage(userCoverImage bool) *ListProjectsService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ListProjectsService) SetUserPastCoverImage(userPastCoverImage bool) *ListProjectsService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ListProjectsService) SetUserRecommendations(userRecommendations bool) *ListProjectsService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ListProjectsService) SetMarketingMobileNumber(marketingMobileNumber bool) *ListProjectsService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ListProjectsService) SetSanctionDetails(sanctionDetails bool) *ListProjectsService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ListProjectsService) SetLimitedAccount(limitedAccount bool) *ListProjectsService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ListProjectsService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ListProjectsService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

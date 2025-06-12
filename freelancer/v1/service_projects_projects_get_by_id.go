package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Get information about a specific project.
type projectsGetByIDService struct {
	client                       *Client
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

func (s *projectsGetByIDService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

func (s *projectsGetByIDService) SetFullDescription(val bool) *projectsGetByIDService {
	s.fullDescription = &val
	return s
}

func (s *projectsGetByIDService) SetJobDetails(val bool) *projectsGetByIDService {
	s.jobDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUpgradeDetails(val bool) *projectsGetByIDService {
	s.upgradeDetails = &val
	return s
}

func (s *projectsGetByIDService) SetAttachmentDetails(val bool) *projectsGetByIDService {
	s.attachmentDetails = &val
	return s
}

func (s *projectsGetByIDService) SetFileDetails(val bool) *projectsGetByIDService {
	s.fileDetails = &val
	return s
}

func (s *projectsGetByIDService) SetQualificationDetails(val bool) *projectsGetByIDService {
	s.qualificationDetails = &val
	return s
}

func (s *projectsGetByIDService) SetSelectedBids(val bool) *projectsGetByIDService {
	s.selectedBids = &val
	return s
}

func (s *projectsGetByIDService) SetHiremeDetails(val bool) *projectsGetByIDService {
	s.hiremeDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserDetails(val bool) *projectsGetByIDService {
	s.userDetails = &val
	return s
}

func (s *projectsGetByIDService) SetInvitedFreelancerDetails(val bool) *projectsGetByIDService {
	s.invitedFreelancerDetails = &val
	return s
}

func (s *projectsGetByIDService) SetRecommendedFreelancerDetails(val bool) *projectsGetByIDService {
	s.recommendedFreelancerDetails = &val
	return s
}

func (s *projectsGetByIDService) SetHourlyDetails(val bool) *projectsGetByIDService {
	s.hourlyDetails = &val
	return s
}

func (s *projectsGetByIDService) SetSupportSessionDetails(val bool) *projectsGetByIDService {
	s.supportSessionDetails = &val
	return s
}

func (s *projectsGetByIDService) SetLocationDetails(val bool) *projectsGetByIDService {
	s.locationDetails = &val
	return s
}

func (s *projectsGetByIDService) SetNdaSignatureDetails(val bool) *projectsGetByIDService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *projectsGetByIDService) SetDriveFileDetails(val bool) *projectsGetByIDService {
	s.driveFileDetails = &val
	return s
}

func (s *projectsGetByIDService) SetNdaDetails(val bool) *projectsGetByIDService {
	s.ndaDetails = &val
	return s
}

func (s *projectsGetByIDService) SetLocalDetails(val bool) *projectsGetByIDService {
	s.localDetails = &val
	return s
}

func (s *projectsGetByIDService) SetEquipmentDetails(val bool) *projectsGetByIDService {
	s.equipmentDetails = &val
	return s
}

func (s *projectsGetByIDService) SetClientEngagementDetails(val bool) *projectsGetByIDService {
	s.clientEngagementDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserResponsiveness(val bool) *projectsGetByIDService {
	s.userResponsiveness = &val
	return s
}

func (s *projectsGetByIDService) SetServiceOfferingDetails(val bool) *projectsGetByIDService {
	s.serviceOfferingDetails = &val
	return s
}

func (s *projectsGetByIDService) SetCorporateUsers(val bool) *projectsGetByIDService {
	s.corporateUsers = &val
	return s
}

func (s *projectsGetByIDService) SetCompact(val bool) *projectsGetByIDService {
	s.compact = &val
	return s
}

func (s *projectsGetByIDService) SetLimit(val int) *projectsGetByIDService {
	s.limit = &val
	return s
}

func (s *projectsGetByIDService) SetOffset(val int) *projectsGetByIDService {
	s.offset = &val
	return s
}

func (s *projectsGetByIDService) SetProximityDetails(val bool) *projectsGetByIDService {
	s.proximityDetails = &val
	return s
}

func (s *projectsGetByIDService) SetReviewAvailabilityDetails(val bool) *projectsGetByIDService {
	s.reviewAvailabilityDetails = &val
	return s
}

func (s *projectsGetByIDService) SetNegotiatedDetails(val bool) *projectsGetByIDService {
	s.negotiatedDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserAvatar(val bool) *projectsGetByIDService {
	s.userAvatar = &val
	return s
}

func (s *projectsGetByIDService) SetUserCountryDetails(val bool) *projectsGetByIDService {
	s.userCountryDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserProfileDescription(val bool) *projectsGetByIDService {
	s.userProfileDescription = &val
	return s
}

func (s *projectsGetByIDService) SetProjectCollaborationDetails(val bool) *projectsGetByIDService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *projectsGetByIDService) SetTrackDetails(val bool) *projectsGetByIDService {
	s.trackDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserDisplayInfo(val bool) *projectsGetByIDService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectsGetByIDService) SetUserJobs(val bool) *projectsGetByIDService {
	s.userJobs = &val
	return s
}

func (s *projectsGetByIDService) SetUserBalanceDetails(val bool) *projectsGetByIDService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserQualificationDetails(val bool) *projectsGetByIDService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserMembershipDetails(val bool) *projectsGetByIDService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserFinancialDetails(val bool) *projectsGetByIDService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserLocationDetails(val bool) *projectsGetByIDService {
	s.userLocationDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserPortfolioDetails(val bool) *projectsGetByIDService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserPreferredDetails(val bool) *projectsGetByIDService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserBadgeDetails(val bool) *projectsGetByIDService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectsGetByIDService) SetUserStatus(val bool) *projectsGetByIDService {
	s.userStatus = &val
	return s
}

func (s *projectsGetByIDService) SetUserReputation(val bool) *projectsGetByIDService {
	s.userReputation = &val
	return s
}

func (s *projectsGetByIDService) SetUserEmployerReputation(val bool) *projectsGetByIDService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectsGetByIDService) SetUserReputationExtra(val bool) *projectsGetByIDService {
	s.userReputationExtra = &val
	return s
}

func (s *projectsGetByIDService) SetUserEmployerReputationExtra(val bool) *projectsGetByIDService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectsGetByIDService) SetUserCoverImage(val bool) *projectsGetByIDService {
	s.userCoverImage = &val
	return s
}

func (s *projectsGetByIDService) SetUserPastCoverImage(val bool) *projectsGetByIDService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectsGetByIDService) SetUserRecommendations(val bool) *projectsGetByIDService {
	s.userRecommendations = &val
	return s
}

func (s *projectsGetByIDService) SetMarketingMobileNumber(val bool) *projectsGetByIDService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectsGetByIDService) SetSanctionDetails(val bool) *projectsGetByIDService {
	s.sanctionDetails = &val
	return s
}

func (s *projectsGetByIDService) SetLimitedAccount(val bool) *projectsGetByIDService {
	s.limitedAccount = &val
	return s
}

func (s *projectsGetByIDService) SetEquipmentGroupDetails(val bool) *projectsGetByIDService {
	s.equipmentGroupDetails = &val
	return s
}

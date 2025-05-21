package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Get information about a specific project.
type GetProjectService struct {
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
	trackDetatils                *bool
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

type GetProjectResponse struct {
	Status    string  `json:"status"`
	RequsetID string  `json:"request_id"`
	Result    Project `json:"result"`
}

func (s *GetProjectService) Do(ctx context.Context, id int64) (*GetProjectResponse, error) {
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
	res := &GetProjectResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetProjectService) SetFullDescription(val bool) *GetProjectService {
	s.fullDescription = &val
	return s
}

func (s *GetProjectService) SetJobDetails(val bool) *GetProjectService {
	s.jobDetails = &val
	return s
}

func (s *GetProjectService) SetUpgradeDetails(val bool) *GetProjectService {
	s.upgradeDetails = &val
	return s
}

func (s *GetProjectService) SetAttachmentDetails(val bool) *GetProjectService {
	s.attachmentDetails = &val
	return s
}

func (s *GetProjectService) SetFileDetails(val bool) *GetProjectService {
	s.fileDetails = &val
	return s
}

func (s *GetProjectService) SetQualificationDetails(val bool) *GetProjectService {
	s.qualificationDetails = &val
	return s
}

func (s *GetProjectService) SetSelectedBids(val bool) *GetProjectService {
	s.selectedBids = &val
	return s
}

func (s *GetProjectService) SetHireeDetails(val bool) *GetProjectService {
	s.hiremeDetails = &val
	return s
}

func (s *GetProjectService) SetUserDetails(val bool) *GetProjectService {
	s.userDetails = &val
	return s
}

func (s *GetProjectService) SetInvitedFreelancerDetails(val bool) *GetProjectService {
	s.invitedFreelancerDetails = &val
	return s
}

func (s *GetProjectService) SetRecommendedFreelancerDetails(val bool) *GetProjectService {
	s.recommendedFreelancerDetails = &val
	return s
}

func (s *GetProjectService) SetHourlyDetails(val bool) *GetProjectService {
	s.hourlyDetails = &val
	return s
}

func (s *GetProjectService) SetSupportSessionDetails(val bool) *GetProjectService {
	s.supportSessionDetails = &val
	return s
}

func (s *GetProjectService) SetLocationDetails(val bool) *GetProjectService {
	s.locationDetails = &val
	return s
}

func (s *GetProjectService) SetNdaSignatureDetails(val bool) *GetProjectService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *GetProjectService) SetDriveFileDetails(val bool) *GetProjectService {
	s.driveFileDetails = &val
	return s
}

func (s *GetProjectService) SetNdaDetails(val bool) *GetProjectService {
	s.ndaDetails = &val
	return s
}

func (s *GetProjectService) SetLocalDetails(val bool) *GetProjectService {
	s.localDetails = &val
	return s
}

func (s *GetProjectService) SetEquipmentDetails(val bool) *GetProjectService {
	s.equipmentDetails = &val
	return s
}

func (s *GetProjectService) SetClientEngagementDetails(val bool) *GetProjectService {
	s.clientEngagementDetails = &val
	return s
}

func (s *GetProjectService) SetUserResponsiveness(val bool) *GetProjectService {
	s.userResponsiveness = &val
	return s
}

func (s *GetProjectService) SetServiceOfferingDetails(val bool) *GetProjectService {
	s.serviceOfferingDetails = &val
	return s
}

func (s *GetProjectService) SetCorporateUsers(val bool) *GetProjectService {
	s.corporateUsers = &val
	return s
}

func (s *GetProjectService) SetCompact(val bool) *GetProjectService {
	s.compact = &val
	return s
}

func (s *GetProjectService) SetLimit(val int) *GetProjectService {
	s.limit = &val
	return s
}

func (s *GetProjectService) SetOffset(val int) *GetProjectService {
	s.offset = &val
	return s
}

func (s *GetProjectService) SetProximityDetails(val bool) *GetProjectService {
	s.proximityDetails = &val
	return s
}

func (s *GetProjectService) SetReviewAvailabilityDetails(val bool) *GetProjectService {
	s.reviewAvailabilityDetails = &val
	return s
}

func (s *GetProjectService) SetNegotiatedDetails(val bool) *GetProjectService {
	s.negotiatedDetails = &val
	return s
}

func (s *GetProjectService) SetUserAvatar(val bool) *GetProjectService {
	s.userAvatar = &val
	return s
}

func (s *GetProjectService) SetUserCountryDetails(val bool) *GetProjectService {
	s.userCountryDetails = &val
	return s
}

func (s *GetProjectService) SetUserProfileDescription(val bool) *GetProjectService {
	s.userProfileDescription = &val
	return s
}

func (s *GetProjectService) SetProjectCollaborationDetails(val bool) *GetProjectService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *GetProjectService) SetTrackDetails(val bool) *GetProjectService {
	s.trackDetatils = &val
	return s
}

func (s *GetProjectService) SetUserDisplayInfo(val bool) *GetProjectService {
	s.userDisplayInfo = &val
	return s
}

func (s *GetProjectService) SetUserJobs(val bool) *GetProjectService {
	s.userJobs = &val
	return s
}

func (s *GetProjectService) SetUserBalanceDetails(val bool) *GetProjectService {
	s.userBalanceDetails = &val
	return s
}

func (s *GetProjectService) SetUserQualificationDetails(val bool) *GetProjectService {
	s.userQualificationDetails = &val
	return s
}

func (s *GetProjectService) SetUserMembershipDetails(val bool) *GetProjectService {
	s.userMembershipDetails = &val
	return s
}

func (s *GetProjectService) SetUserFinancialDetails(val bool) *GetProjectService {
	s.userFinancialDetails = &val
	return s
}

func (s *GetProjectService) SetUserLocationDetails(val bool) *GetProjectService {
	s.userLocationDetails = &val
	return s
}

func (s *GetProjectService) SetUserPortfolioDetails(val bool) *GetProjectService {
	s.userPortfolioDetails = &val
	return s
}

func (s *GetProjectService) SetUserPreferredDetails(val bool) *GetProjectService {
	s.userPreferredDetails = &val
	return s
}

func (s *GetProjectService) SetUserBadgeDetails(val bool) *GetProjectService {
	s.userBadgeDetails = &val
	return s
}

func (s *GetProjectService) SetUserStatus(val bool) *GetProjectService {
	s.userStatus = &val
	return s
}

func (s *GetProjectService) SetUserReputation(val bool) *GetProjectService {
	s.userReputation = &val
	return s
}

func (s *GetProjectService) SetUserEmployerReputation(val bool) *GetProjectService {
	s.userEmployerReputation = &val
	return s
}

func (s *GetProjectService) SetUserReputationExtra(val bool) *GetProjectService {
	s.userReputationExtra = &val
	return s
}

func (s *GetProjectService) SetUserEmployerReputationExtra(val bool) *GetProjectService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *GetProjectService) SetUserCoverImage(val bool) *GetProjectService {
	s.userCoverImage = &val
	return s
}

func (s *GetProjectService) SetUserPastCoverImage(val bool) *GetProjectService {
	s.userPastCoverImage = &val
	return s
}

func (s *GetProjectService) SetUserRecommendations(val bool) *GetProjectService {
	s.userRecommendations = &val
	return s
}

func (s *GetProjectService) SetMarketingMobileNumber(val bool) *GetProjectService {
	s.marketingMobileNumber = &val
	return s
}

func (s *GetProjectService) SetSanctionDetails(val bool) *GetProjectService {
	s.sanctionDetails = &val
	return s
}

func (s *GetProjectService) SetLimitedAccount(val bool) *GetProjectService {
	s.limitedAccount = &val
	return s
}

func (s *GetProjectService) SetEquipmentGroupDetails(val bool) *GetProjectService {
	s.equipmentGroupDetails = &val
	return s
}

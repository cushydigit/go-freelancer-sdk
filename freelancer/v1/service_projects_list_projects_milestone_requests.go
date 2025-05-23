package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns a list of milestone requests by freelancer for a project
type ListProjectMilestoneRequestsService struct {
	client                      *Client
	projectID                   *int64
	statuses                    []MilestoneStatusType
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

func (s *ListProjectMilestoneRequestsService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/milestone_requests", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
	}

	if s.projectID != nil {
		r.addParam("project_id", *s.projectID)
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

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	resp := &BaseResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp, nil

}

func (s *ListProjectMilestoneRequestsService) SetProjectID(val int64) *ListProjectMilestoneRequestsService {
	s.projectID = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetStatuses(values []MilestoneStatusType) *ListProjectMilestoneRequestsService {
	s.statuses = values
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserAvatar(val bool) *ListProjectMilestoneRequestsService {
	s.userAvatar = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserCountryDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userCountryDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserProfileDescription(val bool) *ListProjectMilestoneRequestsService {
	s.userProfileDescription = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserDisplayInfo(val bool) *ListProjectMilestoneRequestsService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserJobs(val bool) *ListProjectMilestoneRequestsService {
	s.userJobs = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserBalanceDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserQualificationDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserMembershipDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserFinancialDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserLocationDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userLocationDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserPortfolioDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserPreferredDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserBadgeDetails(val bool) *ListProjectMilestoneRequestsService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserStatus(val bool) *ListProjectMilestoneRequestsService {
	s.userStatus = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserReputation(val bool) *ListProjectMilestoneRequestsService {
	s.userReputation = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserEmployerReputation(val bool) *ListProjectMilestoneRequestsService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserReputationExtra(val bool) *ListProjectMilestoneRequestsService {
	s.userReputationExtra = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserEmployerReputationExtra(val bool) *ListProjectMilestoneRequestsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserCoverImage(val bool) *ListProjectMilestoneRequestsService {
	s.userCoverImage = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserPastCoverImage(val bool) *ListProjectMilestoneRequestsService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserRecommendations(val bool) *ListProjectMilestoneRequestsService {
	s.userRecommendations = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetUserResponsiveness(val bool) *ListProjectMilestoneRequestsService {
	s.userResponsiveness = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetCorporateUsers(val bool) *ListProjectMilestoneRequestsService {
	s.corporateUsers = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetMarketingMobileNumber(val bool) *ListProjectMilestoneRequestsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetSanctionDetails(val bool) *ListProjectMilestoneRequestsService {
	s.sanctionDetails = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetLimitedAccount(val bool) *ListProjectMilestoneRequestsService {
	s.limitedAccount = &val
	return s
}

func (s *ListProjectMilestoneRequestsService) SetEquipmentGroupDetails(val bool) *ListProjectMilestoneRequestsService {
	s.equipmentGroupDetails = &val
	return s
}

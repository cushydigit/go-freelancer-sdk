package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns a list of milestones on a project
type ListProjectMilestonesService struct {
	client                      *Client
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

func (s *ListProjectMilestonesService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/milestones", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

func (s *ListProjectMilestonesService) SetStatuses(values []MilestoneStatus) *ListProjectMilestonesService {
	s.statuses = values
	return s
}

func (s *ListProjectMilestonesService) SetUserAvatar(val bool) *ListProjectMilestonesService {
	s.userAvatar = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserCountryDetails(val bool) *ListProjectMilestonesService {
	s.userCountryDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserProfileDescription(val bool) *ListProjectMilestonesService {
	s.userProfileDescription = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserDisplayInfo(val bool) *ListProjectMilestonesService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserJobs(val bool) *ListProjectMilestonesService {
	s.userJobs = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserBalanceDetails(val bool) *ListProjectMilestonesService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserQualificationDetails(val bool) *ListProjectMilestonesService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserMembershipDetails(val bool) *ListProjectMilestonesService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserFinancialDetails(val bool) *ListProjectMilestonesService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserLocationDetails(val bool) *ListProjectMilestonesService {
	s.userLocationDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserPortfolioDetails(val bool) *ListProjectMilestonesService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserPreferredDetails(val bool) *ListProjectMilestonesService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserBadgeDetails(val bool) *ListProjectMilestonesService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserStatus(val bool) *ListProjectMilestonesService {
	s.userStatus = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserReputation(val bool) *ListProjectMilestonesService {
	s.userReputation = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserEmployerReputation(val bool) *ListProjectMilestonesService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserReputationExtra(val bool) *ListProjectMilestonesService {
	s.userReputationExtra = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserEmployerReputationExtra(val bool) *ListProjectMilestonesService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserCoverImage(val bool) *ListProjectMilestonesService {
	s.userCoverImage = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserPastCoverImage(val bool) *ListProjectMilestonesService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserRecommendations(val bool) *ListProjectMilestonesService {
	s.userRecommendations = &val
	return s
}

func (s *ListProjectMilestonesService) SetUserResponsiveness(val bool) *ListProjectMilestonesService {
	s.userResponsiveness = &val
	return s
}

func (s *ListProjectMilestonesService) SetCorporateUsers(val bool) *ListProjectMilestonesService {
	s.corporateUsers = &val
	return s
}

func (s *ListProjectMilestonesService) SetMarketingMobileNumber(val bool) *ListProjectMilestonesService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListProjectMilestonesService) SetSanctionDetails(val bool) *ListProjectMilestonesService {
	s.sanctionDetails = &val
	return s
}

func (s *ListProjectMilestonesService) SetLimitedAccount(val bool) *ListProjectMilestonesService {
	s.limitedAccount = &val
	return s
}

func (s *ListProjectMilestonesService) SetEquipmentGroupDetails(val bool) *ListProjectMilestonesService {
	s.equipmentGroupDetails = &val
	return s
}

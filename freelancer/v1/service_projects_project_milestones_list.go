package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type projectMilestonesListService struct {
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

func (s *projectMilestonesListService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
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

func (s *projectMilestonesListService) SetStatuses(values []MilestoneStatus) *projectMilestonesListService {
	s.statuses = values
	return s
}

func (s *projectMilestonesListService) SetUserAvatar(val bool) *projectMilestonesListService {
	s.userAvatar = &val
	return s
}

func (s *projectMilestonesListService) SetUserCountryDetails(val bool) *projectMilestonesListService {
	s.userCountryDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserProfileDescription(val bool) *projectMilestonesListService {
	s.userProfileDescription = &val
	return s
}

func (s *projectMilestonesListService) SetUserDisplayInfo(val bool) *projectMilestonesListService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectMilestonesListService) SetUserJobs(val bool) *projectMilestonesListService {
	s.userJobs = &val
	return s
}

func (s *projectMilestonesListService) SetUserBalanceDetails(val bool) *projectMilestonesListService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserQualificationDetails(val bool) *projectMilestonesListService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserMembershipDetails(val bool) *projectMilestonesListService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserFinancialDetails(val bool) *projectMilestonesListService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserLocationDetails(val bool) *projectMilestonesListService {
	s.userLocationDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserPortfolioDetails(val bool) *projectMilestonesListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserPreferredDetails(val bool) *projectMilestonesListService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserBadgeDetails(val bool) *projectMilestonesListService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectMilestonesListService) SetUserStatus(val bool) *projectMilestonesListService {
	s.userStatus = &val
	return s
}

func (s *projectMilestonesListService) SetUserReputation(val bool) *projectMilestonesListService {
	s.userReputation = &val
	return s
}

func (s *projectMilestonesListService) SetUserEmployerReputation(val bool) *projectMilestonesListService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectMilestonesListService) SetUserReputationExtra(val bool) *projectMilestonesListService {
	s.userReputationExtra = &val
	return s
}

func (s *projectMilestonesListService) SetUserEmployerReputationExtra(val bool) *projectMilestonesListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectMilestonesListService) SetUserCoverImage(val bool) *projectMilestonesListService {
	s.userCoverImage = &val
	return s
}

func (s *projectMilestonesListService) SetUserPastCoverImage(val bool) *projectMilestonesListService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectMilestonesListService) SetUserRecommendations(val bool) *projectMilestonesListService {
	s.userRecommendations = &val
	return s
}

func (s *projectMilestonesListService) SetUserResponsiveness(val bool) *projectMilestonesListService {
	s.userResponsiveness = &val
	return s
}

func (s *projectMilestonesListService) SetCorporateUsers(val bool) *projectMilestonesListService {
	s.corporateUsers = &val
	return s
}

func (s *projectMilestonesListService) SetMarketingMobileNumber(val bool) *projectMilestonesListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectMilestonesListService) SetSanctionDetails(val bool) *projectMilestonesListService {
	s.sanctionDetails = &val
	return s
}

func (s *projectMilestonesListService) SetLimitedAccount(val bool) *projectMilestonesListService {
	s.limitedAccount = &val
	return s
}

func (s *projectMilestonesListService) SetEquipmentGroupDetails(val bool) *projectMilestonesListService {
	s.equipmentGroupDetails = &val
	return s
}

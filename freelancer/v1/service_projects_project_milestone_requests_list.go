package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type projectMilestoneRequestsListService struct {
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

func (s *projectMilestoneRequestsListService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/milestone_requests", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
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

func (s *projectMilestoneRequestsListService) SetStatuses(values []MilestoneStatus) *projectMilestoneRequestsListService {
	s.statuses = values
	return s
}

func (s *projectMilestoneRequestsListService) SetUserAvatar(val bool) *projectMilestoneRequestsListService {
	s.userAvatar = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserCountryDetails(val bool) *projectMilestoneRequestsListService {
	s.userCountryDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserProfileDescription(val bool) *projectMilestoneRequestsListService {
	s.userProfileDescription = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserDisplayInfo(val bool) *projectMilestoneRequestsListService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserJobs(val bool) *projectMilestoneRequestsListService {
	s.userJobs = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserBalanceDetails(val bool) *projectMilestoneRequestsListService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserQualificationDetails(val bool) *projectMilestoneRequestsListService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserMembershipDetails(val bool) *projectMilestoneRequestsListService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserFinancialDetails(val bool) *projectMilestoneRequestsListService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserLocationDetails(val bool) *projectMilestoneRequestsListService {
	s.userLocationDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserPortfolioDetails(val bool) *projectMilestoneRequestsListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserPreferredDetails(val bool) *projectMilestoneRequestsListService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserBadgeDetails(val bool) *projectMilestoneRequestsListService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserStatus(val bool) *projectMilestoneRequestsListService {
	s.userStatus = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserReputation(val bool) *projectMilestoneRequestsListService {
	s.userReputation = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserEmployerReputation(val bool) *projectMilestoneRequestsListService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserReputationExtra(val bool) *projectMilestoneRequestsListService {
	s.userReputationExtra = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserEmployerReputationExtra(val bool) *projectMilestoneRequestsListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserCoverImage(val bool) *projectMilestoneRequestsListService {
	s.userCoverImage = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserPastCoverImage(val bool) *projectMilestoneRequestsListService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserRecommendations(val bool) *projectMilestoneRequestsListService {
	s.userRecommendations = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetUserResponsiveness(val bool) *projectMilestoneRequestsListService {
	s.userResponsiveness = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetCorporateUsers(val bool) *projectMilestoneRequestsListService {
	s.corporateUsers = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetMarketingMobileNumber(val bool) *projectMilestoneRequestsListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetSanctionDetails(val bool) *projectMilestoneRequestsListService {
	s.sanctionDetails = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetLimitedAccount(val bool) *projectMilestoneRequestsListService {
	s.limitedAccount = &val
	return s
}

func (s *projectMilestoneRequestsListService) SetEquipmentGroupDetails(val bool) *projectMilestoneRequestsListService {
	s.equipmentGroupDetails = &val
	return s
}

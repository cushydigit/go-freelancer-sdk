package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getMilestoneRequestService struct {
	client                      *Client
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

func (s *getMilestoneRequestService) Do(ctx context.Context, milestoneRequestID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONE_REQUESTS), milestoneRequestID),
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

func (s *getMilestoneRequestService) SetUserAvatar(val bool) *getMilestoneRequestService {
	s.userAvatar = &val
	return s
}

func (s *getMilestoneRequestService) SetUserCountryDetails(val bool) *getMilestoneRequestService {
	s.userCountryDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserProfileDescription(val bool) *getMilestoneRequestService {
	s.userProfileDescription = &val
	return s
}

func (s *getMilestoneRequestService) SetUserDisplayInfo(val bool) *getMilestoneRequestService {
	s.userDisplayInfo = &val
	return s
}

func (s *getMilestoneRequestService) SetUserJobs(val bool) *getMilestoneRequestService {
	s.userJobs = &val
	return s
}

func (s *getMilestoneRequestService) SetUserBalanceDetails(val bool) *getMilestoneRequestService {
	s.userBalanceDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserQualificationDetails(val bool) *getMilestoneRequestService {
	s.userQualificationDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserMembershipDetails(val bool) *getMilestoneRequestService {
	s.userMembershipDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserFinancialDetails(val bool) *getMilestoneRequestService {
	s.userFinancialDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserLocationDetails(val bool) *getMilestoneRequestService {
	s.userLocationDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserPortfolioDetails(val bool) *getMilestoneRequestService {
	s.userPortfolioDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserPreferredDetails(val bool) *getMilestoneRequestService {
	s.userPreferredDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserBadgeDetails(val bool) *getMilestoneRequestService {
	s.userBadgeDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetUserStatus(val bool) *getMilestoneRequestService {
	s.userStatus = &val
	return s
}

func (s *getMilestoneRequestService) SetUserReputation(val bool) *getMilestoneRequestService {
	s.userReputation = &val
	return s
}

func (s *getMilestoneRequestService) SetUserEmployerReputation(val bool) *getMilestoneRequestService {
	s.userEmployerReputation = &val
	return s
}

func (s *getMilestoneRequestService) SetUserReputationExtra(val bool) *getMilestoneRequestService {
	s.userReputationExtra = &val
	return s
}

func (s *getMilestoneRequestService) SetUserEmployerReputationExtra(val bool) *getMilestoneRequestService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *getMilestoneRequestService) SetUserCoverImage(val bool) *getMilestoneRequestService {
	s.userCoverImage = &val
	return s
}

func (s *getMilestoneRequestService) SetUserPastCoverImage(val bool) *getMilestoneRequestService {
	s.userPastCoverImage = &val
	return s
}

func (s *getMilestoneRequestService) SetUserRecommendations(val bool) *getMilestoneRequestService {
	s.userRecommendations = &val
	return s
}

func (s *getMilestoneRequestService) SetUserResponsiveness(val bool) *getMilestoneRequestService {
	s.userResponsiveness = &val
	return s
}

func (s *getMilestoneRequestService) SetCorporateUsers(val bool) *getMilestoneRequestService {
	s.corporateUsers = &val
	return s
}

func (s *getMilestoneRequestService) SetMarketingMobileNumber(val bool) *getMilestoneRequestService {
	s.marketingMobileNumber = &val
	return s
}

func (s *getMilestoneRequestService) SetSanctionDetails(val bool) *getMilestoneRequestService {
	s.sanctionDetails = &val
	return s
}

func (s *getMilestoneRequestService) SetLimitedAccount(val bool) *getMilestoneRequestService {
	s.limitedAccount = &val
	return s
}

func (s *getMilestoneRequestService) SetEquipmentGroupDetails(val bool) *getMilestoneRequestService {
	s.equipmentGroupDetails = &val
	return s
}

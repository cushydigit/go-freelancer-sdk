package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getMilestoneService struct {
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

func (s *getMilestoneService) Do(ctx context.Context, milestoneID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_MILESTONES), milestoneID),
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

func (s *getMilestoneService) SetUserAvatar(val bool) *getMilestoneService {
	s.userAvatar = &val
	return s
}

func (s *getMilestoneService) SetUserCountryDetails(val bool) *getMilestoneService {
	s.userCountryDetails = &val
	return s
}

func (s *getMilestoneService) SetUserProfileDescription(val bool) *getMilestoneService {
	s.userProfileDescription = &val
	return s
}

func (s *getMilestoneService) SetUserDisplayInfo(val bool) *getMilestoneService {
	s.userDisplayInfo = &val
	return s
}

func (s *getMilestoneService) SetUserJobs(val bool) *getMilestoneService {
	s.userJobs = &val
	return s
}

func (s *getMilestoneService) SetUserBalanceDetails(val bool) *getMilestoneService {
	s.userBalanceDetails = &val
	return s
}

func (s *getMilestoneService) SetUserQualificationDetails(val bool) *getMilestoneService {
	s.userQualificationDetails = &val
	return s
}

func (s *getMilestoneService) SetUserMembershipDetails(val bool) *getMilestoneService {
	s.userMembershipDetails = &val
	return s
}

func (s *getMilestoneService) SetUserFinancialDetails(val bool) *getMilestoneService {
	s.userFinancialDetails = &val
	return s
}

func (s *getMilestoneService) SetUserLocationDetails(val bool) *getMilestoneService {
	s.userLocationDetails = &val
	return s
}

func (s *getMilestoneService) SetUserPortfolioDetails(val bool) *getMilestoneService {
	s.userPortfolioDetails = &val
	return s
}

func (s *getMilestoneService) SetUserPreferredDetails(val bool) *getMilestoneService {
	s.userPreferredDetails = &val
	return s
}

func (s *getMilestoneService) SetUserBadgeDetails(val bool) *getMilestoneService {
	s.userBadgeDetails = &val
	return s
}

func (s *getMilestoneService) SetUserStatus(val bool) *getMilestoneService {
	s.userStatus = &val
	return s
}

func (s *getMilestoneService) SetUserReputation(val bool) *getMilestoneService {
	s.userReputation = &val
	return s
}

func (s *getMilestoneService) SetUserEmployerReputation(val bool) *getMilestoneService {
	s.userEmployerReputation = &val
	return s
}

func (s *getMilestoneService) SetUserReputationExtra(val bool) *getMilestoneService {
	s.userReputationExtra = &val
	return s
}

func (s *getMilestoneService) SetUserEmployerReputationExtra(val bool) *getMilestoneService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *getMilestoneService) SetUserCoverImage(val bool) *getMilestoneService {
	s.userCoverImage = &val
	return s
}

func (s *getMilestoneService) SetUserPastCoverImage(val bool) *getMilestoneService {
	s.userPastCoverImage = &val
	return s
}

func (s *getMilestoneService) SetUserRecommendations(val bool) *getMilestoneService {
	s.userRecommendations = &val
	return s
}

func (s *getMilestoneService) SetUserResponsiveness(val bool) *getMilestoneService {
	s.userResponsiveness = &val
	return s
}

func (s *getMilestoneService) SetCorporateUsers(val bool) *getMilestoneService {
	s.corporateUsers = &val
	return s
}

func (s *getMilestoneService) SetMarketingMobileNumber(val bool) *getMilestoneService {
	s.marketingMobileNumber = &val
	return s
}

func (s *getMilestoneService) SetSanctionDetails(val bool) *getMilestoneService {
	s.sanctionDetails = &val
	return s
}

func (s *getMilestoneService) SetLimitedAccount(val bool) *getMilestoneService {
	s.limitedAccount = &val
	return s
}

func (s *getMilestoneService) SetEquipmentGroupDetails(val bool) *getMilestoneService {
	s.equipmentGroupDetails = &val
	return s
}

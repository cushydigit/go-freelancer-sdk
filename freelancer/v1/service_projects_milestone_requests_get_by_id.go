package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type milestonesRequestsGetByIDService struct {
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

func (s *milestonesRequestsGetByIDService) Do(ctx context.Context, milestoneRequestID int) (*BaseResponse, error) {
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

func (s *milestonesRequestsGetByIDService) SetUserAvatar(val bool) *milestonesRequestsGetByIDService {
	s.userAvatar = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserCountryDetails(val bool) *milestonesRequestsGetByIDService {
	s.userCountryDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserProfileDescription(val bool) *milestonesRequestsGetByIDService {
	s.userProfileDescription = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserDisplayInfo(val bool) *milestonesRequestsGetByIDService {
	s.userDisplayInfo = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserJobs(val bool) *milestonesRequestsGetByIDService {
	s.userJobs = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserBalanceDetails(val bool) *milestonesRequestsGetByIDService {
	s.userBalanceDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserQualificationDetails(val bool) *milestonesRequestsGetByIDService {
	s.userQualificationDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserMembershipDetails(val bool) *milestonesRequestsGetByIDService {
	s.userMembershipDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserFinancialDetails(val bool) *milestonesRequestsGetByIDService {
	s.userFinancialDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserLocationDetails(val bool) *milestonesRequestsGetByIDService {
	s.userLocationDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserPortfolioDetails(val bool) *milestonesRequestsGetByIDService {
	s.userPortfolioDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserPreferredDetails(val bool) *milestonesRequestsGetByIDService {
	s.userPreferredDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserBadgeDetails(val bool) *milestonesRequestsGetByIDService {
	s.userBadgeDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserStatus(val bool) *milestonesRequestsGetByIDService {
	s.userStatus = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserReputation(val bool) *milestonesRequestsGetByIDService {
	s.userReputation = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserEmployerReputation(val bool) *milestonesRequestsGetByIDService {
	s.userEmployerReputation = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserReputationExtra(val bool) *milestonesRequestsGetByIDService {
	s.userReputationExtra = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserEmployerReputationExtra(val bool) *milestonesRequestsGetByIDService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserCoverImage(val bool) *milestonesRequestsGetByIDService {
	s.userCoverImage = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserPastCoverImage(val bool) *milestonesRequestsGetByIDService {
	s.userPastCoverImage = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserRecommendations(val bool) *milestonesRequestsGetByIDService {
	s.userRecommendations = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetUserResponsiveness(val bool) *milestonesRequestsGetByIDService {
	s.userResponsiveness = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetCorporateUsers(val bool) *milestonesRequestsGetByIDService {
	s.corporateUsers = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetMarketingMobileNumber(val bool) *milestonesRequestsGetByIDService {
	s.marketingMobileNumber = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetSanctionDetails(val bool) *milestonesRequestsGetByIDService {
	s.sanctionDetails = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetLimitedAccount(val bool) *milestonesRequestsGetByIDService {
	s.limitedAccount = &val
	return s
}

func (s *milestonesRequestsGetByIDService) SetEquipmentGroupDetails(val bool) *milestonesRequestsGetByIDService {
	s.equipmentGroupDetails = &val
	return s
}

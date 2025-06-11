package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns bids for a single project.
type getBidService struct {
	client                      *Client
	reputation                  *bool
	buyerProjectFee             *bool
	awardStatusPossibilities    *bool
	projectDetails              *bool
	userDetails                 *bool
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
	limit                       *int
	offset                      *int
	compact                     *bool
	quotations                  *bool
}

func (s *getBidService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%d", string(PROJECTS_BIDS), bidID),
	}

	if s.reputation != nil {
		r.addParam("reputation", *s.reputation)
	}
	if s.buyerProjectFee != nil {
		r.addParam("buyer_project_fee", *s.buyerProjectFee)
	}
	if s.awardStatusPossibilities != nil {
		r.addParam("award_status_possibilities", *s.awardStatusPossibilities)
	}
	if s.projectDetails != nil {
		r.addParam("project_details", *s.projectDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
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
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.setParam("compact", *s.compact)
	}
	if s.quotations != nil {
		r.setParam("quotations", *s.quotations)
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

func (s *getBidService) SetReputation(val bool) *getBidService {
	s.reputation = &val
	return s
}

func (s *getBidService) SetBuyerProjectFee(val bool) *getBidService {
	s.buyerProjectFee = &val
	return s
}

func (s *getBidService) SetAwardStatusPossibilities(val bool) *getBidService {
	s.awardStatusPossibilities = &val
	return s
}

func (s *getBidService) SetProjectDetails(val bool) *getBidService {
	s.projectDetails = &val
	return s
}

func (s *getBidService) SetUserDetails(val bool) *getBidService {
	s.userDetails = &val
	return s
}

func (s *getBidService) SetUserAvatar(val bool) *getBidService {
	s.userAvatar = &val
	return s
}

func (s *getBidService) SetUserCountryDetails(val bool) *getBidService {
	s.userCountryDetails = &val
	return s
}

func (s *getBidService) SetUserProfileDescription(val bool) *getBidService {
	s.userProfileDescription = &val
	return s
}

func (s *getBidService) SetUserDisplayInfo(val bool) *getBidService {
	s.userDisplayInfo = &val
	return s
}

func (s *getBidService) SetUserJobs(val bool) *getBidService {
	s.userJobs = &val
	return s
}

func (s *getBidService) SetUserBalanceDetails(val bool) *getBidService {
	s.userBalanceDetails = &val
	return s
}

func (s *getBidService) SetUserQualificationDetails(val bool) *getBidService {
	s.userQualificationDetails = &val
	return s
}

func (s *getBidService) SetUserMembershipDetails(val bool) *getBidService {
	s.userMembershipDetails = &val
	return s
}

func (s *getBidService) SetUserFinancialDetails(val bool) *getBidService {
	s.userFinancialDetails = &val
	return s
}

func (s *getBidService) SetUserLocationDetails(val bool) *getBidService {
	s.userLocationDetails = &val
	return s
}

func (s *getBidService) SetUserPortfolioDetails(val bool) *getBidService {
	s.userPortfolioDetails = &val
	return s
}

func (s *getBidService) SetUserPreferredDetails(val bool) *getBidService {
	s.userPreferredDetails = &val
	return s
}

func (s *getBidService) SetUserBadgeDetails(val bool) *getBidService {
	s.userBadgeDetails = &val
	return s
}

func (s *getBidService) SetUserStatus(val bool) *getBidService {
	s.userStatus = &val
	return s
}

func (s *getBidService) SetUserReputation(val bool) *getBidService {
	s.userReputation = &val
	return s
}

func (s *getBidService) SetUserEmployerReputation(val bool) *getBidService {
	s.userEmployerReputation = &val
	return s
}

func (s *getBidService) SetUserReputationExtra(val bool) *getBidService {
	s.userReputationExtra = &val
	return s
}

func (s *getBidService) SetUserEmployerReputationExtra(val bool) *getBidService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *getBidService) SetUserCoverImage(val bool) *getBidService {
	s.userCoverImage = &val
	return s
}

func (s *getBidService) SetUserPastCoverImage(val bool) *getBidService {
	s.userPastCoverImage = &val
	return s
}

func (s *getBidService) SetUserRecommendations(val bool) *getBidService {
	s.userRecommendations = &val
	return s
}

func (s *getBidService) SetUserResponsiveness(val bool) *getBidService {
	s.userResponsiveness = &val
	return s
}

func (s *getBidService) SetCorporateUsers(val bool) *getBidService {
	s.corporateUsers = &val
	return s
}

func (s *getBidService) SetMarketingMobileNumber(val bool) *getBidService {
	s.marketingMobileNumber = &val
	return s
}

func (s *getBidService) SetSanctionDetails(val bool) *getBidService {
	s.sanctionDetails = &val
	return s
}

func (s *getBidService) SetLimitedAccount(val bool) *getBidService {
	s.limitedAccount = &val
	return s
}

func (s *getBidService) SetEquipmentGroupDetails(val bool) *getBidService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *getBidService) SetLimit(val int) *getBidService {
	s.limit = &val
	return s
}

func (s *getBidService) SetOffset(val int) *getBidService {
	s.offset = &val
	return s
}

func (s *getBidService) SetCompact(val bool) *getBidService {
	s.compact = &val
	return s
}

func (s *getBidService) SetQuotations(val bool) *getBidService {
	s.quotations = &val
	return s
}

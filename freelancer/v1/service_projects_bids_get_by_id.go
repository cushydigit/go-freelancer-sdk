package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type bidsGetByIDService struct {
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

func (s *bidsGetByIDService) Do(ctx context.Context, bidID int) (*BaseResponse, error) {
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

func (s *bidsGetByIDService) SetReputation(val bool) *bidsGetByIDService {
	s.reputation = &val
	return s
}

func (s *bidsGetByIDService) SetBuyerProjectFee(val bool) *bidsGetByIDService {
	s.buyerProjectFee = &val
	return s
}

func (s *bidsGetByIDService) SetAwardStatusPossibilities(val bool) *bidsGetByIDService {
	s.awardStatusPossibilities = &val
	return s
}

func (s *bidsGetByIDService) SetProjectDetails(val bool) *bidsGetByIDService {
	s.projectDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserDetails(val bool) *bidsGetByIDService {
	s.userDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserAvatar(val bool) *bidsGetByIDService {
	s.userAvatar = &val
	return s
}

func (s *bidsGetByIDService) SetUserCountryDetails(val bool) *bidsGetByIDService {
	s.userCountryDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserProfileDescription(val bool) *bidsGetByIDService {
	s.userProfileDescription = &val
	return s
}

func (s *bidsGetByIDService) SetUserDisplayInfo(val bool) *bidsGetByIDService {
	s.userDisplayInfo = &val
	return s
}

func (s *bidsGetByIDService) SetUserJobs(val bool) *bidsGetByIDService {
	s.userJobs = &val
	return s
}

func (s *bidsGetByIDService) SetUserBalanceDetails(val bool) *bidsGetByIDService {
	s.userBalanceDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserQualificationDetails(val bool) *bidsGetByIDService {
	s.userQualificationDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserMembershipDetails(val bool) *bidsGetByIDService {
	s.userMembershipDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserFinancialDetails(val bool) *bidsGetByIDService {
	s.userFinancialDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserLocationDetails(val bool) *bidsGetByIDService {
	s.userLocationDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserPortfolioDetails(val bool) *bidsGetByIDService {
	s.userPortfolioDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserPreferredDetails(val bool) *bidsGetByIDService {
	s.userPreferredDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserBadgeDetails(val bool) *bidsGetByIDService {
	s.userBadgeDetails = &val
	return s
}

func (s *bidsGetByIDService) SetUserStatus(val bool) *bidsGetByIDService {
	s.userStatus = &val
	return s
}

func (s *bidsGetByIDService) SetUserReputation(val bool) *bidsGetByIDService {
	s.userReputation = &val
	return s
}

func (s *bidsGetByIDService) SetUserEmployerReputation(val bool) *bidsGetByIDService {
	s.userEmployerReputation = &val
	return s
}

func (s *bidsGetByIDService) SetUserReputationExtra(val bool) *bidsGetByIDService {
	s.userReputationExtra = &val
	return s
}

func (s *bidsGetByIDService) SetUserEmployerReputationExtra(val bool) *bidsGetByIDService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *bidsGetByIDService) SetUserCoverImage(val bool) *bidsGetByIDService {
	s.userCoverImage = &val
	return s
}

func (s *bidsGetByIDService) SetUserPastCoverImage(val bool) *bidsGetByIDService {
	s.userPastCoverImage = &val
	return s
}

func (s *bidsGetByIDService) SetUserRecommendations(val bool) *bidsGetByIDService {
	s.userRecommendations = &val
	return s
}

func (s *bidsGetByIDService) SetUserResponsiveness(val bool) *bidsGetByIDService {
	s.userResponsiveness = &val
	return s
}

func (s *bidsGetByIDService) SetCorporateUsers(val bool) *bidsGetByIDService {
	s.corporateUsers = &val
	return s
}

func (s *bidsGetByIDService) SetMarketingMobileNumber(val bool) *bidsGetByIDService {
	s.marketingMobileNumber = &val
	return s
}

func (s *bidsGetByIDService) SetSanctionDetails(val bool) *bidsGetByIDService {
	s.sanctionDetails = &val
	return s
}

func (s *bidsGetByIDService) SetLimitedAccount(val bool) *bidsGetByIDService {
	s.limitedAccount = &val
	return s
}

func (s *bidsGetByIDService) SetEquipmentGroupDetails(val bool) *bidsGetByIDService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *bidsGetByIDService) SetLimit(val int) *bidsGetByIDService {
	s.limit = &val
	return s
}

func (s *bidsGetByIDService) SetOffset(val int) *bidsGetByIDService {
	s.offset = &val
	return s
}

func (s *bidsGetByIDService) SetCompact(val bool) *bidsGetByIDService {
	s.compact = &val
	return s
}

func (s *bidsGetByIDService) SetQuotations(val bool) *bidsGetByIDService {
	s.quotations = &val
	return s
}

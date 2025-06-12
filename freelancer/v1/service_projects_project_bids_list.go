package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type projectBidsListService struct {
	client                      *Client
	isShortlisted               *bool
	reputation                  *bool
	recommendedBid              *bool
	shortlistedBid              *bool
	distance                    *bool
	userDetails                 *bool
	expertGuarantees            *bool
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

func (s *projectBidsListService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s/bids", string(PROJECTS_PROJECTS), strconv.FormatInt(id, 10)),
	}

	if s.isShortlisted != nil {
		r.addParam("is_shortlisted", *s.isShortlisted)
	}
	if s.reputation != nil {
		r.addParam("reputation", *s.reputation)
	}
	if s.recommendedBid != nil {
		r.addParam("recommended_bid", *s.recommendedBid)
	}
	if s.shortlistedBid != nil {
		r.addParam("shortlisted_bid", *s.shortlistedBid)
	}
	if s.distance != nil {
		r.addParam("distance", *s.distance)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.expertGuarantees != nil {
		r.addParam("expert_guarantees", *s.expertGuarantees)
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

func (s *projectBidsListService) SetIsShortlisted(val bool) *projectBidsListService {
	s.isShortlisted = &val
	return s
}

func (s *projectBidsListService) SetReputation(val bool) *projectBidsListService {
	s.reputation = &val
	return s
}

func (s *projectBidsListService) SetRecommendedBid(val bool) *projectBidsListService {
	s.recommendedBid = &val
	return s
}

func (s *projectBidsListService) SetShortlistedBid(val bool) *projectBidsListService {
	s.shortlistedBid = &val
	return s
}

func (s *projectBidsListService) SetDistance(val bool) *projectBidsListService {
	s.distance = &val
	return s
}

func (s *projectBidsListService) SetUserDetails(val bool) *projectBidsListService {
	s.userDetails = &val
	return s
}

func (s *projectBidsListService) SetExpertGuarantees(val bool) *projectBidsListService {
	s.expertGuarantees = &val
	return s
}

func (s *projectBidsListService) SetUserAvatar(val bool) *projectBidsListService {
	s.userAvatar = &val
	return s
}

func (s *projectBidsListService) SetUserCountryDetails(val bool) *projectBidsListService {
	s.userCountryDetails = &val
	return s
}

func (s *projectBidsListService) SetUserProfileDescription(val bool) *projectBidsListService {
	s.userProfileDescription = &val
	return s
}

func (s *projectBidsListService) SetUserDisplayInfo(val bool) *projectBidsListService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectBidsListService) SetUserJobs(val bool) *projectBidsListService {
	s.userJobs = &val
	return s
}

func (s *projectBidsListService) SetUserBalanceDetails(val bool) *projectBidsListService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectBidsListService) SetUserQualificationDetails(val bool) *projectBidsListService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectBidsListService) SetUserMembershipDetails(val bool) *projectBidsListService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectBidsListService) SetUserFinancialDetails(val bool) *projectBidsListService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectBidsListService) SetUserLocationDetails(val bool) *projectBidsListService {
	s.userLocationDetails = &val
	return s
}

func (s *projectBidsListService) SetUserPortfolioDetails(val bool) *projectBidsListService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectBidsListService) SetUserPreferredDetails(val bool) *projectBidsListService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectBidsListService) SetUserBadgeDetails(val bool) *projectBidsListService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectBidsListService) SetUserStatus(val bool) *projectBidsListService {
	s.userStatus = &val
	return s
}

func (s *projectBidsListService) SetUserReputation(val bool) *projectBidsListService {
	s.userReputation = &val
	return s
}

func (s *projectBidsListService) SetUserEmployerReputation(val bool) *projectBidsListService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectBidsListService) SetUserReputationExtra(val bool) *projectBidsListService {
	s.userReputationExtra = &val
	return s
}

func (s *projectBidsListService) SetUserEmployerReputationExtra(val bool) *projectBidsListService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectBidsListService) SetUserCoverImage(val bool) *projectBidsListService {
	s.userCoverImage = &val
	return s
}

func (s *projectBidsListService) SetUserPastCoverImage(val bool) *projectBidsListService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectBidsListService) SetUserRecommendations(val bool) *projectBidsListService {
	s.userRecommendations = &val
	return s
}

func (s *projectBidsListService) SetUserResponsiveness(val bool) *projectBidsListService {
	s.userResponsiveness = &val
	return s
}

func (s *projectBidsListService) SetCorporateUsers(val bool) *projectBidsListService {
	s.corporateUsers = &val
	return s
}

func (s *projectBidsListService) SetMarketingMobileNumber(val bool) *projectBidsListService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectBidsListService) SetSanctionDetails(val bool) *projectBidsListService {
	s.sanctionDetails = &val
	return s
}

func (s *projectBidsListService) SetLimitedAccount(val bool) *projectBidsListService {
	s.limitedAccount = &val
	return s
}

func (s *projectBidsListService) SetEquipmentGroupDetails(val bool) *projectBidsListService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *projectBidsListService) SetLimit(val int) *projectBidsListService {
	s.limit = &val
	return s
}

func (s *projectBidsListService) SetOffset(val int) *projectBidsListService {
	s.offset = &val
	return s
}

func (s *projectBidsListService) SetCompact(val bool) *projectBidsListService {
	s.compact = &val
	return s
}

func (s *projectBidsListService) SetQuotations(val bool) *projectBidsListService {
	s.quotations = &val
	return s
}

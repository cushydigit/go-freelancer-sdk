package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Returns bids for a single project.
type ListProjectBidsService struct {
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

func (s *ListProjectBidsService) Do(ctx context.Context, id int64) (*BaseResponse, error) {
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

func (s *ListProjectBidsService) SetIsShortlisted(val bool) *ListProjectBidsService {
	s.isShortlisted = &val
	return s
}

func (s *ListProjectBidsService) SetReputation(val bool) *ListProjectBidsService {
	s.reputation = &val
	return s
}

func (s *ListProjectBidsService) SetRecommendedBid(val bool) *ListProjectBidsService {
	s.recommendedBid = &val
	return s
}

func (s *ListProjectBidsService) SetShortlistedBid(val bool) *ListProjectBidsService {
	s.shortlistedBid = &val
	return s
}

func (s *ListProjectBidsService) SetDistance(val bool) *ListProjectBidsService {
	s.distance = &val
	return s
}

func (s *ListProjectBidsService) SetUserDetails(val bool) *ListProjectBidsService {
	s.userDetails = &val
	return s
}

func (s *ListProjectBidsService) SetExpertGuarantess(val bool) *ListProjectBidsService {
	s.expertGuarantees = &val
	return s
}

func (s *ListProjectBidsService) SetUserAvatar(val bool) *ListProjectBidsService {
	s.userAvatar = &val
	return s
}

func (s *ListProjectBidsService) SetUserCountryDetails(val bool) *ListProjectBidsService {
	s.userCountryDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserProfileDescription(val bool) *ListProjectBidsService {
	s.userProfileDescription = &val
	return s
}

func (s *ListProjectBidsService) SetUserDisplayInfo(val bool) *ListProjectBidsService {
	s.userDisplayInfo = &val
	return s
}

func (s *ListProjectBidsService) SetUserJobs(val bool) *ListProjectBidsService {
	s.userJobs = &val
	return s
}

func (s *ListProjectBidsService) SetUserBalanceDetails(val bool) *ListProjectBidsService {
	s.userBalanceDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserQualificationDetails(val bool) *ListProjectBidsService {
	s.userQualificationDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserMembershipDetails(val bool) *ListProjectBidsService {
	s.userMembershipDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserFinancialDetails(val bool) *ListProjectBidsService {
	s.userFinancialDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserLocationDetails(val bool) *ListProjectBidsService {
	s.userLocationDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserPortfolioDetails(val bool) *ListProjectBidsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserPreferredDetails(val bool) *ListProjectBidsService {
	s.userPreferredDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserBadgeDetails(val bool) *ListProjectBidsService {
	s.userBadgeDetails = &val
	return s
}

func (s *ListProjectBidsService) SetUserStatus(val bool) *ListProjectBidsService {
	s.userStatus = &val
	return s
}

func (s *ListProjectBidsService) SetUserReputation(val bool) *ListProjectBidsService {
	s.userReputation = &val
	return s
}

func (s *ListProjectBidsService) SetUserEmployerReputation(val bool) *ListProjectBidsService {
	s.userEmployerReputation = &val
	return s
}

func (s *ListProjectBidsService) SetUserReputationExtra(val bool) *ListProjectBidsService {
	s.userReputationExtra = &val
	return s
}

func (s *ListProjectBidsService) SetUserEmployerReputationExtra(val bool) *ListProjectBidsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *ListProjectBidsService) SetUserCoverImage(val bool) *ListProjectBidsService {
	s.userCoverImage = &val
	return s
}

func (s *ListProjectBidsService) SetUserPastCoverImage(val bool) *ListProjectBidsService {
	s.userPastCoverImage = &val
	return s
}

func (s *ListProjectBidsService) SetUserRecommendations(val bool) *ListProjectBidsService {
	s.userRecommendations = &val
	return s
}

func (s *ListProjectBidsService) SetUserResponsiveness(val bool) *ListProjectBidsService {
	s.userResponsiveness = &val
	return s
}

func (s *ListProjectBidsService) SetCorporateUsers(val bool) *ListProjectBidsService {
	s.corporateUsers = &val
	return s
}

func (s *ListProjectBidsService) SetMarketingMobileNumber(val bool) *ListProjectBidsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListProjectBidsService) SetSanctionDetails(val bool) *ListProjectBidsService {
	s.sanctionDetails = &val
	return s
}

func (s *ListProjectBidsService) SetLimitedAccount(val bool) *ListProjectBidsService {
	s.limitedAccount = &val
	return s
}

func (s *ListProjectBidsService) SetEquipmentGroupDetails(val bool) *ListProjectBidsService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *ListProjectBidsService) SetLimit(val int) *ListProjectBidsService {
	s.limit = &val
	return s
}

func (s *ListProjectBidsService) SetOffset(val int) *ListProjectBidsService {
	s.offset = &val
	return s
}

func (s *ListProjectBidsService) SetCompact(val bool) *ListProjectBidsService {
	s.compact = &val
	return s
}

func (s *ListProjectBidsService) SetQuotations(val bool) *ListProjectBidsService {
	s.quotations = &val
	return s
}

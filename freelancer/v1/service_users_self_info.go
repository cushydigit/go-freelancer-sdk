package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Return information about current user
type GetSelfInfoService struct {
	client                        *Client
	avatar                        *bool
	countryDetails                *bool
	profileDescription            *bool
	displayInfo                   *bool
	jobs                          *bool
	balanceDetails                *bool
	qualificationDetails          *bool
	membershipDetails             *bool
	financialDetails              *bool
	locationDetails               *bool
	portfolioDetails              *bool
	preferredDetails              *bool
	badgeDetails                  *bool
	status                        *bool
	reputation                    *bool
	employerReputation            *bool
	reputationExtra               *bool
	employerReputationExtra       *bool
	coverImage                    *bool
	pastCoverImage                *bool
	pastCoverImages               *bool
	mobileTracking                *bool
	bidQualityDetails             *bool
	depositMethods                *bool
	userRecommendations           *bool
	marketingMobileNumber         *bool
	sanctionDetails               *bool
	limitedAccount                *bool
	completedUserRelevantJobCount *bool
	equipmentGroupDetails         *bool
	jobRanks                      *bool
	jobSeoDetails                 *bool
	risingStar                    *bool
	shareholderDetails            *bool
	staffDetails                  *bool
	limit                         *int
	offset                        *int
	compact                       *bool
}

type GetSelfInfoResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

func (s *GetSelfInfoService) DO(ctx context.Context) (*GetSelfInfoResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF),
	}

	if s.avatar != nil {
		r.setParam("avatar", *s.avatar)
	}
	if s.countryDetails != nil {
		r.setParam("country_details", *s.countryDetails)
	}
	if s.profileDescription != nil {
		r.setParam("profile_description", *s.profileDescription)
	}
	if s.displayInfo != nil {
		r.setParam("display_info", *s.displayInfo)
	}
	if s.jobs != nil {
		r.setParam("jobs", *s.jobs)
	}
	if s.balanceDetails != nil {
		r.setParam("balance_details", *s.balanceDetails)
	}
	if s.qualificationDetails != nil {
		r.setParam("qualification_details", *s.qualificationDetails)
	}
	if s.membershipDetails != nil {
		r.setParam("membership_details", *s.membershipDetails)
	}
	if s.financialDetails != nil {
		r.setParam("financial_details", *s.financialDetails)
	}
	if s.locationDetails != nil {
		r.setParam("location_details", *s.locationDetails)
	}
	if s.portfolioDetails != nil {
		r.setParam("portfolio_details", *s.portfolioDetails)
	}
	if s.preferredDetails != nil {
		r.setParam("preferred_details", *s.preferredDetails)
	}
	if s.badgeDetails != nil {
		r.setParam("badge_details", *s.badgeDetails)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.reputation != nil {
		r.setParam("reputation", *s.reputation)
	}
	if s.employerReputation != nil {
		r.setParam("employer_reputation", *s.employerReputation)
	}
	if s.reputationExtra != nil {
		r.setParam("reputation_extra", *s.reputationExtra)
	}
	if s.employerReputationExtra != nil {
		r.setParam("employer_reputation_extra", *s.employerReputationExtra)
	}
	if s.coverImage != nil {
		r.setParam("cover_image", *s.coverImage)
	}
	if s.pastCoverImage != nil {
		r.setParam("past_cover_image", *s.pastCoverImage)
	}
	if s.mobileTracking != nil {
		r.setParam("mobile_tracking", *s.mobileTracking)
	}
	if s.bidQualityDetails != nil {
		r.setParam("bid_quality_details", *s.bidQualityDetails)
	}
	if s.depositMethods != nil {
		r.setParam("deposit_methods", *s.depositMethods)
	}
	if s.userRecommendations != nil {
		r.setParam("user_recommendations", *s.userRecommendations)
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
	if s.completedUserRelevantJobCount != nil {
		r.setParam("completed_user_relevant_job_count", *s.completedUserRelevantJobCount)
	}
	if s.equipmentGroupDetails != nil {
		r.setParam("equipment_group_details", *s.equipmentGroupDetails)
	}
	if s.jobRanks != nil {
		r.setParam("job_ranks", *s.jobRanks)
	}
	if s.jobSeoDetails != nil {
		r.setParam("job_seo_details", *s.jobSeoDetails)
	}
	if s.risingStar != nil {
		r.setParam("rising_star", *s.risingStar)
	}
	if s.shareholderDetails != nil {
		r.setParam("shareholder_details", *s.shareholderDetails)
	}
	if s.staffDetails != nil {
		r.setParam("staff_details", *s.staffDetails)
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &GetSelfInfoResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil

}

func (s *GetSelfInfoService) SetAvatar(val bool) *GetSelfInfoService {
	s.avatar = &val
	return s
}

func (s *GetSelfInfoService) SetCountryDetails(val bool) *GetSelfInfoService {
	s.countryDetails = &val
	return s
}

func (s *GetSelfInfoService) SetProfileDescription(val bool) *GetSelfInfoService {
	s.profileDescription = &val
	return s
}

func (s *GetSelfInfoService) SetDisplayInfo(val bool) *GetSelfInfoService {
	s.displayInfo = &val
	return s
}

func (s *GetSelfInfoService) SetJobs(val bool) *GetSelfInfoService {
	s.jobs = &val
	return s
}

func (s *GetSelfInfoService) SetBalanceDetails(val bool) *GetSelfInfoService {
	s.balanceDetails = &val
	return s
}

func (s *GetSelfInfoService) SetQualificationDetails(val bool) *GetSelfInfoService {
	s.qualificationDetails = &val
	return s
}

func (s *GetSelfInfoService) SetMembershipDetails(val bool) *GetSelfInfoService {
	s.membershipDetails = &val
	return s
}

func (s *GetSelfInfoService) SetFinancialDetails(val bool) *GetSelfInfoService {
	s.financialDetails = &val
	return s
}

func (s *GetSelfInfoService) SetLocationDetails(val bool) *GetSelfInfoService {
	s.locationDetails = &val
	return s
}

func (s *GetSelfInfoService) SetPortfolioDetails(val bool) *GetSelfInfoService {
	s.portfolioDetails = &val
	return s
}

func (s *GetSelfInfoService) SetPreferredDetails(val bool) *GetSelfInfoService {
	s.preferredDetails = &val
	return s
}

func (s *GetSelfInfoService) SetBadgeDetails(val bool) *GetSelfInfoService {
	s.badgeDetails = &val
	return s
}

func (s *GetSelfInfoService) SetStatus(val bool) *GetSelfInfoService {
	s.status = &val
	return s
}

func (s *GetSelfInfoService) SetReputation(val bool) *GetSelfInfoService {
	s.reputation = &val
	return s
}

func (s *GetSelfInfoService) SetEmployerReputation(val bool) *GetSelfInfoService {
	s.employerReputation = &val
	return s
}

func (s *GetSelfInfoService) SetReputationExtra(val bool) *GetSelfInfoService {
	s.reputationExtra = &val
	return s
}

func (s *GetSelfInfoService) SetEmployerReputationExtra(val bool) *GetSelfInfoService {
	s.employerReputationExtra = &val
	return s
}

func (s *GetSelfInfoService) SetCoverImage(val bool) *GetSelfInfoService {
	s.coverImage = &val
	return s
}

func (s *GetSelfInfoService) SetPastCoverImage(val bool) *GetSelfInfoService {
	s.pastCoverImage = &val
	return s
}

func (s *GetSelfInfoService) SetMobileTracking(val bool) *GetSelfInfoService {
	s.mobileTracking = &val
	return s
}

func (s *GetSelfInfoService) SetBidQualityDetails(val bool) *GetSelfInfoService {
	s.bidQualityDetails = &val
	return s
}

func (s *GetSelfInfoService) SetDepositMethods(val bool) *GetSelfInfoService {
	s.depositMethods = &val
	return s
}

func (s *GetSelfInfoService) SetUserRecommendations(val bool) *GetSelfInfoService {
	s.userRecommendations = &val
	return s
}

func (s *GetSelfInfoService) SetMarketingMobileNumber(val bool) *GetSelfInfoService {
	s.marketingMobileNumber = &val
	return s
}

func (s *GetSelfInfoService) SetSanctionDetails(val bool) *GetSelfInfoService {
	s.sanctionDetails = &val
	return s
}

func (s *GetSelfInfoService) SetLimitedAccount(val bool) *GetSelfInfoService {
	s.limitedAccount = &val
	return s
}

func (s *GetSelfInfoService) SetCompletedUserRelevantJobCount(val bool) *GetSelfInfoService {
	s.completedUserRelevantJobCount = &val
	return s
}

func (s *GetSelfInfoService) SetEquipmentGroupDetails(val bool) *GetSelfInfoService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *GetSelfInfoService) SetJobRanks(val bool) *GetSelfInfoService {
	s.jobRanks = &val
	return s
}

func (s *GetSelfInfoService) SetShareholderDetails(val bool) *GetSelfInfoService {
	s.shareholderDetails = &val
	return s
}

func (s *GetSelfInfoService) SetRisingStar(val bool) *GetSelfInfoService {
	s.risingStar = &val
	return s
}

func (s *GetSelfInfoService) SetStaffDetails(val bool) *GetSelfInfoService {
	s.staffDetails = &val
	return s
}

func (s *GetSelfInfoService) SetJobSeoDetails(val bool) *GetSelfInfoService {
	s.jobSeoDetails = &val
	return s
}

func (s *GetSelfInfoService) SetLimit(val int) *GetSelfInfoService {
	s.limit = &val
	return s
}

func (s *GetSelfInfoService) SetOffset(val int) *GetSelfInfoService {
	s.offset = &val
	return s
}

func (s *GetSelfInfoService) SetCompact(val bool) *GetSelfInfoService {
	s.compact = &val
	return s
}

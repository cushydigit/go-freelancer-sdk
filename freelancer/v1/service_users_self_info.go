package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type selfInfoService struct {
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

func (s *selfInfoService) Do(ctx context.Context) (*BaseResponse, error) {
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
	if s.pastCoverImages != nil {
		r.setParam("past_cover_images", *s.pastCoverImages)
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
	res := &BaseResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil

}

// SetAvatar sets whether to return the avatar of the selected user/users.
func (s *selfInfoService) SetAvatar(val bool) *selfInfoService {
	s.avatar = &val
	return s
}

// SetCountryDetails sets whether to return the country flag/code of the selected user/users.
func (s *selfInfoService) SetCountryDetails(val bool) *selfInfoService {
	s.countryDetails = &val
	return s
}

// SetProfileDescription sets whether to return the profile blurb of the selected user/users.
func (s *selfInfoService) SetProfileDescription(val bool) *selfInfoService {
	s.profileDescription = &val
	return s
}

// SetDisplayInfo sets whether to return the display name of the selected user/users.
func (s *selfInfoService) SetDisplayInfo(val bool) *selfInfoService {
	s.displayInfo = &val
	return s
}

// SetJobs sets whether to return the jobs of the selected user/users.
func (s *selfInfoService) SetJobs(val bool) *selfInfoService {
	s.jobs = &val
	return s
}

// SetBalanceDetails sets whether to return the currency balance of the selected user/users.
func (s *selfInfoService) SetBalanceDetails(val bool) *selfInfoService {
	s.balanceDetails = &val
	return s
}

// SetQualificationDetails sets whether to return qualification exams completed by the user/users.
func (s *selfInfoService) SetQualificationDetails(val bool) *selfInfoService {
	s.qualificationDetails = &val
	return s
}

// SetMembershipDetails sets whether to return the membership information of the user/users.
func (s *selfInfoService) SetMembershipDetails(val bool) *selfInfoService {
	s.membershipDetails = &val
	return s
}

// SetFinancialDetails sets whether to return the financial information of the user/users.
func (s *selfInfoService) SetFinancialDetails(val bool) *selfInfoService {
	s.financialDetails = &val
	return s
}

// SetLocationDetails sets whether to return the location information of the user/users.
func (s *selfInfoService) SetLocationDetails(val bool) *selfInfoService {
	s.locationDetails = &val
	return s
}

// SetPortfolioDetails sets whether to return the portfolio information of the user/users.
func (s *selfInfoService) SetPortfolioDetails(val bool) *selfInfoService {
	s.portfolioDetails = &val
	return s
}

// SetPreferredDetails sets whether to return the preferred information of the user/users.
func (s *selfInfoService) SetPreferredDetails(val bool) *selfInfoService {
	s.preferredDetails = &val
	return s
}

// SetBadgeDetails sets whether to return the badges earned by the user/users.
func (s *selfInfoService) SetBadgeDetails(val bool) *selfInfoService {
	s.badgeDetails = &val
	return s
}

// SetStatus sets whether to return additional status information about the user/users.
func (s *selfInfoService) SetStatus(val bool) *selfInfoService {
	s.status = &val
	return s
}

// SetReputation sets whether to return the freelancer reputation of the selected user/users.
func (s *selfInfoService) SetReputation(val bool) *selfInfoService {
	s.reputation = &val
	return s
}

// SetEmployerReputation sets whether to return the employer reputation of the selected user/users.
func (s *selfInfoService) SetEmployerReputation(val bool) *selfInfoService {
	s.employerReputation = &val
	return s
}

// SetReputationExtra sets whether to return the full freelancer reputation of the selected user/users.
func (s *selfInfoService) SetReputationExtra(val bool) *selfInfoService {
	s.reputationExtra = &val
	return s
}

// SetEmployerReputationExtra sets whether to return the full employer reputation of the selected user/users.
func (s *selfInfoService) SetEmployerReputationExtra(val bool) *selfInfoService {
	s.employerReputationExtra = &val
	return s
}

// SetCoverImage sets whether to return the profile picture of the user.
func (s *selfInfoService) SetCoverImage(val bool) *selfInfoService {
	s.coverImage = &val
	return s
}

// SetPastCoverImages sets whether to return previous profile pictures of the user.
func (s *selfInfoService) SetPastCoverImages(val bool) *selfInfoService {
	s.pastCoverImages = &val
	return s
}

// SetMobileTracking sets whether to return the mobile platforms used by the selected user/users.
func (s *selfInfoService) SetMobileTracking(val bool) *selfInfoService {
	s.mobileTracking = &val
	return s
}

// SetBidQualityDetails sets whether to return the user’s bid quality details, including bid quality score.
func (s *selfInfoService) SetBidQualityDetails(val bool) *selfInfoService {
	s.bidQualityDetails = &val
	return s
}

// SetDepositMethods sets whether to return the list of methods that the user can use to deposit funds, including any fees.
func (s *selfInfoService) SetDepositMethods(val bool) *selfInfoService {
	s.depositMethods = &val
	return s
}

// SetUserRecommendations sets whether to return recommendations count of selected user/users.
func (s *selfInfoService) SetUserRecommendations(val bool) *selfInfoService {
	s.userRecommendations = &val
	return s
}

// SetMarketingMobileNumber sets whether to return the mobile number of the user used by recruiters to contact the user.
func (s *selfInfoService) SetMarketingMobileNumber(val bool) *selfInfoService {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails sets whether to return the end timestamp of the sanction given to the user.
func (s *selfInfoService) SetSanctionDetails(val bool) *selfInfoService {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount sets whether to return the limit account status of the user.
func (s *selfInfoService) SetLimitedAccount(val bool) *selfInfoService {
	s.limitedAccount = &val
	return s
}

// SetCompletedUserRelevantJobCount sets whether to return the user’s number of completed relevant jobs.
func (s *selfInfoService) SetCompletedUserRelevantJobCount(val bool) *selfInfoService {
	s.completedUserRelevantJobCount = &val
	return s
}

// SetEquipmentGroupDetails sets whether to return the equipment groups and items attached to the user.
func (s *selfInfoService) SetEquipmentGroupDetails(val bool) *selfInfoService {
	s.equipmentGroupDetails = &val
	return s
}

// SetJobRanks sets whether to return the user’s job ranks (requires anyreputation.* and jobs projections to be true).
func (s *selfInfoService) SetJobRanks(val bool) *selfInfoService {
	s.jobRanks = &val
	return s
}

// SetShareholderDetails sets whether to return if the user is part of the Foundation Shareholder Program.
func (s *selfInfoService) SetShareholderDetails(val bool) *selfInfoService {
	s.shareholderDetails = &val
	return s
}

// SetRisingStar sets whether to return the user’s enrollment status in the Rising Star program.
func (s *selfInfoService) SetRisingStar(val bool) *selfInfoService {
	s.risingStar = &val
	return s
}

// SetStaffDetails sets whether to return if the user is a Freelancer staff.
func (s *selfInfoService) SetStaffDetails(val bool) *selfInfoService {
	s.staffDetails = &val
	return s
}

// SetJobSeoDetails sets whether to return the SEO details of the user’s jobs (requires jobs projection to be true).
func (s *selfInfoService) SetJobSeoDetails(val bool) *selfInfoService {
	s.jobSeoDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
func (s *selfInfoService) SetLimit(val int) *selfInfoService {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip, allowing pagination.
func (s *selfInfoService) SetOffset(val int) *selfInfoService {
	s.offset = &val
	return s
}

// SetCompact sets whether to strip all null and empty values from the response.
func (s *selfInfoService) SetCompact(val bool) *selfInfoService {
	s.compact = &val
	return s
}

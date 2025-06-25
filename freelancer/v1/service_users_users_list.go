package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type usersListService struct {
	client                        *Client
	users                         []int64
	usernames                     []string
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
}

func (s *usersListService) Do(ctx context.Context) (*BaseResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_USERS),
	}
	for _, userID := range s.users {
		r.addParam("users[]", userID)
	}
	for _, username := range s.usernames {
		r.addParam("usernames[]", username)
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

	res := &BaseResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SetUsers sets the list of user IDs to search.
func (s *usersListService) SetUsers(vals []int64) *usersListService {
	s.users = vals
	return s
}

// SetUsernames sets the list of usernames to search.
func (s *usersListService) SetUsernames(vals []string) *usersListService {
	s.usernames = vals
	return s
}

// SetAvatar enables returning the avatar of the selected user/users.
func (s *usersListService) SetAvatar(val bool) *usersListService {
	s.avatar = &val
	return s
}

// SetCountryDetails enables returning the country flag/code of the selected user/users.
func (s *usersListService) SetCountryDetails(val bool) *usersListService {
	s.countryDetails = &val
	return s
}

// SetProfileDescription enables returning the profile blurb of the selected user/users.
func (s *usersListService) SetProfileDescription(val bool) *usersListService {
	s.profileDescription = &val
	return s
}

// SetDisplayInfo enables returning the display name of the selected user/users.
func (s *usersListService) SetDisplayInfo(val bool) *usersListService {
	s.displayInfo = &val
	return s
}

// SetJobs enables returning the jobs of the selected user/users.
func (s *usersListService) SetJobs(val bool) *usersListService {
	s.jobs = &val
	return s
}

// SetBalanceDetails enables returning the currency balance of the selected user/users.
func (s *usersListService) SetBalanceDetails(val bool) *usersListService {
	s.balanceDetails = &val
	return s
}

// SetQualificationDetails enables returning qualification exams completed by the user/users.
func (s *usersListService) SetQualificationDetails(val bool) *usersListService {
	s.qualificationDetails = &val
	return s
}

// SetMembershipDetails enables returning the membership information of the user/users.
func (s *usersListService) SetMembershipDetails(val bool) *usersListService {
	s.membershipDetails = &val
	return s
}

// SetFinancialDetails enables returning the financial information of the user/users.
func (s *usersListService) SetFinancialDetails(val bool) *usersListService {
	s.financialDetails = &val
	return s
}

// SetLocationDetails enables returning the location information of the user/users.
func (s *usersListService) SetLocationDetails(val bool) *usersListService {
	s.locationDetails = &val
	return s
}

// SetPortfolioDetails enables returning the portfolio information of the user/users.
func (s *usersListService) SetPortfolioDetails(val bool) *usersListService {
	s.portfolioDetails = &val
	return s
}

// SetPreferredDetails enables returning the preferred information of the user/users.
func (s *usersListService) SetPreferredDetails(val bool) *usersListService {
	s.preferredDetails = &val
	return s
}

// SetBadgeDetails enables returning the badges earned by the user/users.
func (s *usersListService) SetBadgeDetails(val bool) *usersListService {
	s.badgeDetails = &val
	return s
}

// SetStatus enables returning additional status information about the user/users.
func (s *usersListService) SetStatus(val bool) *usersListService {
	s.status = &val
	return s
}

// SetReputation enables returning the freelancer reputation of the selected user/users.
func (s *usersListService) SetReputation(val bool) *usersListService {
	s.reputation = &val
	return s
}

// SetEmployerReputation enables returning the employer reputation of the selected user/users.
func (s *usersListService) SetEmployerReputation(val bool) *usersListService {
	s.employerReputation = &val
	return s
}

// SetReputationExtra enables returning the full freelancer reputation of the selected user/users.
func (s *usersListService) SetReputationExtra(val bool) *usersListService {
	s.reputationExtra = &val
	return s
}

// SetEmployerReputationExtra enables returning the full employer reputation of the selected user/users.
func (s *usersListService) SetEmployerReputationExtra(val bool) *usersListService {
	s.employerReputationExtra = &val
	return s
}

// SetCoverImage enables returning the profile picture of the user.
func (s *usersListService) SetCoverImage(val bool) *usersListService {
	s.coverImage = &val
	return s
}

// SetPastCoverImage enables returning previous profile pictures of the user.
func (s *usersListService) SetPastCoverImage(val bool) *usersListService {
	s.pastCoverImage = &val
	return s
}

// SetMobileTracking enables returning the mobile platforms used by the selected user/users.
func (s *usersListService) SetMobileTracking(val bool) *usersListService {
	s.mobileTracking = &val
	return s
}

// SetBidQualityDetails enables returning the user's bid quality details, including bid quality score.
func (s *usersListService) SetBidQualityDetails(val bool) *usersListService {
	s.bidQualityDetails = &val
	return s
}

// SetDepositMethods enables returning the list of deposit methods and associated fees.
func (s *usersListService) SetDepositMethods(val bool) *usersListService {
	s.depositMethods = &val
	return s
}

// SetUserRecommendations enables returning recommendations count of selected user/users.
func (s *usersListService) SetUserRecommendations(val bool) *usersListService {
	s.userRecommendations = &val
	return s
}

// SetMarketingMobileNumber enables returning the mobile number used by recruiters to contact the user.
func (s *usersListService) SetMarketingMobileNumber(val bool) *usersListService {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables returning the end timestamp of any sanction given to the user.
func (s *usersListService) SetSanctionDetails(val bool) *usersListService {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables returning the limited account status of the user.
func (s *usersListService) SetLimitedAccount(val bool) *usersListService {
	s.limitedAccount = &val
	return s
}

// SetCompletedUserRelevantJobCount enables returning the number of completed relevant jobs.
func (s *usersListService) SetCompletedUserRelevantJobCount(val bool) *usersListService {
	s.completedUserRelevantJobCount = &val
	return s
}

// SetEquipmentGroupDetails enables returning the user's equipment groups and items.
func (s *usersListService) SetEquipmentGroupDetails(val bool) *usersListService {
	s.equipmentGroupDetails = &val
	return s
}

// SetJobRanks enables returning the user's job ranks (requires reputation and jobs projections).
func (s *usersListService) SetJobRanks(val bool) *usersListService {
	s.jobRanks = &val
	return s
}

// SetShareholderDetails enables returning whether the user is part of the Shareholder Program.
func (s *usersListService) SetShareholderDetails(val bool) *usersListService {
	s.shareholderDetails = &val
	return s
}

// SetRisingStar enables returning the user's Rising Star program status.
func (s *usersListService) SetRisingStar(val bool) *usersListService {
	s.risingStar = &val
	return s
}

// SetStaffDetails enables returning whether the user is a Freelancer staff member.
func (s *usersListService) SetStaffDetails(val bool) *usersListService {
	s.staffDetails = &val
	return s
}

// SetJobSeoDetails enables returning SEO details of the user’s jobs (requires jobs projection).
func (s *usersListService) SetJobSeoDetails(val bool) *usersListService {
	s.jobSeoDetails = &val
	return s
}

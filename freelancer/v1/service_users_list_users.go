package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

// Returns a list of users
type ListUsersService struct {
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

type ListUsersResponse struct {
	Status    string      `json:"status"`
	RequestID string      `json:"request_id"`
	Result    UsersResult `json:"result"`
}

type UsersResult struct {
	Users map[string]User `json:"users"`
}

func (s *ListUsersService) Do(ctx context.Context) (*ListUsersResponse, error) {
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

	res := &ListUsersResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ListUsersService) SetUsers(users []int64) *ListUsersService {
	s.users = users
	return s
}

func (s *ListUsersService) SetUsernames(usernames []string) *ListUsersService {
	s.usernames = usernames
	return s
}

func (s *ListUsersService) SetAvatar(val bool) *ListUsersService {
	s.avatar = &val
	return s
}

func (s *ListUsersService) SetCountryDetails(val bool) *ListUsersService {
	s.countryDetails = &val
	return s
}

func (s *ListUsersService) SetProfileDescription(val bool) *ListUsersService {
	s.profileDescription = &val
	return s
}

func (s *ListUsersService) SetDisplayInfo(val bool) *ListUsersService {
	s.displayInfo = &val
	return s
}

func (s *ListUsersService) SetJobs(val bool) *ListUsersService {
	s.jobs = &val
	return s
}

func (s *ListUsersService) SetBalanceDetails(val bool) *ListUsersService {
	s.balanceDetails = &val
	return s
}

func (s *ListUsersService) SetQualificationDetails(val bool) *ListUsersService {
	s.qualificationDetails = &val
	return s
}

func (s *ListUsersService) SetMembershipDetails(val bool) *ListUsersService {
	s.membershipDetails = &val
	return s
}

func (s *ListUsersService) SetFinancialDetails(val bool) *ListUsersService {
	s.financialDetails = &val
	return s
}

func (s *ListUsersService) SetLocationDetails(val bool) *ListUsersService {
	s.locationDetails = &val
	return s
}

func (s *ListUsersService) SetPortfolioDetails(val bool) *ListUsersService {
	s.portfolioDetails = &val
	return s
}

func (s *ListUsersService) SetPreferredDetails(val bool) *ListUsersService {
	s.preferredDetails = &val
	return s
}

func (s *ListUsersService) SetBadgeDetails(val bool) *ListUsersService {
	s.badgeDetails = &val
	return s
}

func (s *ListUsersService) SetStatus(val bool) *ListUsersService {
	s.status = &val
	return s
}

func (s *ListUsersService) SetReputation(val bool) *ListUsersService {
	s.reputation = &val
	return s
}

func (s *ListUsersService) SetEmployerReputation(val bool) *ListUsersService {
	s.employerReputation = &val
	return s
}

func (s *ListUsersService) SetReputationExtra(val bool) *ListUsersService {
	s.reputationExtra = &val
	return s
}

func (s *ListUsersService) SetEmployerReputationExtra(val bool) *ListUsersService {
	s.employerReputationExtra = &val
	return s
}

func (s *ListUsersService) SetCoverImage(val bool) *ListUsersService {
	s.coverImage = &val
	return s
}

func (s *ListUsersService) SetPastCoverImage(val bool) *ListUsersService {
	s.pastCoverImage = &val
	return s
}

func (s *ListUsersService) SetMobileTracking(val bool) *ListUsersService {
	s.mobileTracking = &val
	return s
}

func (s *ListUsersService) SetBidQualityDetails(val bool) *ListUsersService {
	s.bidQualityDetails = &val
	return s
}

func (s *ListUsersService) SetDepositMethods(val bool) *ListUsersService {
	s.depositMethods = &val
	return s
}

func (s *ListUsersService) SetUserRecommendations(val bool) *ListUsersService {
	s.userRecommendations = &val
	return s
}

func (s *ListUsersService) SetMarketingMobileNumber(val bool) *ListUsersService {
	s.marketingMobileNumber = &val
	return s
}

func (s *ListUsersService) SetSanctionDetails(val bool) *ListUsersService {
	s.sanctionDetails = &val
	return s
}

func (s *ListUsersService) SetLimitedAccount(val bool) *ListUsersService {
	s.limitedAccount = &val
	return s
}

func (s *ListUsersService) SetCompletedUserRelevantJobCount(val bool) *ListUsersService {
	s.completedUserRelevantJobCount = &val
	return s
}

func (s *ListUsersService) SetEquipmentGroupDetails(val bool) *ListUsersService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *ListUsersService) SetJobRanks(val bool) *ListUsersService {
	s.jobRanks = &val
	return s
}

func (s *ListUsersService) SetShareholderDetails(val bool) *ListUsersService {
	s.shareholderDetails = &val
	return s
}

func (s *ListUsersService) SetRisingStar(val bool) *ListUsersService {
	s.risingStar = &val
	return s
}

func (s *ListUsersService) SetStaffDetails(val bool) *ListUsersService {
	s.staffDetails = &val
	return s
}

func (s *ListUsersService) SetJobSeoDetails(val bool) *ListUsersService {
	s.jobSeoDetails = &val
	return s
}

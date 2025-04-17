package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUsersService struct {
	client                        *Client
	users                         []int
	usernames                     []string
	avatar                        bool
	countryDetails                bool
	profileDescription            bool
	displayInfo                   bool
	jobs                          bool
	balanceDetails                bool
	qualificationDetails          bool
	membershipDetails             bool
	financialDetails              bool
	locationDetails               bool
	portfolioDetails              bool
	preferredDetails              bool
	badgeDetails                  bool
	status                        bool
	reputation                    bool
	employerReputation            bool
	reputationExtra               bool
	employerReputationExtra       bool
	coverImage                    bool
	pastCoverImage                bool
	mobileTracking                bool
	bidQualityDetails             bool
	depositMethods                bool
	userRecommendations           bool
	marketingMobileNumber         bool
	sanctionDetails               bool
	limitedAccount                bool
	completedUserRelevantJobCount bool
	equipmentGroupDetails         bool
	jobRanks                      bool
	jobSeoDetails                 bool
	risingStar                    bool
	shareholderDetails            bool
	staffDetails                  bool
}

func (s *ListUsersService) Do(ctx context.Context) (*ListUsersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/users/0.1/users",
	}
	for _, userID := range s.users {
		r.addParam("users[]", userID)
	}
	for _, username := range s.usernames {
		r.addParam("usernames[]", username)
	}
	if s.avatar {
		r.setParam("avatar", s.avatar)
	}
	if s.countryDetails {
		r.setParam("country_details", s.countryDetails)
	}
	if s.profileDescription {
		r.setParam("profile_description", s.profileDescription)
	}
	if s.displayInfo {
		r.setParam("display_info", s.displayInfo)
	}
	if s.jobs {
		r.setParam("jobs", s.jobs)
	}
	if s.balanceDetails {
		r.setParam("balance_details", s.balanceDetails)
	}
	if s.qualificationDetails {
		r.setParam("qualification_details", s.qualificationDetails)
	}
	if s.membershipDetails {
		r.setParam("membership_details", s.membershipDetails)
	}
	if s.financialDetails {
		r.setParam("financial_details", s.financialDetails)
	}
	if s.locationDetails {
		r.setParam("location_details", s.locationDetails)
	}
	if s.portfolioDetails {
		r.setParam("portfolio_details", s.portfolioDetails)
	}
	if s.preferredDetails {
		r.setParam("preferred_details", s.preferredDetails)
	}
	if s.badgeDetails {
		r.setParam("badge_details", s.badgeDetails)
	}
	if s.status {
		r.setParam("status", s.status)
	}
	if s.reputation {
		r.setParam("reputation", s.reputation)
	}
	if s.employerReputation {
		r.setParam("employer_reputation", s.employerReputation)
	}
	if s.reputationExtra {
		r.setParam("reputation_extra", s.reputationExtra)
	}
	if s.employerReputationExtra {
		r.setParam("employer_reputation_extra", s.employerReputationExtra)
	}
	if s.coverImage {
		r.setParam("cover_image", s.coverImage)
	}
	if s.pastCoverImage {
		r.setParam("past_cover_image", s.pastCoverImage)
	}
	if s.mobileTracking {
		r.setParam("mobile_tracking", s.mobileTracking)
	}
	if s.bidQualityDetails {
		r.setParam("bid_quality_details", s.bidQualityDetails)
	}
	if s.depositMethods {
		r.setParam("deposit_methods", s.depositMethods)
	}
	if s.userRecommendations {
		r.setParam("user_recommendations", s.userRecommendations)
	}
	if s.marketingMobileNumber {
		r.setParam("marketing_mobile_number", s.marketingMobileNumber)
	}
	if s.sanctionDetails {
		r.setParam("sanction_details", s.sanctionDetails)
	}
	if s.limitedAccount {
		r.setParam("limited_account", s.limitedAccount)
	}
	if s.completedUserRelevantJobCount {
		r.setParam("completed_user_relevant_job_count", s.completedUserRelevantJobCount)
	}
	if s.equipmentGroupDetails {
		r.setParam("equipment_group_details", s.equipmentGroupDetails)
	}
	if s.jobRanks {
		r.setParam("job_ranks", s.jobRanks)
	}
	if s.jobSeoDetails {
		r.setParam("job_seo_details", s.jobSeoDetails)
	}
	if s.risingStar {
		r.setParam("rising_star", s.risingStar)
	}
	if s.shareholderDetails {
		r.setParam("shareholder_details", s.shareholderDetails)
	}
	if s.staffDetails {
		r.setParam("staff_details", s.staffDetails)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListUsersResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListUsersService) SetUsers(users []int) *ListUsersService {
	s.users = users
	return s
}

func (s *ListUsersService) SetUsernames(usernames []string) *ListUsersService {
	s.usernames = usernames
	return s
}

func (s *ListUsersService) SetAvatar(avatar bool) *ListUsersService {
	s.avatar = avatar
	return s
}

func (s *ListUsersService) SetCountryDetails(countryDetails bool) *ListUsersService {
	s.countryDetails = countryDetails
	return s
}

func (s *ListUsersService) SetProfileDescription(profileDescription bool) *ListUsersService {
	s.profileDescription = profileDescription
	return s
}

func (s *ListUsersService) SetDisplayInfo(displayInfo bool) *ListUsersService {
	s.displayInfo = displayInfo
	return s
}

func (s *ListUsersService) SetJobs(jobs bool) *ListUsersService {
	s.jobs = jobs
	return s
}

func (s *ListUsersService) SetBalanceDetails(balanceDetails bool) *ListUsersService {
	s.balanceDetails = balanceDetails
	return s
}

func (s *ListUsersService) SetQualificationDetails(qualificationDetails bool) *ListUsersService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *ListUsersService) SetMembershipDetails(membershipDetails bool) *ListUsersService {
	s.membershipDetails = membershipDetails
	return s
}

func (s *ListUsersService) SetFinancialDetails(financialDetails bool) *ListUsersService {
	s.financialDetails = financialDetails
	return s
}

func (s *ListUsersService) SetLocationDetails(locationDetails bool) *ListUsersService {
	s.locationDetails = locationDetails
	return s
}

func (s *ListUsersService) SetPortfolioDetails(portfolioDetails bool) *ListUsersService {
	s.portfolioDetails = portfolioDetails
	return s
}

func (s *ListUsersService) SetPreferredDetails(preferredDetails bool) *ListUsersService {
	s.preferredDetails = preferredDetails
	return s
}

func (s *ListUsersService) SetBadgeDetails(badgeDetails bool) *ListUsersService {
	s.badgeDetails = badgeDetails
	return s
}

func (s *ListUsersService) SetStatus(status bool) *ListUsersService {
	s.status = status
	return s
}

func (s *ListUsersService) SetReputation(reputation bool) *ListUsersService {
	s.reputation = reputation
	return s
}

func (s *ListUsersService) SetEmployerReputation(employerReputation bool) *ListUsersService {
	s.employerReputation = employerReputation
	return s
}

func (s *ListUsersService) SetReputationExtra(reputationExtra bool) *ListUsersService {
	s.reputationExtra = reputationExtra
	return s
}

func (s *ListUsersService) SetEmployerReputationExtra(employerReputationExtra bool) *ListUsersService {
	s.employerReputationExtra = employerReputationExtra
	return s
}

func (s *ListUsersService) SetCoverImage(coverImage bool) *ListUsersService {
	s.coverImage = coverImage
	return s
}

func (s *ListUsersService) SetPastCoverImage(pastCoverImage bool) *ListUsersService {
	s.pastCoverImage = pastCoverImage
	return s
}

func (s *ListUsersService) SetMobileTracking(mobileTracking bool) *ListUsersService {
	s.mobileTracking = mobileTracking
	return s
}

func (s *ListUsersService) SetBidQualityDetails(bidQualityDetails bool) *ListUsersService {
	s.bidQualityDetails = bidQualityDetails
	return s
}

func (s *ListUsersService) SetDepositMethods(depositMethods bool) *ListUsersService {
	s.depositMethods = depositMethods
	return s
}

func (s *ListUsersService) SetUserRecommendations(userRecommendations bool) *ListUsersService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ListUsersService) SetMarketingMobileNumber(marketingMobileNumber bool) *ListUsersService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ListUsersService) SetSanctionDetails(sanctionDetails bool) *ListUsersService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ListUsersService) SetLimitedAccount(limitedAccount bool) *ListUsersService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ListUsersService) SetCompletedUserRelevantJobCount(completedUserRelevantJobCount bool) *ListUsersService {
	s.completedUserRelevantJobCount = completedUserRelevantJobCount
	return s
}

func (s *ListUsersService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ListUsersService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *ListUsersService) SetJobRanks(jobRanks bool) *ListUsersService {
	s.jobRanks = jobRanks
	return s
}

func (s *ListUsersService) SetShareholderDetails(shareholderDetails bool) *ListUsersService {
	s.shareholderDetails = shareholderDetails
	return s
}

func (s *ListUsersService) SetRisingStar(risingStar bool) *ListUsersService {
	s.risingStar = risingStar
	return s
}

func (s *ListUsersService) SetStaffDetails(staffDetails bool) *ListUsersService {
	s.staffDetails = staffDetails
	return s
}

func (s *ListUsersService) SetJobSeoDetails(jobSeoDetails bool) *ListUsersService {
	s.jobSeoDetails = jobSeoDetails
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListSelfDevicesService struct {
	client *Client
}

type ListSelfDevicesResponse struct {
	Status    string        `json:"status"`
	RequestID string        `json:"requset_id"`
	Result    DevicesResult `json:"result"`
}

type DevicesResult struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	UserAgent string `json:"user_agent"`
	Platform  string `json:"platform"`
	City      string `json:"city"`
	Country   string `json:"country"`
	LastLogin int64  `json:"last_login"`
}

func (s *ListSelfDevicesService) Do(ctx context.Context) (*ListSelfDevicesResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF_DEVICES),
	}

	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListSelfDevicesResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type GetSelfInfoService struct {
	client                        *Client
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
	pastCoverImages               bool
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
	limit                         int
	offset                        int
	compact                       bool
}

type GetSelfInfoResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

func (s *GetSelfInfoService) Do(ctx context.Context) (*GetSelfInfoResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_SELF),
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
	res := &GetSelfInfoResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil

}

func (s *GetSelfInfoService) SetAvatar(avatar bool) *GetSelfInfoService {
	s.avatar = avatar
	return s
}

func (s *GetSelfInfoService) SetCountryDetails(countryDetails bool) *GetSelfInfoService {
	s.countryDetails = countryDetails
	return s
}

func (s *GetSelfInfoService) SetProfileDescription(profileDescription bool) *GetSelfInfoService {
	s.profileDescription = profileDescription
	return s
}

func (s *GetSelfInfoService) SetDisplayInfo(displayInfo bool) *GetSelfInfoService {
	s.displayInfo = displayInfo
	return s
}

func (s *GetSelfInfoService) SetJobs(jobs bool) *GetSelfInfoService {
	s.jobs = jobs
	return s
}

func (s *GetSelfInfoService) SetBalanceDetails(balanceDetails bool) *GetSelfInfoService {
	s.balanceDetails = balanceDetails
	return s
}

func (s *GetSelfInfoService) SetQualificationDetails(qualificationDetails bool) *GetSelfInfoService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *GetSelfInfoService) SetMembershipDetails(membershipDetails bool) *GetSelfInfoService {
	s.membershipDetails = membershipDetails
	return s
}

func (s *GetSelfInfoService) SetFinancialDetails(financialDetails bool) *GetSelfInfoService {
	s.financialDetails = financialDetails
	return s
}

func (s *GetSelfInfoService) SetLocationDetails(locationDetails bool) *GetSelfInfoService {
	s.locationDetails = locationDetails
	return s
}

func (s *GetSelfInfoService) SetPortfolioDetails(portfolioDetails bool) *GetSelfInfoService {
	s.portfolioDetails = portfolioDetails
	return s
}

func (s *GetSelfInfoService) SetPreferredDetails(preferredDetails bool) *GetSelfInfoService {
	s.preferredDetails = preferredDetails
	return s
}

func (s *GetSelfInfoService) SetBadgeDetails(badgeDetails bool) *GetSelfInfoService {
	s.badgeDetails = badgeDetails
	return s
}

func (s *GetSelfInfoService) SetStatus(status bool) *GetSelfInfoService {
	s.status = status
	return s
}

func (s *GetSelfInfoService) SetReputation(reputation bool) *GetSelfInfoService {
	s.reputation = reputation
	return s
}

func (s *GetSelfInfoService) SetEmployerReputation(employerReputation bool) *GetSelfInfoService {
	s.employerReputation = employerReputation
	return s
}

func (s *GetSelfInfoService) SetReputationExtra(reputationExtra bool) *GetSelfInfoService {
	s.reputationExtra = reputationExtra
	return s
}

func (s *GetSelfInfoService) SetEmployerReputationExtra(employerReputationExtra bool) *GetSelfInfoService {
	s.employerReputationExtra = employerReputationExtra
	return s
}

func (s *GetSelfInfoService) SetCoverImage(coverImage bool) *GetSelfInfoService {
	s.coverImage = coverImage
	return s
}

func (s *GetSelfInfoService) SetPastCoverImage(pastCoverImage bool) *GetSelfInfoService {
	s.pastCoverImage = pastCoverImage
	return s
}

func (s *GetSelfInfoService) SetMobileTracking(mobileTracking bool) *GetSelfInfoService {
	s.mobileTracking = mobileTracking
	return s
}

func (s *GetSelfInfoService) SetBidQualityDetails(bidQualityDetails bool) *GetSelfInfoService {
	s.bidQualityDetails = bidQualityDetails
	return s
}

func (s *GetSelfInfoService) SetDepositMethods(depositMethods bool) *GetSelfInfoService {
	s.depositMethods = depositMethods
	return s
}

func (s *GetSelfInfoService) SetUserRecommendations(userRecommendations bool) *GetSelfInfoService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *GetSelfInfoService) SetMarketingMobileNumber(marketingMobileNumber bool) *GetSelfInfoService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *GetSelfInfoService) SetSanctionDetails(sanctionDetails bool) *GetSelfInfoService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *GetSelfInfoService) SetLimitedAccount(limitedAccount bool) *GetSelfInfoService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *GetSelfInfoService) SetCompletedUserRelevantJobCount(completedUserRelevantJobCount bool) *GetSelfInfoService {
	s.completedUserRelevantJobCount = completedUserRelevantJobCount
	return s
}

func (s *GetSelfInfoService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *GetSelfInfoService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *GetSelfInfoService) SetJobRanks(jobRanks bool) *GetSelfInfoService {
	s.jobRanks = jobRanks
	return s
}

func (s *GetSelfInfoService) SetShareholderDetails(shareholderDetails bool) *GetSelfInfoService {
	s.shareholderDetails = shareholderDetails
	return s
}

func (s *GetSelfInfoService) SetRisingStar(risingStar bool) *GetSelfInfoService {
	s.risingStar = risingStar
	return s
}

func (s *GetSelfInfoService) SetStaffDetails(staffDetails bool) *GetSelfInfoService {
	s.staffDetails = staffDetails
	return s
}

func (s *GetSelfInfoService) SetJobSeoDetails(jobSeoDetails bool) *GetSelfInfoService {
	s.jobSeoDetails = jobSeoDetails
	return s
}

func (s *GetSelfInfoService) SetLimit(limit int) *GetSelfInfoService {
	s.limit = limit
	return s
}

func (s *GetSelfInfoService) SetOffset(offset int) *GetSelfInfoService {
	s.offset = offset
	return s
}

func (s *GetSelfInfoService) SetCompact(compact bool) *GetSelfInfoService {
	s.compact = compact
	return s
}

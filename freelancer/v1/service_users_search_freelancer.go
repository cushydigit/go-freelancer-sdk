package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type SearchFreelancersService struct {
	client                        *Client
	query                         *string
	jobsIDs                       []int32 // I have to rename the field (jobs[]) to jobsIDs we have another jobs as bool
	skills                        []int32
	countries                     []string
	hourlyRateMin                 *int
	hourlyRateMax                 *int
	reviewCountMin                *int
	reviewCountMax                *int
	onlineOnly                    *bool
	locationLatitude              *float64
	locationLongitude             *float64
	insignias                     []int32
	poolIds                       []int32
	ratings                       *float32
	sortField                     *int // DirectorySortFieldEnum? could not found in freelancer.com docs
	reverseSort                   *bool
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
	poolDetails                   *bool
	limit                         *int
	offset                        *int
	compact                       *bool
}

type SearchFreelancersResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id"`
	Result    FreelancerResult `json:"result"`
}

type FreelancerResult struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
}

func (s *SearchFreelancersService) Do(ctx context.Context) (*SearchFreelancersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_FREELANCERS),
	}

	if s.query != nil {
		r.addParam("query", *s.query)
	}
	for _, jobID := range s.jobsIDs {
		r.addParam("jobs[]", jobID)
	}
	for _, skill := range s.skills {
		r.addParam("skills[]", skill)
	}
	for _, country := range s.countries {
		r.addParam("countries[]", country)
	}
	if s.hourlyRateMin != nil {
		r.addParam("hourly_rate_min", *s.hourlyRateMin)
	}
	if s.hourlyRateMax != nil {
		r.addParam("hourly_rate_max", *s.hourlyRateMax)
	}
	if s.reviewCountMin != nil {
		r.addParam("review_count_min", *s.reviewCountMin)
	}
	if s.reviewCountMax != nil {
		r.addParam("review_count_max", *s.reviewCountMax)
	}
	if s.onlineOnly != nil {
		r.addParam("online_only", *s.onlineOnly)
	}
	if s.locationLatitude != nil {
		r.addParam("location_latitude", *s.locationLatitude)
	}
	if s.locationLongitude != nil {
		r.addParam("location_longitude", *s.locationLongitude)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
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
	if s.poolDetails != nil {
		r.setParam("pool_details", *s.poolDetails)
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
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &SearchFreelancersResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SearchFreelancersService) SetQuery(val string) *SearchFreelancersService {
	s.query = &val
	return s
}

func (s *SearchFreelancersService) SetJobsIDs(jobsIDs []int32) *SearchFreelancersService {
	s.jobsIDs = jobsIDs
	return s
}

func (s *SearchFreelancersService) SetSkills(skills []int32) *SearchFreelancersService {
	s.skills = skills
	return s
}

func (s *SearchFreelancersService) SetCountries(countries []string) *SearchFreelancersService {
	s.countries = countries
	return s
}

func (s *SearchFreelancersService) SetHourlyRateMin(val int) *SearchFreelancersService {
	s.hourlyRateMin = &val
	return s
}

func (s *SearchFreelancersService) SetHourlyRateMax(val int) *SearchFreelancersService {
	s.hourlyRateMax = &val
	return s
}

func (s *SearchFreelancersService) SetReviewCountMin(val int) *SearchFreelancersService {
	s.reviewCountMin = &val
	return s
}

func (s *SearchFreelancersService) SetReviewCountMax(val int) *SearchFreelancersService {
	s.reviewCountMax = &val
	return s
}

func (s *SearchFreelancersService) SetOnlineOnly(val bool) *SearchFreelancersService {
	s.onlineOnly = &val
	return s
}

func (s *SearchFreelancersService) SetLocationLatitude(locationLatitude float64) *SearchFreelancersService {
	s.locationLatitude = &locationLatitude
	return s
}

func (s *SearchFreelancersService) SetLocationLongitude(locationLongitude float64) *SearchFreelancersService {
	s.locationLongitude = &locationLongitude
	return s
}

func (s *SearchFreelancersService) SetInsignias(insignias []int32) *SearchFreelancersService {
	s.insignias = insignias
	return s
}

func (s *SearchFreelancersService) SetPoolIds(poolIds []int32) *SearchFreelancersService {
	s.poolIds = poolIds
	return s
}

func (s *SearchFreelancersService) SetRatings(val float32) *SearchFreelancersService {
	s.ratings = &val
	return s
}

func (s *SearchFreelancersService) SetSortField(val int) *SearchFreelancersService {
	s.sortField = &val
	return s
}

func (s *SearchFreelancersService) SetReverseSort(val bool) *SearchFreelancersService {
	s.reverseSort = &val
	return s
}

func (s *SearchFreelancersService) SetAvatar(val bool) *SearchFreelancersService {
	s.avatar = &val
	return s
}

func (s *SearchFreelancersService) SetCountryDetails(val bool) *SearchFreelancersService {
	s.countryDetails = &val
	return s
}

func (s *SearchFreelancersService) SetProfileDescription(val bool) *SearchFreelancersService {
	s.profileDescription = &val
	return s
}

func (s *SearchFreelancersService) SetDisplayInfo(val bool) *SearchFreelancersService {
	s.displayInfo = &val
	return s
}

func (s *SearchFreelancersService) SetJobs(val bool) *SearchFreelancersService {
	s.jobs = &val
	return s
}

func (s *SearchFreelancersService) SetBalanceDetails(val bool) *SearchFreelancersService {
	s.balanceDetails = &val
	return s
}

func (s *SearchFreelancersService) SetQualificationDetails(val bool) *SearchFreelancersService {
	s.qualificationDetails = &val
	return s
}

func (s *SearchFreelancersService) SetMembershipDetails(val bool) *SearchFreelancersService {
	s.membershipDetails = &val
	return s
}

func (s *SearchFreelancersService) SetFinancialDetails(val bool) *SearchFreelancersService {
	s.financialDetails = &val
	return s
}

func (s *SearchFreelancersService) SetLocationDetails(val bool) *SearchFreelancersService {
	s.locationDetails = &val
	return s
}

func (s *SearchFreelancersService) SetPortfolioDetails(val bool) *SearchFreelancersService {
	s.portfolioDetails = &val
	return s
}

func (s *SearchFreelancersService) SetPreferredDetails(val bool) *SearchFreelancersService {
	s.preferredDetails = &val
	return s
}

func (s *SearchFreelancersService) SetBadgeDetails(val bool) *SearchFreelancersService {
	s.badgeDetails = &val
	return s
}

func (s *SearchFreelancersService) SetStatus(val bool) *SearchFreelancersService {
	s.status = &val
	return s
}

func (s *SearchFreelancersService) SetReputation(val bool) *SearchFreelancersService {
	s.reputation = &val
	return s
}

func (s *SearchFreelancersService) SetEmployerReputation(val bool) *SearchFreelancersService {
	s.employerReputation = &val
	return s
}

func (s *SearchFreelancersService) SetReputationExtra(val bool) *SearchFreelancersService {
	s.reputationExtra = &val
	return s
}

func (s *SearchFreelancersService) SetEmployerReputationExtra(val bool) *SearchFreelancersService {
	s.employerReputationExtra = &val
	return s
}

func (s *SearchFreelancersService) SetCoverImage(val bool) *SearchFreelancersService {
	s.coverImage = &val
	return s
}

func (s *SearchFreelancersService) SetPastCoverImage(val bool) *SearchFreelancersService {
	s.pastCoverImage = &val
	return s
}

func (s *SearchFreelancersService) SetMobileTracking(val bool) *SearchFreelancersService {
	s.mobileTracking = &val
	return s
}

func (s *SearchFreelancersService) SetBidQualityDetails(val bool) *SearchFreelancersService {
	s.bidQualityDetails = &val
	return s
}

func (s *SearchFreelancersService) SetDepositMethods(val bool) *SearchFreelancersService {
	s.depositMethods = &val
	return s
}

func (s *SearchFreelancersService) SetUserRecommendations(val bool) *SearchFreelancersService {
	s.userRecommendations = &val
	return s
}

func (s *SearchFreelancersService) SetMarketingMobileNumber(val bool) *SearchFreelancersService {
	s.marketingMobileNumber = &val
	return s
}

func (s *SearchFreelancersService) SetSanctionDetails(val bool) *SearchFreelancersService {
	s.sanctionDetails = &val
	return s
}

func (s *SearchFreelancersService) SetLimitedAccount(val bool) *SearchFreelancersService {
	s.limitedAccount = &val
	return s
}

func (s *SearchFreelancersService) SetCompletedUserRelevantJobCount(val bool) *SearchFreelancersService {
	s.completedUserRelevantJobCount = &val
	return s
}

func (s *SearchFreelancersService) SetEquipmentGroupDetails(val bool) *SearchFreelancersService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *SearchFreelancersService) SetJobRanks(val bool) *SearchFreelancersService {
	s.jobRanks = &val
	return s
}

func (s *SearchFreelancersService) SetJobSeoDetails(val bool) *SearchFreelancersService {
	s.jobSeoDetails = &val
	return s
}

func (s *SearchFreelancersService) SetRisingStar(val bool) *SearchFreelancersService {
	s.risingStar = &val
	return s
}

func (s *SearchFreelancersService) SetShareholderDetails(val bool) *SearchFreelancersService {
	s.shareholderDetails = &val
	return s
}

func (s *SearchFreelancersService) SetStaffDetails(val bool) *SearchFreelancersService {
	s.staffDetails = &val
	return s
}

func (s *SearchFreelancersService) SetPoolDetails(val bool) *SearchFreelancersService {
	s.poolDetails = &val
	return s
}

func (s *SearchFreelancersService) SetLimit(val int) *SearchFreelancersService {
	s.limit = &val
	return s
}

func (s *SearchFreelancersService) SetOffset(val int) *SearchFreelancersService {
	s.offset = &val
	return s
}

func (s *SearchFreelancersService) SetCompact(val bool) *SearchFreelancersService {
	s.compact = &val
	return s
}

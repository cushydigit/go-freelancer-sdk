package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type freelancersSearchService struct {
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

func (s *freelancersSearchService) Do(ctx context.Context) (*BaseResponse, error) {
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
	res := &BaseResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *freelancersSearchService) SetQuery(val string) *freelancersSearchService {
	s.query = &val
	return s
}

func (s *freelancersSearchService) SetJobsIDs(jobsIDs []int32) *freelancersSearchService {
	s.jobsIDs = jobsIDs
	return s
}

func (s *freelancersSearchService) SetSkills(skills []int32) *freelancersSearchService {
	s.skills = skills
	return s
}

func (s *freelancersSearchService) SetCountries(countries []string) *freelancersSearchService {
	s.countries = countries
	return s
}

func (s *freelancersSearchService) SetHourlyRateMin(val int) *freelancersSearchService {
	s.hourlyRateMin = &val
	return s
}

func (s *freelancersSearchService) SetHourlyRateMax(val int) *freelancersSearchService {
	s.hourlyRateMax = &val
	return s
}

func (s *freelancersSearchService) SetReviewCountMin(val int) *freelancersSearchService {
	s.reviewCountMin = &val
	return s
}

func (s *freelancersSearchService) SetReviewCountMax(val int) *freelancersSearchService {
	s.reviewCountMax = &val
	return s
}

func (s *freelancersSearchService) SetOnlineOnly(val bool) *freelancersSearchService {
	s.onlineOnly = &val
	return s
}

func (s *freelancersSearchService) SetLocationLatitude(locationLatitude float64) *freelancersSearchService {
	s.locationLatitude = &locationLatitude
	return s
}

func (s *freelancersSearchService) SetLocationLongitude(locationLongitude float64) *freelancersSearchService {
	s.locationLongitude = &locationLongitude
	return s
}

func (s *freelancersSearchService) SetInsignias(insignias []int32) *freelancersSearchService {
	s.insignias = insignias
	return s
}

func (s *freelancersSearchService) SetPoolIds(poolIds []int32) *freelancersSearchService {
	s.poolIds = poolIds
	return s
}

func (s *freelancersSearchService) SetRatings(val float32) *freelancersSearchService {
	s.ratings = &val
	return s
}

func (s *freelancersSearchService) SetSortField(val int) *freelancersSearchService {
	s.sortField = &val
	return s
}

func (s *freelancersSearchService) SetReverseSort(val bool) *freelancersSearchService {
	s.reverseSort = &val
	return s
}

func (s *freelancersSearchService) SetAvatar(val bool) *freelancersSearchService {
	s.avatar = &val
	return s
}

func (s *freelancersSearchService) SetCountryDetails(val bool) *freelancersSearchService {
	s.countryDetails = &val
	return s
}

func (s *freelancersSearchService) SetProfileDescription(val bool) *freelancersSearchService {
	s.profileDescription = &val
	return s
}

func (s *freelancersSearchService) SetDisplayInfo(val bool) *freelancersSearchService {
	s.displayInfo = &val
	return s
}

func (s *freelancersSearchService) SetJobs(val bool) *freelancersSearchService {
	s.jobs = &val
	return s
}

func (s *freelancersSearchService) SetBalanceDetails(val bool) *freelancersSearchService {
	s.balanceDetails = &val
	return s
}

func (s *freelancersSearchService) SetQualificationDetails(val bool) *freelancersSearchService {
	s.qualificationDetails = &val
	return s
}

func (s *freelancersSearchService) SetMembershipDetails(val bool) *freelancersSearchService {
	s.membershipDetails = &val
	return s
}

func (s *freelancersSearchService) SetFinancialDetails(val bool) *freelancersSearchService {
	s.financialDetails = &val
	return s
}

func (s *freelancersSearchService) SetLocationDetails(val bool) *freelancersSearchService {
	s.locationDetails = &val
	return s
}

func (s *freelancersSearchService) SetPortfolioDetails(val bool) *freelancersSearchService {
	s.portfolioDetails = &val
	return s
}

func (s *freelancersSearchService) SetPreferredDetails(val bool) *freelancersSearchService {
	s.preferredDetails = &val
	return s
}

func (s *freelancersSearchService) SetBadgeDetails(val bool) *freelancersSearchService {
	s.badgeDetails = &val
	return s
}

func (s *freelancersSearchService) SetStatus(val bool) *freelancersSearchService {
	s.status = &val
	return s
}

func (s *freelancersSearchService) SetReputation(val bool) *freelancersSearchService {
	s.reputation = &val
	return s
}

func (s *freelancersSearchService) SetEmployerReputation(val bool) *freelancersSearchService {
	s.employerReputation = &val
	return s
}

func (s *freelancersSearchService) SetReputationExtra(val bool) *freelancersSearchService {
	s.reputationExtra = &val
	return s
}

func (s *freelancersSearchService) SetEmployerReputationExtra(val bool) *freelancersSearchService {
	s.employerReputationExtra = &val
	return s
}

func (s *freelancersSearchService) SetCoverImage(val bool) *freelancersSearchService {
	s.coverImage = &val
	return s
}

func (s *freelancersSearchService) SetPastCoverImage(val bool) *freelancersSearchService {
	s.pastCoverImage = &val
	return s
}

func (s *freelancersSearchService) SetMobileTracking(val bool) *freelancersSearchService {
	s.mobileTracking = &val
	return s
}

func (s *freelancersSearchService) SetBidQualityDetails(val bool) *freelancersSearchService {
	s.bidQualityDetails = &val
	return s
}

func (s *freelancersSearchService) SetDepositMethods(val bool) *freelancersSearchService {
	s.depositMethods = &val
	return s
}

func (s *freelancersSearchService) SetUserRecommendations(val bool) *freelancersSearchService {
	s.userRecommendations = &val
	return s
}

func (s *freelancersSearchService) SetMarketingMobileNumber(val bool) *freelancersSearchService {
	s.marketingMobileNumber = &val
	return s
}

func (s *freelancersSearchService) SetSanctionDetails(val bool) *freelancersSearchService {
	s.sanctionDetails = &val
	return s
}

func (s *freelancersSearchService) SetLimitedAccount(val bool) *freelancersSearchService {
	s.limitedAccount = &val
	return s
}

func (s *freelancersSearchService) SetCompletedUserRelevantJobCount(val bool) *freelancersSearchService {
	s.completedUserRelevantJobCount = &val
	return s
}

func (s *freelancersSearchService) SetEquipmentGroupDetails(val bool) *freelancersSearchService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *freelancersSearchService) SetJobRanks(val bool) *freelancersSearchService {
	s.jobRanks = &val
	return s
}

func (s *freelancersSearchService) SetJobSeoDetails(val bool) *freelancersSearchService {
	s.jobSeoDetails = &val
	return s
}

func (s *freelancersSearchService) SetRisingStar(val bool) *freelancersSearchService {
	s.risingStar = &val
	return s
}

func (s *freelancersSearchService) SetShareholderDetails(val bool) *freelancersSearchService {
	s.shareholderDetails = &val
	return s
}

func (s *freelancersSearchService) SetStaffDetails(val bool) *freelancersSearchService {
	s.staffDetails = &val
	return s
}

func (s *freelancersSearchService) SetPoolDetails(val bool) *freelancersSearchService {
	s.poolDetails = &val
	return s
}

func (s *freelancersSearchService) SetLimit(val int) *freelancersSearchService {
	s.limit = &val
	return s
}

func (s *freelancersSearchService) SetOffset(val int) *freelancersSearchService {
	s.offset = &val
	return s
}

func (s *freelancersSearchService) SetCompact(val bool) *freelancersSearchService {
	s.compact = &val
	return s
}

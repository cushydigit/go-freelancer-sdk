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

// query	string (optional) Example: janedoe design logo
// filterReturns freelancers with the query appearing in the username, public name, profile description, or skill name
func (s *freelancersSearchService) SetQuery(val string) *freelancersSearchService {
	s.query = &val
	return s
}

// jobs[]	array[number] (optional) Example: 1, 2
// filterReturns freelancers with the specified jobs
func (s *freelancersSearchService) SetJobsIDs(jobsIDs []int32) *freelancersSearchService {
	s.jobsIDs = jobsIDs
	return s
}

// skills[]	array[number] (optional) Example: 3, 4
// filterReturns freelancers with the specified skills. Note that while jobs[] is still used, this filter is preferred to avoid any potential errors that may arise in the webapp since jobs[] shares the same name as the jobs projection.
func (s *freelancersSearchService) SetSkills(skills []int32) *freelancersSearchService {
	s.skills = skills
	return s
}

// countries[]	array[string] (optional) Example: Philippines, Australia
// filterReturns freelancers with the specified countries
func (s *freelancersSearchService) SetCountries(countries []string) *freelancersSearchService {
	s.countries = countries
	return s
}

// hourly_rate_min	number (optional) Example: 10
// filterReturns freelancers with the specified minimum hourly rate
func (s *freelancersSearchService) SetHourlyRateMin(val int) *freelancersSearchService {
	s.hourlyRateMin = &val
	return s
}

// hourly_rate_max	number (optional) Example: 100
// filterReturns freelancers with the specified maximum hourly rate
func (s *freelancersSearchService) SetHourlyRateMax(val int) *freelancersSearchService {
	s.hourlyRateMax = &val
	return s
}

// review_count_min	number (optional) Example: 1
// filterReturns freelancers with the specified minimum review count
func (s *freelancersSearchService) SetReviewCountMin(val int) *freelancersSearchService {
	s.reviewCountMin = &val
	return s
}

// review_count_max	number (optional) Example: 100
// filterReturns freelancers with the specified maximum review count
func (s *freelancersSearchService) SetReviewCountMax(val int) *freelancersSearchService {
	s.reviewCountMax = &val
	return s
}

// online_only	boolean (optional)
// filterReturns only freelancers that are online
func (s *freelancersSearchService) SetOnlineOnly(val bool) *freelancersSearchService {
	s.onlineOnly = &val
	return s
}

// location_latitude	decimal (optional) Example: 12.31
// filterReturns freelancers that are close to the specified latitude
func (s *freelancersSearchService) SetLocationLatitude(locationLatitude float64) *freelancersSearchService {
	s.locationLatitude = &locationLatitude
	return s
}

// location_longitude	decimal (optional) Example: 12.31
// filterReturns freelancers that are close to the specified longitude
func (s *freelancersSearchService) SetLocationLongitude(locationLongitude float64) *freelancersSearchService {
	s.locationLongitude = &locationLongitude
	return s
}

// insignias[]	array[number] (optional) Example: 1, 2
// filterReturns freelancers with the specified examination IDs complete
func (s *freelancersSearchService) SetInsignias(insignias []int32) *freelancersSearchService {
	s.insignias = insignias
	return s
}

// pool_ids[]	array[string] (optional) Example: 1, 2
// filterReturns freelancers with the specified pool ids
func (s *freelancersSearchService) SetPoolIds(poolIds []int32) *freelancersSearchService {
	s.poolIds = poolIds
	return s
}

// ratings	decimal (optional) Example: 4.5
// filterReturns freelancers with a minimum rating
func (s *freelancersSearchService) SetRatings(val float32) *freelancersSearchService {
	s.ratings = &val
	return s
}

// sort_field	DirectorySortFieldEnum (optional)
// filterSorting field, by default searches by relevance
func (s *freelancersSearchService) SetSortField(val int) *freelancersSearchService {
	s.sortField = &val
	return s
}

// reverse_sort	boolean (optional)
// filterIf true, results appear in ascending order instead of descending order
func (s *freelancersSearchService) SetReverseSort(val bool) *freelancersSearchService {
	s.reverseSort = &val
	return s
}

// avatar	boolean (optional)
// projectionReturns the avatar of the selected user/users.
func (s *freelancersSearchService) SetAvatar(val bool) *freelancersSearchService {
	s.avatar = &val
	return s
}

// country_details	boolean (optional)
// projectionReturns the country flag/code of selected user/users.
func (s *freelancersSearchService) SetCountryDetails(val bool) *freelancersSearchService {
	s.countryDetails = &val
	return s
}

// profile_description	boolean (optional)
// projectionReturns the profile blurb of selected user/users.
func (s *freelancersSearchService) SetProfileDescription(val bool) *freelancersSearchService {
	s.profileDescription = &val
	return s
}

// display_info	boolean (optional)
// projectionReturns the display name of the selected user/users.
func (s *freelancersSearchService) SetDisplayInfo(val bool) *freelancersSearchService {
	s.displayInfo = &val
	return s
}

// jobs	boolean (optional)
// projectionReturns the jobs of the selected user/users.
func (s *freelancersSearchService) SetJobs(val bool) *freelancersSearchService {
	s.jobs = &val
	return s
}

// balance_details	boolean (optional)
// projectionReturns the currency balance of selected user/users.
func (s *freelancersSearchService) SetBalanceDetails(val bool) *freelancersSearchService {
	s.balanceDetails = &val
	return s
}

// qualification_details	boolean (optional)
// projectionReturns qualification exams completed by the user/users.
func (s *freelancersSearchService) SetQualificationDetails(val bool) *freelancersSearchService {
	s.qualificationDetails = &val
	return s
}

// membership_details	boolean (optional)
// projectionReturns the membership information of the user/users.
func (s *freelancersSearchService) SetMembershipDetails(val bool) *freelancersSearchService {
	s.membershipDetails = &val
	return s
}

// financial_details	boolean (optional)
// projectionReturns the financial information of the user/users.
func (s *freelancersSearchService) SetFinancialDetails(val bool) *freelancersSearchService {
	s.financialDetails = &val
	return s
}

// location_details	boolean (optional)
// projectionReturns the location information of the user/users.
func (s *freelancersSearchService) SetLocationDetails(val bool) *freelancersSearchService {
	s.locationDetails = &val
	return s
}

// portfolio_details	boolean (optional)
// projectionReturns the portfolio information of the user/users.
func (s *freelancersSearchService) SetPortfolioDetails(val bool) *freelancersSearchService {
	s.portfolioDetails = &val
	return s
}

// preferred_details	boolean (optional)
// projectionReturns the preferred information of the user/users.
func (s *freelancersSearchService) SetPreferredDetails(val bool) *freelancersSearchService {
	s.preferredDetails = &val
	return s
}

// badge_details	boolean (optional)
// projectionReturns the badges earned by the user/users.
func (s *freelancersSearchService) SetBadgeDetails(val bool) *freelancersSearchService {
	s.badgeDetails = &val
	return s
}

// status	boolean (optional)
// projectionReturns additional status information about the user/users.
func (s *freelancersSearchService) SetStatus(val bool) *freelancersSearchService {
	s.status = &val
	return s
}

// reputation	boolean (optional)
// projectionReturns the freelancer reputation of the selected user/users.
func (s *freelancersSearchService) SetReputation(val bool) *freelancersSearchService {
	s.reputation = &val
	return s
}

// employer_reputation	boolean (optional)
// projectionReturns the employer reputation of the selected user/users.
func (s *freelancersSearchService) SetEmployerReputation(val bool) *freelancersSearchService {
	s.employerReputation = &val
	return s
}

// reputation_extra	boolean (optional)
// projectionReturns the full freelancer reputation of the selected user/users.
func (s *freelancersSearchService) SetReputationExtra(val bool) *freelancersSearchService {
	s.reputationExtra = &val
	return s
}

// employer_reputation_extra	boolean (optional)
// projectionReturns the full employer reputation of the selected user/users.
func (s *freelancersSearchService) SetEmployerReputationExtra(val bool) *freelancersSearchService {
	s.employerReputationExtra = &val
	return s
}

// cover_image	boolean (optional)
// projectionReturns the profile picture of the user.
func (s *freelancersSearchService) SetCoverImage(val bool) *freelancersSearchService {
	s.coverImage = &val
	return s
}

// past_cover_images	boolean (optional)
// projectionReturns previous profile pictures of the user.
func (s *freelancersSearchService) SetPastCoverImage(val bool) *freelancersSearchService {
	s.pastCoverImage = &val
	return s
}

// mobile_tracking	boolean (optional)
// projectionReturns the mobile platforms used by the selected user/users.
func (s *freelancersSearchService) SetMobileTracking(val bool) *freelancersSearchService {
	s.mobileTracking = &val
	return s
}

// bid_quality_details	boolean (optional)
// projectionReturns the user’s bid quality details, includes bid quality score.
func (s *freelancersSearchService) SetBidQualityDetails(val bool) *freelancersSearchService {
	s.bidQualityDetails = &val
	return s
}

// deposit_methods	boolean (optional)
// projectionReturns the deposit methods accepted by the user.
func (s *freelancersSearchService) SetDepositMethods(val bool) *freelancersSearchService {
	s.depositMethods = &val
	return s
}

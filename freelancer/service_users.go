package freelancer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type listUsers struct {
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

func (s *listUsers) Do(ctx context.Context) (*ListUsersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersUsers),
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
	return execute[*ListUsersResponse](ctx, s.client, r)
}

// SetUsers sets the list of user IDs to search.
func (s *listUsers) SetUsers(vals []int64) *listUsers {
	s.users = vals
	return s
}

// SetUsernames sets the list of usernames to search.
func (s *listUsers) SetUsernames(vals []string) *listUsers {
	s.usernames = vals
	return s
}

// SetAvatar enables returning the avatar of the selected user/users.
func (s *listUsers) SetAvatar(val bool) *listUsers {
	s.avatar = &val
	return s
}

// SetCountryDetails enables returning the country flag/code of the selected user/users.
func (s *listUsers) SetCountryDetails(val bool) *listUsers {
	s.countryDetails = &val
	return s
}

// SetProfileDescription enables returning the profile blurb of the selected user/users.
func (s *listUsers) SetProfileDescription(val bool) *listUsers {
	s.profileDescription = &val
	return s
}

// SetDisplayInfo enables returning the display name of the selected user/users.
func (s *listUsers) SetDisplayInfo(val bool) *listUsers {
	s.displayInfo = &val
	return s
}

// SetJobs enables returning the jobs of the selected user/users.
func (s *listUsers) SetJobs(val bool) *listUsers {
	s.jobs = &val
	return s
}

// SetBalanceDetails enables returning the currency balance of the selected user/users.
func (s *listUsers) SetBalanceDetails(val bool) *listUsers {
	s.balanceDetails = &val
	return s
}

// SetQualificationDetails enables returning qualification exams completed by the user/users.
func (s *listUsers) SetQualificationDetails(val bool) *listUsers {
	s.qualificationDetails = &val
	return s
}

// SetMembershipDetails enables returning the membership information of the user/users.
func (s *listUsers) SetMembershipDetails(val bool) *listUsers {
	s.membershipDetails = &val
	return s
}

// SetFinancialDetails enables returning the financial information of the user/users.
func (s *listUsers) SetFinancialDetails(val bool) *listUsers {
	s.financialDetails = &val
	return s
}

// SetLocationDetails enables returning the location information of the user/users.
func (s *listUsers) SetLocationDetails(val bool) *listUsers {
	s.locationDetails = &val
	return s
}

// SetPortfolioDetails enables returning the portfolio information of the user/users.
func (s *listUsers) SetPortfolioDetails(val bool) *listUsers {
	s.portfolioDetails = &val
	return s
}

// SetPreferredDetails enables returning the preferred information of the user/users.
func (s *listUsers) SetPreferredDetails(val bool) *listUsers {
	s.preferredDetails = &val
	return s
}

// SetBadgeDetails enables returning the badges earned by the user/users.
func (s *listUsers) SetBadgeDetails(val bool) *listUsers {
	s.badgeDetails = &val
	return s
}

// SetStatus enables returning additional status information about the user/users.
func (s *listUsers) SetStatus(val bool) *listUsers {
	s.status = &val
	return s
}

// SetReputation enables returning the freelancer reputation of the selected user/users.
func (s *listUsers) SetReputation(val bool) *listUsers {
	s.reputation = &val
	return s
}

// SetEmployerReputation enables returning the employer reputation of the selected user/users.
func (s *listUsers) SetEmployerReputation(val bool) *listUsers {
	s.employerReputation = &val
	return s
}

// SetReputationExtra enables returning the full freelancer reputation of the selected user/users.
func (s *listUsers) SetReputationExtra(val bool) *listUsers {
	s.reputationExtra = &val
	return s
}

// SetEmployerReputationExtra enables returning the full employer reputation of the selected user/users.
func (s *listUsers) SetEmployerReputationExtra(val bool) *listUsers {
	s.employerReputationExtra = &val
	return s
}

// SetCoverImage enables returning the profile picture of the user.
func (s *listUsers) SetCoverImage(val bool) *listUsers {
	s.coverImage = &val
	return s
}

// SetPastCoverImage enables returning previous profile pictures of the user.
func (s *listUsers) SetPastCoverImage(val bool) *listUsers {
	s.pastCoverImage = &val
	return s
}

// SetMobileTracking enables returning the mobile platforms used by the selected user/users.
func (s *listUsers) SetMobileTracking(val bool) *listUsers {
	s.mobileTracking = &val
	return s
}

// SetBidQualityDetails enables returning the user's bid quality details, including bid quality score.
func (s *listUsers) SetBidQualityDetails(val bool) *listUsers {
	s.bidQualityDetails = &val
	return s
}

// SetDepositMethods enables returning the list of deposit methods and associated fees.
func (s *listUsers) SetDepositMethods(val bool) *listUsers {
	s.depositMethods = &val
	return s
}

// SetUserRecommendations enables returning recommendations count of selected user/users.
func (s *listUsers) SetUserRecommendations(val bool) *listUsers {
	s.userRecommendations = &val
	return s
}

// SetMarketingMobileNumber enables returning the mobile number used by recruiters to contact the user.
func (s *listUsers) SetMarketingMobileNumber(val bool) *listUsers {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails enables returning the end timestamp of any sanction given to the user.
func (s *listUsers) SetSanctionDetails(val bool) *listUsers {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount enables returning the limited account status of the user.
func (s *listUsers) SetLimitedAccount(val bool) *listUsers {
	s.limitedAccount = &val
	return s
}

// SetCompletedUserRelevantJobCount enables returning the number of completed relevant jobs.
func (s *listUsers) SetCompletedUserRelevantJobCount(val bool) *listUsers {
	s.completedUserRelevantJobCount = &val
	return s
}

// SetEquipmentGroupDetails enables returning the user's equipment groups and items.
func (s *listUsers) SetEquipmentGroupDetails(val bool) *listUsers {
	s.equipmentGroupDetails = &val
	return s
}

// SetJobRanks enables returning the user's job ranks (requires reputation and jobs projections).
func (s *listUsers) SetJobRanks(val bool) *listUsers {
	s.jobRanks = &val
	return s
}

// SetShareholderDetails enables returning whether the user is part of the Shareholder Program.
func (s *listUsers) SetShareholderDetails(val bool) *listUsers {
	s.shareholderDetails = &val
	return s
}

// SetRisingStar enables returning the user's Rising Star program status.
func (s *listUsers) SetRisingStar(val bool) *listUsers {
	s.risingStar = &val
	return s
}

// SetStaffDetails enables returning whether the user is a Freelancer staff member.
func (s *listUsers) SetStaffDetails(val bool) *listUsers {
	s.staffDetails = &val
	return s
}

// SetJobSeoDetails enables returning SEO details of the user’s jobs (requires jobs projection).
func (s *listUsers) SetJobSeoDetails(val bool) *listUsers {
	s.jobSeoDetails = &val
	return s
}

type getUserByID struct {
	client *Client
	userID int64
}

func (s *getUserByID) Do(ctx context.Context) (*GetUserResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(UsersUsers), strconv.FormatInt(s.userID, 10)),
	}
	return execute[*GetUserResponse](ctx, s.client, r)
}

type searchFreelancers struct {
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

func (s *searchFreelancers) Do(ctx context.Context) (*SearchFreelancersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(UsersFreelancers),
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
	return execute[*SearchFreelancersResponse](ctx, s.client, r)
}

// query	string (optional) Example: janedoe design logo
// filterReturns freelancers with the query appearing in the username, public name, profile description, or skill name
func (s *searchFreelancers) SetQuery(val string) *searchFreelancers {
	s.query = &val
	return s
}

// jobs[]	array[number] (optional) Example: 1, 2
// filterReturns freelancers with the specified jobs
func (s *searchFreelancers) SetJobsIDs(jobsIDs []int32) *searchFreelancers {
	s.jobsIDs = jobsIDs
	return s
}

// skills[]	array[number] (optional) Example: 3, 4
// filterReturns freelancers with the specified skills. Note that while jobs[] is still used, this filter is preferred to avoid any potential errors that may arise in the webapp since jobs[] shares the same name as the jobs projection.
func (s *searchFreelancers) SetSkills(skills []int32) *searchFreelancers {
	s.skills = skills
	return s
}

// countries[]	array[string] (optional) Example: Philippines, Australia
// filterReturns freelancers with the specified countries
func (s *searchFreelancers) SetCountries(countries []string) *searchFreelancers {
	s.countries = countries
	return s
}

// hourly_rate_min	number (optional) Example: 10
// filterReturns freelancers with the specified minimum hourly rate
func (s *searchFreelancers) SetHourlyRateMin(val int) *searchFreelancers {
	s.hourlyRateMin = &val
	return s
}

// hourly_rate_max	number (optional) Example: 100
// filterReturns freelancers with the specified maximum hourly rate
func (s *searchFreelancers) SetHourlyRateMax(val int) *searchFreelancers {
	s.hourlyRateMax = &val
	return s
}

// review_count_min	number (optional) Example: 1
// filterReturns freelancers with the specified minimum review count
func (s *searchFreelancers) SetReviewCountMin(val int) *searchFreelancers {
	s.reviewCountMin = &val
	return s
}

// review_count_max	number (optional) Example: 100
// filterReturns freelancers with the specified maximum review count
func (s *searchFreelancers) SetReviewCountMax(val int) *searchFreelancers {
	s.reviewCountMax = &val
	return s
}

// online_only	boolean (optional)
// filterReturns only freelancers that are online
func (s *searchFreelancers) SetOnlineOnly(val bool) *searchFreelancers {
	s.onlineOnly = &val
	return s
}

// location_latitude	decimal (optional) Example: 12.31
// filterReturns freelancers that are close to the specified latitude
func (s *searchFreelancers) SetLocationLatitude(locationLatitude float64) *searchFreelancers {
	s.locationLatitude = &locationLatitude
	return s
}

// location_longitude	decimal (optional) Example: 12.31
// filterReturns freelancers that are close to the specified longitude
func (s *searchFreelancers) SetLocationLongitude(locationLongitude float64) *searchFreelancers {
	s.locationLongitude = &locationLongitude
	return s
}

// insignias[]	array[number] (optional) Example: 1, 2
// filterReturns freelancers with the specified examination IDs complete
func (s *searchFreelancers) SetInsignias(insignias []int32) *searchFreelancers {
	s.insignias = insignias
	return s
}

// pool_ids[]	array[string] (optional) Example: 1, 2
// filterReturns freelancers with the specified pool ids
func (s *searchFreelancers) SetPoolIds(poolIds []int32) *searchFreelancers {
	s.poolIds = poolIds
	return s
}

// ratings	decimal (optional) Example: 4.5
// filterReturns freelancers with a minimum rating
func (s *searchFreelancers) SetRatings(val float32) *searchFreelancers {
	s.ratings = &val
	return s
}

// sort_field	DirectorySortFieldEnum (optional)
// filterSorting field, by default searches by relevance
func (s *searchFreelancers) SetSortField(val int) *searchFreelancers {
	s.sortField = &val
	return s
}

// reverse_sort	boolean (optional)
// filterIf true, results appear in ascending order instead of descending order
func (s *searchFreelancers) SetReverseSort(val bool) *searchFreelancers {
	s.reverseSort = &val
	return s
}

// avatar	boolean (optional)
// projectionReturns the avatar of the selected user/users.
func (s *searchFreelancers) SetAvatar(val bool) *searchFreelancers {
	s.avatar = &val
	return s
}

// country_details	boolean (optional)
// projectionReturns the country flag/code of selected user/users.
func (s *searchFreelancers) SetCountryDetails(val bool) *searchFreelancers {
	s.countryDetails = &val
	return s
}

// profile_description	boolean (optional)
// projectionReturns the profile blurb of selected user/users.
func (s *searchFreelancers) SetProfileDescription(val bool) *searchFreelancers {
	s.profileDescription = &val
	return s
}

// display_info	boolean (optional)
// projectionReturns the display name of the selected user/users.
func (s *searchFreelancers) SetDisplayInfo(val bool) *searchFreelancers {
	s.displayInfo = &val
	return s
}

// jobs	boolean (optional)
// projectionReturns the jobs of the selected user/users.
func (s *searchFreelancers) SetJobs(val bool) *searchFreelancers {
	s.jobs = &val
	return s
}

// balance_details	boolean (optional)
// projectionReturns the currency balance of selected user/users.
func (s *searchFreelancers) SetBalanceDetails(val bool) *searchFreelancers {
	s.balanceDetails = &val
	return s
}

// qualification_details	boolean (optional)
// projectionReturns qualification exams completed by the user/users.
func (s *searchFreelancers) SetQualificationDetails(val bool) *searchFreelancers {
	s.qualificationDetails = &val
	return s
}

// membership_details	boolean (optional)
// projectionReturns the membership information of the user/users.
func (s *searchFreelancers) SetMembershipDetails(val bool) *searchFreelancers {
	s.membershipDetails = &val
	return s
}

// financial_details	boolean (optional)
// projectionReturns the financial information of the user/users.
func (s *searchFreelancers) SetFinancialDetails(val bool) *searchFreelancers {
	s.financialDetails = &val
	return s
}

// location_details	boolean (optional)
// projectionReturns the location information of the user/users.
func (s *searchFreelancers) SetLocationDetails(val bool) *searchFreelancers {
	s.locationDetails = &val
	return s
}

// portfolio_details	boolean (optional)
// projectionReturns the portfolio information of the user/users.
func (s *searchFreelancers) SetPortfolioDetails(val bool) *searchFreelancers {
	s.portfolioDetails = &val
	return s
}

// preferred_details	boolean (optional)
// projectionReturns the preferred information of the user/users.
func (s *searchFreelancers) SetPreferredDetails(val bool) *searchFreelancers {
	s.preferredDetails = &val
	return s
}

// badge_details	boolean (optional)
// projectionReturns the badges earned by the user/users.
func (s *searchFreelancers) SetBadgeDetails(val bool) *searchFreelancers {
	s.badgeDetails = &val
	return s
}

// status	boolean (optional)
// projectionReturns additional status information about the user/users.
func (s *searchFreelancers) SetStatus(val bool) *searchFreelancers {
	s.status = &val
	return s
}

// reputation	boolean (optional)
// projectionReturns the freelancer reputation of the selected user/users.
func (s *searchFreelancers) SetReputation(val bool) *searchFreelancers {
	s.reputation = &val
	return s
}

// employer_reputation	boolean (optional)
// projectionReturns the employer reputation of the selected user/users.
func (s *searchFreelancers) SetEmployerReputation(val bool) *searchFreelancers {
	s.employerReputation = &val
	return s
}

// reputation_extra	boolean (optional)
// projectionReturns the full freelancer reputation of the selected user/users.
func (s *searchFreelancers) SetReputationExtra(val bool) *searchFreelancers {
	s.reputationExtra = &val
	return s
}

// employer_reputation_extra	boolean (optional)
// projectionReturns the full employer reputation of the selected user/users.
func (s *searchFreelancers) SetEmployerReputationExtra(val bool) *searchFreelancers {
	s.employerReputationExtra = &val
	return s
}

// cover_image	boolean (optional)
// projectionReturns the profile picture of the user.
func (s *searchFreelancers) SetCoverImage(val bool) *searchFreelancers {
	s.coverImage = &val
	return s
}

// past_cover_images	boolean (optional)
// projectionReturns previous profile pictures of the user.
func (s *searchFreelancers) SetPastCoverImage(val bool) *searchFreelancers {
	s.pastCoverImage = &val
	return s
}

// mobile_tracking	boolean (optional)
// projectionReturns the mobile platforms used by the selected user/users.
func (s *searchFreelancers) SetMobileTracking(val bool) *searchFreelancers {
	s.mobileTracking = &val
	return s
}

// bid_quality_details	boolean (optional)
// projectionReturns the user’s bid quality details, includes bid quality score.
func (s *searchFreelancers) SetBidQualityDetails(val bool) *searchFreelancers {
	s.bidQualityDetails = &val
	return s
}

// deposit_methods	boolean (optional)
// projectionReturns the deposit methods accepted by the user.
func (s *searchFreelancers) SetDepositMethods(val bool) *searchFreelancers {
	s.depositMethods = &val
	return s
}

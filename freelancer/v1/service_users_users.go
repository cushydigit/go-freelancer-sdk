package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GetUserService struct {
	client *Client
}

type GetUserResponse struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

func (s *GetUserService) Do(ctx context.Context, userID int) (*GetUserResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", string(USERS_USERS), strconv.Itoa(userID)),
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &GetUserResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

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
	if err := json.Unmarshal(data, res); err != nil {
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

type ListFreelancersService struct {
	client *Client
	query  string
	// jobs[]
	jobsIDs                       []int // I have to rename the field to jobsIDs we have another jobs as bool
	skills                        []int
	coutries                      []string
	hourlyRateMin                 *int
	hourlyRateMax                 *int
	reviewCountMin                *int
	reviewCountMax                *int
	onlineOnly                    bool
	locationLatitude              *float64
	locationLongitude             *float64
	insignias                     []int
	poolIds                       []int
	ratings                       float32
	sortField                     *int // DirectorySortFieldEnum? could not found in freelancer.com docs
	reverseSort                   bool
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
	poolDetails                   bool
	limit                         *int
	offset                        *int
	compact                       bool
}

type ListFreelancersResponse struct {
	Status    string           `json:"status"`
	RequestID string           `json:"request_id"`
	Result    FreelancerResult `json:"result"`
}

type FreelancerResult struct {
	Users []User `json:"users"`
}

func (s *ListFreelancersService) Do(ctx context.Context) (*ListFreelancersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(USERS_FREELANCERS),
	}
	if s.query != "" {
		r.addParam("query", s.query)
	}

	for _, jobID := range s.jobsIDs {
		r.addParam("jobs[]", jobID)
	}
	for _, skill := range s.skills {
		r.addParam("skills[]", skill)
	}
	for _, country := range s.coutries {
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
	if s.onlineOnly {
		r.addParam("online_only", s.onlineOnly)
	}
	if s.locationLatitude != nil {
		r.addParam("location_latitude", *s.locationLatitude)
	}
	if s.locationLongitude != nil {
		r.addParam("location_longitude", *s.locationLongitude)
	}
	if s.reverseSort {
		r.addParam("reverse_sort", s.reverseSort)
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
	if s.poolDetails {
		r.setParam("pool_details", s.poolDetails)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.compact {
		r.setParam("compact", s.compact)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListFreelancersResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListFreelancersService) SetQuery(query string) *ListFreelancersService {
	s.query = query
	return s
}

func (s *ListFreelancersService) SetJobsIDs(jobsIDs []int) *ListFreelancersService {
	s.jobsIDs = jobsIDs
	return s
}

func (s *ListFreelancersService) SetSkills(skills []int) *ListFreelancersService {
	s.skills = skills
	return s
}

func (s *ListFreelancersService) SetCoutries(coutries []string) *ListFreelancersService {
	s.coutries = coutries
	return s
}

func (s *ListFreelancersService) SetHourlyRateMin(hourlyRateMin int) *ListFreelancersService {
	s.hourlyRateMin = &hourlyRateMin
	return s
}

func (s *ListFreelancersService) SetHourlyRateMax(hourlyRateMax int) *ListFreelancersService {
	s.hourlyRateMax = &hourlyRateMax
	return s
}

func (s *ListFreelancersService) SetReviewCountMin(reviewCountMin int) *ListFreelancersService {
	s.reviewCountMin = &reviewCountMin
	return s
}

func (s *ListFreelancersService) SetReviewCountMax(reviewCountMax int) *ListFreelancersService {
	s.reviewCountMax = &reviewCountMax
	return s
}

func (s *ListFreelancersService) SetOnlineOnly(onlineOnly bool) *ListFreelancersService {
	s.onlineOnly = onlineOnly
	return s
}

func (s *ListFreelancersService) SetLocationLatitude(locationLatitude float64) *ListFreelancersService {
	s.locationLatitude = &locationLatitude
	return s
}

func (s *ListFreelancersService) SetLocationLongitude(locationLongitude float64) *ListFreelancersService {
	s.locationLongitude = &locationLongitude
	return s
}

func (s *ListFreelancersService) SetInsignias(insignias []int) *ListFreelancersService {
	s.insignias = insignias
	return s
}

func (s *ListFreelancersService) SetPoolIds(poolIds []int) *ListFreelancersService {
	s.poolIds = poolIds
	return s
}

func (s *ListFreelancersService) SetRatings(ratings float32) *ListFreelancersService {
	s.ratings = ratings
	return s
}

func (s *ListFreelancersService) SetSortField(sortField *int) *ListFreelancersService {
	s.sortField = sortField
	return s
}

func (s *ListFreelancersService) SetReverseSort(reverseSort bool) *ListFreelancersService {
	s.reverseSort = reverseSort
	return s
}

func (s *ListFreelancersService) SetAvatar(avatar bool) *ListFreelancersService {
	s.avatar = avatar
	return s
}

func (s *ListFreelancersService) SetCountryDetails(countryDetails bool) *ListFreelancersService {
	s.countryDetails = countryDetails
	return s
}

func (s *ListFreelancersService) SetProfileDescription(profileDescription bool) *ListFreelancersService {
	s.profileDescription = profileDescription
	return s
}

func (s *ListFreelancersService) SetDisplayInfo(displayInfo bool) *ListFreelancersService {
	s.displayInfo = displayInfo
	return s
}

func (s *ListFreelancersService) SetJobs(jobs bool) *ListFreelancersService {
	s.jobs = jobs
	return s
}

func (s *ListFreelancersService) SetBalanceDetails(balanceDetails bool) *ListFreelancersService {
	s.balanceDetails = balanceDetails
	return s
}

func (s *ListFreelancersService) SetQualificationDetails(qualificationDetails bool) *ListFreelancersService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *ListFreelancersService) SetMembershipDetails(membershipDetails bool) *ListFreelancersService {
	s.membershipDetails = membershipDetails
	return s
}

func (s *ListFreelancersService) SetFinancialDetails(financialDetails bool) *ListFreelancersService {
	s.financialDetails = financialDetails
	return s
}

func (s *ListFreelancersService) SetLocationDetails(locationDetails bool) *ListFreelancersService {
	s.locationDetails = locationDetails
	return s
}

func (s *ListFreelancersService) SetPortfolioDetails(portfolioDetails bool) *ListFreelancersService {
	s.portfolioDetails = portfolioDetails
	return s
}

func (s *ListFreelancersService) SetPreferredDetails(preferredDetails bool) *ListFreelancersService {
	s.preferredDetails = preferredDetails
	return s
}

func (s *ListFreelancersService) SetBadgeDetails(badgeDetails bool) *ListFreelancersService {
	s.badgeDetails = badgeDetails
	return s
}

func (s *ListFreelancersService) SetStatus(status bool) *ListFreelancersService {
	s.status = status
	return s
}

func (s *ListFreelancersService) SetReputation(reputation bool) *ListFreelancersService {
	s.reputation = reputation
	return s
}

func (s *ListFreelancersService) SetEmployerReputation(employerReputation bool) *ListFreelancersService {
	s.employerReputation = employerReputation
	return s
}

func (s *ListFreelancersService) SetReputationExtra(reputationExtra bool) *ListFreelancersService {
	s.reputationExtra = reputationExtra
	return s
}

func (s *ListFreelancersService) SetEmployerReputationExtra(employerReputationExtra bool) *ListFreelancersService {
	s.employerReputationExtra = employerReputationExtra
	return s
}

func (s *ListFreelancersService) SetCoverImage(coverImage bool) *ListFreelancersService {
	s.coverImage = coverImage
	return s
}

func (s *ListFreelancersService) SetPastCoverImage(pastCoverImage bool) *ListFreelancersService {
	s.pastCoverImage = pastCoverImage
	return s
}

func (s *ListFreelancersService) SetMobileTracking(mobileTracking bool) *ListFreelancersService {
	s.mobileTracking = mobileTracking
	return s
}

func (s *ListFreelancersService) SetBidQualityDetails(bidQualityDetails bool) *ListFreelancersService {
	s.bidQualityDetails = bidQualityDetails
	return s
}

func (s *ListFreelancersService) SetDepositMethods(depositMethods bool) *ListFreelancersService {
	s.depositMethods = depositMethods
	return s
}

func (s *ListFreelancersService) SetUserRecommendations(userRecommendations bool) *ListFreelancersService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ListFreelancersService) SetMarketingMobileNumber(marketingMobileNumber bool) *ListFreelancersService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ListFreelancersService) SetSanctionDetails(sanctionDetails bool) *ListFreelancersService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ListFreelancersService) SetLimitedAccount(limitedAccount bool) *ListFreelancersService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ListFreelancersService) SetCompletedUserRelevantJobCount(completedUserRelevantJobCount bool) *ListFreelancersService {
	s.completedUserRelevantJobCount = completedUserRelevantJobCount
	return s
}

func (s *ListFreelancersService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ListFreelancersService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *ListFreelancersService) SetJobRanks(jobRanks bool) *ListFreelancersService {
	s.jobRanks = jobRanks
	return s
}

func (s *ListFreelancersService) SetJobSeoDetails(jobSeoDetails bool) *ListFreelancersService {
	s.jobSeoDetails = jobSeoDetails
	return s
}

func (s *ListFreelancersService) SetRisingStar(risingStar bool) *ListFreelancersService {
	s.risingStar = risingStar
	return s
}

func (s *ListFreelancersService) SetShareholderDetails(shareholderDetails bool) *ListFreelancersService {
	s.shareholderDetails = shareholderDetails
	return s
}

func (s *ListFreelancersService) SetStaffDetails(staffDetails bool) *ListFreelancersService {
	s.staffDetails = staffDetails
	return s
}

func (s *ListFreelancersService) SetPoolDetails(poolDetails bool) *ListFreelancersService {
	s.poolDetails = poolDetails
	return s
}

func (s *ListFreelancersService) SetLimit(limit *int) *ListFreelancersService {
	s.limit = limit
	return s
}

func (s *ListFreelancersService) SetOffset(offset *int) *ListFreelancersService {
	s.offset = offset
	return s
}

func (s *ListFreelancersService) SetCompact(compact bool) *ListFreelancersService {
	s.compact = compact
	return s
}

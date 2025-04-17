package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListActiveProjectsService struct {
	client                      *Client
	query                       string
	projectTypes                []ProjectType
	projectUpgrades             []ProjectUpgradeType
	sortFields                  []SortFieldsType
	contestUpgrades             []ContestUpgradeType
	minAvgPrice                 float64
	maxAvgPrice                 float64
	minAvgHourlyRate            float64
	maxAvgHourlyRate            float64
	minPrice                    float64
	maxPrice                    float64
	minHourlyRate               float64
	maxHourlyRate               float64
	jobs                        []int64
	countries                   []string
	languages                   []string
	latitude                    float64
	longitude                   float64
	fromTime                    int64
	toTime                      int64
	sortField                   SortFieldsType
	projectIDs                  []int64
	topRightLatitude            float64
	topRightLongitude           float64
	bottomLeftLatitude          float64
	bottomLeftLongitude         float64
	reverseSort                 bool
	orSearchQuery               string
	highlightPreTags            string
	highlightPostTags           string
	fullDescription             bool
	jobDetails                  bool
	upgradeDetails              bool
	userDetails                 bool
	locationDetails             bool
	ndaSignatureDetails         bool
	projectCollaborationDetails bool
	userAvatar                  bool
	userCountryDetails          bool
	userProfileDescription      bool
	userDisplayInfo             bool
	userJobs                    bool
	userBalanceDetails          bool
	userQualificationDetails    bool
	userMembershipDetails       bool
	userFinancialDetails        bool
	userLocationDetails         bool
	userPortfolioDetails        bool
	userPreferredDetails        bool
	userBadgeDetails            bool
	userStatus                  bool
	userReputation              bool
	userEmployerReputation      bool
	userReputationExtra         bool
	userEmployerReputationExtra bool
	userCoverImage              bool
	userPastCoverImage          bool
	userRecommendations         bool
	userResponsiveness          bool
	corporateUsers              bool
	marketingMobileNumber       bool
	sanctionDetails             bool
	limitedAccount              bool
	equipmentGroupDetails       bool
	limit                       int
	offset                      int
	compact                     bool
}

// Do perform GET request on endpoint "projects/0.1/projects/active/""
func (s *ListActiveProjectsService) Do(ctx context.Context) (*ListProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/projects/active",
	}
	if s.query != "" {
		r.addParam("query", s.query)
	}
	if s.projectTypes != nil {
		r.addParam("project_types", s.projectTypes)
	}
	if s.projectUpgrades != nil {
		r.addParam("project_upgrades", s.projectUpgrades)
	}
	if s.sortFields != nil {
		r.addParam("sort_fields", s.sortFields)
	}
	if s.contestUpgrades != nil {
		r.addParam("contest_upgrades", s.contestUpgrades)
	}
	if s.minAvgPrice != 0 {
		r.addParam("min_avg_price", s.minAvgPrice)
	}
	if s.maxAvgPrice != 0 {
		r.addParam("max_avg_price", s.maxAvgPrice)
	}
	if s.minAvgHourlyRate != 0 {
		r.addParam("min_avg_hourly_rate", s.minAvgHourlyRate)
	}
	if s.maxAvgHourlyRate != 0 {
		r.addParam("max_avg_hourly_rate", s.maxAvgHourlyRate)
	}
	if s.minPrice != 0 {
		r.addParam("min_price", s.minPrice)
	}
	if s.maxPrice != 0 {
		r.addParam("max_price", s.maxPrice)
	}
	if s.minHourlyRate != 0 {
		r.addParam("min_hourly_rate", s.minHourlyRate)
	}
	if s.maxHourlyRate != 0 {
		r.addParam("max_hourly_rate", s.maxHourlyRate)
	}
	if s.jobs != nil {
		r.addParam("jobs", s.jobs)
	}
	if s.countries != nil {
		r.addParam("countries", s.countries)
	}
	if s.languages != nil {
		r.addParam("languages", s.languages)
	}
	if s.latitude != 0 {
		r.addParam("latitude", s.latitude)
	}
	if s.longitude != 0 {
		r.addParam("longitude", s.longitude)
	}
	if s.fromTime != 0 {
		r.addParam("from_time", s.fromTime)
	}
	if s.toTime != 0 {
		r.addParam("to_time", s.toTime)
	}
	if s.sortField != "" {
		r.addParam("sort_field", s.sortField)
	}
	if s.projectIDs != nil {
		r.addParam("project_ids", s.projectIDs)
	}
	if s.topRightLatitude != 0 {
		r.addParam("top_right_latitude", s.topRightLatitude)
	}
	if s.topRightLongitude != 0 {
		r.addParam("top_right_longitude", s.topRightLongitude)
	}
	if s.bottomLeftLatitude != 0 {
		r.addParam("bottom_left_latitude", s.bottomLeftLatitude)
	}
	if s.bottomLeftLongitude != 0 {
		r.addParam("bottom_left_longitude", s.bottomLeftLongitude)
	}
	if s.reverseSort {
		r.addParam("reverse_sort", s.reverseSort)
	}
	if s.orSearchQuery != "" {
		r.addParam("or_search_query", s.orSearchQuery)
	}
	if s.highlightPreTags != "" {
		r.addParam("highlight_pre_tags", s.highlightPreTags)
	}
	if s.highlightPostTags != "" {
		r.addParam("highlight_post_tags", s.highlightPostTags)
	}
	if s.fullDescription {
		r.addParam("full_description", s.fullDescription)
	}
	if s.jobDetails {
		r.addParam("job_details", s.jobDetails)
	}
	if s.upgradeDetails {
		r.addParam("upgrade_details", s.upgradeDetails)
	}
	if s.userStatus {
		r.addParam("user_status", s.userStatus)
	}
	if s.userEmployerReputation {
		r.addParam("user_employer_reputation", s.userEmployerReputation)
	}
	if s.userReputationExtra {
		r.addParam("user_reputation_extra", s.userReputationExtra)
	}
	if s.userEmployerReputationExtra {
		r.addParam("user_employer_reputation_extra", s.userEmployerReputationExtra)
	}
	if s.userCoverImage {
		r.addParam("user_cover_image", s.userCoverImage)
	}
	if s.userPastCoverImage {
		r.addParam("user_past_cover_image", s.userPastCoverImage)
	}
	if s.userRecommendations {
		r.addParam("user_recommendations", s.userRecommendations)
	}
	if s.userResponsiveness {
		r.addParam("user_responsiveness", s.userResponsiveness)
	}
	if s.corporateUsers {
		r.addParam("corporate_users", s.corporateUsers)
	}
	if s.marketingMobileNumber {
		r.addParam("marketing_mobile_number", s.marketingMobileNumber)
	}
	if s.sanctionDetails {
		r.addParam("sanction_details", s.sanctionDetails)
	}
	if s.limitedAccount {
		r.addParam("limited_account", s.limitedAccount)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ListProjectsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ListActiveProjectsService) SetUserResponsiveness(userResponsiveness bool) *ListActiveProjectsService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ListActiveProjectsService) SetCorporateUsers(corporateUsers bool) *ListActiveProjectsService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ListActiveProjectsService) SetMarketingMobileNumber(marketingMobileNumber bool) *ListActiveProjectsService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ListActiveProjectsService) SetSanctionDetails(sanctionDetails bool) *ListActiveProjectsService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ListActiveProjectsService) SetLimitedAccount(limitedAccount bool) *ListActiveProjectsService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ListActiveProjectsService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ListActiveProjectsService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *ListActiveProjectsService) SetUserDetails(userDetails bool) *ListActiveProjectsService {
	s.userDetails = userDetails
	return s
}
func (s *ListActiveProjectsService) SetLocationDetails(locationDetails bool) *ListActiveProjectsService {
	s.locationDetails = locationDetails
	return s
}

func (s *ListActiveProjectsService) SetNdaSignatureDetails(ndaSignatureDetails bool) *ListActiveProjectsService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ListActiveProjectsService) SetProjectCollaborationDetails(projectCollaborationDetails bool) *ListActiveProjectsService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ListActiveProjectsService) SetUserAvatar(userAvatar bool) *ListActiveProjectsService {
	s.userAvatar = userAvatar
	return s
}

func (s *ListActiveProjectsService) SetUserCountryDetails(userCountryDetails bool) *ListActiveProjectsService {
	s.userCountryDetails = userCountryDetails
	return s
}

func (s *ListActiveProjectsService) SetUserProfileDescription(userProfileDescription bool) *ListActiveProjectsService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ListActiveProjectsService) SetUserDisplayInfo(userDisplayInfo bool) *ListActiveProjectsService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ListActiveProjectsService) SetUserJobs(userJobs bool) *ListActiveProjectsService {
	s.userJobs = userJobs
	return s
}

func (s *ListActiveProjectsService) SetUserBalanceDetails(userBalanceDetails bool) *ListActiveProjectsService {
	s.userBalanceDetails = userBalanceDetails
	return s
}

func (s *ListActiveProjectsService) SetUserQualificationDetails(userQualificationDetails bool) *ListActiveProjectsService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ListActiveProjectsService) SetUserMembershipDetails(userMembershipDetails bool) *ListActiveProjectsService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ListActiveProjectsService) SetUserFinancialDetails(userFinancialDetails bool) *ListActiveProjectsService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ListActiveProjectsService) SetUserLocationDetails(userLocationDetails bool) *ListActiveProjectsService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ListActiveProjectsService) SetUserPortfolioDetails(userPortfolioDetails bool) *ListActiveProjectsService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ListActiveProjectsService) SetUserPreferredDetails(userPreferredDetails bool) *ListActiveProjectsService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ListActiveProjectsService) SetUserBadgeDetails(userBadgeDetails bool) *ListActiveProjectsService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ListActiveProjectsService) SetUserStatus(userStatus bool) *ListActiveProjectsService {
	s.userStatus = userStatus
	return s
}

func (s *ListActiveProjectsService) SetUserReputation(userReputation bool) *ListActiveProjectsService {
	s.userReputation = userReputation
	return s
}

func (s *ListActiveProjectsService) SetUserEmployerReputation(userEmployerReputation bool) *ListActiveProjectsService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ListActiveProjectsService) SetUserReputationExtra(userReputationExtra bool) *ListActiveProjectsService {
	s.userReputationExtra = userReputationExtra
	return s
}

func (s *ListActiveProjectsService) SetUserEmployerReputationExtra(userEmployerReputationExtra bool) *ListActiveProjectsService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ListActiveProjectsService) SetUserCoverImage(userCoverImage bool) *ListActiveProjectsService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ListActiveProjectsService) SetUserPastCoverImage(userPastCoverImage bool) *ListActiveProjectsService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ListActiveProjectsService) SetUserRecommendations(userRecommendations bool) *ListActiveProjectsService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ListActiveProjectsService) SetProjectTypes(projectTypes []ProjectType) *ListActiveProjectsService {
	s.projectTypes = projectTypes
	return s
}

func (s *ListActiveProjectsService) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *ListActiveProjectsService {
	s.projectUpgrades = projectUpgrades
	return s
}

func (s *ListActiveProjectsService) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *ListActiveProjectsService {
	s.contestUpgrades = contestUpgrades
	return s
}
func (s *ListActiveProjectsService) SetCompact(compact bool) *ListActiveProjectsService {
	s.compact = compact
	return s
}
func (s *ListActiveProjectsService) SetLimit(limit int) *ListActiveProjectsService {
	s.limit = limit
	return s
}

func (s *ListActiveProjectsService) SetOffset(offset int) *ListActiveProjectsService {
	s.offset = offset
	return s
}

func (s *ListActiveProjectsService) SetReverseSort(reverseSort bool) *ListActiveProjectsService {
	s.reverseSort = reverseSort
	return s
}

func (s *ListActiveProjectsService) SetFullDescription(fullDescription bool) *ListActiveProjectsService {
	s.fullDescription = fullDescription
	return s
}

func (s *ListActiveProjectsService) SetJobDetails(jobDetails bool) *ListActiveProjectsService {
	s.jobDetails = jobDetails
	return s
}

func (s *ListActiveProjectsService) SetUpgradeDetails(upgradeDetails bool) *ListActiveProjectsService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ListActiveProjectsService) SetQuery(query string) *ListActiveProjectsService {
	s.query = query
	return s
}

func (s *ListActiveProjectsService) SetMinAvgPrice(minAvgPrice float64) *ListActiveProjectsService {
	s.minAvgPrice = minAvgPrice
	return s
}

func (s *ListActiveProjectsService) SetMaxAvgPrice(maxAvgPrice float64) *ListActiveProjectsService {
	s.maxAvgPrice = maxAvgPrice
	return s
}

func (s *ListActiveProjectsService) SetMinAvgHourlyRate(minAvgHourlyRate float64) *ListActiveProjectsService {
	s.minAvgHourlyRate = minAvgHourlyRate
	return s
}

func (s *ListActiveProjectsService) SetMaxAvgHourlyRate(maxAvgHourlyRate float64) *ListActiveProjectsService {
	s.maxAvgHourlyRate = maxAvgHourlyRate
	return s
}

func (s *ListActiveProjectsService) SetMinPrice(minPrice float64) *ListActiveProjectsService {
	s.minPrice = minPrice
	return s
}

func (s *ListActiveProjectsService) SetMaxPrice(maxPrice float64) *ListActiveProjectsService {
	s.maxPrice = maxPrice
	return s
}

func (s *ListActiveProjectsService) SetMinHourlyRate(minHourlyRate float64) *ListActiveProjectsService {
	s.minHourlyRate = minHourlyRate
	return s
}

func (s *ListActiveProjectsService) SetMaxHourlyRate(maxHourlyRate float64) *ListActiveProjectsService {
	s.maxHourlyRate = maxHourlyRate
	return s
}

func (s *ListActiveProjectsService) SetJobs(jobs []int64) *ListActiveProjectsService {
	s.jobs = jobs
	return s
}

func (s *ListActiveProjectsService) SetCountries(countries []string) *ListActiveProjectsService {
	s.countries = countries
	return s
}

func (s *ListActiveProjectsService) SetLanguages(languages []string) *ListActiveProjectsService {
	s.languages = languages
	return s
}

func (s *ListActiveProjectsService) SetLatitude(latitude float64) *ListActiveProjectsService {
	s.latitude = latitude
	return s
}

func (s *ListActiveProjectsService) SetLongitude(longitude float64) *ListActiveProjectsService {
	s.longitude = longitude
	return s
}

func (s *ListActiveProjectsService) SetFromTime(fromTime int64) *ListActiveProjectsService {
	s.fromTime = fromTime
	return s
}

func (s *ListActiveProjectsService) SetToTime(toTime int64) *ListActiveProjectsService {
	s.toTime = toTime
	return s
}

func (s *ListActiveProjectsService) SetSortField(sortField SortFieldsType) *ListActiveProjectsService {
	s.sortField = sortField
	return s
}

func (s *ListActiveProjectsService) SetProjectIDs(projectIDs []int64) *ListActiveProjectsService {
	s.projectIDs = projectIDs
	return s
}

func (s *ListActiveProjectsService) SetTopRightLatitude(topRightLatitude float64) *ListActiveProjectsService {
	s.topRightLatitude = topRightLatitude
	return s
}

func (s *ListActiveProjectsService) SetTopRightLongitude(topRightLongitude float64) *ListActiveProjectsService {
	s.topRightLongitude = topRightLongitude
	return s
}

func (s *ListActiveProjectsService) SetBottomLeftLatitude(bottomLeftLatitude float64) *ListActiveProjectsService {
	s.bottomLeftLatitude = bottomLeftLatitude
	return s
}

func (s *ListActiveProjectsService) SetBottomLeftLongitude(bottomLeftLongitude float64) *ListActiveProjectsService {
	s.bottomLeftLongitude = bottomLeftLongitude
	return s
}

func (s *ListActiveProjectsService) SetOrSearchQuery(orSearchQuery string) *ListActiveProjectsService {
	s.orSearchQuery = orSearchQuery
	return s
}

func (s *ListActiveProjectsService) SetHighlightPreTags(highlightPreTags string) *ListActiveProjectsService {
	s.highlightPreTags = highlightPreTags
	return s
}

func (s *ListActiveProjectsService) SetHighlightPostTags(highlightPostTags string) *ListActiveProjectsService {
	s.highlightPostTags = highlightPostTags
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type SearchActiveProjectsService struct {
	client                      *Client
	query                       *string
	projectTypes                []ProjectType
	projectUpgrades             []ProjectUpgradeType
	contestUpgrades             []ContestUpgradeType
	minAvgPrice                 *float64
	maxAvgPrice                 *float64
	minAvgHourlyRate            *float64
	maxAvgHourlyRate            *float64
	minPrice                    *float64
	maxPrice                    *float64
	minHourlyRate               *float64
	maxHourlyRate               *float64
	jobs                        []int32
	countries                   []string
	languages                   []string
	latitude                    *float64
	longitude                   *float64
	fromTime                    *int64
	toTime                      *int64
	sortField                   *SortField
	projectIDs                  []int64
	topRightLatitude            *float64
	topRightLongitude           *float64
	bottomLeftLatitude          *float64
	bottomLeftLongitude         *float64
	reverseSort                 *bool
	orSearchQuery               *string
	highlightPreTags            *string
	highlightPostTags           *string
	fullDescription             *bool
	jobDetails                  *bool
	upgradeDetails              *bool
	userDetails                 *bool
	locationDetails             *bool
	ndaSignatureDetails         *bool
	projectCollaborationDetails *bool
	userAvatar                  *bool
	userCountryDetails          *bool
	userProfileDescription      *bool
	userDisplayInfo             *bool
	userJobs                    *bool
	userBalanceDetails          *bool
	userQualificationDetails    *bool
	userMembershipDetails       *bool
	userFinancialDetails        *bool
	userLocationDetails         *bool
	userPortfolioDetails        *bool
	userPreferredDetails        *bool
	userBadgeDetails            *bool
	userStatus                  *bool
	userReputation              *bool
	userEmployerReputation      *bool
	userReputationExtra         *bool
	userEmployerReputationExtra *bool
	userCoverImage              *bool
	userPastCoverImage          *bool
	userRecommendations         *bool
	userResponsiveness          *bool
	corporateUsers              *bool
	marketingMobileNumber       *bool
	sanctionDetails             *bool
	limitedAccount              *bool
	equipmentGroupDetails       *bool
	limit                       *int
	offset                      *int
	compact                     *bool
}

type SearchActiveProjectsResponse struct {
	Status    string               `json:"status"`
	RequestID string               `json:"request_id"`
	Result    ActiveProjectsResult `json:"result"`
}

type ActiveProjectsResult struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}

// Do perform GET request on endpoint "projects/0.1/projects/active/""
func (s *SearchActiveProjectsService) Do(ctx context.Context) (*SearchActiveProjectsResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_PROJECTS_ACTIVE),
	}
	if s.query != nil {
		r.addParam("query", s.query)
	}
	for _, projectType := range s.projectTypes {
		r.addParam("project_types[]", projectType)
	}
	for _, projectUpgrade := range s.projectUpgrades {
		r.addParam("project_upgrades[]", projectUpgrade)
	}
	for _, contestUpgrade := range s.contestUpgrades {
		r.addParam("contest_upgrades[]", contestUpgrade)
	}
	if s.contestUpgrades != nil {
		r.addParam("contest_upgrades", s.contestUpgrades)
	}
	if s.minAvgPrice != nil {
		r.addParam("min_avg_price", *s.minAvgPrice)
	}
	if s.maxAvgPrice != nil {
		r.addParam("max_avg_price", *s.maxAvgPrice)
	}
	if s.minAvgHourlyRate != nil {
		r.addParam("min_avg_hourly_rate", *s.minAvgHourlyRate)
	}
	if s.maxAvgHourlyRate != nil {
		r.addParam("max_avg_hourly_rate", *s.maxAvgHourlyRate)
	}
	if s.minPrice != nil {
		r.addParam("min_price", *s.minPrice)
	}
	if s.maxPrice != nil {
		r.addParam("max_price", *s.maxPrice)
	}
	if s.minHourlyRate != nil {
		r.addParam("min_hourly_rate", *s.minHourlyRate)
	}
	if s.maxHourlyRate != nil {
		r.addParam("max_hourly_rate", *s.maxHourlyRate)
	}
	for _, job := range s.jobs {
		r.addParam("jobs[]", job)
	}
	for _, country := range s.countries {
		r.addParam("countries[]", country)
	}
	for _, language := range s.languages {
		r.addParam("languages[]", language)
	}
	if s.latitude != nil {
		r.addParam("latitude", *s.latitude)
	}
	if s.longitude != nil {
		r.addParam("longitude", *s.longitude)
	}
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	if s.sortField != nil {
		r.addParam("sort_field", *s.sortField)
	}
	for _, projectID := range s.projectIDs {
		r.addParam("project_ids[]", projectID)
	}
	if s.topRightLatitude != nil {
		r.addParam("top_right_latitude", *s.topRightLatitude)
	}
	if s.topRightLongitude != nil {
		r.addParam("top_right_longitude", *s.topRightLongitude)
	}
	if s.bottomLeftLatitude != nil {
		r.addParam("bottom_left_latitude", *s.bottomLeftLatitude)
	}
	if s.bottomLeftLongitude != nil {
		r.addParam("bottom_left_longitude", *s.bottomLeftLongitude)
	}
	if s.reverseSort != nil {
		r.addParam("reverse_sort", *s.reverseSort)
	}
	if s.orSearchQuery != nil {
		r.addParam("or_search_query", *s.orSearchQuery)
	}
	if s.highlightPreTags != nil {
		r.addParam("highlight_pre_tags", *s.highlightPreTags)
	}
	if s.highlightPostTags != nil {
		r.addParam("highlight_post_tags", *s.highlightPostTags)
	}
	if s.fullDescription != nil {
		r.addParam("full_description", *s.fullDescription)
	}
	if s.jobDetails != nil {
		r.addParam("job_details", *s.jobDetails)
	}
	if s.upgradeDetails != nil {
		r.addParam("upgrade_details", *s.upgradeDetails)
	}
	if s.userStatus != nil {
		r.addParam("user_status", *s.userStatus)
	}
	if s.userEmployerReputation != nil {
		r.addParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.addParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.addParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.addParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.addParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.addParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.addParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.addParam("corporate_users", *s.corporateUsers)
	}
	if s.marketingMobileNumber != nil {
		r.addParam("marketing_mobile_number", *s.marketingMobileNumber)
	}
	if s.sanctionDetails != nil {
		r.addParam("sanction_details", *s.sanctionDetails)
	}
	if s.limitedAccount != nil {
		r.addParam("limited_account", *s.limitedAccount)
	}
	if s.limit != nil {
		r.addParam("limit", *s.limit)
	}
	if s.offset != nil {
		r.addParam("offset", *s.offset)
	}
	if s.compact != nil {
		r.addParam("compact", *s.compact)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &SearchActiveProjectsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SearchActiveProjectsService) SetUserResponsiveness(val bool) *SearchActiveProjectsService {
	s.userResponsiveness = &val
	return s
}

func (s *SearchActiveProjectsService) SetCorporateUsers(val bool) *SearchActiveProjectsService {
	s.corporateUsers = &val
	return s
}

func (s *SearchActiveProjectsService) SetMarketingMobileNumber(val bool) *SearchActiveProjectsService {
	s.marketingMobileNumber = &val
	return s
}

func (s *SearchActiveProjectsService) SetSanctionDetails(val bool) *SearchActiveProjectsService {
	s.sanctionDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetLimitedAccount(val bool) *SearchActiveProjectsService {
	s.limitedAccount = &val
	return s
}

func (s *SearchActiveProjectsService) SetEquipmentGroupDetails(val bool) *SearchActiveProjectsService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserDetails(val bool) *SearchActiveProjectsService {
	s.userDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetLocationDetails(val bool) *SearchActiveProjectsService {
	s.locationDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetNdaSignatureDetails(val bool) *SearchActiveProjectsService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetProjectCollaborationDetails(val bool) *SearchActiveProjectsService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserAvatar(val bool) *SearchActiveProjectsService {
	s.userAvatar = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserCountryDetails(val bool) *SearchActiveProjectsService {
	s.userCountryDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserProfileDescription(val bool) *SearchActiveProjectsService {
	s.userProfileDescription = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserDisplayInfo(val bool) *SearchActiveProjectsService {
	s.userDisplayInfo = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserJobs(val bool) *SearchActiveProjectsService {
	s.userJobs = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserBalanceDetails(val bool) *SearchActiveProjectsService {
	s.userBalanceDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserQualificationDetails(val bool) *SearchActiveProjectsService {
	s.userQualificationDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserMembershipDetails(val bool) *SearchActiveProjectsService {
	s.userMembershipDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserFinancialDetails(val bool) *SearchActiveProjectsService {
	s.userFinancialDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserLocationDetails(val bool) *SearchActiveProjectsService {
	s.userLocationDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserPortfolioDetails(val bool) *SearchActiveProjectsService {
	s.userPortfolioDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserPreferredDetails(val bool) *SearchActiveProjectsService {
	s.userPreferredDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserBadgeDetails(val bool) *SearchActiveProjectsService {
	s.userBadgeDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserStatus(val bool) *SearchActiveProjectsService {
	s.userStatus = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserReputation(val bool) *SearchActiveProjectsService {
	s.userReputation = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserEmployerReputation(val bool) *SearchActiveProjectsService {
	s.userEmployerReputation = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserReputationExtra(val bool) *SearchActiveProjectsService {
	s.userReputationExtra = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserEmployerReputationExtra(val bool) *SearchActiveProjectsService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserCoverImage(val bool) *SearchActiveProjectsService {
	s.userCoverImage = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserPastCoverImage(val bool) *SearchActiveProjectsService {
	s.userPastCoverImage = &val
	return s
}

func (s *SearchActiveProjectsService) SetUserRecommendations(val bool) *SearchActiveProjectsService {
	s.userRecommendations = &val
	return s
}

func (s *SearchActiveProjectsService) SetProjectTypes(projectTypes []ProjectType) *SearchActiveProjectsService {
	s.projectTypes = projectTypes
	return s
}

func (s *SearchActiveProjectsService) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *SearchActiveProjectsService {
	s.projectUpgrades = projectUpgrades
	return s
}

func (s *SearchActiveProjectsService) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *SearchActiveProjectsService {
	s.contestUpgrades = contestUpgrades
	return s
}

func (s *SearchActiveProjectsService) SetCompact(val bool) *SearchActiveProjectsService {
	s.compact = &val
	return s
}

func (s *SearchActiveProjectsService) SetLimit(val int) *SearchActiveProjectsService {
	s.limit = &val
	return s
}

func (s *SearchActiveProjectsService) SetOffset(val int) *SearchActiveProjectsService {
	s.offset = &val
	return s
}

func (s *SearchActiveProjectsService) SetReverseSort(val bool) *SearchActiveProjectsService {
	s.reverseSort = &val
	return s
}

func (s *SearchActiveProjectsService) SetFullDescription(val bool) *SearchActiveProjectsService {
	s.fullDescription = &val
	return s
}

func (s *SearchActiveProjectsService) SetJobDetails(val bool) *SearchActiveProjectsService {
	s.jobDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetUpgradeDetails(val bool) *SearchActiveProjectsService {
	s.upgradeDetails = &val
	return s
}

func (s *SearchActiveProjectsService) SetQuery(val string) *SearchActiveProjectsService {
	s.query = &val
	return s
}

func (s *SearchActiveProjectsService) SetMinAvgPrice(val float64) *SearchActiveProjectsService {
	s.minAvgPrice = &val
	return s
}

func (s *SearchActiveProjectsService) SetMaxAvgPrice(val float64) *SearchActiveProjectsService {
	s.maxAvgPrice = &val
	return s
}

func (s *SearchActiveProjectsService) SetMinAvgHourlyRate(val float64) *SearchActiveProjectsService {
	s.minAvgHourlyRate = &val
	return s
}

func (s *SearchActiveProjectsService) SetMaxAvgHourlyRate(val float64) *SearchActiveProjectsService {
	s.maxAvgHourlyRate = &val
	return s
}

func (s *SearchActiveProjectsService) SetMinPrice(val float64) *SearchActiveProjectsService {
	s.minPrice = &val
	return s
}

func (s *SearchActiveProjectsService) SetMaxPrice(val float64) *SearchActiveProjectsService {
	s.maxPrice = &val
	return s
}

func (s *SearchActiveProjectsService) SetMinHourlyRate(val float64) *SearchActiveProjectsService {
	s.minHourlyRate = &val
	return s
}

func (s *SearchActiveProjectsService) SetMaxHourlyRate(val float64) *SearchActiveProjectsService {
	s.maxHourlyRate = &val
	return s
}

func (s *SearchActiveProjectsService) SetJobs(jobs []int32) *SearchActiveProjectsService {
	s.jobs = jobs
	return s
}

func (s *SearchActiveProjectsService) SetCountries(countries []string) *SearchActiveProjectsService {
	s.countries = countries
	return s
}

func (s *SearchActiveProjectsService) SetLanguages(languages []string) *SearchActiveProjectsService {
	s.languages = languages
	return s
}

func (s *SearchActiveProjectsService) SetLatitude(val float64) *SearchActiveProjectsService {
	s.latitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetLongitude(val float64) *SearchActiveProjectsService {
	s.longitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetFromTime(val int64) *SearchActiveProjectsService {
	s.fromTime = &val
	return s
}

func (s *SearchActiveProjectsService) SetToTime(val int64) *SearchActiveProjectsService {
	s.toTime = &val
	return s
}

func (s *SearchActiveProjectsService) SetSortField(val SortField) *SearchActiveProjectsService {
	s.sortField = &val
	return s
}

func (s *SearchActiveProjectsService) SetProjectIDs(projectIDs []int64) *SearchActiveProjectsService {
	s.projectIDs = projectIDs
	return s
}

func (s *SearchActiveProjectsService) SetTopRightLatitude(val float64) *SearchActiveProjectsService {
	s.topRightLatitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetTopRightLongitude(val float64) *SearchActiveProjectsService {
	s.topRightLongitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetBottomLeftLatitude(val float64) *SearchActiveProjectsService {
	s.bottomLeftLatitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetBottomLeftLongitude(val float64) *SearchActiveProjectsService {
	s.bottomLeftLongitude = &val
	return s
}

func (s *SearchActiveProjectsService) SetOrSearchQuery(val string) *SearchActiveProjectsService {
	s.orSearchQuery = &val
	return s
}

func (s *SearchActiveProjectsService) SetHighlightPreTags(val string) *SearchActiveProjectsService {
	s.highlightPreTags = &val
	return s
}

func (s *SearchActiveProjectsService) SetHighlightPostTags(val string) *SearchActiveProjectsService {
	s.highlightPostTags = &val
	return s
}

package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type projectsSearchActiveService struct {
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

// Do perform GET request on endpoint "projects/0.1/projects/active/""
func (s *projectsSearchActiveService) Do(ctx context.Context) (*BaseResponse, error) {
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
	res := &BaseResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *projectsSearchActiveService) SetUserResponsiveness(val bool) *projectsSearchActiveService {
	s.userResponsiveness = &val
	return s
}

func (s *projectsSearchActiveService) SetCorporateUsers(val bool) *projectsSearchActiveService {
	s.corporateUsers = &val
	return s
}

func (s *projectsSearchActiveService) SetMarketingMobileNumber(val bool) *projectsSearchActiveService {
	s.marketingMobileNumber = &val
	return s
}

func (s *projectsSearchActiveService) SetSanctionDetails(val bool) *projectsSearchActiveService {
	s.sanctionDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetLimitedAccount(val bool) *projectsSearchActiveService {
	s.limitedAccount = &val
	return s
}

func (s *projectsSearchActiveService) SetEquipmentGroupDetails(val bool) *projectsSearchActiveService {
	s.equipmentGroupDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserDetails(val bool) *projectsSearchActiveService {
	s.userDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetLocationDetails(val bool) *projectsSearchActiveService {
	s.locationDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetNdaSignatureDetails(val bool) *projectsSearchActiveService {
	s.ndaSignatureDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetProjectCollaborationDetails(val bool) *projectsSearchActiveService {
	s.projectCollaborationDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserAvatar(val bool) *projectsSearchActiveService {
	s.userAvatar = &val
	return s
}

func (s *projectsSearchActiveService) SetUserCountryDetails(val bool) *projectsSearchActiveService {
	s.userCountryDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserProfileDescription(val bool) *projectsSearchActiveService {
	s.userProfileDescription = &val
	return s
}

func (s *projectsSearchActiveService) SetUserDisplayInfo(val bool) *projectsSearchActiveService {
	s.userDisplayInfo = &val
	return s
}

func (s *projectsSearchActiveService) SetUserJobs(val bool) *projectsSearchActiveService {
	s.userJobs = &val
	return s
}

func (s *projectsSearchActiveService) SetUserBalanceDetails(val bool) *projectsSearchActiveService {
	s.userBalanceDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserQualificationDetails(val bool) *projectsSearchActiveService {
	s.userQualificationDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserMembershipDetails(val bool) *projectsSearchActiveService {
	s.userMembershipDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserFinancialDetails(val bool) *projectsSearchActiveService {
	s.userFinancialDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserLocationDetails(val bool) *projectsSearchActiveService {
	s.userLocationDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserPortfolioDetails(val bool) *projectsSearchActiveService {
	s.userPortfolioDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserPreferredDetails(val bool) *projectsSearchActiveService {
	s.userPreferredDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserBadgeDetails(val bool) *projectsSearchActiveService {
	s.userBadgeDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUserStatus(val bool) *projectsSearchActiveService {
	s.userStatus = &val
	return s
}

func (s *projectsSearchActiveService) SetUserReputation(val bool) *projectsSearchActiveService {
	s.userReputation = &val
	return s
}

func (s *projectsSearchActiveService) SetUserEmployerReputation(val bool) *projectsSearchActiveService {
	s.userEmployerReputation = &val
	return s
}

func (s *projectsSearchActiveService) SetUserReputationExtra(val bool) *projectsSearchActiveService {
	s.userReputationExtra = &val
	return s
}

func (s *projectsSearchActiveService) SetUserEmployerReputationExtra(val bool) *projectsSearchActiveService {
	s.userEmployerReputationExtra = &val
	return s
}

func (s *projectsSearchActiveService) SetUserCoverImage(val bool) *projectsSearchActiveService {
	s.userCoverImage = &val
	return s
}

func (s *projectsSearchActiveService) SetUserPastCoverImage(val bool) *projectsSearchActiveService {
	s.userPastCoverImage = &val
	return s
}

func (s *projectsSearchActiveService) SetUserRecommendations(val bool) *projectsSearchActiveService {
	s.userRecommendations = &val
	return s
}

func (s *projectsSearchActiveService) SetProjectTypes(projectTypes []ProjectType) *projectsSearchActiveService {
	s.projectTypes = projectTypes
	return s
}

func (s *projectsSearchActiveService) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *projectsSearchActiveService {
	s.projectUpgrades = projectUpgrades
	return s
}

func (s *projectsSearchActiveService) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *projectsSearchActiveService {
	s.contestUpgrades = contestUpgrades
	return s
}

func (s *projectsSearchActiveService) SetCompact(val bool) *projectsSearchActiveService {
	s.compact = &val
	return s
}

func (s *projectsSearchActiveService) SetLimit(val int) *projectsSearchActiveService {
	s.limit = &val
	return s
}

func (s *projectsSearchActiveService) SetOffset(val int) *projectsSearchActiveService {
	s.offset = &val
	return s
}

func (s *projectsSearchActiveService) SetReverseSort(val bool) *projectsSearchActiveService {
	s.reverseSort = &val
	return s
}

func (s *projectsSearchActiveService) SetFullDescription(val bool) *projectsSearchActiveService {
	s.fullDescription = &val
	return s
}

func (s *projectsSearchActiveService) SetJobDetails(val bool) *projectsSearchActiveService {
	s.jobDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetUpgradeDetails(val bool) *projectsSearchActiveService {
	s.upgradeDetails = &val
	return s
}

func (s *projectsSearchActiveService) SetQuery(val string) *projectsSearchActiveService {
	s.query = &val
	return s
}

func (s *projectsSearchActiveService) SetMinAvgPrice(val float64) *projectsSearchActiveService {
	s.minAvgPrice = &val
	return s
}

func (s *projectsSearchActiveService) SetMaxAvgPrice(val float64) *projectsSearchActiveService {
	s.maxAvgPrice = &val
	return s
}

func (s *projectsSearchActiveService) SetMinAvgHourlyRate(val float64) *projectsSearchActiveService {
	s.minAvgHourlyRate = &val
	return s
}

func (s *projectsSearchActiveService) SetMaxAvgHourlyRate(val float64) *projectsSearchActiveService {
	s.maxAvgHourlyRate = &val
	return s
}

func (s *projectsSearchActiveService) SetMinPrice(val float64) *projectsSearchActiveService {
	s.minPrice = &val
	return s
}

func (s *projectsSearchActiveService) SetMaxPrice(val float64) *projectsSearchActiveService {
	s.maxPrice = &val
	return s
}

func (s *projectsSearchActiveService) SetMinHourlyRate(val float64) *projectsSearchActiveService {
	s.minHourlyRate = &val
	return s
}

func (s *projectsSearchActiveService) SetMaxHourlyRate(val float64) *projectsSearchActiveService {
	s.maxHourlyRate = &val
	return s
}

func (s *projectsSearchActiveService) SetJobs(jobs []int32) *projectsSearchActiveService {
	s.jobs = jobs
	return s
}

func (s *projectsSearchActiveService) SetCountries(countries []string) *projectsSearchActiveService {
	s.countries = countries
	return s
}

func (s *projectsSearchActiveService) SetLanguages(languages []string) *projectsSearchActiveService {
	s.languages = languages
	return s
}

func (s *projectsSearchActiveService) SetLatitude(val float64) *projectsSearchActiveService {
	s.latitude = &val
	return s
}

func (s *projectsSearchActiveService) SetLongitude(val float64) *projectsSearchActiveService {
	s.longitude = &val
	return s
}

func (s *projectsSearchActiveService) SetFromTime(val int64) *projectsSearchActiveService {
	s.fromTime = &val
	return s
}

func (s *projectsSearchActiveService) SetToTime(val int64) *projectsSearchActiveService {
	s.toTime = &val
	return s
}

func (s *projectsSearchActiveService) SetSortField(val SortField) *projectsSearchActiveService {
	s.sortField = &val
	return s
}

func (s *projectsSearchActiveService) SetProjectIDs(projectIDs []int64) *projectsSearchActiveService {
	s.projectIDs = projectIDs
	return s
}

func (s *projectsSearchActiveService) SetTopRightLatitude(val float64) *projectsSearchActiveService {
	s.topRightLatitude = &val
	return s
}

func (s *projectsSearchActiveService) SetTopRightLongitude(val float64) *projectsSearchActiveService {
	s.topRightLongitude = &val
	return s
}

func (s *projectsSearchActiveService) SetBottomLeftLatitude(val float64) *projectsSearchActiveService {
	s.bottomLeftLatitude = &val
	return s
}

func (s *projectsSearchActiveService) SetBottomLeftLongitude(val float64) *projectsSearchActiveService {
	s.bottomLeftLongitude = &val
	return s
}

func (s *projectsSearchActiveService) SetOrSearchQuery(val string) *projectsSearchActiveService {
	s.orSearchQuery = &val
	return s
}

func (s *projectsSearchActiveService) SetHighlightPreTags(val string) *projectsSearchActiveService {
	s.highlightPreTags = &val
	return s
}

func (s *projectsSearchActiveService) SetHighlightPostTags(val string) *projectsSearchActiveService {
	s.highlightPostTags = &val
	return s
}

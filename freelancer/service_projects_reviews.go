package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type listReviews struct {
	client                      *Client
	projects                    []int64
	fromUsers                   []int64
	toUsers                     []int64
	contests                    []int64
	reviewTypes                 []ReviewType
	jobIds                      []int64
	completionStatuses          []CompletionStatus
	fromTime                    *int64
	toTime                      *int64
	reviewStatus                []string
	projectDetails              *bool
	ratings                     *bool
	reviewCount                 *bool
	role                        *RoleType
	contestDetails              *bool
	userDetails                 *bool
	projectFullDescription      *bool
	projectUpgradeDetails       *bool
	projectJobDetails           *bool
	projectSelectedBids         *bool
	projectQualificationDetails *bool
	projectAttachmentDetails    *bool
	projectHiremeDetails        *bool
	contestJobDetails           *bool
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

// TODO: refine with typed response
func (s *listReviews) Do(ctx context.Context) (*RawResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: string(PROJECTS_REVIEWS),
	}

	for _, val := range s.projects {
		r.addParam("projects[]", val)
	}
	for _, val := range s.fromUsers {
		r.addParam("from_users[]", val)
	}
	for _, val := range s.toUsers {
		r.addParam("to_users[]", val)
	}
	for _, val := range s.contests {
		r.addParam("contests[]", val)
	}
	for _, val := range s.reviewTypes {
		r.addParam("review_types[]", val)
	}
	for _, val := range s.jobIds {
		r.addParam("job_ids[]", val)
	}
	for _, val := range s.completionStatuses {
		r.addParam("completion_statuses[]", val)
	}
	if s.fromTime != nil {
		r.addParam("from_time", *s.fromTime)
	}
	if s.toTime != nil {
		r.addParam("to_time", *s.toTime)
	}
	for _, val := range s.reviewStatus {
		r.addParam("review_status[]", val)
	}
	if s.projectDetails != nil {
		r.addParam("project_details", *s.projectDetails)
	}
	if s.ratings != nil {
		r.addParam("ratings", *s.ratings)
	}
	if s.reviewCount != nil {
		r.addParam("review_count", *s.reviewCount)
	}
	if s.role != nil {
		r.addParam("role", *s.role)
	}
	if s.contestDetails != nil {
		r.addParam("contest_details", *s.contestDetails)
	}
	if s.userDetails != nil {
		r.addParam("user_details", *s.userDetails)
	}
	if s.projectFullDescription != nil {
		r.addParam("project_full_description", *s.projectFullDescription)
	}
	if s.projectUpgradeDetails != nil {
		r.addParam("project_upgrade_details", *s.projectUpgradeDetails)
	}
	if s.projectJobDetails != nil {
		r.addParam("project_job_details", *s.projectJobDetails)
	}
	if s.projectSelectedBids != nil {
		r.addParam("project_selected_bids", *s.projectSelectedBids)
	}
	if s.projectQualificationDetails != nil {
		r.addParam("project_qualification_details", *s.projectQualificationDetails)
	}
	if s.projectAttachmentDetails != nil {
		r.addParam("project_attachment_details", *s.projectAttachmentDetails)
	}
	if s.projectHiremeDetails != nil {
		r.addParam("project_hireme_details", *s.projectHiremeDetails)
	}
	if s.contestJobDetails != nil {
		r.addParam("contest_job_details", *s.contestJobDetails)
	}
	if s.userAvatar != nil {
		r.addParam("user_avatar", *s.userAvatar)
	}
	if s.userCountryDetails != nil {
		r.setParam("user_country_details", *s.userCountryDetails)
	}
	if s.userProfileDescription != nil {
		r.setParam("user_profile_description", *s.userProfileDescription)
	}
	if s.userDisplayInfo != nil {
		r.setParam("user_display_info", *s.userDisplayInfo)
	}
	if s.userJobs != nil {
		r.setParam("user_jobs", *s.userJobs)
	}
	if s.userBalanceDetails != nil {
		r.setParam("user_balance_details", *s.userBalanceDetails)
	}
	if s.userQualificationDetails != nil {
		r.setParam("user_qualification_details", *s.userQualificationDetails)
	}
	if s.userMembershipDetails != nil {
		r.setParam("user_membership_details", *s.userMembershipDetails)
	}
	if s.userFinancialDetails != nil {
		r.setParam("user_financial_details", *s.userFinancialDetails)
	}
	if s.userLocationDetails != nil {
		r.setParam("user_location_details", *s.userLocationDetails)
	}
	if s.userPortfolioDetails != nil {
		r.setParam("user_portfolio_details", *s.userPortfolioDetails)
	}
	if s.userPreferredDetails != nil {
		r.setParam("user_preferred_details", *s.userPreferredDetails)
	}
	if s.userBadgeDetails != nil {
		r.setParam("user_badge_details", *s.userBadgeDetails)
	}
	if s.userStatus != nil {
		r.setParam("user_status", *s.userStatus)
	}
	if s.userReputation != nil {
		r.setParam("user_reputation", *s.userReputation)
	}
	if s.userEmployerReputation != nil {
		r.setParam("user_employer_reputation", *s.userEmployerReputation)
	}
	if s.userReputationExtra != nil {
		r.setParam("user_reputation_extra", *s.userReputationExtra)
	}
	if s.userEmployerReputationExtra != nil {
		r.setParam("user_employer_reputation_extra", *s.userEmployerReputationExtra)
	}
	if s.userCoverImage != nil {
		r.setParam("user_cover_image", *s.userCoverImage)
	}
	if s.userPastCoverImage != nil {
		r.setParam("user_past_cover_image", *s.userPastCoverImage)
	}
	if s.userRecommendations != nil {
		r.setParam("user_recommendations", *s.userRecommendations)
	}
	if s.userResponsiveness != nil {
		r.setParam("user_responsiveness", *s.userResponsiveness)
	}
	if s.corporateUsers != nil {
		r.setParam("corporate_users", *s.corporateUsers)
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
	if s.equipmentGroupDetails != nil {
		r.setParam("equipment_group_details", *s.equipmentGroupDetails)
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
	return execute[*RawResponse](ctx, s.client, r)
}

// SetProjects sets the list of project IDs to filter reviews on.
// Equivalent to `projects[]`.
func (s *listReviews) SetProjects(vals []int64) *listReviews {
	s.projects = vals
	return s
}

// SetFromUsers sets user IDs who wrote the reviews.
// Equivalent to `from_users[]`.
func (s *listReviews) SetFromUsers(vals []int64) *listReviews {
	s.fromUsers = vals
	return s
}

// SetToUsers sets user IDs who are targets of the review.
// Equivalent to `to_users[]`.
func (s *listReviews) SetToUsers(vals []int64) *listReviews {
	s.toUsers = vals
	return s
}

// SetContests sets contest IDs to filter reviews on.
// Equivalent to `contests[]`.
func (s *listReviews) SetContests(vals []int64) *listReviews {
	s.contests = vals
	return s
}

// SetReviewTypes sets the review types to filter by (e.g., project, contest).
// Equivalent to `review_types[]`.
func (s *listReviews) SetReviewTypes(vals []ReviewType) *listReviews {
	s.reviewTypes = vals
	return s
}

// SetJobIds sets job IDs to filter reviews to only those with these jobs.
// Equivalent to `job_ids[]`.
func (s *listReviews) SetJobIds(vals []int64) *listReviews {
	s.jobIds = vals
	return s
}

// SetCompletionStatuses sets the completion status filter (e.g., complete/incomplete).
// Equivalent to `completion_statuses[]`.
func (s *listReviews) SetCompletionStatuses(vals []CompletionStatus) *listReviews {
	s.completionStatuses = vals
	return s
}

// SetFromTime sets the minimum timestamp to filter reviews created after (inclusive).
// Equivalent to `from_time`.
func (s *listReviews) SetFromTime(val int64) *listReviews {
	s.fromTime = &val
	return s
}

// SetToTime sets the maximum timestamp to filter reviews created before (exclusive).
// Equivalent to `to_time`.
func (s *listReviews) SetToTime(val int64) *listReviews {
	s.toTime = &val
	return s
}

// SetProjectDetails includes project info in the review response.
// Equivalent to `project_details`.
func (s *listReviews) SetProjectDetails(val bool) *listReviews {
	s.projectDetails = &val
	return s
}

// SetRatings includes individual ratings in the response.
// Equivalent to `ratings`.
func (s *listReviews) SetRatings(val bool) *listReviews {
	s.ratings = &val
	return s
}

// SetReviewCount includes the total review count.
// Equivalent to `reviews_count`.
func (s *listReviews) SetReviewCount(val bool) *listReviews {
	s.reviewCount = &val
	return s
}

// SetRole sets the role type (freelancer or employer).
// Equivalent to `role`.
func (s *listReviews) SetRole(val RoleType) *listReviews {
	s.role = &val
	return s
}

// SetUserDetails includes basic user information.
// Equivalent to `user_details`.
func (s *listReviews) SetUserDetails(val bool) *listReviews {
	s.userDetails = &val
	return s
}

// SetProjectFullDescription includes full project description.
// Equivalent to `project_full_description`.
func (s *listReviews) SetProjectFullDescription(val bool) *listReviews {
	s.projectFullDescription = &val
	return s
}

// SetProjectUpgradeDetails includes upgrade information.
// Equivalent to `project_upgrade_details`.
func (s *listReviews) SetProjectUpgradeDetails(val bool) *listReviews {
	s.projectUpgradeDetails = &val
	return s
}

// SetProjectJobDetails includes job information.
// Equivalent to `project_job_details`.
func (s *listReviews) SetProjectJobDetails(val bool) *listReviews {
	s.projectJobDetails = &val
	return s
}

// SetProjectSelectedBids includes selected or pending bids.
// Equivalent to `project_selected_bids`.
func (s *listReviews) SetProjectSelectedBids(val bool) *listReviews {
	s.projectSelectedBids = &val
	return s
}

// SetProjectQualificationDetails includes exams required to qualify.
// Equivalent to `project_qualification_details`.
func (s *listReviews) SetProjectQualificationDetails(val bool) *listReviews {
	s.projectQualificationDetails = &val
	return s
}

// SetProjectAttachmentDetails includes project attachments.
// Equivalent to `project_attachment_details`.
func (s *listReviews) SetProjectAttachmentDetails(val bool) *listReviews {
	s.projectAttachmentDetails = &val
	return s
}

// SetProjectHiremeDetails includes hireme offer details.
// Equivalent to `project_hireme_details`.
func (s *listReviews) SetProjectHiremeDetails(val bool) *listReviews {
	s.projectHiremeDetails = &val
	return s
}

// SetContestJobDetails includes job details for each contest.
// Equivalent to `contest_job_details`.
func (s *listReviews) SetContestJobDetails(val bool) *listReviews {
	s.contestJobDetails = &val
	return s
}

// SetUserAvatar includes user avatar.
// Equivalent to `user_avatar`.
func (s *listReviews) SetUserAvatar(val bool) *listReviews {
	s.userAvatar = &val
	return s
}

// SetUserCountryDetails includes user country info.
// Equivalent to `user_country_details`.
func (s *listReviews) SetUserCountryDetails(val bool) *listReviews {
	s.userCountryDetails = &val
	return s
}

// SetUserProfileDescription includes the user's profile blurb.
// Equivalent to `user_profile_description`.
func (s *listReviews) SetUserProfileDescription(val bool) *listReviews {
	s.userProfileDescription = &val
	return s
}

// SetUserDisplayInfo includes user display name.
// Equivalent to `user_display_info`.
func (s *listReviews) SetUserDisplayInfo(val bool) *listReviews {
	s.userDisplayInfo = &val
	return s
}

// SetUserJobs includes user's jobs.
// Equivalent to `user_jobs`.
func (s *listReviews) SetUserJobs(val bool) *listReviews {
	s.userJobs = &val
	return s
}

// SetUserBalanceDetails includes user's balance info.
// Equivalent to `user_balance_details`.
func (s *listReviews) SetUserBalanceDetails(val bool) *listReviews {
	s.userBalanceDetails = &val
	return s
}

// SetUserQualificationDetails includes user's passed exams.
// Equivalent to `user_qualification_details`.
func (s *listReviews) SetUserQualificationDetails(val bool) *listReviews {
	s.userQualificationDetails = &val
	return s
}

// SetUserMembershipDetails includes user's membership info.
// Equivalent to `user_membership_details`.
func (s *listReviews) SetUserMembershipDetails(val bool) *listReviews {
	s.userMembershipDetails = &val
	return s
}

// SetUserFinancialDetails includes user's financial data.
// Equivalent to `user_financial_details`.
func (s *listReviews) SetUserFinancialDetails(val bool) *listReviews {
	s.userFinancialDetails = &val
	return s
}

// SetUserLocationDetails includes user's location info.
// Equivalent to `user_location_details`.
func (s *listReviews) SetUserLocationDetails(val bool) *listReviews {
	s.userLocationDetails = &val
	return s
}

// SetUserPortfolioDetails includes user's portfolio.
// Equivalent to `user_portfolio_details`.
func (s *listReviews) SetUserPortfolioDetails(val bool) *listReviews {
	s.userPortfolioDetails = &val
	return s
}

// SetUserPreferredDetails includes preferred info.
// Equivalent to `user_preferred_details`.
func (s *listReviews) SetUserPreferredDetails(val bool) *listReviews {
	s.userPreferredDetails = &val
	return s
}

// SetUserBadgeDetails includes badges earned by the user.
// Equivalent to `user_badge_details`.
func (s *listReviews) SetUserBadgeDetails(val bool) *listReviews {
	s.userBadgeDetails = &val
	return s
}

// SetUserStatus includes user's status info.
// Equivalent to `user_status`.
func (s *listReviews) SetUserStatus(val bool) *listReviews {
	s.userStatus = &val
	return s
}

// SetUserReputation includes freelancer reputation.
// Equivalent to `user_reputation`.
func (s *listReviews) SetUserReputation(val bool) *listReviews {
	s.userReputation = &val
	return s
}

// SetUserEmployerReputation includes employer reputation.
// Equivalent to `user_employer_reputation`.
func (s *listReviews) SetUserEmployerReputation(val bool) *listReviews {
	s.userEmployerReputation = &val
	return s
}

// SetUserReputationExtra includes full freelancer reputation.
// Equivalent to `user_reputation_extra`.
func (s *listReviews) SetUserReputationExtra(val bool) *listReviews {
	s.userReputationExtra = &val
	return s
}

// SetUserEmployerReputationExtra includes full employer reputation.
// Equivalent to `user_employer_reputation_extra`.
func (s *listReviews) SetUserEmployerReputationExtra(val bool) *listReviews {
	s.userEmployerReputationExtra = &val
	return s
}

// SetUserCoverImage includes user's current profile image.
// Equivalent to `user_cover_image`.
func (s *listReviews) SetUserCoverImage(val bool) *listReviews {
	s.userCoverImage = &val
	return s
}

// SetUserPastCoverImage includes previous profile images.
// Equivalent to `user_past_cover_images`.
func (s *listReviews) SetUserPastCoverImage(val bool) *listReviews {
	s.userPastCoverImage = &val
	return s
}

// SetUserRecommendations includes user recommendations count.
// Equivalent to `user_recommendations`.
func (s *listReviews) SetUserRecommendations(val bool) *listReviews {
	s.userRecommendations = &val
	return s
}

// SetUserResponsiveness includes user responsiveness score.
// Equivalent to `user_responsiveness`.
func (s *listReviews) SetUserResponsiveness(val bool) *listReviews {
	s.userResponsiveness = &val
	return s
}

// SetCorporateUsers includes corporate accounts info.
// Equivalent to `corporate_users`.
func (s *listReviews) SetCorporateUsers(val bool) *listReviews {
	s.corporateUsers = &val
	return s
}

// SetMarketingMobileNumber includes user's recruiter mobile number.
// Equivalent to `marketing_mobile_number`.
func (s *listReviews) SetMarketingMobileNumber(val bool) *listReviews {
	s.marketingMobileNumber = &val
	return s
}

// SetSanctionDetails includes user sanction timestamp.
// Equivalent to `sanction_details`.
func (s *listReviews) SetSanctionDetails(val bool) *listReviews {
	s.sanctionDetails = &val
	return s
}

// SetLimitedAccount includes limited account status.
// Equivalent to `limited_account`.
func (s *listReviews) SetLimitedAccount(val bool) *listReviews {
	s.limitedAccount = &val
	return s
}

// SetEquipmentGroupDetails includes equipment group info.
// Equivalent to `equipment_group_details`.
func (s *listReviews) SetEquipmentGroupDetails(val bool) *listReviews {
	s.equipmentGroupDetails = &val
	return s
}

// SetLimit sets the maximum number of results to return.
// Equivalent to `limit`.
func (s *listReviews) SetLimit(val int) *listReviews {
	s.limit = &val
	return s
}

// SetOffset sets the number of results to skip (pagination).
// Equivalent to `offset`.
func (s *listReviews) SetOffset(val int) *listReviews {
	s.offset = &val
	return s
}

// SetCompact strips null and empty values from the response.
// Equivalent to `compact`.
func (s *listReviews) SetCompact(val bool) *listReviews {
	s.compact = &val
	return s
}

type createReview struct {
	client *Client
}

type CreateReviewBody struct {
	ProjectID  int64      `json:"project_id"`
	ToUserID   int64      `json:"to_user_id"`
	FromUserID int64      `json:"from_user_id"`
	ReviewType ReviewType `json:"review_type"`
	Comment    string     `json:"comment"`
	Role       RoleType   `json:"role"`
}

func (s *createReview) Do(ctx context.Context, b CreateReviewBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: string(PROJECTS_REVIEWS),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

type actionReview struct {
	client   *Client
	reviewID int64
}

type ReviewActionBody struct {
	Action     ReviewAction `json:"action"`
	ReviewType ReviewType   `json:"review_type"`
}

func (s *actionReview) Do(ctx context.Context, b ReviewActionBody) (*RawResponse, error) {
	m, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("%s/%s", string(PROJECTS_REVIEWS), strconv.FormatInt(s.reviewID, 10)),
		body:     bytes.NewBuffer(m),
	}
	return execute[*RawResponse](ctx, s.client, r)
}

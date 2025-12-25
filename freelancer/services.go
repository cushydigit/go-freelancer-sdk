package freelancer

type Services struct {
	client   *Client
	Projects *ProjectsService
	Users    *UsersService
	Common   *CommonService
}

func newServices(c *Client) *Services {

	s := &Services{client: c}
	// init projects services
	s.Projects = &ProjectsService{client: c}
	s.Projects.Collaborations = &CollaborationsService{client: c}
	s.Projects.Services = &ServicesService{client: c}
	s.Projects.Bids = &BidsService{client: c}
	s.Projects.Bids.EditRequests = &BidEditRequestService{client: c}
	s.Projects.Bids.Ratings = &BidRatingsService{client: c}
	s.Projects.Milestones = &MilestonesService{client: c}
	s.Projects.Milestones.Requests = &MilestoneRequestsService{client: c}
	s.Projects.Reviews = &ReviewsService{client: c}
	s.Projects.Jobs = &JobsService{client: c}
	s.Projects.Extras = &ProjectsExtrasService{client: c}
	s.Projects.Extras.Budgets = &BudgetsService{client: c}
	s.Projects.Extras.Categories = &CategoriesService{client: c}
	s.Projects.Extras.Currencies = &CurrenciesService{client: c}
	s.Projects.Extras.ExpertGuarantees = &ExpertGuaranteesService{client: c}

	// init users services
	s.Users = &UsersService{client: c}
	s.Users.Self = &SelfService{client: c}
	s.Users.SelfJob = &SelfJobsService{client: c}
	s.Users.Profiles = &ProfilesService{client: c}
	s.Users.Extras = &UsersExtrasService{client: c}

	// init common services
	s.Common = &CommonService{client: c}

	return s
}

type ProjectsService struct {
	client         *Client
	Collaborations *CollaborationsService
	Services       *ServicesService
	Bids           *BidsService
	Jobs           *JobsService
	Milestones     *MilestonesService
	Reviews        *ReviewsService
	Extras         *ProjectsExtrasService
}

type CollaborationsService struct{ client *Client }
type ServicesService struct{ client *Client }

type BidsService struct {
	client       *Client
	EditRequests *BidEditRequestService
	Ratings      *BidRatingsService
}

type BidEditRequestService struct{ client *Client }

type BidRatingsService struct{ client *Client }

type JobsService struct{ client *Client }

type MilestonesService struct {
	client   *Client
	Requests *MilestoneRequestsService
}

type ReviewsService struct{ client *Client }

type MilestoneRequestsService struct{ client *Client }

type ProjectsExtrasService struct {
	client           *Client
	ExpertGuarantees *ExpertGuaranteesService
	Budgets          *BudgetsService
	Currencies       *CurrenciesService
	Categories       *CategoriesService
}

type ExpertGuaranteesService struct{ client *Client }

type BudgetsService struct{ client *Client }

type CurrenciesService struct{ client *Client }

type CategoriesService struct{ client *Client }

type UsersService struct {
	client   *Client
	Self     *SelfService
	SelfJob  *SelfJobsService
	Profiles *ProfilesService
	Extras   *UsersExtrasService
}

type SelfService struct{ client *Client }

type SelfJobsService struct{ client *Client }

type ProfilesService struct{ client *Client }

type UsersExtrasService struct{ client *Client }

type CommonService struct{ client *Client }

// -------------------------------------------------------------------------------
// Projects - Projects
// -------------------------------------------------------------------------------

// The returned service prepares a `POST` request to the Projects endpoint
// `/projects/0.1/projects`. The request is executed when Do is called.
//
// The service creates a new project
func (s *ProjectsService) Create() *createProject {
	return &createProject{client: s.client}
}

// The returned service prepares a `POST` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}`. The request is executed when Do is called.
//
// The service performs an action on a project
func (s *ProjectsService) Action(projectID int64) *actionProject {
	return &actionProject{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects`. The request is executed when Do is called.
//
// The service retrieve information about multiple projects. Will be ordered by descending submit date (newest-to-oldest).
func (s *ProjectsService) List() *listProjects { return &listProjects{client: s.client} }

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/self`. The request is executed when Do is called.
//
// The service retrieve the logged in user's projects/contest they either created or participated in (by bidding or submitting an entry)
func (s *ProjectsService) ListSelf() *listSelfProjects { return &listSelfProjects{client: s.client} }

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}`. The request is executed when Do is called.
//
// The service retrieve information about a specific project. The full range of users projection options can be specified as part of this request by first setting the user_details parameter to true.
func (s *ProjectsService) GetByID(projectID int64) *getProjectByID {
	return &getProjectByID{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/active`. The request is executed when Do is called.
//
// The service searches for active projects matching the desired query
func (s *ProjectsService) SearchActive() *searchActiveProjects {
	return &searchActiveProjects{client: s.client}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/active`. The request is executed when Do is called.
//
// The service searches for active projects matching the desired query
func (s *ProjectsService) SearchAll() *searchAllProjects {
	return &searchAllProjects{client: s.client}
}

// The returned service prepares a `POST` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/invite`. The request is executed when Do is called.
//
// The service invites specific freelancers to bid on a project
func (s *ProjectsService) InviteFreelancer(projectID int64) *inviteFreelancer {
	return &inviteFreelancer{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/fees`. The request is executed when Do is called.
//
// The service returns the project fees for a given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
func (s *ProjectsService) ListUpgradesFees() *listUpgradeFees {
	return &listUpgradeFees{client: s.client}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/bids`. The request is executed when Do is called.
//
// The Service retrieve bids for a single project. Returns bids for a single project.
// Employers will see bids in a sorted order, which begins with sponsored bids, then by bid ranking. Freelancers will receive the bid list ordered by date.
// Note: This method is expensive to compute so it is recommended that reputation and user projection options are not set.
func (s *ProjectsService) ListBids(projectID int64) *listProjectBids {
	return &listProjectBids{client: s.client, projectID: projectID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids/fees`. The request is executed when Do is called.
//
// The service returns a list of bid upgrade fees for given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
func (s *ProjectsService) ListBidUpgradeFees() *listBidUpgradeFees {
	return &listBidUpgradeFees{client: s.client}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/bid_info`. The request is executed when Do is called.
//
// The service returns information for posting bids on project
func (s *ProjectsService) GetBidInfo(projectID int64) *getProjectBidInfo {
	return &getProjectBidInfo{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/milestones`. The request is executed when Do is called.
//
// The service retrieve a list of milestones on a project after e
func (s *ProjectsService) ListMilestones(projectID int64) *listProjectMilestones {
	return &listProjectMilestones{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/milestone_requests`. The request is executed when Do is called.
//
// The service retrieve a list of milestone requests by freelancer for a project.
func (s *ProjectsService) ListMilestoneRequests(projectID int64) *listProjectMilestoneRequests {
	return &listProjectMilestoneRequests{client: s.client, projectID: projectID}
}

// The returned service prepares a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/hourly_contract_info`. The request is executed when Do is called.
//
// The service fetch the hourly contract matching the desired query
func (s *ProjectsService) GetHourlyContractInfo() *getHourlyContractInfo {
	return &getHourlyContractInfo{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/ip_contract_info`. The request is executed when Do is called.
//
// The service fetch the IP contract matching for the project id.If you are an employer it will return all of the contracts, ELSE we will return contract details specific to the logged-in user.
func (s *ProjectsService) GetIPContractInfo(projectID int64) *getIPContractInfo {
	return &getIPContractInfo{client: s.client, projectID: projectID}
}

// The returned service prepare a `DELETE` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}`. The request is executed when Do is called.
//
// The service able to delete a project by id, Only projects that are in pending or rejected state may be deleted.
func (s *ProjectsService) Delete(projectID int64) *deleteProject {
	return &deleteProject{client: s.client, projectID: projectID}
}

// -------------------------------------------------------------------------------
// Projects - Collaborations
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/collaborations`. The request is executed when Do is called.
//
// The service returns a list of project collaboration data for a project
func (s *CollaborationsService) List(projectID int64) *listCollaborations {
	return &listCollaborations{client: s.client, projectID: projectID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/collaborations`. The request is executed when Do is called.
//
// The service creates a new project collaboration
func (s *CollaborationsService) Create(projectID int64) *createCollaboration {
	return &createCollaboration{client: s.client, projectID: projectID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions`. The request is executed when Do is called.
//
// The service performs an action on a collaboration
func (s *CollaborationsService) Action(projectID int64, collaborationID int) *actionCollaboration {
	return &actionCollaboration{client: s.client, projectID: projectID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/collaborations`. The request is executed when Do is called.
//
// The service retrieves a list of all collaboration data for a user
func (s *CollaborationsService) ListAll() *listAllCollaborations {
	return &listAllCollaborations{client: s.client}
}

// -------------------------------------------------------------------------------
// Projects - Services
// -------------------------------------------------------------------------------

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/services/{service_type}/{service_id}/order`. The request is executed when Do is called.
//
// The service orders one of the available services
func (s *ServicesService) Order(serviceID int, serviceType ServiceType) *orderService {
	return &orderService{client: s.client, serviceID: serviceID, serviceType: serviceType}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/services`. The request is executed when Do is called.
//
// The service retrieves a list of services
func (s *ServicesService) List() *listServices {
	return &listServices{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/services/active`. The request is executed when Do is called.
//
// The service retrieves active services
func (s *ServicesService) SearchActive() *searchActiveServices {
	return &searchActiveServices{client: s.client}
}

// -------------------------------------------------------------------------------
// Projects - Bids
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids`. The request is executed when Do is called.
//
// The service retrieves a list of bids that match the specified criteria
func (s *BidsService) List() *listBids {
	return &listBids{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}`. The request is executed when Do is called.
//
// The service retrieve information about a specific bid
func (s *BidsService) GetByID(bidID int) *getBidByID {
	return &getBidByID{client: s.client, bidID: bidID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/bids`. The request is executed when Do is called.
//
// The service creates a new bid
func (s *BidsService) Create() *createBid {
	return &createBid{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}`. The request is executed when Do is called.
//
// The service updates an existing bid
func (s *BidsService) Update(bidID int) *updateBid {
	return &updateBid{client: s.client, bidID: bidID}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}`. The request is executed when Do is called.
//
// The service performs an action on a bid
func (s *BidsService) Action(bidID int) *actionBid {
	return &actionBid{client: s.client, bidID: bidID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/time_tracking`. The request is executed when Do is called.
//
// The service returns a list of aggregate time tracking data for a bid
func (s *BidsService) GetTimeTracking(bidID int) *getTimeTracking {
	return &getTimeTracking{client: s.client, bidID: bidID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/time_tracking`. The request is executed when Do is called.
//
// The service creates a time tracking session for a specific bid
func (s *BidsService) CreateTimeTracking(bidID int) *createTimeTracking {
	return &createTimeTracking{client: s.client, bidID: bidID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/edit_requests`. The request is executed when Do is called.
//
// The service returns bid edit requests by bid id.
func (s *BidEditRequestService) Get(bidID int) *getBidEditRequest {
	return &getBidEditRequest{client: s.client, bidID: bidID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/edit_requests`. The request is executed when Do is called.
//
// The service create a bid edit request on a post accept awarded bid. With no pending bid edit request
func (s *BidEditRequestService) Create(bidID int) *createBidEditRequest {
	return &createBidEditRequest{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}`. The request is executed when Do is called.
//
// The service able employer to perform action on a PENDING bid edit request
func (s *BidEditRequestService) Action(bidID, bidEditRequestID int) *actionBidEditRequest {
	return &actionBidEditRequest{client: s.client, bidID: bidID, bidEditRequestID: bidEditRequestID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/bid_ratings`. The request is executed when Do is called.
//
// The service returns bid rating by bid id
func (s *BidRatingsService) Get(bidID int) *getBidRatings {
	return &getBidRatings{client: s.client, bidID: bidID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/bid_ratings`. The request is executed when Do is called.
//
// The service rate a bid (create a bid rating)
func (s *BidRatingsService) Create(bidID int) *createBidRating {
	return &createBidRating{client: s.client, bidID: bidID}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}`. The request is executed when Do is called.
//
// The service updates an existing bid rating
func (s *BidRatingsService) Update(bidID int) *updateBidRating {
	return &updateBidRating{client: s.client, bidID: bidID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/bid_ratings`. The request is executed when Do is called.
//
// The service returns bid rating for multiple bids
func (s *BidRatingsService) GetByListOfBids() *getBidRatingsByListOfBids {
	return &getBidRatingsByListOfBids{client: s.client}
}

// -------------------------------------------------------------------------------
// Projects - Jobs
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/jobs`. The request is executed when Do is called.
//
// The service returns a list of jobs
func (s *JobsService) List() *listJobs {
	return &listJobs{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/jobs/search`. The request is executed when Do is called.
//
// The service returns a list of jobs
// Note: This performs a sub-string search for all the parameters specified on the jobs
func (s *JobsService) Search() *searchJobs {
	return &searchJobs{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/job_bundles`. The request is executed when Do is called.
//
// The service returns a list of job bundles
// Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
func (s *JobsService) ListJobBundles() *listJobBundles {
	return &listJobBundles{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/job_bundle_categories`. The request is executed when Do is called.
//
// The service returns a list of job bundle categories
func (s *JobsService) ListJobBundleCategories() *listJobBundleCategories {
	return &listJobBundleCategories{client: s.client}
}

// -------------------------------------------------------------------------------
// Projects - Milestones
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/milestones`. The request is executed when Do is called.
//
// The service returns a list of milestones. Does not return un-awarded prepaid milestones
// Note that for the parameters, you cannot specify both users[]and bidders[]/project_owners[]
func (s *MilestonesService) List() *listMilestones {
	return &listMilestones{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/milestones/{milestone_id}`. The request is executed when Do is called.
//
// The service returns information about a specific milestone
func (s *MilestonesService) GetByID(milestoneID int) *getMilestoneByID {
	return &getMilestoneByID{client: s.client, milestoneID: milestoneID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/milestones`. The request is executed when Do is called.
//
// The service create a milestone
func (s *MilestonesService) Create() *createMilestone {
	return &createMilestone{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/milestones/{milestone_id}`. The request is executed when Do is called.
//
// The service performs an action on a milestone
func (s *MilestonesService) Action(milestoneID int) *actionMilestone {
	return &actionMilestone{client: s.client, milestoneID: milestoneID}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/milestone_requests`. The request is executed when Do is called.
//
// The service returns a list of milestone requests
func (s *MilestoneRequestsService) List() *listMilestoneRequests {
	return &listMilestoneRequests{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/milestone_requests/{milestone_request_id}`. The request is executed when Do is called.
//
// The service returns information about a specific milestone request
func (s *MilestoneRequestsService) GetByID(milestoneRequestID int) *getMilestoneRequestByID {
	return &getMilestoneRequestByID{client: s.client, milestoneRequestID: milestoneRequestID}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/milestone_requests`. The request is executed when Do is called.
//
// The service creates a milestone request
func (s *MilestoneRequestsService) Create() *createMilestoneRequest {
	return &createMilestoneRequest{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/milestone_requests/{milestone_request_id}`. The request is executed when Do is called.
//
// The service performs an action on a milestone request
func (s *MilestoneRequestsService) Action(milestoneRequestID int) *actionMilestoneRequest {
	return &actionMilestoneRequest{client: s.client, milestoneRequestID: milestoneRequestID}
}

// -------------------------------------------------------------------------------
// Projects - Reviews
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/reviews`. The request is executed when Do is called.
//
// The service returns a list of project reviews.
func (s *ReviewsService) List() *listReviews {
	return &listReviews{client: s.client}
}

// The returned service prepare a `POST` request to the Projects endpoint
// `/projects/0.1/reviews`. The request is executed when Do is called.
//
// The service creates a new review of a user
func (s *ReviewsService) Create() *createReview {
	return &createReview{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/reviews/{review_id}`. The request is executed when Do is called.
//
// The service performs an action on a review
// Note that Reviews are uniquely identified by a combination of review id and review type.
func (s *ReviewsService) Action(reviewID int64) *actionReview {
	return &actionReview{client: s.client, reviewID: reviewID}
}

// -------------------------------------------------------------------------------
// Projects - Extras
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/expert_guarantees`. The request is executed when Do is called.
//
// The service returns a list of expert guarantees
func (s *ExpertGuaranteesService) List() *listExpertGuarantees {
	return &listExpertGuarantees{client: s.client}
}

// The returned service prepare a `PUT` request to the Projects endpoint
// `/projects/0.1/expert_guarantees/{expert_guarantee_id}`. The request is executed when Do is called.
//
// The service performs an action on a expert guarantee
func (s *ExpertGuaranteesService) Action() *actionExpertGuarantee {
	return &actionExpertGuarantee{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/budgets`. The request is executed when Do is called.
//
// The service returns a list of budgets with the specified currencies. currency_codes and currency_ids are incompatible with each other
func (s *BudgetsService) List() *listBudgets {
	return &listBudgets{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/currencies`. The request is executed when Do is called.
//
// The service returns a list of currencies, `currency_codes` and `currency_ids` are incompatible with each other
func (s *CurrenciesService) List() *listCurrencies {
	return &listCurrencies{client: s.client}
}

// The returned service prepare a `GET` request to the Projects endpoint
// `/projects/0.1/categories`. The request is executed when Do is called.
//
// The service returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
func (s *CategoriesService) List() *listCategories {
	return &listCategories{client: s.client}
}

// -------------------------------------------------------------------------------
// Common
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Users - Users
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/users`. The request is executed when Do is called.
//
// The service returns a list of users
func (s *UsersService) List() *listUsers {
	return &listUsers{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/users/{user_id}`. The request is executed when Do is called.
//
// The service returns information about a specific user
func (s *UsersService) GetByID(userID int64) *getUserByID {
	return &getUserByID{client: s.client, userID: userID}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/users/directory`. The request is executed when Do is called.
//
// The service returns a list of freelancers.
func (s *UsersService) SearchFreelancer() *searchFreelancers {
	return &searchFreelancers{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Authenticated (self)
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/self`. The request is executed when Do is called.
//
// The service returns information about current user
func (s *SelfService) Get() *getInfo {
	return &getInfo{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/self/devices`. The request is executed when Do is called.
//
// The service returns a list of current user's recent logged in devices
func (s *SelfService) ListDevices() *listDevices {
	return &listDevices{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Authenticated Jobs (self)
// -------------------------------------------------------------------------------

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service adds a list of jobs to the job list of current User
func (s *SelfJobsService) Add() *addJobs {
	return &addJobs{client: s.client}
}

// The returned service prepare a `PUT` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service sets a list of jobs to the job list of current User
func (s *SelfJobsService) Update() *updateJobs {
	return &updateJobs{client: s.client}
}

// The returned service prepare a `DELETE` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service removes a list of jobs from the job list of current User
func (s *SelfJobsService) Delete() *deleteJobs {
	return &deleteJobs{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Profiles
// -------------------------------------------------------------------------------

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service creates a new profile for a user
func (s *ProfilesService) Create() *createProfile {
	return &createProfile{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service gets the profiles for a list of users
func (s *ProfilesService) Get() *getProfile {
	return &getProfile{client: s.client}
}

// The returned service prepare a `PUT` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service updates a profile
func (s *ProfilesService) Update() *updateProfile {
	return &updateProfile{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Extras
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/reputations`. The request is executed when Do is called.
//
// The service gets the reputations for a list of users
func (s *UsersExtrasService) ListReputations() *listReputations {
	return &listReputations{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/portfolios`. The request is executed when Do is called.
//
// The service gets the portfolios for a list of users
func (s *UsersExtrasService) ListPortfolios() *listPortfolios {
	return &listPortfolios{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/enterprises`. The request is executed when Do is called.
//
// The service returns a list of enterprises
func (s *UsersExtrasService) ListEnterprises() *listEnterprises {
	return &listEnterprises{client: s.client}
}

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/violation_reports`. The request is executed when Do is called.
//
// The service creates a new violation
func (s *UsersExtrasService) CreateViolation() *createViolation {
	return &createViolation{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/pools`. The request is executed when Do is called.
//
// The service returns a list of pools belonging to the current user
func (s *UsersExtrasService) ListPools() *listPools {
	return &listPools{client: s.client}
}

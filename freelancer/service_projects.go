package freelancer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/query"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// PROJECTS
// --------------------------------------

// Create a new project
// It maps to the `POST` `/projects/0.1/projects` endpoint.
func (s *ProjectsService) Create(ctx context.Context, b rr.CreateProjectBody) (*rr.CreateProjectResponse, error) {
	p := endpoints.Projects
	return execute[*rr.CreateProjectResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Perform an action on a project
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Action(ctx context.Context, projectID int64, action rr.ActionProjectBody) (*rr.RawResponse, error) {

	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, action)
}

// Returns information about multiple projects. Will be ordered by descending submit date (newest-to-oldest).
// it maps to the `GET` `/projects/0.1/projects` endpoint
func (s *ProjectsService) List(ctx context.Context, opts *rr.ListProjectsOptions) (*rr.ListProjectsResponse, error) {
	p := endpoints.Projects
	q := query.Values(opts)
	return execute[*rr.ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Returns the logged in user’s projects/contests they either created or participated in (by bidding or submitting an entry).
// it maps to the `GET` `/projects/0.1/self` endpoint
func (s *ProjectsService) ListSelf(ctx context.Context, opts *rr.ListSelfProjectsOptions) (*rr.ListProjectsResponse, error) {
	p := endpoints.ProjectsSelf
	q := query.Values(opts)
	return execute[*rr.ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Get information about a specific project. The full range of users projection options can be specified as part of this request by first setting theuser_detailsparameter to true.
// It maps to the `GET` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Get(ctx context.Context, projectID int64, opts *rr.GetProjectOptions) (*rr.GetProjectResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*rr.GetProjectResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Searches for active projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/active` endpoint
func (s *ProjectsService) SearchActive(ctx context.Context, opts *rr.SearchActiveProjectsOptions) (*rr.ListProjectsResponse, error) {
	p := endpoints.ProjectsActive
	q := query.Values(opts)
	return execute[*rr.ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Searches for all projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/all` endpoint
func (s *ProjectsService) SearchAll(ctx context.Context, opts *rr.SearchAllProjectsOptions) (*rr.ListProjectsResponse, error) {
	p := endpoints.ProjectsAll
	q := query.Values(opts)
	return execute[*rr.ListProjectsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Invites specific freelancers to bid on a project.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/invite` endpoint
func (s *ProjectsService) InviteFreelancer(ctx context.Context, projectID int64, b rr.InviteFreelancersBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/invite", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: Refine the typed response with

// Returns the project upgrade fees for a given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
// It maps to the `GET` `/projects/0.1/projects/fees` endpoint
func (s *ProjectsService) ListUpgradesFees(ctx context.Context, opts *rr.ListUpgradesFeesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsFees
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: Refine the typed response with

// Returns bids for a single project. Employers will see bids in a sorted order, which begins with sponsored bids, then by bid ranking. Freelancers will receive the bid list ordered by date. Note: This method is expensive to compute so it is recommended that reputation and user projection options are not set.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids` endpoint
func (s *ProjectsService) ListBids(ctx context.Context, projectID int64, opts *rr.ListProjectBidsOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bids", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns information for posting bids on a project.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids_info` endpoint
func (s *ProjectsService) GetBidInfo(ctx context.Context, projectID int64) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bids_info", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Returns a list of milestones on a project. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/milestones` endpoint
func (s *ProjectsService) ListMilestones(ctx context.Context, projectID int64, opts *rr.ListProjectMilestonesOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/milestones", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of milestone requests by freelancers for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/milestone_requests` endpoint
func (s *ProjectsService) ListMilestoneRequests(ctx context.Context, projectID int64, opts *rr.ListProjectsMilestoneRequestsOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/milestone_requests", endpoints.Projects, projectID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Fetch the hourly contract matching the desired query.
// It maps to the `GET` `/projects/0.1/hourly_contract_info` endpoint
func (s *ProjectsService) GetHourlyContractInfo(ctx context.Context, opts *rr.GetHourlyContractInfoOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsHourlyContract
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response
// TODO: check the endpoint

// Fetch the IP contract matching for the project id. If you are an employer it will return all of the contracts, ELSE we will return contract details specific to the logged-in user.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/ip_contract_info` endpoint
func (s *ProjectsService) GetIPContractInfo(ctx context.Context, projectID int64) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/ip_contract_info", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Delete a project by id. Only projects that are in pending or rejected states may be deleted. This will close the project and remove its visibility.
// it maps to the `DELETE` `/projects/0.1/projects/{project_id}` endpoint
func (s *ProjectsService) Delete(ctx context.Context, projectID int64) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodDelete, p, nil, nil)
}

// --------------------------------------
// PROJECTS-COLLABORATIONS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of project collaboration data for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) List(ctx context.Context, projectID int64) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/collaborations", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// Creates a new project collaboration.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (s *CollaborationsService) Create(ctx context.Context, projectID int64, b rr.CreateCollaborationBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/collaborations", endpoints.Projects, projectID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// TODO: check the method

// Performs an action on a collaboration.
// it maps to the `PUT` `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` endpoint
func (s *CollaborationsService) Action(ctx context.Context, projectID int64, collaborationID int64, b rr.ActionCollaborationBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/collaborations/%d/actions", endpoints.Projects, projectID, collaborationID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// TODO: refine with typed response
// TODO: check the endpoint

// Returns a list of all collaboration data for a user.
// It maps toi the `GET` `/projects/0.1/collaborations` endpoint
func (s *CollaborationsService) ListAll(ctx context.Context) (*rr.RawResponse, error) {
	p := endpoints.ProjectsCollaborations
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// --------------------------------------
// PROJECTS-SERVICES
// --------------------------------------

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (s *ServicesService) Order(ctx context.Context, serviceID int64, serviceType rr.ServiceType) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%s/%d/order", endpoints.ProjectsServices, serviceType, serviceID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, nil)
}

// TODO: refine with typed response

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (s *ServicesService) List(ctx context.Context, opts *rr.ListServicesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsServices
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (s *ServicesService) SearchActive(ctx context.Context, opts *rr.SearchActiveServicesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsServicesActive
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// PROJECTS-BIDS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// It maps to the `GET` `/projects/0.1/bids` endpoint
func (s *BidsService) List(ctx context.Context, opts *rr.ListBidsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsBids
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of bids that match the specified criteria.
// it maps to the `GET` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Get(ctx context.Context, bidID int64, opts *rr.GetBidOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Creates a bid on a project. Accepts a JSON object in the style described in the Bid struct (with enums as strings, and objects as dictionaries).
// It maps to the `POST` `/projects/0.1/bids` endpoint
func (s *BidsService) Create(ctx context.Context, b rr.CreateBidBody) (*rr.RawResponse, error) {
	p := endpoints.ProjectsBids
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Action(ctx context.Context, bidID int64, b rr.ActionBidBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// Updates an existing bid on a project. An existing bids information (description,amount,milestone_percentage) can be updated by sending a JSON encoded Bid struct.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (s *BidsService) Update(ctx context.Context, bidID int64, b rr.UpdateBidBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsBids, bidID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// TODO: refine with typed response

// Returns a list of aggregate time tracking data for a bid.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (s *BidsService) GetTimeTracking(ctx context.Context, bidID int64, opts *rr.GetTimeTrackingOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Creates a time tracking session for a specific bid.
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (s *BidsService) CreateTimeTracking(ctx context.Context, bidID int64, b rr.CreateTimeTrackingBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/time_tracking", endpoints.ProjectsBids, bidID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Return bid edit requests by bid id.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/edit_requests` endpoint
// the original name was Get but it was renamed to List due to returning list of edit requests
func (s *BidEditRequestsService) List(ctx context.Context, bidID int64, opts *rr.ListBidEditRequestsOptions) (*rr.ListBidEditRequestsResponse, error) {
	p := fmt.Sprintf("%s/%d/edit_requests", endpoints.ProjectsBids, bidID)
	q := query.Values(opts)
	return execute[*rr.ListBidEditRequestsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request.
// It maps to the `POST` `/projects/0.1/bids/edit_requests` endpoint
func (s *BidEditRequestsService) Create(ctx context.Context, b rr.CreateBidEditRequestBody) (*rr.CreateBidEditRequestResponse, error) {
	p := endpoints.ProjectsBidEditRequests
	return execute[*rr.CreateBidEditRequestResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (s *BidEditRequestsService) Action(ctx context.Context, bidID, bidEditRequestID int64, b rr.ActionBidEditRequestBody) (*rr.ActionBidEditRequestResponse, error) {
	p := fmt.Sprintf("%s/%d/edit_requests/%d", endpoints.ProjectsBids, bidID, bidEditRequestID)
	return execute[*rr.ActionBidEditRequestResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// TODO: refine with typed response

// Fetch bid rating for a bid
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Get(ctx context.Context, bidID int64) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, nil, nil)
}

// TODO: refine with typed response

// Fetch bid ratings for multiple bids
// it maps to the `GET` `/projects/0.1/bid_ratings` endpoint
func (s *BidRatingsService) GetByListOfBids(ctx context.Context, opts *rr.GetByListOfBidsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsBidRatings
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Rates a bid (creates a bid rating)
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (s *BidRatingsService) Create(ctx context.Context, bidID int64, b rr.CreateBidRatingBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings", endpoints.ProjectsBids, bidID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Updates an existing bid rating
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` endpoint
func (s *BidRatingsService) Update(ctx context.Context, bidID int64, bidRatingID int64, b rr.UpdateBidRatingBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d/bid_ratings/%d", endpoints.ProjectsBids, bidID, bidRatingID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// --------------------------------------
// PROJECTS-JOBS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (s *JobsService) List(ctx context.Context, opts *rr.ListJobsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsJobs
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (s *JobsService) Search(ctx context.Context, opts *rr.SearchJobsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsJobsSearch
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (s *JobBundlesService) List(ctx context.Context, opts *rr.ListJobBundlesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsJobBundles
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (s *JobBundleCategoriesService) List(ctx context.Context, opts *rr.ListJobBundleCategoriesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsJobBundleCategories
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// --------------------------------------
// PROJECTS-MILESTONES
// --------------------------------------

// TODO: refine with typed response

// Returns a list of milestones. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/milestones` endpoint
func (s *MilestonesService) List(ctx context.Context, opts *rr.ListMilestonesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsMilestones
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns information about a specific milestone.
// It maps to the `GET` `/projects/0.1/milestones/{milestone_id}` endpoint
func (s *MilestonesService) GetByID(ctx context.Context, milestoneID int, opts *rr.GetMilestoneOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Post a review of a user.
// It maps to the `POST` `/projects/0.1/milestones` endpoint
func (s *MilestonesService) Create(ctx context.Context, b rr.CreateMilestoneBody) (*rr.RawResponse, error) {
	p := endpoints.ProjectsMilestones
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (s *MilestonesService) Action(ctx context.Context, milestoneID int, b rr.ActionMilestoneBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestones, milestoneID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// TODO: refine with typed response

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/milestone_requests` endpoint
func (s *MilestoneRequestsService) List(ctx context.Context, opts *rr.ListMilestonesRequestsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsMilestoneRequests
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// TODO: refine with typed response

// Returns information about a specific milestone request.
// It maps to the `GET` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Get(ctx context.Context, milestoneRequestID int, opts *rr.GetMilestoneRequestOptions) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Creates a milestone request from a given JSON object.
// It maps to the `POST` `/projects/0.1/milestone_requests` endpoint
func (s *MilestoneRequestsService) Create(ctx context.Context, b rr.CreateMilestoneRequestBody) (*rr.RawResponse, error) {
	p := endpoints.ProjectsMilestoneRequests
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (s *MilestoneRequestsService) Action(ctx context.Context, milestoneRequestID int, b rr.ActionMilestoneRequestBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsMilestoneRequests, milestoneRequestID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// --------------------------------------
// PROJECTS-REVIEWS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of project reviews.
// It maps to the `GET` `/projects/0.1/reviews` endpoint
func (s *ReviewsService) List(ctx context.Context, opts *rr.ListReviewsOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsReviews
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Post a review of a user.
// It maps to the `POST` `/projects/0.1/reviews` endpoint
func (s *ReviewsService) Create(ctx context.Context, b rr.CreateReviewBody) (*rr.RawResponse, error) {
	p := endpoints.ProjectsReviews
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPost, p, nil, b)
}

// Performs an action on a review. Note that Reviews are uniquely identified by a combination of review id and review type.
// It maps to the `PUT` `/projects/0.1/reviews/{review_id}` endpoint
func (s *ReviewsService) Action(ctx context.Context, reviewID int64, b rr.ReviewActionBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsReviews, reviewID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// --------------------------------------
// PROJECTS-EXTRAS
// --------------------------------------

// TODO: refine with typed response

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (s *ExpertGuaranteesService) List(ctx context.Context, opts *rr.ListExpertGuaranteesOptions) (*rr.RawResponse, error) {
	p := endpoints.ProjectsExpertGuarantees
	q := query.Values(opts)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (s *ExpertGuaranteesService) Action(ctx context.Context, expertGuaranteesID int64, b rr.ExpertGuaranteesActionRequestBody) (*rr.RawResponse, error) {
	p := fmt.Sprintf("%s/%d", endpoints.ProjectsExpertGuarantees, expertGuaranteesID)
	return execute[*rr.RawResponse](ctx, s.client, http.MethodPut, p, nil, b)
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/currencies` endpoint
func (s *CurrenciesService) List(ctx context.Context, opts *rr.ListCurrenciesOptions) (*rr.ListCurrenciesResponse, error) {
	p := endpoints.ProjectsCurrencies
	q := query.Values(opts)
	return execute[*rr.ListCurrenciesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
// it maps to the `GET` `/projects/0.1/categories` endpoint
func (s *CategoriesService) List(ctx context.Context, opts *rr.ListCategoriesOptions) (*rr.ListCategoriesResponse, error) {
	p := endpoints.ProjectsCategories
	q := query.Values(opts)
	return execute[*rr.ListCategoriesResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
// It maps to the `GET` `/projects/0.1/budgets` endpoint
func (s *BudgetsService) List(ctx context.Context, opts *rr.ListBudgetsOptions) (*rr.ListBudgetsResponse, error) {
	p := endpoints.ProjectsBudgets
	q := query.Values(opts)
	return execute[*rr.ListBudgetsResponse](ctx, s.client, http.MethodGet, p, q, nil)
}

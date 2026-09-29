package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// PROJECTS
// --------------------------------------

// Create a new project
// It maps to the `POST` `/projects/0.1/projects` endpoint.
func (r *Projects) Create(
	ctx context.Context,
	b rr.CreateProjectBody,
) (*rr.CreateProjectResponse, *ResponseMeta, error) {
	return execute[*rr.CreateProjectResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Projects,
		nil,
		b,
	)
}

// Perform an action on a project
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Action(
	ctx context.Context,
	projectID int64,
	b rr.ActionProjectBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Project(projectID),
		nil,
		b,
	)
}

// Returns information about multiple projects. Will be ordered by descending submit date (newest-to-oldest).
// it maps to the `GET` `/projects/0.1/projects` endpoint
func (r *Projects) List(
	ctx context.Context,
	opts *rr.ListProjectsOptions,
) (*rr.ListProjectsResponse, *ResponseMeta, error) {
	return execute[*rr.ListProjectsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Projects,
		opts,
		nil,
	)
}

// Returns the logged in user’s projects/contests they either created or participated in (by bidding or submitting an entry).
// it maps to the `GET` `/projects/0.1/self` endpoint
func (r *Projects) ListSelf(
	ctx context.Context,
	opts *rr.ListSelfProjectsOptions,
) (*rr.ListProjectsResponse, *ResponseMeta, error) {
	return execute[*rr.ListProjectsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsSelf,
		opts,
		nil,
	)
}

// Get information about a specific project. The full range of users projection options can be specified as part of this request by first setting theuser_detailsparameter to true.
// It maps to the `GET` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Get(
	ctx context.Context,
	projectID int64,
	opts *rr.GetProjectOptions,
) (*rr.GetProjectResponse, *ResponseMeta, error) {
	return execute[*rr.GetProjectResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Project(projectID),
		opts,
		nil,
	)
}

// Searches for active projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/active` endpoint
func (r *Projects) SearchActive(
	ctx context.Context,
	opts *rr.SearchActiveProjectsOptions,
) (*rr.ListProjectsResponse, *ResponseMeta, error) {
	return execute[*rr.ListProjectsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsActive,
		opts,
		nil,
	)
}

// Searches for all projects matching the desired query.
// It maps to the `GET` `/projects/0.1/projects/all` endpoint
func (r *Projects) SearchAll(
	ctx context.Context,
	opts *rr.SearchAllProjectsOptions,
) (*rr.ListProjectsResponse, *ResponseMeta, error) {
	return execute[*rr.ListProjectsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsAll,
		opts,
		nil,
	)
}

// Invites specific freelancers to bid on a project.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/invite` endpoint
func (r *Projects) InviteFreelancer(
	ctx context.Context,
	projectID int64,
	b rr.InviteFreelancersBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ProjectInvite(projectID),
		nil,
		b,
	)
}

// Returns the project upgrade fees for a given list of currencies. Also checks if the current user is eligible for free upgrades if requested.
// It maps to the `GET` `/projects/0.1/projects/fees` endpoint
func (r *Projects) ListUpgradesFees(
	ctx context.Context,
	opts *rr.ListUpgradesFeesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsFees,
		opts,
		nil,
	)
}

// Returns bids for a single project. Employers will see bids in a sorted order, which begins with sponsored bids, then by bid ranking. Freelancers will receive the bid list ordered by date. Note: This method is expensive to compute so it is recommended that reputation and user projection options are not set.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids` endpoint
func (r *Projects) ListBids(
	ctx context.Context,
	projectID int64,
	opts *rr.ListProjectBidsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectBids(projectID),
		opts,
		nil,
	)
}

// Returns information for posting bids on a project.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/bids_info` endpoint
func (r *Projects) GetBidInfo(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectBidsInfo(projectID),
		nil,
		nil,
	)
}

// Returns a list of milestones on a project. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/milestones` endpoint
func (r *Projects) ListMilestones(
	ctx context.Context,
	projectID int64,
	opts *rr.ListProjectMilestonesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectMilestones(projectID),
		opts,
		nil,
	)
}

// Returns a list of milestone requests by freelancers for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/milestone_requests` endpoint
func (r *Projects) ListMilestoneRequests(
	ctx context.Context,
	projectID int64,
	opts *rr.ListProjectsMilestoneRequestsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectMilestoneRequests(projectID),
		opts,
		nil,
	)
}

// Fetch the hourly contract matching the desired query.
// It maps to the `GET` `/projects/0.1/hourly_contract_info` endpoint
func (r *Projects) GetHourlyContractInfo(
	ctx context.Context,
	opts *rr.GetHourlyContractInfoOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.HourlyContractInfo,
		opts,
		nil,
	)
}

// Fetch the IP contract matching for the project id. If you are an employer it will return all of the contracts, ELSE we will return contract details specific to the logged-in user.
// It maps to the `GET` `/projects/0.1/projects/{project_id}/ip_contract_info` endpoint
func (r *Projects) GetIPContractInfo(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectIPContractInfo(projectID),
		nil,
		nil,
	)
}

// Delete a project by id. Only projects that are in pending or rejected states may be deleted. This will close the project and remove its visibility.
// it maps to the `DELETE` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Delete(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodDelete,
		endpoints.Project(projectID),
		nil,
		nil,
	)
}

// Returns a list of expert guarantees.
// It maps to the `GET` `/projects/0.1/expert_guarantees` endpoint
func (r *Projects) ListExpertGuarantees(
	ctx context.Context,
	opts *rr.ListExpertGuaranteesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ExpertGuarantees,
		opts,
		nil,
	)
}

// Perform an action on a expert guarantee.
// It maps to the `PUT` `/projects/0.1/expert_guarantees/{expert_guarantee_id}` endpoint
func (r *Projects) ActionExpertGuarantee(
	ctx context.Context,
	expertGuaranteesID int64,
	b rr.ActionExpertGuaranteesBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.ExpertGuarantee(expertGuaranteesID),
		nil,
		b,
	)
}

// --------------------------------------
// PROJECTS-COLLABORATIONS
// --------------------------------------

// Returns a list of project collaboration data for a project.
// it maps to the `GET` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (r *Collaborations) List(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectCollaborations(projectID),
		nil,
		nil,
	)
}

// Creates a new project collaboration.
// It maps to the `POST` `/projects/0.1/projects/{project_id}/collaborations` endpoint
func (r *Collaborations) Create(
	ctx context.Context,
	projectID int64, b rr.CreateCollaborationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ProjectCollaborations(projectID),
		nil,
		b,
	)
}

// Performs an action on a collaboration.
// it maps to the `PUT` `/projects/0.1/projects/{project_id}/collaborations/{collaboration_id}/actions` endpoint
func (r *Collaborations) Action(
	ctx context.Context,
	projectID int64,
	collaborationID int64,
	b rr.ActionCollaborationBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.ProjectCollaborationsActions(projectID, collaborationID),
		nil,
		b,
	)
}

// Returns a list of all collaboration data for a user.
// It maps toi the `GET` `/projects/0.1/projects/collaborations` endpoint
func (r *Collaborations) ListAll(
	ctx context.Context,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ProjectsCollaborations,
		nil,
		nil,
	)
}

// --------------------------------------
// PROJECTS-SERVICES
// --------------------------------------

// Orders one of the available services.
// It maps to the `POST` `/projects/0.1/services/{service_type}/{service_id}/order` endpoint
func (r *Services) Order(
	ctx context.Context,
	serviceID int64,
	serviceType rr.ServiceType,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.ServicesOrder(string(serviceType), serviceID),
		nil,
		nil,
	)
}

// Returns a list of services.
// it maps to the `GET` `/projects/0.1/services` endpoint
func (r *Services) List(
	ctx context.Context,
	opts *rr.ListServicesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Services,
		opts,
		nil,
	)
}

// Returns active services.
// it maps to the `GET` `/projects/0.1/services/active` endpoint
func (r *Services) SearchActive(
	ctx context.Context,
	opts *rr.SearchActiveServicesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.ServicesActive,
		opts,
		nil,
	)
}

// --------------------------------------
// PROJECTS-BIDS
// --------------------------------------

// Returns a list of bids that match the specified criteria.
// It maps to the `GET` `/projects/0.1/bids` endpoint
func (r *Bids) List(
	ctx context.Context,
	opts *rr.ListBidsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Bids,
		opts,
		nil,
	)
}

// Returns a list of bids that match the specified criteria.
// it maps to the `GET` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Get(
	ctx context.Context,
	bidID int64,
	opts *rr.GetBidOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Bid(bidID),
		opts,
		nil,
	)
}

// Creates a bid on a project. Accepts a JSON object in the style described in the Bid struct (with enums as strings, and objects as dictionaries).
// It maps to the `POST` `/projects/0.1/bids` endpoint
func (r *Bids) Create(
	ctx context.Context,
	b rr.CreateBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Bids,
		nil,
		b,
	)
}

// Performs an action on a bid.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Action(
	ctx context.Context,
	bidID int64,
	b rr.ActionBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Bid(bidID),
		nil,
		b,
	)
}

// Updates an existing bid on a project. An existing bids information (description,amount,milestone_percentage) can be updated by sending a JSON encoded Bid struct.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}` endpoint
func (r *Bids) Update(
	ctx context.Context,
	bidID int64,
	b rr.UpdateBidBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Bid(bidID),
		nil,
		b,
	)
}

// Returns a list of aggregate time tracking data for a bid.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (r *Bids) GetTimeTracking(
	ctx context.Context,
	bidID int64,
	opts *rr.GetTimeTrackingOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidTimeTracking(bidID),
		opts,
		nil,
	)
}

// Creates a time tracking session for a specific bid.
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/time_tracking` endpoint
func (r *Bids) CreateTimeTracking(
	ctx context.Context,
	bidID int64,
	b rr.CreateTimeTrackingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidTimeTracking(bidID),
		nil,
		b,
	)
}

// Return bid edit requests by bid id.
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/edit_requests` endpoint
// the original name was Get but it was renamed to List due to returning list of edit requests
func (r *Bids) ListEditRequests(
	ctx context.Context,
	bidID int64,
	opts *rr.ListBidEditRequestsOptions,
) (*rr.ListBidEditRequestsResponse, *ResponseMeta, error) {
	return execute[*rr.ListBidEditRequestsResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidEditRequests(bidID),
		opts,
		nil,
	)
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request.
// It maps to the `POST` `/projects/0.1/bids/edit_requests` endpoint
func (r *Bids) CreateEditRequest(
	ctx context.Context,
	b rr.CreateBidEditRequestBody,
) (*rr.CreateBidEditRequestResponse, *ResponseMeta, error) {
	return execute[*rr.CreateBidEditRequestResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidsEditRequests,
		nil,
		b,
	)
}

// Employer perform action on a PENDING bid edit request.
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/edit_requests/{edit_request_id}` endpoint
func (r *Bids) ActionEditRequest(
	ctx context.Context,
	bidID, bidEditRequestID int64,
	b rr.ActionBidEditRequestBody,
) (*rr.ActionBidEditRequestResponse, *ResponseMeta, error) {
	return execute[*rr.ActionBidEditRequestResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.BidEditRequest(bidID, bidEditRequestID),
		nil,
		b,
	)
}

// Fetch bid rating for a bid
// It maps to the `GET` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (r *Bids) GetRating(
	ctx context.Context,
	bidID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidsRatings(bidID),
		nil,
		nil,
	)
}

// Fetch bid ratings for multiple bids
// it maps to the `GET` `/projects/0.1/bid_ratings` endpoint
func (r *Bids) ListRatings(
	ctx context.Context,
	opts *rr.GetByListOfBidsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.BidRatings,
		opts,
		nil,
	)
}

// Rates a bid (creates a bid rating)
// It maps to the `POST` `/projects/0.1/bids/{bid_id}/bid_ratings` endpoint
func (r *Bids) CreateRating(
	ctx context.Context,
	bidID int64,
	b rr.CreateBidRatingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.BidsRatings(bidID),
		nil,
		b,
	)
}

// Updates an existing bid rating
// It maps to the `PUT` `/projects/0.1/bids/{bid_id}/bid_ratings/{bid_rating_id}` endpoint
func (r *Bids) UpdateRating(
	ctx context.Context,
	bidID int64,
	bidRatingID int64,
	b rr.UpdateBidRatingBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.BidRating(bidID, bidRatingID),
		nil,
		b,
	)
}

// --------------------------------------
// PROJECTS-JOBS
// --------------------------------------

// Returns a list of jobs.
// It maps to the `GET` `/projects/0.1/jobs` endpoint
func (r *Jobs) List(
	ctx context.Context,
	opts *rr.ListJobsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Jobs,
		opts,
		nil,
	)
}

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs.
// It maps to the `GET` `/projects/0.1/jobs/search` endpoint
func (r *Jobs) Search(
	ctx context.Context,
	opts *rr.SearchJobsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobsSearch,
		opts,
		nil,
	)
}

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name.
// It maps to the `GET` `/projects/0.1/job_bundles` endpoint
func (r *Jobs) ListBundles(
	ctx context.Context,
	opts *rr.ListJobBundlesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobBundles,
		opts,
		nil,
	)
}

// Returns a list of job bundle categories.
// It maps to the `GET` `/projects/0.1/job_bundle_categories` endpoint
func (r *Jobs) ListBundleCategories(
	ctx context.Context,
	opts *rr.ListJobBundleCategoriesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.JobBundleCategories,
		opts,
		nil,
	)
}

// --------------------------------------
// PROJECTS-MILESTONES
// --------------------------------------

// Returns a list of milestones. Does not return un-awarded prepaid milestones.
// It maps to the `GET` `/projects/0.1/milestones` endpoint
func (r *Milestones) List(
	ctx context.Context,
	opts *rr.ListMilestonesOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Milestones,
		opts,
		nil,
	)
}

// Returns information about a specific milestone.
// It maps to the `GET` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Get(
	ctx context.Context,
	milestoneID int64,
	opts *rr.GetMilestoneOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.Milestone(milestoneID),
		opts,
		nil,
	)
}

// Creates a milestone.
// It maps to the `POST` `/projects/0.1/milestones` endpoint
func (r *Milestones) Create(
	ctx context.Context,
	b rr.CreateMilestoneBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Milestones,
		nil,
		b,
	)
}

// Actions to be performed on a milestone.
// It maps to the `PUT` `/projects/0.1/milestones/{milestone_id}` endpoint
func (r *Milestones) Action(
	ctx context.Context,
	milestoneID int64,
	b rr.ActionMilestoneBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.Milestone(milestoneID),
		nil,
		b,
	)
}

// Returns a list of milestone requests.
// It maps to the `GET` `/projects/0.1/milestone_requests` endpoint
func (r *Milestones) ListRequests(
	ctx context.Context,
	opts *rr.ListMilestoneRequestsOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.MilestoneRequests,
		opts,
		nil,
	)
}

// Returns information about a specific milestone request.
// It maps to the `GET` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) GetRequest(
	ctx context.Context,
	milestoneRequestID int64,
	opts *rr.GetMilestoneRequestOptions,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodGet,
		endpoints.MilestoneRequest(milestoneRequestID),
		opts,
		nil,
	)
}

// Creates a milestone request from a given JSON object.
// It maps to the `POST` `/projects/0.1/milestone_requests` endpoint
func (r *Milestones) CreateRequest(
	ctx context.Context,
	b rr.CreateMilestoneRequestBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.MilestoneRequests,
		nil,
		b,
	)
}

// Perform an action on a milestone request.
// It maps to the `PUT` `/projects/0.1/milestone_requests/{milestone_request_id}` endpoint
func (r *Milestones) ActionRequest(
	ctx context.Context,
	milestoneRequestID int64,
	b rr.ActionMilestoneRequestBody,
) (*rr.RawResponse, *ResponseMeta, error) {
	return execute[*rr.RawResponse](
		ctx,
		r.client,
		http.MethodPut,
		endpoints.MilestoneRequest(milestoneRequestID),
		nil,
		b,
	)
}

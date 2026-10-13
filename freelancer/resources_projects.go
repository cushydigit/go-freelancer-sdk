package freelancer

import (
	"context"
	"net/http"

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
	rr "github.com/cushydigit/go-freelancer-sdk/freelancer/reqres"
)

// --------------------------------------
// Resource-Projects
// --------------------------------------

// Create a new project
// It maps to the `POST` `/projects/0.1/projects` endpoint.
func (r *Projects) Create(
	ctx context.Context,
	b rr.CreateProjectBody,
) (*rr.ProjectResponse, *ResponseMeta, error) {
	return execute[*rr.ProjectResponse](
		ctx,
		r.client,
		http.MethodPost,
		endpoints.Projects,
		nil,
		b,
	)
}

// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) SignNDA(
	ctx context.Context,
	projectID int64,
	fullName, address, city, state, phone, country string,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		projectID,
		rr.ActionProjectBody{
			Action:   rr.ProjectActionSignNDA,
			FullName: fullName,
			Address:  address,
			City:     city,
			State:    state,
			Phone:    phone,
			Country:  country,
		},
	)
}

// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Upgrade(
	ctx context.Context,
	projectID int64,
	upgrades []rr.ProjectUpgradeType,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		projectID,
		rr.ActionProjectBody{
			Action:   rr.ProjectActionUpgrade,
			Upgrades: upgrades,
		},
	)
}

// Update Project’s Description and Skill.
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Update(
	ctx context.Context,
	projectID int64,
	jobIDs []int64,
	description *string,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		projectID,
		rr.ActionProjectBody{
			Action:      rr.ProjectActionUpdate,
			Description: description,
			JonIDs:      jobIDs,
		},
	)
}

// Close a project which is still open for bidding. Freelancers cannot bid on a project which has been closed.
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Close(
	ctx context.Context,
	projectID int64,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		projectID,
		rr.ActionProjectBody{
			Action: rr.ProjectActionClose,
		},
	)
}

// End a project that has been awarded to, and accepted by a freelancer. Ending a project cancels the contract between the freelancer and employer.
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) End(
	ctx context.Context,
	projectID int64,
	bidID int64,
	status rr.ProjectEndStatusType,
) (*rr.RawResponse, *ResponseMeta, error) {
	return r.action(
		ctx,
		projectID,
		rr.ActionProjectBody{
			Action: rr.ProjectActionEnd,
			BidID:  bidID,
			Status: status,
		},
	)
}

// Perform an action on a project
// It maps to the `PUT` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) action(
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

// Get information about a specific project. The full range of users projection options can be specified as part of this request by first setting theuser_detailsparameter to true.
// It maps to the `GET` `/projects/0.1/projects/{project_id}` endpoint
func (r *Projects) Get(
	ctx context.Context,
	projectID int64,
	opts *rr.GetProjectOptions,
) (*rr.ProjectResponse, *ResponseMeta, error) {
	return execute[*rr.ProjectResponse](
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
func (r *Projects) ListUpgradeFees(
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

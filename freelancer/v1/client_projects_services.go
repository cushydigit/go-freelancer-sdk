package freelancer

// Return information about multiple projects. Will be ordered by descending submit data (newest-to-oldest)
func (c *Client) NewListProjectsService() *ListProjectsService {
	return &ListProjectsService{client: c}
}

// Get information about a specific project.
func (c *Client) NewGetProjectService() *GetProjectService {
	return &GetProjectService{client: c}
}

// Creates a new project
func (c *Client) NewCreateProjectService() *CreateProjectService {
	return &CreateProjectService{client: c}
}

// Delete a project by id, Only projects that are in pending or rejected state may be deleted.
func (c *Client) NewDeleteProjectService() *DeleteProjectService {
	return &DeleteProjectService{client: c}
}

// Searches for active projects matching the desired query
func (c *Client) NewSearchActiveProjectsService() *SearchActiveProjectsService {
	return &SearchActiveProjectsService{client: c}
}

// Returns the logged in user's projects/contests they either created or participated in (by bidding or submiting an entry)
func (c *Client) NewListSelfProjectsService() *ListSelfProjectsService {
	return &ListSelfProjectsService{client: c}
}

// Invite specific freelancers to bid on a project
func (c *Client) NewInviteFreelancerService() *InviteFreelancersService {
	return &InviteFreelancersService{client: c}
}

// Returns the project fees for a given list of currencies
func (c *Client) NewListUpgradeFeesService() *ListUpgradeFeesService {
	return &ListUpgradeFeesService{client: c}
}

// Returns bids for a single project.
func (c *Client) NewListProjectBidsService() *ListProjectBidsService {
	return &ListProjectBidsService{client: c}
}

// Returns information for posting bids on project
func (c *Client) NewGetProjectBidInformationService() *GetProjectBidInformationService {
	return &GetProjectBidInformationService{client: c}
}

// Returns a list of milestones on a project
func (c *Client) NewListProjectMilestonesService() *ListProjectMilestonesService {
	return &ListProjectMilestonesService{client: c}
}

// Returns a list of milestone requests by freelancer for a project
func (c *Client) NewListProjectMilestoneRequestsService() *ListProjectMilestoneRequestsService {
	return &ListProjectMilestoneRequestsService{client: c}
}

// Fetch the hourly contract matching the desired query
func (c *Client) NewGetHourlyContractInformationService() *GetHourlyContractInformationService {
	return &GetHourlyContractInformationService{client: c}
}

// Fetch the IP contract matching for the project id
func (c *Client) NewGetProjectIPContractInformationService() *GetProjectIPContractInformationService {
	return &GetProjectIPContractInformationService{client: c}
}

// Returns a list of project collaboration data for a project
func (c *Client) NewListProjectCollaborationsService() *ListProjectCollaborationsService {
	return &ListProjectCollaborationsService{client: c}
}

// Creates a new project collaborations
func (c *Client) NewCreateProjectCollaborationService() *CreateProjectCollaborationService {
	return &CreateProjectCollaborationService{client: c}
}

// Performs and action on a collaboration
func (c *Client) NewUpdateProjectCollaborationService() *UpdateProjectCollaborationService {
	return &UpdateProjectCollaborationService{client: c}
}

// Returns a list of all collaboration data for a user
func (c *Client) NewListAllProjectsCollaborationsService() *ListAllProjectsCollaborationsService {
	return &ListAllProjectsCollaborationsService{client: c}
}

// Orders one of the available services
func (c *Client) NewOrderServiceService() *OrderServiceService {
	return &OrderServiceService{client: c}
}

// Returns a list of services
func (c *Client) NewListProjectsServicesService() *ListProjectsServicesService {
	return &ListProjectsServicesService{client: c}
}

// Returns active services
func (c *Client) NewListProjectsActiveServicesService() *ListProjectsActiveServicesService {
	return &ListProjectsActiveServicesService{client: c}
}

// Returns a list of bids that match the specified criteria
func (c *Client) NewListBidsService() *ListBidsService {
	return &ListBidsService{client: c}
}

// Returns information about a specific bid
func (c *Client) NewGetBidService() *getBidService {
	return &getBidService{client: c}
}

// Creates a bid on a project
func (c *Client) NewCreateBidService() *createBidService {
	return &createBidService{client: c}
}

// Update an existing bid on project
func (c *Client) NewUpdateBidService() *updateBidService {
	return &updateBidService{client: c}
}

// Performs an action on a bid
func (c *Client) NewActionBidService() *actionBidService {
	return &actionBidService{client: c}
}

// Returns a list of aggregate time tracking data for a bid
func (c *Client) NewGetBidTimeTrackingService() *getBidTimeTrackingService {
	return &getBidTimeTrackingService{client: c}
}

// Creates a time tracking session for a specific bid
func (c *Client) NewCreateTimeTrackingBidService() *createTimeTrackingBidService {
	return &createTimeTrackingBidService{client: c}
}

// Returns bid requests by bid id
func (c *Client) NewGetBidEditRequestService() *getBidEditRequestService {
	return &getBidEditRequestService{client: c}
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request
func (c *Client) NewCreateBidEditRequestService() *createBidEditRequestService {
	return &createBidEditRequestService{client: c}
}

// Employer perform action on a PENDING bid edit reqeust
func (c *Client) NewActionBidEditRequestService() *actionBidEditRequestService {
	return &actionBidEditRequestService{client: c}
}

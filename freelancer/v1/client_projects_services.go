package freelancer

// Return information about multiple projects. Will be ordered by descending submit data (newest-to-oldest)
func (c *Client) NewListProjectsService() *ListProjectsService {
	return &ListProjectsService{client: c}
}

// Get information about a specific project.
func (c *Client) NewGetProjectService() *GetProjectService {
	return &GetProjectService{client: c}
}

// Creats a new project
func (c *Client) NewCreateProjectService() *CreateProjectService {
	return &CreateProjectService{client: c}
}

// Delete a project by id, Only pojects that are in pending or rejected state may be deleted.
func (c *Client) NewDeleteProjectService() *DeleteProjectService {
	return &DeleteProjectService{client: c}
}

// Searchesfor active projects matching the desired query
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

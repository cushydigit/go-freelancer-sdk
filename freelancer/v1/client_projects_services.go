package freelancer

// Return information about multiple projects. Will be ordered by descending submit data (newest-to-oldest)
func (c *Client) NewProjectsListService() *projectsListService {
	return &projectsListService{client: c}
}

// Get information about a specific project.
func (c *Client) NewProjectsGetByIDService() *projectsGetByIDService {
	return &projectsGetByIDService{client: c}
}

// Creates a new project
func (c *Client) NewProjectsCreateService() *projectsCreateService {
	return &projectsCreateService{client: c}
}

// Returns a list of categories. If job_details is set, a map of category IDs to jobs in those categories.
func (c *Client) NewCategoriesListService() *categoriesListService {
	return &categoriesListService{client: c}
}

// Delete a project by id, Only projects that are in pending or rejected state may be deleted.
func (c *Client) NewProjectsDeleteService() *projectsDeleteService {
	return &projectsDeleteService{client: c}
}

// Searches for active projects matching the desired query
func (c *Client) NewProjectsSearchActiveService() *projectsSearchActiveService {
	return &projectsSearchActiveService{client: c}
}

// Returns the logged in user's projects/contests they either created or participated in (by bidding or submitting an entry)
func (c *Client) NewSelfProjectsListService() *selfProjectsListService {
	return &selfProjectsListService{client: c}
}

// Invite specific freelancers to bid on a project
func (c *Client) NewInviteFreelancerService() *inviteFreelancersService {
	return &inviteFreelancersService{client: c}
}

// Returns the project fees for a given list of currencies
func (c *Client) NewUpgradeFeesListService() *upgradeFeesListService {
	return &upgradeFeesListService{client: c}
}

// Returns bids for a single project.
func (c *Client) NewProjectBidsListService() *projectBidsListService {
	return &projectBidsListService{client: c}
}

// Returns information for posting bids on project
func (c *Client) NewprojectBidInfoGetService() *projectBidInfoGetService {
	return &projectBidInfoGetService{client: c}
}

// Returns a list of milestones on a project
func (c *Client) NewProjectMilestonesListService() *projectMilestonesListService {
	return &projectMilestonesListService{client: c}
}

// Returns a list of milestone requests by freelancer for a project
func (c *Client) NewProjectMilestoneRequestsListService() *projectMilestoneRequestsListService {
	return &projectMilestoneRequestsListService{client: c}
}

// Fetch the hourly contract matching the desired query
func (c *Client) NewHourlyContractInfoGetService() *hourlyContractInfoGetService {
	return &hourlyContractInfoGetService{client: c}
}

// Fetch the IP contract matching for the project id
func (c *Client) NewProjectIPContractInfoGetService() *projectIPContractInfoGetService {
	return &projectIPContractInfoGetService{client: c}
}

// Returns a list of project collaboration data for a project
func (c *Client) NewCollaborationsListService() *collaborationsListService {
	return &collaborationsListService{client: c}
}

// Creates a new project collaborations
func (c *Client) NewCollaborationsCreateService() *collaborationsCreateService {
	return &collaborationsCreateService{client: c}
}

// Performs and action on a collaboration
func (c *Client) NewCollaborationsActionsService() *collaborationsActionsService {
	return &collaborationsActionsService{client: c}
}

// Returns a list of all collaboration data for a user
func (c *Client) NewCollaborationsListAllService() *collaborationsListAllService {
	return &collaborationsListAllService{client: c}
}

// Orders one of the available services
func (c *Client) NewServicesOrderService() *servicesOrderService {
	return &servicesOrderService{client: c}
}

// Returns a list of services
func (c *Client) NewServicesListService() *servicesListService {
	return &servicesListService{client: c}
}

// Returns active services
func (c *Client) NewServicesSearchActiveService() *servicesSearchActiveService {
	return &servicesSearchActiveService{client: c}
}

// Returns a list of bids that match the specified criteria
func (c *Client) NewBidsListService() *bidsListService {
	return &bidsListService{client: c}
}

// Returns information about a specific bid
func (c *Client) NewBidsGetByIDService() *bidsGetByIDService {
	return &bidsGetByIDService{client: c}
}

// Creates a bid on a project
func (c *Client) NewBidsCreateService() *bidsCreateService {
	return &bidsCreateService{client: c}
}

// Update an existing bid on project
func (c *Client) NewBidsUpdateService() *bidsUpdateService {
	return &bidsUpdateService{client: c}
}

// Performs an action on a bid
func (c *Client) NewBidsActionService() *bidsActionService {
	return &bidsActionService{client: c}
}

// Returns a list of aggregate time tracking data for a bid
func (c *Client) NewBidTimeTrackingGetService() *bidTimeTrackingGetService {
	return &bidTimeTrackingGetService{client: c}
}

// Creates a time tracking session for a specific bid
func (c *Client) NewBidTimeTrackingCreateService() *bidTimeTrackingCreateService {
	return &bidTimeTrackingCreateService{client: c}
}

// Returns bid requests by bid id
func (c *Client) NewBidEditRequestsGetService() *bidEditRequestsGetService {
	return &bidEditRequestsGetService{client: c}
}

// Create a bid edit request on a post accept awarded bid. With no pending bid edit request
func (c *Client) NewBidEditRequestsCreateService() *bidEditRequestsCreateService {
	return &bidEditRequestsCreateService{client: c}
}

// Employer perform action on a PENDING bid edit request
func (c *Client) NewBidEditRequestsActionService() *bidEditRequestsActionService {
	return &bidEditRequestsActionService{client: c}
}

// Fetch bid rating for a bid
func (c *Client) NewBidsRatingsGetService() *bidsRatingsGetService {
	return &bidsRatingsGetService{client: c}
}

// Fetch bid ratings for multiple bids
func (c *Client) NewBidsRatingsListService() *bidsRatingsListService {
	return &bidsRatingsListService{client: c}
}

// Rate a bid (create a bid rating)
func (c *Client) NewBidRatingsCreateService() *bidRatingsCreateService {
	return &bidRatingsCreateService{client: c}
}

// Updates an existing bid rating
func (c *Client) NewBidsRatingsUpdateService() *bidsRatingsUpdateService {
	return &bidsRatingsUpdateService{client: c}
}

// Returns a list of jobs
func (c *Client) NewJobsListService() *jobsListService {
	return &jobsListService{client: c}
}

// Returns a list of jobs. Note: This performs a sub-string search for all the parameters specified on the jobs
func (c *Client) NewJobsSearchService() *JobsSearchService {
	return &JobsSearchService{client: c}
}

// Returns a list of job bundles. Note: Categories in this context are job bundle categories. These are not the same as job categories even though they share the same name
func (c *Client) NewJobBundlesListService() *jobBundlesListService {
	return &jobBundlesListService{client: c}
}

// Returns a list of job bundle categories
func (c *Client) NewJobBundleCategoriesListService() *jobBundleCategoriesListService {
	return &jobBundleCategoriesListService{client: c}
}

// Returns a list of milestones. Does not return un-awarded prepaid milestones
// Note that for the parameters, you cannot specify both users[]and bidders[]/project_owners[]
func (c *Client) NewMilestonesListService() *milestonesListService {
	return &milestonesListService{client: c}
}

// Returns information about a specific milestone
func (c *Client) NewMilestonesGetByIDService() *milestonesGetByIDService {
	return &milestonesGetByIDService{client: c}
}

// Actions to be performed on a milestone
func (c *Client) NewMilestonesCreateService() *milestonesCreateService {
	return &milestonesCreateService{client: c}
}

// Actions to be performed on a milestone
func (c *Client) NewMilestonesActionsService() *milestonesActionsService {
	return &milestonesActionsService{client: c}
}

// Returns a list of milestone requests
func (c *Client) NewMilestoneRequestsListService() *milestoneRequestsListService {
	return &milestoneRequestsListService{client: c}
}

// Returns information about a specific milestone request
func (c *Client) NewMilestonesRequestsGetByIDService() *milestonesRequestsGetByIDService {
	return &milestonesRequestsGetByIDService{client: c}
}

// Creates a milestone request from a given JSON object.
func (c *Client) NewMilestoneRequestsCreateService() *milestoneRequestsCreateService {
	return &milestoneRequestsCreateService{client: c}
}

// Perform an action on a milestone request
func (c *Client) NewMilestoneRequestsActionsService() *milestoneRequestsActionsService {
	return &milestoneRequestsActionsService{client: c}
}

// Returns a list of expert guarantees
func (c *Client) NewExpertGuaranteesListService() *expertGuaranteesListService {
	return &expertGuaranteesListService{client: c}
}

// Perform an action on a expert guarantee
func (c *Client) NewExpertGuaranteesActionsService() *expertGuaranteesActionsService {
	return &expertGuaranteesActionsService{client: c}
}

// Returns a list of budgets with the specified currencies.currency_codes and currency_ids are incompatible with each other.
func (c *Client) NewBudgetsListService() *budgetsListService {
	return &budgetsListService{client: c}
}

// Returns a list of currencies.currency_codes and currency_ids are incompatible with each other.
func (c *Client) NewCurrenciesListService() *currenciesListService {
	return &currenciesListService{client: c}
}

func (c *Client) NewBidUpgradeFeesListService() *bidUpgradeFeesListService {
	return &bidUpgradeFeesListService{client: c}
}

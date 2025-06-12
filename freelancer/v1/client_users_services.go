package freelancer

// Returns a list of users
func (c *Client) NewUsersListService() *usersListService {
	return &usersListService{client: c}
}

// Return a list of current user's recent logged in devices
func (c *Client) NewSelfDevicesService() *selfDevicesListService {
	return &selfDevicesListService{client: c}
}

// Return information about current user
func (c *Client) NewSelfInfoService() *selfInfoService {
	return &selfInfoService{client: c}
}

// Returns information about a specific user
func (c *Client) NewUsersGetByIDService() *usersGetByIDService {
	return &usersGetByIDService{client: c}
}

// Returns a list of Freelancers
func (c *Client) NewFreelancersSearchService() *freelancersSearchService {
	return &freelancersSearchService{client: c}
}

// Add a list of jobs to the job list of current User
func (c *Client) NewSelfJobsAddService() *selfJobsAddService {
	return &selfJobsAddService{client: c}
}

// Sets a list of jobs to the job list of the current user
func (c *Client) NewSelfJobsUpdateService() *selfJobsUpdateService {
	return &selfJobsUpdateService{client: c}
}

// Removes a list of job list of the current user
func (c *Client) NewSelfJobsDeleteService() *selfJobsDeleteService {
	return &selfJobsDeleteService{client: c}
}

// Gets the reputations for a list of users
func (c *Client) NewReputationListService() *reputationListService {
	return &reputationListService{client: c}
}

// Returns a list of portfolios of users
func (c *Client) NewPortfoliosListService() *portfoliosListService {
	return &portfoliosListService{client: c}
}

// Get profile(s)
func (c *Client) NewProfilesGetService() *profilesGetService {
	return &profilesGetService{client: c}
}

// Update a profile
func (c *Client) NewProfilesUpdateService() *profilesUpdateService {
	return &profilesUpdateService{client: c}
}

// Create a new profile for a user. Returns the created profile
func (c *Client) NewProfilesCreateService() *profilesCreateService {
	return &profilesCreateService{client: c}
}

// Returns a list of enterprises
func (c *Client) NewEnterprisesListService() *enterprisesListService {
	return &enterprisesListService{client: c}
}

// Report user violations
func (c *Client) NewViolationCreateService() *violationCreateService {
	return &violationCreateService{client: c}
}

// Returns a list of pools belonging to the current user
func (c *Client) NewPoolsListService() *poolsListService {
	return &poolsListService{client: c}
}

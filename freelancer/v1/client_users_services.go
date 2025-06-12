package freelancer

// Returns a list of users
func (c *Client) NewUsersListService() *usersListService {
	return &usersListService{client: c}
}

// Return a list of current user's recent logged in devices
func (c *Client) NewListUserLoginDevicesService() *ListSelfLoginDevicesService {
	return &ListSelfLoginDevicesService{client: c}
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
func (c *Client) NewAddSelfJobsService() *AddSelfJobsService {
	return &AddSelfJobsService{client: c}
}

// Sets a list of jobs to the job list of the current user
func (c *Client) NewUpdateSelfJobsService() *UpdateSelfJobsService {
	return &UpdateSelfJobsService{client: c}
}

// Removes a list of job list of the current user
func (c *Client) NewDeleteSelfJobsService() *DeleteSelfJobsService {
	return &DeleteSelfJobsService{client: c}
}

// Gets the reputations for a list of users
func (c *Client) NewReputationListService() *reputationListService {
	return &reputationListService{client: c}
}

// Returns a list of portfolios of users
func (c *Client) NewListUsersPortfoliosService() *ListUsersPortfoliosService {
	return &ListUsersPortfoliosService{client: c}
}

// Get profile(s)
func (c *Client) NewGetProfilesService() *GetProfilesService {
	return &GetProfilesService{client: c}
}

// Update a profile
func (c *Client) NewUpdateProfilesService() *UpdateProfilesService {
	return &UpdateProfilesService{client: c}
}

// Create a new profile for a user. Returns the created profile
func (c *Client) NewCreateProfilesService() *CreateProfilesService {
	return &CreateProfilesService{client: c}
}

// Returns a list of enterprises
func (c *Client) NewUsersEnterprisesService() *ListEnterprisesService {
	return &ListEnterprisesService{client: c}
}

// Report user violations
func (c *Client) NewReportUserViolationService() *ReportUserViolationService {
	return &ReportUserViolationService{client: c}
}

// Returns a list of pools belonging to the current user
func (c *Client) NewListPoolsService() *ListPoolsService {
	return &ListPoolsService{client: c}
}

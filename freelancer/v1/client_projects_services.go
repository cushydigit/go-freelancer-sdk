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
func (c *Client) NewSerachActiveProjectsService() *SearchActiveProjectsService {
	return &SearchActiveProjectsService{client: c}
}

// Returns the logged in user's projects/contests they either created or participated in (by bidding or submiting an entry)
func (c *Client) NewListSelfProjectsService() *ListSelfProjectsService {
	return &ListSelfProjectsService{client: c}
}

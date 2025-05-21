package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ProjectFrontendStatus string

const (
	ProjectFrontendStatusOpen           ProjectFrontendStatus = "open"
	ProjectFrontendStatusComplete       ProjectFrontendStatus = "complete"
	ProjectFrontendStatusPending        ProjectFrontendStatus = "pending"
	ProjectFrontendStatusDraft          ProjectFrontendStatus = "draft"
	ProjectFrontendStatusWorkInProgress ProjectFrontendStatus = "work_in_progress"
)

const (
	BaseAPIMainURL    = "https://www.freelancer.com/api"
	BaseAPISandBoxURL = "https://api-sandbox.freelancer.com"
)

func (c *Client) GetBaseUrl() string {
	return c.baseURL
}
func (c *Client) SetBaseUrl(url string) {
	c.baseURL = url
}
func (c *Client) SetUseRateLimit(enabled bool) {
	c.useRateLimit = enabled
}

type Client struct {
	logger       *log.Logger
	HTTPClient   *http.Client
	rateLimiter  *RateLimiter
	apiToken     string
	baseURL      string
	useRateLimit bool
	Debug        bool
}

func NewClient(apiToken string) *Client {
	return &Client{
		logger:       log.Default(),
		HTTPClient:   &http.Client{},
		apiToken:     apiToken,
		baseURL:      BaseAPIMainURL,
		Debug:        false,
		useRateLimit: true,
		rateLimiter:  NewRateLimiter(),
	}
}

func (c *Client) debug(format string, v ...any) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request) (err error) {

	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.baseURL, r.endpoint)
	queryString := r.query.Encode()
	header := http.Header{}
	header.Set("freelancer-oauth-v1", c.apiToken)
	header.Set("Content-Type", "application/json")
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s\n", fullURL)

	r.fullURL = fullURL
	r.header = header
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request) (data []byte, err error) {
	c.debug("calling api endpoint: %s\n", r.fullURL)
	err = c.parseRequest(r)
	if err != nil {
		c.debug("failed to parse request: %s\n", err)
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		c.debug("failed to create http request: %s\n", err)
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("http request: %v\n", req)

	if c.useRateLimit {
		c.rateLimiter.WaitIfNeeded()
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.debug("failed to send http request: %s\n", err)
		return []byte{}, err
	}
	if c.useRateLimit {
		c.rateLimiter.UpdateFromHeader(res)
	}

	c.debug("http response: %s\n", res.Status)
	data, err = io.ReadAll(res.Body)
	if err != nil {
		c.debug("failed to read http response body: %s\n", err)
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if res.StatusCode >= http.StatusBadRequest {
		c.debug("received bad status code: %s\n", res.Status)
		apiErr := new(APIError2)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, apiErr
	}

	return data, nil
}

//### USERS ENDPOINT SEVICES ###

// Returns a list of users
func (c *Client) NewListUsersService() *ListUsersService {
	return &ListUsersService{client: c}
}

// Return a list of current user's recent logged in devices
func (c *Client) NewListUserLoginDevicesService() *ListSelfLoginDevicesService {
	return &ListSelfLoginDevicesService{client: c}
}

// Return information about current user
func (c *Client) NewGetSelfInfoService() *GetSelfInfoService {
	return &GetSelfInfoService{client: c}
}

// Returns information about a specific user
func (c *Client) NewGetUserService() *GetUserService {
	return &GetUserService{client: c}
}

// Returns a list of Freelancers
func (c *Client) NewSearchFreelancersService() *SearchFreelancersService {
	return &SearchFreelancersService{client: c}
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
func (c *Client) NewListUsersReputationsService() *ListUsersReputationsService {
	return &ListUsersReputationsService{clinet: c}
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

// Returns a list of poools belognig to the current user
func (c *Client) NewListPoolsService() *ListPoolsService {
	return &ListPoolsService{client: c}
}

//### PROJECTS ENDPOINT SEVICES ###

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

func (c *Client) NewListActiveProjectsService() *ListActiveProjectsService {
	return &ListActiveProjectsService{client: c}
}

func (c *Client) NewListCategoriesService() *ListCategoriesService {
	return &ListCategoriesService{client: c}
}

func (c *Client) NewListCurrenciesService() *ListCurrenciesService {
	return &ListCurrenciesService{client: c}
}

func (c *Client) NewListBudgetsService() *ListBudgetsService {
	return &ListBudgetsService{client: c}
}

func (c *Client) NewListTimezonesService() *ListTimezonesService {
	return &ListTimezonesService{client: c}
}

func (c *Client) NewListCountriesService() *ListCountriesService {
	return &ListCountriesService{client: c}
}

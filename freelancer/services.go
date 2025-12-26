package freelancer

type Services struct {
	client   *Client
	Projects *ProjectsService
	Users    *UsersService
	Common   *CommonService
}

func newServices(c *Client) *Services {

	s := &Services{client: c}
	// init projects services
	s.Projects = &ProjectsService{client: c}
	s.Projects.Collaborations = &CollaborationsService{client: c}
	s.Projects.Services = &ServicesService{client: c}
	s.Projects.Bids = &BidsService{client: c}
	s.Projects.Bids.EditRequests = &BidEditRequestService{client: c}
	s.Projects.Bids.Ratings = &BidRatingsService{client: c}
	s.Projects.Milestones = &MilestonesService{client: c}
	s.Projects.Milestones.Requests = &MilestoneRequestsService{client: c}
	s.Projects.Reviews = &ReviewsService{client: c}
	s.Projects.Jobs = &JobsService{client: c}
	s.Projects.Extras = &ProjectsExtrasService{client: c}
	s.Projects.Extras.Budgets = &BudgetsService{client: c}
	s.Projects.Extras.Categories = &CategoriesService{client: c}
	s.Projects.Extras.Currencies = &CurrenciesService{client: c}
	s.Projects.Extras.ExpertGuarantees = &ExpertGuaranteesService{client: c}

	// init users services
	s.Users = &UsersService{client: c}
	s.Users.Self = &SelfService{client: c}
	s.Users.SelfJob = &SelfJobsService{client: c}
	s.Users.Profiles = &ProfilesService{client: c}
	s.Users.Extras = &UsersExtrasService{client: c}

	// init common services
	s.Common = &CommonService{client: c}

	return s
}

type ProjectsService struct {
	client         *Client
	Collaborations *CollaborationsService
	Services       *ServicesService
	Bids           *BidsService
	Jobs           *JobsService
	Milestones     *MilestonesService
	Reviews        *ReviewsService
	Extras         *ProjectsExtrasService
}

type CollaborationsService struct{ client *Client }
type ServicesService struct{ client *Client }

type BidsService struct {
	client       *Client
	EditRequests *BidEditRequestService
	Ratings      *BidRatingsService
}

type BidEditRequestService struct{ client *Client }

type BidRatingsService struct{ client *Client }

type JobsService struct{ client *Client }

type MilestonesService struct {
	client   *Client
	Requests *MilestoneRequestsService
}

type ReviewsService struct{ client *Client }

type MilestoneRequestsService struct{ client *Client }

type ProjectsExtrasService struct {
	client           *Client
	ExpertGuarantees *ExpertGuaranteesService
	Budgets          *BudgetsService
	Currencies       *CurrenciesService
	Categories       *CategoriesService
}

type ExpertGuaranteesService struct{ client *Client }

type BudgetsService struct{ client *Client }

type CurrenciesService struct{ client *Client }

type CategoriesService struct{ client *Client }

type UsersService struct {
	client   *Client
	Self     *SelfService
	SelfJob  *SelfJobsService
	Profiles *ProfilesService
	Extras   *UsersExtrasService
}

type SelfService struct{ client *Client }

type SelfJobsService struct{ client *Client }

type ProfilesService struct{ client *Client }

type UsersExtrasService struct{ client *Client }

type CommonService struct{ client *Client }

// -------------------------------------------------------------------------------
// Projects - Projects
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Collaborations
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Services
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Bids
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Jobs
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Milestones
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Reviews
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Projects - Extras
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Common
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Users - Users
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Users - Authenticated (self)
// -------------------------------------------------------------------------------

// -------------------------------------------------------------------------------
// Users - Authenticated Jobs (self)
// -------------------------------------------------------------------------------

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service adds a list of jobs to the job list of current User
func (s *SelfJobsService) Add() *addJobs {
	return &addJobs{client: s.client}
}

// The returned service prepare a `PUT` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service sets a list of jobs to the job list of current User
func (s *SelfJobsService) Update() *updateJobs {
	return &updateJobs{client: s.client}
}

// The returned service prepare a `DELETE` request to the Users endpoint
// `/users/0.1/self/jobs`. The request is executed when Do is called.
//
// The service removes a list of jobs from the job list of current User
func (s *SelfJobsService) Delete() *deleteJobs {
	return &deleteJobs{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Profiles
// -------------------------------------------------------------------------------

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service creates a new profile for a user
func (s *ProfilesService) Create() *createProfile {
	return &createProfile{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service gets the profiles for a list of users
func (s *ProfilesService) Get() *getProfile {
	return &getProfile{client: s.client}
}

// The returned service prepare a `PUT` request to the Users endpoint
// `/users/0.1/profiles`. The request is executed when Do is called.
//
// The service updates a profile
func (s *ProfilesService) Update() *updateProfile {
	return &updateProfile{client: s.client}
}

// -------------------------------------------------------------------------------
// Users - Extras
// -------------------------------------------------------------------------------

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/reputations`. The request is executed when Do is called.
//
// The service gets the reputations for a list of users
func (s *UsersExtrasService) ListReputations() *listReputations {
	return &listReputations{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/portfolios`. The request is executed when Do is called.
//
// The service gets the portfolios for a list of users
func (s *UsersExtrasService) ListPortfolios() *listPortfolios {
	return &listPortfolios{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/enterprises`. The request is executed when Do is called.
//
// The service returns a list of enterprises
func (s *UsersExtrasService) ListEnterprises() *listEnterprises {
	return &listEnterprises{client: s.client}
}

// The returned service prepare a `POST` request to the Users endpoint
// `/users/0.1/violation_reports`. The request is executed when Do is called.
//
// The service creates a new violation
func (s *UsersExtrasService) CreateViolation() *createViolation {
	return &createViolation{client: s.client}
}

// The returned service prepare a `GET` request to the Users endpoint
// `/users/0.1/pools`. The request is executed when Do is called.
//
// The service returns a list of pools belonging to the current user
func (s *UsersExtrasService) ListPools() *listPools {
	return &listPools{client: s.client}
}

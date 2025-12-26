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

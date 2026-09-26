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
	s.Projects.Milestones = &MilestonesService{client: c}
	s.Projects.Reviews = &ReviewsService{client: c}
	s.Projects.Jobs = &JobsService{client: c}

	// init users services
	s.Users = &UsersService{client: c}
	s.Users.Self = &SelfService{client: c}
	s.Users.SelfJob = &SelfJobsService{client: c}
	s.Users.Profiles = &ProfilesService{client: c}
	s.Users.Reputations = &ReputationsService{client: c}
	s.Users.Enterprises = &EnterprisesService{client: c}
	s.Users.Portfolios = &PortfoliosService{client: c}
	s.Users.Violations = &ViolationsService{client: c}
	s.Users.Pools = &PoolsService{client: c}

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
}

type CollaborationsService struct{ client *Client }
type ServicesService struct{ client *Client }
type BidsService struct{ client *Client }
type JobsService struct{ client *Client }
type MilestonesService struct{ client *Client }
type ReviewsService struct{ client *Client }

type UsersService struct {
	client      *Client
	Self        *SelfService
	SelfJob     *SelfJobsService
	Profiles    *ProfilesService
	Reputations *ReputationsService
	Enterprises *EnterprisesService
	Portfolios  *PortfoliosService
	Violations  *ViolationsService
	Pools       *PoolsService
}

type SelfService struct{ client *Client }
type SelfJobsService struct{ client *Client }
type ProfilesService struct{ client *Client }
type ReputationsService struct{ client *Client }
type EnterprisesService struct{ client *Client }
type PortfoliosService struct{ client *Client }
type ViolationsService struct{ client *Client }
type PoolsService struct{ client *Client }

type CommonService struct{ client *Client }

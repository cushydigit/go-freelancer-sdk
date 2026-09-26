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
	client *Client
	Self   *SelfService
}

type SelfService struct{ client *Client }

type CommonService struct{ client *Client }
